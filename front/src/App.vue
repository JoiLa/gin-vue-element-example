<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style lang="scss">
  .set-language {
    color: #fff;
    position: absolute;
    top: 3px;
    font-size: 18px;
    right: 0px;
    cursor: pointer;
  }
</style>
