<template>
  <div class="iden">
    <el-table :data="identities">
      <el-table-column prop="id" label="id" />
      <el-table-column prop="type" label="type" />
      <el-table-column prop="affiliation" label="affiliation" />
      <el-table-column prop="max_enrollments" label="max_enrollments" />
      <el-table-column label="attrs" min-width="300px">
        <template slot-scope="{ row }">
          <div class="attr" v-for="(item, i) in row.attrs" :key="i">
            {{ item }}
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      identities: []
    };
  },
  created() {
    this.fetchIdentities();
  },
  methods: {
    fetchIdentities() {
      this.$request
        .get("/identities")
        .then(res => {
          this.identities = res.data;
        })
        .catch(() => {});
    }
  }
};
</script>
