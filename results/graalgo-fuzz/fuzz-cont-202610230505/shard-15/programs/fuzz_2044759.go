package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 293
    v2 := make(map[int]uint32)
    var v3 int8 = 100
    v4 := make(map[uint32]string)
    var v5 int8 = v3
    v5 += 100
    v2[9223372036854775807] = (v1 ^ 4294967295)
    v3 = v5
    _ = len(v2)
    if (float64(0.1) <= 3.14159) {
        v6 := []int{128, 9223372036854775807, 0, 32767}
        _ = v6
    }
    v5 += v3
    v1 += v1
    v5 += (v5 ^ (v3 << ((18446744073709551615) & 63)))
    if (!(!(1000 <= 900))) {
        var v7 int8 = v5
        _ = len(v2)
        _ = v7
    } else {
        if false {
            var v8 int8 = (v5 << ((1) & 63))
            _ = v8
        } else {
            v5 = (-(v3))
            for fi9 := 0; fi9 < 18; fi9++ {
                v3 += v3
                _ = fi9
            }
            if ((71 >= 2) || false) {
                fmt.Println(v5, v1, v3)
            } else {
                var v10 uint16 = uint16(v1)
                v4[(^(v1))] = ("\n\t" + "foo bar")
                var v11 bool = (0.0 >= 0.1)
                _ = v10
                _ = v11
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
