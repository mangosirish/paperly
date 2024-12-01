let profileCounter = 0; // Contador de perfiles académicos

function addAcademicProfile() {
	profileCounter++;
	const profileContainer = document.getElementById("profileContainer");

	// Crear un nuevo contenedor de perfil académico
	const newProfile = document.createElement("div");
	newProfile.className = "profile-container";
	newProfile.id = `profile-${profileCounter}`;
	newProfile.style.marginBottom = "1em";

	// Contenido del nuevo perfil
	newProfile.innerHTML = `
	<br/>
	<h4>Perfil Académico Extra</h4>
	<input type="text" name="profiles[${profileCounter}][institution]" placeholder="Institución" required>
	<br/>

	<input type="text" name="profiles[${profileCounter}][faculty]" placeholder="Facultad" required>
	<br/>

	<input type="text" name="profiles[${profileCounter}][career]" placeholder="Carrera" required>
	<br/>

	<input type="text" name="profiles[${profileCounter}][specialization]" placeholder="Especialidad">
	<br/>

	<button type="button" onclick="removeAcademicProfile(${profileCounter})">Eliminar Perfil</button>
	<br/>
`;

	profileContainer.appendChild(newProfile);
}

function removeAcademicProfile(id) {
	const profile = document.getElementById(`profile-${id}`);
	if (profile) {
		profile.remove();
	}
}

function searchPeople(searchQuery) {
	this.debounceTimer = setTimeout(() => {
		apiRoute = searchQuery.trim() !== '' ? `/people/search/${searchQuery}` : '/people'
		fetch(apiRoute)
			.then(response => response.json())
			.then(data => {
				const ulElement = document.getElementById('search_results');
				ulElement.innerHTML = '';
				const defaultLi = document.createElement('li')
				defaultLi.textContent = 'Crear un nuevo perfil personal +'
				defaultLi.setAttribute('x-on:click', 'setCreationMode()')
				ulElement.appendChild(defaultLi)

				data.forEach(person => {
					const li = document.createElement('li');
					li.textContent = person.first_name;
					li.setAttribute('x-on:click', `setResultMode();
						first_name='${person.first_name}'; 
						second_name='${person.middle_name}'; 
						first_surname='${person.first_surname}'; 
						last_surname='${person.second_surname}'; 
						email_pers='${person.personal_email}'; 
						email_inst='${person.institutional_email}';
						person_id='${person.person_id}' `)
					ulElement.appendChild(li);
				});
				this.show_search_results = true;
			})
			.catch(error => {
				console.error('Error al hacer la solicitud:', error);
			});
	}, 500);
}
//Cerrar ventana de form, logica

function createFormHandler() {
	return {
		open_form_create: false,
		confirmClose() {
			if (confirm("¿Realmente desea salir sin guardar?")) {
				this.open_form_create = false;
			}
		}
	};
}

function restoreProfileCounter() {
	profileCounter = 0;
}

// Journals

let debounceTimer = null;

// Función para buscar autores
function searchAuthors(searchQuery) {
	debounceTimer = setTimeout(() => {
		let apiRoute = searchQuery.trim() !== '' ? `/people/search/${searchQuery}` : '/people';
		fetch(apiRoute)
			.then(response => response.json())
			.then(data => {
				const ulElement = document.getElementById('author_search_results');
				ulElement.innerHTML = '';

				data.forEach(author => {
					const li = document.createElement('li');
					li.textContent = `${author.first_name} ${author.first_surname}`;
					li.setAttribute('x-on:click', `setResultAuthor(${JSON.stringify(author)})`);
					ulElement.appendChild(li);
				});
			})
			.catch(error => {
				console.error('Error al hacer la solicitud:', error);
			});
	}, 500);
}

// Función para buscar journals
function searchJournals(searchQuery) {
	debounceTimer = setTimeout(() => {
		let apiRoute = '/journals/details';
		fetch(apiRoute)
			.then(response => response.json())
			.then(raw_data => {
				const data = searchQuery.trim()
					? raw_data.filter(item =>
						item.Ejemplar?.toLowerCase().includes(searchQuery.toLowerCase())
					)
					: raw_data;
					console.log(raw_data)
				const ulElement = document.getElementById('journal_search_results');
				ulElement.innerHTML = '';

				data.forEach(journal => {
					const li = document.createElement('li');
					li.textContent = `${journal.Ejemplar} (${journal.Nombre})`;
					li.setAttribute('x-on:click', `setResultModeJournal(${JSON.stringify(journal)})`);
					ulElement.appendChild(li);
				});
			})
			.catch(error => {
				console.error('Error al hacer la solicitud:', error);
			});
	}, 500);
}
