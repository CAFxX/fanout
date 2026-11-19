package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []uint32{65535}
    v2 := []int{-100, 4294967295, -2147483648}
    var v3 string = "hello"
    v4 := make(map[int]bool)
    v5 := make(map[string]uint32)
    var v6 float32 = ((0.1 - ((-2.25 / 1e10) * -2.25)) - -2.25)
    if (!true) {
        if (((v3 + "foo bar") + v3) != ((v3 + "") + "a")) {
            var v7 string = (v3 + v3)
            _ = v7
        }
    } else {
        var v8 uint64 = 0
        switch 4294967295 {
        case 855:
            var v9 uint32 = 1
            _ = v9
        case -421:
            v6 = (v6 + v6)
        case -736:
            switch 32768 {
            case 4294967295:
                fmt.Println(v6, v8, v8)
            case -2147483648:
                v2 = append(v2, 42)
            }
        }
        v10 := GenS1{F0: -45, F1: 1}
        _ = v10
        _ = v8
    }
    if true {
        if (2147483647 >= -659) {
            v5[v3] = 65535
        }
        var v11 int16 = 0
        _ = v11
    }
    var v12 bool = (("\n\t" + v3) <= v3)
    v2 = append(v2, 9223372036854775807)
    var v13 float32 = 1.5
    fmt.Println(v3)
    var v14 float64 = -2.25
    if (!(v6 > (0.0 / v13))) {
        _ = len(v1)
        if v12 {
            v14 = (v14 * (v14 / v14))
        } else {
            for fi15 := 0; fi15 < 14; fi15++ {
                if v12 {
                    var v16 uint16 = 255
                    _ = v16
                }
                _ = fi15
            }
            var v17 float32 = 1.5
            for fi18 := 0; fi18 < 10; fi18++ {
                if v12 {
                    var v19 float64 = float64(v17)
                    _ = v19
                }
                switch (fi18 - (fi18 >> ((11) & 63))) {
                case 305:
                    v17 -= (float32(v14) - v6)
                case -90:
                    var v20 float64 = 1e10
                    _ = v20
                default:
                    if (((fi18 * 9223372036854775807) <= 32767) && v12) {
                        if v12 {
                            v21 := GenS2{F0: 195, F1: -1}
                            _ = v21
                        } else {
                            for fi22 := 0; fi22 < 20; fi22++ {
                                var v23 int64 = 132
                                v24 := v5["foo bar"]
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            fmt.Println(fi18, v13, v13)
                        }
                    } else {
                        v25 := GenS1{F0: 1000, F1: 0}
                        fmt.Println(v13)
                        for fi26 := 0; fi26 < 17; fi26++ {
                            var v27 uint16 = 65535
                            _ = fi26
                            _ = v27
                        }
                        _ = v25
                    }
                }
                _ = fi18
            }
            _ = v17
        }
        fmt.Println(v3, v14, v14)
    }
    if v12 {
        var v28 uint16 = 0
        _ = v28
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        switch v3 {
        case "hello":
            if v12 {
                var v30 int16 = -1
                _ = v30
            }
        case "\n\t":
            fmt.Println(v6)
        }
        var v31 int8 = 0
        _ = fi29
        _ = v31
    }
    var v32 int64 = -9223372036854775808
    if (v13 < v13) {
        var v33 int8 = 1
        _ = v33
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        if v12 {
            var v35 int = fi34
            v13 -= 0.1
            _ = v35
        }
        fmt.Println(v32)
        if (1e10 == v14) {
            var v36 uint64 = 1
            _ = v36
        } else {
            var v37 int8 = 127
            v38 := GenS3{F0: -298}
            _ = v37
            _ = v38
        }
        _ = fi34
    }
    var v39 int = -394
    for fi40 := 0; fi40 < 4; fi40++ {
        fmt.Println(v39)
        _ = fi40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
