package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 bool; F2 bool }
func genFn0() uint8 {
    if false {
        v6 := GenS1{F0: 100}
        v7 := GenS1{F0: -1}
        _ = v6
        _ = v7
    } else {
        var v8 uint8 = 128
        var v9 int = -1
        v9 += (v9 / (((v9 >> ((134) & 63))) | 1))
        _ = v8
        _ = v9
    }
    switch 4294967295 {
    case -1:
        v10 := []int{56, 390}
        _ = v10
    case -46:
        var v11 uint64 = 442
        _ = v11
    default:
        if (!(true || true)) {
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 int64 = -9223372036854775808
                for fi14 := 0; fi14 < 3; fi14++ {
                    v15 := GenS2{F0: 0}
                    v13 = (v13 * v13)
                    var v16 int64 = (v13 | v13)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
            v17 := GenS1{F0: 1}
            _ = v17
        } else {
            var v18 string = ("hello" + "")
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 uint32 = 2
                _ = fi19
                _ = v20
            }
            _ = v18
        }
    }
    return 239
}
func genFn1(p0 int64) uint8 {
    p0 = 0
    p0 = p0
    switch "foo bar" {
    case "a":
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 int16 = 42
            for fi23 := 0; fi23 < 17; fi23++ {
                if (!false) {
                    v22 = 1000
                } else {
                    var v24 uint8 = 2
                    _ = v24
                }
                var v25 float64 = ((3.14159 + 1e10) + float64(float32(0.1)))
                _ = fi23
                _ = v25
            }
            _ = fi21
            _ = v22
        }
    case "":
        p0 += 114
    case "foo bar":
        for fi26 := 0; fi26 < 14; fi26++ {
            switch 128 {
            case -53:
                if false {
                    var v27 uint64 = 579
                    var v28 uint32 = 2147483647
                    var v29 int32 = 42
                    _ = v27
                    _ = v28
                    _ = v29
                }
            }
            fmt.Println(p0)
            _ = fi26
        }
    }
    v30 := GenS1{F0: 127}
    _ = p0
    _ = v30
    return 1
}

func main() {
    v1 := []int{-9223372036854775808}
    v2 := []uint32{2147483647}
    v3 := []uint32{}
    var v4 float64 = (0.0 / 1e10)
    var v5 bool = false
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    if (0 == 0) {
        v5 = true
        fmt.Println(v5, v5)
        v31 := GenS2{F0: 11}
        _ = v31
    } else {
        fmt.Println(v5, v5)
        var v32 bool = (42 != 4294967295)
        _ = v32
    }
    var v33 uint64 = 109
    var v34 int8 = 21
    var v35 int8 = (v34 - v34)
    var v36 uint8 = genFn1(9223372036854775807)
    if v5 {
        var v37 int16 = 1
        _ = v37
    } else {
        v38 := GenS3{F0: v36, F1: ((^(v36)) < (v36 >> ((434) & 63))), F2: true}
        _ = v38
    }
    var v39 float64 = v4
    var v40 string = "hello"
    v41 := GenS2{F0: v35}
    fmt.Println(v4)
    if (0 == 530) {
        for fi42 := 0; fi42 < 2; fi42++ {
            v39 -= (float64(0.1) - 3.14159)
            var v43 int16 = 42
            _ = fi42
            _ = v43
        }
        var v44 int = -2147483648
        _ = v44
    } else {
        for fi45 := 0; fi45 < 13; fi45++ {
            var v46 float64 = -2.25
            _ = fi45
            _ = v46
        }
    }
    if ((1.5 != (1.5 + 3.14159)) && (2147483647 == -1)) {
        for fi47 := 0; fi47 < 16; fi47++ {
            v48 := GenS3{F0: (v36 ^ v36), F1: false, F2: ((v34 * v34) <= (v35 / ((-1) | 1)))}
            var v49 int8 = v35
            var v50 int8 = (v35 & v35)
            _ = fi47
            _ = v48
            _ = v49
            _ = v50
        }
        for fi51 := 0; fi51 < 17; fi51++ {
            v52 := GenS3{F0: (v36 << ((0) & 63)), F1: ((v40 + ("" + ((v40 + "\n\t") + v40))) < v40), F2: false}
            var v53 bool = v5
            v4 = v4
            _ = fi51
            _ = v52
            _ = v53
        }
        switch v40 {
        case "":
            if len(v3) > 0 {
                v3[0] = 718
            }
        case "a":
            v54 := GenS1{F0: v35}
            _ = v54
        default:
            switch 277 {
            case 9223372036854775807:
                v55 := GenS3{F0: 128, F1: (!((v36 & 255) <= 128)), F2: v5}
                _ = v55
            case 130:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v56 bool = true
        var v57 int32 = 372
        _ = v56
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
