package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0() int8 {
    var v6 float32 = ((1.5 + -2.25) * 1e10)
    var v7 int32 = 1
    _ = v6
    _ = v7
    return 43
}

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]bool)
    var v3 float64 = ((float64((((0.0 / (3.14159 / float32(1.5))) - 3.14159) - (float32(0.0) + (1e10 * (1e10 + 1.5))))) + float64((1e10 * -2.25))) / 1.5)
    v4 := make(map[int]bool)
    v5 := make(map[string]uint32)
    v1 = (v1 / ((v1) | 1))
    var v8 float64 = (((float64(1e10) / v3) - v3) - (0.1 / ((v3 / float64(float32(v3))) * v3)))
    switch 560 {
    case -556:
        if false {
            v3 = ((1e10 - ((float64(-2.25) / float64((3.14159 / float32(v3)))) / 1e10)) * v3)
            switch 1 {
            case -1:
                v9 := GenS1{F0: 0, F1: 32767}
                _ = v9
            case 32767:
                for fi10 := 0; fi10 < 5; fi10++ {
                    fmt.Println(fi10)
                    if false {
                        _ = len(v2)
                    }
                    if true {
                        v11 := GenS2{F0: true}
                        v12 := GenS2{F0: (18446744073709551615 <= 1)}
                        _ = v11
                        _ = v12
                    } else {
                        v13 := GenS3{F0: 2, F1: 1000}
                        _ = v13
                    }
                    _ = fi10
                }
            }
        }
    case 127:
        v1 += (v1 / (((v1 - 0)) | 1))
    }
    var v14 bool = true
    v15 := []int{32767, 2147483647}
    v8 += v8
    fmt.Println(v14, v8)
    if v14 {
        var v16 float64 = v3
        switch ("hello" + ("\n\t" + "foo bar")) {
        case "\n\t":
            v17 := GenS2{F0: v14}
            _ = v17
        case "a":
            if (v1 <= v1) {
                if len(v15) > 0 {
                    v15[0] = 128
                }
                fmt.Println(v3, v3, v8)
            }
        }
        _ = v16
    }
    v14 = ((!(true || (float64(0.0) > 0.0))) || (!(-728 >= 1)))
    v1 = v1
    if (("hello" == "\n\t") || v14) {
        if len(v15) > 0 {
            v15[0] = 42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
