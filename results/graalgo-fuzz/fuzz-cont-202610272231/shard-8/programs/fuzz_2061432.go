package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int16; F1 uint16; F2 uint }
type GenS6 struct { F0 int; F1 int64; F2 int8 }
type GenS7 struct { F0 uint; F1 int32; F2 uint }
type GenS8 struct { F0 int; F1 string; F2 uint32 }
func genFn0() int64 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint32 = 496
        if true {
            var v8 int8 = -1
            v9 := []int{fi6}
            if (1 <= -1) {
                fmt.Println(v8, fi6, v8)
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            }
            _ = v8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 uint64 = 4294967295
    for fi11 := 0; fi11 < 4; fi11++ {
        switch fi11 {
        case 32768:
            switch 1 {
            case 128:
                var v12 uint16 = 255
                _ = v12
            case -919:
                switch (fi11 / (((fi11 / (((fi11 & fi11)) | 1))) | 1)) {
                case -9223372036854775808:
                    var v13 uint64 = v10
                    _ = v13
                }
            case -1:
                v10 -= v10
            }
        case -910:
            var v14 string = "foo bar"
            _ = v14
        }
        var v15 int = (fi11 << ((42) & 63))
        _ = fi11
        _ = v15
    }
    v16 := GenS1{F0: v10}
    _ = v10
    _ = v16
    return 0
}
func genFn1() string {
    switch "\n\t" {
    case "":
        if false {
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 uint8 = 128
                v19 := []int{}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    case "\n\t":
        var v20 string = "hello"
        _ = v20
    }
    var v21 int = -2147483648
    v22 := []int{v21, v21}
    _ = v21
    _ = v22
    return ""
}

func main() {
    var v1 int32 = 500
    v2 := []string{((("hello" + (("\n\t" + "hello") + "foo bar")) + (("hello" + ("" + "hello")) + ("foo bar" + "\n\t"))) + ("foo bar" + "a"))}
    var v3 uint64 = 720
    var v4 uint16 = 1
    var v5 uint64 = (v3 >> ((44) & 63))
    if len(v2) > 0 {
        v2[0] = ("" + (("foo bar" + "") + ("\n\t" + "a")))
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    var v23 float32 = 3.14159
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        if true {
            for fi25 := 0; fi25 < 13; fi25++ {
                if len(v2) > 0 {
                    v2[0] = "a"
                }
                var v26 float64 = -2.25
                if ((((v26 + (v26 + v26)) / v26) + v26) == 1.5) {
                    var v27 uint = 0
                    v28 := GenS2{F0: v4}
                    v29 := GenS3{F0: (^(v1)), F1: 3}
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    var v30 int = 573
                    v31 := GenS1{F0: (v5 + (v5 / ((1) | 1)))}
                    fmt.Println(v4, v26, v5)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v30
                    _ = v31
                }
                _ = fi25
                _ = v26
            }
        }
        v4 = 0
        var v32 float64 = -2.25
        _ = fi24
        _ = v32
    }
    v1 -= v1
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v33 bool = (2 >= 0)
        _ = v33
    case -9223372036854775808:
        fmt.Println(v4, v5, v1)
    case 4294967295:
        var v34 int8 = 1
        _ = v34
    default:
        fmt.Println(v3)
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        v1 = (v1 / ((42) | 1))
        for fi36 := 0; fi36 < 10; fi36++ {
            if true {
                switch (("foo bar" + "") + ((("\n\t" + "foo bar") + "") + ("a" + ("foo bar" + "hello")))) {
                case "foo bar":
                    v4 = 2
                case "\n\t":
                    if ((3.14159 - 0.1) != (-2.25 * -2.25)) {
                        v37 := GenS4{F0: int64(v4), F1: 2}
                        v3 -= v5
                        v23 = (((v23 + (v23 + (3.14159 / v23))) / 3.14159) * v23)
                        _ = v37
                    }
                case "":
                    v3 = uint64(v4)
                }
                v38 := GenS3{F0: (^(v1)), F1: 127}
                var v39 uint32 = 1
                _ = v38
                _ = v39
            }
            _ = fi36
        }
        _ = fi35
    }
    if (v1 < (v1 & v1)) {
        if (!((((v23 * (v23 + v23)) / v23) + 1.5) < (((v23 * (((v23 - v23) + v23) / v23)) / -2.25) + v23))) {
            v40 := GenS5{F0: -696, F1: 65535, F2: 0}
            v2 = append(v2, "a")
            var v41 uint = 2
            _ = v40
            _ = v41
        } else {
            var v42 int16 = -58
            for fi43 := 0; fi43 < 2; fi43++ {
                if true {
                    fmt.Println(v23, fi43)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi44 := 0; fi44 < 13; fi44++ {
                        var v45 int64 = 0
                        v4 -= (^(v4))
                        _ = fi44
                        _ = v45
                    }
                } else {
                    for fi46 := 0; fi46 < 14; fi46++ {
                        fmt.Println(v5, v42)
                        _ = fi46
                    }
                }
                v47 := GenS6{F0: (^(fi43)), F1: -396, F2: 127}
                _ = fi43
                _ = v47
            }
            _ = v42
        }
        switch "\n\t" {
        case "":
            var v48 uint8 = 2
            _ = v48
        case "\n\t":
            v5 = v3
        case "foo bar":
            switch "\n\t" {
            case "a":
                var v49 float32 = (v23 + (v23 + 3.14159))
                _ = v49
            case "foo bar":
                for fi50 := 0; fi50 < 11; fi50++ {
                    if false {
                        switch (-(fi50)) {
                        case 2147483647:
                            for fi51 := 0; fi51 < 18; fi51++ {
                                v4 = v4
                                v2 = append(v2, ("foo bar" + (("foo bar" + "a") + "\n\t")))
                                v3 = (v3 & uint64(v4))
                                _ = fi51
                            }
                        default:
                            _ = len(v2)
                        }
                        v23 -= 0.1
                    } else {
                        var v52 float32 = v23
                        _ = v52
                    }
                    _ = fi50
                }
            }
        default:
            if false {
                var v53 float64 = 0.0
                v2 = append(v2, "a")
                v23 -= 0.0
                _ = v53
            } else {
                if (128 <= -606) {
                    v4 = 255
                    v5 += v3
                    for fi54 := 0; fi54 < 16; fi54++ {
                        v55 := GenS7{F0: 65535, F1: (v1 % ((2147483647) | 1)), F2: 455}
                        v3 = v5
                        _ = fi54
                        _ = v55
                    }
                }
            }
        }
    } else {
        var v56 uint32 = 1
        _ = v56
    }
    var v57 uint8 = 42
    v2 = append(v2, "hello")
    switch 42 {
    case 9223372036854775807:
        var v58 uint = 0
        _ = v58
    case -731:
        v3 += (v3 - v5)
    case 128:
        var v59 float32 = (v23 - 0.0)
        _ = v59
    default:
        var v60 int8 = 18
        _ = v60
    }
    fmt.Println(v57)
    if ((v5 ^ v5) != (v3 % (((v5 - v3)) | 1))) {
        var v61 bool = false
        _ = v61
    }
    v2 = append(v2, (("foo bar" + "foo bar") + "foo bar"))
    v5 -= (^(v3))
    if true {
        v57 += (v57 | (v57 << ((427) & 63)))
        var v62 uint32 = 47
        v63 := GenS7{F0: uint(v57), F1: (v1 * (v1 / (((v1 >> ((65535) & 63))) | 1))), F2: 255}
        _ = v62
        _ = v63
    } else {
        v64 := GenS8{F0: 32768, F1: "hello", F2: 2}
        fmt.Println(v1, v1, v1)
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v57
}
