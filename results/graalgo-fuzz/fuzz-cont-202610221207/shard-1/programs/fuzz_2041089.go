package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int, p1 bool) uint32 {
    fmt.Println(p0, p0)
    p0 = p0
    p1 = (((1e10 * 1e10) * 1e10) < 0.0)
    fmt.Println(p1, p0)
    p1 = ((1.5 <= 1.5) && ((("foo bar" + "foo bar") + ("\n\t" + "\n\t")) == "hello"))
    _ = p0
    _ = p1
    return 97
}

func main() {
    var v1 uint8 = 113
    v2 := make(map[uint32]string)
    var v3 string = ("\n\t" + ("hello" + "foo bar"))
    if true {
        var v4 uint = 0
        var v5 int = 544
        _ = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        v1 += 128
        for fi7 := 0; fi7 < 7; fi7++ {
            v1 += (v1 % (((^(v1))) | 1))
            _ = fi7
        }
        _ = fi6
    }
    v8 := []int{32767, -289, 32767, 169}
    delete(v2, 2147483648)
    var v9 uint = 18446744073709551615
    var v10 int32 = -2147483648
    v3 = (v3 + "\n\t")
    v9 = v9
    var v11 uint32 = 2
    v12 := GenS1{F0: 9223372036854775807, F1: 924}
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 uint16 = 255
        v15 := GenS2{F0: 0, F1: v14}
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 uint64 = 1
            if true {
                fmt.Println(v3)
                var v18 uint8 = (v1 * 2)
                _ = v18
            }
            _ = fi16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = len(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
