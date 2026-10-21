package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    v3 := make(map[int]string)
    v2 = (((3.14159 - (0.1 + (float32(3.14159) - 1.5))) * (1.5 + float32(float64(((-2.25 / 1.0) / -2.25))))) != 1.5)
    v2 = v2
    fmt.Println(v2)
    v4 := []int{9223372036854775807, 2147483647, 32767, -204}
    fmt.Println(v2)
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := GenS1{F0: false, F1: "a"}
            var v8 bool = (0 != 0)
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        _ = len(v3)
        _ = fi9
    }
    v2 = (235 == -2147483648)
    var v10 uint8 = 2
    v2 = ((0 >= 100) || v2)
    for fi11 := 0; fi11 < 17; fi11++ {
        if (0 == -1) {
            for fi12 := 0; fi12 < 19; fi12++ {
                fmt.Println(fi12, v2)
                v13 := GenS1{F0: v2, F1: "\n\t"}
                fmt.Println(fi12, v2)
                var v14 uint = 65535
                switch "hello" {
                case "hello":
                    v14 += v14
                default:
                    delete(v1, (fi11 % ((-1) | 1)))
                }
                fmt.Println(fi12, fi12)
                if (-9223372036854775808 <= -9223372036854775808) {
                    var v15 int16 = -783
                    var v16 uint16 = 2
                    v17 := GenS1{F0: v2, F1: "a"}
                    _ = v15
                    _ = v16
                    _ = v17
                }
                var v18 int16 = int16(v14)
                var v19 int64 = 42
                _ = fi12
                _ = v13
                _ = v14
                _ = v18
                _ = v19
            }
        }
        var v20 int = fi11
        _ = fi11
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
