package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int64 }
func genFn0(p0 int32, p1 int64) int8 {
    if true {
        fmt.Println(p1)
    } else {
        if ("a" != ("a" + "\n\t")) {
            v5 := GenS1{F0: "", F1: 862, F2: 65535}
            fmt.Println(p1, p0)
            _ = v5
        }
        var v6 float32 = 0.1
        switch "foo bar" {
        case "":
            var v7 uint16 = 255
            _ = v7
        default:
            v8 := GenS1{F0: ("" + "a"), F1: 9223372036854775807, F2: 0}
            _ = v8
        }
        _ = v6
    }
    v9 := []int{0}
    if (3.14159 <= float64(1.5)) {
        v10 := GenS1{F0: "", F1: 661, F2: 256}
        _ = v10
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v9
    return 72
}

func main() {
    var v1 int16 = 42
    v2 := []string{"a"}
    v3 := make(map[int]uint32)
    var v4 int8 = 0
    var v11 uint16 = 374
    var v12 int64 = 0
    _ = len(v3)
    switch 2147483647 {
    case -233:
        v13 := GenS1{F0: ("" + "a"), F1: 42, F2: v11}
        _ = v13
    case 127:
        v11 -= v11
    }
    v12 += -1
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS2{F0: (genFn0(-592, int64(fi14)) < (v4 | v4)), F1: (true || (true && (-128 <= v4))), F2: v12}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        if false {
            fmt.Println(v1, v4)
            for fi17 := 0; fi17 < 12; fi17++ {
                v18 := GenS3{F0: (fi16 | fi17)}
                _ = fi17
                _ = v18
            }
        }
        _ = fi16
    }
    v11 = (^(v11))
    fmt.Println(v11, v4)
    v4 = (v4 + (v4 >> ((4294967295) & 63)))
    v1 -= v1
    switch ("hello" + "") {
    case "\n\t":
        switch ("" + "foo bar") {
        case "\n\t":
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 uint16 = v11
                for fi21 := 0; fi21 < 6; fi21++ {
                    if false {
                        v12 += 0
                        switch "a" {
                        case "foo bar":
                            var v22 int32 = 906
                            _ = v22
                        }
                    } else {
                        fmt.Println(v4)
                    }
                    fmt.Println(v12, v1)
                    switch (^(fi19)) {
                    case 128:
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    default:
                        switch "a" {
                        case "":
                            v23 := v3[(^(fi21))]
                            _ = v23
                        case "hello":
                            fmt.Println(v12, v12, v11)
                        }
                    }
                    _ = fi21
                }
                _ = fi19
                _ = v20
            }
        case "foo bar":
            for fi24 := 0; fi24 < 13; fi24++ {
                v25 := GenS3{F0: 497}
                _ = fi24
                _ = v25
            }
        }
    case "":
        for fi26 := 0; fi26 < 18; fi26++ {
            _ = len(v3)
            var v27 int64 = v12
            _ = fi26
            _ = v27
        }
    case "hello":
        v28 := GenS4{F0: false, F1: (v12 * (-(v12)))}
        _ = v28
    }
    switch 42 {
    case 443:
        fmt.Println(v11, v11, v11)
    default:
        var v29 uint = 901
        _ = v29
    }
    fmt.Println(v12)
    fmt.Println(v12, v12, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
