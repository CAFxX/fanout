package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint8; F1 string }
func genFn0(p0 int64, p1 int, p2 bool) int8 {
    p2 = p2
    var v4 int16 = -32768
    v5 := GenS1{F0: (p0 / (((p0 - (p0 / (((p0 | (p0 / (((p0 / (((p0 >> ((873) & 63))) | 1))) | 1)))) | 1)))) | 1)), F1: 0}
    p0 -= p0
    var v6 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 105
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ""
    v3 := make(map[string]string)
    switch 221 {
    case 127:
        var v7 uint16 = 556
        _ = v7
    case -1:
        var v8 int64 = 42
        _ = v8
    case 32767:
        var v9 uint8 = 49
        _ = v9
    }
    if (0 <= 790) {
        if ((1e10 == 0.1) || false) {
            if (42 < -1) {
                v10 := GenS2{F0: -4, F1: -2147483648}
                for fi11 := 0; fi11 < 8; fi11++ {
                    v12 := GenS2{F0: 100, F1: fi11}
                    v13 := GenS2{F0: -1, F1: fi11}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                var v14 uint32 = 342
                _ = v10
                _ = v14
            } else {
                switch v2 {
                case "hello":
                    var v15 float32 = 0.1
                    _ = v15
                case "foo bar":
                    v16 := []int{1}
                    _ = v16
                }
                fmt.Println(v2)
            }
            v2 = (((("foo bar" + v2) + "") + (("hello" + "\n\t") + ((v2 + "\n\t") + v2))) + v2)
        }
    } else {
        delete(v3, v2)
    }
    var v17 uint32 = 571
    var v18 int32 = 2147483647
    switch 32768 {
    case 378:
        var v19 int = -2147483648
        _ = v19
    }
    v18 -= v18
    var v20 int = 4294967295
    fmt.Println(v17, v18, v18)
    v20 = (-(v20))
    if true {
        delete(v3, "")
        v20 += -339
        switch v20 {
        case -160:
            var v21 uint32 = (^(v17))
            _ = v21
        case -290:
            var v22 uint8 = 0
            _ = v22
        default:
            var v23 int32 = -1
            _ = v23
        }
    } else {
        v24 := []int{v20, v20, (v20 * v20), 9223372036854775807}
        v2 = v2
        switch "hello" {
        case "":
            v17 = 983
        case "a":
            v25 := GenS3{F0: 39, F1: "foo bar"}
            _ = v25
        case "hello":
            var v26 int16 = 32767
            _ = v26
        }
        _ = v24
    }
    if true {
        v18 = v18
        v27 := []int{v20, v20, v20}
        _ = v27
    }
    var v28 uint8 = 119
    var v29 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
}
