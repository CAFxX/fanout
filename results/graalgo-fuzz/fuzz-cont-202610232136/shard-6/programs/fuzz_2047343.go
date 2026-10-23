package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint64) uint16 {
    v7 := []int{4294967295, 4294967295, 4294967295}
    var v8 int64 = 0
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int16, p1 string) int32 {
    var v9 float64 = 1e10
    var v10 int8 = 63
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 371
}

func main() {
    var v1 uint64 = 0
    var v2 float32 = 0.1
    v3 := make(map[uint32]bool)
    v4 := []int{32767, -9223372036854775808, 674}
    var v5 int64 = 1
    var v6 int64 = v5
    switch "foo bar" {
    case "hello":
        if (1e10 >= -2.25) {
            var v11 int64 = -388
            _ = v11
        } else {
            for fi12 := 0; fi12 < 20; fi12++ {
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(fi12, v1, v1)
                case "\n\t":
                    var v13 int8 = -1
                    _ = v13
                default:
                    v14 := v3[2147483647]
                    _ = v14
                }
                switch fi12 {
                case 996:
                    v15 := GenS1{F0: 1}
                    _ = v15
                case 691:
                    var v16 uint16 = genFn0(2)
                    _ = v16
                default:
                    v2 = 3.14159
                }
                _ = fi12
            }
            var v17 uint32 = 919
            var v18 uint8 = 202
            _ = v17
            _ = v18
        }
    default:
        if len(v4) > 0 {
            v4[0] = 0
        }
    }
    v1 = (^(v1))
    var v19 uint32 = 2147483648
    var v20 uint8 = 127
    v5 += (v6 / ((1) | 1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 = int64(v20)
    if (!(0.1 == (0.1 + 1e10))) {
        switch 4294967295 {
        case 128:
            switch 42 {
            case -850:
                var v21 uint8 = (^(v20))
                _ = v21
            }
        }
        switch 0 {
        case -2147483648:
            var v22 bool = false
            _ = v22
        case 4294967295:
            var v23 uint16 = 32767
            _ = v23
        case 127:
            var v24 float32 = (v2 - (((float32((0.0 * 1.5)) / v2) / ((float32((3.14159 + (1e10 + ((-2.25 * ((-2.25 + 0.0) - ((((3.14159 / ((1.5 * 3.14159) * 3.14159)) / 0.1) * ((float64(v2) * 0.1) / float64((v2 * 1e10)))) / 1.5))) * float64(v2))))) + (float32(float64(0.0)) / 1.5)) + v2)) - v2))
            _ = v24
        }
    } else {
        _ = len(v3)
        fmt.Println(v20, v5)
        v25 := GenS1{F0: 255}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
