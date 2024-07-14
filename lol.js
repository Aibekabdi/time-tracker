export const setupSearch = (data) => {
    const searchInput = document.getElementById('search');
    searchInput.addEventListener('input', event => {
        const searchTerm = String(event.target.value).toLowerCase();
        let hasOperator = searchTerm.split(':', 2)
        if (hasOperator.length == 1){
            filteredData = data.filter(hero => hero.name.toLowerCase().includes(searchTerm))
        }else {
            hasOperator[0] = hasOperator[0].trim()
            hasOperator[1] = hasOperator[1].trim()
            switch (hasOperator[0]){
                case 'include':
                    filteredData = data.filter(hero => hero.name.toLowerCase().includes(hasOperator[1]))
                    setupPagination(filteredData);
                    return;
                case 'exclude':
                    filteredData = data.filter(hero => !hero.name.toLowerCase().includes(hasOperator[1]))
                    setupPagination(filteredData);
                    return;
                case 'fuzzy':
                    filteredData = data
                    setupPagination(filteredData);
                    return;
                case 'equal':
                    if (hasOperator[1].includes("'") && hasOperator[1].length >= 3){
                        filteredData = data.filter(hero => hero.appearance.height[0] === hasOperator[1])
                    }else if (hasOperator[1].includes("cm") && hasOperator[1].split(/\s+/g).length == 2){
                        filteredData = data.filter(hero => hero.appearance.height[1] === hasOperator[1].replace(/\s+/g, ' '))
                    }else if (hasOperator[1].includes("lb") && hasOperator[1].split(/\s+/g).length == 2){
                        filteredData = data.filter(hero => hero.appearance.weight[0] === hasOperator[1].replace(/\s+/g, ' '))
                    }else if (hasOperator[1].includes("kg") && hasOperator[1].split(/\s+/g).length == 2){
                        filteredData = data.filter(hero => hero.appearance.weight[1] === hasOperator[1].replace(/\s+/g, ' '))
                    }
                    setupPagination(filteredData);
                    return;
                case 'not equal':
                    if (hasOperator[1].includes("'") && hasOperator[1].length == 3){
                        filteredData = data.filter(hero => hero.appearance.height[0] !== hasOperator[1])
                    }else if (hasOperator[1].includes("cm") && hasOperator[1].split(/\s+/g).length == 2){
                        filteredData = data.filter(hero => hero.appearance.height[1] !== hasOperator[1].replace(/\s+/g, ' '))
                    }else if (hasOperator[1].includes("meters") && hasOperator[1].split(/\s+/g).length == 2){
                        filteredData = data.filter(hero => hero.appearance.height[1] !== hasOperator[1].replace(/\s+/g, ' '))
                    }else if (hasOperator[1].includes("lb") && hasOperator[1].split(/\s+/g).length == 2){
                        filteredData = data.filter(hero => hero.appearance.weight[0] !== hasOperator[1].replace(/\s+/g, ' '))
                    }else if (hasOperator[1].includes("kg") && hasOperator[1].split(/\s+/g).length == 2){
                        filteredData = data.filter(hero => hero.appearance.weight[1] !== hasOperator[1].replace(/\s+/g, ' '))
                    }
                    setupPagination(filteredData);
                    return;
                case 'greater than':
                    if (hasOperator[1].includes("'") && hasOperator[1].length >= 3 && !Number.isNaN(parseFloat(hasOperator[1].replace("'", '.')))){
                        filteredData = data.filter(hero => isGreater(hasOperator[1], hero.appearance.height[0]))
                    }else if (hasOperator[1].includes("cm") && hasOperator[1].split(/\s+/g).length == 2){
                        // filteredData = data.filter(hero => isGreater(hasOperator[1].split(/\s+/g)[0], hero.appearance.height[1].split(/\s+/g)[0] ? hero.appearance.height[1].split(/\s+/g)[0] : '0'))
                        filteredData = data.filter(hero => hero.appearance.height.length > 1 && isGreater(hasOperator[1].split(/\s+/g)[0], hero.appearance.height[1].split(/\s+/g)[0]))
                    }
                    setupPagination(filteredData);
                    return;
                default:
                    filteredData = []
                    setupPagination(filteredData);
                    return;
            }
        };
        setupPagination(filteredData);
    });
};

function makeToNumHeight(arg){
    switch (arg){
        case arg.includes('cm'):
            return Number(arg)
        case arg.includes('meters'):
            return Number(arg) * 100
        case arg.includes("'"):
            return parseFloat(arg) * 30.48
        
    }
    return -1
}

function isGreater(srch, height){
    if (Number.isNaN(parseFloat(height.replace("'", '.')))){
        return false
    }
    if (parseFloat(srch.replace("'", '.')) < parseFloat(height.replace("'", '.'))){
        return true
    }
    return false
}
