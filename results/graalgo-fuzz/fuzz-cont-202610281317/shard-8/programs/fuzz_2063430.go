package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint16, p1 int32, p2 uint) int64 {
    var v6 uint16 = p0
    var v7 uint64 = 771
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint8 = 0
    var v2 bool = (0 == 334)
    var v3 int16 = 0
    var v4 uint8 = v1
    v5 := make(map[string]bool)
    if false {
        if (-498 <= 347) {
            var v8 string = ("a" + ("foo bar" + ""))
            var v9 float64 = -2.25
            v10 := GenS1{F0: 128}
            _ = v10
            _ = v8
            _ = v9
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := []int{}
            _ = fi11
            _ = v12
        }
        fmt.Println(v3)
    }
    if (127 >= -128) {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint8 = 117
            _ = fi13
            _ = v14
        }
    }
    v15 := GenS1{F0: (^(v4))}
    v16 := []int{127, 127, -9223372036854775808}
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 uint = 0
        delete(v5, "a")
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 uint8 = (v4 << (((v18 | v18)) & 63))
            v21 := GenS1{F0: 42}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi17
        _ = v18
    }
    v22 := GenS1{F0: v1}
    v3 += -1
    var v23 uint16 = 42
    v24 := GenS2{F0: -2147483648, F1: (v4 + (^(v1)))}
    v23 = (v23 * 255)
    v23 -= (v23 >> ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
