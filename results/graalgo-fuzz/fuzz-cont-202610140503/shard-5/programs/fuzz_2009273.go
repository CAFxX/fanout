package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 string; F1 uint; F2 int }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 int16; F1 int }
func genFn0(p0 string, p1 int8, p2 uint8) string {
    switch p0 {
    case "foo bar":
        if ("\n\t" <= p0) {
            v4 := GenS1{F0: 0, F1: 3}
            _ = v4
        }
    case "hello":
        p0 = "hello"
    }
    v5 := GenS1{F0: 128, F1: -737}
    var v6 float32 = float32((0.1 + 3.14159))
    var v7 string = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := []string{("\n\t" + ("hello" + ("\n\t" + "a"))), "hello"}
    var v2 string = ("foo bar" + "\n\t")
    var v3 int8 = -128
    fmt.Println(v3)
    switch 32767 {
    case -2147483648:
        _ = len(v1)
    case 4294967295:
        switch "hello" {
        case "foo bar":
            if true {
                var v8 uint64 = 509
                _ = v8
            } else {
                _ = len(v1)
                v3 = v3
            }
        case "\n\t":
            v3 = (v3 * v3)
        case "a":
            v3 = 42
        }
    case -378:
        var v9 int64 = 0
        _ = v9
    }
    v10 := GenS2{F0: v2, F1: 0}
    v1 = append(v1, ("\n\t" + v2))
    var v11 int16 = -133
    for fi12 := 0; fi12 < 10; fi12++ {
        v11 = v11
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 float64 = (float64((1e10 - 1.5)) * 0.0)
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 int = 885
    v16 := GenS1{F0: v15, F1: (v15 % ((v15) | 1))}
    if ((0 < 0) && false) {
        if false {
            v2 = v2
            v17 := GenS1{F0: v15, F1: v15}
            _ = v17
        }
    } else {
        switch ("\n\t" + "a") {
        case "hello":
            fmt.Println(v2, v2)
        default:
            if (127 >= 2) {
                v18 := GenS2{F0: "", F1: 1}
                for fi19 := 0; fi19 < 10; fi19++ {
                    v20 := GenS3{F0: v15}
                    v2 = ((v2 + v2) + v2)
                    var v21 bool = false
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 4; fi22++ {
                    var v23 int = fi22
                    switch "" {
                    case "hello":
                        if true {
                            var v24 int8 = (v3 | v3)
                            _ = v24
                        }
                    case "":
                        var v25 float32 = (0.1 + (1e10 - -2.25))
                        _ = v25
                    case "a":
                        v26 := GenS4{F0: (v11 << ((120) & 63)), F1: -1}
                        _ = v26
                    }
                    _ = fi22
                    _ = v23
                }
                _ = v18
            }
        }
    }
    if (false && false) {
        v3 += -50
    }
    v3 = (v3 - -115)
    v1 = append(v1, v2)
    v2 = v2
    if (v15 > 1) {
        v27 := GenS5{F0: "hello", F1: 33, F2: -98}
        _ = v27
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        for fi29 := 0; fi29 < 16; fi29++ {
            v30 := GenS2{F0: v2, F1: -1}
            v31 := GenS6{F0: (fi28 & (fi28 | v15)), F1: 65535}
            switch v2 {
            case "\n\t":
                var v32 uint32 = 2147483648
                _ = v32
            case "a":
                fmt.Println(fi29, fi29)
            default:
                fmt.Println(v2, v15, v3)
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        v33 := GenS7{F0: false, F1: -32768}
        _ = fi28
        _ = v33
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        v35 := GenS8{F0: (v11 << ((2) & 63)), F1: fi34}
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 int = (fi34 * fi34)
            _ = fi36
            _ = v37
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
