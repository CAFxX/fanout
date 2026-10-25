package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }

func main() {
    var v1 int8 = 104
    v2 := make(map[int]int)
    v3 := []int{127}
    v4 := make(map[uint32]string)
    var v5 uint = 501
    var v6 float64 = float64(-2.25)
    var v7 int64 = -47
    var v8 int32 = 2147483647
    var v9 bool = true
    v10 := GenS1{F0: (-(v7)), F1: (v5 ^ 2), F2: v8}
    if false {
        var v11 float32 = 3.14159
        var v12 int64 = -1
        for fi13 := 0; fi13 < 17; fi13++ {
            if (1 > 0) {
                switch ("" + "") {
                case "":
                    var v14 bool = v9
                    _ = v14
                case "hello":
                    switch 973 {
                    case 42:
                        var v15 int64 = v12
                        _ = v15
                    case 123:
                        var v16 int16 = 32767
                        _ = v16
                    case 9223372036854775807:
                        fmt.Println(fi13, v8)
                    default:
                        var v17 uint8 = 2
                        _ = v17
                    }
                case "foo bar":
                    switch 128 {
                    case 0:
                        delete(v2, (fi13 / ((fi13) | 1)))
                    case -844:
                        v5 = v5
                    }
                default:
                    var v18 uint64 = 4294967295
                    _ = v18
                }
                var v19 string = "\n\t"
                if len(v3) > 0 {
                    v3[0] = -2147483648
                }
                _ = v19
            }
            fmt.Println(v8, v5)
            switch "" {
            case "a":
                fmt.Println(v7, v1, v6)
            case "hello":
                var v20 uint8 = 0
                _ = v20
            default:
                fmt.Println(v6)
            }
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    v21 := GenS1{F0: (v7 << (((v5 / (((^(v5))) | 1))) & 63)), F1: 0, F2: (v8 % (((v8 & 1)) | 1))}
    v22 := v2[0]
    var v23 int32 = (v8 & v8)
    var v24 bool = (!true)
    v25 := GenS1{F0: 1, F1: (v5 | v5), F2: v8}
    for fi26 := 0; fi26 < 18; fi26++ {
        v7 += v7
        fmt.Println(v9, v7)
        _ = fi26
    }
    if v24 {
        var v27 float64 = float64(3.14159)
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 float64 = (v6 * 3.14159)
            if true {
                v4[1] = (("" + "a") + ("a" + "foo bar"))
                v22 = -612
                fmt.Println(v8)
            } else {
                v7 = (v7 % ((-9223372036854775808) | 1))
            }
            v30 := GenS1{F0: -849, F1: 255, F2: 0}
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = v27
    }
    switch "hello" {
    case "a":
        v5 = 660
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 uint8 = 127
        v5 -= 432
        v24 = (1e10 > 3.14159)
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        v9 = true
        v7 -= v7
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
