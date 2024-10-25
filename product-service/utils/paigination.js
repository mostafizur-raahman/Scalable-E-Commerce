function paginate(data, page = 1, limit = 10) {
    const totalItems = data.length;
    const totalPages = Math.ceil(totalItems / limit);
    const startIndex = (page - 1) * limit;
    const paginatedData = data.slice(startIndex, startIndex + limit);

    return {
        data: paginatedData,
        pagination: {
            totalItems,
            page,
            totalPages,
            limit,
        },
    };
}

module.exports = paginate;
