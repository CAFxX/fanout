package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int8 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
func genFn0() int8 {
    v7 := []int{1, 4294967295}
    var v8 uint32 = 2
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 string) uint {
    switch -1 {
    case -543:
        fmt.Println(p0)
    }
    switch 492 {
    case -569:
        p0 = (p0 + p0)
    default:
        v9 := GenS1{F0: -905, F1: -256, F2: 94}
        _ = v9
    }
    fmt.Println(p0)
    if true {
        v10 := []int{}
        p0 = (("" + ("" + (("foo bar" + "hello") + p0))) + ("" + p0))
        var v11 int = 9223372036854775807
        _ = v10
        _ = v11
    } else {
        fmt.Println(p0)
        v12 := []int{701, 42, 9223372036854775807, 0}
        fmt.Println(p0)
        _ = v12
    }
    v13 := GenS2{F0: 0, F1: 485, F2: 208}
    _ = p0
    _ = v13
    return 890
}

func main() {
    var v1 uint16 = 65535
    v2 := []uint32{65535, 582}
    var v3 int32 = -2147483648
    var v4 uint32 = 264
    v5 := []string{"a", "hello", ""}
    v6 := make(map[string]bool)
    v6["hello"] = ((416 == 1) && (true || (("\n\t" + "") == "a")))
    v4 += 4294967295
    var v14 int32 = (v3 >> ((65535) & 63))
    var v15 uint8 = 208
    var v16 int = 2147483647
    v4 = (v4 >> ((65535) & 63))
    var v17 int16 = 42
    delete(v6, "a")
    var v18 bool = (!true)
    var v19 string = ("" + (("" + "") + "a"))
    v18 = false
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
