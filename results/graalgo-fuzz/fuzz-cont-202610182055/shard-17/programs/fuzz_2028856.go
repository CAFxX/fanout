package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int64; F2 string }
type GenS3 struct { F0 int; F1 uint8; F2 uint16 }
func genFn0(p0 uint8, p1 string, p2 uint32) uint {
    v7 := []int{4294967295}
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 1
    var v6 float32 = 1.5
    v6 = 1.5
    v8 := GenS1{F0: 65535}
    var v9 int16 = 0
    if ((42 > 2147483647) || false) {
        var v10 int16 = (^(v9))
        switch "" {
        case "hello":
            for fi11 := 0; fi11 < 1; fi11++ {
                fmt.Println(v10)
                switch "hello" {
                case "hello":
                    v12 := GenS1{F0: 735}
                    _ = v12
                default:
                    v10 -= (v10 / ((1) | 1))
                }
                _ = fi11
            }
        default:
            var v13 uint16 = 886
            _ = v13
        }
        if (!(-2147483648 >= 42)) {
            var v14 uint = 18446744073709551615
            fmt.Println(v5, v9)
            v15 := GenS2{F0: 1, F1: 312, F2: "foo bar"}
            _ = v14
            _ = v15
        }
        _ = v10
    } else {
        v16 := GenS3{F0: 4294967295, F1: 42, F2: 255}
        var v17 bool = false
        var v18 string = ("\n\t" + "hello")
        _ = v16
        _ = v17
        _ = v18
    }
    if true {
        v6 = float32(1e10)
        var v19 int16 = v9
        _ = v19
    }
    v5 = v5
    v20 := GenS3{F0: -1, F1: 1, F2: 32767}
    var v21 float32 = v6
    var v22 float32 = v21
    v23 := GenS2{F0: -49, F1: 0, F2: ("\n\t" + "foo bar")}
    var v24 int16 = 1
    if (9223372036854775807 > -666) {
        var v25 uint = 2
        fmt.Println(v9)
        _ = v25
    } else {
        if true {
            v6 = (float32(1e10) + v6)
            v1[828] = 2
            var v26 string = ""
            _ = v26
        } else {
            v27 := GenS2{F0: 127, F1: -621, F2: "\n\t"}
            _ = v27
        }
        if (48 > 1) {
            var v28 string = "a"
            v24 -= 259
            _ = v28
        }
        v1[42] = 640
    }
    var v29 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
