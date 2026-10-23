package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint }
type GenS5 struct { F0 uint8; F1 uint32; F2 bool }
func genFn0() uint8 {
    v4 := GenS1{F0: 0}
    var v5 int = -1
    for fi6 := 0; fi6 < 11; fi6++ {
        if (934 != 65535) {
            if (!(!(42 != -1))) {
                v7 := GenS1{F0: 54}
                v8 := GenS2{F0: 93}
                var v9 bool = true
                _ = v7
                _ = v8
                _ = v9
            }
            if (false && false) {
                switch "a" {
                case "\n\t":
                    v10 := GenS1{F0: 0}
                    _ = v10
                case "hello":
                    v11 := []int{(v5 & (fi6 - (fi6 & 736))), (^(v5)), (v5 * (-(v5))), (fi6 | 2147483647)}
                    _ = v11
                case "foo bar":
                    var v12 string = "a"
                    _ = v12
                }
                if ((true && (!true)) && (595 <= 1)) {
                    v5 -= fi6
                    var v13 bool = ("foo bar" != ("foo bar" + ""))
                    _ = v13
                } else {
                    v14 := []int{fi6, v5}
                    _ = v14
                }
            }
        }
        _ = fi6
    }
    v15 := []int{(v5 << ((443) & 63)), v5}
    switch 127 {
    case 32768:
        v16 := GenS3{F0: 65535, F1: -431}
        _ = v16
    default:
        var v17 uint8 = 255
        _ = v17
    }
    _ = v15
    _ = v4
    _ = v5
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 127
    var v3 uint64 = 1
    v18 := []int{417, 370, 441, 444}
    v19 := GenS4{F0: 2, F1: v2, F2: 18446744073709551615}
    v20 := GenS5{F0: v2, F1: 2, F2: (true && ("foo bar" <= ("\n\t" + "\n\t")))}
    var v21 string = ("foo bar" + ("hello" + ("" + (("foo bar" + ("foo bar" + ("hello" + ("hello" + ("a" + ("foo bar" + "")))))) + "foo bar"))))
    _ = len(v18)
    v3 += (^(v3))
    var v22 bool = true
    if (2 > 42) {
        _ = len(v18)
        var v23 uint = 4294967295
        fmt.Println(v21, v3)
        _ = v23
    }
    delete(v1, v21)
    var v24 int32 = 2147483647
    fmt.Println(v3, v2)
    v22 = (!false)
    _ = len(v1)
    v25 := GenS2{F0: 1}
    v26 := v1[v21]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
