package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 bool; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint64) int8 {
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            p1 -= p1
            p0 = (p0 / (((p0 << ((1) & 63))) | 1))
            p0 -= (p0 % ((p0) | 1))
            _ = fi7
        }
        var v8 bool = (!true)
        _ = fi6
        _ = v8
    }
    var v9 int8 = 0
    v10 := GenS1{F0: 0}
    for fi11 := 0; fi11 < 18; fi11++ {
        switch fi11 {
        case -2147483648:
            var v12 uint64 = (p1 << ((255) & 63))
            _ = v12
        case -994:
            var v13 int32 = 42
            _ = v13
        }
        _ = fi11
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (v9 + v9)
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    var v3 uint16 = 590
    var v4 uint8 = 1
    var v5 int8 = int8(v4)
    _ = len(v2)
    var v14 string = "a"
    v15 := GenS2{F0: 9223372036854775807}
    var v16 int32 = -23
    v14 = v14
    if (v1 && v1) {
        var v17 float32 = 3.14159
        var v18 int64 = int64(v4)
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 14; fi19++ {
            v14 = v14
            var v20 uint64 = 0
            v14 = v14
            _ = fi19
            _ = v20
        }
        fmt.Println(v16, v4, v14)
    }
    switch v14 {
    case "a":
        var v21 float32 = float32((-2.25 - (1.5 - (float64(0.1) * (-2.25 + float64((1.5 / float32(1.5))))))))
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        if (fi22 == int(v3)) {
            for fi23 := 0; fi23 < 19; fi23++ {
                fmt.Println(v4, v4, v5)
                _ = fi23
            }
            switch fi22 {
            case 127:
                v24 := v2[32768]
                _ = v24
            case 768:
                var v25 int = fi22
                _ = v25
            case -1:
                v26 := GenS3{F0: 76}
                _ = v26
            }
            v27 := GenS4{F0: 980, F1: 269}
            _ = v27
        }
        var v28 int64 = 619
        _ = fi22
        _ = v28
    }
    var v29 float64 = 1e10
    fmt.Println(v14)
    if v1 {
        var v30 uint32 = 65535
        v14 = v14
        _ = v30
    } else {
        var v31 float32 = float32(v29)
        var v32 bool = (int64(v16) >= 9223372036854775807)
        _ = v31
        _ = v32
    }
    if true {
        v1 = (-473 > (v16 ^ v16))
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        switch (v14 + v14) {
        case "foo bar":
            v34 := GenS5{F0: v1, F1: (v16 | (^(v16))), F2: 42}
            _ = v34
        }
        _ = fi33
    }
    if (v1 || v1) {
        if (v1 && (((v1 || (3.14159 >= (0.0 / 1.0))) || (2147483647 >= -664)) && true)) {
            if v1 {
                var v35 uint64 = 4294967295
                _ = v35
            } else {
                if (((v16 | v16) >= v16) || v1) {
                    fmt.Println(v29, v16, v3)
                } else {
                    v36 := []int{}
                    var v37 string = ""
                    v37 = v14
                    _ = v36
                    _ = v37
                }
                switch ("foo bar" + v14) {
                case "hello":
                    v38 := GenS4{F0: 65535, F1: 0}
                    _ = v38
                default:
                    switch 128 {
                    case 1:
                        var v39 int8 = int8(v16)
                        _ = v39
                    default:
                        v40 := GenS6{F0: (v3 & (^(v3)))}
                        _ = v40
                    }
                }
                v14 = "hello"
            }
            v41 := []int{2147483647, 796, 4294967295}
            var v42 float64 = float64((3.14159 / 1e10))
            _ = v41
            _ = v42
        } else {
            v43 := GenS3{F0: -1}
            for fi44 := 0; fi44 < 13; fi44++ {
                var v45 int8 = v5
                _ = fi44
                _ = v45
            }
            _ = v43
        }
    } else {
        var v46 float64 = v29
        switch "" {
        case "":
            v29 += ((float64(-2.25) - v29) + float64(0.1))
        case "\n\t":
            var v47 int16 = int16(v4)
            _ = v47
        case "a":
            v16 = int32(v4)
        default:
            for fi48 := 0; fi48 < 15; fi48++ {
                var v49 float64 = ((1.5 * (((v29 - v29) / v46) - v46)) - v29)
                switch (v14 + ("a" + v14)) {
                case "":
                    if v1 {
                        if v1 {
                            v50 := GenS5{F0: (0 != 18446744073709551615), F1: 0, F2: 42}
                            v29 = 0.0
                            for fi51 := 0; fi51 < 15; fi51++ {
                                fmt.Println(v49)
                                var v52 uint64 = 18446744073709551615
                                var v53 uint16 = 974
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                            _ = v50
                        } else {
                            v54 := v2[fi48]
                            _ = v54
                        }
                    }
                default:
                    v5 += (-(v5))
                }
                _ = fi48
                _ = v49
            }
        }
        _ = v46
    }
    var v55 uint32 = 0
    for fi56 := 0; fi56 < 1; fi56++ {
        v5 = (^(v5))
        switch v14 {
        case "":
            switch ("a" + "\n\t") {
            case "foo bar":
                v57 := GenS5{F0: v1, F1: v16, F2: v55}
                _ = v57
            case "":
                var v58 int = fi56
                _ = v58
            }
        }
        _ = fi56
    }
    v2[-2147483648] = 1
    v59 := v2[-262]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v59
}
