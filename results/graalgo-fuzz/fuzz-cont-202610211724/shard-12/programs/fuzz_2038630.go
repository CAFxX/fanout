package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 string; F2 int }
type GenS3 struct { F0 int8; F1 int32; F2 int16 }
func genFn0(p0 uint32) bool {
    p0 = (p0 & p0)
    var v4 int64 = -9223372036854775808
    _ = p0
    _ = v4
    return (((3.14159 == 1.5) && ((65535 >= 153) || (false && (!false)))) && (0 < 42))
}

func main() {
    v1 := []uint32{2147483647, 2147483648}
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    var v5 int16 = 42
    var v6 bool = true
    switch "foo bar" {
    case "hello":
        v6 = false
    }
    if (!(4294967295 <= -336)) {
        if v6 {
            v7 := GenS1{F0: 116}
            _ = v7
        }
        fmt.Println(v5, v6)
        switch 32768 {
        case 4294967295:
            v3[-618] = 65535
        }
    }
    var v8 float32 = ((((0.0 * (1e10 * ((1e10 / (0.1 / (1.5 * ((((0.1 + 0.1) + 3.14159) * 1e10) / 0.1)))) - 0.0))) * -2.25) / 3.14159) + 0.0)
    var v9 uint16 = 32767
    switch ("a" + "foo bar") {
    case "":
        for fi10 := 0; fi10 < 5; fi10++ {
            switch fi10 {
            case 128:
                switch ("\n\t" + "") {
                case "":
                    var v11 bool = v6
                    _ = v11
                case "hello":
                    var v12 string = "a"
                    _ = v12
                default:
                    if (!v6) {
                        var v13 uint = 255
                        var v14 int8 = 0
                        _ = v13
                        _ = v14
                    }
                }
            case 647:
                var v15 int8 = 18
                _ = v15
            default:
                var v16 int8 = 53
                _ = v16
            }
            switch "a" {
            case "\n\t":
                if v6 {
                    var v17 int = (^(fi10))
                    v18 := GenS1{F0: 18446744073709551615}
                    v8 = (-2.25 / (float32(0.1) * (v8 - ((3.14159 - v8) - v8))))
                    _ = v17
                    _ = v18
                }
            case "hello":
                var v19 int64 = 9223372036854775807
                _ = v19
            }
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := GenS1{F0: 1}
                _ = fi20
                _ = v21
            }
            _ = fi10
        }
    case "foo bar":
        v9 += v9
    }
    switch 128 {
    case 1:
        v22 := GenS2{F0: 2, F1: ("a" + "\n\t"), F2: -376}
        _ = v22
    case -618:
        fmt.Println(v8, v9, v5)
    case -441:
        v5 = (^(v5))
    }
    v6 = (45 > 127)
    var v23 int64 = -9223372036854775808
    v24 := GenS1{F0: 255}
    v25 := GenS2{F0: 421, F1: "", F2: 1}
    for fi26 := 0; fi26 < 5; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            v28 := GenS3{F0: -128, F1: 1, F2: v5}
            var v29 int16 = (v5 << ((696) & 63))
            var v30 string = "foo bar"
            fmt.Println(v5, v8)
            v3[fi27] = 0
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        v8 += v8
        switch ("" + "") {
        case "hello":
            if ("a" <= "a") {
                if v6 {
                    var v31 uint64 = 2
                    v6 = v6
                    var v32 int16 = v5
                    _ = v31
                    _ = v32
                } else {
                    switch (("hello" + ("a" + "hello")) + "a") {
                    case "":
                        v5 += (v5 & (v5 % ((-483) | 1)))
                    default:
                        var v33 int16 = v5
                        _ = v33
                    }
                }
            } else {
                var v34 uint32 = 2147483648
                _ = v34
            }
        case "\n\t":
            v8 = v8
        default:
            var v35 uint16 = v9
            _ = v35
        }
        _ = fi26
    }
    var v36 int16 = v5
    v36 = v5
    switch "foo bar" {
    case "a":
        var v37 float64 = (-2.25 + ((1.5 / ((1.5 + (-2.25 - 1e10)) / -2.25)) - 3.14159))
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        fmt.Println(v8, v9, fi38)
        _ = fi38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
