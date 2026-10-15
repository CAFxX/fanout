package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }

func main() {
    var v1 uint16 = 2
    var v2 uint8 = 1
    v3 := make(map[uint32]uint32)
    var v4 uint8 = (v2 * v2)
    fmt.Println(v4)
    if (false && true) {
        for fi5 := 0; fi5 < 20; fi5++ {
            v2 = v2
            v6 := GenS1{F0: fi5, F1: "hello"}
            v2 = 190
            _ = fi5
            _ = v6
        }
        var v7 uint8 = 193
        var v8 float32 = ((1.5 - (1.5 + 3.14159)) + 0.1)
        _ = v7
        _ = v8
    } else {
        v9 := v3[2]
        _ = v9
    }
    if true {
        v3[2147483648] = 797
    }
    if (false && false) {
        if false {
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 float64 = float64(1.5)
                var v12 int32 = 1
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := GenS1{F0: fi10, F1: ("foo bar" + "")}
                    var v15 int8 = -128
                    var v16 int8 = v15
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            if (0 >= 993) {
                v17 := []int{-333, -9223372036854775808}
                fmt.Println(v1, v2)
                _ = v17
            }
        }
        fmt.Println(v2, v4)
    } else {
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS1{F0: (fi18 + 2147483647), F1: ("foo bar" + ("a" + ("\n\t" + "")))}
            var v20 int16 = -1
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v21 := []int{980}
    if len(v21) > 0 {
        v21[0] = 2147483647
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    switch -912 {
    case -194:
        var v22 string = "a"
        _ = v22
    case 32767:
        for fi23 := 0; fi23 < 12; fi23++ {
            if false {
                if (!(18446744073709551615 != 4294967295)) {
                    var v24 bool = false
                    fmt.Println(v24, v24, v1)
                    _ = v24
                } else {
                    var v25 int = (fi23 * 42)
                    v1 = v1
                    var v26 float64 = 3.14159
                    _ = v25
                    _ = v26
                }
                v2 = (v2 % (((v2 ^ v4)) | 1))
            } else {
                var v27 int32 = -2147483648
                _ = v27
            }
            v2 = (v4 | v2)
            _ = fi23
        }
    default:
        v2 += v4
    }
    var v28 uint = uint(v4)
    v21 = append(v21, 9223372036854775807)
    v21 = append(v21, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
}
