package mapcopy

type GenericMap map[interface{}]interface{}

func Copymap(dest GenericMap, src GenericMap)  {
	for k,v := range src {
		dest[k] = v
	}
}
