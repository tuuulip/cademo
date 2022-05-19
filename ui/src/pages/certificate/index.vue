<template>
  <div class="cert">
    <div class="cert-header">
      <div class="cert-search">
        <el-input placeholder="Please input id" v-model="search.id">
          <template slot="prepend">Id</template>
        </el-input>
        <el-input placeholder="Please input serial" v-model="search.serial">
          <template slot="prepend">Serial</template>
        </el-input>
        <el-button @click="reset">Reset</el-button>
        <el-button type="primary" @click="fetchIdentities">Search</el-button>
      </div>
      <el-button type="primary">Create Certificate</el-button>
    </div>
    <div class="cert-body">
      <el-table :data="certificates" border height="80vh" class="cert-table">
        <el-table-column prop="id" label="id" width="100px" />
        <el-table-column prop="serialNumber" label="serialNumber" />
        <el-table-column prop="pem" label="pem" width="230px">
          <div slot-scope="{ row }" class="cert-column-attr">
            <span class="cert-attr">{{ row.pem }}</span>
            <el-popover trigger="hover" placement="top-start" width="600">
              <div class="cert-pem">{{ row.pem }}</div>
              <i
                slot="reference"
                class="iden-attr-icon el-icon-warning-outline"
              ></i>
            </el-popover>
          </div>
        </el-table-column>
        <el-table-column prop="notBefore" label="not before" />
        <el-table-column prop="notAfter" label="not after" />
        <el-table-column label="operation">
          <el-button type="text">Reenroll</el-button>
          <el-button type="text" class="cert-btn-del">Delete</el-button>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      certificates: [],
      search: {
        id: "",
        serial: ""
      }
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
        .post("/certificates", this.search)
        .then(res => {
          this.certificates = res.data;
        })
        .catch(() => {});
    },
    reset() {
      this.search = {};
      return this.fetchIdentities();
    },
    onDateChange() {
      console.log(this.dateRange);
    }
  }
};
</script>

<style lang="postcss" scoped>
.cert-header {
  width: 100%;
  height: 100px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  box-sizing: border-box;
}

.cert-body {
  padding: 0 20px;
  box-sizing: border-box;
}

.cert-title {
  text-align: left;
}

.cert-table {
  width: 99%;
}

.cert-attr {
  width: 200px;
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
}

.cert-attr-icon {
  font-size: 18px;
  cursor: pointer;
  color: #409eff;
}

.cert-column-attr {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.cert-pem {
  font-weight: 100;
  color: #909399;
  font-size: 14px;
  white-space: pre;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
}

.cert-search {
  display: flex;
  justify-content: flex-start;
}

.cert-search .el-input {
  margin-right: 10px;
}

.cert-btn-del {
  color: red;
}
</style>
