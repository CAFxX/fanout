package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint32; F1 int; F2 int16 }
func genFn0() uint16 {
    var v5 int32 = -1
    switch -1 {
    case 653:
        if ((3.14159 + ((3.14159 / 0.1) * float64(1.5))) != (((1e10 + (0.0 / 0.1)) - (-2.25 / (0.1 - (0.1 + 0.1)))) * 1e10)) {
            var v6 uint16 = 42
            v6 = v6
            v5 -= -1
            _ = v6
        } else {
            fmt.Println(v5)
        }
    case -2147483648:
        v7 := []int{-2147483648, 4294967295, 9223372036854775807}
        _ = v7
    case 0:
        var v8 uint64 = 4294967295
        _ = v8
    }
    _ = v5
    return 32767
}

func main() {
    var v1 int = -502
    var v2 string = ("foo bar" + ("a" + "foo bar"))
    v3 := make(map[int]int)
    var v4 int64 = 794
    v9 := []int{v1}
    var v10 int32 = 42
    v2 = (v2 + "\n\t")
    v11 := GenS1{F0: 4294967295}
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: 2147483647}
        var v14 int64 = v4
        _ = fi12
        _ = v13
        _ = v14
    }
    if true {
        switch 42 {
        case -2147483648:
            v2 = "foo bar"
        case -51:
            v3[v1] = 128
        }
        var v15 int8 = -7
        if false {
            v16 := GenS1{F0: 42}
            var v17 uint64 = 2
            if (139 <= 1) {
                if (v10 < v10) {
                    v2 = "a"
                    delete(v3, (v1 >> ((239) & 63)))
                    if (1e10 != 1.5) {
                        v1 += (^(v1))
                        v15 = int8(v17)
                    }
                }
                v3[int(v4)] = 128
                v2 = ("" + "a")
            }
            _ = v16
            _ = v17
        } else {
            if false {
                if (0 >= 232) {
                    var v18 uint64 = 42
                    v19 := GenS2{F0: 32767, F1: -1}
                    _ = v18
                    _ = v19
                }
            } else {
                v2 = "\n\t"
            }
        }
        _ = v15
    }
    var v20 int8 = 42
    v10 -= (v10 % ((v10) | 1))
    switch v2 {
    case "\n\t":
        if len(v9) > 0 {
            v9[0] = (v1 - (v1 & int(v4)))
        }
    case "":
        delete(v3, 909)
    case "a":
        if (1 <= 2) {
            var v21 int32 = 2147483647
            v21 = v10
            _ = v21
        } else {
            fmt.Println(v2)
            v22 := GenS2{F0: 351, F1: -468}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 10; fi23++ {
            for fi24 := 0; fi24 < 17; fi24++ {
                v1 = 821
                fmt.Println(v1, v1)
                v25 := GenS3{F0: 996, F1: (v1 / ((fi23) | 1)), F2: 1}
                _ = fi24
                _ = v25
            }
            v4 = 0
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi23
        }
    }
    v26 := v3[v1]
    var v27 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
