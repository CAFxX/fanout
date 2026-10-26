package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 string; F1 uint16; F2 uint32 }
func genFn0(p0 int16, p1 uint, p2 uint16) int32 {
    v4 := GenS1{F0: 417, F1: -1}
    var v5 int16 = (p0 / (((p0 % (((-(p0))) | 1))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -2147483648
}

func main() {
    var v1 float64 = 0.1
    v2 := []string{("a" + "a"), "a"}
    var v3 float32 = (1e10 / 0.1)
    if (926 >= 198) {
        v2 = append(v2, "foo bar")
    }
    if false {
        v6 := GenS1{F0: -2147483648, F1: -128}
        _ = v6
    }
    v3 = float32(((v1 + float64(v3)) * float64(v3)))
    if true {
        v3 -= v3
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch 2147483647 {
        case 9223372036854775807:
            if len(v2) > 0 {
                v2[0] = ("foo bar" + "\n\t")
            }
        case 127:
            switch -515 {
            case -555:
                v7 := GenS2{F0: "\n\t", F1: ("a" + ("" + ""))}
                _ = v7
            case -52:
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v3 = v3
                }
            case 516:
                _ = len(v2)
            }
        case -9223372036854775808:
            v8 := GenS1{F0: -159, F1: 42}
            _ = v8
        default:
            var v9 int64 = 0
            _ = v9
        }
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    fmt.Println(v3)
    v10 := GenS3{F0: 128, F1: 42}
    v11 := GenS2{F0: (("foo bar" + ("a" + ("\n\t" + "foo bar"))) + "hello"), F1: "foo bar"}
    fmt.Println(v1, v3)
    switch 0 {
    case 128:
        var v12 int8 = 0
        _ = v12
    case -900:
        fmt.Println(v1, v1)
    case 127:
        if false {
            fmt.Println(v1, v3)
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 int = fi13
                v15 := GenS4{F0: ("a" + (("hello" + "") + "")), F1: 0, F2: 187}
                fmt.Println(v1, v3, v14)
                _ = fi13
                _ = v14
                _ = v15
            }
            if len(v2) > 0 {
                v2[0] = "a"
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
