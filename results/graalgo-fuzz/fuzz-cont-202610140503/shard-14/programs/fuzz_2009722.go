package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint8; F2 bool }
func genFn0(p0 int16, p1 string, p2 bool) bool {
    v6 := GenS1{F0: false, F1: 1, F2: 1}
    fmt.Println(p2, p1, p2)
    if (!(!true)) {
        p0 = (p0 * p0)
    } else {
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(p1)
            v8 := []int{fi7}
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint32 = 65535
                if len(v8) > 0 {
                    v8[0] = fi9
                }
                _ = len(v8)
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        v11 := GenS1{F0: p2, F1: 0, F2: 834}
        p1 = ("\n\t" + "")
        _ = v11
    }
    if (2147483647 != 42) {
        v12 := GenS2{F0: p1}
        p0 -= (-(p0))
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1(p0 uint32) int {
    var v13 uint = 18446744073709551615
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 uint16 = 65535
        v15 -= (v15 + (v15 >> ((v13) & 63)))
        var v16 uint8 = 255
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = v13
    return 0
}

func main() {
    v1 := []uint32{4294967295, 151}
    v2 := make(map[int]int)
    v3 := []uint32{1, 2147483647}
    v4 := make(map[string]bool)
    var v5 float32 = -2.25
    fmt.Println(v5)
    var v17 uint = 2
    for fi18 := 0; fi18 < 3; fi18++ {
        _ = len(v2)
        if (false || (!(-62 > 63))) {
            var v19 int = (fi18 + (^(fi18)))
            _ = v19
        } else {
            v20 := GenS2{F0: "\n\t"}
            var v21 float64 = ((1e10 - ((((0.0 / -2.25) + 0.0) / 3.14159) - -2.25)) + 0.0)
            for fi22 := 0; fi22 < 8; fi22++ {
                if true {
                    v23 := GenS3{F0: uint64(v17), F1: 255, F2: false}
                    _ = v23
                } else {
                    var v24 uint64 = 4294967295
                    v17 -= (v17 + 750)
                    _ = v24
                }
                v25 := GenS3{F0: uint64(v17), F1: 169, F2: (true || false)}
                _ = fi22
                _ = v25
            }
            _ = v20
            _ = v21
        }
        if true {
            if (!false) {
                v5 = 1.5
                v5 = 0.0
            }
            v17 -= v17
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            _ = len(v4)
            for fi26 := 0; fi26 < 9; fi26++ {
                switch (fi18 % (((fi18 % ((987) | 1))) | 1)) {
                case -1:
                    var v27 int64 = int64(v17)
                    _ = v27
                case 0:
                    if len(v1) > 0 {
                        v1[0] = 93
                    }
                case -9223372036854775808:
                    v28 := GenS2{F0: ("" + "foo bar")}
                    _ = v28
                }
                var v29 int32 = -163
                var v30 float32 = 1.5
                _ = fi26
                _ = v29
                _ = v30
            }
            var v31 bool = (1 < 0)
            _ = v31
        }
        _ = fi18
    }
    var v32 bool = true
    v17 -= v17
    fmt.Println(v5)
    var v33 float64 = 0.0
    if (!(1 <= (v17 ^ (^(v17))))) {
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 uint64 = 925
            _ = fi34
            _ = v35
        }
    } else {
        fmt.Println(v33, v5)
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
