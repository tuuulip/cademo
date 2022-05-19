<template>
  <div class="iden">
    <div class="iden-header">
      <el-button type="primary">Create Identity</el-button>
    </div>
    <div class="iden-body">
      <h3 class="iden-title">Identities</h3>
      <el-table :data="identities" border>
        <el-table-column prop="id" label="id" />
        <el-table-column prop="type" label="type" />
        <el-table-column prop="affiliation" label="affiliation" />
        <el-table-column prop="max_enrollments" label="max_enrollments" />
        <el-table-column label="attrs" width="230px">
          <div slot-scope="{ row }" class="iden-column-attr">
            <span class="iden-attr">{{ row.attrs | attrJson }}</span>
            <el-popover trigger="hover" placement="top-start">
              <div>
                <div v-for="(item, i) in row.attrs" :key="i">{{ item }}</div>
              </div>
              <i
                slot="reference"
                class="iden-attr-icon el-icon-warning-outline"
              ></i>
            </el-popover>
          </div>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      identities: []
    };
  },
  filters: {
    attrJson(obj) {
      return JSON.stringify(obj);
    }
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

<style lang="postcss" scoped>
.iden-header {
  width: 100%;
  height: 10vh;
  display: flex;
  align-items: flex-end;
  justify-content: flex-end;
  padding: 0 20px;
  box-sizing: border-box;
}

.iden-body {
  padding: 0 20px;
  box-sizing: border-box;
}

.iden-title {
  text-align: left;
}

.iden-attr {
  width: 200px;
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
}

.iden-attr-icon {
  font-size: 18px;
  cursor: pointer;
  color: #409eff;
}

.iden-column-attr {
  display: flex;
  align-items: center;
  justify-content: space-between;
}
</style>
