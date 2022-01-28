<template>
	<SpSidebar v-on:sidebar-open="sidebarOpen = !sidebarOpen">
		<template v-slot:default>
			<div class="custom-header">
				<img :src="image" />
			</div>

			<div class="bond"  >
				<router-link to="/">
					<img :src="dashboard" />
					<p>Dashboard</p>
				</router-link>
			</div>
			<!-- <div class="bond">
				<a href="#">
					<img :src="bond" />
					<p>Bond</p>
				</a>
			</div> -->
			<div class="bond">
				<router-link to="/staking">
					<img :src="staking" />
					<p>Staking</p>
				</router-link>
			</div>
			<div class="bond">
				<router-link to="/unstaking">
					<img :src="unstaking" />
					<p>UnStaking</p>
				</router-link>
			</div>
			<div class="bond">
				<router-link to="/Withdraw">
					<img :src="withdraw" />
					<p>Withdraw Reward</p>
				</router-link>
			</div>
			<!-- <SpLinkIcon link="/" text="Dashboard" icon="Dashboard" />
			<SpLinkIcon href="#" target="_blank" text="Bond (Buy CST change to Bond)" icon="Form"/>
			<SpLinkIcon link="/Staking" text="Staking" icon="Transactions" />
			<SpLinkIcon href="http://explorer.crypstarter.network/" target="_blank" text="Explorer" icon="Form" /> -->
		
			<hr style="height: 1px; border-width: 0; color: black; background-color: black" />
			<br />
			<div class="bond">
				<a href="http://explorer.crypstarter.network/" target="_blank">
					<img :src="explorer" />
					<p>Explorer</p>
				</a>
			</div>
			<!-- <div class="bond">
				<a href="#">
					<img :src="forum" />
					<p>Forum</p>
				</a>
			</div>
			<div class="bond">
				<a href="#">
					<img :src="governance" />
					<p>Governance</p>
				</a>
			</div> -->
			<div class="bond">
				<a href="https://docs.crypstarter.network/" target="_blank">
					<img :src="documents" />
					<p>Documents</p>
				</a>
			</div>

			<!-- <SpLinkIcon href="#" target="_blank" text="Forum" icon="Form" />
			<SpLinkIcon href="#" target="_blank" text="Governance" icon="Form"/>
			<SpLinkIcon
				href="https://docs.crypstarter.network/whitepaper/introduction/what-is-crypstarter"
				target="_blank"
				text="Documents"
				icon="Docs"
			/> -->
		</template>
		<template v-slot:footer>
			<!-- <SpStatusAPI :showText="sidebarOpen" />
			<SpStatusRPC :showText="sidebarOpen" />
			<SpStatusWS :showText="sidebarOpen" /> -->
			<div class="orugnsasd">
				<div class="social-media">
					<a href="https://twitter.com/cst_blockchain" target="_blank"><img :src="twitter" /></a>
					<a href="#"><img :src="medium" /></a>
					<a href="#"><img :src="discord" /></a>
					<a href="https://www.reddit.com/user/crypstarternetwork" target="_blank"><img :src="reddit" /></a>
					<a href="#"><img :src="telegram" /></a>
				</div>
				<!-- <div class="copyright">
					<p><strong style="display:block;" :showText="sidebarOpen">Crypstarter</strong> @2022. All Right Reserved.</p>

				</div> -->
			</div>
			<!-- <div class="sp-text">Build: v0.3.8</div> -->
		</template>
	</SpSidebar>
</template>

<style>
.copyright p {
	font-size: 12px;
	color: #171622;
}

.social-media {
	display: flex;
	justify-content: space-between;
	border: none;
	/* margin-bottom: -13px; */
	width: 85%;
	margin-top: 27px;
}

.social-media a {
	text-decoration: none;
	color: #171622;
}

.bond img {
	width: 21px;
	margin-bottom: 20px;
}

.bond a {
	text-decoration: none;
}

.bond p {
	font-size: 16px;
	margin-bottom: 30px;
	margin-top: -42px;
	margin-left: 35px;
	color: #171622;
	font-weight: 400;
}
</style>

<script>
export default {
	name: 'Sidebar',
	data() {
		var elements = document.getElementsByClassName("sp-sidebar__content");

			var myFunction = function() {
				// document.getElementsByClassName('sp-sidebar')[0].classList.add('sp-mob-opened')
				// that.sidebarOpen = false
				var link = document.getElementsByClassName('sp-hamburger')[0];
				link.click();

				
			};
			setTimeout(()=>{
				 elements[0].addEventListener('click', myFunction, false);
				
			},1500)

		
		return {
			sidebarOpen: true,
			image: require('/public/Wide_logo-black.png'),
			dashboard: require('/public/dashboard.svg'),
			bond: require('/public/bond.svg'),
			staking: require('/public/staking.svg'),
			unstaking: require('/public/Unstaking.svg'),
			withdraw: require('/public/withdraw.svg'),
			explorer: require('/public/explorer.svg'),
			forum: require('/public/forum.svg'),
			governance: require('/public/governance.svg'),
			documents: require('/public/documents.svg'),
			twitter: require('/public/twitter.svg'),
			medium: require('/public/medium.svg'),
			discord: require('/public/discord.svg'),
			reddit: require('/public/Reddit.svg'),
			telegram: require('/public/telegram.svg')
		}
	},

	computed: {
		hasWallet() {
			return this.$store.hasModule(['common', 'wallet'])
		}
	}
}
</script>
