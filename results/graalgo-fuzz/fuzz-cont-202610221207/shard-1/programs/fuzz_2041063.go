package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int64; F1 int8 }

func main() {
    var v1 bool = false
    var v2 uint16 = 2
    v3 := make(map[int]int)
    v4 := GenS1{F0: 0, F1: ""}
    var v5 int16 = 130
    if false {
        for fi6 := 0; fi6 < 15; fi6++ {
            delete(v3, (-(fi6)))
            v2 = (v2 ^ 32767)
            if (-128 == 127) {
                switch "hello" {
                case "a":
                    var v7 uint16 = (^(v2))
                    _ = v7
                case "":
                    var v8 string = "hello"
                    _ = v8
                case "foo bar":
                    if v1 {
                        v1 = v1
                        v9 := []int{fi6, (fi6 ^ fi6), (fi6 >> ((255) & 63)), (fi6 * (fi6 / (((^(fi6))) | 1)))}
                        _ = v9
                    } else {
                        var v10 int = (fi6 ^ (fi6 >> ((2) & 63)))
                        var v11 uint16 = (v2 >> ((114) & 63))
                        _ = v10
                        _ = v11
                    }
                }
                fmt.Println(v5)
            } else {
                if (65535 >= 2) {
                    for fi12 := 0; fi12 < 12; fi12++ {
                        v13 := GenS2{F0: 81, F1: 91}
                        v5 -= -32768
                        v14 := []int{fi6}
                        fmt.Println(v2, v5, v2)
                        var v15 uint32 = 65535
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                }
            }
            _ = fi6
        }
        v16 := GenS1{F0: 18446744073709551615, F1: "hello"}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            if (-86 == -18) {
                v18 := []int{fi17}
                v5 -= v5
                var v19 uint32 = 2
                _ = v18
                _ = v19
            } else {
                fmt.Println(v1, v2, v5)
            }
            if (32767 <= (^(v5))) {
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 uint8 = 2
                    v5 = v5
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi17
        }
        v2 = 591
    }
    var v22 float32 = 1.5
    v23 := []int{128, 32768}
    var v24 string = "\n\t"
    var v25 bool = (0 > 1)
    v1 = false
    v22 = 0.0
    switch (("foo bar" + (v24 + v24)) + v24) {
    case "hello":
        var v26 uint32 = 2147483647
        _ = v26
    case "":
        if (v25 || true) {
            v5 = v5
            var v27 int32 = 1
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = v27
        }
    default:
        var v28 string = ((v24 + (("hello" + ("foo bar" + v24)) + v24)) + v24)
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
