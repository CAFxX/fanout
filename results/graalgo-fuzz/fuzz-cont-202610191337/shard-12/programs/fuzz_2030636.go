package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint64 = 852
    var v2 float32 = (float32(1.5) / (-2.25 - 3.14159))
    v3 := make(map[string]int)
    var v4 uint64 = (v1 % (((v1 & 18446744073709551615)) | 1))
    v5 := make(map[uint32]bool)
    var v6 int32 = 0
    v7 := []int{}
    var v8 float64 = float64(v2)
    if true {
        v9 := GenS1{F0: (465 > (v1 << ((4294967295) & 63)))}
        switch "" {
        case "a":
            switch 1 {
            case 0:
                var v10 int16 = 32767
                _ = v10
            case 2147483647:
                fmt.Println(v4)
            }
        case "\n\t":
            var v11 int = -9223372036854775808
            _ = v11
        case "":
            v12 := GenS2{F0: 28, F1: 2}
            _ = v12
        }
        v13 := GenS1{F0: false}
        _ = v13
        _ = v9
    } else {
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 float32 = v2
            fmt.Println(v15, v4, v8)
            _ = fi14
            _ = v15
        }
    }
    switch -9223372036854775808 {
    case -778:
        delete(v5, 2147483648)
    case -9223372036854775808:
        fmt.Println(v8, v2, v4)
    case -849:
        v8 = (-2.25 - v8)
    }
    if (936 < 1) {
        if true {
            for fi16 := 0; fi16 < 15; fi16++ {
                if (1 == (^(v6))) {
                    _ = len(v7)
                    var v17 uint8 = 128
                    _ = v17
                }
                for fi18 := 0; fi18 < 8; fi18++ {
                    _ = len(v3)
                    var v19 float32 = v2
                    v19 = v2
                    switch -883 {
                    case 0:
                        if true {
                            v1 = v4
                            var v20 uint16 = 255
                            _ = v20
                        } else {
                            fmt.Println(v2, v1)
                            var v21 int8 = 100
                            var v22 float32 = (3.14159 - (-2.25 - float32(0.1)))
                            switch ("" + "hello") {
                            case "":
                                _ = len(v7)
                            case "a":
                                fmt.Println(v22, v21, v21)
                            case "foo bar":
                                var v23 int64 = 1
                                _ = v23
                            }
                            _ = v21
                            _ = v22
                        }
                    }
                    fmt.Println(v6, fi18, v8)
                    _ = fi18
                    _ = v19
                }
                _ = fi16
            }
            var v24 int64 = 9223372036854775807
            _ = v24
        }
        v6 -= -2147483648
    }
    v4 -= (v4 >> ((2) & 63))
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint = 4294967295
        _ = fi25
        _ = v26
    }
    if false {
        v1 += (v4 ^ 2)
        fmt.Println(v8)
    } else {
        v27 := GenS3{F0: 0, F1: 64}
        v6 = v6
        v4 += v1
        _ = v27
    }
    v8 = (0.1 + -2.25)
    switch 1 {
    case 0:
        fmt.Println(v1, v6, v2)
    case 127:
        v8 += v8
    case 4294967295:
        fmt.Println(v2)
    default:
        v8 = (0.1 + v8)
    }
    v28 := v5[42]
    v29 := GenS2{F0: 0, F1: 1}
    var v30 uint64 = v4
    v30 += (v1 | v4)
    v31 := GenS2{F0: 127, F1: 249}
    var v32 uint8 = 99
    for fi33 := 0; fi33 < 15; fi33++ {
        var v34 int8 = 0
        v32 -= 127
        if v28 {
            v35 := GenS2{F0: -4, F1: v32}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v35
        }
        _ = fi33
        _ = v34
    }
    v8 = (v8 + (v8 + (0.0 - v8)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
