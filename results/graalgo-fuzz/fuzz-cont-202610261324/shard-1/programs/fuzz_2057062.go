package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := []uint32{1}
    var v2 int8 = 112
    var v3 uint32 = 2
    var v4 uint32 = v3
    v4 -= (^(v4))
    v5 := GenS1{F0: "\n\t", F1: 0}
    v3 += 2147483648
    var v6 string = "hello"
    v7 := GenS2{F0: -230, F1: 1, F2: (^(v3))}
    for fi8 := 0; fi8 < 10; fi8++ {
        switch 32767 {
        case 9223372036854775807:
            v3 += 0
        case 128:
            v1 = append(v1, (v3 | 2))
        case 4294967295:
            var v9 float32 = (0.1 - (1e10 + float32(1e10)))
            _ = v9
        default:
            if (true || false) {
                fmt.Println(v2)
                v4 -= v3
                var v10 uint8 = 97
                _ = v10
            } else {
                v4 -= (v3 * (v3 * v3))
                var v11 uint32 = (v3 & v3)
                _ = v11
            }
        }
        var v12 uint8 = 255
        _ = fi8
        _ = v12
    }
    if (0 <= 522) {
        v2 -= 91
        fmt.Println(v6, v2, v4)
        var v13 uint32 = v4
        _ = v13
    } else {
        v14 := GenS2{F0: 1, F1: (v3 | v4), F2: v3}
        var v15 int = 4294967295
        v1 = append(v1, 1)
        _ = v14
        _ = v15
    }
    v4 += 0
    var v16 int = -2147483648
    v17 := GenS2{F0: -1, F1: (^(v3)), F2: (v3 ^ 2147483647)}
    for fi18 := 0; fi18 < 10; fi18++ {
        v4 += (v4 ^ v4)
        v19 := GenS2{F0: -606, F1: 42, F2: v3}
        for fi20 := 0; fi20 < 9; fi20++ {
            _ = len(v1)
            var v21 int16 = 0
            if true {
                var v22 uint64 = 2
                if len(v1) > 0 {
                    v1[0] = v4
                }
                if (true || ((65535 > 65535) && false)) {
                    var v23 int32 = -1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v23
                } else {
                    v21 = -1
                    v4 = v3
                    fmt.Println(v21)
                    if false {
                        if ((("\n\t" + "hello") + v6) >= "\n\t") {
                            fmt.Println(v16)
                            fmt.Println(v16, v2, v22)
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                        }
                        var v24 string = v6
                        _ = v24
                    }
                    var v25 float64 = 0.0
                    _ = v25
                }
                _ = v22
            } else {
                v21 = -1
                var v26 string = ("hello" + "foo bar")
                v2 -= v2
                _ = v26
            }
            var v27 uint64 = 42
            var v28 uint16 = 1
            if (128 != 112) {
                var v29 uint8 = uint8(v28)
                var v30 float32 = (1.5 * -2.25)
                _ = v29
                _ = v30
            } else {
                fmt.Println(v16, fi20)
                if (true || false) {
                    v2 = (v2 << ((4294967295) & 63))
                    switch v6 {
                    case "hello":
                        v31 := GenS1{F0: v6, F1: 9223372036854775807}
                        v2 -= 42
                        _ = v31
                    default:
                        fmt.Println(v4, v4, v27)
                    }
                    var v32 int8 = v2
                    v21 = (v21 * -1)
                    fmt.Println(v6)
                    var v33 int = (fi18 ^ 1)
                    _ = v32
                    _ = v33
                } else {
                    fmt.Println(v3, v2)
                    if true {
                        v2 += (v2 % ((v2) | 1))
                        v34 := GenS1{F0: "\n\t", F1: int64(v28)}
                        _ = len(v1)
                        fmt.Println(v21, fi18, v2)
                        _ = v34
                    }
                    v3 += 4294967295
                }
            }
            v35 := GenS3{F0: v28, F1: v28}
            var v36 int16 = v21
            _ = fi20
            _ = v21
            _ = v27
            _ = v28
            _ = v35
            _ = v36
        }
        _ = fi18
        _ = v19
    }
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
