package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 string; F2 int64 }
func genFn0(p0 int16) uint16 {
    var v6 string = ""
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 32767
}
func genFn1() uint32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint32 = 0
            _ = fi8
            _ = v9
        }
        var v10 uint8 = 255
        _ = fi7
        _ = v10
    }
    switch 32768 {
    case 42:
        var v11 string = "\n\t"
        _ = v11
    }
    if false {
        var v12 uint = 0
        var v13 int16 = 1000
        _ = v12
        _ = v13
    } else {
        if (100 >= -49) {
            v14 := GenS1{F0: 622, F1: "foo bar", F2: 125}
            var v15 int8 = 100
            _ = v14
            _ = v15
        } else {
            var v16 bool = (!(148 != 2))
            v17 := []int{-9223372036854775808, 32767, 42}
            _ = v16
            _ = v17
        }
    }
    return 241
}

func main() {
    v1 := []uint32{42, 1}
    var v2 uint = 18446744073709551615
    var v3 uint8 = 255
    var v4 uint = (v2 ^ v2)
    v5 := []string{"\n\t"}
    v4 = 63
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 += (v3 << (((v2 ^ (v4 << (((v4 % ((0) | 1))) & 63)))) & 63))
    v4 = v4
    v18 := GenS1{F0: 1, F1: "foo bar", F2: 127}
    for fi19 := 0; fi19 < 12; fi19++ {
        switch (fi19 & (fi19 * 476)) {
        case -135:
            fmt.Println(v4, v2, v2)
        default:
            var v20 int = (fi19 ^ (fi19 / ((fi19) | 1)))
            _ = v20
        }
        if false {
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 int16 = 375
                v1 = append(v1, 4294967295)
                v4 += uint(v3)
                v2 = v2
                _ = fi21
                _ = v22
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(fi19, fi19)
            v2 = v2
        }
        var v23 uint16 = 2
        _ = fi19
        _ = v23
    }
    fmt.Println(v3, v4)
    switch "hello" {
    case "hello":
        var v24 int64 = -9223372036854775808
        _ = v24
    case "":
        var v25 int8 = 15
        _ = v25
    case "\n\t":
        v2 = v4
    default:
        for fi26 := 0; fi26 < 2; fi26++ {
            v4 += 895
            _ = fi26
        }
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        if (v2 == (v4 | 255)) {
            _ = len(v1)
        }
        if true {
            v1 = append(v1, 685)
            var v28 int64 = 958
            _ = v28
        } else {
            if (-2.25 < 1.5) {
                if len(v1) > 0 {
                    v1[0] = 2
                }
                var v29 uint16 = 32767
                _ = v29
            } else {
                v30 := GenS2{F0: 1, F1: 2147483648}
                _ = v30
            }
            v31 := GenS2{F0: 1, F1: 2147483648}
            var v32 uint16 = 1
            _ = v31
            _ = v32
        }
        _ = fi27
    }
    switch "hello" {
    case "foo bar":
        if len(v5) > 0 {
            v5[0] = ((("" + "hello") + "foo bar") + "hello")
        }
    case "a":
        var v33 bool = (4 >= -2147483648)
        _ = v33
    default:
        if false {
            for fi34 := 0; fi34 < 15; fi34++ {
                for fi35 := 0; fi35 < 8; fi35++ {
                    v1 = append(v1, 4294967295)
                    v3 = (^(v3))
                    var v36 int64 = 1
                    _ = fi35
                    _ = v36
                }
                _ = fi34
            }
            var v37 int64 = -9223372036854775808
            _ = v37
        } else {
            fmt.Println(v2, v2, v3)
            for fi38 := 0; fi38 < 18; fi38++ {
                switch "foo bar" {
                case "":
                    v39 := GenS1{F0: 1, F1: ("\n\t" + ("\n\t" + ("" + ("" + "\n\t")))), F2: -128}
                    _ = v39
                case "a":
                    v2 = (v2 * (v2 ^ (^(v2))))
                case "\n\t":
                    fmt.Println(fi38, v3, fi38)
                default:
                    var v40 uint = 18446744073709551615
                    _ = v40
                }
                v41 := GenS3{F0: 42}
                _ = fi38
                _ = v41
            }
        }
    }
    v42 := GenS4{F0: v3, F1: (((("a" + "a") + "foo bar") + "a") + "\n\t"), F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
