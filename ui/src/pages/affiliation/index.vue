<template>
  <div class="aff">
    <div class="aff-header"></div>
    <div class="aff-body">
      affiliation
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      affiliationInfo: {}
    };
  },
  created() {
    this.getAllAffiliations();
  },
  methods: {
    getAllAffiliations() {
      this.$request.get("/affiliations").then(res => {
        this.affiliations = res.data;
      });
    }
  }
};
</script>

<style lang="postcss" scoped>
.aff {
  text-align: left;
}

.aff-header {
  height: 100px;
}

.aff-body {
  padding: 0 20px;
  text-align: left;
}
</style>
