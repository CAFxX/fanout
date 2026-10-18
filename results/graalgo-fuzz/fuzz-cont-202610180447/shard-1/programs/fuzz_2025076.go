package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 bool }
type GenS3 struct { F0 string }
func genFn0(p0 uint8, p1 uint32) uint {
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := []int{fi4, fi4, fi4}
        _ = fi4
        _ = v5
    }
    if ((637 > 226) && (-9223372036854775808 >= 845)) {
        v6 := []int{32767, 42, -9223372036854775808}
        fmt.Println(p0, p1)
        v7 := GenS1{F0: 32767, F1: 1}
        _ = v6
        _ = v7
    } else {
        if (!(((1e10 / 1.5) >= (float32(0.0) / -2.25)) && false)) {
            fmt.Println(p1, p0)
            if (true && true) {
                fmt.Println(p0, p1)
            }
            v8 := GenS2{F0: -878, F1: -251, F2: false}
            _ = v8
        } else {
            p1 = (^(p1))
            var v9 int8 = 127
            p0 -= (p0 | (p0 / ((65) | 1)))
            _ = v9
        }
    }
    v10 := GenS3{F0: "hello"}
    _ = p0
    _ = p1
    _ = v10
    return uint(p0)
}

func main() {
    var v1 bool = (2 < 1)
    v2 := make(map[int]uint32)
    var v3 int = 2147483647
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            if v1 {
                var v13 string = (("" + "foo bar") + ("a" + "foo bar"))
                var v14 uint64 = 2
                _ = v13
                _ = v14
            }
            switch (v3 | fi12) {
            case 128:
                v15 := GenS3{F0: "hello"}
                _ = v15
            case 32768:
                v3 = (fi11 / ((fi11) | 1))
            default:
                var v16 uint = genFn0(0, 4294967295)
                _ = v16
            }
            v1 = v1
            _ = fi12
        }
        _ = fi11
    }
    v3 -= (v3 << ((708) & 63))
    var v17 float64 = ((0.0 - (1e10 / 0.1)) / 1.5)
    if v1 {
        v17 -= v17
        delete(v2, 499)
    }
    _ = len(v2)
    v3 = (v3 >> ((124) & 63))
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(v3, v17)
        var v19 float64 = v17
        v3 += (v3 - -2147483648)
        _ = fi18
        _ = v19
    }
    v1 = v1
    _ = len(v2)
    fmt.Println(v1, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
}
