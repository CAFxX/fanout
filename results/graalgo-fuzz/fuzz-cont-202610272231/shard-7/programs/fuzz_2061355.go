package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
func genFn0(p0 string, p1 uint64, p2 uint64) uint16 {
    if (4294967295 > 2147483648) {
        if true {
            fmt.Println(p2, p0)
            v7 := GenS1{F0: 42, F1: 42, F2: 1}
            _ = v7
        }
        if ((-396 == 1) || (true && true)) {
            for fi8 := 0; fi8 < 19; fi8++ {
                p0 = p0
                for fi9 := 0; fi9 < 8; fi9++ {
                    v10 := []int{fi9}
                    _ = fi9
                    _ = v10
                }
                if true {
                    switch 2147483647 {
                    case 461:
                        if (!(901 <= -2147483648)) {
                            p0 = p0
                        } else {
                            if ((fi8 * 0) <= fi8) {
                                if true {
                                    var v11 int8 = 26
                                    var v12 float32 = 1.5
                                    _ = v11
                                    _ = v12
                                } else {
                                    var v13 uint = 366
                                    p1 += p2
                                    _ = v13
                                }
                            } else {
                                p2 += 2
                                switch (p0 + p0) {
                                case "\n\t":
                                    v14 := []int{}
                                    _ = v14
                                }
                                p0 = "foo bar"
                            }
                            v15 := []int{-458, 339, (fi8 + 0), fi8}
                            _ = len(v15)
                            _ = v15
                        }
                    case -1:
                        var v16 float64 = 3.14159
                        _ = v16
                    default:
                        var v17 uint16 = uint16(p2)
                        _ = v17
                    }
                    switch fi8 {
                    case -312:
                        p1 = 893
                    case -9223372036854775808:
                        v18 := GenS1{F0: 42, F1: 219, F2: fi8}
                        _ = v18
                    }
                }
                _ = fi8
            }
            var v19 uint32 = 538
            var v20 int8 = 70
            _ = v19
            _ = v20
        }
        if true {
            v21 := []int{-1, int(p1)}
            _ = v21
        }
    }
    if true {
        v22 := []int{}
        _ = len(v22)
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := GenS1{F0: 1, F1: 979, F2: (fi23 | 127)}
            fmt.Println(p1)
            if len(v22) > 0 {
                v22[0] = fi23
            }
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    v25 := GenS1{F0: 0, F1: 878, F2: 2147483647}
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := []int{-485, fi26}
        if false {
            fmt.Println(p0, p1, fi26)
            var v28 float64 = 1e10
            _ = v28
        }
        if (true || ((!false) || ((9 != 24) || ((4294967295 > 825) && false)))) {
            var v29 int = 0
            fmt.Println(p0, p0, p1)
            _ = v29
        } else {
            _ = len(v27)
            p0 = ("\n\t" + p0)
            var v30 int32 = 2147483647
            _ = v30
        }
        _ = fi26
        _ = v27
    }
    v31 := []int{127, -9223372036854775808, 32768, 122}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v31
    return 32767
}
func genFn1() int16 {
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 int8 = -1
        var v34 int64 = -9223372036854775808
        v35 := []int{}
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    v36 := GenS1{F0: 1, F1: 65535, F2: 4294967295}
    _ = v36
    return 32767
}

func main() {
    v1 := []uint32{878, 42, 65535, 65535}
    var v2 int16 = 42
    var v3 uint16 = 409
    v4 := []uint32{4294967295, 42, 4294967295}
    var v5 int16 = 1000
    var v6 uint = 18446744073709551615
    v2 = 1
    var v37 int32 = 724
    switch int(v6) {
    case -2147483648:
        for fi38 := 0; fi38 < 4; fi38++ {
            _ = len(v4)
            _ = fi38
        }
    default:
        fmt.Println(v3)
    }
    v5 = 42
    v5 = -603
    v5 = 0
    if false {
        v2 = (^(v5))
    }
    for fi39 := 0; fi39 < 2; fi39++ {
        v6 -= 0
        fmt.Println(v2)
        _ = fi39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
