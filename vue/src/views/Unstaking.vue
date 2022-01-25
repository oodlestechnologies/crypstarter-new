<template>
	<div>
		<div class="container-fluid">
			<h3 class="staking-text">UnStake Assets (Earn POS)</h3>
			<div class="form-div">
				<form @submit="postData" method="post">
					<label for="fname">Validator Address</label>
			  <input type="text" id="validatoraddress1" name="delegate_address" v-model="posts.delegate_address" placeholder="Add Validator address.." required>
			  <input type="text" id="validatoraddress2" disabled name="delegate_address" v-model="posts.delegate_address" placeholder="Add Validator address.." required>
					<label for="lname">Amount</label>
					<input type="number" id="amount1" name="amount" v-model="posts.amount" placeholder="Enter Amount.." required min="1" />
					<input type="number" id="amount2" disabled name="amount" v-model="posts.amount" placeholder="Enter Amount.." required min="1" />
					<!-- <label for="lname">From Address</label>
			  <input type="text" id="fromaddress" name="from_address" v-model="posts.from_address" placeholder="From address.." required> -->
					<!-- <label for="lname">Gas</label>
			  <input type="text" id="gas" name="gas" v-model="posts.gas" placeholder="Gas.." required>
			  <label for="lname">Gas Price</label> -->
					<!-- <input type="text" id="gasprice" name="gas_price" v-model="posts.gas_price" placeholder="Gas price.." required>       -->
					<input type="submit" value="Submit" id="bts1" />
					<input type="submit" value="Submit" disabled id="bts2" />
				</form>
			</div>
		</div>
	</div>
</template>
<style>
.staking-text {
	color: #ffffff;
	margin-bottom: 20px;
}
#bts1, #validatoraddress1, #amount1 {
	display: none !important;
}
#bts2, #validatoraddress2, #amount2 {
	opacity: 0.5;
	display: none !important;
	cursor: not-allowed;
}
input,
select {
	width: 100%;
	padding: 12px 20px;
	/* margin: 8px 0; */
	display: inline-block;
	border-radius: 10px;
	border: none;
	box-sizing: border-box;
	background-color: #34344b;
	color: #ffffff;
	font-size: 15px;
	margin-bottom: 25px;
	font-weight: 700;
}

input[type='submit'] {
	width: 100%;
	background-color: #fec007;
	color: #171622 !important;
	padding: 14px 20px;
	/* margin: 8px 0; */
	border: none;
	border-radius: 5px;
	cursor: pointer;
	font-size: 20px;
	font-weight: 600;
}
label {
	display: block;
	font-weight: 400 !important;
	margin-bottom: 5px;
}

.form-div {
	border-radius: 5px;
	background-color: #212130;
	padding: 50px;
}

label {
	color: #fff;
	font-weight: 600;
	font-size: 18px;
}

.my-costom-class-success {
	background: #fec007;
	color: white;
	font-size: 18px;
	width: 50%;
	border-radius: 5px;
	position: absolute;
	margin-left: 50%;
}
.mystyle#bts1,
.mystyle#bts2,
.mystyle#validatoraddress1,
.mystyle#validatoraddress2,
.mystyle#amount1,
.mystyle#amount2 {
	display: block !important;
}
.my-costom-class-error {
	background: red;
	color: white;
	font-size: 18px;
	width: 50%;
	border-radius: 5px;
	position: absolute;
	margin-left: 50%;
}
.sp-sidebar__footer {
	height: 50px;
	position: absolute;
	bottom: 20px;
	left: 20px;
}

.container-fluid {
	margin-right: 25px;
}
</style>

<script>
import axios from 'axios'
import VueAxios from 'vue-axios'

export default {
	name: 'UnStaking',
	data() {
		return {
			errors: [],
			posts: {
				delegate_address:'',
				amount: ''
				// from_address:'',
				// gas:'',
				// gas_price:'',
				// pwd:'aman6726'
			}
		}
	},
	methods: {
		postData(e) {
			var gridcell0 = document.getElementsByClassName('sp-text')
			// console.log( gridcell0[1].getAttribute('title') );
			let myAddress = gridcell0[1].getAttribute('title')
			e.preventDefault()
			if (myAddress == null) {
				this.$toast.open({
					message: 'Wallet is locked',
					type: 'error',
					position: 'top-right',
					duration: 3000,
					dismissible: true,
					className: ['my-costom-class-error']
				})
				return
			}

			axios
				.post('http://faucet.crypstarter.network/api/v1/unstalking', {
					delegate_address:this.posts.delegate_address,
					amount: this.posts.amount,
					from_address: myAddress
					// gas:this.posts.gas,
					// gas_price:this.posts.gas_price,
					// pwd:'aman6726'
				})
				.then((res) => {
					const mes = res.data.message
					const txhash = res.data.txhash
					const status = JSON.parse(res.data.status_code)
					if (status == '200') {
						this.$router.push('/')
						this.$toast.open({
							message: mes,
							type: 'success',
							position: 'top-right',
							duration: 3000,
							dismissible: true,
							className: ['my-costom-class-success']
						})
					} else {
						this.$toast.open({
							message: mes,
							type: 'error',
							position: 'top-right',
							duration: 3000,
							dismissible: true,
							className: ['my-costom-class-error']
						})
					}
				})
			// .catch((err)=>{
			//     this.$toast.open({
			//   message: "Error Show",
			//   type: "error",
			//   position: "top-right",
			//   duration: 5000,
			//   dismissible: true,
			// })

			// })
			e.preventDefault()
		}
	}
}
</script>
