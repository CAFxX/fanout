package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 bool }
type GenS3 struct { F0 int16; F1 bool }

func main() {
    var v1 float32 = (0.0 + 3.14159)
    var v2 int16 = -1
    var v3 float32 = v1
    var v4 uint16 = 175
    v5 := []string{("hello" + "\n\t"), ""}
    var v6 int32 = -1
    v5 = append(v5, ("\n\t" + "hello"))
    switch 1 {
    case 42:
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    }
    switch -603 {
    case 185:
        var v7 int64 = -9223372036854775808
        _ = v7
    case -725:
        if len(v5) > 0 {
            v5[0] = ("foo bar" + "\n\t")
        }
    default:
        for fi8 := 0; fi8 < 13; fi8++ {
            if ((208 < 128) && (uint(v4) <= 42)) {
                for fi9 := 0; fi9 < 5; fi9++ {
                    switch fi9 {
                    case 32768:
                        var v10 float32 = 1.5
                        _ = v10
                    }
                    _ = fi9
                }
            }
            var v11 uint32 = 2147483647
            _ = fi8
            _ = v11
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = v4
    v12 := GenS1{F0: -1, F1: (v6 + (-(v6))), F2: (0.0 == 0.1)}
    var v13 float64 = 0.0
    var v14 int8 = 42
    if len(v5) > 0 {
        v5[0] = ((("a" + ((("hello" + "foo bar") + ("hello" + ("\n\t" + (("\n\t" + ("foo bar" + (((("a" + "hello") + "hello") + "foo bar") + ("" + "\n\t")))) + (("a" + "hello") + "hello"))))) + "")) + "") + "a")
    }
    if false {
        switch 32767 {
        case 32767:
            v2 = v2
        case 2147483647:
            var v15 string = ((("\n\t" + "foo bar") + ("\n\t" + ((("\n\t" + "\n\t") + "a") + ("\n\t" + "foo bar")))) + "\n\t")
            _ = v15
        }
        if false {
            fmt.Println(v14)
            for fi16 := 0; fi16 < 3; fi16++ {
                _ = len(v5)
                for fi17 := 0; fi17 < 13; fi17++ {
                    if true {
                        v18 := GenS1{F0: 9223372036854775807, F1: v6, F2: false}
                        _ = v18
                    }
                    _ = fi17
                }
                _ = fi16
            }
        }
    } else {
        for fi19 := 0; fi19 < 7; fi19++ {
            if (!(491 != 1)) {
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
            } else {
                var v20 string = "\n\t"
                _ = v20
            }
            if true {
                var v21 uint = 146
                var v22 int16 = (v2 % ((v2) | 1))
                v23 := GenS2{F0: 42, F1: 837, F2: true}
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi19
        }
        var v24 string = "foo bar"
        for fi25 := 0; fi25 < 18; fi25++ {
            switch ("a" + "") {
            case "hello":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                if true {
                    v26 := GenS3{F0: v2, F1: (32767 > 1000)}
                    _ = v26
                } else {
                    v5 = append(v5, v24)
                    if len(v5) > 0 {
                        v5[0] = v24
                    }
                    if ((v14 / ((61) | 1)) <= (v14 * (v14 % (((v14 * (v14 & (v14 & (v14 + 100))))) | 1)))) {
                        var v27 bool = false
                        var v28 int8 = (-(v14))
                        _ = v27
                        _ = v28
                    } else {
                        _ = len(v5)
                        fmt.Println(fi25, v2, v4)
                        var v29 uint16 = 65535
                        _ = v29
                    }
                }
            }
            v2 = (v2 / (((v2 >> ((4294967295) & 63))) | 1))
            _ = fi25
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
