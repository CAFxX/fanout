package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 int32, p2 uint8) uint32 {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: 0, F1: 814, F2: uint32(p2)}
        v6 := GenS1{F0: 255, F1: 219, F2: 1000}
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        p2 += (p2 - (p2 << ((1) & 63)))
        _ = fi7
    }
    if ((0.1 + 0.1) >= float32(0.0)) {
        var v8 string = "foo bar"
        fmt.Println(p1, p2, v8)
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 float32 = 1e10
            fmt.Println(fi9, p0, p0)
            v11 := []int{(fi9 - (p0 & fi9)), (p0 / ((fi9) | 1))}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    } else {
        if false {
            for fi12 := 0; fi12 < 18; fi12++ {
                if true {
                    v13 := []int{p0, p0}
                    _ = v13
                } else {
                    switch (p0 | p0) {
                    case 32767:
                        v14 := []int{(fi12 - fi12), -434}
                        _ = v14
                    }
                    fmt.Println(p2, p1)
                    if false {
                        p1 += -747
                    }
                }
                _ = fi12
            }
            fmt.Println(p1, p0, p1)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}

func main() {
    var v1 uint16 = 1
    var v2 int32 = 0
    var v3 string = ("" + "hello")
    if (!false) {
        var v15 int64 = -192
        _ = v15
    } else {
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 float64 = 1e10
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: v1, F1: 0, F2: 593}
        _ = v18
    }
    var v19 uint8 = 1
    var v20 bool = false
    var v21 uint32 = 259
    fmt.Println(v1, v1, v3)
    for fi22 := 0; fi22 < 16; fi22++ {
        if v20 {
            fmt.Println(v19)
            v23 := GenS1{F0: (v1 ^ 255), F1: 148, F2: uint32(v19)}
            if (115 <= 2) {
                if v20 {
                    v24 := []int{fi22}
                    _ = v24
                } else {
                    var v25 int32 = v2
                    switch (v3 + ("" + (v3 + v3))) {
                    case "\n\t":
                        v26 := GenS1{F0: v1, F1: 4294967295, F2: v21}
                        _ = v26
                    }
                    v1 -= v1
                    _ = v25
                }
                switch "\n\t" {
                case "":
                    for fi27 := 0; fi27 < 12; fi27++ {
                        v28 := []int{163, fi27, fi27, fi22}
                        if (3.14159 > 0.1) {
                            v28 = append(v28, (fi27 - 32767))
                        } else {
                            var v29 string = v3
                            _ = v29
                        }
                        fmt.Println(v1)
                        _ = fi27
                        _ = v28
                    }
                case "\n\t":
                    v30 := GenS2{F0: 42, F1: 18446744073709551615, F2: (v19 / ((124) | 1))}
                    _ = v30
                default:
                    v21 -= (v21 | v21)
                }
            }
            _ = v23
        }
        var v31 float32 = -2.25
        v32 := []int{}
        _ = fi22
        _ = v31
        _ = v32
    }
    fmt.Println(v3)
    v33 := []int{4294967295, 477, int(v19), 42}
    v34 := GenS3{F0: v2}
    v35 := GenS1{F0: v1, F1: 4294967295, F2: (v21 * 1)}
    for fi36 := 0; fi36 < 17; fi36++ {
        if ((v20 && (v20 || ((42 <= uint8(v21)) || false))) || (((!v20) && (fi36 < (fi36 & 9223372036854775807))) && ((v21 & (v21 | 4294967295)) != v21))) {
            if true {
                v2 += -37
            }
            v1 += v1
        }
        if len(v33) > 0 {
            v33[0] = fi36
        }
        _ = fi36
    }
    v37 := GenS1{F0: (^(v1)), F1: 0, F2: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v37
}
