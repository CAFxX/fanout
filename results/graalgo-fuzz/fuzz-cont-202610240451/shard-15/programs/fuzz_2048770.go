package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int; F1 uint16; F2 int32 }

func main() {
    var v1 uint64 = 0
    var v2 int64 = 9223372036854775807
    v3 := []string{("hello" + "foo bar"), ((("" + "hello") + "\n\t") + "foo bar"), ("\n\t" + (("hello" + ("\n\t" + ((("" + ("\n\t" + (("foo bar" + "foo bar") + ("" + "\n\t")))) + ("\n\t" + "a")) + ("hello" + "\n\t")))) + "hello"))}
    var v4 float64 = (-2.25 / 3.14159)
    v5 := GenS1{F0: "a", F1: -1}
    v1 += v1
    var v6 int64 = (v2 / ((v2) | 1))
    switch "hello" {
    case "":
        v4 = ((v4 - -2.25) - (float64(float32((0.1 / 1.5))) + v4))
    }
    v7 := GenS1{F0: "hello", F1: 1}
    v2 -= v2
    var v8 uint32 = 32
    switch 830 {
    case 32768:
        v1 = v1
    default:
        if true {
            var v9 int16 = 0
            _ = v9
        }
    }
    if (int32(v6) == -1) {
        var v10 float32 = (0.0 / float32(1.5))
        for fi11 := 0; fi11 < 1; fi11++ {
            _ = len(v3)
            _ = fi11
        }
        _ = v10
    } else {
        v12 := GenS2{F0: 1, F1: -128}
        _ = v12
    }
    v13 := GenS1{F0: "a", F1: 179}
    if false {
        _ = len(v3)
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            switch (fi14 / (((fi14 - (fi14 >> ((2) & 63)))) | 1)) {
            case 32767:
                switch "\n\t" {
                case "\n\t":
                    v15 := GenS2{F0: 519, F1: 72}
                    _ = v15
                case "":
                    var v16 float32 = float32(3.14159)
                    _ = v16
                case "hello":
                    v17 := GenS2{F0: 922, F1: 42}
                    _ = v17
                }
            case 127:
                var v18 string = ("" + ((("foo bar" + (("hello" + (("\n\t" + "hello") + "\n\t")) + ("hello" + ("\n\t" + "hello")))) + "") + "\n\t"))
                _ = v18
            default:
                var v19 float64 = v4
                _ = v19
            }
            fmt.Println(v8, v8)
            v2 = -738
            _ = fi14
        }
        var v20 uint32 = 809
        _ = v20
    }
    if (-27 != 100) {
        switch "hello" {
        case "\n\t":
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 uint8 = 1
                _ = fi21
                _ = v22
            }
        case "hello":
            for fi23 := 0; fi23 < 3; fi23++ {
                v24 := GenS3{F0: -128}
                var v25 int16 = 42
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    } else {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        if (!(65535 >= 4294967295)) {
            switch "foo bar" {
            case "":
                v6 = v6
            case "hello":
                v3 = append(v3, "foo bar")
            }
            _ = len(v3)
            v3 = append(v3, "a")
        }
    }
    v2 += v6
    var v26 int8 = -128
    for fi27 := 0; fi27 < 3; fi27++ {
        switch "foo bar" {
        case "a":
            var v28 uint16 = 0
            _ = v28
        default:
            fmt.Println(v4)
        }
        _ = fi27
    }
    var v29 int = 726
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 7; fi31++ {
            v3 = append(v3, ("a" + ("" + "foo bar")))
            _ = fi31
        }
        var v32 float32 = 0.1
        _ = fi30
        _ = v32
    }
    if (((float32(3.14159) * 1e10) / 3.14159) == 1.5) {
        v33 := GenS4{F0: (v8 ^ 65535), F1: uint16(v1), F2: (v1 & v1)}
        for fi34 := 0; fi34 < 12; fi34++ {
            fmt.Println(v2, v6)
            var v35 float32 = ((-2.25 - (0.0 / (0.0 - ((float32(((((v4 * (v4 * v4)) + v4) + 0.0) * 3.14159)) * (1e10 / 0.1)) - 1.5)))) / 3.14159)
            v36 := GenS5{F0: 32768, F1: 716, F2: 2147483647}
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v33
    } else {
        var v37 int16 = -1
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
