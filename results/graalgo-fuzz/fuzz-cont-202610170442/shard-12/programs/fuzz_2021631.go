package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 string }
func genFn0(p0 uint16) uint8 {
    var v7 uint32 = 1
    var v8 int8 = -128
    var v9 int32 = 0
    switch ("a" + "") {
    case "hello":
        v8 += -81
    case "":
        switch "hello" {
        case "hello":
            v10 := GenS1{F0: -418}
            _ = v10
        }
    case "foo bar":
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := []int{(fi11 | 0), int(v7), 0, (fi11 - fi12)}
                if false {
                    v8 = (v8 - v8)
                } else {
                    v9 += v9
                    v7 += 65535
                }
                _ = fi12
                _ = v13
            }
            if (88 != 164) {
                var v14 string = ("hello" + "a")
                _ = v14
            }
            v15 := GenS1{F0: fi11}
            _ = fi11
            _ = v15
        }
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 58
}
func genFn1() int {
    switch 93 {
    case -2147483648:
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 int = (fi16 % (((fi16 + -553)) | 1))
            var v18 bool = false
            _ = fi16
            _ = v17
            _ = v18
        }
    case -9223372036854775808:
        var v19 int16 = 32767
        _ = v19
    default:
        v20 := GenS2{F0: 42, F1: "hello"}
        _ = v20
    }
    switch "" {
    case "a":
        var v21 bool = false
        _ = v21
    default:
        var v22 string = "foo bar"
        _ = v22
    }
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 409
    v3 := make(map[int]int)
    v4 := []string{}
    var v5 float32 = 0.0
    var v6 float32 = 0.1
    var v23 float32 = v6
    var v24 int64 = 42
    v3[genFn1()] = -2147483648
    switch "foo bar" {
    case "hello":
        if false {
            v6 = v5
            fmt.Println(v24, v6)
        }
    case "foo bar":
        if (!true) {
            delete(v1, 2147483648)
        }
    case "a":
        _ = len(v3)
    }
    v25 := GenS2{F0: v24, F1: "a"}
    v23 -= ((float32(1.5) * (float32((0.0 / -2.25)) / float32(float64(v5)))) + -2.25)
    var v26 int8 = 42
    v2 = (v2 & (v2 >> ((436) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
