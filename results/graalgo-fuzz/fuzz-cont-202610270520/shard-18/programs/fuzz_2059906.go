package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint8, p1 uint, p2 uint) int16 {
    v7 := []int{9223372036854775807, 0, 0, 128}
    for fi8 := 0; fi8 < 13; fi8++ {
        p1 += p2
        var v9 uint8 = p0
        p2 = p2
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1() uint8 {
    v10 := GenS1{F0: 2147483647}
    v11 := GenS1{F0: 9223372036854775807}
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int16 = 32767
        switch "" {
        case "":
            fmt.Println(fi12, v13)
        default:
            v13 = v13
        }
        _ = fi12
        _ = v13
    }
    _ = v10
    _ = v11
    return 136
}

func main() {
    var v1 uint16 = 1
    v2 := []int{4294967295, -168, 713, 32768}
    v3 := make(map[int]bool)
    v4 := []uint32{2}
    var v5 int32 = 1
    var v6 uint64 = 2
    switch ((("foo bar" + "") + "") + "foo bar") {
    case "foo bar":
        v1 += (v1 - v1)
    case "a":
        v1 = (v1 * (v1 << ((18446744073709551615) & 63)))
    case "\n\t":
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 18; fi15++ {
                v1 += (^(v1))
                v5 += 42
                _ = fi15
            }
            _ = fi14
        }
    }
    v6 = 743
    var v16 uint8 = 157
    fmt.Println(v1, v16)
    _ = len(v3)
    v16 += 128
    switch ("hello" + "foo bar") {
    case "a":
        for fi17 := 0; fi17 < 12; fi17++ {
            v1 += (v1 & (v1 + v1))
            switch ("hello" + "\n\t") {
            case "foo bar":
                v16 = v16
            case "hello":
                v18 := GenS1{F0: fi17}
                _ = v18
            default:
                for fi19 := 0; fi19 < 18; fi19++ {
                    v20 := GenS1{F0: -9223372036854775808}
                    v21 := GenS1{F0: fi19}
                    var v22 int = fi19
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
            _ = fi17
        }
    }
    var v23 int32 = (^(v5))
    var v24 int8 = int8(v1)
    if false {
        if ((v6 * (v6 | (v6 & (v6 << ((0) & 63))))) > 4294967295) {
            v24 = v24
        }
    } else {
        if false {
            delete(v3, -9223372036854775808)
            fmt.Println(v23, v1)
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 uint8 = v16
                _ = fi25
                _ = v26
            }
        }
        v27 := v3[4294967295]
        var v28 int64 = 587
        _ = v27
        _ = v28
    }
    fmt.Println(v6, v16, v16)
    if (1000 >= -1) {
        v1 -= v1
        var v29 int = 127
        _ = v29
    } else {
        v30 := GenS1{F0: 2147483647}
        var v31 int16 = genFn0(v16, 591, 2)
        _ = v30
        _ = v31
    }
    if true {
        for fi32 := 0; fi32 < 6; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                v34 := GenS1{F0: fi32}
                switch "\n\t" {
                case "a":
                    fmt.Println(v1, v16)
                default:
                    delete(v3, fi33)
                }
                _ = fi33
                _ = v34
            }
            var v35 uint16 = (v1 | (^(v1)))
            var v36 uint8 = 127
            _ = fi32
            _ = v35
            _ = v36
        }
        var v37 uint = uint(v6)
        _ = len(v2)
        _ = v37
    } else {
        v38 := GenS2{F0: 1, F1: false, F2: uint64(v16)}
        fmt.Println(v1, v5, v16)
        var v39 uint64 = 18446744073709551615
        _ = v38
        _ = v39
    }
    v2 = append(v2, 127)
    var v40 uint32 = 65535
    if false {
        switch "" {
        case "foo bar":
            if true {
                var v41 float32 = (1.5 / float32(3.14159))
                v5 -= v5
                _ = v41
            }
        }
    } else {
        v5 = (v5 | int32(v6))
        v42 := GenS3{F0: v16, F1: v40}
        v24 = v24
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
