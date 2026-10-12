package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int64 }
type GenS2 struct { F0 uint8; F1 int32 }
func genFn0(p0 int32) int8 {
    var v5 int8 = -117
    switch 42 {
    case -188:
        p0 -= p0
    default:
        v5 = -72
    }
    v6 := GenS1{F0: "\n\t", F1: (v5 * 1), F2: 0}
    var v7 uint = 65535
    v8 := []int{-2147483648, 824, 4294967295, 179}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v5
}
func genFn1(p0 int8) int64 {
    v9 := GenS1{F0: "foo bar", F1: (p0 - (p0 >> ((255) & 63))), F2: 64}
    v10 := []int{127}
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 string = ("hello" + (("hello" + "foo bar") + "a"))
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: (((("" + "a") + ("\n\t" + "a")) + "\n\t") + ("a" + "hello")), F1: (p0 / ((p0) | 1)), F2: 647}
    _ = p0
    _ = v10
    _ = v13
    _ = v9
    return -252
}

func main() {
    v1 := []string{}
    v2 := []int{}
    var v3 uint = 647
    var v4 int8 = 100
    fmt.Println(v4)
    fmt.Println(v4, v4)
    var v14 bool = ("foo bar" == "a")
    v4 -= 100
    var v15 uint = (v3 & v3)
    var v16 uint8 = 2
    for fi17 := 0; fi17 < 15; fi17++ {
        if v14 {
            var v18 float32 = 1.5
            if len(v2) > 0 {
                v2[0] = (fi17 | 0)
            }
            _ = v18
        }
        _ = fi17
    }
    switch "\n\t" {
    case "":
        if (0 > 746) {
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 int16 = -32768
                _ = fi19
                _ = v20
            }
            if true {
                switch 128 {
                case 4294967295:
                    v2 = append(v2, 128)
                }
                for fi21 := 0; fi21 < 18; fi21++ {
                    v3 -= (v3 | (v3 | (v3 << ((0) & 63))))
                    _ = fi21
                }
                if false {
                    _ = len(v2)
                    v15 = (v3 & v15)
                } else {
                    v3 -= 0
                    var v22 uint8 = v16
                    _ = v22
                }
            }
        }
    case "foo bar":
        v23 := GenS2{F0: (v16 >> (((v3 / (((v3 * (v3 % ((38) | 1)))) | 1))) & 63)), F1: 0}
        _ = v23
    }
    fmt.Println(v14, v16)
    var v24 int8 = (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
