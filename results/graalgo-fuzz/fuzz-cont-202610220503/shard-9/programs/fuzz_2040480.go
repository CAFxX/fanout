package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint64; F1 int8; F2 int16 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint) int32 {
    v4 := GenS1{F0: 1, F1: 1, F2: ""}
    p0 = 255
    p0 = p0
    _ = p0
    _ = v4
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    var v5 int32 = -2147483648
    fmt.Println(v5)
    v5 -= (v5 >> ((299) & 63))
    v5 += (v5 - v5)
    v6 := []int{127, 9223372036854775807, -407, -718}
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(v5)
        var v8 int8 = 100
        if false {
            v9 := GenS1{F0: 2147483648, F1: 4294967295, F2: "foo bar"}
            switch fi7 {
            case -9223372036854775808:
                v10 := GenS1{F0: 42, F1: (fi7 - (fi7 - int(v8))), F2: ((("a" + "hello") + "hello") + (("foo bar" + "\n\t") + ("foo bar" + "hello")))}
                _ = v10
            case 0:
                v2[(fi7 - -119)] = ((!true) || (!(((1.5 - 0.0) / (float64((3.14159 + (1.5 - 1.5))) + 0.1)) < 3.14159)))
            case 127:
                v11 := GenS2{F0: 65535, F1: 661}
                _ = v11
            default:
                var v12 int = (-(fi7))
                _ = v12
            }
            var v13 uint8 = 250
            _ = v13
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v14 float32 = float32((3.14159 - 1e10))
    _ = len(v3)
    for fi15 := 0; fi15 < 16; fi15++ {
        v6 = append(v6, fi15)
        _ = len(v6)
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := GenS3{F0: 823, F1: -128, F2: 54}
            v14 += 0.1
            switch 128 {
            case -9223372036854775808:
                v18 := GenS4{F0: 255}
                _ = v18
            }
            var v19 uint32 = 65535
            if false {
                var v20 uint8 = 12
                _ = v20
            }
            _ = fi16
            _ = v17
            _ = v19
        }
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
