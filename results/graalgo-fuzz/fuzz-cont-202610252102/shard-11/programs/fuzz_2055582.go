package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 int }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string }
func genFn0() string {
    switch 2147483647 {
    case 127:
        switch 32767 {
        case 0:
            var v4 uint32 = 0
            _ = v4
        }
    case 4294967295:
        for fi5 := 0; fi5 < 12; fi5++ {
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 int16 = 1000
                var v8 uint16 = 348
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi5
        }
    case 2147483647:
        var v9 uint = 42
        _ = v9
    }
    var v10 float64 = 1.5
    v10 = (v10 - v10)
    _ = v10
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 uint64 = 18446744073709551615
    var v11 bool = true
    v3 += v3
    v12 := GenS1{F0: 32768, F1: 65535}
    v11 = (42 > 127)
    v13 := v1[2147483648]
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            if (float32(0.0) >= 1e10) {
                var v16 float64 = 1.5
                fmt.Println(v16, fi14, v13)
                v13 = (!v11)
                if ((("a" + ("a" + "")) + "hello") >= ("a" + ("a" + ("" + "a")))) {
                    v11 = (true && v13)
                    v17 := GenS2{F0: 179, F1: "a", F2: (fi14 ^ (fi15 >> ((2) & 63)))}
                    fmt.Println(v11)
                    v11 = (v3 != (^(v3)))
                    _ = v17
                }
                v18 := GenS3{F0: 0, F1: 4294967295}
                _ = v16
                _ = v18
            }
            fmt.Println(fi15, fi14)
            _ = fi15
        }
        switch -2147483648 {
        case 0:
            v11 = v13
        }
        if true {
            v19 := GenS4{F0: 226}
            v11 = (((3.14159 + -2.25) >= 3.14159) || v13)
            _ = v19
        }
        _ = fi14
    }
    if v11 {
        switch "foo bar" {
        case "a":
            v13 = (2 > 994)
        case "hello":
            switch "a" {
            case "hello":
                var v20 string = ("a" + ("foo bar" + "hello"))
                _ = v20
            case "foo bar":
                v13 = (4294967295 > 18446744073709551615)
            default:
                v3 += (v3 * (v3 + 42))
            }
        case "\n\t":
            v3 = (v3 / ((801) | 1))
        }
        switch (("\n\t" + ("hello" + "\n\t")) + "") {
        case "":
            v21 := GenS3{F0: 484, F1: 18446744073709551615}
            _ = v21
        case "a":
            v22 := GenS5{F0: 1}
            _ = v22
        case "hello":
            var v23 float32 = (0.1 * (float32(0.1) - 0.0))
            _ = v23
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            if true {
                v13 = (v13 || v13)
                v25 := []int{(fi24 / (((fi24 & 9223372036854775807)) | 1))}
                _ = v25
            } else {
                if v11 {
                    var v26 uint16 = 42
                    for fi27 := 0; fi27 < 3; fi27++ {
                        fmt.Println(v26, v11)
                        v28 := GenS4{F0: (^(v26))}
                        _ = fi27
                        _ = v28
                    }
                    v11 = v11
                    _ = v26
                } else {
                    v29 := GenS6{F0: 42}
                    _ = v29
                }
            }
            _ = fi24
        }
    }
    var v30 float32 = (1.5 + 0.0)
    fmt.Println(v11, v11)
    var v31 uint8 = 255
    for fi32 := 0; fi32 < 1; fi32++ {
        v13 = v11
        _ = fi32
    }
    var v33 int = 9223372036854775807
    if ((true || (v11 || (v31 != (v31 | v31)))) || v11) {
        v34 := GenS7{F0: (("\n\t" + ("" + ("" + "a"))) + "a")}
        _ = v34
    } else {
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 uint16 = 2
            fmt.Println(v33, v3)
            _ = fi35
            _ = v36
        }
        var v37 int64 = -665
        _ = v37
    }
    fmt.Println(v30)
    fmt.Println(v31, v31, v13)
    v38 := []int{v33, (-(v33))}
    v33 += 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v38
}
