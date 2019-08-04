export default {
  name: 'ShareSecret',
  data() {
    return {
      slug: "",
      secretText: "",
      expiresAt: 0,
      remainingViews: 0,
    }
  },
  methods: {
    handleSubmit: function() {
      const params = new URLSearchParams();
      params.append('secretText', this.secretText);
      params.append('expiresAt', this.expiresAt)
      params.append('remainingViews', this.remainingViews)

      axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
      axios.post('/api/v1/secret', params)
        .then(res => {
          this.slug = res.data.hash;
          this.$router.push({ name: 'reveal', params: { slug: this.slug }})
        }).catch(err => {
          this.show = true;
          console.log(err)
        })
    }
  },
  template: `
    <div id="share-secret">
        <form>
            <label for="textarea">Secret Message</label>
            <textarea v-model="secretText" id="textarea" class="u-full-width" placeholder="Type your secret message here" required></textarea>
            <div class="row">
                <div class="six columns">
                    <label for="remainingViews">Views Allowed</label>
                    <input v-model.number="remainingViews" class="input u-full-width" type="number" min="1" placeholder="Select number of views" id="remainingViews" oninput="validity.valid||(value='');">
                </div>
                <div class="six columns">
                    <label for="expiresAt">Expires in Minutes</label>
                    <input v-model.number="expiresAt" class="input u-full-width" type="number" min="0" placeholder="Expire after minutes" id="expiresAt" oninput="validity.valid||(value='');">
                </div>
            </div>
            <a @click="handleSubmit" class="button button-enigma">Share <i class="fa fa-paper-plane" aria-hidden="true"></i></a>
        </form>
    </div>
  `,
};
