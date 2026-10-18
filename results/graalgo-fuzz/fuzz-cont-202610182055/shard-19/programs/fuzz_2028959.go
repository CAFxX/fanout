package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0() int16 {
    if (!(-617 != 405)) {
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 uint64 = 1
            _ = fi5
            _ = v6
        }
        var v7 int16 = 42
        fmt.Println(v7)
        _ = v7
    }
    var v8 string = "foo bar"
    var v9 uint8 = 1
    if false {
        v10 := GenS1{F0: 330}
        var v11 string = "\n\t"
        _ = v10
        _ = v11
    } else {
        var v12 uint16 = 65535
        v13 := GenS2{F0: (v12 % ((32767) | 1)), F1: 9223372036854775807}
        for fi14 := 0; fi14 < 19; fi14++ {
            if true {
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 uint16 = (^(v12))
                    var v17 float64 = float64((0.1 / 1.0))
                    if true {
                        v12 += 101
                        var v18 uint8 = v9
                        _ = v18
                    } else {
                        v19 := GenS3{F0: (v12 % (((^(v12))) | 1)), F1: 42}
                        v20 := []int{-81, fi14}
                        v21 := GenS1{F0: 2}
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    var v22 int16 = -957
                    var v23 uint8 = (v9 ^ 2)
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v22
                    _ = v23
                }
                var v24 uint32 = 308
                fmt.Println(v12)
                _ = v24
            } else {
                switch fi14 {
                case -9223372036854775808:
                    fmt.Println(v12, v8)
                case 128:
                    var v25 float64 = 0.1
                    _ = v25
                }
                var v26 int32 = 546
                _ = v26
            }
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 bool = true
                v12 = (v12 / ((v12) | 1))
                v29 := GenS2{F0: (v12 | v12), F1: (^(fi27))}
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi14
        }
        _ = v12
        _ = v13
    }
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 float64 = 3.14159
    var v2 float64 = v1
    v3 := []int{535, 1}
    var v4 int = -322
    v1 -= 0.1
    var v30 int32 = 273
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 bool = false
        switch "" {
        case "a":
            var v33 uint8 = 255
            _ = v33
        case "foo bar":
            v3 = append(v3, (fi31 << ((764) & 63)))
        case "hello":
            fmt.Println(v32, v4)
        default:
            fmt.Println(v2, v4, v32)
        }
        _ = fi31
        _ = v32
    }
    var v34 uint = 18446744073709551615
    fmt.Println(v30, v30, v30)
    var v35 uint64 = 572
    _ = len(v3)
    v30 = v30
    v34 -= 977
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
}
