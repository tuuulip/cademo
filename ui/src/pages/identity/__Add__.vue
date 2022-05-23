<template>
  <el-dialog
    class="add"
    title="Add Identity"
    :visible.sync="dialogVisible"
    width="30%"
    :before-close="hide"
  >
    <el-form class="add-form" ref="form" :model="form" label-width="80px">
      <el-form-item prop="id" label="ID">
        <el-input v-model="form.id" placeholder="Please input id"></el-input>
      </el-form-item>
      <el-form-item prop="type" label="Type">
        <el-select v-model="form.type" placeholder="Please select type">
          <el-option
            v-for="item in typeOptions"
            :key="item"
            :label="item"
            :value="item"
          >
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item prop="affiliation" label="Affiliation">
        <el-input
          v-model="form.affiliation"
          placeholder="Please input affiliation"
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
  props: ["addIdentity"],
  data() {
    return {
      dialogVisible: false,
      form: {
        id: "",
        type: "",
        affiliation: ""
      },
      typeOptions: ["client", "orderer", "peer", "user"]
    };
  },
  methods: {
    show() {
      this.dialogVisible = true;
    },
    hide() {
      this.beforeClose();
      this.dialogVisible = false;
    },
    beforeClose() {
      this.$refs["form"].resetFields();
    },
    onConfirm() {
      this.$emit("addIdentity", this.form);
    }
  }
};
</script>

<style lang="postcss" scoped>
.add-form .el-select {
  width: 220px;
}
.add-form .el-input {
  width: 220px;
}
</style>
