package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint32) int16 {
    v5 := []int{-1, 0}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 66
}

func main() {
    var v1 uint16 = 65535
    var v2 int8 = -1
    var v3 uint64 = 1
    v4 := make(map[string]bool)
    if (true || ("a" == "\n\t")) {
        v3 = v3
    }
    var v6 string = "hello"
    for fi7 := 0; fi7 < 18; fi7++ {
        if (v6 == v6) {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 bool = false
                v10 := []int{(fi7 * (fi7 << ((738) & 63))), (fi7 + 127), 0}
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := []int{(-(fi7))}
            v3 += 4294967295
            _ = v11
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            v6 = (v6 + "foo bar")
            _ = fi12
        }
        var v13 float64 = (-2.25 / float64(3.14159))
        _ = fi7
        _ = v13
    }
    v14 := []int{9223372036854775807, 9223372036854775807, 4294967295, 163}
    var v15 int16 = -199
    v14 = append(v14, -2147483648)
    var v16 uint8 = 2
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v3)
        v2 -= (v2 % ((0) | 1))
        var v18 float64 = float64((3.14159 + 1.5))
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS1{F0: 0}
        var v21 uint8 = uint8(v1)
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 uint32 = 2147483647
    var v23 uint16 = v1
    switch (v6 + "") {
    case "":
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 float32 = ((3.14159 + (((0.0 - 0.1) - 1.5) + -2.25)) - float32(1e10))
            v26 := GenS2{F0: (v22 ^ uint32(v16)), F1: v22}
            v22 += 42
            _ = fi24
            _ = v25
            _ = v26
        }
    case "a":
        var v27 bool = false
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v6
}
