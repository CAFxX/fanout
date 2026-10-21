package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 uint; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int8; F1 string; F2 uint32 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 int32; F1 uint }
func genFn0() uint16 {
    v5 := []int{-9223372036854775808, -1, 32767}
    var v6 uint16 = 2
    v6 += (^(v6))
    v7 := GenS1{F0: 4294967295}
    _ = v5
    _ = v6
    _ = v7
    return (v6 << ((18446744073709551615) & 63))
}
func genFn1(p0 bool, p1 int8, p2 uint64) uint32 {
    switch "foo bar" {
    case "a":
        var v8 int = 127
        _ = v8
    }
    var v9 uint = 42
    var v10 string = ("a" + (((("foo bar" + "a") + "hello") + (("" + "foo bar") + "foo bar")) + (("foo bar" + "\n\t") + "")))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 2147483648
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    v3 := []uint32{65535, 877, 4294967295}
    var v4 uint32 = 2147483647
    switch ((("hello" + ("\n\t" + "a")) + "hello") + "\n\t") {
    case "foo bar":
        switch (("foo bar" + "") + (("\n\t" + "a") + ("" + "a"))) {
        case "":
            var v11 uint64 = 18446744073709551615
            _ = v11
        case "hello":
            if true {
                if true {
                    fmt.Println(v4)
                } else {
                    var v12 uint16 = 42
                    var v13 uint32 = (v4 >> ((18446744073709551615) & 63))
                    v13 = v13
                    _ = v12
                    _ = v13
                }
            }
        case "foo bar":
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int64 = -1
                v4 = (v4 ^ v4)
                _ = fi14
                _ = v15
            }
        default:
            var v16 uint8 = 255
            _ = v16
        }
    case "\n\t":
        _ = len(v3)
    default:
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS2{F0: 1, F1: false}
            _ = fi17
            _ = v18
        }
    }
    if true {
        fmt.Println(v4)
        v4 -= v4
        v19 := GenS1{F0: 65535}
        _ = v19
    } else {
        v4 = 132
        fmt.Println(v4)
        v20 := GenS3{F0: 0, F1: 65535, F2: -128}
        _ = v20
    }
    v4 -= (v4 >> ((65535) & 63))
    v21 := GenS3{F0: 524, F1: 65535, F2: -128}
    v3 = append(v3, (v4 << ((671) & 63)))
    v4 = 2147483648
    if (1.5 >= (0.1 + -2.25)) {
        switch ("a" + (("foo bar" + ("" + "hello")) + ("\n\t" + "hello"))) {
        case "foo bar":
            v22 := GenS4{F0: 32767}
            _ = v22
        case "a":
            var v23 int32 = 85
            _ = v23
        }
    } else {
        v24 := GenS5{F0: ("\n\t" + ""), F1: 18446744073709551615}
        _ = v24
    }
    switch "foo bar" {
    case "\n\t":
        v4 -= 308
    case "":
        _ = len(v1)
    case "a":
        switch (("a" + ("" + "foo bar")) + ("foo bar" + "")) {
        case "foo bar":
            if false {
                v25 := v2["\n\t"]
                v25 = false
                _ = v25
            } else {
                var v26 int64 = 1
                v4 = 547
                v3 = append(v3, 0)
                _ = v26
            }
        case "\n\t":
            if true {
                var v27 uint64 = 243
                var v28 uint32 = (^(v4))
                _ = v27
                _ = v28
            } else {
                _ = len(v3)
            }
        }
    }
    switch "\n\t" {
    case "":
        switch -9223372036854775808 {
        case -1:
            v29 := GenS6{F0: 96, F1: (("hello" + "hello") + "a"), F2: (v4 ^ (v4 % ((genFn1(false, 127, 2)) | 1)))}
            _ = v29
        case 952:
            _ = len(v2)
        }
    default:
        v4 = v4
    }
    v4 = genFn1((0 > 1), 0, 1)
    fmt.Println(v4)
    v4 -= genFn1(true, -128, 4294967295)
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 13; fi31++ {
            if true {
                switch fi31 {
                case 32767:
                    if (!(0 < 255)) {
                        var v32 float32 = 0.0
                        _ = v32
                    } else {
                        fmt.Println(fi30, v4)
                        var v33 uint8 = 127
                        _ = v33
                    }
                case 42:
                    v34 := GenS3{F0: -660, F1: 2, F2: 67}
                    _ = v34
                case 4294967295:
                    v35 := GenS7{F0: 872, F1: (^(v4))}
                    _ = v35
                default:
                    v4 = (v4 % ((v4) | 1))
                }
                fmt.Println(fi30, fi30, fi30)
                if (0 > (-(fi30))) {
                    v4 = 441
                    var v36 uint8 = 255
                    var v37 uint8 = v36
                    if (0 == 255) {
                        var v38 float64 = float64(3.14159)
                        _ = v38
                    }
                    switch "\n\t" {
                    case "a":
                        var v39 int64 = int64(fi31)
                        _ = v39
                    }
                    var v40 int = 128
                    v41 := GenS8{F0: -2147483648, F1: 556}
                    _ = v36
                    _ = v37
                    _ = v40
                    _ = v41
                } else {
                    var v42 int16 = -1
                    _ = v42
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v4 = v4
            fmt.Println(v4)
            _ = fi31
        }
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
