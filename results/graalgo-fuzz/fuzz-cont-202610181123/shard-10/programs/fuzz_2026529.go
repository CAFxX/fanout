package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 string, p1 uint) int16 {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: (fi4 | -9223372036854775808)}
        _ = fi4
        _ = v5
    }
    switch "" {
    case "\n\t":
        for fi6 := 0; fi6 < 10; fi6++ {
            if (0 == 127) {
                p1 -= (p1 >> ((65535) & 63))
                var v7 uint = (^(p1))
                _ = v7
            } else {
                var v8 uint64 = 134
                _ = v8
            }
            switch (fi6 & 1) {
            case 42:
                fmt.Println(fi6)
            }
            var v9 uint64 = 4294967295
            _ = fi6
            _ = v9
        }
    default:
        v10 := GenS1{F0: 1}
        _ = v10
    }
    v11 := []int{42}
    _ = p0
    _ = p1
    _ = v11
    return -1
}
func genFn1(p0 int8, p1 string) uint {
    fmt.Println(p1, p1)
    fmt.Println(p0)
    if false {
        var v12 uint32 = 158
        _ = v12
    }
    if true {
        v13 := []int{128}
        if false {
            v14 := GenS2{F0: 2, F1: 4294967295}
            _ = v14
        } else {
            var v15 uint32 = 0
            _ = v15
        }
        _ = v13
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 float64 = -2.25
    v2 := []int{2147483647, 32768}
    v3 := make(map[int]int)
    var v16 uint16 = 0
    var v17 string = "a"
    v1 = float64(3.14159)
    if (!false) {
        var v18 uint = 0
        var v19 uint64 = 153
        _ = v18
        _ = v19
    } else {
        v17 = "foo bar"
    }
    var v20 int = 128
    switch (v20 - 42) {
    case 162:
        v16 -= v16
    case 32767:
        v17 = v17
    default:
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS1{F0: v20}
            v1 += (v1 - float64(1e10))
            fmt.Println(fi21)
            _ = fi21
            _ = v22
        }
    }
    v23 := GenS1{F0: v20}
    var v24 float32 = 1.5
    delete(v3, v20)
    v1 -= (v1 / 3.14159)
    if (2 <= 255) {
        switch v20 {
        case -5:
            var v25 int16 = 32767
            _ = v25
        case 9223372036854775807:
            switch (v20 * v20) {
            case 376:
                if ((117 > 172) && true) {
                    v26 := GenS2{F0: 18446744073709551615, F1: 65535}
                    _ = v26
                }
            case -1:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                v20 -= -9223372036854775808
            }
        case 32767:
            if false {
                if (!false) {
                    var v27 float64 = 0.1
                    var v28 uint64 = 244
                    _ = v27
                    _ = v28
                }
            } else {
                v29 := GenS1{F0: v20}
                v16 = 0
                for fi30 := 0; fi30 < 17; fi30++ {
                    v16 = v16
                    var v31 uint32 = 42
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
        }
    }
    v17 = (v17 + v17)
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
}
