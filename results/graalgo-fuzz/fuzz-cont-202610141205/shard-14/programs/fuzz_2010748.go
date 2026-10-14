package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int }
type GenS4 struct { F0 int16; F1 int8; F2 bool }
type GenS5 struct { F0 int64; F1 int; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 32767
    var v3 uint64 = 18446744073709551615
    v4 := GenS1{F0: 1}
    v3 = (^(v3))
    v2 = 2
    var v5 string = "a"
    v6 := GenS2{F0: (v3 & v3)}
    if (266 > 9223372036854775807) {
        v2 = v2
        if (true || false) {
            switch 987 {
            case 32768:
                v3 += (v3 * 2)
            case -817:
                for fi7 := 0; fi7 < 7; fi7++ {
                    for fi8 := 0; fi8 < 19; fi8++ {
                        delete(v1, (fi7 * 32767))
                        v9 := GenS3{F0: (v3 ^ v3), F1: 0, F2: (fi8 / ((127) | 1))}
                        v10 := GenS2{F0: v3}
                        v2 += 65535
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    _ = fi7
                }
            default:
                if false {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v3 -= (^(v3))
                        v5 = v5
                        var v12 uint64 = (^(v3))
                        _ = fi11
                        _ = v12
                    }
                    v5 = (v5 + (v5 + v5))
                } else {
                    switch v5 {
                    case "":
                        if (-2147483648 <= -1) {
                            v13 := []int{-1, -2147483648, 42}
                            if false {
                                v3 += (v3 ^ (^(v3)))
                            }
                            _ = v13
                        } else {
                            var v14 int32 = 1
                            _ = v14
                        }
                    case "foo bar":
                        if (844 >= 990) {
                            v5 = v5
                            for fi15 := 0; fi15 < 6; fi15++ {
                                var v16 uint = 2
                                for fi17 := 0; fi17 < 2; fi17++ {
                                    if false {
                                        v16 += (^(v16))
                                        if (true || (42 == 2147483647)) {
                                            var v18 int64 = 1
                                            _ = v18
                                        }
                                        v3 = v3
                                        var v19 uint8 = 254
                                        v20 := []int{fi17, int(v2)}
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = fi17
                                }
                                _ = fi15
                                _ = v16
                            }
                        }
                    }
                }
            }
            v3 = (v3 / ((v3) | 1))
        }
    }
    v21 := GenS4{F0: int16(v2), F1: 0, F2: (!(v2 < v2))}
    v22 := GenS3{F0: v3, F1: 32767, F2: 4294967295}
    v3 -= v3
    v2 = (^(v2))
    fmt.Println(v5)
    switch v5 {
    case "":
        fmt.Println(v2, v5)
    case "a":
        switch "" {
        case "hello":
            var v23 uint64 = v3
            _ = v23
        }
    default:
        switch 1 {
        case -151:
            var v24 bool = false
            _ = v24
        case 948:
            switch v5 {
            case "hello":
                for fi25 := 0; fi25 < 2; fi25++ {
                    var v26 string = v5
                    if true {
                        var v27 uint32 = 42
                        _ = v27
                    } else {
                        v5 = v26
                        var v28 bool = (true && true)
                        v26 = v5
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
            case "a":
                fmt.Println(v2, v2)
            case "\n\t":
                v3 = (v3 & (v3 - v3))
            }
        }
    }
    v29 := GenS5{F0: -9223372036854775808, F1: -194, F2: 32767}
    switch 185 {
    case 32768:
        var v30 uint8 = 128
        _ = v30
    case 0:
        v31 := GenS3{F0: v3, F1: v2, F2: -2147483648}
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
