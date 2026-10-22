package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 int32; F1 int64 }
func genFn0(p0 int16) int {
    p0 += p0
    var v4 int16 = 1
    _ = p0
    _ = v4
    return -822
}

func main() {
    var v1 int64 = -1
    var v2 uint64 = 2
    v3 := []int{}
    v1 += v1
    var v5 uint8 = 127
    var v6 int64 = (v1 | (^(v1)))
    for fi7 := 0; fi7 < 2; fi7++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v5)
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v1, fi8)
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint = 255
            switch "hello" {
            case "":
                _ = len(v3)
            case "hello":
                v11 := GenS1{F0: true, F1: 2147483647, F2: (^(v1))}
                _ = v11
            }
            if true {
                switch genFn0(-1) {
                case -552:
                    if false {
                        var v12 string = "\n\t"
                        v10 -= v10
                        _ = v12
                    }
                default:
                    v3 = append(v3, fi8)
                }
                v6 += (-(v1))
                var v13 uint8 = (v5 + (v5 / ((0) | 1)))
                _ = v13
            } else {
                var v14 int32 = 2147483647
                var v15 int = fi9
                _ = v14
                _ = v15
            }
            if (!(!true)) {
                v6 = v6
                v16 := GenS2{F0: false}
                v17 := GenS3{F0: 42, F1: 42}
                _ = v16
                _ = v17
            }
            v3 = append(v3, -1)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if true {
        _ = len(v3)
        var v18 int8 = 100
        _ = v18
    }
    if len(v3) > 0 {
        v3[0] = genFn0(1)
    }
    switch "\n\t" {
    case "\n\t":
        switch (("foo bar" + ("a" + "foo bar")) + "foo bar") {
        case "":
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS4{F0: 42, F1: 32767, F2: v5}
                fmt.Println(v5, v5)
                _ = fi19
                _ = v20
            }
        case "foo bar":
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 int = 4294967295
                var v23 uint32 = uint32(v2)
                _ = fi21
                _ = v22
                _ = v23
            }
        case "hello":
            if len(v3) > 0 {
                v3[0] = genFn0(1)
            }
        }
    case "":
        if true {
            var v24 float64 = 3.14159
            _ = v24
        }
    default:
        fmt.Println(v2, v2, v2)
    }
    switch 42 {
    case 9223372036854775807:
        if (!true) {
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 uint32 = 2147483647
                var v27 int8 = 42
                _ = fi25
                _ = v26
                _ = v27
            }
        }
    }
    var v28 string = "foo bar"
    for fi29 := 0; fi29 < 6; fi29++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v30 float32 = ((((0.1 / 3.14159) * -2.25) - (3.14159 + (1.5 - float32(-2.25)))) - (float32(0.0) + 0.0))
        _ = fi29
        _ = v30
    }
    v31 := GenS5{F0: 1, F1: 115}
    switch ((v28 + v28) + v28) {
    case "a":
        v1 = v1
    case "\n\t":
        if true {
            switch v28 {
            case "a":
                v5 = (v5 + 42)
            default:
                if (1 > 42) {
                    var v32 int8 = 42
                    _ = v32
                } else {
                    fmt.Println(v2)
                    var v33 uint16 = 387
                    _ = v33
                }
            }
            v3 = append(v3, 759)
        } else {
            v2 = 4294967295
            fmt.Println(v5, v2)
        }
    }
    var v34 int64 = (v6 & -9223372036854775808)
    for fi35 := 0; fi35 < 3; fi35++ {
        if false {
            v3 = append(v3, -1)
            for fi36 := 0; fi36 < 3; fi36++ {
                v3 = append(v3, (fi36 ^ 0))
                _ = fi36
            }
        }
        switch (v28 + "\n\t") {
        case "":
            v28 = "a"
        case "\n\t":
            var v37 uint = 439
            _ = v37
        default:
            v38 := GenS6{F0: -1, F1: v6}
            _ = v38
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v34
    _ = v5
    _ = v6
}
