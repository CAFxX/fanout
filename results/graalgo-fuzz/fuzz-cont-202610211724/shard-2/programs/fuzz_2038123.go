package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int32, p1 int64, p2 uint8) int64 {
    if false {
        for fi4 := 0; fi4 < 7; fi4++ {
            if false {
                if true {
                    v5 := []int{-601}
                    _ = v5
                }
                p0 = p0
            }
            for fi6 := 0; fi6 < 14; fi6++ {
                if (p2 > (p2 + (p2 % ((p2) | 1)))) {
                    fmt.Println(p1, p2)
                    var v7 uint16 = 511
                    var v8 uint64 = 286
                    _ = v7
                    _ = v8
                }
                v9 := GenS1{F0: (p2 * 127), F1: (!(127 == int8(fi4))), F2: 0}
                v10 := []int{(-(fi4)), fi4, 341, fi6}
                _ = fi6
                _ = v10
                _ = v9
            }
            v11 := GenS2{F0: 953}
            _ = fi4
            _ = v11
        }
        var v12 int = 908
        for fi13 := 0; fi13 < 19; fi13++ {
            if (-1 < 1) {
                var v14 float32 = -2.25
                v15 := []int{fi13, (fi13 ^ fi13), (fi13 << ((961) & 63))}
                v16 := GenS2{F0: 42}
                _ = v14
                _ = v15
                _ = v16
            }
            p1 = (^(p1))
            _ = fi13
        }
        _ = v12
    } else {
        var v17 int8 = int8(p2)
        _ = v17
    }
    v18 := GenS3{F0: 32767}
    switch 0 {
    case 127:
        var v19 uint32 = 4294967295
        _ = v19
    case -850:
        p2 -= p2
    case -529:
        var v20 int16 = -46
        _ = v20
    default:
        var v21 int32 = (p0 >> ((705) & 63))
        _ = v21
    }
    var v22 int8 = 61
    v23 := GenS4{F0: false, F1: (p1 | 1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v22
    _ = v23
    return 9223372036854775807
}
func genFn1(p0 uint64) string {
    for fi24 := 0; fi24 < 1; fi24++ {
        v25 := GenS3{F0: 1}
        _ = fi24
        _ = v25
    }
    if true {
        var v26 uint8 = 202
        _ = v26
    }
    var v27 uint64 = (p0 - (p0 >> ((42) & 63)))
    _ = p0
    _ = v27
    return "a"
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 42
    var v3 int = 42
    v28 := GenS1{F0: 247, F1: ((v3 + (v3 | 0)) >= v3), F2: 42}
    var v29 uint32 = 65535
    v30 := GenS1{F0: 2, F1: ((!true) && (!true)), F2: (v2 * (v2 ^ (v2 & 456)))}
    var v31 uint32 = (v29 * 4294967295)
    for fi32 := 0; fi32 < 16; fi32++ {
        v33 := v1[("foo bar" + "foo bar")]
        v33 = "foo bar"
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        v3 -= (-(fi34))
        v35 := GenS5{F0: -1, F1: v29, F2: 18446744073709551615}
        _ = fi34
        _ = v35
    }
    var v36 uint32 = (^(v31))
    switch ("hello" + "a") {
    case "\n\t":
        switch 9223372036854775807 {
        case 42:
            var v37 uint16 = (^(v2))
            _ = v37
        case 127:
            for fi38 := 0; fi38 < 20; fi38++ {
                if true {
                    v39 := GenS6{F0: 127}
                    var v40 int = v3
                    _ = v39
                    _ = v40
                } else {
                    v31 += v36
                    var v41 uint16 = (^(v2))
                    fmt.Println(v36, fi38)
                    _ = v41
                }
                v36 = 212
                _ = fi38
            }
        case 355:
            if true {
                v2 = 361
                v42 := GenS4{F0: (-509 <= 787), F1: 9223372036854775807}
                _ = v42
            } else {
                v2 += v2
                v31 -= v29
            }
        }
    }
    v43 := GenS3{F0: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v43
}
