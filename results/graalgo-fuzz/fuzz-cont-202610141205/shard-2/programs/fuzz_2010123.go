package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint8; F2 uint }
type GenS4 struct { F0 string }
func genFn0(p0 uint64, p1 int64, p2 uint16) int16 {
    p1 -= 42
    var v5 int = 441
    var v6 uint32 = 31
    v5 -= -2147483648
    v5 -= 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint64) string {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{(fi7 + (^(fi7))), fi7, -193}
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "a":
        var v9 int8 = 100
        _ = v9
    case "hello":
        p0 = p0
    case "":
        if (32767 <= -1) {
            switch 32768 {
            case 127:
                if (180 == -1) {
                    if (false || (873 <= 9223372036854775807)) {
                        switch -412 {
                        case -288:
                            var v10 int64 = -189
                            _ = v10
                        case 127:
                            var v11 bool = false
                            _ = v11
                        default:
                            var v12 uint = 651
                            _ = v12
                        }
                        var v13 float64 = (-2.25 / ((1.5 + ((0.0 - (3.14159 + float64((-2.25 / ((3.14159 * ((float32((float64(((float32(1.5) / 1.0) * 1e10)) - 1.5)) - float32(float64((-2.25 / 3.14159)))) + -2.25)) / 1.5))))) * 0.0)) / -2.25))
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 32767}
                        p0 = p0
                        var v15 int32 = -2147483648
                        _ = v14
                        _ = v15
                    }
                    if false {
                        for fi16 := 0; fi16 < 18; fi16++ {
                            var v17 bool = (255 == 1)
                            var v18 uint64 = p0
                            switch fi16 {
                            case 0:
                                if v17 {
                                    v19 := GenS1{F0: 1}
                                    _ = v19
                                } else {
                                    v17 = false
                                    p0 -= (p0 % (((v18 << ((255) & 63))) | 1))
                                    fmt.Println(fi16)
                                }
                            case 32768:
                                v20 := GenS1{F0: 42}
                                _ = v20
                            default:
                                var v21 float64 = 0.1
                                _ = v21
                            }
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                    v22 := []int{-1, -9223372036854775808, 42, 32767}
                    _ = v22
                }
            case -988:
                v23 := GenS2{F0: 805}
                _ = v23
            }
        } else {
            var v24 int64 = -9223372036854775808
            v25 := []int{}
            v26 := GenS1{F0: -996}
            _ = v24
            _ = v25
            _ = v26
        }
    }
    if ((18446744073709551615 == 4294967295) || false) {
        v27 := GenS2{F0: 627}
        _ = v27
    }
    _ = p0
    return "foo bar"
}

func main() {
    var v1 uint8 = 2
    var v2 uint32 = 917
    v3 := make(map[string]int)
    var v4 float32 = float32(float64(3.14159))
    for fi28 := 0; fi28 < 20; fi28++ {
        for fi29 := 0; fi29 < 13; fi29++ {
            v4 += (float32(-2.25) * (v4 * v4))
            v4 += (v4 / v4)
            v1 += (v1 << ((42) & 63))
            _ = fi29
        }
        var v30 bool = false
        _ = fi28
        _ = v30
    }
    if (v2 > 2147483648) {
        switch 127 {
        case 9223372036854775807:
            var v31 int = 0
            _ = v31
        case -405:
            v32 := []int{32768}
            _ = v32
        default:
            var v33 uint32 = (v2 >> ((18446744073709551615) & 63))
            _ = v33
        }
        v2 -= (^(v2))
    } else {
        switch 128 {
        case 42:
            fmt.Println(v1, v4)
        case -1:
            if (true && ((v1 / (((^(v1))) | 1)) <= (v1 | 178))) {
                v4 = 0.0
                var v34 float64 = ((((1e10 + (0.0 - ((3.14159 * float64(v4)) / (0.1 + 1e10)))) / -2.25) / 3.14159) - 3.14159)
                v4 = (0.0 * v4)
                _ = v34
            } else {
                v35 := GenS3{F0: -128, F1: v1, F2: 65535}
                _ = v35
            }
        }
        fmt.Println(v2)
    }
    if (v2 == v2) {
        var v36 uint32 = v2
        v1 += 111
        v36 += v2
        _ = v36
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        var v38 int16 = -32768
        _ = fi37
        _ = v38
    }
    v39 := GenS2{F0: 255}
    if ((-113 < -751) || (0.1 < float64(v4))) {
        v40 := v3[("a" + "foo bar")]
        _ = v40
    } else {
        v41 := GenS4{F0: "hello"}
        _ = v41
    }
    v42 := GenS4{F0: "hello"}
    var v43 float64 = (0.1 / 1.0)
    if (255 <= (^(v1))) {
        if (1.5 < v4) {
            switch -9223372036854775808 {
            case 9223372036854775807:
                v2 = (v2 & (v2 % ((1) | 1)))
            case 128:
                var v44 bool = (1 >= 42)
                _ = v44
            }
            var v45 int64 = 1
            _ = v45
        } else {
            var v46 int8 = -128
            _ = v46
        }
        var v47 int32 = 1
        _ = v47
    }
    if true {
        v48 := GenS3{F0: 0, F1: v1, F2: 589}
        _ = v48
    }
    delete(v3, "")
    var v49 int = 42
    v50 := GenS1{F0: -32768}
    for fi51 := 0; fi51 < 3; fi51++ {
        v52 := GenS3{F0: 38, F1: (v1 + v1), F2: 0}
        fmt.Println(v4)
        _ = fi51
        _ = v52
    }
    v2 = (v2 << ((969) & 63))
    v43 = ((float64(float32(1e10)) - v43) / 3.14159)
    v2 = (^(v2))
    for fi53 := 0; fi53 < 2; fi53++ {
        switch fi53 {
        case 32768:
            v54 := []int{-9223372036854775808, (fi53 | v49), fi53}
            _ = v54
        case -683:
            if (false && (-9223372036854775808 >= -9223372036854775808)) {
                v43 -= 1e10
                for fi55 := 0; fi55 < 13; fi55++ {
                    v56 := GenS3{F0: 42, F1: v1, F2: 255}
                    _ = fi55
                    _ = v56
                }
                if (42 != (v2 & (v2 % ((v2) | 1)))) {
                    var v57 int32 = 0
                    for fi58 := 0; fi58 < 6; fi58++ {
                        var v59 float64 = v43
                        switch "a" {
                        case "":
                            var v60 int64 = -1
                            _ = v60
                        }
                        _ = fi58
                        _ = v59
                    }
                    _ = v57
                }
            }
        default:
            v43 = v43
        }
        v1 = 100
        v2 = v2
        _ = fi53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v49
    _ = v50
}
