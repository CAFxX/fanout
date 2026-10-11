package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
func genFn0() uint32 {
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(fi5)
        v6 := GenS1{F0: false}
        fmt.Println(fi5)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: (true || false)}
        v9 := GenS1{F0: true}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 2147483648
    v11 := GenS2{F0: 2147483648, F1: true, F2: -32768}
    var v12 uint8 = 2
    _ = v10
    _ = v11
    _ = v12
    return v10
}
func genFn1(p0 uint32, p1 uint8, p2 uint16) uint {
    var v13 string = ""
    p2 = (p2 ^ 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 255
}

func main() {
    var v1 bool = (true || true)
    var v2 float64 = 0.1
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    v14 := []int{2147483647, -2147483648, 1}
    v2 = v2
    var v15 uint64 = 1
    v1 = true
    switch 4294967295 {
    case 373:
        for fi16 := 0; fi16 < 12; fi16++ {
            v15 -= 4294967295
            _ = fi16
        }
    }
    switch 552 {
    case 2147483647:
        v2 = (((v2 + 0.0) + float64(((-2.25 - ((-2.25 * 1e10) + ((0.0 / float32(v2)) - 1e10))) * 0.1))) + 0.0)
    case 0:
        v1 = (false || true)
    case -642:
        var v17 bool = true
        _ = v17
    }
    if v1 {
        var v18 int8 = 0
        _ = v18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS1{F0: ("foo bar" < "foo bar")}
        if (!(-32768 == -1)) {
            for fi21 := 0; fi21 < 14; fi21++ {
                if v1 {
                    v1 = true
                }
                _ = fi21
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 uint16 = 1
                _ = fi22
                _ = v23
            }
        } else {
            if len(v14) > 0 {
                v14[0] = 218
            }
            fmt.Println(v15)
            var v24 string = "\n\t"
            _ = v24
        }
        _ = fi19
        _ = v20
    }
    fmt.Println(v2, v2, v15)
    v1 = (50 == 1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
