package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }

func main() {
    v1 := []int{4294967295, 32767, 1}
    v2 := []uint32{}
    var v3 uint32 = 4294967295
    var v4 int8 = -26
    var v5 int = -2147483648
    var v6 uint16 = 605
    fmt.Println(v5, v4)
    v7 := GenS1{F0: v6, F1: 2147483647, F2: "foo bar"}
    if len(v1) > 0 {
        v1[0] = v5
    }
    v3 = v3
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v6, v3, fi8)
        _ = fi8
    }
    _ = len(v2)
    var v9 int32 = 2147483647
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v6, v5)
    v1 = append(v1, 893)
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            fmt.Println(v9)
            _ = len(v2)
            _ = fi11
        }
        var v12 int = v5
        fmt.Println(v3)
        _ = fi10
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        if ((^(v3)) == (v3 + (v3 << ((658) & 63)))) {
            v3 += (^(v3))
        } else {
            if (false || true) {
                var v14 uint16 = 65535
                switch "foo bar" {
                case "":
                    var v15 uint16 = 1
                    _ = v15
                case "\n\t":
                    v1 = append(v1, -9223372036854775808)
                case "a":
                    fmt.Println(v5, v9)
                }
                fmt.Println(fi13)
                _ = v14
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 = append(v1, -2147483648)
            }
            var v16 uint8 = 26
            if len(v2) > 0 {
                v2[0] = (v3 / ((v3) | 1))
            }
            _ = v16
        }
        _ = fi13
    }
    v17 := GenS1{F0: 32767, F1: (^(v9)), F2: "foo bar"}
    var v18 uint8 = 127
    v18 += v18
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 float64 = ((3.14159 + ((((0.0 / (1e10 / 1e10)) / float64((3.14159 - float32(float64((-2.25 * (0.0 - (((((1.5 * 0.1) * (float32((3.14159 - (0.1 * 3.14159))) + ((0.1 - 0.0) - 0.1))) / float32(3.14159)) / -2.25) - 0.0)))))))) - ((((1e10 * 1.5) - (float64((1e10 + 0.0)) + 1.5)) * 3.14159) / (((float64((0.0 - -2.25)) * 1.5) - 0.1) + float64(-2.25)))) * ((-2.25 - (float64(0.0) / float64((float32(float64(((0.0 + 0.1) - 1.5))) + (float32(1.5) + float32(3.14159)))))) * (0.1 - (0.1 + 1e10))))) * float64(-2.25))
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        v9 -= -2147483648
        fmt.Println(v6)
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
