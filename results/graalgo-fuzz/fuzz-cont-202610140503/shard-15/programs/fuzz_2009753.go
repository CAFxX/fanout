package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 string }
type GenS2 struct { F0 int8; F1 int32; F2 int }
type GenS3 struct { F0 string; F1 uint64 }
func genFn0(p0 int) uint8 {
    switch "a" {
    case "foo bar":
        p0 = (p0 * -1)
    case "":
        fmt.Println(p0)
    }
    if true {
        fmt.Println(p0)
    } else {
        var v5 uint32 = 4294967295
        _ = v5
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(p0)
        if true {
            fmt.Println(fi6)
        }
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    if (2 == 1) {
        v8 := GenS1{F0: true, F1: "a", F2: ("foo bar" + ("foo bar" + "a"))}
        v9 := GenS1{F0: false, F1: ("\n\t" + "foo bar"), F2: "hello"}
        var v10 int64 = -9223372036854775808
        _ = v10
        _ = v8
        _ = v9
    } else {
        p0 -= (p0 - p0)
        fmt.Println(p0)
    }
    _ = p0
    return 36
}

func main() {
    var v1 float32 = (((1e10 - (float32((1e10 - ((3.14159 + (0.0 / (float64((-2.25 * 1.5)) + -2.25))) * -2.25))) * -2.25)) / 1.5) - (1.5 + 0.0))
    var v2 int32 = -1
    v3 := make(map[uint32]string)
    var v4 float64 = float64(((((1.5 + 3.14159) / v1) * v1) / v1))
    switch 42 {
    case 458:
        v11 := GenS1{F0: ((103 < 0) || false), F1: "hello", F2: ("a" + (("foo bar" + (("a" + ((((((("" + "foo bar") + ("" + "")) + "") + "hello") + "foo bar") + ("hello" + "")) + ("\n\t" + ""))) + ("a" + "a"))) + "a"))}
        _ = v11
    case 2147483647:
        if (!(-9223372036854775808 < -9223372036854775808)) {
            var v12 int = -2147483648
            var v13 int = (v12 / ((-193) | 1))
            _ = v12
            _ = v13
        }
    }
    switch "foo bar" {
    case "":
        var v14 uint64 = 0
        _ = v14
    case "a":
        v1 += v1
    }
    var v15 int32 = v2
    v16 := []int{32768, 0, 184}
    v15 = (v15 * -1)
    v17 := GenS1{F0: (2147483647 != 2), F1: "hello", F2: ((((("" + "hello") + (("a" + "hello") + ("foo bar" + "a"))) + ("foo bar" + ("\n\t" + "hello"))) + (("\n\t" + "") + "")) + "a")}
    v4 = -2.25
    v4 += v4
    if false {
        fmt.Println(v4, v15, v4)
    }
    switch 32768 {
    case 670:
        fmt.Println(v15, v1)
    case 32767:
        var v18 int8 = 76
        _ = v18
    }
    v15 -= v15
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 uint = 910
        v20 += (v20 ^ (v20 >> (((v20 ^ (v20 + 255))) & 63)))
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: -87, F1: (v2 >> ((42) & 63)), F2: -156}
    v15 += 0
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v2)
        _ = fi22
    }
    switch "a" {
    case "foo bar":
        if true {
            v23 := GenS1{F0: (-1 <= 85), F1: "a", F2: "a"}
            v24 := GenS3{F0: "a", F1: 0}
            fmt.Println(v1, v15)
            _ = v23
            _ = v24
        } else {
            if true {
                v1 -= 0.1
            }
            for fi25 := 0; fi25 < 14; fi25++ {
                if true {
                    var v26 bool = (532 <= 42)
                    _ = len(v3)
                    _ = v26
                }
                if ("\n\t" != (("foo bar" + "a") + "\n\t")) {
                    v27 := GenS1{F0: (!true), F1: ("\n\t" + ("foo bar" + (("a" + "a") + ""))), F2: "\n\t"}
                    _ = v27
                }
                for fi28 := 0; fi28 < 5; fi28++ {
                    var v29 uint = 121
                    _ = fi28
                    _ = v29
                }
                _ = fi25
            }
        }
    case "\n\t":
        if (2 > 32767) {
            v3[2147483648] = ((((("a" + "hello") + "\n\t") + "a") + "a") + "")
            v30 := GenS3{F0: "hello", F1: 2}
            var v31 float64 = v4
            _ = v30
            _ = v31
        }
    case "hello":
        v2 += v15
    }
    var v32 uint8 = 128
    switch "hello" {
    case "foo bar":
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v4
}
