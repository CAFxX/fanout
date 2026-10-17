package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 string) uint64 {
    var v5 float32 = -2.25
    if ((!((1 > -101) || (v5 >= (v5 * (-2.25 - v5))))) && (29 < 42)) {
        v6 := GenS1{F0: p0, F1: -470}
        v7 := GenS1{F0: "a", F1: 1}
        v8 := GenS1{F0: ("a" + p0), F1: 1}
        _ = v6
        _ = v7
        _ = v8
    } else {
        if (0 < -584) {
            for fi9 := 0; fi9 < 9; fi9++ {
                v5 -= 0.1
                _ = fi9
            }
        } else {
            var v10 uint = 1
            _ = v10
        }
    }
    if ((32767 != 1) && ((v5 == 0.1) && true)) {
        var v11 float64 = ((-2.25 / float64((v5 + float32((1.5 + 0.0))))) + -2.25)
        v12 := []int{}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v5
    return 869
}

func main() {
    var v1 uint32 = 34
    var v2 int8 = 127
    var v3 int64 = -859
    v4 := make(map[uint32]string)
    v13 := []int{}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v14 := GenS1{F0: "hello", F1: 42}
    var v15 int16 = 42
    if true {
        switch -2147483648 {
        case -617:
            _ = len(v4)
        case -279:
            v3 -= 250
        default:
            if len(v13) > 0 {
                v13[0] = 10
            }
        }
        if (42 == 32767) {
            var v16 uint16 = 0
            _ = v16
        }
        switch "hello" {
        case "a":
            v17 := GenS1{F0: "a", F1: v3}
            _ = v17
        case "\n\t":
            v2 += (v2 * (^(v2)))
        }
    } else {
        v4[v1] = "a"
    }
    if (true || true) {
        switch 842 {
        case 42:
            v1 -= 569
        default:
            v18 := GenS1{F0: "\n\t", F1: 0}
            _ = v18
        }
    }
    if (false || (v3 > v3)) {
        var v19 int8 = (^(v2))
        for fi20 := 0; fi20 < 2; fi20++ {
            v2 -= (v19 >> ((161) & 63))
            _ = fi20
        }
        _ = v19
    } else {
        var v21 int8 = v2
        v22 := GenS2{F0: (v1 << ((377) & 63))}
        if false {
            var v23 uint64 = genFn0((("" + ((((((("a" + "") + ("foo bar" + "hello")) + "foo bar") + "hello") + ("\n\t" + "hello")) + "foo bar") + "hello")) + "\n\t"))
            _ = v23
        } else {
            v24 := GenS2{F0: 1}
            fmt.Println(v15)
            _ = v24
        }
        _ = v21
        _ = v22
    }
    switch 2147483647 {
    case 32767:
        v25 := GenS2{F0: 0}
        _ = v25
    case 675:
        v26 := GenS3{F0: 4294967295}
        _ = v26
    default:
        for fi27 := 0; fi27 < 15; fi27++ {
            v2 = (v2 | 127)
            var v28 uint16 = 2
            _ = fi27
            _ = v28
        }
    }
    switch ("hello" + "a") {
    case "hello":
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 int32 = int32(v15)
            v31 := GenS2{F0: 0}
            _ = fi29
            _ = v30
            _ = v31
        }
    case "a":
        var v32 int8 = v2
        _ = v32
    case "\n\t":
        _ = len(v13)
    default:
        if ((false || ((42 < 1) && false)) && (!true)) {
            for fi33 := 0; fi33 < 4; fi33++ {
                v3 = 0
                fmt.Println(v1)
                var v34 uint8 = 43
                _ = fi33
                _ = v34
            }
        } else {
            if ((^(v2)) <= int8(v1)) {
                var v35 string = (("" + "foo bar") + "a")
                _ = v35
            }
            v3 -= (v3 | 42)
            v15 = (v15 >> ((42) & 63))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
