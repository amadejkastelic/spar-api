package sparsi

import "encoding/json"

type FacetElement struct {
	Text         string          `json:"text,omitempty"`
	TotalHits    int             `json:"totalHits,omitempty"`
	Selected     string          `json:"selected,omitempty"`
	ClusterLevel int             `json:"clusterLevel,omitempty"`
	SearchParams json.RawMessage `json:"searchParams,omitempty"`
}

type Facet struct {
	Name                string             `json:"name,omitempty"`
	DetailedLinks       int                `json:"detailedLinks,omitempty"`
	Unit                string             `json:"unit,omitempty"`
	Type                string             `json:"type,omitempty"`
	ShowPreviewImages   bool               `json:"showPreviewImages,omitempty"`
	FilterStyle         string             `json:"filterStyle,omitempty"`
	SelectionType       string             `json:"selectionType,omitempty"`
	AssociatedFieldName string             `json:"associatedFieldName,omitempty"`
	SelectedElements    []*json.RawMessage `json:"selectedElements,omitempty"`
	Elements            []*FacetElement    `json:"elements,omitempty"`
}

type Product struct {
	IsOnPromotion       string     `json:"is-on-promotion,omitempty"`
	CategoryNames       string     `json:"category-names,omitempty"`
	BadgeIcon           string     `json:"badge-icon,omitempty"`
	Description         string     `json:"description,omitempty"`
	SalesUnit           string     `json:"sales-unit,omitempty"`
	Title               string     `json:"title,omitempty"`
	BadgeNames          string     `json:"badge-names,omitempty"`
	ItemType            string     `json:"item-type,omitempty"`
	CodeInternal        string     `json:"code-internal,omitempty"`
	CategoryName        string     `json:"category-name,omitempty"`
	AllergensFilter     []string   `json:"allergens-filter,omitempty"`
	Price               float64    `json:"price,omitempty"`
	BestPrice           float64    `json:"best-price,omitempty"`
	BadgeShortName      string     `json:"badge-short-name,omitempty"`
	CreatedAt           string     `json:"created-at,omitempty"`
	Categories          []string   `json:"categories,omitempty"`
	CategoryPath        [][]string `json:"category-path,omitempty"`
	ShortDescription    string     `json:"short-description,omitempty"`
	ShortDescription2   string     `json:"short-description-2,omitempty"`
	EcrCategoryNumber   string     `json:"ecr-category-number,omitempty"`
	ProductFeatures     []string   `json:"product-features,omitempty"`
	StockStatus         string     `json:"stock-status,omitempty"`
	IsNew               string     `json:"is-new,omitempty"`
	ImageURL            string     `json:"image-url,omitempty"`
	CategoryID          string     `json:"category-id,omitempty"`
	ApproxWeightProduct string     `json:"approx-weight-product,omitempty"`
	URL                 string     `json:"url,omitempty"`
	PosVisible          []string   `json:"pos-visible,omitempty"`
	Name                string     `json:"name,omitempty"`
	ProductNumber       string     `json:"product-number,omitempty"`
	ProcePerUnit        string     `json:"price-per-unit,omitempty"`
	RegularPrice        float64    `json:"regular-price,omitempty"`
	PricePerUnitNumber  float64    `json:"price-per-unit-number,omitempty"`
	PosPurchasable      string     `json:"pos-purchasable,omitempty"`
}

type Hit struct {
	Product       *Product `json:"masterValues,omitempty"`
	VariantValues []string `json:"variant-values,omitempty"`
	ID            string   `json:"id,omitempty"`
	Score         float64  `json:"score,omitempty"`
	Position      int      `json:"position,omitempty"`
	FoundWords    []string `json:"found-words,omitempty"`
}

type Paging struct {
	CurrentPage        int             `json:"currentPage,omitempty"`
	PageCount          int             `json:"pageCount,omitempty"`
	HitsPerPage        int             `json:"hitsPerPage,omitempty"`
	DefaultHitsPerPage int             `json:"defaultHitsPerPage,omitempty"`
	NextLink           json.RawMessage `json:"nextLink,omitempty"`
}

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

type SortField string

const (
	SortFieldRelevancy    SortField = "relevancy"
	SortFieldCreatedAt    SortField = "created-at"
	SortFieldPrice        SortField = "best-price"
	SortFieldTitle        SortField = "title"
	SortFieldPricePerUnit SortField = "price-per-unit"
)

type Sort struct {
	Field SortField `json:"field,omitempty"`
	Order Order     `json:"order,omitempty"`
}

type FilterValue struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Exclude bool   `json:"exclude,omitempty"`
}

type Filter struct {
	Name      string        `json:"name,omitempty"`
	Values    []FilterValue `json:"values,omitempty"`
	Substring bool          `json:"substring,omitempty"`
}

type PriceRangeFilter struct {
	MinPrice float64
	MaxPrice float64
}

type SearchRequest struct {
	Q           string
	Query       string
	HitsPerPage int
	Page        int
	Filters     []Filter
	Sort        *Sort
	PriceRange  *PriceRangeFilter
}

type SearchResponse struct {
	Hits                []*Hit   `json:"hits,omitempty"`
	TotalHits           int      `json:"totalHits,omitempty"`
	ArticleNumberSearch bool     `json:"articleNumberSearch,omitempty"`
	ScoreFirstHit       float64  `json:"scoreFirstHit,omitempty"`
	ScoreLastHit        float64  `json:"scoreLastHit,omitempty"`
	TookTotal           int      `json:"tookTotal,omitempty"`
	TookWorldmatch      int      `json:"tookWorldmatch,omitempty"`
	TookLoop54          int      `json:"tookLoop54,omitempty"`
	TookAtlasAi         int      `json:"tookAtlasAi,omitempty"`
	TookGPTSynonyms     int      `json:"tookGPTSynonyms,omitempty"`
	TimedOut            bool     `json:"timedOut,omitempty"`
	SplitDocuments      bool     `json:"splitDocuments,omitempty"`
	Paging              *Paging  `json:"paging,omitempty"`
	Facets              []*Facet `json:"facets,omitempty"`
}

type Category struct {
	Name         string   `json:"name,omitempty"`
	FilterName   string   `json:"filterName,omitempty"`
	FilterValues []string `json:"filterValues,omitempty"`
}

type CategoriesResponse struct {
	Values []*Category `json:"categories,omitempty"`
}
