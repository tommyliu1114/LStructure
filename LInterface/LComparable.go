package LInterface

type LComparable interface {
	Compare(interface{}) int // lower -1 equal 0  bigger 1
}
