package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int64; F1 uint8; F2 int }

func main() {
    var v1 int64 = 0
    v2 := []int{1, 1, 42}
    var v3 uint = 4294967295
    var v4 int8 = -102
    for fi5 := 0; fi5 < 3; fi5++ {
        fmt.Println(fi5)
        v1 -= v1
        _ = fi5
    }
    if (!true) {
        v1 += v1
    } else {
        v1 = 0
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v6 int = 1
    fmt.Println(v3, v3)
    var v7 float32 = float32(-2.25)
    v8 := GenS1{F0: 0, F1: 0}
    switch (("hello" + "") + "foo bar") {
    case "\n\t":
        switch "\n\t" {
        case "":
            if true {
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            }
        case "\n\t":
            switch (v6 | (^(v6))) {
            case 1:
                switch "hello" {
                case "":
                    var v9 int16 = 42
                    _ = v9
                default:
                    if (175 <= 0) {
                        switch "hello" {
                        case "a":
                            var v10 uint8 = 0
                            _ = v10
                        case "":
                            _ = len(v2)
                        default:
                            if (true && false) {
                                v6 += v6
                                var v11 int16 = -1
                                _ = v11
                            } else {
                                var v12 float32 = v7
                                var v13 int32 = -2147483648
                                if false {
                                    if true {
                                        v6 -= v6
                                        var v14 uint8 = 2
                                        _ = v14
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = v6
                                    }
                                    if true {
                                        if ((366 != 2147483647) || (1e10 != 0.0)) {
                                            switch 128 {
                                            case 32768:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                var v15 uint = v3
                                                _ = v15
                                            }
                                            v16 := GenS1{F0: v13, F1: 127}
                                            var v17 uint32 = 1
                                            _ = v16
                                            _ = v17
                                        } else {
                                            var v18 int8 = (v4 << (((^(v3))) & 63))
                                            var v19 int64 = v1
                                            v1 = int64(v13)
                                            _ = v18
                                            _ = v19
                                        }
                                    }
                                } else {
                                    var v20 int = v6
                                    if (!true) {
                                        if (true && (2147483647 == 4294967295)) {
                                            var v21 uint8 = 127
                                            for fi22 := 0; fi22 < 18; fi22++ {
                                                v23 := GenS1{F0: (v13 - 1), F1: v21}
                                                _ = fi22
                                                _ = v23
                                            }
                                            v24 := GenS2{F0: v1}
                                            _ = v21
                                            _ = v24
                                        }
                                    }
                                    v3 -= (v3 / ((v3) | 1))
                                    _ = v20
                                }
                                _ = v12
                                _ = v13
                            }
                        }
                        fmt.Println(v6, v1)
                        switch (v6 * (^(v6))) {
                        case 543:
                            var v25 int = (v6 - v6)
                            _ = v25
                        case -1:
                            var v26 uint = 153
                            _ = v26
                        case -304:
                            v1 += v1
                        default:
                            v27 := GenS1{F0: 1, F1: 42}
                            _ = v27
                        }
                    } else {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            for fi29 := 0; fi29 < 3; fi29++ {
                                if true {
                                    v6 = (fi28 & -2147483648)
                                } else {
                                    if false {
                                        var v30 uint = v3
                                        _ = v30
                                    } else {
                                        if false {
                                            var v31 uint64 = 42
                                            v32 := GenS3{F0: v3}
                                            _ = v31
                                            _ = v32
                                        } else {
                                            _ = len(v2)
                                            fmt.Println(v4, v1)
                                        }
                                        var v33 bool = (false && false)
                                        fmt.Println(fi29, fi28)
                                        v2 = append(v2, fi28)
                                        v34 := GenS2{F0: (v1 * v1)}
                                        _ = v33
                                        _ = v34
                                    }
                                    var v35 int = (fi28 + (fi29 / (((fi29 ^ fi29)) | 1)))
                                    if len(v2) > 0 {
                                        v2[0] = fi28
                                    }
                                    _ = v35
                                }
                                _ = fi29
                            }
                            for fi36 := 0; fi36 < 20; fi36++ {
                                v37 := GenS2{F0: v1}
                                _ = fi36
                                _ = v37
                            }
                            for fi38 := 0; fi38 < 17; fi38++ {
                                v2 = append(v2, -9223372036854775808)
                                v3 = 42
                                _ = fi38
                            }
                            _ = fi28
                        }
                        if len(v2) > 0 {
                            v2[0] = 128
                        }
                        v6 = 4294967295
                    }
                }
            case 32768:
                v3 += v3
            case 127:
                for fi39 := 0; fi39 < 2; fi39++ {
                    v40 := GenS3{F0: v3}
                    var v41 int = fi39
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            }
        default:
            v2 = append(v2, v6)
        }
    default:
        v7 = v7
    }
    fmt.Println(v7)
    v42 := GenS4{F0: 100, F1: v4}
    var v43 string = "\n\t"
    v44 := GenS5{F0: v1, F1: 0, F2: v6}
    var v45 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v6
    _ = v7
    _ = v8
}
