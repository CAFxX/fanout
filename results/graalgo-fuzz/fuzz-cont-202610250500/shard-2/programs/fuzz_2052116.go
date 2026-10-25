package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int; F1 bool; F2 int64 }
func genFn0(p0 int, p1 int16, p2 int32) uint32 {
    var v5 float64 = 0.1
    var v6 string = ""
    fmt.Println(p2, p0)
    var v7 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 uint8 = 128
    var v2 float64 = 0.1
    v3 := []int{0, -2147483648}
    v4 := []string{}
    if len(v3) > 0 {
        v3[0] = -437
    }
    var v8 bool = (!(-885 != 904))
    var v9 float64 = v2
    v8 = (42 < 102)
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 string = "\n\t"
            _ = len(v4)
            v12 = v12
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v13 := GenS1{F0: 832, F1: genFn0(2147483647, -32768, -1), F2: 4294967295}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if v8 {
        if v8 {
            var v14 uint = 18446744073709551615
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "" {
            case "foo bar":
                for fi15 := 0; fi15 < 4; fi15++ {
                    v1 = v1
                    if v8 {
                        v8 = (1 < 100)
                        var v16 uint32 = 4294967295
                        v14 = (^(v14))
                        _ = v16
                    }
                    _ = fi15
                }
            case "a":
                var v17 float32 = (1e10 + float32(v2))
                _ = v17
            case "hello":
                v2 += float64(float32((v2 / ((3.14159 / (v9 - ((v2 - 1.5) * float64(float32(v9))))) / (v9 * (v2 / (v9 / (float64(1e10) + (float64(0.1) / v9)))))))))
            }
            _ = v14
        }
    }
    var v18 int64 = -9223372036854775808
    var v19 uint64 = 1
    for fi20 := 0; fi20 < 20; fi20++ {
        fmt.Println(v19, v19, v9)
        switch (-(fi20)) {
        case -973:
            v21 := GenS1{F0: 32767, F1: 990, F2: v19}
            _ = v21
        case -2147483648:
            fmt.Println(v19, v8, v9)
        default:
            fmt.Println(v9)
        }
        _ = fi20
    }
    _ = len(v3)
    var v22 float64 = -2.25
    var v23 float64 = 0.0
    if len(v3) > 0 {
        v3[0] = 32768
    }
    _ = len(v3)
    if v8 {
        v1 = uint8(v19)
        if v8 {
            var v24 uint8 = (v1 | v1)
            v18 = 1
            _ = v24
        } else {
            var v25 float64 = v23
            v3 = append(v3, -9223372036854775808)
            switch "" {
            case "foo bar":
                switch 0 {
                case 4294967295:
                    if len(v3) > 0 {
                        v3[0] = int(v18)
                    }
                }
            case "a":
                fmt.Println(v1)
            default:
                v2 -= 1e10
            }
            _ = v25
        }
        fmt.Println(v8, v23, v18)
    } else {
        v26 := GenS2{F0: 636, F1: (1000 != 314), F2: (v18 | (-(v18)))}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
