<template>
  <div class="menu">
    <el-menu :default-active="defaultActive" @select="onSelect">
      <el-menu-item index="identity">Identity</el-menu-item>
      <el-menu-item index="certificate">Certificate</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  data() {
    return {
      defaultActive: "1"
    };
  },
  methods: {
    onSelect(index) {
      switch (index) {
        case "identity":
          this.$router.push("/identity").catch(() => {});
          break;
        case "certificate":
          this.$router.push("/certificate").catch(() => {});
          break;
      }
    }
  }
};
</script>
