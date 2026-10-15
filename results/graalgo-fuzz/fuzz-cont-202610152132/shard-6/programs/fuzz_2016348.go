package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int8; F1 uint32; F2 int16 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 int) int8 {
    switch (p0 & (p0 & p0)) {
    case -363:
        switch (-(p0)) {
        case 128:
            switch "foo bar" {
            case "foo bar":
                fmt.Println(p0)
            default:
                v5 := GenS1{F0: 18446744073709551615, F1: 1}
                _ = v5
            }
        case 127:
            p0 = -98
        }
    }
    fmt.Println(p0)
    v6 := GenS2{F0: -1, F1: 42, F2: 42}
    var v7 uint16 = 2
    switch p0 {
    case 4294967295:
        if true {
            v8 := GenS2{F0: 0, F1: 781, F2: -766}
            _ = v8
        } else {
            if false {
                p0 -= (-(p0))
            } else {
                var v9 float32 = 0.0
                _ = v9
            }
            var v10 uint16 = (v7 & (v7 << ((255) & 63)))
            v11 := GenS3{F0: "\n\t", F1: 2147483648}
            _ = v10
            _ = v11
        }
    case 128:
        if (188 != 0) {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int32 = -373
                _ = fi12
                _ = v13
            }
        }
    default:
        var v14 uint8 = 92
        _ = v14
    }
    _ = p0
    _ = v6
    _ = v7
    return -128
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{-1, 333, -9223372036854775808, int(v1)}
    var v3 int32 = 1
    v4 := make(map[int]uint32)
    v15 := GenS1{F0: 0, F1: -118}
    var v16 int16 = -867
    var v17 bool = true
    v3 = 42
    v17 = (!v17)
    var v18 int64 = (v1 - v1)
    v3 -= v3
    for fi19 := 0; fi19 < 11; fi19++ {
        if (v17 && (3.14159 <= (-2.25 + 1.5))) {
            if (!(1e10 == (1.5 * 1.5))) {
                if (127 != 255) {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        delete(v4, (fi20 % ((fi19) | 1)))
                        _ = fi20
                    }
                    v3 = v3
                } else {
                    if v17 {
                        delete(v4, -2147483648)
                        for fi21 := 0; fi21 < 9; fi21++ {
                            _ = len(v4)
                            _ = fi21
                        }
                    }
                    var v22 uint8 = 1
                    _ = v22
                }
                v23 := GenS1{F0: 18446744073709551615, F1: v18}
                _ = v23
            }
            v18 = (v18 | v1)
        } else {
            v24 := GenS1{F0: 4294967295, F1: 281}
            for fi25 := 0; fi25 < 5; fi25++ {
                v2 = append(v2, fi25)
                v18 -= 0
                v16 += (-(v16))
                _ = fi25
            }
            _ = v24
        }
        var v26 uint64 = 163
        for fi27 := 0; fi27 < 12; fi27++ {
            v18 = (v18 - (-(v1)))
            var v28 uint = 4294967295
            _ = fi27
            _ = v28
        }
        _ = fi19
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
