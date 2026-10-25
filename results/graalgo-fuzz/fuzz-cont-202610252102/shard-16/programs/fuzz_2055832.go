package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int8 }
func genFn0(p0 string) uint16 {
    if (3.14159 > ((3.14159 / float64(float32(float64(1e10)))) / ((1e10 + ((0.1 / (0.1 + 0.1)) / 1.5)) - -2.25))) {
        switch p0 {
        case "\n\t":
            v6 := GenS1{F0: -9223372036854775808}
            _ = v6
        case "hello":
            v7 := []int{42, 4294967295}
            _ = v7
        default:
            if (64 > 645) {
                v8 := []int{32767, 24, -2147483648}
                _ = v8
            } else {
                switch p0 {
                case "foo bar":
                    p0 = p0
                }
            }
        }
        var v9 int32 = 240
        var v10 int = 2147483647
        _ = v10
        _ = v9
    } else {
        var v11 uint16 = 65535
        _ = v11
    }
    if false {
        p0 = "\n\t"
        p0 = p0
    }
    _ = p0
    return 0
}

func main() {
    var v1 int16 = -844
    var v2 uint8 = 0
    v3 := make(map[string]bool)
    var v4 uint32 = 616
    v5 := make(map[string]uint32)
    var v12 uint16 = 0
    v13 := GenS2{F0: "\n\t", F1: (v1 ^ v1)}
    var v14 uint8 = 127
    v15 := GenS1{F0: 9223372036854775807}
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS1{F0: 1}
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    var v19 uint16 = genFn0(("hello" + "a"))
    v20 := []int{1, 128}
    v21 := GenS3{F0: v12, F1: 9223372036854775807}
    var v22 int32 = 740
    v23 := GenS4{F0: -1}
    var v24 float32 = 3.14159
    v5[("foo bar" + (("a" + ("\n\t" + "foo bar")) + "a"))] = v4
    var v25 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
