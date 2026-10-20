package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint64 }
type GenS2 struct { F0 int16 }
func genFn0() int16 {
    var v7 uint = 4294967295
    switch -96 {
    case 128:
        v8 := GenS1{F0: "\n\t", F1: ("foo bar" + ("foo bar" + ("hello" + "foo bar"))), F2: 42}
        _ = v8
    default:
        v9 := []int{-1}
        _ = v9
    }
    switch ("foo bar" + ("foo bar" + "hello")) {
    case "\n\t":
        v10 := []int{127, -743, 0, -2147483648}
        _ = v10
    case "":
        fmt.Println(v7)
    case "a":
        var v11 int32 = 1
        _ = v11
    default:
        var v12 uint32 = 1
        _ = v12
    }
    fmt.Println(v7)
    var v13 uint8 = 128
    _ = v13
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -9223372036854775808
    v3 := make(map[string]string)
    var v4 uint64 = 199
    v5 := make(map[uint32]string)
    v6 := []uint32{220, 567, 2147483647}
    v2 = v2
    delete(v3, "\n\t")
    v2 = -1
    delete(v3, ("" + "\n\t"))
    fmt.Println(v4)
    switch v2 {
    case -994:
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS1{F0: "foo bar", F1: (("hello" + "hello") + ("a" + "foo bar")), F2: v4}
            v4 = (v4 | 4294967295)
            _ = fi14
            _ = v15
        }
    default:
        fmt.Println(v2, v2)
    }
    fmt.Println(v4)
    v2 += (v2 >> ((42) & 63))
    var v16 int16 = 918
    switch 0 {
    case 1:
        var v17 uint64 = (^(v4))
        _ = v17
    case 32767:
        _ = len(v6)
    case -802:
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
    }
    if (true || (((1e10 >= 0.0) && (!("" != ""))) && false)) {
        if false {
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 int8 = -82
                v20 := GenS2{F0: int16(v19)}
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 string = "hello"
            var v22 int16 = (-(v16))
            _ = v21
            _ = v22
        }
        v23 := GenS1{F0: "", F1: "\n\t", F2: v4}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 4; fi24++ {
            switch "" {
            case "":
                var v25 uint64 = 2
                _ = v25
            default:
                v2 = (v2 / ((fi24) | 1))
            }
            _ = fi24
        }
    }
    v5[4294967295] = ((("\n\t" + "\n\t") + "foo bar") + "\n\t")
    var v26 int32 = 42
    v27 := GenS2{F0: v16}
    v28 := GenS1{F0: ("hello" + ("a" + ((("a" + "hello") + "foo bar") + (("hello" + ("\n\t" + ("hello" + "a"))) + ("\n\t" + ("foo bar" + "a")))))), F1: "foo bar", F2: (^(v4))}
    if false {
        for fi29 := 0; fi29 < 4; fi29++ {
            v16 -= (v16 >> ((255) & 63))
            _ = fi29
        }
        delete(v5, 1)
    } else {
        var v30 int = int(v4)
        switch v2 {
        case -1:
            switch ("" + "hello") {
            case "\n\t":
                var v31 float32 = 1e10
                _ = v31
            }
        case -2147483648:
            v32 := GenS1{F0: "foo bar", F1: "a", F2: 1}
            _ = v32
        case 127:
            v33 := GenS1{F0: "\n\t", F1: "", F2: (v4 + v4)}
            _ = v33
        default:
            var v34 int16 = 1
            _ = v34
        }
        var v35 int16 = genFn0()
        _ = v30
        _ = v35
    }
    var v36 uint32 = 3
    var v37 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
