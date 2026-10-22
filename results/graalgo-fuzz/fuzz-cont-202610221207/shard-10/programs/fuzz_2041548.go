package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 bool; F1 string; F2 uint8 }

func main() {
    var v1 int16 = -781
    v2 := []string{}
    v3 := []int{4294967295, 4294967295}
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]uint32)
    v1 += (^(v1))
    v2 = append(v2, ("" + (("" + ("a" + "hello")) + "a")))
    v1 = (v1 / (((v1 - v1)) | 1))
    switch "\n\t" {
    case "a":
        switch "\n\t" {
        case "a":
            v3 = append(v3, 128)
        case "\n\t":
            switch "" {
            case "\n\t":
                v1 += -32768
            default:
                for fi7 := 0; fi7 < 6; fi7++ {
                    v1 = v1
                    _ = fi7
                }
            }
        }
    case "":
        v1 = (v1 & (v1 + v1))
    case "foo bar":
        var v8 uint16 = 42
        _ = v8
    default:
        fmt.Println(v1)
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint = 255
        v1 = 1000
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 147, F1: 2147483647}
    v1 -= v1
    var v12 uint64 = 42
    var v13 uint64 = v12
    v14 := GenS1{F0: 1, F1: 0}
    var v15 int64 = 42
    fmt.Println(v12)
    var v16 uint16 = 2
    if ("" != ("a" + ("a" + (("foo bar" + "\n\t") + ("a" + "\n\t"))))) {
        v17 := GenS1{F0: 1, F1: 1}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v17
    } else {
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 float32 = 0.1
            var v20 uint16 = v16
            for fi21 := 0; fi21 < 6; fi21++ {
                v12 -= 523
                var v22 int32 = -2147483648
                v23 := v5[0]
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    if false {
        fmt.Println(v12)
        v3 = append(v3, -2147483648)
    } else {
        if true {
            var v24 int = 0
            _ = v24
        } else {
            v25 := GenS2{F0: (((!(!true)) || false) && (127 >= 0)), F1: "hello", F2: 75}
            _ = v25
        }
        if false {
            var v26 uint32 = 43
            _ = v26
        }
        v15 += 0
    }
    v27 := GenS2{F0: true, F1: "a", F2: 190}
    v13 += v13
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
