package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 int }
type GenS4 struct { F0 uint32; F1 uint; F2 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 uint; F2 uint32 }
type GenS8 struct { F0 int8 }
func genFn0(p0 int8, p1 int64, p2 uint16) string {
    var v4 int = 0
    if true {
        fmt.Println(v4)
        p1 += 1
        var v5 string = ("foo bar" + ("a" + ("hello" + ("" + "foo bar"))))
        _ = v5
    }
    fmt.Println(p0, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "a"
}
func genFn1() uint8 {
    v6 := []int{32768, -2147483648, 127}
    switch -2147483648 {
    case 128:
        var v7 float64 = 0.0
        _ = v7
    case -1:
        var v8 int = -1
        _ = v8
    }
    v6 = append(v6, 877)
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint64 = 966
        fmt.Println(v10, fi9)
        _ = fi9
        _ = v10
    }
    _ = v6
    return 127
}

func main() {
    var v1 uint32 = 1
    var v2 uint8 = 0
    v3 := []uint32{v1, 88, 820, 2147483648}
    v3 = append(v3, 2)
    var v11 int16 = 42
    for fi12 := 0; fi12 < 3; fi12++ {
        if (-1 != 9223372036854775807) {
            if (!false) {
                v13 := GenS1{F0: v2}
                var v14 int16 = (v11 % ((-303) | 1))
                _ = v13
                _ = v14
            }
            if false {
                switch "foo bar" {
                case "foo bar":
                    if false {
                        v15 := GenS2{F0: 18446744073709551615, F1: true, F2: 1}
                        fmt.Println(v2, fi12, fi12)
                        _ = v15
                    } else {
                        var v16 int64 = -1
                        _ = v16
                    }
                case "a":
                    v11 = 32767
                case "\n\t":
                    if true {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            var v18 string = "a"
                            var v19 uint64 = 42
                            var v20 float32 = 0.1
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    } else {
                        v2 += v2
                        var v21 uint = 1
                        for fi22 := 0; fi22 < 13; fi22++ {
                            fmt.Println(fi22)
                            var v23 int = (fi12 << ((v21) & 63))
                            if len(v3) > 0 {
                                v3[0] = v1
                            }
                            _ = fi22
                            _ = v23
                        }
                        _ = v21
                    }
                default:
                    v3 = append(v3, (v1 << ((1) & 63)))
                }
                var v24 int8 = 100
                _ = v24
            }
            v2 = 92
        } else {
            v25 := GenS1{F0: (v2 ^ v2)}
            v1 += uint32(v2)
            var v26 uint16 = 412
            _ = v25
            _ = v26
        }
        _ = fi12
    }
    fmt.Println(v11, v2)
    v27 := GenS2{F0: 42, F1: false, F2: 1}
    var v28 string = "\n\t"
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 uint64 = 18446744073709551615
        v2 -= v2
        var v31 float32 = float32(0.1)
        _ = fi29
        _ = v30
        _ = v31
    }
    v1 += v1
    if true {
        if (2 == 1) {
            var v32 uint64 = 0
            v32 = (v32 & uint64(v2))
            _ = v32
        }
        if (("\n\t" + "foo bar") > "\n\t") {
            if true {
                v11 -= (v11 << ((4294967295) & 63))
                if true {
                    _ = len(v3)
                    v2 -= 127
                    if false {
                        var v33 bool = false
                        _ = v33
                    }
                }
                for fi34 := 0; fi34 < 18; fi34++ {
                    if (!(((v28 + "foo bar") + ("\n\t" + v28)) >= "hello")) {
                        if false {
                            if len(v3) > 0 {
                                v3[0] = v1
                            }
                            if (0 != v1) {
                                v1 = v1
                                if (true && false) {
                                    v35 := GenS3{F0: v2, F1: int16(fi34), F2: fi34}
                                    switch v28 {
                                    case "":
                                        v36 := GenS3{F0: (v2 * (v2 >> ((383) & 63))), F1: v11, F2: (-(fi34))}
                                        _ = v36
                                    default:
                                        v1 += (v1 ^ v1)
                                    }
                                    switch (v28 + "foo bar") {
                                    case "\n\t":
                                        v37 := GenS4{F0: v1, F1: 18446744073709551615, F2: ((18446744073709551615 <= 1) || (v1 <= (v1 - uint32(v2))))}
                                        _ = v37
                                    }
                                    _ = v35
                                }
                                for fi38 := 0; fi38 < 17; fi38++ {
                                    v39 := GenS2{F0: 61, F1: false, F2: 343}
                                    _ = fi38
                                    _ = v39
                                }
                            } else {
                                v28 = (v28 + "")
                                v40 := GenS5{F0: 42}
                                _ = v40
                            }
                            v41 := GenS6{F0: (4294967295 >= -80)}
                            _ = v41
                        } else {
                            fmt.Println(v11)
                        }
                        var v42 int = fi34
                        v11 += -32768
                        _ = v42
                    }
                    _ = fi34
                }
            }
            fmt.Println(v2, v11)
            var v43 string = (("foo bar" + (v28 + ((v28 + v28) + "a"))) + ("hello" + v28))
            _ = v43
        }
        v11 += v11
    }
    for fi44 := 0; fi44 < 18; fi44++ {
        if (1 < -570) {
            v45 := GenS2{F0: 1, F1: (!(fi44 == 939)), F2: 9223372036854775807}
            _ = v45
        }
        switch v28 {
        case "foo bar":
            if true {
                v46 := GenS7{F0: fi44, F1: 379, F2: (v1 & uint32(v2))}
                _ = v46
            }
        }
        _ = fi44
    }
    switch 32768 {
    case 4294967295:
        v1 += v1
    case 42:
        if (1.5 <= (0.0 + float32((0.0 - (3.14159 + (3.14159 - (0.0 / 1e10))))))) {
            for fi47 := 0; fi47 < 5; fi47++ {
                v1 = v1
                fmt.Println(v28, v28, fi47)
                _ = fi47
            }
            var v48 uint8 = (^(v2))
            _ = v48
        } else {
            fmt.Println(v11, v28, v11)
        }
    }
    if (2147483647 < 1) {
        v28 = v28
    } else {
        var v49 int16 = (^(v11))
        _ = v49
    }
    switch ("a" + v28) {
    case "foo bar":
        for fi50 := 0; fi50 < 6; fi50++ {
            switch "\n\t" {
            case "":
                for fi51 := 0; fi51 < 16; fi51++ {
                    if (true && true) {
                        var v52 bool = ((v28 + v28) >= (("\n\t" + v28) + "\n\t"))
                        if v52 {
                            v11 = 0
                            v11 = (v11 << ((uint(v1)) & 63))
                        }
                        v53 := GenS8{F0: 8}
                        _ = v52
                        _ = v53
                    } else {
                        var v54 float64 = -2.25
                        if (!(uint64(v2) != 1)) {
                            fmt.Println(v1, fi51)
                            switch v28 {
                            case "\n\t":
                                var v55 int = 127
                                _ = v55
                            }
                        } else {
                            switch fi51 {
                            case 9223372036854775807:
                                v2 = (v2 | (v2 + genFn1()))
                            case -29:
                                v3 = append(v3, (^(v1)))
                            }
                        }
                        if (0 > v1) {
                            fmt.Println(fi50)
                        }
                        _ = v54
                    }
                    switch fi50 {
                    case -9223372036854775808:
                        fmt.Println(v1)
                    case -383:
                        v11 -= 564
                    case 0:
                        switch (v28 + "hello") {
                        case "foo bar":
                            if (!false) {
                                fmt.Println(v28, fi50)
                                fmt.Println(fi51)
                                var v56 uint32 = v1
                                var v57 int32 = int32(v56)
                                v57 = (v57 & v57)
                                _ = v56
                                _ = v57
                            } else {
                                fmt.Println(v1, v1, v11)
                                var v58 bool = (!(127 <= 127))
                                _ = v58
                            }
                        case "hello":
                            var v59 string = v28
                            _ = v59
                        default:
                            v3 = append(v3, (v1 << ((0) & 63)))
                        }
                    }
                    v2 += (v2 * (v2 + (v2 >> ((42) & 63))))
                    _ = fi51
                }
            }
            var v60 int16 = v11
            fmt.Println(v60, fi50, v1)
            _ = fi50
            _ = v60
        }
    case "":
        v61 := GenS3{F0: v2, F1: 344, F2: 1}
        _ = v61
    }
    v11 -= v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
