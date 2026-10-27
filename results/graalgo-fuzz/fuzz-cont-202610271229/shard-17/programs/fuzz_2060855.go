package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int64; F1 uint32; F2 int8 }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0(p0 uint, p1 string) bool {
    v4 := GenS1{F0: 1, F1: 2, F2: 0}
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(fi5, p0, p1)
        p0 = p0
        _ = fi5
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        if (true || false) {
            fmt.Println(p1, p0, fi6)
            fmt.Println(p0, p1)
        } else {
            v7 := GenS1{F0: 626, F1: 65535, F2: 42}
            if false {
                var v8 int8 = -128
                var v9 int = fi6
                var v10 uint64 = 0
                _ = v10
                _ = v8
                _ = v9
            } else {
                v11 := []int{(fi6 | (fi6 & (fi6 + -2147483648))), fi6, fi6, -909}
                if ((3.14159 <= ((1e10 / (-2.25 / ((0.1 + ((0.1 * 1.5) / (-2.25 - 1.5))) - 3.14159))) * float32(((0.0 / -2.25) + (1e10 + (3.14159 / 1e10)))))) && (255 > 32767)) {
                    v12 := GenS1{F0: -915, F1: 4294967295, F2: -9223372036854775808}
                    _ = v12
                }
                _ = v11
            }
            _ = v7
        }
        p0 = 42
        if false {
            v13 := GenS2{F0: 32767, F1: -311}
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := []int{fi14, fi6, (fi14 / ((fi14) | 1)), fi6}
                p1 = p1
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    return true
}
func genFn1() uint64 {
    var v16 int64 = 42
    v17 := []int{-521, -9223372036854775808}
    v16 -= (v16 | v16)
    _ = v16
    _ = v17
    return 0
}

func main() {
    var v1 int64 = 968
    v2 := []int{}
    var v3 int8 = 1
    var v18 uint8 = 69
    if (v18 <= 127) {
        if ((!true) || (4294967295 == 926)) {
            if false {
                fmt.Println(v1, v18)
            }
            if false {
                fmt.Println(v18, v1, v1)
                v1 += (v1 % ((-9223372036854775808) | 1))
                var v19 float64 = 1e10
                _ = v19
            }
        } else {
            if ((0 > 1000) && ((65 >= 128) && true)) {
                var v20 uint64 = 42
                _ = v20
            }
            var v21 uint = 65535
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 int32 = -1
        v3 -= 23
        _ = fi22
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = 715
    }
    v3 += int8(v18)
    v24 := GenS3{F0: v1, F1: 0, F2: (v3 % ((42) | 1))}
    for fi25 := 0; fi25 < 4; fi25++ {
        v18 = (v18 * (v18 | (v18 ^ v18)))
        _ = fi25
    }
    switch 0 {
    case 973:
        fmt.Println(v18, v18)
    case 1:
        switch "hello" {
        case "a":
            var v26 float32 = 0.0
            _ = v26
        }
    }
    switch "a" {
    case "hello":
        if true {
            fmt.Println(v1, v3, v18)
        } else {
            v27 := GenS4{F0: 42, F1: 933}
            if len(v2) > 0 {
                v2[0] = -1
            }
            switch 0 {
            case -978:
                for fi28 := 0; fi28 < 1; fi28++ {
                    switch (fi28 & 2147483647) {
                    case -144:
                        var v29 int = 32768
                        _ = v29
                    default:
                        v1 = 42
                    }
                    var v30 float64 = (float64(0.0) / 1.5)
                    v30 -= ((v30 * float64((1e10 * float32(0.1)))) + (v30 / v30))
                    _ = fi28
                    _ = v30
                }
            default:
                var v31 string = "hello"
                _ = v31
            }
            _ = v27
        }
    default:
        v3 = v3
    }
    var v32 int8 = -1
    var v33 int64 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v33
}
