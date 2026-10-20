package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int8; F1 bool; F2 int32 }
func genFn0(p0 uint, p1 uint) int16 {
    v6 := GenS1{F0: int(p1), F1: true}
    if true {
        p1 += 0
        fmt.Println(p0, p1)
    } else {
        v7 := GenS1{F0: -2147483648, F1: (2 == 42)}
        _ = v7
    }
    v8 := GenS1{F0: 81, F1: ((85 > 127) && ((3.14159 - (-2.25 + 3.14159)) >= (-2.25 + ((float64(((3.14159 / (0.0 - (1e10 / 3.14159))) / 1e10)) - (0.1 * 0.0)) + -2.25))))}
    p0 = (^(p1))
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 1000
}
func genFn1(p0 int8, p1 int16) bool {
    var v9 uint16 = 65535
    v10 := GenS1{F0: 0, F1: true}
    if ((0.0 - 3.14159) >= -2.25) {
        v9 += v9
    } else {
        var v11 int64 = 42
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS1{F0: (fi12 * (fi12 ^ (fi12 * (fi12 << ((255) & 63))))), F1: true}
            _ = fi12
            _ = v13
        }
        v14 := []int{32768, 315}
        _ = v11
        _ = v14
    }
    fmt.Println(p0, v9, v9)
    if false {
        v15 := []int{-440, int(v9), 907, -9223372036854775808}
        p0 = (p0 % ((p0) | 1))
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (2 != v9)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    var v3 int32 = -2147483648
    var v4 int8 = 1
    var v5 uint64 = 109
    v3 -= 2147483647
    if false {
        var v16 uint32 = 948
        switch int(v16) {
        case 614:
            if (169 >= (^(v5))) {
                v17 := GenS2{F0: (v5 & v5), F1: -1}
                _ = v17
            } else {
                var v18 int16 = 32767
                _ = v18
            }
        case 1:
            var v19 float32 = 0.1
            _ = v19
        }
        _ = v16
    }
    v4 = (v4 << ((0) & 63))
    if true {
        v20 := []int{}
        fmt.Println(v5, v4)
        v4 = (v4 % ((-128) | 1))
        _ = v20
    } else {
        fmt.Println(v5, v3, v5)
        if ((false && (255 != 42)) && (!true)) {
            v21 := GenS1{F0: 32768, F1: false}
            _ = v21
        }
    }
    v3 += (v3 - 2147483647)
    v4 -= v4
    v4 = v4
    var v22 int64 = -1
    if true {
        v22 -= -1
    } else {
        for fi23 := 0; fi23 < 18; fi23++ {
            v24 := v2[(fi23 / (((fi23 % ((32768) | 1))) | 1))]
            fmt.Println(v3, v3)
            var v25 string = "foo bar"
            _ = fi23
            _ = v24
            _ = v25
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            for fi27 := 0; fi27 < 5; fi27++ {
                v28 := GenS3{F0: (v4 + v4), F1: true, F2: 0}
                _ = fi27
                _ = v28
            }
            v29 := GenS1{F0: fi26, F1: false}
            v30 := GenS3{F0: v4, F1: true, F2: v3}
            _ = fi26
            _ = v29
            _ = v30
        }
    }
    v4 -= (v4 << ((42) & 63))
    delete(v1, 65535)
    v3 += v3
    v22 = v22
    v22 = 233
    fmt.Println(v5, v4, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
