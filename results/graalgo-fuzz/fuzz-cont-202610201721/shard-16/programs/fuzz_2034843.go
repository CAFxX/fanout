package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int32 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8; F1 bool; F2 uint }
func genFn0() uint16 {
    if (!true) {
        v5 := GenS1{F0: 118, F1: "\n\t", F2: 2147483647}
        if (255 <= 1) {
            v6 := GenS1{F0: -1, F1: ("" + "foo bar"), F2: 42}
            switch 32767 {
            case 128:
                var v7 int8 = 1
                _ = v7
            case 590:
                v8 := GenS2{F0: 42, F1: 42, F2: 32767}
                _ = v8
            case -298:
                v9 := GenS2{F0: 42, F1: 1000, F2: 671}
                _ = v9
            }
            _ = v6
        }
        _ = v5
    }
    switch 4294967295 {
    case 9223372036854775807:
        var v10 uint = 2
        _ = v10
    }
    return 664
}
func genFn1(p0 string) uint8 {
    v11 := []int{4294967295, -652, 4294967295}
    var v12 bool = (-2147483648 != 2147483647)
    var v13 uint8 = 2
    switch 4294967295 {
    case -145:
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    case 0:
        v12 = (1.5 >= 3.14159)
    case 32768:
        v11 = append(v11, 0)
    default:
        switch -721 {
        case -81:
            if (913 != 4294967295) {
                for fi14 := 0; fi14 < 1; fi14++ {
                    switch (fi14 & fi14) {
                    case -2147483648:
                        _ = len(v11)
                    case 29:
                        for fi15 := 0; fi15 < 4; fi15++ {
                            _ = len(v11)
                            v16 := GenS1{F0: 1, F1: "\n\t", F2: 927}
                            var v17 string = ("a" + "foo bar")
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    default:
                        _ = len(v11)
                    }
                    _ = fi14
                }
                if (!(-128 >= 42)) {
                    v13 = v13
                } else {
                    var v18 float32 = -2.25
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 16; fi19++ {
                    var v20 float64 = 1.5
                    fmt.Println(p0, v13, p0)
                    p0 = "hello"
                    _ = fi19
                    _ = v20
                }
                p0 = "hello"
            }
        case 459:
            switch 777 {
            case 9223372036854775807:
                if (255 <= 0) {
                    v12 = (!v12)
                }
            default:
                v21 := GenS2{F0: 18446744073709551615, F1: 32767, F2: 0}
                _ = v21
            }
        case 32768:
            var v22 uint64 = 2
            _ = v22
        }
    }
    if (42 > 981) {
        v12 = (995 >= uint64(v13))
        var v23 uint8 = v13
        p0 = "hello"
        _ = v23
    } else {
        if true {
            var v24 float32 = (1.5 - (-2.25 / 1.0))
            v24 += float32(0.0)
            _ = v24
        } else {
            p0 = p0
            p0 = ("" + p0)
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        var v25 int32 = 42
        _ = v25
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return (v13 & v13)
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]int)
    var v4 float64 = 0.0
    switch "" {
    case "a":
        v26 := GenS2{F0: (^(v2)), F1: 0, F2: 0}
        _ = v26
    default:
        switch -689 {
        case -2147483648:
            delete(v3, 473)
        case 128:
            fmt.Println(v2, v2)
        case -486:
            v27 := GenS3{F0: 42, F1: false, F2: 766}
            _ = v27
        }
    }
    v28 := v1["foo bar"]
    var v29 int16 = -897
    var v30 uint64 = v2
    fmt.Println(v30, v4)
    v31 := GenS3{F0: 0, F1: false, F2: 949}
    v32 := GenS1{F0: 0, F1: (("a" + "") + "a"), F2: -2147483648}
    switch (v28 | v28) {
    case 42:
        v33 := GenS3{F0: -73, F1: (!((v2 | v2) >= v2)), F2: 18446744073709551615}
        _ = v33
    case 127:
        for fi34 := 0; fi34 < 17; fi34++ {
            fmt.Println(v2)
            switch "" {
            case "a":
                v4 = (float64((((1.5 + 0.1) / ((1e10 - float32(v4)) - float32(0.0))) * -2.25)) / (v4 + v4))
            }
            _ = fi34
        }
    case 0:
        var v35 uint = 255
        _ = v35
    default:
        for fi36 := 0; fi36 < 2; fi36++ {
            v4 -= -2.25
            _ = fi36
        }
    }
    var v37 uint16 = 32767
    v2 += v30
    v2 += 42
    fmt.Println(v29)
    var v38 uint16 = 1
    if false {
        v4 = v4
        fmt.Println(v38, v30, v37)
        switch v28 {
        case 4294967295:
            var v39 uint = 4294967295
            _ = v39
        default:
            if true {
                var v40 float32 = -2.25
                var v41 int = (v28 - (-(v28)))
                fmt.Println(v4)
                _ = v40
                _ = v41
            } else {
                if ((v30 >> ((4294967295) & 63)) != (v30 & 1)) {
                    v28 += (v28 & (v28 << ((18446744073709551615) & 63)))
                    var v42 uint16 = (v38 << ((659) & 63))
                    var v43 uint64 = (v30 | 1)
                    _ = v42
                    _ = v43
                }
            }
        }
    }
    switch v28 {
    case -9223372036854775808:
        for fi44 := 0; fi44 < 15; fi44++ {
            switch ("" + ("" + "")) {
            case "":
                _ = len(v3)
            case "foo bar":
                var v45 int = fi44
                _ = v45
            default:
                v3[fi44] = 9223372036854775807
            }
            _ = fi44
        }
    case 9223372036854775807:
        fmt.Println(v28)
    case 753:
        v46 := v1["hello"]
        _ = v46
    default:
        var v47 bool = (true || false)
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v4
}
