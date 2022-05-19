<template>
  <div class="default">
    <div class="left">
      <Menu />
    </div>
    <div class="mid">
      <router-view />
    </div>
  </div>
</template>

<script>
import Menu from "@/components/Menu.vue";
export default {
  components: { Menu }
};
</script>

<style lang="postcss" scoped>
.default {
  display: flex;
  min-height: 99vh;
}

.left {
  flex: 1;
}

.mid {
  flex: 4;
  border: 1px solid #ccc;
}
</style>
