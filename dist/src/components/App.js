export default {
  name: 'App',
  template: `
    <div id="app" class="container">

      <div class="row">
        <router-link v-bind:to="'/'">
          <div class="one-half column jutsu-box ">
            <p class="jutsu">ENIGMA</p>
            <h3>Share a secret with the universe!</h3>
          </div>
          <div class="one-half column image-box">
            <img class="sealing-jutsu" src="/dist/static/images/logo.png" alt="Eight Trigrams Sealing Jutsu">
          </div>
        </router-link>
      </div>

      <div class="row">
        <div class="one-half column">
          <router-view></router-view>
        </div>
      </div>
    </div>
  `,
};
