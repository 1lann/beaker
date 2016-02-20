package element

type elementInfo struct {
	period    int
	group     int
	electrons []int
	category  string
	name      string
	symbol    string
	mass      float64
}

var allElements = []elementInfo{
	{
		period:    1,
		group:     1,
		electrons: []int{1},
		category:  "nonmetal",
		name:      "Hydrogen",
		symbol:    "H",
		mass:      1.00794,
	},
	{
		period:    1,
		group:     18,
		electrons: []int{2},
		category:  "noble",
		name:      "Helium",
		symbol:    "He",
		mass:      4.002602,
	},
	{
		period:    2,
		group:     1,
		electrons: []int{2, 1},
		category:  "alkali",
		name:      "Lithium",
		symbol:    "Li",
		mass:      6.941,
	},
	{
		period:    2,
		group:     2,
		electrons: []int{2, 2},
		category:  "alkaline",
		name:      "Beryllium",
		symbol:    "Be",
		mass:      9.012182,
	},
	{
		period:    2,
		group:     13,
		electrons: []int{2, 3},
		category:  "metalloid",
		name:      "Boron",
		symbol:    "B",
		mass:      10.811,
	},
	{
		period:    2,
		group:     14,
		electrons: []int{2, 4},
		category:  "nonmetal",
		name:      "Carbon",
		symbol:    "C",
		mass:      12.0107,
	},
	{
		period:    2,
		group:     15,
		electrons: []int{2, 5},
		category:  "nonmetal",
		name:      "Nitrogen",
		symbol:    "N",
		mass:      14.0067,
	},
	{
		period:    2,
		group:     16,
		electrons: []int{2, 6},
		category:  "nonmetal",
		name:      "Oxygen",
		symbol:    "O",
		mass:      15.9994,
	},
	{
		period:    2,
		group:     17,
		electrons: []int{2, 7},
		category:  "halogen",
		name:      "Fluorine",
		symbol:    "F",
		mass:      18.998404,
	},
	{
		period:    2,
		group:     18,
		electrons: []int{2, 8},
		category:  "noble",
		name:      "Neon",
		symbol:    "Ne",
		mass:      20.1797,
	},
	{
		period:    3,
		group:     1,
		electrons: []int{2, 8, 1},
		category:  "alkali",
		name:      "Sodium",
		symbol:    "Na",
		mass:      22.989769,
	},
	{
		period:    3,
		group:     2,
		electrons: []int{2, 8, 2},
		category:  "alkaline",
		name:      "Magnesium",
		symbol:    "Mg",
		mass:      24.305,
	},
	{
		period:    3,
		group:     13,
		electrons: []int{2, 8, 3},
		category:  "posttransition",
		name:      "Aluminium",
		symbol:    "Al",
		mass:      26.981539,
	},
	{
		period:    3,
		group:     14,
		electrons: []int{2, 8, 4},
		category:  "metalloid",
		name:      "Silicon",
		symbol:    "Si",
		mass:      28.0855,
	},
	{
		period:    3,
		group:     15,
		electrons: []int{2, 8, 5},
		category:  "nonmetal",
		name:      "Phosphorus",
		symbol:    "P",
		mass:      30.973763,
	},
	{
		period:    3,
		group:     16,
		electrons: []int{2, 8, 6},
		category:  "nonmetal",
		name:      "Sulfur",
		symbol:    "S",
		mass:      32.065,
	},
	{
		period:    3,
		group:     17,
		electrons: []int{2, 8, 7},
		category:  "halogen",
		name:      "Chlorine",
		symbol:    "Cl",
		mass:      35.453,
	},
	{
		period:    3,
		group:     18,
		electrons: []int{2, 8, 8},
		category:  "noble",
		name:      "Argon",
		symbol:    "Ar",
		mass:      39.948,
	},
	{
		period:    4,
		group:     1,
		electrons: []int{2, 8, 8, 1},
		category:  "alkali",
		name:      "Potassium",
		symbol:    "K",
		mass:      39.0983,
	},
	{
		period:    4,
		group:     2,
		electrons: []int{2, 8, 8, 2},
		category:  "alkaline",
		name:      "Calcium",
		symbol:    "Ca",
		mass:      40.078,
	},
	{
		period:    4,
		group:     3,
		electrons: []int{2, 8, 9, 2},
		category:  "transition",
		name:      "Scandium",
		symbol:    "Sc",
		mass:      44.955914,
	},
	{
		period:    4,
		group:     4,
		electrons: []int{2, 8, 10, 2},
		category:  "transition",
		name:      "Titanium",
		symbol:    "Ti",
		mass:      47.867,
	},
	{
		period:    4,
		group:     5,
		electrons: []int{2, 8, 11, 2},
		category:  "transition",
		name:      "Vanadium",
		symbol:    "V",
		mass:      50.9415,
	},
	{
		period:    4,
		group:     6,
		electrons: []int{2, 8, 13, 1},
		category:  "transition",
		name:      "Chromium",
		symbol:    "Cr",
		mass:      51.9961,
	},
	{
		period:    4,
		group:     7,
		electrons: []int{2, 8, 13, 2},
		category:  "transition",
		name:      "Manganese",
		symbol:    "Mn",
		mass:      54.938046,
	},
	{
		period:    4,
		group:     8,
		electrons: []int{2, 8, 14, 2},
		category:  "transition",
		name:      "Iron",
		symbol:    "Fe",
		mass:      55.845,
	},
	{
		period:    4,
		group:     9,
		electrons: []int{2, 8, 15, 2},
		category:  "transition",
		name:      "Cobalt",
		symbol:    "Co",
		mass:      58.933193,
	},
	{
		period:    4,
		group:     10,
		electrons: []int{2, 8, 16, 2},
		category:  "transition",
		name:      "Nickel",
		symbol:    "Ni",
		mass:      58.6934,
	},
	{
		period:    4,
		group:     11,
		electrons: []int{2, 8, 18, 1},
		category:  "transition",
		name:      "Copper",
		symbol:    "Cu",
		mass:      63.546,
	},
	{
		period:    4,
		group:     12,
		electrons: []int{2, 8, 18, 2},
		category:  "transition",
		name:      "Zinc",
		symbol:    "Zn",
		mass:      65.38,
	},
	{
		period:    4,
		group:     13,
		electrons: []int{2, 8, 18, 3},
		category:  "posttransition",
		name:      "Gallium",
		symbol:    "Ga",
		mass:      69.723,
	},
	{
		period:    4,
		group:     14,
		electrons: []int{2, 8, 18, 4},
		category:  "metalloid",
		name:      "Germanium",
		symbol:    "Ge",
		mass:      72.63,
	},
	{
		period:    4,
		group:     15,
		electrons: []int{2, 8, 18, 5},
		category:  "metalloid",
		name:      "Arsenic",
		symbol:    "As",
		mass:      74.9216,
	},
	{
		period:    4,
		group:     16,
		electrons: []int{2, 8, 18, 6},
		category:  "nonmetal",
		name:      "Selenium",
		symbol:    "Se",
		mass:      78.96,
	},
	{
		period:    4,
		group:     17,
		electrons: []int{2, 8, 18, 7},
		category:  "halogen",
		name:      "Bromine",
		symbol:    "Br",
		mass:      79.904,
	},
	{
		period:    4,
		group:     18,
		electrons: []int{2, 8, 18, 8},
		category:  "noble",
		name:      "Krypton",
		symbol:    "Kr",
		mass:      83.798,
	},
	{
		period:    5,
		group:     1,
		electrons: []int{2, 8, 18, 8, 1},
		category:  "alkali",
		name:      "Rubidium",
		symbol:    "Rb",
		mass:      85.4678,
	},
	{
		period:    5,
		group:     2,
		electrons: []int{2, 8, 18, 8, 2},
		category:  "alkaline",
		name:      "Strontium",
		symbol:    "Sr",
		mass:      87.62,
	},
	{
		period:    5,
		group:     3,
		electrons: []int{2, 8, 18, 9, 2},
		category:  "transition",
		name:      "Yttrium",
		symbol:    "Y",
		mass:      88.90585,
	},
	{
		period:    5,
		group:     4,
		electrons: []int{2, 8, 18, 10, 2},
		category:  "transition",
		name:      "Zirconium",
		symbol:    "Zr",
		mass:      91.224,
	},
	{
		period:    5,
		group:     5,
		electrons: []int{2, 8, 18, 12, 1},
		category:  "transition",
		name:      "Niobium",
		symbol:    "Nb",
		mass:      92.90638,
	},
	{
		period:    5,
		group:     6,
		electrons: []int{2, 8, 18, 13, 1},
		category:  "transition",
		name:      "Molybdenum",
		symbol:    "Mo",
		mass:      95.96,
	},
	{
		period:    5,
		group:     7,
		electrons: []int{2, 8, 18, 13, 2},
		category:  "transition",
		name:      "Technetium",
		symbol:    "Tc",
		mass:      98,
	},
	{
		period:    5,
		group:     8,
		electrons: []int{2, 8, 18, 15, 1},
		category:  "transition",
		name:      "Ruthenium",
		symbol:    "Ru",
		mass:      101.07,
	},
	{
		period:    5,
		group:     9,
		electrons: []int{2, 8, 18, 16, 1},
		category:  "transition",
		name:      "Rhodium",
		symbol:    "Rh",
		mass:      102.9055,
	},
	{
		period:    5,
		group:     10,
		electrons: []int{2, 8, 18, 18},
		category:  "transition",
		name:      "Palladium",
		symbol:    "Pd",
		mass:      106.42,
	},
	{
		period:    5,
		group:     11,
		electrons: []int{2, 8, 18, 18, 1},
		category:  "transition",
		name:      "Silver",
		symbol:    "Ag",
		mass:      107.8682,
	},
	{
		period:    5,
		group:     12,
		electrons: []int{2, 8, 18, 18, 2},
		category:  "transition",
		name:      "Cadmium",
		symbol:    "Cd",
		mass:      112.411,
	},
	{
		period:    5,
		group:     13,
		electrons: []int{2, 8, 18, 18, 3},
		category:  "posttransition",
		name:      "Indium",
		symbol:    "In",
		mass:      114.818,
	},
	{
		period:    5,
		group:     14,
		electrons: []int{2, 8, 18, 18, 4},
		category:  "posttransition",
		name:      "Tin",
		symbol:    "Sn",
		mass:      118.71,
	},
	{
		period:    5,
		group:     15,
		electrons: []int{2, 8, 18, 18, 5},
		category:  "metalloid",
		name:      "Antimony",
		symbol:    "Sb",
		mass:      121.76,
	},
	{
		period:    5,
		group:     16,
		electrons: []int{2, 8, 18, 18, 6},
		category:  "metalloid",
		name:      "Tellurium",
		symbol:    "Te",
		mass:      127.6,
	},
	{
		period:    5,
		group:     17,
		electrons: []int{2, 8, 18, 18, 7},
		category:  "halogen",
		name:      "Iodine",
		symbol:    "I",
		mass:      126.90447,
	},
	{
		period:    5,
		group:     18,
		electrons: []int{2, 8, 18, 18, 8},
		category:  "noble",
		name:      "Xenon",
		symbol:    "Xe",
		mass:      131.293,
	},
	{
		period:    6,
		group:     1,
		electrons: []int{2, 8, 18, 18, 8, 1},
		category:  "alkali",
		name:      "Caesium",
		symbol:    "Cs",
		mass:      132.90546,
	},
	{
		period:    6,
		group:     2,
		electrons: []int{2, 8, 18, 18, 8, 2},
		category:  "alkaline",
		name:      "Barium",
		symbol:    "Ba",
		mass:      137.327,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 18, 9, 2},
		category:  "lanthanoid",
		name:      "Lanthanum",
		symbol:    "La",
		mass:      138.90547,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 19, 9, 2},
		category:  "lanthanoid",
		name:      "Cerium",
		symbol:    "Ce",
		mass:      140.116,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 21, 8, 2},
		category:  "lanthanoid",
		name:      "Praseodymium",
		symbol:    "Pr",
		mass:      140.90765,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 22, 8, 2},
		category:  "lanthanoid",
		name:      "Neodymium",
		symbol:    "Nd",
		mass:      144.242,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 23, 8, 2},
		category:  "lanthanoid",
		name:      "Promethium",
		symbol:    "Pm",
		mass:      145,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 24, 8, 2},
		category:  "lanthanoid",
		name:      "Samarium",
		symbol:    "Sm",
		mass:      150.36,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 25, 8, 2},
		category:  "lanthanoid",
		name:      "Europium",
		symbol:    "Eu",
		mass:      151.964,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 25, 9, 2},
		category:  "lanthanoid",
		name:      "Gadolinium",
		symbol:    "Gd",
		mass:      157.25,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 27, 8, 2},
		category:  "lanthanoid",
		name:      "Terbium",
		symbol:    "Tb",
		mass:      158.92535,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 28, 8, 2},
		category:  "lanthanoid",
		name:      "Dysprosium",
		symbol:    "Dy",
		mass:      162.5,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 29, 8, 2},
		category:  "lanthanoid",
		name:      "Holmium",
		symbol:    "Ho",
		mass:      164.93031,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 30, 8, 2},
		category:  "lanthanoid",
		name:      "Erbium",
		symbol:    "Er",
		mass:      167.259,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 31, 8, 2},
		category:  "lanthanoid",
		name:      "Thulium",
		symbol:    "Tm",
		mass:      168.9342,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 8, 2},
		category:  "lanthanoid",
		name:      "Ytterbium",
		symbol:    "Yb",
		mass:      173.054,
	},
	{
		period:    6,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 9, 2},
		category:  "lanthanoid",
		name:      "Lutetium",
		symbol:    "Lu",
		mass:      174.9668,
	},
	{
		period:    6,
		group:     4,
		electrons: []int{2, 8, 18, 32, 10, 2},
		category:  "transition",
		name:      "Hafnium",
		symbol:    "Hf",
		mass:      178.49,
	},
	{
		period:    6,
		group:     5,
		electrons: []int{2, 8, 18, 32, 11, 2},
		category:  "transition",
		name:      "Tantalum",
		symbol:    "Ta",
		mass:      180.94788,
	},
	{
		period:    6,
		group:     6,
		electrons: []int{2, 8, 18, 32, 12, 2},
		category:  "transition",
		name:      "Tungsten",
		symbol:    "W",
		mass:      183.84,
	},
	{
		period:    6,
		group:     7,
		electrons: []int{2, 8, 18, 32, 13, 2},
		category:  "transition",
		name:      "Rhenium",
		symbol:    "Re",
		mass:      186.207,
	},
	{
		period:    6,
		group:     8,
		electrons: []int{2, 8, 18, 32, 14, 2},
		category:  "transition",
		name:      "Osmium",
		symbol:    "Os",
		mass:      190.23,
	},
	{
		period:    6,
		group:     9,
		electrons: []int{2, 8, 18, 32, 15, 2},
		category:  "transition",
		name:      "Iridium",
		symbol:    "Ir",
		mass:      192.217,
	},
	{
		period:    6,
		group:     10,
		electrons: []int{2, 8, 18, 32, 17, 1},
		category:  "transition",
		name:      "Platinum",
		symbol:    "Pt",
		mass:      195.084,
	},
	{
		period:    6,
		group:     11,
		electrons: []int{2, 8, 18, 32, 18, 1},
		category:  "transition",
		name:      "Gold",
		symbol:    "Au",
		mass:      196.96657,
	},
	{
		period:    6,
		group:     12,
		electrons: []int{2, 8, 18, 32, 18, 2},
		category:  "transition",
		name:      "Mercury",
		symbol:    "Hg",
		mass:      200.59,
	},
	{
		period:    6,
		group:     13,
		electrons: []int{2, 8, 18, 32, 18, 3},
		category:  "posttransition",
		name:      "Thallium",
		symbol:    "Tl",
		mass:      204.3833,
	},
	{
		period:    6,
		group:     14,
		electrons: []int{2, 8, 18, 32, 18, 4},
		category:  "posttransition",
		name:      "Lead",
		symbol:    "Pb",
		mass:      207.2,
	},
	{
		period:    6,
		group:     15,
		electrons: []int{2, 8, 18, 32, 18, 5},
		category:  "posttransition",
		name:      "Bismuth",
		symbol:    "Bi",
		mass:      208.9804,
	},
	{
		period:    6,
		group:     16,
		electrons: []int{2, 8, 18, 32, 18, 6},
		category:  "metalloid",
		name:      "Polonium",
		symbol:    "Po",
		mass:      209,
	},
	{
		period:    6,
		group:     17,
		electrons: []int{2, 8, 18, 32, 18, 7},
		category:  "halogen",
		name:      "Astatine",
		symbol:    "At",
		mass:      210,
	},
	{
		period:    6,
		group:     18,
		electrons: []int{2, 8, 18, 32, 18, 8},
		category:  "noble",
		name:      "Radon",
		symbol:    "Rn",
		mass:      222,
	},
	{
		period:    7,
		group:     1,
		electrons: []int{2, 8, 18, 32, 18, 8, 1},
		category:  "alkali",
		name:      "Francium",
		symbol:    "Fr",
		mass:      223,
	},
	{
		period:    7,
		group:     2,
		electrons: []int{2, 8, 18, 32, 18, 8, 2},
		category:  "alkaline",
		name:      "Radium",
		symbol:    "Ra",
		mass:      226,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 18, 9, 2},
		category:  "actinoid",
		name:      "Actinium",
		symbol:    "Ac",
		mass:      227,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 18, 10, 2},
		category:  "actinoid",
		name:      "Thorium",
		symbol:    "Th",
		mass:      232.03806,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 20, 9, 2},
		category:  "actinoid",
		name:      "Protactinium",
		symbol:    "Pa",
		mass:      231.03587,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 21, 9, 2},
		category:  "actinoid",
		name:      "Uranium",
		symbol:    "U",
		mass:      238.02892,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 22, 9, 2},
		category:  "actinoid",
		name:      "Neptunium",
		symbol:    "Np",
		mass:      237,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 24, 8, 2},
		category:  "actinoid",
		name:      "Plutonium",
		symbol:    "Pu",
		mass:      244,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 25, 8, 2},
		category:  "actinoid",
		name:      "Americium",
		symbol:    "Am",
		mass:      243,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 25, 9, 2},
		category:  "actinoid",
		name:      "Curium",
		symbol:    "Cm",
		mass:      247,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 27, 8, 2},
		category:  "actinoid",
		name:      "Berkelium",
		symbol:    "Bk",
		mass:      247,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 28, 8, 2},
		category:  "actinoid",
		name:      "Californium",
		symbol:    "Cf",
		mass:      251,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 29, 8, 2},
		category:  "actinoid",
		name:      "Einsteinium",
		symbol:    "Es",
		mass:      252,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 30, 8, 2},
		category:  "actinoid",
		name:      "Fermium",
		symbol:    "Fm",
		mass:      257,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 31, 8, 2},
		category:  "actinoid",
		name:      "Mendelevium",
		symbol:    "Md",
		mass:      258,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 32, 8, 2},
		category:  "actinoid",
		name:      "Nobelium",
		symbol:    "No",
		mass:      259,
	},
	{
		period:    7,
		group:     -1,
		electrons: []int{2, 8, 18, 32, 32, 8, 3},
		category:  "actinoid",
		name:      "Lawrencium",
		symbol:    "Lr",
		mass:      262,
	},
	{
		period:    7,
		group:     4,
		electrons: []int{2, 8, 18, 32, 32, 10, 2},
		category:  "transition",
		name:      "Rutherfordium",
		symbol:    "Rf",
		mass:      267,
	},
	{
		period:    7,
		group:     5,
		electrons: []int{2, 8, 18, 32, 32, 11, 2},
		category:  "transition",
		name:      "Dubnium",
		symbol:    "Db",
		mass:      268,
	},
	{
		period:    7,
		group:     6,
		electrons: []int{2, 8, 18, 32, 32, 12, 2},
		category:  "transition",
		name:      "Seaborgium",
		symbol:    "Sg",
		mass:      271,
	},
	{
		period:    7,
		group:     7,
		electrons: []int{2, 8, 18, 32, 32, 13, 2},
		category:  "transition",
		name:      "Bohrium",
		symbol:    "Bh",
		mass:      272,
	},
	{
		period:    7,
		group:     8,
		electrons: []int{2, 8, 18, 32, 32, 14, 2},
		category:  "transition",
		name:      "Hassium",
		symbol:    "Hs",
		mass:      270,
	},
	{
		period:    7,
		group:     9,
		electrons: []int{2, 8, 18, 32, 32, 15, 2},
		category:  "transition",
		name:      "Meitnerium",
		symbol:    "Mt",
		mass:      276,
	},
	{
		period:    7,
		group:     10,
		electrons: []int{2, 8, 18, 32, 32, 17, 1},
		category:  "transition",
		name:      "Darmstadtium",
		symbol:    "Ds",
		mass:      281,
	},
	{
		period:    7,
		group:     11,
		electrons: []int{2, 8, 18, 32, 32, 18, 1},
		category:  "transition",
		name:      "Roentgenium",
		symbol:    "Rg",
		mass:      280,
	},
	{
		period:    7,
		group:     12,
		electrons: []int{2, 8, 18, 32, 32, 18, 2},
		category:  "transition",
		name:      "Copernicium",
		symbol:    "Cn",
		mass:      285,
	},
	{
		period:    7,
		group:     13,
		electrons: []int{2, 8, 18, 32, 32, 18, 3},
		category:  "posttransition",
		name:      "Ununtrium",
		symbol:    "Uut",
		mass:      284,
	},
	{
		period:    7,
		group:     14,
		electrons: []int{2, 8, 18, 32, 32, 18, 4},
		category:  "posttransition",
		name:      "Flerovium",
		symbol:    "Fl",
		mass:      289,
	},
	{
		period:    7,
		group:     15,
		electrons: []int{2, 8, 18, 32, 32, 18, 5},
		category:  "posttransition",
		name:      "Ununpentium",
		symbol:    "Uup",
		mass:      288,
	},
	{
		period:    7,
		group:     16,
		electrons: []int{2, 8, 18, 32, 32, 18, 6},
		category:  "posttransition",
		name:      "Livermorium",
		symbol:    "Lv",
		mass:      293,
	},
	{
		period:    7,
		group:     17,
		electrons: []int{2, 8, 18, 32, 32, 18, 7},
		category:  "halogen",
		name:      "Ununseptium",
		symbol:    "Uus",
		mass:      294,
	},
	{
		period:    7,
		group:     18,
		electrons: []int{2, 8, 18, 32, 32, 18, 8},
		category:  "noble",
		name:      "Ununoctium",
		symbol:    "Uuo",
		mass:      294,
	},
}

func numElements() int {
	return len(allElements)
}

var (
	Hydrogen      = Element{1}
	Helium        = Element{2}
	Lithium       = Element{3}
	Beryllium     = Element{4}
	Boron         = Element{5}
	Carbon        = Element{6}
	Nitrogen      = Element{7}
	Oxygen        = Element{8}
	Fluorine      = Element{9}
	Neon          = Element{10}
	Sodium        = Element{11}
	Magnesium     = Element{12}
	Aluminium     = Element{13}
	Silicon       = Element{14}
	Phosphorus    = Element{15}
	Sulfur        = Element{16}
	Chlorine      = Element{17}
	Argon         = Element{18}
	Potassium     = Element{19}
	Calcium       = Element{20}
	Scandium      = Element{21}
	Titanium      = Element{22}
	Vanadium      = Element{23}
	Chromium      = Element{24}
	Manganese     = Element{25}
	Iron          = Element{26}
	Cobalt        = Element{27}
	Nickel        = Element{28}
	Copper        = Element{29}
	Zinc          = Element{30}
	Gallium       = Element{31}
	Germanium     = Element{32}
	Arsenic       = Element{33}
	Selenium      = Element{34}
	Bromine       = Element{35}
	Krypton       = Element{36}
	Rubidium      = Element{37}
	Strontium     = Element{38}
	Yttrium       = Element{39}
	Zirconium     = Element{40}
	Niobium       = Element{41}
	Molybdenum    = Element{42}
	Technetium    = Element{43}
	Ruthenium     = Element{44}
	Rhodium       = Element{45}
	Palladium     = Element{46}
	Silver        = Element{47}
	Cadmium       = Element{48}
	Indium        = Element{49}
	Tin           = Element{50}
	Antimony      = Element{51}
	Tellurium     = Element{52}
	Iodine        = Element{53}
	Xenon         = Element{54}
	Caesium       = Element{55}
	Barium        = Element{56}
	Lanthanum     = Element{57}
	Cerium        = Element{58}
	Praseodymium  = Element{59}
	Neodymium     = Element{60}
	Promethium    = Element{61}
	Samarium      = Element{62}
	Europium      = Element{63}
	Gadolinium    = Element{64}
	Terbium       = Element{65}
	Dysprosium    = Element{66}
	Holmium       = Element{67}
	Erbium        = Element{68}
	Thulium       = Element{69}
	Ytterbium     = Element{70}
	Lutetium      = Element{71}
	Hafnium       = Element{72}
	Tantalum      = Element{73}
	Tungsten      = Element{74}
	Rhenium       = Element{75}
	Osmium        = Element{76}
	Iridium       = Element{77}
	Platinum      = Element{78}
	Gold          = Element{79}
	Mercury       = Element{80}
	Thallium      = Element{81}
	Lead          = Element{82}
	Bismuth       = Element{83}
	Polonium      = Element{84}
	Astatine      = Element{85}
	Radon         = Element{86}
	Francium      = Element{87}
	Radium        = Element{88}
	Actinium      = Element{89}
	Thorium       = Element{90}
	Protactinium  = Element{91}
	Uranium       = Element{92}
	Neptunium     = Element{93}
	Plutonium     = Element{94}
	Americium     = Element{95}
	Curium        = Element{96}
	Berkelium     = Element{97}
	Californium   = Element{98}
	Einsteinium   = Element{99}
	Fermium       = Element{100}
	Mendelevium   = Element{101}
	Nobelium      = Element{102}
	Lawrencium    = Element{103}
	Rutherfordium = Element{104}
	Dubnium       = Element{105}
	Seaborgium    = Element{106}
	Bohrium       = Element{107}
	Hassium       = Element{108}
	Meitnerium    = Element{109}
	Darmstadtium  = Element{110}
	Roentgenium   = Element{111}
	Copernicium   = Element{112}
	Ununtrium     = Element{113}
	Flerovium     = Element{114}
	Ununpentium   = Element{115}
	Livermorium   = Element{116}
	Ununseptium   = Element{117}
	Ununoctium    = Element{118}
)
