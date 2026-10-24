package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int32; F2 int8 }

func main() {
    var v1 float32 = -2.25
    var v2 uint = 1
    var v3 uint32 = 2
    var v4 uint32 = (v3 + (v3 / (((^(v3))) | 1)))
    var v5 float64 = ((1e10 - float64(v1)) + (((float64(((v1 - 3.14159) - 3.14159)) / ((0.1 / ((0.0 / 0.1) + 1e10)) / 0.1)) * 1e10) / -2.25))
    var v6 uint = 1
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    fmt.Println(v1, v3, v6)
    v8 := GenS1{F0: ((!false) || true), F1: v2}
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS2{F0: (671 <= 2), F1: v4, F2: v4}
        _ = fi9
        _ = v10
    }
    fmt.Println(v4, v6)
    var v11 float64 = v5
    if ((v6 + v6) > v2) {
        var v12 uint = v2
        fmt.Println(v12, v2)
        v13 := GenS3{F0: 67, F1: 0, F2: 948}
        _ = v12
        _ = v13
    } else {
        switch (("a" + "\n\t") + "hello") {
        case "foo bar":
            if (658 != 0) {
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 uint32 = (v3 * uint32(v2))
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := GenS1{F0: false, F1: 1}
                    v1 = float32(v5)
                    fmt.Println(v11, v11)
                    _ = fi16
                    _ = v17
                }
            } else {
                for fi18 := 0; fi18 < 4; fi18++ {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v20 := []int{int(v3), 762}
                        var v21 int32 = 250
                        fmt.Println(v2, v5, fi18)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    if (!(-332 <= 741)) {
                        var v22 int32 = 139
                        _ = v22
                    } else {
                        var v23 uint = v2
                        _ = v23
                    }
                    if false {
                        v24 := GenS2{F0: (!true), F1: (^(v4)), F2: (v3 - (v4 | v3))}
                        _ = v24
                    } else {
                        v11 = (-2.25 * (v11 - 0.1))
                    }
                    _ = fi18
                }
            }
        case "\n\t":
            for fi25 := 0; fi25 < 13; fi25++ {
                var v26 uint8 = 171
                _ = fi25
                _ = v26
            }
        case "a":
            for fi27 := 0; fi27 < 3; fi27++ {
                v28 := []int{(fi27 - 838)}
                var v29 string = "hello"
                switch 0 {
                case 32768:
                    if (42 != 0) {
                        switch -57 {
                        case 32768:
                            if len(v28) > 0 {
                                v28[0] = (fi27 >> ((v2) & 63))
                            }
                        }
                    }
                case 127:
                    v6 = (v2 ^ v6)
                }
                _ = fi27
                _ = v28
                _ = v29
            }
        default:
            var v30 int32 = 42
            _ = v30
        }
        v6 = (v6 / ((359) | 1))
        if true {
            switch 4294967295 {
            case -359:
                var v31 int64 = 9223372036854775807
                _ = v31
            case -83:
                for fi32 := 0; fi32 < 3; fi32++ {
                    v33 := []int{-637}
                    v2 -= v6
                    switch -9223372036854775808 {
                    case -1:
                        var v34 uint16 = 0
                        _ = v34
                    case 128:
                        v11 = (v11 * v5)
                    case 0:
                        v35 := GenS2{F0: (((^(v3)) >= 1) && (false && false)), F1: v4, F2: v3}
                        _ = v35
                    default:
                        v6 = 18446744073709551615
                    }
                    _ = fi32
                    _ = v33
                }
            case -2147483648:
                v36 := []int{-87, 0}
                _ = v36
            default:
                var v37 float64 = -2.25
                _ = v37
            }
            for fi38 := 0; fi38 < 1; fi38++ {
                var v39 uint32 = v4
                v4 = v3
                _ = fi38
                _ = v39
            }
        }
    }
    var v40 int64 = 0
    if (-2147483648 != 2147483647) {
        var v41 uint64 = 42
        v42 := GenS4{F0: 2, F1: int32(v41), F2: -128}
        v11 = float64(v1)
        _ = v41
        _ = v42
    } else {
        switch -1 {
        case 42:
            for fi43 := 0; fi43 < 18; fi43++ {
                v44 := []int{}
                var v45 int64 = -1
                var v46 uint32 = 65535
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
        case -1:
            v47 := []int{2147483647, 9223372036854775807, 32767}
            _ = v47
        case 9223372036854775807:
            v48 := []int{498, 42, -2147483648, 0}
            _ = v48
        }
    }
    v49 := []int{0, 298, 4294967295}
    v40 -= v40
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v49
    _ = v5
    _ = v6
    _ = v8
}
