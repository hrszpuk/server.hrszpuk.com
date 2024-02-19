class ExpandableList extends HTMLElement {
    static get observedAttributes() {
        return ['endpoint'];
    }

    constructor() {
        super();
        this.attachShadow({ mode: 'open' });
        const template = document.getElementById('expandable-list-template');
        this.shadowRoot.appendChild(template.content.cloneNode(true));
        this.endpoint = "./api/v1/"
    }

    connectedCallback() {
        this.fetchData();
        this.shadowRoot.addEventListener('click', this.toggleDetails.bind(this));
    }

    async fetchData() {
        try {
            // Placeholder for fetching data
            const response = await fetch(this.endpoint);
            const items = await response.json();
            this.populateList(items);
        } catch (error) {
            console.error('Failed to fetch items:', error);
        }
    }

    populateList(items) {
        const ul = this.shadowRoot.querySelector('ul');
        // Clear default items
        ul.innerHTML = '';
        items.forEach(item => {
            const li = document.createElement('li');
            li.classList.add('list-item');
            li.textContent = item.title; // Assuming each item has a 'title'

            const details = document.createElement('div');
            details.classList.add('list-item-details');
            details.textContent = item.details; // Assuming each item has 'details'
            li.appendChild(details);

            ul.appendChild(li);
        });
    }

    toggleDetails(event) {
        if (event.target.classList.contains('list-item')) {
            const details = event.target.querySelector('.list-item-details');
            details.style.display = details.style.display === 'none' ? 'block' : 'none';
        }
    }

    attributeChangedCallback(name, oldValue, newValue) {
        console.log(`Attribute ${name} changed from ${oldValue} to ${newValue}`);
        if (name === 'endpoint') {
            this.endpoint = newValue
            this.fetchData()
        }
    }
}

customElements.define('ExpandableList', ExpandableList);
