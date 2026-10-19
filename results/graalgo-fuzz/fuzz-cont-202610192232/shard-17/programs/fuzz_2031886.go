package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint64; F1 int64 }

func main() {
    var v1 uint16 = 646
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    v4 := GenS1{F0: "", F1: -213, F2: "\n\t"}
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v5 int32 = 2147483647
        _ = v5
    case 127:
        for fi6 := 0; fi6 < 14; fi6++ {
            switch ("a" + "foo bar") {
            case "hello":
                if false {
                    var v7 uint16 = (^(v1))
                    var v8 int8 = 0
                    _ = v7
                    _ = v8
                } else {
                    v1 += 48
                }
            case "\n\t":
                v1 -= v1
            }
            _ = fi6
        }
    default:
        _ = len(v3)
    }
    v9 := []int{1, 42, -2147483648, 0}
    var v10 uint = 245
    fmt.Println(v1, v10)
    for fi11 := 0; fi11 < 7; fi11++ {
        switch fi11 {
        case 1:
            var v12 int64 = -1
            _ = v12
        case 32768:
            switch (fi11 & fi11) {
            case -1:
                v1 -= 314
            case 1:
                v10 = (v10 >> ((v10) & 63))
            default:
                var v13 int32 = -177
                _ = v13
            }
        case 42:
            fmt.Println(v1)
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS1{F0: "foo bar", F1: -157, F2: ((("hello" + "a") + (("hello" + ("\n\t" + (("foo bar" + ("\n\t" + "a")) + ("\n\t" + (("a" + "a") + ""))))) + "")) + "a")}
            var v16 uint8 = 0
            if ((false || true) && (v10 == (v10 >> ((2) & 63)))) {
                var v17 int16 = -32768
                v10 = 753
                var v18 uint = (^(v10))
                _ = v17
                _ = v18
            } else {
                if false {
                    var v19 float64 = -2.25
                    v20 := GenS1{F0: ("\n\t" + "a"), F1: 9223372036854775807, F2: "\n\t"}
                    _ = v19
                    _ = v20
                }
                fmt.Println(v1, fi14, v16)
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi11
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        v22 := GenS1{F0: "hello", F1: 0, F2: "\n\t"}
        v23 := GenS2{F0: (v1 & uint16(v10))}
        v24 := GenS1{F0: "hello", F1: -9223372036854775808, F2: ""}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    var v25 float32 = (1.5 / 0.1)
    switch (("" + "hello") + ("hello" + "")) {
    case "\n\t":
        for fi26 := 0; fi26 < 19; fi26++ {
            fmt.Println(fi26, v1, v25)
            fmt.Println(v10, fi26)
            for fi27 := 0; fi27 < 11; fi27++ {
                var v28 bool = ((false && (false || (!(0 >= 42)))) && (false && false))
                if ((^(fi27)) < fi26) {
                    v10 = 1
                    _ = len(v3)
                    var v29 uint = (v10 % ((v10) | 1))
                    var v30 uint8 = 2
                    fmt.Println(v29)
                    v31 := GenS3{F0: (^(v30)), F1: -32768}
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    case "foo bar":
        var v32 string = "hello"
        _ = v32
    }
    fmt.Println(v10, v1)
    for fi33 := 0; fi33 < 18; fi33++ {
        v34 := GenS4{F0: 1, F1: 1}
        _ = fi33
        _ = v34
    }
    var v35 float64 = 1.5
    v36 := GenS3{F0: 2, F1: 1000}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v25 = (((0.0 + (v25 / v25)) - v25) / v25)
    if len(v9) > 0 {
        v9[0] = 1
    }
    if false {
        v1 = (v1 << (((v10 * v10)) & 63))
        var v37 int64 = 0
        v35 = (v35 / ((v35 / v35) / 1.0))
        _ = v37
    }
    var v38 int = -167
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v9
}
