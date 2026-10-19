package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int64; F1 bool; F2 uint64 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 string) int16 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint8 = 42
        v9 := GenS1{F0: (v8 << ((387) & 63)), F1: 251}
        switch 1 {
        case 0:
            p0 = ("a" + ("" + (p0 + "hello")))
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        p0 = ("" + "hello")
    } else {
        p0 = p0
    }
    var v10 bool = (!false)
    var v11 int = -1
    _ = p0
    _ = v10
    _ = v11
    return -1
}

func main() {
    var v1 int32 = 42
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 int8 = -52
    v5 := make(map[string]bool)
    v6 := []string{"foo bar", "hello", (("" + (("hello" + "\n\t") + "hello")) + ("foo bar" + (("hello" + "foo bar") + "hello"))), "\n\t"}
    v5[""] = (!(false || true))
    var v12 int8 = (v4 + (v4 - v4))
    for fi13 := 0; fi13 < 9; fi13++ {
        v4 = v4
        var v14 uint64 = 424
        _ = fi13
        _ = v14
    }
    switch 1 {
    case 9223372036854775807:
        var v15 uint32 = 42
        _ = v15
    case 4294967295:
        switch 33 {
        case 42:
            v6 = append(v6, ("a" + (((("foo bar" + "hello") + ("hello" + "\n\t")) + "a") + "hello")))
        case 780:
            for fi16 := 0; fi16 < 8; fi16++ {
                v17 := GenS1{F0: 128, F1: -1}
                v18 := GenS2{F0: 2147483647, F1: -9223372036854775808}
                _ = fi16
                _ = v17
                _ = v18
            }
        case 155:
            _ = len(v5)
        default:
            v19 := GenS3{F0: 112, F1: 65535}
            _ = v19
        }
    case -1:
        switch ("hello" + "hello") {
        case "hello":
            v20 := GenS4{F0: 42, F1: (false && (128 < 32768)), F2: 42}
            _ = v20
        default:
            var v21 uint32 = 685
            _ = v21
        }
    default:
        switch 32768 {
        case -1:
            if (false && (65535 >= 1)) {
                if len(v6) > 0 {
                    v6[0] = ("a" + "a")
                }
            }
        case -2147483648:
            v5["foo bar"] = false
        case 128:
            if ((!(false || false)) || false) {
                v12 -= v12
                v4 = v4
            }
        default:
            var v22 uint64 = 780
            _ = v22
        }
    }
    v23 := GenS2{F0: 0, F1: 601}
    v24 := GenS2{F0: v1, F1: 306}
    v25 := GenS2{F0: (v1 >> ((255) & 63)), F1: 478}
    v26 := GenS5{F0: -9223372036854775808, F1: -9223372036854775808}
    v27 := GenS6{F0: 42}
    switch -441 {
    case 2147483647:
        var v28 uint32 = 2147483648
        _ = v28
    case 42:
        for fi29 := 0; fi29 < 15; fi29++ {
            fmt.Println(v1, v12)
            fmt.Println(v1, v12, fi29)
            _ = fi29
        }
    }
    var v30 string = (("a" + "") + (("" + "hello") + ("foo bar" + "hello")))
    fmt.Println(v12, v4)
    fmt.Println(v12, v4, v12)
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 float32 = (float32((float64(-2.25) - 0.0)) - -2.25)
        _ = fi31
        _ = v32
    }
    _ = len(v6)
    switch "" {
    case "a":
        v2["hello"] = 42
    default:
        for fi33 := 0; fi33 < 1; fi33++ {
            v4 = v12
            _ = fi33
        }
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        if len(v6) > 0 {
            v6[0] = ""
        }
        var v35 int8 = 42
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
