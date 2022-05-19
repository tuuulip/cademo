<template>
  <div class="layout">
    <div class="layout-body">
      <div class="left">
        <Header />
        <Menu />
      </div>
      <div class="mid">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script>
import Menu from "@/components/Menu.vue";
import Header from "@/components/Header.vue";
export default {
  components: { Header, Menu }
};
</script>

<style lang="postcss" scoped>
.layout {
  min-height: 99vh;
}

.layout-body {
  display: flex;
}

.left {
  flex: 1;
  min-height: 100vh;
  background-color: #f3f6fb;
}

.mid {
  flex: 5;
  flex-shrink: 1;
}
</style>
