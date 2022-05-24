<template>
  <el-dialog
    class="enroll"
    title="Enroll"
    :visible.sync="dialogVisible"
    width="30%"
    :before-close="hide"
  >
    <el-form class="enroll-form" ref="form" :model="form" label-width="80px">
      <el-form-item prop="user" label="ID">
        <el-select
          v-model="form.user"
          placeholder="Please select identity"
          @change="onChange"
        >
          <el-option
            v-for="item in identities"
            :key="item.id"
            :label="item.id"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item prop="type" label="type">
        <el-select
          v-model="form.type"
          placeholder="Please select enrollment type"
        >
          <el-option
            v-for="item in typeOptions"
            :key="item"
            :label="item"
            :value="item"
          />
        </el-select>
      </el-form-item>
      <el-form-item prop="organization" label="Organ">
        <el-input
          readonly
          v-model="form.organization"
          placeholder="Please select organization"
        ></el-input>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button @click="hide">Cancel</el-button>
      <el-button type="primary" @click="onConfirm">Confirm</el-button>
    </span>
  </el-dialog>
</template>

<script>
export default {
  props: ["enroll"],
  data() {
    return {
      identities: [],
      states: [],
      stateMap: {},
      dialogVisible: false,
      typeOptions: ["x509", "idemix"],
      form: {
        user: "",
        type: "x509",
        organization: ""
      }
    };
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
    show() {
      Promise.all([this.fetchIdentities(), this.fetchStates()]).then(() => {
        this.dialogVisible = true;
      });
    },
    hide() {
      this.beforeClose();
      this.dialogVisible = false;
    },
    beforeClose() {
      this.$refs["form"].resetFields();
    },
    onConfirm() {
      if (this.form.user === "") {
        this.$message.warning("Please select identity.");
        return;
      }
      if (this.form.type === "") {
        this.$message.warning("Please select enrollment type.");
        return;
      }
      const count = this.stateMap[this.form.user];
      if (count && count > 0) {
        const content = `This identity already enrolled, after the operation, 
        the old one will be covered, please make a backup and continue! `;
        this.$confirm(content, "Warning")
          .then(() => {
            this.$emit("enroll", this.form);
          })
          .catch(() => {});
      } else {
        this.$emit("enroll", this.form);
      }
    },
    onChange() {
      const matchers = this.identities.filter(e => {
        return e.id === this.form.user;
      });
      if (matchers.length > 0) {
        this.form.organization = matchers[0].affiliation;
      }
    }
  }
};
</script>

<style lang="postcss" scoped>
.enroll-form .el-select {
  width: 220px;
}
.enroll-form .el-input {
  width: 220px;
}
</style>
