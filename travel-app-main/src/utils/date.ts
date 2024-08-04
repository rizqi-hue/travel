const months = ["Januari", "Februari", "Maret", "April", "Mei", "Juni",
    "Juli", "August", "September", "Oktober", "November", "Desember"];

export const toDate = (date: string) => {
    try {
        const t = new Date(date); // Epoch
        return t.getDate().toString() + ", " + months[t.getMonth()] + " " + t.getFullYear().toString();
    } catch (e) {
        return "Invalid date";
    }
};