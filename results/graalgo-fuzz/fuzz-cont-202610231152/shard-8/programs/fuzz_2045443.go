package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 int16; F1 uint; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 698
    var v3 float64 = 0.1
    var v4 uint16 = 42
    if (false || true) {
        switch -2147483648 {
        case 32768:
            if (-2.25 != ((0.1 - 0.0) + 1.5)) {
                var v5 int64 = 1
                _ = v5
            }
        default:
            var v6 int64 = 1
            _ = v6
        }
        var v7 int8 = 42
        _ = v7
    } else {
        v8 := GenS1{F0: 0, F1: ((("foo bar" + ("" + "")) + ("a" + ("" + ("\n\t" + "\n\t")))) + "hello"), F2: 103}
        var v9 int16 = 166
        v10 := GenS2{F0: 0, F1: 944, F2: -4}
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        v2 = 1
        _ = fi11
    }
    v12 := GenS1{F0: 32767, F1: ("" + (("" + (((("foo bar" + "a") + "\n\t") + "\n\t") + "\n\t")) + "\n\t")), F2: 100}
    fmt.Println(v2, v2, v4)
    var v13 uint8 = 1
    for fi14 := 0; fi14 < 12; fi14++ {
        if true {
            fmt.Println(v4, v13, fi14)
            v15 := []int{fi14, fi14, (-(fi14))}
            _ = v15
        } else {
            var v16 bool = false
            _ = v16
        }
        _ = fi14
    }
    if (-1 < 0) {
        v13 = (v13 | v13)
    }
    switch "a" {
    case "hello":
        var v17 int64 = 403
        _ = v17
    case "\n\t":
        for fi18 := 0; fi18 < 7; fi18++ {
            v13 = 128
            if true {
                var v19 int8 = 91
                _ = v19
            } else {
                if true {
                    fmt.Println(v13, v3, v3)
                    v4 -= v4
                    var v20 uint = 42
                    _ = v20
                }
            }
            _ = fi18
        }
    }
    switch "hello" {
    case "":
        switch 32768 {
        case -708:
            v4 -= 42
        default:
            v21 := GenS1{F0: 1000, F1: "foo bar", F2: -128}
            _ = v21
        }
    case "foo bar":
        v4 = 769
    case "\n\t":
        delete(v1, "foo bar")
    default:
        v22 := GenS3{F0: 42, F1: v4}
        _ = v22
    }
    switch "foo bar" {
    case "foo bar":
        var v23 float32 = (0.0 / float32(v3))
        _ = v23
    default:
        v24 := GenS2{F0: 32767, F1: 4294967295, F2: -341}
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
