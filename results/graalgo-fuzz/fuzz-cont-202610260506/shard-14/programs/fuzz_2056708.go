package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 string }

func main() {
    var v1 bool = ((127 == -106) && ((float64(0.0) < 1e10) || false))
    var v2 uint8 = 42
    var v3 int32 = 450
    var v4 uint32 = 42
    v5 := GenS1{F0: 255}
    switch ("\n\t" + ("a" + (("foo bar" + "a") + "\n\t"))) {
    case "foo bar":
        v4 += 4294967295
    default:
        v3 = v3
    }
    v6 := []int{128, 776, -335}
    var v7 uint16 = 255
    v7 -= (v7 | v7)
    v6 = append(v6, 876)
    var v8 uint32 = 4294967295
    fmt.Println(v2, v7, v1)
    for fi9 := 0; fi9 < 10; fi9++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if v1 {
            _ = len(v6)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            switch "foo bar" {
            case "foo bar":
                var v10 float64 = ((((float64(1.5) * 0.0) / (1.5 - 3.14159)) + 1.5) * 0.1)
                _ = v10
            case "\n\t":
                var v11 string = ("" + "")
                _ = v11
            }
        }
        _ = fi9
    }
    switch 0 {
    case -2147483648:
        v4 = (v4 - 2147483648)
    case 128:
        v7 = v7
    default:
        if v1 {
            var v12 int64 = -146
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS1{F0: 461}
                for fi15 := 0; fi15 < 8; fi15++ {
                    if v1 {
                        v16 := GenS2{F0: -128, F1: "foo bar"}
                        var v17 string = ("a" + "hello")
                        fmt.Println(v17, fi13)
                        _ = len(v6)
                        _ = v16
                        _ = v17
                    }
                    v3 = (v3 / ((-1) | 1))
                    _ = fi15
                }
                _ = fi13
                _ = v14
            }
            var v18 uint32 = (v8 % (((v4 / (((v4 | 65535)) | 1))) | 1))
            _ = v12
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
