<template>
  <div class="iden">
    <div class="iden-header">
      <el-button type="primary" @click="showDialog">Add Identity</el-button>
    </div>
    <div class="iden-body">
      <el-table :data="identities" border height="80vh">
        <el-table-column prop="id" label="id" />
        <el-table-column label="state">
          <template slot-scope="{ row }">
            <span>{{ row.id | stateName(stateMap) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="type" label="type" />
        <el-table-column prop="affiliation" label="affiliation" />
        <el-table-column label="enrollments">
          <template slot-scope="{ row }">
            <span>{{ row.id | enrollentCount(stateMap) }}</span>
          </template>
        </el-table-column>
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
        <el-table-column label="operation" width="160">
          <template slot-scope="{ row }">
            <div>
              <el-button type="text">Edit</el-button>
              <el-button type="text" @click="revokeIdentity(row.id)"
                >Revoke</el-button
              >
              <el-button
                type="text"
                class="iden-btn-del"
                @click="deleteIdentity(row.id)"
                >Delete</el-button
              >
            </div>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- dialogs -->
    <Add ref="addDialog" @addIdentity="addIdentity" />
  </div>
</template>

<script>
import Add from "./__Add__.vue";
export default {
  components: { Add },
  data() {
    return {
      identities: [],
      states: [],
      stateMap: {}
    };
  },
  filters: {
    attrJson(obj) {
      return JSON.stringify(obj);
    },
    enrollentCount(id, stateMap) {
      const count = stateMap[id];
      if (count !== null) {
        return count;
      }
      return 0;
    },
    stateName(id, stateMap) {
      const state = stateMap[id];
      switch (state) {
        case "-1":
          return `Revoked`;
        default:
          return `Valid`;
      }
    }
  },
  created() {
    Promise.all([this.fetchIdentities(), this.fetchStates()]);
  },
  methods: {
    fetchIdentities() {
      this.$request
        .get("/id/all")
        .then(res => {
          this.identities = res.data;
        })
        .catch(() => {});
    },
    fetchStates() {
      this.$request
        .get("/id/state")
        .then(res => {
          this.states = res.data;
          this.updateStateMap();
        })
        .catch(() => {});
    },
    updateStateMap() {
      this.stateMap = {};
      this.states.forEach(element => {
        this.stateMap[element.id] = element.state;
      });
    },
    showDialog() {
      this.$refs["addDialog"].show();
    },
    hideDialog() {
      this.$refs["addDialog"].hide();
    },
    addIdentity(postData) {
      this.$request
        .post("/id/register", postData)
        .then(() => {
          this.hideDialog();
          return this.fetchIdentities();
        })
        .catch(() => {});
    },
    deleteIdentity(id) {
      this.$request
        .post("/id/del", { id })
        .then(() => {
          return this.fetchIdentities();
        })
        .catch(() => {});
    },
    revokeIdentity(id) {
      this.$request
        .post("/id/revoke", { id })
        .then(() => {
          return this.fetchIdentities();
        })
        .catch(() => {});
    }
  }
};
</script>

<style lang="postcss" scoped>
.iden-header {
  width: 100%;
  height: 100px;
  display: flex;
  align-items: center;
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

.iden-column-attr {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.iden-btn-del {
  color: red;
}
</style>
