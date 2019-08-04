export default {
  name: 'RevealSecret',
  data() {
    return {
      show: false,
      notFound: true,
      countdownToggle: true,
      slug: "",
      secretText: "",
      createdAt: 0,
      expiresAt: 0,
      remainingViews: 0,
      shareableUrl: "",
      nowEpoch: Math.trunc(Date.now() / 1000)
    }
  },
  computed: {
    countDownSec () {
      if (this.expiresAt == this.createdAt) {
        return "∞;" // infinity
      }
      let sec = Math.trunc(this.expiresAt - this.nowEpoch);
      return sec >= 0 ? sec : "Expired";
    },
    countDownMin () {
      if (this.expiresAt == this.createdAt) {
        return "∞;" // infinity
      }
      let min = Math.trunc((this.expiresAt - this.nowEpoch) / 60) % 60
      return min >= -1 ? min : "Expired";
    }
  },
  methods: {
    copy() {
      var copyText = document.getElementById("shareableUrl");

      /* Select the text field */
      copyText.select();

      /* Copy the text inside the text field */
      document.execCommand("copy");
    }
  },
  mounted() {
    // update the date for countdown to expiry
    window.setInterval(() => {
        this.nowEpoch = Math.trunc(Date.now() / 1000);
    },1000);

    this.slug = window.location.pathname.split("/secret/")[1]
    let request_url = '/api/v1/secret/' + this.slug

    axios
      .get(request_url)
      .then(res => {
        this.show = true;
        this.notFound = false;
        console.log("GET", res);
        this.secretText = res.data.secretText;
        this.createdAt = res.data.createdAt;
        this.expiresAt = res.data.expiresAt;
        this.remainingViews = res.data.remainingViews;
        this.shareableUrl = `${window.location.protocol}//${window.location.host}/secret/${res.data.hash}`;

        // Entire view has been rendered
        // Send a put request to backend to update the number of views
        if (this.remainingViews > 0) {

          axios
            .put(request_url)
            .then(res => {
              console.log("PUT", res);
              this.remainingViews = res.data.remainingViews;
            })
            .catch(err => {
              this.show = true;
              console.log(err)
            })
        }
      })
      .catch(err => {
        this.show = true;
        console.log(err)
      })

  },
  template: `
    <div id="reveal-secret" v-if="show">
      <div v-if="notFound">
        <h1>To infinity and beyond!</h1>
        <label class="info-label">404: secret not found or secret expired</label>
      </div>
      <div v-else>
        <label class="info-label">Your super secret message</label>
        <p class="secret-text">{{ secretText }}</p>

        <div class="row">
          <div class="six columns">
            <label class="info-text">{{ remainingViews }}</label>
            <label class="info-label">Views left</label>
          </div>

          <div class="six columns timer">
            <transition name="fade" mode="out-in">
              <div v-if="countdownToggle" @click="countdownToggle = !countdownToggle">
                <p class="info-text">{{ countDownMin }}</p>
                <p class="info-label">Minutes to expire</p>
              </div>
            </transition>
            <transition name="fade" mode="out-in">
              <div v-if="!countdownToggle" @click="countdownToggle = !countdownToggle">
                <p class="info-text">{{ countDownSec }}</p>
                <p class="info-label">Seconds to expire</p>
              </div>
            </transition>
          </div>
        </div>

        <div class="row share-link" v-if="shareableUrl">
          <label class="info-label">Shareable link</label>
          <div class="ten columns copy-row" @click="copy">
            <input id="shareableUrl" class="input u-full-width" type="text" v-model="shareableUrl" readonly="readonly">
          </div>
          <div class="one columns">
            <button @click="copy" class="button-enigma button-copy"><i class="fa fa-clipboard" aria-hidden="true"></i></button>
          </div>
        </div>
      </div>
    </div>
  `,
};
