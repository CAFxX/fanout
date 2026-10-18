package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int; F1 uint32; F2 int32 }

func main() {
    v1 := []uint32{943}
    var v2 uint32 = 2147483647
    v3 := []uint32{(v2 / ((780) | 1))}
    var v4 uint = 18446744073709551615
    v5 := make(map[int]uint32)
    v6 := GenS1{F0: -286, F1: 288, F2: 29}
    var v7 string = ("foo bar" + "\n\t")
    v2 = 2147483647
    if true {
        var v8 uint8 = 42
        v9 := GenS1{F0: -9223372036854775808, F1: 1, F2: 42}
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: 9223372036854775807, F1: 42, F2: fi10}
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    } else {
        fmt.Println(v7)
        var v12 int8 = 64
        _ = v12
    }
    var v13 string = "hello"
    v2 = 898
    var v14 int8 = 1
    v15 := GenS1{F0: -230, F1: 0, F2: 0}
    v16 := v5[468]
    var v17 uint64 = 1
    if true {
        switch 811 {
        case -9223372036854775808:
            delete(v5, -2147483648)
        case -154:
            fmt.Println(v4)
        case -1:
            if ("foo bar" > v13) {
                var v18 int64 = -9223372036854775808
                v18 -= (-(v18))
                _ = v18
            }
        default:
            var v19 uint8 = 0
            _ = v19
        }
        if true {
            for fi20 := 0; fi20 < 9; fi20++ {
                for fi21 := 0; fi21 < 19; fi21++ {
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    if (-898 <= (fi21 + (fi20 % (((fi20 * fi21)) | 1)))) {
                        var v22 int32 = 1000
                        fmt.Println(v2, v13)
                        var v23 int32 = v22
                        _ = len(v1)
                        _ = v22
                        _ = v23
                    } else {
                        v24 := GenS2{F0: 0, F1: 127}
                        var v25 int32 = -1
                        _ = v24
                        _ = v25
                    }
                    var v26 int16 = 1
                    _ = fi21
                    _ = v26
                }
                if (-2.25 >= 0.1) {
                    v17 += v17
                    v27 := GenS2{F0: 574, F1: 1}
                    _ = v27
                } else {
                    fmt.Println(v13, v16, fi20)
                    var v28 int8 = (v14 >> (((v4 ^ 1)) & 63))
                    _ = v28
                }
                _ = fi20
            }
        }
    } else {
        v29 := GenS1{F0: -1, F1: 963, F2: 378}
        switch v13 {
        case "a":
            switch v7 {
            case "hello":
                v7 = v13
            case "foo bar":
                if (true && false) {
                    v30 := v5[-2147483648]
                    v31 := GenS1{F0: -9223372036854775808, F1: 42, F2: 32768}
                    if (false || false) {
                        v30 -= v2
                    } else {
                        fmt.Println(v16, v7, v14)
                        var v32 uint16 = 32767
                        for fi33 := 0; fi33 < 5; fi33++ {
                            switch fi33 {
                            case 9223372036854775807:
                                v30 += v2
                            case -2147483648:
                                var v34 uint8 = 0
                                _ = v34
                            case -9223372036854775808:
                                var v35 string = v7
                                _ = v35
                            }
                            var v36 uint8 = 0
                            _ = fi33
                            _ = v36
                        }
                        _ = v32
                    }
                    _ = v30
                    _ = v31
                }
            case "":
                v14 -= (v14 << ((494) & 63))
            }
        case "hello":
            var v37 string = "foo bar"
            _ = v37
        case "foo bar":
            v38 := GenS1{F0: 278, F1: 781, F2: -9223372036854775808}
            _ = v38
        default:
            var v39 uint64 = 4294967295
            _ = v39
        }
        _ = v29
    }
    switch 2147483647 {
    case 9223372036854775807:
        v40 := GenS3{F0: 2147483647, F1: v16, F2: 1}
        _ = v40
    case 42:
        var v41 uint16 = 0
        _ = v41
    case 4294967295:
        v14 = (v14 / ((v14) | 1))
    }
    v42 := GenS1{F0: -9223372036854775808, F1: uint16(v4), F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
