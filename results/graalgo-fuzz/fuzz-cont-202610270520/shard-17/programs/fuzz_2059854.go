package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 int32 = -1
    var v2 int32 = -1
    var v3 uint8 = 0
    switch "hello" {
    case "\n\t":
        var v4 bool = (-570 > -91)
        _ = v4
    default:
        v5 := GenS1{F0: 2, F1: 0, F2: (^(v2))}
        _ = v5
    }
    if false {
        v2 += v2
    } else {
        if true {
            var v6 bool = (228 >= 2147483648)
            v7 := GenS2{F0: -9223372036854775808}
            _ = v6
            _ = v7
        }
        v1 -= (v2 - 1)
    }
    v8 := GenS3{F0: "foo bar", F1: 42}
    v3 = (v3 | v3)
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS4{F0: -97, F1: 932}
        v3 += 255
        switch "hello" {
        case "hello":
            var v11 uint8 = 2
            _ = v11
        case "a":
            var v12 bool = true
            _ = v12
        case "foo bar":
            switch ("hello" + "foo bar") {
            case "hello":
                var v13 int8 = 127
                _ = v13
            case "a":
                fmt.Println(v1)
            case "":
                for fi14 := 0; fi14 < 17; fi14++ {
                    v2 = -2147483648
                    _ = fi14
                }
            }
        default:
            v2 = 0
        }
        _ = fi9
        _ = v10
    }
    v3 = v3
    v15 := []int{127, -162}
    var v16 bool = (1 == 0)
    for fi17 := 0; fi17 < 18; fi17++ {
        if false {
            for fi18 := 0; fi18 < 3; fi18++ {
                switch "" {
                case "":
                    v2 = (v2 % ((2147483647) | 1))
                }
                v2 = v2
                v19 := GenS5{F0: 912}
                _ = fi18
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                fmt.Println(fi20)
                _ = fi20
            }
            switch fi17 {
            case 0:
                v16 = v16
            case -81:
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 int = 69
                    var v23 uint64 = 18446744073709551615
                    var v24 uint32 = 0
                    var v25 int8 = -128
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            default:
                for fi26 := 0; fi26 < 18; fi26++ {
                    if (("foo bar" + "hello") <= ("\n\t" + (("foo bar" + "\n\t") + ("hello" + (("\n\t" + "") + "a"))))) {
                        v27 := GenS4{F0: int8(v1), F1: -1}
                        _ = v27
                    } else {
                        v16 = false
                    }
                    v1 = v1
                    _ = fi26
                }
            }
        }
        v3 += (v3 - 42)
        _ = fi17
    }
    fmt.Println(v16, v16)
    switch ("hello" + "\n\t") {
    case "a":
        v3 = v3
    case "foo bar":
        v1 -= 921
    case "\n\t":
        if (((v16 || v16) || ((553 != 1) && v16)) && v16) {
            var v28 bool = (42 > 304)
            _ = v28
        } else {
            var v29 string = (((("" + (("hello" + ("foo bar" + "\n\t")) + (("foo bar" + ("" + "a")) + "a"))) + "") + "") + ("a" + (("foo bar" + "\n\t") + ("foo bar" + "a"))))
            _ = v29
        }
    }
    if false {
        var v30 int = -667
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 int8 = 127
            v33 := GenS4{F0: (^(v32)), F1: 9223372036854775807}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = v30
    }
    fmt.Println(v1, v2)
    for fi34 := 0; fi34 < 4; fi34++ {
        if (v16 && (((v16 && (100 < 0)) && false) || (!(-9223372036854775808 != 843)))) {
            v3 = v3
            if v16 {
                var v35 uint = 65535
                v15 = append(v15, (-(fi34)))
                v3 = 42
                _ = v35
            }
        } else {
            v36 := GenS3{F0: "hello", F1: 4294967295}
            _ = v36
        }
        _ = fi34
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        for fi38 := 0; fi38 < 6; fi38++ {
            fmt.Println(v16, v16, v1)
            _ = fi38
        }
        for fi39 := 0; fi39 < 4; fi39++ {
            var v40 string = ("\n\t" + "")
            _ = fi39
            _ = v40
        }
        fmt.Println(v1)
        _ = fi37
    }
    var v41 float32 = (0.0 - 0.1)
    if ((true || (100 == -128)) && false) {
        var v42 int8 = -99
        _ = v42
    } else {
        v2 = v2
        var v43 int32 = (v1 << ((448) & 63))
        _ = v43
    }
    fmt.Println(v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v41
    _ = v8
}
