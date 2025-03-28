import { useState } from 'react';
import './App.css';

function App() {
  const user = {
    "firstName": "Akash",
    "middleName": "Kumar",
    "lastName": "Sharma",
    "role": "Senior UX Designer",
    "employeeID": "KAI003",
    "performance": "83%",
    "dateOfJoining": "24th Oct, 2024",
    "email": "akash.s@gmail.com",
    "primaryContact": "98760043",
    "alternateContact": "98760043",
    "bloodGroup": "O+",
    "gender": "Male",
    "aadhaar": "1234-5678-9102",
    "aadhaarName": "John Dove",
    "permanenetAddress": {
      "line1": "123, Park street, Newtown, Ahemdabad",
      "line2": "123, Park street, Newtown, Ahemdabad",
      "city": "Ahmedabad",
      "state": "Gujarat",
      "zipcode": "123456",
    },
    "temporaryAddress": {
      "line1": "123, Park street, Newtown, Ahemdabad",
      "line2": "123, Park street, Newtown, Ahemdabad",
      "city": "Ahmedabad",
      "state": "Gujarat",
      "zipcode": "123456",
    },
    "pan": "AAAAA1111E",
    "bankDetails": {
      "bankName": "Axis Bank",
      "accountNumber": "9876543211000",
      "accountName": "John Dove",
      "ifsc": "ABC0000123",
      "branch": "Ahmedabad",

    }
  }

  const [detailType, setDetailType] = useState("personal-details")

  function handleRadioOnClick(detailType) {
    setDetailType(detailType)
  }

  return (
    <div className="page">
      <div className='breadcrumbs'>
        Home &gt; Employee profile
      </div>
      <div id='head-frame'>
        <img id="profile-image" src="Employee.png" alt="Photo" />
        <div id='emp-data'>
          <h3>{user.firstName + " " + user.lastName}</h3>
          <h4>{user.role}</h4>
          <h3>{user.employeeID}</h3>
          <h3>{user.performance}</h3>
        </div>
        <button id='head-frame-button'>Edit</button>
      </div>
      <div id='page-action'>
        <button>Overview  6</button>
        <button>Details  6</button>
        <button>Assets  6</button>
        <button>Flows  6</button>
        <button>Roles & Responsibilites  6</button>
        <button>Users  3</button>
      </div>
      <div>
        <div id='project-data'>
          <DetailBox label="First name" value={user.firstName} />
          <DetailBox label="Middle name" value={user.middleName} />
          <DetailBox label="Last name" value={user.lastName} />
          <DetailBox label="Role" value={user.lastName} />
          <DetailBox label="Date of Joining" value={user.dateOfJoining} />
          <DetailBox label="Email id" value={user.email} />
        </div>
        <div id='stepper'>
          <input type='radio' id='personal-details' name='stepper' value='personal-details' onClick={() => handleRadioOnClick("personal-details")}/>
          <label for='personal-details'>Personal details</label>
          <span className='arrow'>&rarr;</span>
          <input type='radio' id='academic-details' name='stepper' value='academic-details' onClick={() => handleRadioOnClick("academic-details")}/>
          <label for='academic-details'>Academic details</label>
          <span className='arrow'>&rarr;</span>
          <input type='radio' id='professional-details' name='stepper' value='professional-details' onClick={() => handleRadioOnClick("professional-details")}/>
          <label for='professional-details'>Professional details</label>
        </div>
        <OnboardingDetails user={user} detailType={detailType} />
      </div>
    </div>
  );
}

function OnboardingDetails({ user, detailType }) {
  if (detailType == "professional-details") {
    return <ProfessionalDetails user={user}/>
  } else if (detailType == "academic-details") {
    return <AcademicDetails user={user}/>
  } else {
    return <PersonalDetails user={user}/>
  }
}
function PersonalDetails({ user }) {
  return (
    <div id='onboarding-form'>
      <div id="gen-info" className='section'>
        <div className='section-title'>
          <p>General information</p>
          <p>Help us by providing your basic information like (Name, email id, contact details, blood group, gender )</p>
        </div>
        <div className='section-data'>
          <div className='detail-box-container'>
            <DetailBox label="Personal email id" value={user.email} classes="w-100" />
            <DetailBox label="First name" value={user.firstName} />
            <DetailBox label="Middle name" value={user.middleName} />
            <DetailBox label="Last name" value={user.lastName} />
            <DetailBox label="Primary contact number" value={user.primaryContact} classes="w-40" />
            <DetailBox label="Alternate contact number" value={user.alternateContact} classes="w-40" />
            <DetailBox label="Blood group" value={user.bloodGroup} classes="w-40" />
            <DetailBox label="Gender" value={user.gender} classes="w-40" />
          </div>
        </div>
      </div>
      <div id="docs" className='section color-dirtywhite'>
        <div className='section-title'>
          <p>Aadhar details</p>
          <p>Help us with your documents. (Aadhar details, address details.)</p>
        </div>
        <div className='section-data'>
          <div>
            <input className='file-input' type='file' id='degree-certificate' ></input>
            <DetailBox label="Aadhar number" value={user.aadhaar} classes="w-40" />
            <DetailBox label="Name as per aadhar card" value={user.aadhaarName} classes="w-40" />
          </div>
          <div className='detail-box-container'>
            <p>Permanent address</p>
            <DetailBox label="Address Line 1" value={user.permanenetAddress.line1} classes="w-100" />
            <DetailBox label="Address Line 2" value={user.permanenetAddress.line2} classes="w-100" />
            <DetailBox label="City" value={user.permanenetAddress.city} />
            <DetailBox label="State" value={user.permanenetAddress.state} />
            <DetailBox label="Zip code" value={user.permanenetAddress.zipcode} />
          </div>
          <div className='detail-box-container'>
            <p>Temporary address</p>
            <input type='checkbox' name='same-address'></input>
            <label for='same-address'> Same as permanent address</label>
            <DetailBox label="Address Line 1" value={user.temporaryAddress.line1} classes="w-100" />
            <DetailBox label="Address Line 2" value={user.temporaryAddress.line2} classes="w-100" />
            <DetailBox label="City" value={user.temporaryAddress.city} />
            <DetailBox label="State" value={user.temporaryAddress.state} />
            <DetailBox label="Zip code" value={user.temporaryAddress.zipcode} />
          </div>
        </div>
      </div>
      <div id="bank-details" className='section'>
        <div className='section-title'>
          <p>Bank details</p>
          <p>Help us with your bank account details. (Pan card, account number, name on account, IFSC code, branch name)</p>
        </div>
        <div className='section-data'>
          <div className='detail-box-container'>
            <input className='file-input' type='file' id='bank-detail-1' ></input>
            <DetailBox label="PAN number" value={user.pan} classes="w-100" />
            <p>Bank Account details</p>
            <input className='file-input' type='file' id='bank-detail-1' ></input>
            <DetailBox label="Bank name" value={user.bankDetails.bankName} classes="w-100" />
            <DetailBox label="Account number" value={user.bankDetails.accountNumber} classes="w-40" />
            <DetailBox label="Name on account" value={user.bankDetails.accountName} classes="w-40" />
            <DetailBox label="IFSC code" value={user.bankDetails.ifsc} classes="w-40" />
            <DetailBox label="Branch" value={user.bankDetails.branch} classes="w-40" />
          </div>
        </div>
      </div>
    </div>
  )
}

function AcademicDetails({ user }) {
  return (
    <div id='onboarding-form'>
    </div>
  )
}

function ProfessionalDetails({ user }) {
  return (
    <div id='onboarding-form'>
    </div>
  )
}

function DetailBox({ label, value, classes }) {
  let classnames = 'detail-box' + ' ' + classes
  return (
    <div className={classnames}>
      <p>{label}</p>
      <p>{value}</p>
    </div>
  )
}
export default App;
