package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 uint }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint64, p1 int16, p2 uint8) string {
    if true {
        for fi4 := 0; fi4 < 13; fi4++ {
            v5 := []int{}
            if true {
                p2 -= p2
                p1 += (-(p1))
            }
            if (!false) {
                v6 := GenS1{F0: 1}
                _ = v6
            }
            _ = fi4
            _ = v5
        }
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: 255}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    p1 += 0
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}
func genFn1(p0 uint8, p1 uint64, p2 bool) int {
    v10 := GenS1{F0: 795}
    switch 9223372036854775807 {
    case 42:
        p2 = p2
    case -1:
        if (-566 != -1) {
            var v11 int8 = 1
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 32767
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 32767
    v3 := make(map[int]bool)
    if (true && true) {
        v1 = (v1 - 2)
        if (false && (47 < 128)) {
            var v12 uint8 = 128
            v13 := GenS1{F0: 32767}
            var v14 uint8 = v12
            _ = v12
            _ = v13
            _ = v14
        }
    } else {
        if false {
            for fi15 := 0; fi15 < 12; fi15++ {
                if (418 <= -2147483648) {
                    v2 = v2
                    var v16 int16 = -32768
                    _ = v16
                } else {
                    v2 = v2
                }
                v17 := GenS2{F0: 1, F1: 0}
                var v18 uint32 = 422
                _ = fi15
                _ = v17
                _ = v18
            }
        } else {
            v3[42] = false
        }
        if false {
            if false {
                var v19 int16 = 32767
                v2 += (v2 * 255)
                v20 := GenS3{F0: (^(v19)), F1: 42, F2: 276}
                _ = v19
                _ = v20
            }
        }
    }
    v1 = v1
    v2 = v2
    v1 += v1
    var v21 bool = (3.14159 < -2.25)
    var v22 float64 = (1e10 - (0.1 + 0.0))
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 float32 = ((-2.25 / 1.0) + 0.0)
        v2 -= v2
        var v25 int32 = int32(fi23)
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        fmt.Println(v21, v21)
        fmt.Println(v2, v21)
        _ = fi26
    }
    if (!v21) {
        v27 := []int{42}
        var v28 string = ""
        v29 := GenS2{F0: -1, F1: 1}
        _ = v27
        _ = v28
        _ = v29
    }
    v22 = v22
    v30 := GenS1{F0: 702}
    v22 = v22
    if (2 >= 162) {
        for fi31 := 0; fi31 < 7; fi31++ {
            if v21 {
                for fi32 := 0; fi32 < 2; fi32++ {
                    v3[fi32] = v21
                    _ = fi32
                }
                for fi33 := 0; fi33 < 8; fi33++ {
                    switch fi33 {
                    case -1:
                        v34 := GenS4{F0: 1}
                        _ = v34
                    default:
                        var v35 int8 = -128
                        _ = v35
                    }
                    _ = fi33
                }
            } else {
                var v36 uint8 = 42
                switch ("hello" + "a") {
                case "a":
                    var v37 int8 = 1
                    _ = v37
                case "\n\t":
                    var v38 int64 = -9223372036854775808
                    _ = v38
                case "foo bar":
                    v36 -= v36
                default:
                    var v39 uint8 = v36
                    _ = v39
                }
                _ = v36
            }
            _ = fi31
        }
    }
    if true {
        var v40 float64 = ((v22 * v22) - v22)
        if (255 <= v2) {
            if ((4294967295 >= 969) && false) {
                v22 -= v40
                var v41 int64 = 404
                if (-1 == 118) {
                    v2 += 255
                }
                _ = v41
            }
        } else {
            var v42 uint64 = 302
            switch ("" + ("hello" + "foo bar")) {
            case "foo bar":
                v1 += (v42 % (((^(v42))) | 1))
            case "\n\t":
                v43 := GenS3{F0: -32768, F1: -128, F2: 42}
                _ = v43
            default:
                v21 = v21
            }
            v40 -= v40
            _ = v42
        }
        v2 += v2
        _ = v40
    }
    v22 -= v22
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
}
