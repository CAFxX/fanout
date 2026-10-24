package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int; F2 int }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    if (("" <= ("\n\t" + "a")) || (!(18446744073709551615 == 18446744073709551615))) {
        v4 := GenS1{F0: 0, F1: 1, F2: 245}
        for fi5 := 0; fi5 < 14; fi5++ {
            v6 := []int{-89}
            if false {
                v7 := GenS1{F0: -9223372036854775808, F1: 0, F2: 128}
                _ = v7
            }
            _ = fi5
            _ = v6
        }
        _ = v4
    }
    var v8 uint = 18446744073709551615
    var v9 float32 = 0.0
    fmt.Println(v9, v9)
    var v10 uint = v8
    v11 := GenS1{F0: -880, F1: 1, F2: 42}
    if (18446744073709551615 != 0) {
        switch 2147483647 {
        case 32767:
            if false {
                var v12 string = "hello"
                var v13 uint = (v8 >> ((v8) & 63))
                _ = v12
                _ = v13
            } else {
                v10 += 2
            }
        case 2147483647:
            var v14 int32 = 2147483647
            _ = v14
        default:
            v8 -= v10
        }
    }
    v10 = 42
    v9 = (v9 + v9)
    v8 += 1
    switch -1 {
    case 128:
        v2[65535] = ("\n\t" + "foo bar")
    case 88:
        switch 32767 {
        case 127:
            v9 -= (v9 + v9)
        default:
            var v15 uint64 = 2
            _ = v15
        }
    }
    var v16 float64 = 1e10
    var v17 uint = v10
    v17 = 13
    v9 = v9
    var v18 uint = (^(v10))
    if (1000 <= -711) {
        if (924 == 2147483647) {
            fmt.Println(v10, v8)
            v8 = 42
        } else {
            if (-128 < -1) {
                if ((921 <= 1) || false) {
                    var v19 int64 = -1
                    if true {
                        switch "a" {
                        case "foo bar":
                            v3[-9223372036854775808] = (!(0 < 42))
                        case "":
                            fmt.Println(v9, v17)
                        case "a":
                            v20 := GenS2{F0: -1}
                            _ = v20
                        }
                        v10 = v17
                    } else {
                        if (-102 == -1) {
                            v18 = (v8 + v10)
                        } else {
                            var v21 uint8 = 2
                            _ = v21
                        }
                    }
                    var v22 uint8 = uint8(v8)
                    _ = v19
                    _ = v22
                } else {
                    if (128 == -616) {
                        var v23 int64 = 42
                        fmt.Println(v23)
                        _ = v23
                    } else {
                        v24 := GenS2{F0: 42}
                        if false {
                            if (v9 >= (0.1 - v9)) {
                                v25 := GenS2{F0: -1}
                                var v26 uint = (^(v8))
                                if (-1 > 9223372036854775807) {
                                    switch -143 {
                                    case -9223372036854775808:
                                        v27 := []int{}
                                        _ = v27
                                    case 127:
                                        var v28 uint16 = 42
                                        _ = v28
                                    case 2147483647:
                                        v29 := GenS3{F0: 1, F1: -9223372036854775808, F2: 32768}
                                        _ = v29
                                    default:
                                        v16 = 0.1
                                    }
                                    v9 = (v9 + ((1.5 * float32(v16)) * -2.25))
                                    v26 -= 0
                                } else {
                                    for fi30 := 0; fi30 < 19; fi30++ {
                                        v18 += (^(v17))
                                        _ = fi30
                                    }
                                }
                                _ = v25
                                _ = v26
                            }
                            fmt.Println(v10, v16, v18)
                        }
                        _ = v24
                    }
                    var v31 int8 = 127
                    if false {
                        if false {
                            var v32 int = 128
                            v16 = (v16 / 1.0)
                            _ = v32
                        } else {
                            v10 = (v8 << (((v18 << ((v17) & 63))) & 63))
                            var v33 float32 = -2.25
                            _ = v33
                        }
                    }
                    _ = v31
                }
                if false {
                    if false {
                        v34 := GenS4{F0: -9223372036854775808}
                        _ = v34
                    }
                }
            }
        }
        for fi35 := 0; fi35 < 13; fi35++ {
            fmt.Println(v8)
            if ((v10 & (v10 ^ 0)) > v8) {
                var v36 int32 = -386
                v37 := GenS2{F0: v36}
                _ = v36
                _ = v37
            }
            _ = fi35
        }
        for fi38 := 0; fi38 < 9; fi38++ {
            for fi39 := 0; fi39 < 7; fi39++ {
                v40 := GenS4{F0: fi39}
                var v41 float32 = (((0.1 - v9) / (v9 / ((v9 + v9) - v9))) + 1e10)
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi38
        }
    } else {
        var v42 uint8 = 1
        v43 := v3[127]
        v44 := GenS3{F0: v42, F1: -9223372036854775808, F2: int(v8)}
        _ = v42
        _ = v43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
