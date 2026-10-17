package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint32 }
type GenS2 struct { F0 int32; F1 string; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 bool) string {
    v4 := GenS1{F0: 921, F1: ("a" + "foo bar"), F2: 1}
    var v5 int8 = -1
    v6 := GenS1{F0: 32767, F1: ((("hello" + "hello") + "foo bar") + "hello"), F2: 0}
    switch "\n\t" {
    case "hello":
        v7 := []int{}
        _ = v7
    case "foo bar":
        var v8 bool = false
        _ = v8
    case "\n\t":
        var v9 float32 = 1e10
        _ = v9
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (("\n\t" + "\n\t") + ("a" + "a"))
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 162
    var v3 uint32 = uint32(v2)
    if (!(-293 > 9223372036854775807)) {
        v1 = v1
    }
    if (34 == -9223372036854775808) {
        switch 32767 {
        case 2147483647:
            if ((!((float64(0.1) * 1e10) != -2.25)) && (!false)) {
                v10 := GenS2{F0: 0, F1: "foo bar", F2: v2}
                if true {
                    v11 := GenS1{F0: 2147483647, F1: "hello", F2: (v3 + v3)}
                    v3 = (v3 * (^(v3)))
                    if ("hello" != "") {
                        var v12 uint64 = 1
                        _ = v12
                    }
                    _ = v11
                } else {
                    var v13 int8 = -49
                    _ = v13
                }
                var v14 string = "\n\t"
                _ = v10
                _ = v14
            } else {
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := []int{9223372036854775807}
                    var v17 uint32 = 2147483647
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
        case 4294967295:
            if true {
                fmt.Println(v2, v1, v3)
                fmt.Println(v2, v2, v3)
            }
        case 0:
            switch "foo bar" {
            case "a":
                v2 -= (^(v2))
            }
        }
        v1 += (^(v1))
    }
    v18 := []int{829, 128, 9223372036854775807, 881}
    if (2147483647 != -681) {
        v19 := GenS3{F0: (v2 % ((v2) | 1)), F1: 65535, F2: (v2 & 88)}
        v1 -= v1
        _ = v19
    } else {
        switch -391 {
        case -387:
            var v20 uint8 = (v2 | v2)
            _ = v20
        }
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        var v21 bool = true
        _ = v21
    }
    fmt.Println(v2, v2, v3)
    v22 := GenS1{F0: 1, F1: (("" + "") + ("foo bar" + "")), F2: 2147483648}
    var v23 uint = (^(v1))
    if false {
        var v24 int8 = 126
        _ = v24
    } else {
        v1 -= 794
        fmt.Println(v2)
        if len(v18) > 0 {
            v18[0] = 128
        }
    }
    v25 := GenS4{F0: -128}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
}
