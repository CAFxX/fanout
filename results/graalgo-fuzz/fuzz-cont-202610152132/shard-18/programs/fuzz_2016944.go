package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint8; F1 bool; F2 int32 }
type GenS5 struct { F0 bool; F1 int; F2 int8 }
func genFn0() bool {
    var v5 string = ""
    v5 = ((((v5 + "foo bar") + v5) + "foo bar") + v5)
    v6 := GenS1{F0: 24}
    _ = v5
    _ = v6
    return (!true)
}

func main() {
    var v1 int8 = 107
    var v2 bool = (true || (!false))
    var v3 int32 = 784
    v4 := make(map[string]bool)
    if ((1 >= 4294967295) || v2) {
        for fi7 := 0; fi7 < 4; fi7++ {
            delete(v4, "foo bar")
            _ = fi7
        }
    }
    switch 0 {
    case -9223372036854775808:
        v8 := GenS1{F0: 117}
        _ = v8
    case -108:
        v9 := GenS2{F0: 987}
        _ = v9
    default:
        v10 := GenS1{F0: 736}
        _ = v10
    }
    if false {
        v11 := v4[""]
        if (!v2) {
            v12 := v4[((("a" + "foo bar") + "a") + ("foo bar" + "a"))]
            var v13 int32 = 511
            _ = v12
            _ = v13
        }
        v2 = (884 > 882)
        _ = v11
    } else {
        if (2147483647 != 1) {
            for fi14 := 0; fi14 < 3; fi14++ {
                _ = len(v4)
                v15 := GenS1{F0: 65535}
                var v16 int32 = v3
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 int16 = 1
            var v18 uint32 = 2147483647
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 19; fi19++ {
                fmt.Println(fi19, fi19, fi19)
                var v20 uint = 65535
                var v21 int64 = 9223372036854775807
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        v3 += 2147483647
    }
    v22 := []int{32767, 4294967295, 42}
    v1 = (v1 / ((v1) | 1))
    var v23 uint32 = 332
    var v24 int32 = (v3 << ((890) & 63))
    v22 = append(v22, 32767)
    v25 := GenS2{F0: v23}
    v3 = (v3 + (v24 % ((v24) | 1)))
    var v26 bool = (32768 > 128)
    var v27 float32 = (float32(((0.1 * 3.14159) * 1.5)) * 0.1)
    if true {
        fmt.Println(v2, v27, v3)
        if (!v2) {
            fmt.Println(v3, v26)
            var v28 int16 = 1
            var v29 float32 = 3.14159
            _ = v28
            _ = v29
        } else {
            var v30 int8 = 0
            _ = v30
        }
    } else {
        var v31 uint32 = v23
        _ = v31
    }
    v26 = (127 >= -1)
    if (!((v27 + (v27 / 1.0)) < 3.14159)) {
        v27 -= v27
    } else {
        var v32 bool = (!v2)
        if (!v2) {
            switch -82 {
            case 722:
                switch -176 {
                case -2147483648:
                    if v26 {
                        v33 := GenS3{F0: (v23 / ((4294967295) | 1)), F1: 0}
                        v34 := GenS4{F0: 232, F1: (!(1 != 127)), F2: v3}
                        v27 = v27
                        _ = v33
                        _ = v34
                    }
                case 2147483647:
                    switch -149 {
                    case 32768:
                        for fi35 := 0; fi35 < 12; fi35++ {
                            var v36 uint8 = uint8(v23)
                            _ = fi35
                            _ = v36
                        }
                    default:
                        var v37 int16 = -32768
                        _ = v37
                    }
                case -301:
                    v38 := GenS4{F0: 2, F1: v2, F2: v3}
                    _ = v38
                }
            case 4294967295:
                v39 := GenS5{F0: (!(((!false) && true) && false)), F1: 9223372036854775807, F2: (v1 % ((127) | 1))}
                _ = v39
            case 2147483647:
                v40 := GenS4{F0: 1, F1: false, F2: (v3 & 0)}
                _ = v40
            default:
                var v41 uint8 = 53
                _ = v41
            }
            var v42 uint16 = 65535
            _ = v42
        }
        _ = v32
    }
    _ = len(v4)
    switch 0 {
    case -9223372036854775808:
        switch 351 {
        case -629:
            var v43 string = "hello"
            _ = v43
        case -904:
            switch "a" {
            case "a":
                switch 32768 {
                case 0:
                    v44 := GenS1{F0: 2}
                    _ = v44
                }
            case "\n\t":
                switch 32768 {
                case 827:
                    var v45 uint64 = 4294967295
                    _ = v45
                case -322:
                    v26 = v2
                }
            default:
                fmt.Println(v2, v3)
            }
        case 9223372036854775807:
            fmt.Println(v23, v2, v23)
        default:
            fmt.Println(v1, v27)
        }
    case -2147483648:
        _ = len(v4)
    }
    var v46 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v46
}
