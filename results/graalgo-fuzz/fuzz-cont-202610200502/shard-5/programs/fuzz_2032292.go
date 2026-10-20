package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint, p1 int64, p2 string) bool {
    if true {
        p1 = (p1 & 9223372036854775807)
    }
    switch 42 {
    case 128:
        p2 = ("\n\t" + p2)
    case 32768:
        for fi4 := 0; fi4 < 6; fi4++ {
            for fi5 := 0; fi5 < 16; fi5++ {
                v6 := GenS1{F0: 2}
                v7 := GenS2{F0: 42, F1: 2147483647, F2: -1}
                p2 = p2
                _ = fi5
                _ = v6
                _ = v7
            }
            p2 = (("a" + (p2 + "a")) + p2)
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := []int{fi8, (fi4 << ((1) & 63)), (fi4 % ((-9223372036854775808) | 1))}
                var v10 int8 = 0
                _ = len(v9)
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi4
        }
    default:
        p1 -= (p1 >> ((p0) & 63))
    }
    v11 := GenS3{F0: 508}
    v12 := []int{-2147483648, -254, -2147483648}
    v12 = append(v12, -2147483648)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return (2 > uint32(p0))
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 128
    var v3 int = v2
    if false {
        var v13 int64 = 0
        v13 += v13
        _ = v13
    } else {
        fmt.Println(v1)
        var v14 uint64 = 231
        var v15 uint = 972
        _ = v14
        _ = v15
    }
    v3 = 9223372036854775807
    var v16 uint64 = 977
    if (("hello" < v1) || false) {
        if true {
            var v17 uint8 = 246
            v18 := []int{(v3 / (((v3 | v3)) | 1)), int(v16), (v3 * (^(v3)))}
            v17 -= (v17 + (^(v17)))
            _ = v17
            _ = v18
        }
    } else {
        v2 = (v2 / ((v3) | 1))
        for fi19 := 0; fi19 < 16; fi19++ {
            if (3.14159 != ((float64((0.0 - 3.14159)) + 1.5) - (3.14159 * 0.0))) {
                fmt.Println(fi19, v3)
                fmt.Println(v16)
                var v20 int64 = 9223372036854775807
                _ = v20
            }
            _ = fi19
        }
        v21 := []int{v3, (-(v2)), (v2 << ((18446744073709551615) & 63))}
        _ = v21
    }
    v1 = (v1 + v1)
    v22 := []int{}
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 int16 = -192
        v25 := GenS1{F0: v16}
        _ = fi23
        _ = v24
        _ = v25
    }
    v16 = 90
    var v26 float64 = (1e10 / float64(3.14159))
    if true {
        v26 = (v26 / v26)
        fmt.Println(v1)
    }
    switch (("" + v1) + v1) {
    case "":
        if true {
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        } else {
            var v27 int32 = -2147483648
            v28 := GenS4{F0: 59}
            _ = v27
            _ = v28
        }
    }
    if (-2147483648 < 2147483647) {
        _ = len(v22)
        v1 = "foo bar"
    } else {
        switch (v2 * 458) {
        case -456:
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        default:
            v1 = ((("a" + v1) + v1) + (v1 + ("foo bar" + ((v1 + v1) + (v1 + "hello")))))
        }
        fmt.Println(v26, v26, v1)
    }
    if ((942 != 0) || (0 >= 153)) {
        v26 -= (v26 + ((0.1 + v26) * 0.1))
        var v29 int64 = 9223372036854775807
        _ = len(v22)
        _ = v29
    }
    var v30 float64 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v30
}
