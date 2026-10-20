package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint8, p1 int8, p2 int8) uint16 {
    var v6 uint8 = p0
    switch 0 {
    case -2147483648:
        p2 += (p1 / (((p2 - 1)) | 1))
    }
    if (((((((1.5 * 1e10) + float64(3.14159)) / 1.0) + (float64(0.0) + 3.14159)) * ((((1e10 - (0.1 + 1e10)) + 0.1) * -2.25) + 3.14159)) - 1e10) > 1.5) {
        v7 := GenS1{F0: (p1 + (p2 - 0))}
        if (0 != 185) {
            v8 := GenS1{F0: p1}
            _ = v8
        } else {
            fmt.Println(p0, p0)
            for fi9 := 0; fi9 < 9; fi9++ {
                for fi10 := 0; fi10 < 9; fi10++ {
                    v6 = p0
                    _ = fi10
                }
                switch (fi9 << ((52) & 63)) {
                case 32768:
                    var v11 uint64 = 363
                    _ = v11
                case -2147483648:
                    for fi12 := 0; fi12 < 6; fi12++ {
                        v13 := []int{(fi9 * (fi9 * fi9))}
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = fi12
                        _ = v13
                    }
                default:
                    if (true || (!false)) {
                        var v14 int32 = 1
                        for fi15 := 0; fi15 < 6; fi15++ {
                            var v16 int64 = 841
                            _ = fi15
                            _ = v16
                        }
                        var v17 string = "foo bar"
                        _ = v14
                        _ = v17
                    }
                }
                _ = fi9
            }
        }
        _ = v7
    }
    v18 := GenS1{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v6
    return 812
}
func genFn1() int16 {
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := []int{}
            v22 := GenS1{F0: -18}
            _ = fi20
            _ = v21
            _ = v22
        }
        v23 := []int{fi19, fi19, fi19}
        _ = fi19
        _ = v23
    }
    switch -740 {
    case 294:
        var v24 int8 = 42
        _ = v24
    }
    return 0
}

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{105, 4294967295}
    v3 := make(map[int]bool)
    var v4 uint64 = 18446744073709551615
    v5 := []int{127, 1, -9223372036854775808}
    v4 = (v4 ^ v4)
    var v25 uint16 = 65535
    if true {
        v25 = genFn0(2, -50, -1)
        v4 -= (^(v4))
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 int8 = 42
            _ = len(v2)
            switch (fi26 & 9223372036854775807) {
            case 42:
                v28 := GenS1{F0: (^(v27))}
                _ = v28
            default:
                v27 = v27
            }
            _ = fi26
            _ = v27
        }
    }
    v25 = (v25 % ((v25) | 1))
    v25 = v25
    v4 = (v4 | 29)
    fmt.Println(v25, v1, v1)
    var v29 uint16 = genFn0(42, -128, 127)
    var v30 uint = 42
    v1 -= (v1 * (-(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
