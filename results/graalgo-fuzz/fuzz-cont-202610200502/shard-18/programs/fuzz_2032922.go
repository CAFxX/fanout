package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    v2 := []int{103}
    v3 := make(map[uint32]string)
    var v4 uint64 = 646
    v5 := make(map[int]int)
    v6 := make(map[int]uint32)
    fmt.Println(v4)
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 12; fi8++ {
        if (false && false) {
            var v9 uint16 = 2
            switch ("foo bar" + v7) {
            case "a":
                switch "a" {
                case "foo bar":
                    var v10 int64 = 42
                    _ = v10
                case "hello":
                    var v11 uint64 = (v4 >> ((42) & 63))
                    _ = v11
                }
            case "foo bar":
                for fi12 := 0; fi12 < 11; fi12++ {
                    v9 = v9
                    _ = fi12
                }
            case "hello":
                var v13 int8 = -77
                _ = v13
            default:
                v2 = append(v2, fi8)
            }
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 float64 = float64((0.1 * (float32(((1.5 + 0.0) - (((0.0 + (-2.25 * float64((float32(0.0) / -2.25)))) + -2.25) / 1.5))) / 0.1)))
                _ = fi14
                _ = v15
            }
            _ = v9
        }
        _ = fi8
    }
    switch v7 {
    case "\n\t":
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 uint32 = 1
            if true {
                v18 := v5[int(v4)]
                _ = v18
            }
            _ = fi16
            _ = v17
        }
    }
    v4 = 4294967295
    v2 = append(v2, 4294967295)
    if (false && false) {
        v4 -= (v4 - 307)
        v4 -= v4
    }
    fmt.Println(v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
