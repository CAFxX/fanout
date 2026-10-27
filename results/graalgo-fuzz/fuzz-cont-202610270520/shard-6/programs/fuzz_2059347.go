package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int }
func genFn0(p0 uint16, p1 uint8, p2 uint) int8 {
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: false, F1: -9223372036854775808}
        if (604 <= 2) {
            if true {
                for fi7 := 0; fi7 < 14; fi7++ {
                    if (74 != 225) {
                        v8 := GenS1{F0: (0 >= 264), F1: 0}
                        var v9 uint8 = uint8(p0)
                        _ = v8
                        _ = v9
                    }
                    _ = fi7
                }
            } else {
                v10 := GenS2{F0: (fi5 % (((fi5 & fi5)) | 1)), F1: "foo bar"}
                if (-1 <= int(p0)) {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        var v12 bool = (1e10 <= (0.0 * (0.0 + (float32(1e10) * 0.0))))
                        v13 := []int{int(p0), (fi5 >> ((p2) & 63)), fi5, fi11}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                var v14 int16 = 0
                _ = v10
                _ = v14
            }
        }
        var v15 float32 = 1e10
        _ = fi5
        _ = v15
        _ = v6
    }
    p1 += (p1 & 2)
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := []int{(-(fi16)), (fi16 + fi16), 32768, 128}
        _ = fi16
        _ = v17
    }
    if (!false) {
        v18 := []int{}
        v19 := GenS1{F0: (0 <= p0), F1: -1}
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}
func genFn1(p0 int8, p1 uint64, p2 int8) uint {
    fmt.Println(p2, p0, p2)
    v20 := GenS1{F0: true, F1: 9223372036854775807}
    v21 := []int{0, -9223372036854775808, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 65535
}

func main() {
    var v1 uint64 = 42
    var v2 uint = 1
    var v3 int8 = -128
    v4 := make(map[uint32]uint32)
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 13; fi23++ {
            fmt.Println(v3, fi23, v2)
            switch "" {
            case "":
                fmt.Println(v3, fi23)
            case "foo bar":
                var v24 uint32 = 974
                _ = v24
            default:
                v1 -= (v1 * 4294967295)
            }
            _ = fi23
        }
        v25 := GenS1{F0: (1e10 >= -2.25), F1: -9223372036854775808}
        var v26 string = "\n\t"
        _ = fi22
        _ = v25
        _ = v26
    }
    if true {
        v27 := GenS1{F0: true, F1: -414}
        switch 1 {
        case 799:
            var v28 uint64 = (v1 / ((0) | 1))
            _ = v28
        case -2147483648:
            v3 -= (-(v3))
        }
        v2 -= v2
        _ = v27
    } else {
        v1 = (v1 << ((v2) & 63))
    }
    v2 = uint(v1)
    v2 = (v2 << ((65535) & 63))
    var v29 int = 32767
    var v30 uint16 = 2
    for fi31 := 0; fi31 < 2; fi31++ {
        for fi32 := 0; fi32 < 14; fi32++ {
            v33 := []int{}
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    v34 := []int{(v29 << (((v2 + (v2 >> (((v2 & 1)) & 63)))) & 63)), (v29 & (^(v29))), 128, -1}
    v3 = v3
    var v35 uint64 = (v1 + v1)
    v36 := GenS1{F0: false, F1: 42}
    v1 = 1
    v29 = (v29 % (((v29 + -9223372036854775808)) | 1))
    switch (v29 * (v29 | (v29 * (v29 + 1)))) {
    case -2147483648:
        var v37 float64 = 0.1
        _ = v37
    case 127:
        for fi38 := 0; fi38 < 19; fi38++ {
            if true {
                v2 = (^(v2))
                var v39 uint64 = (v1 << ((genFn1((^(v3)), 0, (v3 ^ (v3 | (v3 + v3))))) & 63))
                _ = v39
            } else {
                for fi40 := 0; fi40 < 12; fi40++ {
                    _ = len(v34)
                    fmt.Println(fi40, fi40)
                    var v41 int16 = 1
                    _ = fi40
                    _ = v41
                }
                v42 := GenS1{F0: (!false), F1: -9223372036854775808}
                _ = v42
            }
            _ = fi38
        }
    case -1:
        delete(v4, 577)
    }
    var v43 int = (v29 * (v29 | (v29 << (((^(v2))) & 63))))
    v44 := GenS2{F0: v43, F1: "\n\t"}
    fmt.Println(v30, v29)
    for fi45 := 0; fi45 < 14; fi45++ {
        switch (v29 << (((v2 >> (((v2 & v2)) & 63))) & 63)) {
        case 832:
            v46 := GenS3{F0: v29}
            _ = v46
        default:
            var v47 int32 = 569
            _ = v47
        }
        _ = len(v4)
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v44
}
