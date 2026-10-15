package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 bool; F1 uint }
func genFn0(p0 uint8, p1 int64) string {
    v5 := GenS1{F0: -32768, F1: 1000}
    var v6 uint64 = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (("hello" + "") + ("a" + "a"))
}

func main() {
    var v1 bool = (840 != -9223372036854775808)
    v2 := make(map[uint32]string)
    var v3 int16 = 806
    var v4 uint = 0
    var v7 uint = (v4 % (((^(v4))) | 1))
    fmt.Println(v4, v3, v4)
    v7 = 255
    v8 := []int{127, 792, 0, 127}
    v9 := GenS2{F0: 255, F1: -1}
    v10 := GenS1{F0: (^(v3)), F1: (v3 & (v3 % ((v3) | 1)))}
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := v2[386]
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 string = ("a" + (((("hello" + ("hello" + ("foo bar" + ("hello" + "\n\t")))) + "a") + (("a" + ("\n\t" + ("\n\t" + ("foo bar" + ((("hello" + ("\n\t" + ("hello" + "\n\t"))) + ("a" + "\n\t")) + "hello"))))) + (("hello" + ("foo bar" + "\n\t")) + "hello"))) + "foo bar"))
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := GenS2{F0: 65535, F1: -2147483648}
            var v17 int32 = 817
            _ = fi15
            _ = v16
            _ = v17
        }
        if v1 {
            v2[65535] = v14
            for fi18 := 0; fi18 < 19; fi18++ {
                switch v14 {
                case "hello":
                    var v19 int8 = 0
                    _ = v19
                case "\n\t":
                    v7 = (v4 >> ((v7) & 63))
                default:
                    switch "foo bar" {
                    case "foo bar":
                        var v20 uint8 = 69
                        _ = v20
                    }
                }
                v21 := GenS3{F0: 18446744073709551615, F1: -63}
                _ = fi18
                _ = v21
            }
            switch v14 {
            case "":
                for fi22 := 0; fi22 < 14; fi22++ {
                    v23 := GenS3{F0: 42, F1: -106}
                    switch fi13 {
                    case -2147483648:
                        v24 := GenS2{F0: 1, F1: 915}
                        _ = v24
                    case -9223372036854775808:
                        v25 := v2[839]
                        _ = v25
                    case 32768:
                        v14 = ""
                    }
                    v26 := GenS4{F0: (!false), F1: v4}
                    _ = fi22
                    _ = v23
                    _ = v26
                }
            case "foo bar":
                var v27 int16 = v3
                _ = v27
            case "hello":
                fmt.Println(fi13, v7)
            default:
                var v28 bool = v1
                _ = v28
            }
        }
        _ = fi13
        _ = v14
    }
    var v29 uint = 2
    fmt.Println(v7, v4, v29)
    v30 := GenS5{F0: (1 < 20), F1: v7}
    var v31 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
