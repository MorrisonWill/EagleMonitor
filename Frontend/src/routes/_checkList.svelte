<script>
  function initMultiselect() {
    checkboxStatusChange();

    document.addEventListener("click", function (evt) {
      var flyoutElement = document.getElementById("myMultiselect"),
        targetElement = evt.target; // clicked element

      do {
        if (targetElement == flyoutElement) {
          // This is a click inside. Do nothing, just return.
          //console.log('click inside');
          return;
        }

        // Go up the DOM
        targetElement = targetElement.parentNode;
      } while (targetElement);

      // This is a click outside.
      toggleCheckboxArea(true);
      //console.log('click outside');
    });
  }

  function checkboxStatusChange() {
    var multiselect = document.getElementById("mySelectLabel");
    var multiselectOption = multiselect.getElementsByTagName("option")[0];

    var values = [];
    var checkboxes = document.getElementById("mySelectOptions");
    var checkedCheckboxes = checkboxes.querySelectorAll(
      "input[type=checkbox]:checked"
    );

    for (const item of checkedCheckboxes) {
      var checkboxValue = item.getAttribute("value");
      values.push(checkboxValue);
    }

    var dropdownValue = "Nothing is selected";
    if (values.length > 0) {
      dropdownValue = values.join(", ");
    }

    multiselectOption.innerText = dropdownValue;
  }

  function toggleCheckboxArea(onlyHide = false) {
    var checkboxes = document.getElementById("mySelectOptions");
    var displayValue = checkboxes.style.display;

    if (displayValue != "block") {
      if (onlyHide == false) {
        checkboxes.style.display = "block";
      }
    } else {
      checkboxes.style.display = "none";
    }
    console.log("toggle checkbox area");
  }
</script>

<link
  href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css"
  rel="stylesheet"
/>

<svelte:window on:load={initMultiselect()} />

<div class="container-fluid">
  <div class="form-group col-sm-8">
    <label for="dur">Select School</label>
    <select id="dur" class="form-select">
      <option value="0" selected>None</option>
      <option value="12">MCAS</option>
      <option value="24">CSOM</option>
      <option value="36">School of Advancing Studies</option>
      <option value="48">School of Social Work</option>
      <option value="60">Nursing School</option>
      <option value="72">Law School</option>
    </select>
  </div>

  <div class="form-group col-sm-8">
    <label for="myMultiselect">Options</label>
    <div id="myMultiselect" class="multiselect">
      <div
        id="mySelectLabel"
        class="selectBox"
        on:click={() => toggleCheckboxArea()}
      >
        <select class="form-select">
          <option>Subject</option>
        </select>
        <div class="overSelect" />
      </div>
      <div id="mySelectOptions">
        <label for="one"
          ><input
            type="checkbox"
            id="one"
            on:change={() => checkboxStatusChange()}
            value="one"
          /> First checkbox</label
        >
        <label for="two"
          ><input
            type="checkbox"
            id="two"
            on:change={() => checkboxStatusChange()}
            value="two"
          /> Second checkbox</label
        >
        <label for="three"
          ><input
            type="checkbox"
            id="three"
            on:change={() => checkboxStatusChange()}
            value="three"
          /> Third checkbox</label
        >
        <label for="four"
          ><input
            type="checkbox"
            id="four"
            on:change={() => checkboxStatusChange()}
            value="four"
          /> Third checkbox</label
        >
        <label for="five"
          ><input
            type="checkbox"
            id="five"
            on:change={() => checkboxStatusChange()}
            value="five"
          /> First checkbox</label
        >
        <label for="six"
          ><input
            type="checkbox"
            id="six"
            on:change={() => checkboxStatusChange()}
            value="six"
          /> Second checkbox</label
        >
        <label for="seven"
          ><input
            type="checkbox"
            id="seven"
            on:change={() => checkboxStatusChange()}
            value="seven"
          /> Third checkbox</label
        >
        <label for="eight"
          ><input
            type="checkbox"
            id="eight"
            on:change={() => checkboxStatusChange()}
            value="eight"
          /> First checkbox</label
        >
        <label for="nine"
          ><input
            type="checkbox"
            id="nine"
            on:change={() => checkboxStatusChange()}
            value="nine"
          /> Second checkbox</label
        >
        <label for="ten"
          ><input
            type="checkbox"
            id="ten"
            on:change={() => checkboxStatusChange()}
            value="ten"
          /> Third checkbox</label
        >
      </div>
    </div>
  </div>
</div>

<style>
  .multiselect {
    width: 100%;
  }

  .selectBox {
    position: relative;
  }

  .selectBox select {
    width: 100%;
  }

  .overSelect {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
  }

  #mySelectOptions {
    display: none;
    border: 0.5px #7c7c7c solid;
    background-color: #ffffff;
    max-height: 150px;
    overflow-y: scroll;
  }

  #mySelectOptions label {
    display: block;
    font-weight: normal;
    display: block;
    white-space: nowrap;
    min-height: 1.2em;
    background-color: #ffffff00;
    padding: 0 2.25rem 0 0.75rem;
    /* padding: .375rem 2.25rem .375rem .75rem; */
  }

  #mySelectOptions label:hover {
    background-color: #1e90ff;
  }
</style>
