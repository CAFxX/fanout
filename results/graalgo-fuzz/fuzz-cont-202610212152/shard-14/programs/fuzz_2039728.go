package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int }
func genFn0(p0 uint32, p1 string) uint8 {
    var v5 uint32 = (p0 * (p0 ^ (p0 >> ((uint(p0)) & 63))))
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(p1, v5, p1)
        v7 := []int{0}
        _ = fi6
        _ = v7
    }
    p1 = (p1 + p1)
    if (65535 < 32767) {
        fmt.Println(p0)
    }
    var v8 int8 = -1
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 128
}
func genFn1(p0 int16) bool {
    var v9 int64 = -1
    for fi10 := 0; fi10 < 17; fi10++ {
        if ((v9 & (v9 - (v9 / ((v9) | 1)))) != (v9 % ((v9) | 1))) {
            p0 -= (p0 << ((42) & 63))
            v11 := []int{-1, fi10, (-(fi10))}
            p0 += (p0 * (p0 - 0))
            _ = v11
        }
        _ = fi10
    }
    switch -149 {
    case 2147483647:
        var v12 uint16 = 2
        _ = v12
    }
    v9 = 42
    _ = p0
    _ = v9
    return (42 < 255)
}

func main() {
    var v1 uint32 = 2
    var v2 int8 = -96
    var v3 uint = 244
    var v4 uint32 = (v1 * 1)
    v13 := []int{}
    v13 = append(v13, 32767)
    switch 311 {
    case 128:
        v13 = append(v13, 0)
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            v1 += (^(v4))
            if (632 == 4294967295) {
                var v16 bool = ((true && ((v3 - v3) < 18446744073709551615)) || (!(false && (!(0 < 0)))))
                var v17 int64 = 9223372036854775807
                _ = v16
                _ = v17
            }
            if (632 >= 42) {
                var v18 int32 = 1
                v4 = v4
                _ = v18
            }
            var v19 int64 = -9223372036854775808
            _ = len(v13)
            v20 := GenS1{F0: (^(v3)), F1: -1}
            _ = fi15
            _ = v19
            _ = v20
        }
        var v21 int64 = -9223372036854775808
        _ = fi14
        _ = v21
    }
    var v22 int32 = -1
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := GenS1{F0: v3, F1: 2147483647}
        if ("foo bar" <= "") {
            v1 -= (v1 | 0)
            fmt.Println(v4, v3)
        } else {
            for fi25 := 0; fi25 < 11; fi25++ {
                v2 = (v2 % ((v2) | 1))
                _ = fi25
            }
        }
        _ = fi23
        _ = v24
    }
    if (1e10 < 0.0) {
        var v26 int16 = -1
        _ = v26
    }
    var v27 string = ("" + ("" + ("" + "hello")))
    var v28 bool = (true && (255 < 12))
    switch 4294967295 {
    case 9223372036854775807:
        v13 = append(v13, -2147483648)
    }
    var v29 int = 1
    var v30 int8 = v2
    switch ((v27 + v27) + v27) {
    case "\n\t":
        fmt.Println(v30, v4, v2)
    }
    if v28 {
        v31 := GenS2{F0: (v1 - v1), F1: 930, F2: 32768}
        _ = v31
    }
    v22 += v22
    for fi32 := 0; fi32 < 10; fi32++ {
        if v28 {
            v33 := GenS1{F0: 2, F1: (v22 * -795)}
            if (v28 || ((1 >= (v3 >> ((255) & 63))) && v28)) {
                var v34 int64 = -9223372036854775808
                var v35 uint32 = (v4 * 0)
                var v36 uint32 = 65535
                _ = v34
                _ = v35
                _ = v36
            }
            for fi37 := 0; fi37 < 3; fi37++ {
                if (!v28) {
                    var v38 int = fi32
                    if (!false) {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        var v39 int32 = (v22 << ((494) & 63))
                        _ = v39
                    }
                    _ = len(v13)
                    v30 = 42
                    _ = v38
                }
                _ = fi37
            }
            _ = v33
        }
        for fi40 := 0; fi40 < 14; fi40++ {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi40
        }
        _ = fi32
    }
    v22 += (v22 - 42)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
