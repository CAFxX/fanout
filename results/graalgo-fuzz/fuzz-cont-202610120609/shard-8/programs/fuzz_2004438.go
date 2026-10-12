package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int16 = -294
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 int64 = 0
    v5 := []int{342, 127, -9223372036854775808, 1}
    v4 = v4
    var v6 bool = (127 > 150)
    v6 = (((4294967295 >= 62) || v6) && ((v4 & v4) < (v4 ^ -1)))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 int64 = (v4 - (v4 % (((v4 >> ((255) & 63))) | 1)))
    var v8 int = -248
    v9 := GenS1{F0: 42, F1: 625}
    var v10 float64 = 1.5
    v11 := GenS1{F0: 65535, F1: 42}
    for fi12 := 0; fi12 < 13; fi12++ {
        v8 = (v8 << ((4294967295) & 63))
        var v13 uint64 = 621
        _ = fi12
        _ = v13
    }
    if false {
        if false {
            v10 += float64(1.5)
            v7 = v7
        } else {
            switch (v8 / ((1) | 1)) {
            case 128:
                switch "" {
                case "\n\t":
                    v14 := GenS1{F0: 1, F1: 2}
                    _ = v14
                case "foo bar":
                    if (v6 && (!(true || v6))) {
                        var v15 uint8 = 128
                        if v6 {
                            var v16 bool = (!false)
                            _ = v16
                        } else {
                            if (!true) {
                                v4 = (v7 >> ((2) & 63))
                                fmt.Println(v1, v7)
                            }
                        }
                        _ = v15
                    }
                case "":
                    var v17 float32 = 1e10
                    _ = v17
                }
            case -216:
                if (0 < 1) {
                    fmt.Println(v10, v6)
                    switch (v8 ^ 1) {
                    case -531:
                        v18 := GenS2{F0: (4294967295 < 1)}
                        _ = v18
                    }
                } else {
                    v19 := v3[779]
                    if true {
                        v1 = v1
                        v20 := v3[269]
                        _ = v20
                    } else {
                        v21 := GenS1{F0: 65535, F1: 1}
                        switch 0 {
                        case 925:
                            v19 += (v8 / ((v8) | 1))
                        }
                        _ = v21
                    }
                    _ = v19
                }
            case 2147483647:
                var v22 uint = 2
                _ = v22
            }
            for fi23 := 0; fi23 < 17; fi23++ {
                if false {
                    _ = len(v5)
                    fmt.Println(v8)
                }
                _ = fi23
            }
            var v24 string = "\n\t"
            _ = v24
        }
        var v25 int8 = 71
        _ = v25
    }
    fmt.Println(v8, v8, v10)
    if v6 {
        fmt.Println(v1, v1, v8)
        _ = len(v5)
        if (true && ((!true) && v6)) {
            switch (^(v8)) {
            case -607:
                if (!(165 != 68)) {
                    var v26 int16 = v1
                    switch "\n\t" {
                    case "":
                        var v27 int16 = (v26 ^ v1)
                        _ = v27
                    case "hello":
                        v28 := GenS1{F0: 486, F1: 1}
                        _ = v28
                    default:
                        fmt.Println(v7, v10)
                    }
                    fmt.Println(v1)
                    _ = v26
                } else {
                    v7 = -1
                }
            default:
                if v6 {
                    v4 += 1
                    fmt.Println(v8)
                    fmt.Println(v7, v10)
                }
            }
            v10 = float64(1.5)
        }
    } else {
        v10 -= float64(1e10)
        var v29 uint32 = 65535
        _ = v29
    }
    if ((^(v8)) < 0) {
        var v30 uint8 = 0
        v1 += v1
        switch int(v4) {
        case -564:
            if ((v8 < 0) || v6) {
                switch (-(v8)) {
                case 128:
                    v31 := GenS1{F0: 671, F1: 18446744073709551615}
                    _ = v31
                case -1:
                    var v32 int32 = 388
                    _ = v32
                case -2147483648:
                    v8 += v8
                default:
                    v6 = v6
                }
                if (!false) {
                    fmt.Println(v10, v7, v30)
                    var v33 int16 = (v1 / ((v1) | 1))
                    _ = v33
                } else {
                    for fi34 := 0; fi34 < 13; fi34++ {
                        v35 := GenS1{F0: 2, F1: 18446744073709551615}
                        _ = fi34
                        _ = v35
                    }
                }
            }
        case 0:
            fmt.Println(v6, v8, v8)
        default:
            v36 := v2[-9223372036854775808]
            _ = v36
        }
        _ = v30
    }
    delete(v3, 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
