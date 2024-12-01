<script lang="ts">
	import type { Lemma } from '$lib/types';
	import type { FormEventHandler } from 'svelte/elements';

	async function search(query: string): Promise<Lemma[]> {
		if (query === '') return [];
		const response = await fetch(`search?q=${query}`);
		return response.json();
	}

	let query = '';
	let result: Promise<Lemma[]>;
	let timer: number;
	let debounce: FormEventHandler<HTMLInputElement> = (e) => {
		const { value } = e.target as HTMLInputElement;
		clearTimeout(timer);
		timer = setTimeout(() => {
			query = value;
			result = search(value);
		}, 50);
	};
</script>

<div class="background flex flex-col px-2">
	<header class="flex flex-col w-full max-w-3xl gap-7 mx-auto mt-32 lg:mt-32 xl:mt-56 mb-14">
		<div class="mx-auto text-4xl font-bold">
			<h1>영한사전 / 한영사전</h1>
		</div>
		<div class="w-full max-w-3xl min-w-[200px]">
			<input
				class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-2xl border border-slate-200 rounded-full px-7 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
				placeholder="🔍 검색할 단어를 입력"
				on:input={debounce}
			/>
		</div>
	</header>
	<main class="flex flex-col mx-auto w-full">
		{#await result}
			<section class="flex mx-auto w-full max-w-3xl min-w-[200px]">
				<div role="status" class="m-auto">
					<svg
						aria-hidden="true"
						class="w-8 h-8 text-gray-300 animate-spin fill-sky-600"
						viewBox="0 0 100 101"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
					>
						<path
							d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
							fill="currentColor"
						/>
						<path
							d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
							fill="currentFill"
						/>
					</svg>
					<span class="sr-only">Loading...</span>
				</div>
			</section>
		{:then data}
			{#each data as lemma}
				<section class="mx-auto w-full max-w-3xl min-w-[200px]">
					<div class="flex gap-3">
						<h3 class="text-2xl">{lemma.written_form}</h3>
						<p class="text-base mt-auto">{lemma.part_of_speech}</p>
					</div>

					<div class="flex flex-col gap-3">
						{#each lemma.synsets as synset}
							<div>
								<i class="text-gray-500"> [{synset.words.join(', ')}] </i>
								<p>{synset.definitions}</p>
							</div>
						{/each}
						<hr class="h-px my-8 bg-gray-300 border-0" />
					</div>
				</section>
			{/each}
		{:catch error}
			<p>An error occurred!</p>
		{/await}
	</main>
</div>

<style>
	.background {
		background: #f2f4f7;
		width: 100vw;
		min-height: 100vh;
	}
</style>
