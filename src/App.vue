<template>
  <div id="app">
    <pre v-if="loading">Loading...</pre>
    <b-container fluid v-else>
      <b-row>
        <b-col sm="6">
          <b-form-group horizontal label="Buscar" class="mb-0">
            <b-input-group>
              <b-form-input v-model="filter" placeholder="Escribe para buscar"/>
              <b-btn :disabled="!filter" @click="filter = ''">Limpiar</b-btn>
            </b-input-group>
          </b-form-group>
        </b-col>
        <b-col sm="6">
          <b-button size="lg" class="buttonAlign" @click.stop="create()">Crear</b-button>
        </b-col>
      </b-row>
      <b-table :items="data" :fields="fields" :filter="filter">
        <template slot="opciones" slot-scope="row">
          <b-button size="sm" @click.stop="info(row.item)">Ver</b-button>
          <b-button size="sm" @click.stop="edit(row.item)">Editar</b-button>
          <b-button size="sm" @click.stop="eliminar(row.item)">Eliminar</b-button>
        </template>
      </b-table>
    </b-container>

    <b-modal ref="details" size="lg" :title="temp.name" :ok-only="true" title-tag="h3">
      <b-row>
        <b-col sm="6">
          <img :src="temp.img">
        </b-col>
        <b-col sm="6">
          <b-row>
            <pre>
              <h3>Ingredientes</h3>
              {{temp.ingredients}}

              <h3>Calorias: {{temp.calories}}</h3>
            </pre>
          </b-row>
        </b-col>
      </b-row>
      <b-container>
        <b-row>
          <pre>
            <h3>Instrucciones</h3>
            {{temp.instructions}}
          </pre>
        </b-row>
      </b-container>
    </b-modal>

    <b-modal ref="delete" size="lg" title="Confirmacion" title-tag="h3" @ok="confirmDelete">
      <pre>¿Está seguro que desea eliminar {{temp.name}}?</pre>
    </b-modal>
    <b-modal
      ref="editCreate"
      size="lg"
      :title="`${typeof temp.id === 'string' && temp.id.length > 0? 'Editar' : 'Crear'}`"
      title-tag="h3"
      @ok="confirmEditCreate"
    >
      <b-container fluid>
        <b-row v-for="(input, index) in inputs" :key="index">
          <b-col sm="3">
            <label :for="`type-${index}`">{{ input.label }}:</label>
          </b-col>
          <b-col sm="9">
            <b-input
              :id="`type-${index}`"
              :type="input.type"
              :state="input.state"
              v-model.trim="temp[index]"
              :placeholder="input.label"
              v-if="input.type !== 'text-area'"
            ></b-input>
            <b-textarea
              :id="`type-${index}`"
              :state="input.state"
              v-model="temp[index]"
              :placeholder="input.label"
              rows="7"
              v-else
            ></b-textarea>
            <b-form-text class="textError" v-if="input.state === false">Campo Requerido</b-form-text>
          </b-col>
        </b-row>
      </b-container>
    </b-modal>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "app",
  data() {
    return {
      apiURL: "http://localhost:8082/recipes",
      fields: [
        {
          key: "name",
          label: "Receta",
          sortable: true,
          sortDirection: 'desc'
        },
        {
          key: "calories",
          label: "Calorias",
          sortable: true
        },
        "opciones"
      ],
      data: [],
      temp: {},
      loading: true,
      inputs: {
        name: {
          label: "Nombre Receta",
          type: "text",
          state: null
        },
        calories: {
          label: "Calorias",
          type: "number",
          state: null
        },
        img: {
          label: "URL Imagen",
          type: "text",
          state: null
        },
        ingredients: {
          label: "Ingredientes",
          type: "text-area",
          state: null
        },
        instructions: {
          label: "Instrucciones",
          type: "text-area",
          state: null
        }
      },
      filter: null
    };
  },
  methods: {
    info(item) {
      this.$refs.details.show();
      this.resetChecks();
      this.temp = Object.assign({}, item);
    },
    edit(item) {
      this.$refs.editCreate.show();
      this.resetChecks();
      this.temp = Object.assign({}, item);
    },
    create(item) {
      this.$refs.editCreate.show();
      this.resetChecks();
      this.temp = {};
    },
    confirmEditCreate(evt) {
      evt.preventDefault();
      if (this.checkInputs()) {
        return;
      }
      this.data = [];
      this.loading = true;
      this.$refs.editCreate.hide();
      if (typeof this.temp.id === "string" && this.temp.id.length > 0) {
        axios
          .put(this.apiURL + "/" + this.temp.id, this.temp)
          .then(response => this.mount());
      } else {
        axios.post(this.apiURL, this.temp).then(response => this.mount());
      }
    },
    eliminar(item) {
      this.$refs.delete.show();
      this.temp = item;
    },
    confirmDelete() {
      this.data = [];
      this.loading = true;
      axios
        .delete(this.apiURL + "/" + this.temp.id)
        .then(response => this.mount());
    },
    mount() {
      this.loading = true;
      axios.get(this.apiURL).then(response => {
        this.data = response.data;
        this.loading = false;
      });
    },
    resetChecks() {
      for (let key in this.inputs) {
        let input = this.inputs[key];
        input.state = null;
      }
    },
    checkInputs() {
      let error = false;
      this.resetChecks();
      for (let key in this.inputs) {
        let input = this.inputs[key];
        if (input.type === "number") {
          this.temp[key] = parseFloat(this.temp[key]);
          if (typeof this.temp[key] === "number" && this.temp[key] > 0) {
            input.state = true;
          } else {
            input.state = false;
            error = true;
          }
        } else {
          if (typeof this.temp[key] === "string" && this.temp[key].length > 0) {
            input.state = true;
          } else {
            input.state = false;
            error = true;
          }
        }
      }
      return error;
    }
  },
  mounted() {
    this.mount();
  }
};
</script>

<style lang="scss">
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: right;
  margin-top: 15px;
}
pre {
  text-align: justify;
  white-space: pre-line;
}
img {
  width: 100%;
}
.textError {
  text-align: justify;
  color: red !important;
}
.buttonAlign {
  margin-bottom: 15px;
}
</style>
