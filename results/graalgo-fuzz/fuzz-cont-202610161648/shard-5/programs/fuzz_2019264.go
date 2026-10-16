package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint }

func main() {
    var v1 int16 = 0
    v2 := []uint32{65535, 960}
    v3 := []int{0}
    v1 -= (-(v1))
    v1 = v1
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 string = (((("hello" + "hello") + "") + ((("hello" + ("foo bar" + "a")) + ("hello" + ("\n\t" + "foo bar"))) + "a")) + ("\n\t" + ((("a" + "") + "hello") + ("hello" + "\n\t"))))
        v5 = (v5 + (v5 + v5))
        var v6 bool = ((1 >= 1) || true)
        _ = fi4
        _ = v5
        _ = v6
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        var v7 uint32 = 42
        _ = v7
    case "":
        v1 += (-(v1))
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        switch 42 {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 12; fi8++ {
                for fi9 := 0; fi9 < 2; fi9++ {
                    var v10 int = (fi9 ^ 32768)
                    _ = fi9
                    _ = v10
                }
                var v11 uint16 = 0
                v1 = v1
                _ = fi8
                _ = v11
            }
        }
    }
    var v12 string = "\n\t"
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(fi13)
        v2 = append(v2, 2)
        v12 = "a"
        _ = fi13
    }
    v14 := GenS1{F0: 32767, F1: 0, F2: 0}
    fmt.Println(v12, v1)
    var v15 int = 459
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch v12 {
    case "foo bar":
        var v16 uint = 255
        _ = v16
    default:
        for fi17 := 0; fi17 < 14; fi17++ {
            v1 = v1
            if (!((^(fi17)) < (fi17 + fi17))) {
                switch 4294967295 {
                case 32768:
                    fmt.Println(v12, v12, fi17)
                default:
                    var v18 bool = false
                    _ = v18
                }
                v19 := GenS1{F0: 32767, F1: -128, F2: 255}
                v1 += v1
                _ = v19
            } else {
                _ = len(v3)
                v1 = (v1 >> ((18446744073709551615) & 63))
                for fi20 := 0; fi20 < 2; fi20++ {
                    v21 := GenS1{F0: 42, F1: 42, F2: 255}
                    _ = len(v2)
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi17
        }
    }
    _ = len(v3)
    fmt.Println(v1, v1)
    v1 = (v1 << ((18446744073709551615) & 63))
    if len(v2) > 0 {
        v2[0] = 65535
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
