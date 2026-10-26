package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint8, p2 uint8) int64 {
    v7 := []int{}
    var v8 uint8 = (p1 << ((18446744073709551615) & 63))
    v8 += 1
    fmt.Println(p0, p1, p1)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(p1)
    case "a":
        fmt.Println(p0, p0, v8)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int32) int32 {
    v9 := GenS1{F0: p0}
    for fi10 := 0; fi10 < 2; fi10++ {
        p0 = 530
        var v11 uint16 = 255
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 42, F1: 42, F2: 127}
    switch 128 {
    case -2147483648:
        fmt.Println(p0)
    default:
        var v13 int16 = 42
        _ = v13
    }
    v14 := GenS3{F0: 42, F1: 1}
    _ = p0
    _ = v12
    _ = v14
    _ = v9
    return (p0 * -2147483648)
}

func main() {
    var v1 uint16 = 42
    var v2 bool = (100 != 1)
    var v3 uint16 = (v1 * v1)
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 2
    var v6 int32 = 2147483647
    var v15 int64 = 1
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int32 = (v6 | (v6 ^ v6))
        _ = fi16
        _ = v17
    }
    fmt.Println(v5, v5, v1)
    var v18 int16 = int16(v5)
    v15 = 9223372036854775807
    if v2 {
        if (42 > 65535) {
            v19 := []int{32768, 32768, 128, 42}
            var v20 uint16 = (v1 % ((v3) | 1))
            switch -949 {
            case 9223372036854775807:
                if false {
                    var v21 string = ("foo bar" + ("hello" + "a"))
                    _ = v21
                }
            case -704:
                fmt.Println(v3, v1, v20)
            case 42:
                v2 = (255 >= 2)
            }
            _ = v19
            _ = v20
        } else {
            var v22 int8 = -128
            for fi23 := 0; fi23 < 6; fi23++ {
                _ = len(v4)
                var v24 uint16 = (v1 - (v3 * (v3 + (v1 % ((554) | 1)))))
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        var v25 uint32 = 668
        var v26 int8 = -48
        _ = v25
        _ = v26
    } else {
        var v27 uint16 = v1
        _ = v27
    }
    fmt.Println(v3, v3)
    v6 = genFn1((-(v6)))
    fmt.Println(v18)
    switch 127 {
    case 127:
        fmt.Println(v18, v18)
    default:
        for fi28 := 0; fi28 < 11; fi28++ {
            var v29 bool = v2
            switch ("foo bar" + "a") {
            case "hello":
                for fi30 := 0; fi30 < 15; fi30++ {
                    v31 := []int{int(v15), 42}
                    switch "foo bar" {
                    case "a":
                        v18 = int16(v15)
                    }
                    _ = fi30
                    _ = v31
                }
            case "\n\t":
                if ((-(v15)) == genFn0((v18 / ((int16(v15)) | 1)), 42, 128)) {
                    v32 := []int{int(v15)}
                    _ = v32
                }
            case "foo bar":
                v6 = (v6 + (v6 - (v6 % ((v6) | 1))))
            default:
                if (!(-2.25 > float32(((((0.1 * 1.5) * 0.0) / 1.0) + 0.1)))) {
                    v33 := v4[65535]
                    var v34 uint32 = v33
                    switch "a" {
                    case "hello":
                        var v35 int32 = (v6 & (v6 >> ((2) & 63)))
                        _ = v35
                    default:
                        v18 = v18
                    }
                    _ = v33
                    _ = v34
                }
            }
            v29 = true
            _ = fi28
            _ = v29
        }
    }
    if false {
        v36 := GenS4{F0: (v5 >> ((42) & 63))}
        delete(v4, 4294967295)
        _ = v36
    } else {
        var v37 uint = 18446744073709551615
        _ = v37
    }
    v5 = v5
    if (127 <= 128) {
        switch "a" {
        case "":
            v2 = v2
        case "foo bar":
            v5 += 2
        }
        switch -1 {
        case 42:
            v15 = (v15 << ((2) & 63))
        }
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        var v39 uint16 = (v3 + v3)
        _ = fi38
        _ = v39
    }
    if (v2 || v2) {
        v6 -= -243
    } else {
        var v40 uint64 = v5
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
