package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 string }
func genFn0(p0 int16) uint8 {
    var v7 string = ((("hello" + "hello") + "") + "a")
    switch ((("foo bar" + "a") + ("" + (("foo bar" + "foo bar") + v7))) + v7) {
    case "hello":
        v8 := GenS1{F0: 1, F1: v7, F2: -128}
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        switch v7 {
        case "\n\t":
            for fi10 := 0; fi10 < 6; fi10++ {
                v7 = ((("a" + ("hello" + v7)) + ("foo bar" + ((v7 + (v7 + "foo bar")) + v7))) + v7)
                var v11 string = v7
                v11 = v11
                _ = fi10
                _ = v11
            }
        case "a":
            var v12 int32 = 123
            _ = v12
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    return 23
}

func main() {
    v1 := []string{"\n\t", "a", "hello", ("\n\t" + "hello")}
    v2 := []int{1, 79}
    v3 := make(map[uint32]string)
    v4 := []string{}
    var v5 int32 = -2147483648
    var v6 uint8 = 0
    var v13 uint = 774
    if (!(!(false && true))) {
        for fi14 := 0; fi14 < 1; fi14++ {
            if true {
                var v15 bool = (!false)
                v5 += v5
                _ = v15
            }
            _ = fi14
        }
        fmt.Println(v13)
        for fi16 := 0; fi16 < 6; fi16++ {
            if true {
                v13 += (v13 | (v13 & v13))
                var v17 uint8 = v6
                _ = v17
            }
            _ = fi16
        }
    } else {
        if true {
            var v18 int64 = 1
            v13 += (v13 << ((v13) & 63))
            _ = v18
        }
        v19 := GenS1{F0: (v5 | (v5 << ((v13) & 63))), F1: ((("hello" + "\n\t") + ("" + "foo bar")) + "a"), F2: 127}
        _ = v19
    }
    switch "hello" {
    case "":
        if (-954 != 880) {
            v6 += v6
        } else {
            v13 -= 218
        }
    default:
        var v20 float64 = 0.1
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int16 = -820
        delete(v3, 2)
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint8 = (v6 ^ v6)
        _ = fi23
        _ = v24
    }
    fmt.Println(v6, v5, v6)
    fmt.Println(v13, v13)
    v25 := GenS2{F0: 0, F1: -1, F2: 32767}
    v5 = (v5 % ((-1) | 1))
    v3[2147483648] = "foo bar"
    if (18446744073709551615 <= 42) {
        var v26 uint16 = 32767
        fmt.Println(v5, v13)
        switch "foo bar" {
        case "\n\t":
            v5 = v5
        case "hello":
            _ = len(v2)
        default:
            v26 += (v26 / (((v26 << (((^(v13))) & 63))) | 1))
        }
        _ = v26
    } else {
        if true {
            var v27 uint16 = 32767
            _ = v27
        } else {
            var v28 int64 = 1
            v29 := GenS2{F0: 2147483647, F1: 9223372036854775807, F2: 1000}
            v5 = v5
            _ = v28
            _ = v29
        }
        if false {
            v5 += 1
            switch ("hello" + "hello") {
            case "\n\t":
                var v30 string = ("foo bar" + ((("foo bar" + "\n\t") + ("" + ("\n\t" + (("\n\t" + "hello") + ("foo bar" + ""))))) + (("foo bar" + "foo bar") + "foo bar")))
                _ = v30
            case "":
                var v31 string = ""
                _ = v31
            case "a":
                v32 := v3[2147483648]
                _ = v32
            }
            var v33 uint = (v13 + 1)
            _ = v33
        }
    }
    v5 += v5
    if ((v6 + (v6 % (((v6 >> ((4294967295) & 63))) | 1))) == 2) {
        if (2 <= 4294967295) {
            var v34 int64 = -456
            v35 := GenS3{F0: 1}
            _ = v34
            _ = v35
        } else {
            v13 += v13
            switch "hello" {
            case "\n\t":
                fmt.Println(v5)
            default:
                var v36 int16 = 32767
                _ = v36
            }
            var v37 uint16 = uint16(v6)
            _ = v37
        }
    }
    if true {
        var v38 uint64 = 50
        v38 = (v38 | 0)
        _ = v38
    }
    v39 := GenS4{F0: 0, F1: ("" + (("foo bar" + "hello") + ("foo bar" + (("\n\t" + (("\n\t" + ("\n\t" + "hello")) + "foo bar")) + "a"))))}
    var v40 bool = ((false && (true && (786 > 255))) && true)
    var v41 uint = 2
    fmt.Println(v41, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
