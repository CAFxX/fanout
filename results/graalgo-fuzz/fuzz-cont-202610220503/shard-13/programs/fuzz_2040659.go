package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0(p0 uint, p1 int64, p2 uint) uint64 {
    if (!(-1 != -122)) {
        var v6 uint8 = 128
        p2 += (p0 * p2)
        p2 = (p2 / (((p0 / (((p2 >> ((0) & 63))) | 1))) | 1))
        _ = v6
    }
    fmt.Println(p2, p0, p2)
    p2 = p2
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 0
    v3 := make(map[uint32]uint32)
    v4 := []uint32{42}
    var v5 uint = 65535
    var v7 int64 = (v2 | (^(v2)))
    switch "" {
    case "hello":
        if (4294967295 > 679) {
            v7 += 42
            v3[2147483648] = 502
        } else {
            if (true || false) {
                var v8 string = ("hello" + "hello")
                _ = v8
            }
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint16 = 2
                var v11 int8 = 101
                _ = fi9
                _ = v10
                _ = v11
            }
            v7 = (v2 % ((42) | 1))
        }
    case "\n\t":
        _ = len(v4)
    case "a":
        fmt.Println(v2, v7)
    }
    var v12 uint8 = 127
    var v13 bool = (3.14159 != (0.0 - -2.25))
    v5 += (v5 / (((^(v5))) | 1))
    if (2 == 0) {
        fmt.Println(v5, v5, v7)
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            if v13 {
                fmt.Println(v5)
            }
            fmt.Println(v7, v12, v5)
            _ = fi14
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 int16 = 1
            _ = len(v1)
            _ = fi15
            _ = v16
        }
        v2 = v2
    }
    var v17 int8 = 1
    var v18 float32 = 0.0
    if (2 < 65535) {
        switch 9223372036854775807 {
        case -9223372036854775808:
            v19 := GenS1{F0: (-(v17)), F1: "a"}
            _ = v19
        case 292:
            var v20 uint16 = 484
            _ = v20
        case 2147483647:
            var v21 int16 = -526
            _ = v21
        default:
            fmt.Println(v13)
        }
        if v13 {
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 int8 = v17
                switch "foo bar" {
                case "":
                    var v24 uint64 = 42
                    _ = v24
                case "hello":
                    if (32767 < 42) {
                        var v25 string = ("foo bar" + "")
                        _ = v25
                    }
                }
                _ = fi22
                _ = v23
            }
            v26 := GenS1{F0: 127, F1: "foo bar"}
            _ = v26
        }
    }
    delete(v3, 1)
    _ = len(v4)
    if v13 {
        fmt.Println(v18, v18, v12)
    } else {
        if (!v13) {
            _ = len(v4)
        } else {
            switch -991 {
            case 127:
                v18 = v18
            case 32767:
                var v27 bool = v13
                _ = v27
            case 0:
                fmt.Println(v2, v7, v13)
            }
            var v28 uint = v5
            _ = len(v4)
            _ = v28
        }
        fmt.Println(v2)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, 2147483647)
    var v29 uint = (v5 ^ (v5 >> ((v5) & 63)))
    v5 -= (v5 << ((888) & 63))
    fmt.Println(v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
