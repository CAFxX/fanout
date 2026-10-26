package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 int; F2 int8 }

func main() {
    var v1 bool = (true && true)
    v2 := []string{((("foo bar" + ("\n\t" + ("a" + ""))) + "foo bar") + ("a" + ((("foo bar" + ("\n\t" + "a")) + ("foo bar" + "foo bar")) + ("foo bar" + "\n\t")))), "foo bar", (("" + "") + "a"), "foo bar"}
    v3 := []int{}
    var v4 uint8 = 128
    fmt.Println(v4)
    v1 = (1 != 0)
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            v1 = ((!v1) || (42 >= 33))
            var v7 int64 = 42
            v3 = append(v3, (^(fi6)))
            var v8 uint32 = 65535
            fmt.Println(v4)
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    if (2 == v4) {
        for fi9 := 0; fi9 < 14; fi9++ {
            switch "a" {
            case "foo bar":
                switch "hello" {
                case "\n\t":
                    if (1000 < -1) {
                        var v10 int16 = 42
                        for fi11 := 0; fi11 < 20; fi11++ {
                            var v12 int64 = 0
                            _ = fi11
                            _ = v12
                        }
                        _ = v10
                    }
                default:
                    fmt.Println(v1, fi9)
                }
            case "hello":
                v13 := GenS1{F0: 66, F1: fi9}
                _ = v13
            case "\n\t":
                v14 := GenS1{F0: 942, F1: (fi9 << ((4294967295) & 63))}
                _ = v14
            }
            v4 = (v4 << ((889) & 63))
            var v15 uint16 = uint16(v4)
            _ = fi9
            _ = v15
        }
        if ((v4 + v4) != (v4 ^ v4)) {
            for fi16 := 0; fi16 < 8; fi16++ {
                v4 = (^(v4))
                v17 := GenS1{F0: -2147483648, F1: 32768}
                _ = fi16
                _ = v17
            }
            v18 := GenS2{F0: 42, F1: (v4 << ((4294967295) & 63))}
            _ = v18
        }
        fmt.Println(v4)
    } else {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            switch -529 {
            case -1:
                var v20 int64 = -9223372036854775808
                _ = v20
            case 500:
                if (1e10 > (1.5 + (0.1 / 3.14159))) {
                    var v21 float64 = (-2.25 / 0.1)
                    fmt.Println(v1)
                    var v22 uint16 = 792
                    _ = v21
                    _ = v22
                } else {
                    switch ("a" + (("" + ("a" + "\n\t")) + "a")) {
                    case "":
                        var v23 int32 = -2147483648
                        _ = v23
                    default:
                        v1 = (-1 < -8)
                    }
                    v24 := GenS3{F0: 1}
                    _ = v24
                }
            case 959:
                var v25 int64 = -1
                _ = v25
            default:
                v4 = 140
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v26 := GenS1{F0: 1, F1: 4294967295}
            _ = fi19
            _ = v26
        }
    }
    switch "" {
    case "a":
        if (1 <= 1) {
            v1 = (-801 > 1)
            fmt.Println(v4, v4)
        }
    case "":
        var v27 uint = 1
        _ = v27
    case "foo bar":
        if (0 > 65535) {
            var v28 string = ""
            _ = v28
        } else {
            v29 := GenS3{F0: 2}
            v4 += (v4 + 127)
            _ = v29
        }
    default:
        fmt.Println(v1, v1)
    }
    switch 9223372036854775807 {
    case 42:
        for fi30 := 0; fi30 < 4; fi30++ {
            fmt.Println(v4)
            v31 := GenS1{F0: 2147483647, F1: fi30}
            switch "" {
            case "hello":
                v4 += 127
            case "\n\t":
                _ = len(v3)
            default:
                if v1 {
                    _ = len(v3)
                    v4 += (v4 | (v4 ^ 205))
                } else {
                    var v32 int8 = 127
                    if v1 {
                        v33 := GenS2{F0: 0, F1: 141}
                        if v1 {
                            v1 = v1
                            fmt.Println(v4)
                            v1 = ((v1 || false) || true)
                        }
                        _ = v33
                    } else {
                        v4 = (v4 - 38)
                        v1 = (!true)
                    }
                    _ = v32
                }
            }
            _ = fi30
            _ = v31
        }
    default:
        if v1 {
            if (1e10 <= float32(1e10)) {
                v34 := GenS1{F0: -2147483648, F1: 0}
                for fi35 := 0; fi35 < 14; fi35++ {
                    if v1 {
                        if len(v3) > 0 {
                            v3[0] = (-(fi35))
                        }
                    }
                    v36 := GenS4{F0: fi35}
                    _ = fi35
                    _ = v36
                }
                _ = v34
            } else {
                switch 128 {
                case 726:
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v1, v1)
                    }
                case 32767:
                    switch -1 {
                    case 9223372036854775807:
                        v4 += v4
                    case -9223372036854775808:
                        for fi37 := 0; fi37 < 6; fi37++ {
                            if len(v2) > 0 {
                                v2[0] = ("a" + (("" + "hello") + ("\n\t" + "")))
                            }
                            _ = fi37
                        }
                    }
                default:
                    v3 = append(v3, 9223372036854775807)
                }
                v1 = ("hello" >= "foo bar")
            }
            v38 := GenS1{F0: 0, F1: -9223372036854775808}
            v4 = v4
            _ = v38
        } else {
            v4 -= (v4 % ((0) | 1))
        }
    }
    v1 = (1 != 65535)
    switch ("foo bar" + "foo bar") {
    case "a":
        v39 := GenS5{F0: 88, F1: -9223372036854775808, F2: 127}
        _ = v39
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v40 := GenS2{F0: -2147483648, F1: 176}
    v1 = (0 >= 364)
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        v1 = (v4 <= v4)
        _ = fi41
    }
    var v42 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
}
