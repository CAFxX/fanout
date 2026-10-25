package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint32; F2 int64 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int16, p1 int8, p2 uint16) int {
    switch -13 {
    case 2147483647:
        for fi4 := 0; fi4 < 7; fi4++ {
            v5 := GenS1{F0: (p2 * 0)}
            var v6 int8 = p1
            v7 := []int{(fi4 / ((-468) | 1)), fi4}
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
    case 32768:
        for fi8 := 0; fi8 < 16; fi8++ {
            fmt.Println(p1, p2, p2)
            _ = fi8
        }
    }
    if true {
        var v9 int = -507
        var v10 uint8 = 255
        _ = v10
        _ = v9
    } else {
        v11 := []int{152, 2147483647, 1}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 int = -634
    var v2 int32 = 2147483647
    v3 := []int{(v1 ^ 2147483647), 4294967295, v1}
    _ = len(v3)
    for fi12 := 0; fi12 < 17; fi12++ {
        v1 += -553
        if ((!(false || (v2 == (-(v2))))) && ((42 <= v2) || (-32768 > -32768))) {
            if (!(42 == 0)) {
                v2 = 1
                for fi13 := 0; fi13 < 16; fi13++ {
                    v14 := GenS2{F0: (v2 | (^(v2)))}
                    v1 += v1
                    _ = fi13
                    _ = v14
                }
            }
            for fi15 := 0; fi15 < 10; fi15++ {
                _ = len(v3)
                v16 := GenS3{F0: (v2 ^ v2), F1: 65535, F2: -1}
                if (!(2 != 1)) {
                    v17 := GenS4{F0: 1, F1: 4294967295}
                    switch "hello" {
                    case "hello":
                        var v18 int8 = 127
                        _ = v18
                    case "":
                        var v19 float64 = (1e10 * 1e10)
                        _ = v19
                    }
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
            var v20 uint64 = 82
            _ = v20
        }
        _ = fi12
    }
    v1 += genFn0(32767, 100, 365)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v21 int = -2147483648
    for fi22 := 0; fi22 < 7; fi22++ {
        v3 = append(v3, -120)
        v23 := GenS3{F0: v2, F1: 2147483648, F2: 890}
        _ = fi22
        _ = v23
    }
    fmt.Println(v2, v2)
    if true {
        if ((-89 >= 0) || true) {
            var v24 int8 = int8(v21)
            _ = v24
        }
    }
    var v25 float64 = 1.5
    v26 := GenS5{F0: -32768}
    for fi27 := 0; fi27 < 20; fi27++ {
        v1 -= (fi27 | (fi27 / (((v1 | genFn0(42, 100, 1))) | 1)))
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
}
