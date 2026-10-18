package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 int8; F1 bool }
func genFn0(p0 int64, p1 uint32, p2 bool) string {
    for fi4 := 0; fi4 < 15; fi4++ {
        if ((^(fi4)) > 128) {
            v5 := []int{(fi4 / ((fi4) | 1)), -9223372036854775808, -2147483648, fi4}
            _ = v5
        }
        v6 := GenS1{F0: 18446744073709551615, F1: -128}
        _ = fi4
        _ = v6
    }
    p1 = p1
    v7 := GenS2{F0: -1, F1: 903, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}
func genFn1(p0 string, p1 int8) int64 {
    p0 = p0
    for fi8 := 0; fi8 < 15; fi8++ {
        switch (^(fi8)) {
        case 693:
            v9 := []int{fi8, fi8, (^(fi8))}
            _ = v9
        default:
            if true {
                var v10 bool = false
                var v11 uint32 = 847
                switch "\n\t" {
                case "hello":
                    v12 := []int{}
                    _ = v12
                case "a":
                    v13 := GenS3{F0: 32767}
                    _ = v13
                default:
                    p0 = (("foo bar" + "\n\t") + (((p0 + ("a" + (p0 + (p0 + p0)))) + p0) + "hello"))
                }
                _ = v10
                _ = v11
            } else {
                p1 += (p1 - (-(p1)))
            }
        }
        if (!(42 < 162)) {
            if false {
                switch p0 {
                case "hello":
                    v14 := GenS4{F0: 2147483648, F1: (fi8 / ((fi8) | 1)), F2: 555}
                    _ = v14
                }
            }
            p0 = (p0 + p0)
        }
        _ = fi8
    }
    p0 = p0
    p1 = (p1 * -42)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int32 = 148
    v2 := make(map[uint32]int)
    var v3 uint64 = 235
    var v15 uint16 = 74
    v16 := GenS5{F0: v1, F1: 462, F2: -2147483648}
    var v17 float64 = -2.25
    var v18 string = ("foo bar" + "foo bar")
    var v19 uint16 = (v15 << ((255) & 63))
    switch 9223372036854775807 {
    case 32767:
        if true {
            v15 = (v19 | 2)
            var v20 float32 = 1.5
            _ = v20
        } else {
            v1 = 2147483647
        }
    case 127:
        var v21 int16 = -32768
        _ = v21
    case 9223372036854775807:
        var v22 uint8 = 238
        _ = v22
    }
    v23 := GenS6{F0: uint16(v3), F1: 798}
    if (true || true) {
        fmt.Println(v17)
    } else {
        switch 32767 {
        case 128:
            v24 := []int{32767}
            _ = v24
        case 32768:
            for fi25 := 0; fi25 < 6; fi25++ {
                switch v18 {
                case "":
                    fmt.Println(v15, v18)
                case "foo bar":
                    if false {
                        var v26 uint32 = 2147483648
                        _ = v26
                    }
                default:
                    var v27 int32 = (-(v1))
                    _ = v27
                }
                _ = fi25
            }
        case -289:
            var v28 uint16 = (v15 >> ((628) & 63))
            _ = v28
        }
        v29 := v2[2]
        var v30 bool = (!(-80 >= -128))
        _ = v29
        _ = v30
    }
    if false {
        fmt.Println(v19, v3)
    }
    v31 := GenS6{F0: 424, F1: -32768}
    var v32 int = 4294967295
    v33 := GenS7{F0: 65535, F1: 65535}
    v17 = 0.1
    v34 := GenS8{F0: -128, F1: true}
    fmt.Println(v18)
    var v35 uint64 = v3
    switch (v18 + (v18 + ("hello" + v18))) {
    case "\n\t":
        switch v18 {
        case "foo bar":
            if false {
                v36 := []int{1, v32}
                if true {
                    fmt.Println(v3, v35)
                }
                _ = v36
            }
        case "":
            v37 := []int{v32, v32}
            _ = v37
        case "hello":
            var v38 float64 = v17
            _ = v38
        }
    case "":
        var v39 int32 = v1
        _ = v39
    default:
        v2[999] = (-(v32))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
