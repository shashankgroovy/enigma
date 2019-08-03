export default {
  name: 'ShareSecret',
  data() {
    return {
      hash: "",
      secretText: "",
      expiresAt: 0,
      remainingViews: 0,
    }
  },
  methods: {
    handleSubmit: function() {
      let formData = new FormData()
      console.log(formData)
      // Populate form data with the necessary fields.
      formData.set('secretText', this.secretText);
      formData.set('expiresAt', this.expiresAt);
      formData.set('remainingViews', this.remainingViews);
      console.log(formData)
      axios({
        method: 'post',
        url: '/api/v1/secret',
        data: formData,
        headers: {
          'content-type': `multipart/form-data; boundary=${formData._boundary}`,
        },
      }).then(res => {
        console.log(res)
        this.hash = res.data.hash;
        // this.$router.push()
      });
    }
  },
  template: `
    <div id="share-secret">
        <form>
            <label for="textarea">Secret Message</label>
            <textarea v-model="secretText" id="textarea" class="u-full-width" placeholder="Type your secret message here"></textarea>
            <div class="row">
                <div class="six columns">
                    <label for="remainingViews">Views Allowed</label>
                    <input v-model.number="remainingViews" class="input u-full-width" type="number" min="0" placeholder="Select number of views" id="remainingViews" oninput="validity.valid||(value='');">
                </div>
                <div class="six columns">
                    <label for="expiresAt">Expires in Minutes</label>
                    <input v-model.number="expiresAt" class="input u-full-width" type="number" min="0" placeholder="Expire after minutes" id="expiresAt" oninput="validity.valid||(value='');">
                </div>
            </div>
            <button @click="handleSubmit" class="button-enigma">Share <i class="fa fa-paper-plane" aria-hidden="true"></i></button>
        </form>
    </div>
  `,
};
