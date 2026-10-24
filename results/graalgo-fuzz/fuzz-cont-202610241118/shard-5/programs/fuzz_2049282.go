package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int; F1 uint16 }

func main() {
    v1 := []string{("\n\t" + ("a" + "")), "foo bar", (("" + "") + ("\n\t" + ("\n\t" + "foo bar")))}
    v2 := []uint32{2, 339, 2147483648, 2147483648}
    var v3 int64 = 409
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    v3 -= v3
    for fi6 := 0; fi6 < 7; fi6++ {
        switch "" {
        case "a":
            v3 = (v3 & 256)
        case "foo bar":
            var v7 float64 = 3.14159
            _ = v7
        case "":
            var v8 uint64 = 1
            _ = v8
        default:
            var v9 uint16 = 32767
            _ = v9
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    var v10 int32 = 2147483647
    var v11 uint32 = 1
    var v12 int64 = v3
    v13 := GenS1{F0: false, F1: 2}
    if true {
        fmt.Println(v12, v3, v10)
        for fi14 := 0; fi14 < 15; fi14++ {
            v3 = v3
            v3 -= (v12 / (((v12 ^ -1)) | 1))
            v10 -= 0
            _ = fi14
        }
    } else {
        switch 128 {
        case 127:
            v12 -= (^(v3))
        case -879:
            _ = len(v4)
        default:
            if false {
                var v15 string = "a"
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 float64 = -2.25
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
        }
        switch "a" {
        case "hello":
            v10 = (v10 | (v10 << ((42) & 63)))
        case "foo bar":
            v11 += (v11 / ((v11) | 1))
        case "\n\t":
            if false {
                var v18 int = 983
                _ = v18
            }
        default:
            var v19 int16 = -1
            _ = v19
        }
    }
    var v20 uint32 = (^(v11))
    delete(v4, "a")
    v11 = 4294967295
    for fi21 := 0; fi21 < 17; fi21++ {
        if len(v1) > 0 {
            v1[0] = ("\n\t" + ("hello" + ("a" + "\n\t")))
        }
        var v22 int8 = 108
        _ = fi21
        _ = v22
    }
    if ((-1 >= -1) && true) {
        v23 := GenS2{F0: 32768, F1: 65535}
        if ((4294967295 < 9223372036854775807) || ((("foo bar" + "\n\t") + (("\n\t" + "") + "hello")) >= "foo bar")) {
            v3 = int64(v10)
            fmt.Println(v11, v3)
            var v24 int8 = 100
            _ = v24
        } else {
            if false {
                v3 -= -417
                v2 = append(v2, v20)
                var v25 int = -2147483648
                _ = v25
            }
            v12 -= v3
        }
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
