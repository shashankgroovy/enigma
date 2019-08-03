export default {
  name: 'App',
  template: `
    <div id="app" class="container">
      <router-link v-bind:to="'/'">
        <img class="sealing-jutsu" src="/dist/static/images/enigma.png" alt="Eight Trigrams Sealing Jutsu">
      </router-link>
      <div class="row">
        <div class="one-half column">
          <router-view></router-view>
        </div>
      </div>
    </div>
  `,
};
