package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint64; F1 string; F2 int8 }
func genFn0(p0 int16) int32 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint16 = 1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 317, F1: p0}
    for fi10 := 0; fi10 < 19; fi10++ {
        switch -820 {
        case 127:
            var v11 int32 = 25
            _ = v11
        }
        if (754 <= 746) {
            var v12 int8 = 127
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    _ = v9
    return 1
}
func genFn1(p0 uint16) uint32 {
    var v13 uint64 = 18446744073709551615
    v14 := GenS2{F0: 550, F1: -1}
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(fi15, p0)
        fmt.Println(fi15)
        v16 := []int{(fi15 ^ fi15), int(v13)}
        _ = fi15
        _ = v16
    }
    var v17 uint64 = (v13 & v13)
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := []int{fi18, (fi18 * (fi18 / ((1) | 1))), (fi18 - 1)}
        fmt.Println(v13)
        v17 += 42
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = v13
    _ = v14
    _ = v17
    return 42
}

func main() {
    var v1 uint8 = 2
    var v2 float64 = 0.0
    v3 := []int{2147483647, 32767, 42}
    v4 := []uint32{0}
    v5 := make(map[int]int)
    var v6 uint64 = 344
    fmt.Println(v1, v2)
    var v20 int64 = 1
    if true {
        for fi21 := 0; fi21 < 20; fi21++ {
            for fi22 := 0; fi22 < 2; fi22++ {
                v20 -= 367
                v2 -= 0.0
                var v23 uint16 = 65535
                fmt.Println(v20)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v20 += (v20 ^ 0)
                _ = fi22
                _ = v23
            }
            if (932 != 2) {
                var v24 string = "\n\t"
                _ = v24
            } else {
                if (((-2.25 * v2) < ((v2 * 1e10) / (-2.25 / v2))) || (1 <= 99)) {
                    var v25 bool = false
                    var v26 bool = v25
                    _ = v25
                    _ = v26
                }
                v2 -= (1.5 + v2)
                fmt.Println(v20)
            }
            v1 = (v1 * (v1 % (((v1 & (^(v1)))) | 1)))
            _ = fi21
        }
        if true {
            fmt.Println(v1, v1)
            v27 := GenS3{F0: 296, F1: (32767 >= 42)}
            v1 = (v1 - (v1 * (v1 % ((11) | 1))))
            _ = v27
        }
        var v28 bool = ((true && ((0 > -9223372036854775808) || (true && false))) || false)
        _ = v28
    }
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    if true {
        switch 127 {
        case 9223372036854775807:
            v29 := GenS1{F0: 1, F1: 42}
            _ = v29
        case 0:
            v30 := GenS4{F0: (v6 >> ((18446744073709551615) & 63)), F1: "\n\t", F2: -1}
            _ = v30
        case -178:
            if false {
                var v31 int16 = int16(v20)
                _ = v31
            }
        default:
            var v32 uint8 = v1
            _ = v32
        }
        if (!(-2.25 > v2)) {
            delete(v5, 4294967295)
            _ = len(v5)
        }
    } else {
        for fi33 := 0; fi33 < 16; fi33++ {
            v1 -= (v1 | (^(v1)))
            _ = fi33
        }
        var v34 int = 4294967295
        _ = v34
    }
    v1 += v1
    if false {
        _ = len(v5)
        if true {
            v6 -= (v6 / ((0) | 1))
        }
        v2 += v2
    } else {
        v20 = (v20 / (((v20 << ((478) & 63))) | 1))
        var v35 uint64 = (v6 / ((v6) | 1))
        _ = v35
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        _ = len(v5)
        _ = fi36
    }
    var v37 int32 = 1
    var v38 string = "hello"
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
