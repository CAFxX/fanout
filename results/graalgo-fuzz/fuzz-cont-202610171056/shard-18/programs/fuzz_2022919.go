package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 uint16; F1 bool; F2 int64 }
func genFn0(p0 uint16, p1 int16, p2 uint8) int16 {
    v5 := GenS1{F0: (p0 - p0), F1: -369, F2: p1}
    switch 9223372036854775807 {
    case 561:
        for fi6 := 0; fi6 < 16; fi6++ {
            p0 -= (p0 | p0)
            fmt.Println(fi6, p0, p2)
            _ = fi6
        }
    }
    v7 := GenS2{F0: 2147483647}
    var v8 uint32 = 4294967295
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int8 = 100
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    _ = v8
    return (p1 | (p1 % ((p1) | 1)))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 1
    v3 := make(map[string]bool)
    var v4 int32 = -655
    if false {
        fmt.Println(v2, v4)
        var v11 string = "foo bar"
        _ = v11
    } else {
        switch ("a" + "a") {
        case "a":
            v12 := GenS3{F0: 1}
            _ = v12
        case "hello":
            if (692 <= 4294967295) {
                var v13 int64 = 9223372036854775807
                var v14 int16 = -323
                var v15 uint16 = 42
                _ = v13
                _ = v14
                _ = v15
            } else {
                v16 := GenS2{F0: 42}
                _ = v16
            }
        case "\n\t":
            var v17 int16 = 42
            _ = v17
        }
    }
    v4 = (v4 & v4)
    var v18 float64 = 0.0
    if false {
        var v19 int32 = 1
        v20 := GenS4{F0: 1000}
        v21 := GenS4{F0: 579}
        _ = v19
        _ = v20
        _ = v21
    }
    v2 -= v2
    if ((875 < -1) && (18446744073709551615 != 0)) {
        var v22 int64 = -9223372036854775808
        _ = v22
    }
    fmt.Println(v4, v18, v18)
    v23 := []int{-9223372036854775808, 127, 2147483647}
    for fi24 := 0; fi24 < 16; fi24++ {
        switch (("hello" + ((((("foo bar" + ("foo bar" + "\n\t")) + "\n\t") + ((("" + "hello") + "") + "a")) + ("a" + "")) + (((((("a" + "\n\t") + (((("a" + ("" + "foo bar")) + "\n\t") + ("" + "foo bar")) + "hello")) + "\n\t") + ("\n\t" + ((("hello" + (((("\n\t" + "foo bar") + "foo bar") + (("\n\t" + "hello") + "\n\t")) + "foo bar")) + ((("\n\t" + "\n\t") + ("a" + ("foo bar" + ("a" + "")))) + "a")) + "foo bar"))) + "hello") + "hello"))) + ((("" + "foo bar") + "foo bar") + ("hello" + "a"))) {
        case "foo bar":
            v25 := GenS5{F0: 727}
            _ = v25
        case "hello":
            if ((((3.14159 + 3.14159) == (1e10 / (((float32(float64((((0.0 / float32(v18)) - 0.1) + float32(v18)))) * 1.5) + 1.5) * (3.14159 * 0.1)))) && true) && (((0 <= genFn0(v2, 514, 190)) && (true || (!false))) && (false && true))) {
                v26 := GenS5{F0: 4294967295}
                if (!true) {
                    v4 = v4
                    for fi27 := 0; fi27 < 2; fi27++ {
                        var v28 uint = 913
                        var v29 string = "\n\t"
                        if (!false) {
                            v30 := GenS6{F0: v28, F1: v28}
                            var v31 float64 = v18
                            delete(v1, 65535)
                            _ = v30
                            _ = v31
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                } else {
                    var v32 uint = 218
                    _ = v32
                }
                _ = v26
            } else {
                v33 := GenS7{F0: (v2 | v2), F1: true, F2: -1}
                _ = v33
            }
        }
        v18 -= (1e10 / v18)
        if (1 == -1) {
            fmt.Println(v4)
            fmt.Println(v4, v18, fi24)
            v34 := GenS1{F0: 186, F1: -601, F2: 867}
            _ = v34
        }
        _ = fi24
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        var v36 uint64 = 683
        var v37 uint16 = (v2 ^ v2)
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
