package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint8) uint {
    p0 -= (p0 % (((p0 & (p0 - p0))) | 1))
    var v7 uint16 = 0
    switch 822 {
    case 2147483647:
        v8 := GenS1{F0: 65535}
        _ = v8
    case 43:
        v9 := []int{701, 127, 128, 128}
        _ = v9
    }
    _ = p0
    _ = v7
    return 255
}

func main() {
    var v1 int64 = 0
    var v2 uint64 = 0
    var v3 uint16 = 270
    var v4 float64 = (1e10 - 3.14159)
    var v5 int8 = -15
    var v6 uint32 = uint32(v3)
    switch -928 {
    case 907:
        var v10 bool = ((0.1 == ((3.14159 - (1e10 - 1e10)) * 3.14159)) && false)
        _ = v10
    case -551:
        v11 := []int{}
        _ = v11
    default:
        if (0 > uint8(v2)) {
            if ("\n\t" < "a") {
                v3 -= 1
                var v12 int64 = (^(v1))
                _ = v12
            }
        }
    }
    if true {
        v13 := GenS1{F0: (v6 << ((916) & 63))}
        _ = v13
    }
    v14 := []int{42, 42}
    var v15 uint32 = v6
    v15 = (v15 & (^(v6)))
    var v16 string = ((("" + "hello") + (("a" + ("hello" + "hello")) + "foo bar")) + "\n\t")
    v2 += 18446744073709551615
    var v17 bool = (v4 < ((v4 / v4) / v4))
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
