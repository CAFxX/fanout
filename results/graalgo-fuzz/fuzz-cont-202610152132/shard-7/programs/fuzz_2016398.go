package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int; F1 int32 }

func main() {
    var v1 int32 = -2147483648
    v2 := []int{-2147483648}
    v3 := []int{}
    fmt.Println(v1)
    fmt.Println(v1)
    if (false && (-9223372036854775808 <= -1)) {
        var v4 uint64 = 4294967295
        if ((true && ((!(!true)) && ((!false) || ((((-2.25 / (float32((1e10 + ((3.14159 - 3.14159) - ((1e10 - (0.1 / 1.5)) * (1e10 + 0.1))))) / 0.1)) * float32(-2.25)) / ((((float32(3.14159) + 1e10) / (1e10 + 3.14159)) / 3.14159) / float32(-2.25))) < 1e10)))) && false) {
            fmt.Println(v4)
            v4 = 18446744073709551615
            var v5 int16 = 42
            _ = v5
        }
        var v6 uint16 = 32767
        _ = v4
        _ = v6
    } else {
        v1 = v1
        v1 -= (v1 | -102)
    }
    switch 32768 {
    case -1:
        var v7 uint = 797
        _ = v7
    }
    var v8 int64 = 0
    v1 += 1
    if (1 != -128) {
        var v9 int = -293
        _ = v9
    }
    if (false || false) {
        if true {
            fmt.Println(v1, v8)
            for fi10 := 0; fi10 < 6; fi10++ {
                if (42 <= 618) {
                    v11 := GenS1{F0: 235, F1: fi10}
                    fmt.Println(v8, v8, fi10)
                    var v12 bool = (!(100 > 1))
                    _ = v11
                    _ = v12
                }
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v13 bool = true
                    _ = v13
                } else {
                    v14 := GenS2{F0: (^(fi10)), F1: v1}
                    var v15 bool = false
                    var v16 uint = 902
                    _ = v14
                    _ = v15
                    _ = v16
                }
                var v17 int32 = (v1 & (v1 + v1))
                _ = fi10
                _ = v17
            }
        }
        if len(v2) > 0 {
            v2[0] = 112
        }
        var v18 uint = 255
        _ = v18
    } else {
        v8 += (v8 & (v8 - (v8 / ((v8) | 1))))
    }
    v8 = v8
    var v19 uint8 = 2
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v8
}
