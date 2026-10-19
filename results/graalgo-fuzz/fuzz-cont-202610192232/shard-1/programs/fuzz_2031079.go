package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int16; F1 uint32 }
func genFn0(p0 int) int64 {
    p0 = p0
    p0 += p0
    p0 -= p0
    p0 = 967
    var v5 int16 = -1
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 int8 = 100
    var v2 int64 = -657
    var v3 uint16 = 156
    var v4 int16 = 32767
    var v6 uint8 = 180
    fmt.Println(v1)
    var v7 uint64 = 1
    var v8 uint = 42
    v9 := []int{127, -963, 4294967295}
    if ((("hello" + "\n\t") + "hello") <= (("" + ("\n\t" + ("a" + ("foo bar" + "")))) + "a")) {
        if false {
            if true {
                for fi10 := 0; fi10 < 16; fi10++ {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = fi10
                }
            } else {
                if false {
                    v3 += v3
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v11 uint32 = 2
                    _ = v11
                }
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            }
            for fi12 := 0; fi12 < 7; fi12++ {
                v13 := GenS1{F0: (^(v3)), F1: fi12, F2: v7}
                v8 = v8
                _ = fi12
                _ = v13
            }
            var v14 uint32 = 0
            _ = v14
        } else {
            v15 := GenS2{F0: 1, F1: -1}
            _ = v15
        }
    }
    switch "hello" {
    case "\n\t":
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 float32 = (float32(float64((float32(0.1) * (3.14159 - float32(float64(((1.5 * 1.5) - (((0.0 - 1e10) / (-2.25 + 1e10)) / 1.0)))))))) / (1.5 + 3.14159))
            var v18 uint64 = uint64(v8)
            _ = fi16
            _ = v17
            _ = v18
        }
    case "foo bar":
        var v19 int64 = v2
        _ = v19
    case "hello":
        v20 := GenS3{F0: 170, F1: 2147483648}
        _ = v20
    default:
        if len(v9) > 0 {
            v9[0] = 4294967295
        }
    }
    var v21 uint8 = v6
    v22 := GenS3{F0: (v4 * v4), F1: 2147483647}
    v4 = int16(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
