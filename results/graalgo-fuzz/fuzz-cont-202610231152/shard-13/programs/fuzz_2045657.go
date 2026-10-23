package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint, p1 uint) bool {
    v4 := []int{4294967295, 1}
    v4 = append(v4, -400)
    _ = p0
    _ = p1
    _ = v4
    return false
}

func main() {
    var v1 bool = ((-93 >= 42) || (!((float64((1e10 / float32((0.1 * -2.25)))) * 1e10) > 0.1)))
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v5 := GenS1{F0: 1}
    fmt.Println(v1)
    var v6 int32 = -2147483648
    var v7 int64 = 9223372036854775807
    var v8 int64 = (v7 + v7)
    switch "" {
    case "a":
        v9 := GenS2{F0: -99}
        _ = v9
    case "":
        v8 += v7
    }
    var v10 uint = 255
    v3[531] = 42
    fmt.Println(v6, v6)
    if false {
        v6 = v6
        for fi11 := 0; fi11 < 17; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 string = ""
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    }
    fmt.Println(v1)
    v14 := GenS2{F0: -1}
    v10 -= (v10 / ((v10) | 1))
    v7 += (v7 * (-(v7)))
    for fi15 := 0; fi15 < 14; fi15++ {
        _ = len(v2)
        var v16 uint8 = 232
        _ = fi15
        _ = v16
    }
    v8 += (v7 % ((9223372036854775807) | 1))
    switch "a" {
    case "":
        if v1 {
            if (-1 > int16(v6)) {
                fmt.Println(v6, v6, v6)
            }
        } else {
            for fi17 := 0; fi17 < 4; fi17++ {
                v1 = (!((v6 ^ (v6 ^ (v6 - (v6 - (v6 ^ v6))))) == v6))
                _ = fi17
            }
            var v18 bool = (191 > 2147483648)
            v7 = (v8 >> ((v10) & 63))
            _ = v18
        }
    case "hello":
        v7 = (v7 + v8)
    case "a":
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
