package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint32; F1 bool; F2 int64 }
func genFn0(p0 int64, p1 int8, p2 uint) int {
    var v4 uint8 = 192
    v5 := []int{2147483647, 128}
    var v6 uint64 = 2
    v7 := GenS1{F0: (^(v4))}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 476
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("foo bar" + "foo bar"), "a"}
    v3 := make(map[uint32]string)
    v8 := v3[2147483647]
    v2 = append(v2, v8)
    var v9 float32 = 1e10
    v8 = "\n\t"
    var v10 int16 = -1
    fmt.Println(v8, v9, v10)
    if true {
        if (!false) {
            v3[2147483647] = v8
            var v11 string = ""
            _ = v11
        } else {
            v8 = "foo bar"
            switch "hello" {
            case "hello":
                v8 = v8
            case "\n\t":
                if (false || true) {
                    switch v8 {
                    case "":
                        for fi12 := 0; fi12 < 11; fi12++ {
                            v9 = float32(-2.25)
                            v8 = ""
                            fmt.Println(v9, v8, v9)
                            _ = fi12
                        }
                    case "\n\t":
                        v13 := GenS1{F0: 1}
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 11; fi14++ {
                        fmt.Println(fi14, fi14)
                        var v15 int32 = 2147483647
                        _ = fi14
                        _ = v15
                    }
                    delete(v3, 167)
                }
            case "a":
                if (true || (0 >= 2)) {
                    var v16 int64 = 498
                    var v17 uint64 = 0
                    var v18 int64 = int64(v10)
                    _ = v16
                    _ = v17
                    _ = v18
                }
            default:
                v1[2] = -2147483648
            }
            v19 := GenS2{F0: true, F1: 42}
            _ = v19
        }
        v1[780] = 253
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 int8 = -128
            _ = fi20
            _ = v21
        }
    } else {
        v1[799] = -1
        var v22 int64 = 9223372036854775807
        if false {
            v10 += (^(v10))
            delete(v1, 4294967295)
        }
        _ = v22
    }
    if true {
        for fi23 := 0; fi23 < 11; fi23++ {
            v9 -= v9
            _ = fi23
        }
        v2 = append(v2, "hello")
    }
    v24 := GenS3{F0: 4294967295, F1: true, F2: -741}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v8
    _ = v9
}
