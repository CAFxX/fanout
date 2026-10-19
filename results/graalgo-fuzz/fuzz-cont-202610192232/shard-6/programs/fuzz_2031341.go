package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int; F1 int8 }
func genFn0(p0 uint16, p1 bool) int16 {
    p1 = ((("\n\t" + "") + "") < "hello")
    p0 = p0
    v5 := []int{128, 176, -27, 127}
    _ = p0
    _ = p1
    _ = v5
    return -32768
}
func genFn1() int64 {
    var v6 string = ("hello" + "\n\t")
    v7 := []int{32768, -2147483648, -234}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint16 = 0
        _ = fi8
        _ = v9
    }
    v6 = v6
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int16 = 32767
    var v2 uint = 255
    v3 := []string{"\n\t", ("" + "hello"), (("a" + "hello") + "foo bar"), "hello"}
    var v4 string = "foo bar"
    v4 = v4
    switch "\n\t" {
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        switch -1 {
        case -9223372036854775808:
            v4 = v4
        case 128:
            var v10 bool = false
            _ = v10
        default:
            fmt.Println(v4)
        }
    default:
        fmt.Println(v4, v4, v4)
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint64 = 42
        _ = fi11
        _ = v12
    }
    fmt.Println(v2, v4, v2)
    if false {
        if (!true) {
            switch v4 {
            case "a":
                v2 -= v2
            case "":
                v13 := GenS1{F0: "", F1: 65535}
                _ = v13
            }
            if (2 != 2) {
                for fi14 := 0; fi14 < 14; fi14++ {
                    v4 = v4
                    for fi15 := 0; fi15 < 13; fi15++ {
                        switch (fi15 | 2147483647) {
                        case 32768:
                            v1 = (v1 ^ 42)
                        default:
                            v2 += v2
                        }
                        v16 := GenS1{F0: v4, F1: 0}
                        _ = fi15
                        _ = v16
                    }
                    _ = fi14
                }
                v1 += -32768
            }
        }
        v1 = v1
    } else {
        if false {
            _ = len(v3)
        } else {
            for fi17 := 0; fi17 < 19; fi17++ {
                if (-137 < v1) {
                    var v18 int64 = 9223372036854775807
                    var v19 bool = false
                    _ = len(v3)
                    _ = v18
                    _ = v19
                } else {
                    var v20 int16 = 42
                    _ = v20
                }
                switch v4 {
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "foo bar":
                    if (2 < (v2 >> ((v2) & 63))) {
                        var v21 uint64 = 389
                        _ = v21
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = (v4 + (v4 + v4))
                    }
                }
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + (v4 + ("foo bar" + (v4 + (v4 + (v4 + "hello"))))))
                }
                _ = fi17
            }
            for fi22 := 0; fi22 < 6; fi22++ {
                for fi23 := 0; fi23 < 20; fi23++ {
                    var v24 float32 = 3.14159
                    var v25 uint = (v2 / ((v2) | 1))
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                v3 = append(v3, v4)
                _ = fi22
            }
        }
        v4 = "foo bar"
        v1 = v1
    }
    switch v4 {
    case "hello":
        if (0 >= -159) {
            if true {
                if (((!false) || false) || true) {
                    switch v4 {
                    case "":
                        if false {
                            v26 := GenS2{F0: 2147483647, F1: int16(v2), F2: -1}
                            var v27 int32 = -2147483648
                            v27 = 1
                            _ = v26
                            _ = v27
                        }
                    default:
                        var v28 bool = (3.14159 <= (3.14159 - -2.25))
                        _ = v28
                    }
                } else {
                    var v29 uint = (v2 - v2)
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 19; fi30++ {
                v31 := GenS3{F0: 255, F1: 790}
                _ = fi30
                _ = v31
            }
        } else {
            v4 = (("" + "\n\t") + ("foo bar" + "\n\t"))
            v1 = 0
            for fi32 := 0; fi32 < 17; fi32++ {
                var v33 float64 = 0.0
                fmt.Println(v4, v1)
                _ = fi32
                _ = v33
            }
        }
    case "a":
        var v34 int64 = -176
        _ = v34
    case "":
        switch 128 {
        case 128:
            v3 = append(v3, v4)
        default:
            switch -396 {
            case 116:
                var v35 float64 = ((3.14159 * -2.25) + (1e10 - 3.14159))
                _ = v35
            case -843:
                v2 -= v2
            case 127:
                if true {
                    v36 := GenS4{F0: 666, F1: -128}
                    v3 = append(v3, "")
                    _ = v36
                }
            default:
                _ = len(v3)
            }
        }
    default:
        var v37 uint = 2
        _ = v37
    }
    var v38 int = 1
    v38 -= (v38 / ((v38) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
