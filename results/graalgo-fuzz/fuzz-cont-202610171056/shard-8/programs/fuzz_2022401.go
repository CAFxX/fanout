package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 uint; F2 bool }
type GenS6 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 int64; F1 bool; F2 int32 }
type GenS9 struct { F0 uint; F1 int32; F2 uint32 }
func genFn0() int64 {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 2, F1: 9223372036854775807, F2: 127}
        switch (^(fi5)) {
        case -386:
            fmt.Println(fi5)
        case 9223372036854775807:
            v7 := GenS2{F0: 1, F1: (fi5 / (((fi5 << ((166) & 63))) | 1)), F2: 255}
            _ = v7
        case 127:
            fmt.Println(fi5)
        default:
            v8 := GenS3{F0: true, F1: -625}
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    var v9 uint = 2
    _ = v9
    return -506
}
func genFn1(p0 uint32, p1 uint8, p2 int64) uint64 {
    fmt.Println(p2, p1, p2)
    for fi10 := 0; fi10 < 7; fi10++ {
        p1 = 216
        for fi11 := 0; fi11 < 4; fi11++ {
            if (!true) {
                if ("a" >= ("hello" + "")) {
                    fmt.Println(fi10, p1, fi11)
                    p2 += (^(p2))
                } else {
                    var v12 uint16 = 980
                    _ = v12
                }
            }
            if (3.14159 <= 1e10) {
                v13 := []int{(fi10 >> ((2) & 63)), (fi10 | 32767)}
                fmt.Println(fi10)
                _ = v13
            }
            switch fi11 {
            case 42:
                if (255 > 2) {
                    switch fi10 {
                    case -2147483648:
                        v14 := GenS4{F0: 255}
                        _ = v14
                    case 9223372036854775807:
                        switch "a" {
                        case "\n\t":
                            var v15 int8 = 86
                            _ = v15
                        case "hello":
                            p1 = (p1 << ((21) & 63))
                        case "foo bar":
                            p0 += (p0 >> ((4294967295) & 63))
                        }
                    default:
                        if true {
                            p0 += (^(p0))
                        } else {
                            fmt.Println(p2)
                        }
                    }
                    p1 += (p1 * 42)
                }
            default:
                var v16 string = ("a" + "foo bar")
                _ = v16
            }
            _ = fi11
        }
        fmt.Println(p0)
        _ = fi10
    }
    v17 := GenS5{F0: (p1 ^ (p1 | 128)), F1: 65535, F2: true}
    v18 := []int{-9223372036854775808, 128, 1, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 float32 = ((float32((-2.25 + float64(float32((1.5 + (3.14159 - -2.25)))))) - 1e10) * 1e10)
    if true {
        switch ("hello" + "foo bar") {
        case "foo bar":
            v4 += v4
        default:
            v19 := GenS6{F0: 530, F1: 2147483647, F2: 255}
            _ = v19
        }
        v20 := GenS7{F0: 1, F1: 1}
        v4 = -2.25
        _ = v20
    }
    v21 := GenS4{F0: 247}
    v22 := GenS8{F0: genFn0(), F1: true, F2: 42}
    if (false && (genFn1(42, 127, 655) != 2)) {
        v4 = 0.1
        var v23 uint32 = 919
        v24 := []int{}
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        switch "a" {
        case "":
            var v26 int32 = 1
            _ = v26
        case "a":
            v27 := GenS9{F0: 18446744073709551615, F1: 2147483647, F2: 4294967295}
            _ = v27
        default:
            var v28 uint32 = 2147483648
            _ = v28
        }
        v4 += ((v4 + v4) * v4)
        _ = fi25
    }
    v4 = 3.14159
    switch 127 {
    case 127:
        var v29 int8 = -128
        _ = v29
    case 42:
        switch ("foo bar" + "\n\t") {
        case "":
            var v30 bool = false
            _ = v30
        case "\n\t":
            if (46 == 231) {
                v4 = 1e10
                if true {
                    switch 42 {
                    case 128:
                        v4 += (((0.1 + 1.5) - 1e10) + (0.0 * 0.1))
                    case 213:
                        if (((1e10 + 0.0) >= 0.1) && false) {
                            var v31 uint64 = 1
                            _ = v31
                        }
                    default:
                        var v32 float64 = float64(v4)
                        _ = v32
                    }
                    var v33 int16 = 0
                    for fi34 := 0; fi34 < 9; fi34++ {
                        for fi35 := 0; fi35 < 9; fi35++ {
                            v4 = 1e10
                            v36 := GenS7{F0: 2147483647, F1: -1}
                            _ = fi35
                            _ = v36
                        }
                        _ = fi34
                    }
                    _ = v33
                }
            }
        }
    }
    v4 -= 0.1
    for fi37 := 0; fi37 < 12; fi37++ {
        if true {
            if (3.14159 == 0.1) {
                for fi38 := 0; fi38 < 13; fi38++ {
                    v39 := GenS4{F0: 65535}
                    var v40 uint32 = 354
                    v40 = 2147483647
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            }
        } else {
            if (!(true && (-2147483648 == 42))) {
                v41 := GenS7{F0: -2147483648, F1: 42}
                fmt.Println(fi37, fi37)
                v4 = 3.14159
                _ = v41
            }
            v4 = v4
        }
        _ = fi37
    }
    var v42 uint32 = 266
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v42
}
