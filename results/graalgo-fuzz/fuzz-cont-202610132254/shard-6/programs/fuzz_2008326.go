package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 int; F1 uint32; F2 string }
type GenS3 struct { F0 uint; F1 uint }
func genFn0() uint8 {
    v4 := GenS1{F0: 2, F1: -2147483648, F2: "hello"}
    var v5 uint8 = 2
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 string = ("a" + "")
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS2{F0: 2147483647, F1: 65535, F2: "a"}
            _ = fi8
            _ = v9
        }
        var v10 bool = ((true && (true && true)) && false)
        _ = fi6
        _ = v10
        _ = v7
    }
    v5 += 127
    _ = v4
    _ = v5
    return (v5 | (v5 | 255))
}

func main() {
    v1 := []int{31, 4294967295}
    var v2 int = 32767
    var v3 float32 = float32(float64((-2.25 - 1e10)))
    _ = len(v1)
    var v11 int64 = 9223372036854775807
    fmt.Println(v2, v2)
    fmt.Println(v3, v2)
    var v12 int64 = v11
    fmt.Println(v3, v3)
    v2 -= 855
    var v13 float64 = 3.14159
    switch (v2 + v2) {
    case 599:
        v2 += 2147483647
    default:
        for fi14 := 0; fi14 < 15; fi14++ {
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 int8 = 1
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi15
                _ = v16
            }
            var v17 uint = 877
            _ = fi14
            _ = v17
        }
    }
    v1 = append(v1, (v2 << ((42) & 63)))
    for fi18 := 0; fi18 < 7; fi18++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi18
    }
    switch ("\n\t" + "hello") {
    case "":
        fmt.Println(v2, v13, v3)
    default:
        if (2 <= 52) {
            for fi19 := 0; fi19 < 13; fi19++ {
                fmt.Println(v12, v12, v13)
                _ = fi19
            }
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 int8 = 42
                _ = fi20
                _ = v21
            }
        }
    }
    v22 := GenS3{F0: 18446744073709551615, F1: 1}
    v23 := GenS2{F0: v2, F1: 70, F2: ("" + "")}
    v13 -= v13
    var v24 string = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
