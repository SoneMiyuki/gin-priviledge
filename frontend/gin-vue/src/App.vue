<template>
  <div id="app">
      <b-container>
        <navbar />
      </b-container>
    <router-view/>
  </div>
</template>

<script>
import Navbar from './views/layout/Navbar.vue';

export default {
  components: { Navbar },
  data() {
    return {};
  },
};
</script>

<style lang="scss">

</style>
