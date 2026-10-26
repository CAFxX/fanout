package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int64 }
type GenS2 struct { F0 uint32; F1 string; F2 string }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 uint32 }
func genFn0(p0 int, p1 string, p2 string) int16 {
    var v4 int32 = 449
    v5 := []int{-68, 2147483647}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 531
}
func genFn1() bool {
    v6 := GenS1{F0: -434, F1: 2147483647, F2: -9223372036854775808}
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 string = ""
        _ = fi7
        _ = v8
    }
    if (-256 != -626) {
        var v9 float32 = 1e10
        v9 = 0.0
        _ = v9
    } else {
        var v10 float32 = float32(1.5)
        for fi11 := 0; fi11 < 11; fi11++ {
            v10 = ((3.14159 - (v10 - v10)) / v10)
            var v12 uint16 = 1
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    if false {
        var v13 int32 = -887
        _ = v13
    } else {
        var v14 uint = 65535
        _ = v14
    }
    var v15 int16 = -32768
    _ = v15
    _ = v6
    return (542 >= 2147483647)
}

func main() {
    v1 := []uint32{0, 2147483648, 2147483648}
    var v2 uint32 = 2
    var v3 string = "\n\t"
    v2 += (^(v2))
    v2 -= v2
    _ = len(v1)
    var v16 uint16 = 42
    v17 := GenS2{F0: (v2 ^ v2), F1: v3, F2: v3}
    for fi18 := 0; fi18 < 3; fi18++ {
        fmt.Println(v16)
        v19 := GenS3{F0: 904, F1: -1}
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 string = "\n\t"
            var v22 int64 = 0
            v1 = append(v1, 2147483648)
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    if (128 >= 0) {
        v2 = 714
        fmt.Println(v2, v2, v2)
        _ = len(v1)
    } else {
        v3 = v3
        v2 += 2147483647
    }
    if (!((v16 | (v16 >> ((207) & 63))) != 2)) {
        if len(v1) > 0 {
            v1[0] = v2
        }
    }
    var v23 uint8 = 2
    var v24 int = 32768
    for fi25 := 0; fi25 < 14; fi25++ {
        switch (v3 + v3) {
        case "foo bar":
            v24 -= v24
        case "a":
            if len(v1) > 0 {
                v1[0] = (v2 * (v2 + v2))
            }
        case "\n\t":
            v16 -= (^(v16))
        default:
            v26 := GenS2{F0: 2147483647, F1: ("" + "\n\t"), F2: "hello"}
            _ = v26
        }
        if (!true) {
            if ("" >= v3) {
                fmt.Println(fi25)
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 int8 = 1
                    v29 := GenS4{F0: 191, F1: v2}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            } else {
                v16 += (v16 >> ((879) & 63))
                if false {
                    var v30 int = v24
                    v2 -= v2
                    v31 := GenS5{F0: -1}
                    _ = v30
                    _ = v31
                }
            }
        } else {
            var v32 uint = 42
            v2 = v2
            fmt.Println(v16)
            _ = v32
        }
        if (-128 == 1) {
            if (-1 > 42) {
                for fi33 := 0; fi33 < 12; fi33++ {
                    v34 := GenS2{F0: v2, F1: v3, F2: v3}
                    var v35 float64 = (-2.25 - 1e10)
                    var v36 uint = 236
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                }
            }
            for fi37 := 0; fi37 < 3; fi37++ {
                v2 -= 2147483648
                _ = fi37
            }
        }
        _ = fi25
    }
    if true {
        v38 := GenS6{F0: 1, F1: (v2 % ((v2) | 1))}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v38
    } else {
        for fi39 := 0; fi39 < 2; fi39++ {
            _ = len(v1)
            var v40 uint = 2
            _ = fi39
            _ = v40
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
