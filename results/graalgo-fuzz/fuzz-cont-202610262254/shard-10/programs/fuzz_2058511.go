package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint32 }
func genFn0(p0 uint) int32 {
    for fi4 := 0; fi4 < 11; fi4++ {
        p0 -= 2
        for fi5 := 0; fi5 < 12; fi5++ {
            v6 := []int{(fi5 << ((p0) & 63)), fi4, (fi5 & fi4), (fi5 / ((fi5) | 1))}
            var v7 uint8 = 42
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    v8 := []int{32767, 916}
    p0 = (p0 / (((p0 / (((p0 * (^(p0)))) | 1))) | 1))
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 2147483647
}
func genFn1(p0 int64, p1 uint32) int {
    p0 = p0
    p1 += 1
    var v9 uint8 = 131
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(p0)
        v11 := GenS1{F0: 0}
        _ = fi10
        _ = v11
    }
    switch (("\n\t" + "foo bar") + ("a" + "foo bar")) {
    case "":
        v12 := []int{-9223372036854775808, 128, 42, 1}
        _ = v12
    case "a":
        if true {
            if true {
                var v13 int8 = -1
                _ = v13
            }
        } else {
            var v14 int64 = (p0 / ((p0) | 1))
            _ = v14
        }
    default:
        v9 = (v9 | (v9 - v9))
    }
    _ = p0
    _ = p1
    _ = v9
    return 4294967295
}

func main() {
    var v1 int16 = 1000
    var v2 int = 313
    v3 := make(map[int]string)
    var v15 float64 = float64(1.5)
    fmt.Println(v15)
    fmt.Println(v1)
    if (-9223372036854775808 <= 1) {
        if (false && false) {
            fmt.Println(v15)
        } else {
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 int8 = -87
                var v18 uint32 = 2147483648
                v19 := GenS2{F0: 1}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        }
    } else {
        v20 := GenS1{F0: 1}
        _ = v20
    }
    v1 = v1
    v15 += v15
    var v21 uint32 = 266
    switch (v2 & (^(v2))) {
    case -449:
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := []int{}
            _ = fi22
            _ = v23
        }
    case 9223372036854775807:
        v1 -= v1
    }
    var v24 uint = 834
    if (32767 > 32767) {
        if true {
            for fi25 := 0; fi25 < 9; fi25++ {
                if (true || ((v21 == v21) && (87 <= -514))) {
                    var v26 uint64 = 739
                    _ = v26
                } else {
                    switch (v2 | genFn1(0, (v21 - v21))) {
                    case -2147483648:
                        switch v2 {
                        case -221:
                            var v27 int = fi25
                            _ = v27
                        }
                    case 0:
                        var v28 int = 127
                        _ = v28
                    }
                    v1 = (v1 * int16(v21))
                }
                v29 := GenS2{F0: 957}
                _ = fi25
                _ = v29
            }
            var v30 int16 = (v1 | v1)
            fmt.Println(v2, v30, v1)
            _ = v30
        }
    }
    v31 := GenS2{F0: 42}
    var v32 int8 = 42
    fmt.Println(v32)
    v33 := GenS2{F0: 18446744073709551615}
    for fi34 := 0; fi34 < 16; fi34++ {
        var v35 int64 = 1
        v32 = v32
        v32 += (v32 & (v32 / (((v32 ^ (v32 % ((v32) | 1)))) | 1)))
        _ = fi34
        _ = v35
    }
    if (-1 < -32768) {
        v36 := GenS3{F0: v2, F1: 0}
        v37 := GenS3{F0: v2, F1: v21}
        _ = v36
        _ = v37
    } else {
        v15 += float64(-2.25)
        delete(v3, -1)
    }
    v24 += (v24 - 18446744073709551615)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
