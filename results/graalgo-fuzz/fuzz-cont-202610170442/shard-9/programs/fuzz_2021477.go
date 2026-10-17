package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint32, p1 bool) uint8 {
    p1 = (1 > 1)
    for fi5 := 0; fi5 < 5; fi5++ {
        if (!(-1 >= -1)) {
            if p1 {
                switch fi5 {
                case 2147483647:
                    var v6 int64 = -9223372036854775808
                    _ = v6
                }
            }
        } else {
            var v7 bool = (26 > 127)
            if (!false) {
                v8 := []int{-2147483648, fi5}
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 uint8 = 128
                    switch -925 {
                    case -9223372036854775808:
                        var v11 float32 = 0.1
                        _ = v11
                    }
                    _ = fi9
                    _ = v10
                }
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 int = fi12
                    v13 -= (v13 + 127)
                    _ = fi12
                    _ = v13
                }
                _ = v8
            }
            _ = v7
        }
        fmt.Println(fi5)
        var v14 float32 = 3.14159
        _ = fi5
        _ = v14
    }
    p1 = p1
    v15 := []int{32767, 370}
    _ = p0
    _ = p1
    _ = v15
    return 127
}

func main() {
    var v1 int16 = -805
    v2 := make(map[uint32]bool)
    var v3 int8 = -1
    var v4 bool = (!false)
    v4 = ((v1 + (v1 + v1)) >= (^(v1)))
    var v16 string = ("\n\t" + "\n\t")
    v17 := GenS1{F0: 2147483648}
    v2[2147483648] = v4
    var v18 int32 = -2147483648
    var v19 int64 = 868
    switch "\n\t" {
    case "hello":
        fmt.Println(v1)
    case "":
        var v20 uint = 2
        _ = v20
    default:
        var v21 bool = (1e10 <= (-2.25 * float64((float32((0.1 - 0.0)) / (((1.5 - (1e10 + 0.0)) * ((((1.5 - 0.1) + 0.1) - 1.5) - 1e10)) / ((float32(0.0) / 1.5) - (1e10 + ((((((0.0 + (3.14159 / 1.0)) - (float32(((float64(float32((((3.14159 / 3.14159) + 1e10) * 3.14159))) / (-2.25 * 0.1)) / 1.5)) + -2.25)) * 0.0) * 0.0) - (0.1 + 1e10)) - 0.1))))))))
        _ = v21
    }
    if (538 < 2) {
        for fi22 := 0; fi22 < 16; fi22++ {
            v16 = v16
            v1 -= 42
            var v23 int64 = -739
            _ = fi22
            _ = v23
        }
    } else {
        v19 -= 746
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
