<template>
  <div class="cred">
    <div class="cred-header"></div>
    <div class="cred-body">
      <el-table :data="credentials" border>
        <el-table-column prop="id" label="id"></el-table-column>
        <el-table-column prop="cred" label="cred">
          <div class="cred-cred-item" slot-scope="{ row }">
            <span class="cred-cred">{{ row.cred }}</span>
            <el-popover trigger="hover" placement="top-start" width="500">
              <div class="cred-cred__all">{{ row.cred }}</div>
              <i
                slot="reference"
                class="iden-attr-icon el-icon-warning-outline"
              ></i>
            </el-popover>
          </div>
        </el-table-column>
        <el-table-column prop="status" label="status"></el-table-column>
        <el-table-column prop="expiry" label="expiry"></el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      credentials: []
    };
  },
  created() {
    this.fetchCredentials();
  },
  methods: {
    fetchCredentials() {
      this.$request
        .get("/credential/list")
        .then(res => {
          this.credentials = res.data;
        })
        .catch(() => {});
    }
  }
};
</script>

<style lang="postcss" scoped>
.cred-header {
  width: 100%;
  height: 100px;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding: 0 20px;
  box-sizing: border-box;
}

.cred-body {
  padding: 0 20px;
  box-sizing: border-box;
}

.cred-cred {
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
}

.cred-cred-item {
  display: flex;
}

.cred-cred__all {
  font-weight: 100;
  color: #909399;
  font-size: 14px;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
}
</style>
