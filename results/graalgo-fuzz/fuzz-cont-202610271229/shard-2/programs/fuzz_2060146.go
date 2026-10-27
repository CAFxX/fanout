package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int, p1 uint16) uint64 {
    v4 := GenS1{F0: -128}
    for fi5 := 0; fi5 < 3; fi5++ {
        p1 = (p1 + (^(p1)))
        _ = fi5
    }
    v6 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 45
}
func genFn1(p0 int64, p1 int) int16 {
    var v7 string = "\n\t"
    v8 := GenS1{F0: -128}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []uint32{65535, 1, 763, 1}
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    v9 := GenS2{F0: 503}
    var v10 int8 = 100
    fmt.Println(v10)
    v11 := GenS3{F0: 1}
    for fi12 := 0; fi12 < 20; fi12++ {
        switch "" {
        case "\n\t":
            var v13 float32 = 1.5
            _ = v13
        default:
            var v14 float64 = float64(-2.25)
            _ = v14
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            if (!false) {
                var v16 string = "\n\t"
                v10 -= v10
                _ = v16
            }
            v17 := v2[("a" + ("\n\t" + (("hello" + "hello") + "\n\t")))]
            _ = fi15
            _ = v17
        }
        _ = fi12
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS2{F0: 4294967295}
        _ = fi18
        _ = v19
    }
    var v20 uint32 = 656
    v20 = 4294967295
    v10 = (-(v10))
    v10 = v10
    v20 += (^(v20))
    v1 = append(v1, (v20 << ((42) & 63)))
    var v21 uint8 = 70
    switch 811 {
    case 2147483647:
        for fi22 := 0; fi22 < 6; fi22++ {
            if true {
                var v23 string = ""
                if false {
                    v10 = (v10 * -109)
                    fmt.Println(v10, v10, v10)
                    v3[v20] = (false && (1.5 != ((0.1 - 3.14159) / -2.25)))
                } else {
                    v10 = v10
                    var v24 uint = 65535
                    _ = v24
                }
                _ = v23
            } else {
                v25 := GenS3{F0: 0}
                _ = v25
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                v27 := GenS1{F0: (v10 ^ 67)}
                v10 = v10
                _ = fi26
                _ = v27
            }
            _ = fi22
        }
    case 127:
        if true {
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 string = "\n\t"
                var v30 int8 = v10
                _ = fi28
                _ = v29
                _ = v30
            }
        }
    }
    var v31 int32 = -485
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v9
}
