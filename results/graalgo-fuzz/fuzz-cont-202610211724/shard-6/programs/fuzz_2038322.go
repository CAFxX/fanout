package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 string }

func main() {
    var v1 string = (((("\n\t" + "") + ("\n\t" + "hello")) + "") + "\n\t")
    var v2 uint32 = 1
    v3 := make(map[string]uint32)
    var v4 string = (("hello" + "\n\t") + "foo bar")
    var v5 bool = (42 != 0)
    var v6 uint8 = 127
    if v5 {
        fmt.Println(v6, v5, v1)
        var v7 uint16 = 1
        _ = v7
    } else {
        v8 := GenS1{F0: 1, F1: 18446744073709551615, F2: (v4 + v1)}
        _ = v8
    }
    var v9 int64 = 853
    v6 -= v6
    var v10 uint16 = 1
    if false {
        switch "" {
        case "foo bar":
            var v11 bool = (!true)
            _ = v11
        case "a":
            var v12 int16 = 1
            _ = v12
        }
        v13 := GenS1{F0: uint16(v6), F1: 42, F2: (v1 + "foo bar")}
        _ = v13
    }
    v3["\n\t"] = (^(v2))
    v10 += (v10 + (v10 % ((uint16(v2)) | 1)))
    v14 := []int{0, -579, -1, 32768}
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 15; fi16++ {
            switch v1 {
            case "a":
                v17 := GenS1{F0: (v10 | (v10 + v10)), F1: 907, F2: "foo bar"}
                _ = v17
            case "foo bar":
                var v18 uint8 = v6
                _ = v18
            default:
                fmt.Println(fi16, fi15, v9)
            }
            var v19 int = fi15
            v1 = v1
            _ = fi16
            _ = v19
        }
        var v20 uint32 = v2
        var v21 float32 = 0.0
        _ = fi15
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        switch fi22 {
        case -512:
            v23 := GenS1{F0: (v10 + 65535), F1: 512, F2: ("hello" + v4)}
            _ = v23
        default:
            fmt.Println(v10, v5)
        }
        v24 := GenS1{F0: 65535, F1: 42, F2: v4}
        _ = fi22
        _ = v24
    }
    fmt.Println(v2, v2)
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
