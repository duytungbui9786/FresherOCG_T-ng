<template>
  <form name="calculator">
    <table>
      <tr>
        <td colspan="4">
          <input type="text" name="display" :value="value" disabled />
        </td>
      </tr>
      <tr>
        <td>
          <input type="button" name="one" value="1" @click="addnumber(1)" />
        </td>
        <td>
          <input type="button" name="two" value="2" @click="addnumber(2)" />
        </td>
        <td>
          <input type="button" name="three" value="3" @click="addnumber(3)" />
        </td>
        <td>
          <input
            type="button"
            class="operator"
            name="plus"
            value="+"
            @click="addmethod(1)"
          />
        </td>
      </tr>
      <tr>
        <td>
          <input type="button" name="four" value="4" @click="addnumber(4)" />
        </td>
        <td>
          <input type="button" name="five" value="5" @click="addnumber(5)" />
        </td>
        <td>
          <input type="button" name="six" value="6" @click="addnumber(6)" />
        </td>
        <td>
          <input
            type="button"
            class="operator"
            name="minus"
            value="-"
            @click="addmethod(2)"
          />
        </td>
      </tr>
      <tr>
        <td>
          <input type="button" name="seven" value="7" @click="addnumber(7)" />
        </td>
        <td>
          <input type="button" name="eight" value="8" @click="addnumber(8)" />
        </td>
        <td>
          <input type="button" name="nine" value="9" @click="addnumber(9)" />
        </td>
        <td>
          <input
            type="button"
            class="operator"
            name="times"
            value="x"
            @click="addmethod(3)"
          />
        </td>
      </tr>
      <tr>
        <td>
          <input
            type="button"
            class="operator"
            id="clear"
            name="clear"
            value="C"
            @click="del"
          />
        </td>
        <td>
          <input
            type="button"
            class="operator"
            name="zero"
            value="0"
            onclick="calculator.display.value += '0'"
          />
        </td>
        <td>
          <input
            type="button"
            class="operator"
            name="div"
            value="/"
            @click="addmethod(4)"
          />
        </td>
        <td>
          <input
            type="button"
            class="operator"
            name="doit"
            value="="
            @click="abc"
          />
        </td>
      </tr>
    </table>
  </form>
</template>

<script>
export default {
  data() {
    return {
      value: "",
    };
  },
  methods: {
    del(){
      this.value=""
    },
    addnumber(int) {
      this.value += int;
    },
    addmethod(method) {
      if (method == 1) {
        this.value += "+";
      }
      if (method == 2) {
        this.value += "-";
      }
      if (method == 3) {
        this.value += "x";
      }
      if (method == 4) {
        this.value += "/";
      }
    },
    async abc() {
      if((this.value.includes("+"))==true){
         var sum = this.value.split("+")
       fetch("http://127.0.0.1:8090/sum?result="+sum[0]+","+sum[1])
       .then((response) => response.json())
       .then((json)=>this.value = json)
      .catch((error) => {
          console.error('Error:', error);
        });
      }
      else if((this.value.includes("-"))==true){
        
        
         var minus = this.value.split("-")
         if(minus[0]==""){
  fetch("http://127.0.0.1:8090/minus?result="+"-"+minus[1]+","+minus[2])
       .then((response) => response.json())
       .then((json)=>this.value = json)
       .catch((error) => {
          console.error('Error:', error);
        });
         }else{
             fetch("http://127.0.0.1:8090/minus?result="+minus[0]+","+minus[1])
       .then((response) => response.json())
       .then((json)=>this.value = json)
       .catch((error) => {
          console.error('Error:', error);
        });
         }
     
      }
      else if((this.value.includes("x"))==true){
         var multiply = this.value.split("x")
         console.log(multiply)
       fetch("http://127.0.0.1:8090/multiply?result="+multiply[0]+","+multiply[1])
       .then((response) => response.json())
       .then((json)=>this.value = json)
       .catch((error) => {
          console.error('Error:', error);
        });
      }
       else if((this.value.includes("/"))==true){
         var divide = this.value.split("/")
         console.log(multiply)
         if(divide[1]==0){
           this.value="infinity"
         }
       fetch("http://127.0.0.1:8090/divide?result="+divide[0]+","+divide[1])
       .then((response) => response.json())
       .then((json)=>this.value = json)
       .catch((error) => {
          console.error('Error:', error);
        });
      }
      
      //  fetch("http://127.0.0.1:8090/sum?result=2,3")
      //  .then((response) => response.json())
      //  .then((json)=>this.value = json)
       
    },
  },
  async created() {},
};
</script>

<style>
form {
  background: #ccc;
  margin: 0 auto;
  width: 300px;
  padding: 20px;
  border-radius: 5px;
}
form * {
  box-sizing: border-box;
}
form table {
  width: 100%;
  border-collapse: collapse;
}
form table tr,
form table td {
  border: 1px solid #eee;
}
form table tr:first-child input {
  height: 50px;
  padding: 0 10px;
  font-weight: 700;
}
form table tr:not(:first-child) td {
  height: 70px;
}
form table tr:not(:first-child) td input {
  height: 100%;
}
form table input {
  width: 100%;
  display: block;
  border: 0;
  font-size: 16px;
}
form table input:focus {
  outline: none;
}
form table input:hover {
  background: #eee;
  transition: all 0.2s ease-in-out;
}
form table input.operator {
  font-weight: 700;
  color: #fff;
  background: #ffcc00;
}
</style>
