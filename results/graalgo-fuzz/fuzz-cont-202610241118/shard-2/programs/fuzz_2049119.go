package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 bool; F1 string; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int8; F2 int64 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint8 }
func genFn0() int64 {
    var v7 uint8 = 153
    var v8 uint64 = 337
    v9 := GenS1{F0: (v7 * v7), F1: -979}
    var v10 uint32 = 2147483648
    for fi11 := 0; fi11 < 2; fi11++ {
        switch ("hello" + "foo bar") {
        case "foo bar":
            v8 = (v8 / ((v8) | 1))
        case "":
            if false {
                if ((1e10 - ((float32(float64(1e10)) + (1.5 - (3.14159 - 0.1))) + 1.5)) <= 1e10) {
                    var v12 float32 = 1.5
                    v13 := GenS2{F0: false, F1: ("\n\t" + (("\n\t" + "hello") + "hello")), F2: 65535}
                    _ = v12
                    _ = v13
                } else {
                    if true {
                        var v14 int8 = -1
                        v8 = v8
                        _ = v14
                    }
                }
                fmt.Println(v7, v8)
                for fi15 := 0; fi15 < 14; fi15++ {
                    v8 = 0
                    v7 = (v7 << ((42) & 63))
                    _ = fi15
                }
            } else {
                var v16 uint16 = 0
                _ = v16
            }
        case "hello":
            var v17 float64 = 1.5
            _ = v17
        default:
            v18 := []int{-2147483648, fi11, fi11}
            _ = v18
        }
        var v19 uint32 = v10
        v19 += 2
        _ = fi11
        _ = v19
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -707
}
func genFn1(p0 uint64, p1 uint32, p2 uint16) uint {
    var v20 int16 = 0
    var v21 int8 = -128
    v22 := GenS1{F0: 127, F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v22
    return 698
}

func main() {
    var v1 float32 = -2.25
    v2 := []int{-657, 251}
    v3 := []uint32{2}
    var v4 int32 = -1
    var v5 string = "\n\t"
    v6 := make(map[string]int)
    if (("foo bar" + v5) <= "a") {
        v4 += v4
    } else {
        v5 = ""
        fmt.Println(v5)
        var v23 int8 = -128
        _ = v23
    }
    _ = len(v6)
    var v24 bool = (42 == 2147483648)
    if true {
        v25 := GenS3{F0: 18446744073709551615, F1: -105, F2: 9223372036854775807}
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 int32 = 42
            _ = fi26
            _ = v27
        }
        var v28 uint32 = 560
        _ = v25
        _ = v28
    } else {
        var v29 uint = 279
        var v30 int = -132
        _ = v29
        _ = v30
    }
    switch ((v5 + v5) + "foo bar") {
    case "":
        for fi31 := 0; fi31 < 1; fi31++ {
            if v24 {
                v32 := GenS2{F0: (78 <= 84), F1: ((v5 + (v5 + v5)) + ("a" + v5)), F2: 2}
                v33 := GenS4{F0: v24, F1: "a"}
                v1 -= v1
                _ = v32
                _ = v33
            }
            var v34 uint64 = 1
            _ = fi31
            _ = v34
        }
    case "hello":
        v4 = 0
    case "a":
        fmt.Println(v1, v4)
    default:
        fmt.Println(v4)
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        var v36 int32 = (v4 >> ((65535) & 63))
        v24 = ((1.5 / 0.1) >= 0.0)
        if (42 >= 9) {
            switch "\n\t" {
            case "\n\t":
                v1 = 0.1
            case "hello":
                if (false && (!(true || v24))) {
                    var v37 int = fi35
                    v37 -= 816
                    _ = v37
                }
            }
        } else {
            fmt.Println(v24, v1, v5)
            for fi38 := 0; fi38 < 20; fi38++ {
                var v39 int32 = v36
                var v40 uint = 682
                v4 += v39
                _ = fi38
                _ = v39
                _ = v40
            }
        }
        _ = fi35
        _ = v36
    }
    for fi41 := 0; fi41 < 5; fi41++ {
        fmt.Println(v1, v1, fi41)
        for fi42 := 0; fi42 < 11; fi42++ {
            fmt.Println(v1, v5)
            _ = fi42
        }
        _ = fi41
    }
    fmt.Println(v4, v24)
    switch "" {
    case "foo bar":
        for fi43 := 0; fi43 < 11; fi43++ {
            v44 := GenS5{F0: 255}
            _ = fi43
            _ = v44
        }
    default:
        v4 = (v4 * v4)
    }
    v5 = ("foo bar" + "foo bar")
    v45 := GenS4{F0: (961 >= 981), F1: v5}
    switch "hello" {
    case "":
        var v46 int8 = 100
        _ = v46
    case "foo bar":
        v24 = (!v24)
    default:
        fmt.Println(v1)
    }
    fmt.Println(v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
