package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 int = -2147483648
    v3 := make(map[string]uint32)
    var v4 int8 = 1
    if (0 >= -1) {
        v2 = (v2 - v2)
        if (("" + "") == (((("foo bar" + "\n\t") + "\n\t") + "hello") + "hello")) {
            fmt.Println(v2)
            if ((!true) && false) {
                var v5 uint32 = 0
                for fi6 := 0; fi6 < 19; fi6++ {
                    v7 := GenS1{F0: 255}
                    v8 := []int{(fi6 % ((-2147483648) | 1)), (fi6 >> ((53) & 63)), v2, v2}
                    v2 = fi6
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                _ = v5
            } else {
                for fi9 := 0; fi9 < 16; fi9++ {
                    if true {
                        v10 := GenS2{F0: 475}
                        _ = v10
                    }
                    _ = fi9
                }
            }
        } else {
            v2 = v2
            v11 := GenS2{F0: 1}
            v2 = 128
            _ = v11
        }
    } else {
        v12 := []int{}
        var v13 float32 = float32(float64(-2.25))
        _ = v12
        _ = v13
    }
    switch ("\n\t" + "") {
    case "":
        switch v2 {
        case 128:
            var v14 int = (v2 ^ 127)
            _ = v14
        case 127:
            v15 := v1["a"]
            _ = v15
        default:
            v16 := []int{v2}
            _ = v16
        }
    default:
        v2 = (v2 % ((-978) | 1))
    }
    switch (v2 | v2) {
    case -2147483648:
        fmt.Println(v2, v4)
    }
    v4 = (v4 & (v4 | v4))
    var v17 uint16 = 32767
    var v18 int16 = -25
    v19 := GenS3{F0: (false || (1.5 < 1e10))}
    v20 := GenS2{F0: 2}
    var v21 uint16 = (v17 & v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
