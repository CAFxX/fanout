package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint16; F2 bool }
type GenS3 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[int]int)
    var v2 string = "foo bar"
    var v3 int64 = 42
    var v4 uint16 = 920
    if true {
        var v5 int16 = -594
        v6 := GenS1{F0: 4294967295}
        if true {
            v3 -= -1
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := []int{fi7, fi7}
                var v9 float32 = 0.1
                _ = fi7
                _ = v8
                _ = v9
            }
            switch v2 {
            case "\n\t":
                for fi10 := 0; fi10 < 2; fi10++ {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        v4 += (v4 ^ 255)
                        fmt.Println(v5)
                        v12 := GenS2{F0: int8(fi10), F1: v4, F2: true}
                        _ = fi11
                        _ = v12
                    }
                    var v13 int8 = 69
                    fmt.Println(v13, v5)
                    _ = fi10
                    _ = v13
                }
            default:
                v14 := GenS3{F0: 255, F1: v2}
                _ = v14
            }
        }
        _ = v5
        _ = v6
    } else {
        v3 = (v3 * v3)
    }
    if (true && (289 >= 2)) {
        if false {
            v15 := []int{989, 1, 428}
            v3 = (v3 % (((^(v3))) | 1))
            _ = v15
        } else {
            if true {
                v3 = v3
            }
        }
        var v16 uint16 = (v4 << ((0) & 63))
        _ = v16
    } else {
        v2 = ""
        var v17 uint32 = 1
        _ = v17
    }
    switch -131 {
    case 487:
        fmt.Println(v3, v3)
    case 128:
        v18 := []int{2147483647, -9223372036854775808, 4294967295}
        _ = v18
    case -175:
        v4 -= (v4 | (v4 | 255))
    }
    v3 = -219
    fmt.Println(v2, v3)
    fmt.Println(v4, v2, v3)
    var v19 float64 = 3.14159
    switch 4294967295 {
    case 704:
        fmt.Println(v19, v3)
    case 2147483647:
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 uint = 2
            v2 = ("\n\t" + v2)
            var v22 uint = 1
            _ = fi20
            _ = v21
            _ = v22
        }
    case 4294967295:
        if (true || false) {
            var v23 bool = (728 > 551)
            _ = v23
        } else {
            v4 += (v4 - v4)
            fmt.Println(v2)
            if ((-1 > -1) && true) {
                var v24 uint8 = 134
                var v25 int32 = 550
                _ = v24
                _ = v25
            } else {
                v19 -= v19
                v4 += 65535
                if true {
                    var v26 bool = (!(0 <= 127))
                    fmt.Println(v3)
                    v27 := []int{2147483647, 127, 0}
                    _ = v26
                    _ = v27
                } else {
                    var v28 uint32 = 2147483647
                    _ = v28
                }
            }
        }
    default:
        v3 += v3
    }
    if (true || false) {
        v4 -= (v4 & v4)
        if false {
            delete(v1, -9223372036854775808)
            v1[1] = 4294967295
            var v29 int8 = 0
            _ = v29
        }
        v2 = v2
    } else {
        if true {
            for fi30 := 0; fi30 < 10; fi30++ {
                var v31 uint8 = 2
                _ = fi30
                _ = v31
            }
            var v32 uint16 = (v4 * 2)
            _ = v32
        }
    }
    var v33 uint64 = 1
    fmt.Println(v33, v3)
    v4 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
