package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint16; F1 string; F2 int8 }
type GenS4 struct { F0 int64; F1 uint32 }
func genFn0() bool {
    var v5 int16 = 1000
    var v6 uint16 = 974
    v7 := []int{}
    switch "\n\t" {
    case "\n\t":
        var v8 uint8 = 128
        _ = v8
    case "a":
        switch 0 {
        case -187:
            for fi9 := 0; fi9 < 1; fi9++ {
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch "hello" {
                    case "hello":
                        v5 += (v5 | (v5 + -32768))
                        _ = len(v7)
                        v6 = v6
                        v6 = (v6 % ((v6) | 1))
                    default:
                        v11 := GenS1{F0: -17, F1: 2, F2: false}
                        _ = v11
                    }
                    var v12 float64 = ((3.14159 + 0.1) * 0.1)
                    _ = fi10
                    _ = v12
                }
                if false {
                    if len(v7) > 0 {
                        v7[0] = (fi9 & -2147483648)
                    }
                    var v13 uint = 0
                    var v14 string = "foo bar"
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(fi9, v6)
                    v15 := GenS2{F0: (v6 * v6), F1: v6, F2: -24}
                    _ = v15
                }
                _ = fi9
            }
        }
    case "hello":
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS3{F0: v6, F1: (("foo bar" + "hello") + (("a" + ("" + ("a" + (("hello" + "\n\t") + "a")))) + "hello")), F2: 0}
            var v18 uint8 = 255
            v19 := GenS4{F0: 42, F1: 65535}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]bool)
    var v3 bool = false
    var v4 uint8 = 115
    var v20 int16 = 42
    v21 := []int{0, 2147483647, 2147483647}
    var v22 uint32 = 2147483648
    v4 = 128
    v21 = append(v21, int(v20))
    delete(v2, int(v4))
    if len(v21) > 0 {
        v21[0] = 453
    }
    v1 = -110
    for fi23 := 0; fi23 < 8; fi23++ {
        if (!v3) {
            v3 = v3
        }
        _ = fi23
    }
    v24 := GenS2{F0: 42, F1: 0, F2: 1}
    v3 = (!(int(v22) != -2147483648))
    var v25 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
