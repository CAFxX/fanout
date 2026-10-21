package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 int64; F2 bool }
func genFn0(p0 int16) int16 {
    switch 0 {
    case 9223372036854775807:
        if false {
            if (!false) {
                p0 += -32768
                var v7 int8 = 0
                var v8 uint32 = 2147483647
                _ = v7
                _ = v8
            }
            v9 := []int{128, -9223372036854775808, 32768, 32768}
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 int64 = 0
                v12 := GenS1{F0: -1}
                for fi13 := 0; fi13 < 4; fi13++ {
                    if (1 <= 25) {
                        var v14 uint16 = 32767
                        _ = v14
                    } else {
                        v15 := GenS2{F0: 65535, F1: (p0 * p0), F2: 4294967295}
                        var v16 uint64 = 939
                        var v17 string = ((("" + "foo bar") + "a") + "\n\t")
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = fi13
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        } else {
            var v18 uint = 805
            v18 = (v18 - v18)
            _ = v18
        }
    case 1:
        var v19 int16 = 222
        _ = v19
    default:
        p0 = -184
    }
    v20 := GenS1{F0: -1}
    v21 := []int{127}
    _ = p0
    _ = v20
    _ = v21
    return -32768
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = 2147483647
    var v3 int32 = (-(v2))
    var v4 uint64 = v1
    v5 := []uint32{2147483648, 2147483648, 510}
    v6 := make(map[int]bool)
    if (false || (42 == 829)) {
        switch ("" + "a") {
        case "foo bar":
            fmt.Println(v2, v1)
        case "hello":
            var v22 uint8 = 10
            _ = v22
        }
    }
    v4 += v1
    var v23 float32 = -2.25
    var v24 uint16 = 602
    v24 -= 418
    var v25 int16 = 32767
    v26 := GenS3{F0: 0}
    var v27 uint = 42
    var v28 int8 = -128
    var v29 uint16 = v24
    _ = len(v6)
    var v30 int64 = 1
    v29 += v24
    for fi31 := 0; fi31 < 15; fi31++ {
        if ((!((true || (!((genFn0(v25) < (^(v25))) && false))) || ((v30 & 1) != (v30 * (-(v30)))))) && (1e10 == (float64(1.5) / ((float64(((1.5 + ((v23 / (0.1 * v23)) * (((v23 - ((((float32(1e10) / (-2.25 / 1e10)) + float32(0.1)) + (v23 - 0.1)) / 3.14159)) + 0.0) * (0.1 - v23)))) / v23)) - 3.14159) + 3.14159)))) {
            var v32 uint16 = v29
            _ = v32
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi31
    }
    if true {
        var v33 int64 = v30
        var v34 int = -2147483648
        _ = v33
        _ = v34
    } else {
        v35 := GenS1{F0: v2}
        switch -315 {
        case 0:
            var v36 int32 = v3
            _ = v36
        case 245:
            switch 1 {
            case -57:
                v2 += (-(v3))
            case 401:
                for fi37 := 0; fi37 < 12; fi37++ {
                    switch (-(fi37)) {
                    case 221:
                        for fi38 := 0; fi38 < 9; fi38++ {
                            v39 := v6[fi37]
                            _ = fi38
                            _ = v39
                        }
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    _ = fi37
                }
            }
        case 804:
            if (true && (!true)) {
                v40 := GenS4{F0: 255, F1: v30, F2: true}
                v4 = v1
                _ = v40
            } else {
                v6[32768] = true
            }
        default:
            for fi41 := 0; fi41 < 13; fi41++ {
                var v42 bool = true
                var v43 float64 = ((-2.25 / ((-2.25 - 1.5) / 1.0)) * 3.14159)
                v30 -= v30
                _ = fi41
                _ = v42
                _ = v43
            }
        }
        _ = v35
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
