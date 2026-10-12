package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 string; F1 uint; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 uint16; F2 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool; F1 int; F2 int8 }
type GenS9 struct { F0 int64; F1 uint8 }
func genFn0(p0 int8) bool {
    var v6 int = 4294967295
    v6 = 127
    for fi7 := 0; fi7 < 2; fi7++ {
        switch (^(v6)) {
        case 9223372036854775807:
            v8 := GenS1{F0: "", F1: 42}
            _ = v8
        case 32768:
            fmt.Println(p0, v6, v6)
        case 42:
            fmt.Println(p0, v6)
        default:
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 int8 = p0
                var v11 uint64 = 750
                v12 := GenS1{F0: "", F1: 356}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        }
        _ = fi7
    }
    if true {
        v6 = (v6 | -9223372036854775808)
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 string = "hello"
            v15 := GenS2{F0: v14, F1: 255, F2: 32767}
            fmt.Println(fi13, p0, fi13)
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    _ = p0
    _ = v6
    return true
}
func genFn1() bool {
    v16 := GenS1{F0: "foo bar", F1: 255}
    v17 := GenS1{F0: ("foo bar" + "a"), F1: 358}
    var v18 uint64 = 4294967295
    _ = v16
    _ = v17
    _ = v18
    return false
}

func main() {
    var v1 int16 = -32768
    var v2 float64 = 1.5
    v3 := make(map[string]int)
    var v4 uint64 = 773
    var v5 uint8 = 0
    for fi19 := 0; fi19 < 16; fi19++ {
        fmt.Println(v5, v1, v2)
        _ = fi19
    }
    var v20 int = -9223372036854775808
    var v21 uint = 990
    v4 -= (v4 / ((v4) | 1))
    fmt.Println(v5, v21)
    switch "" {
    case "a":
        v22 := GenS2{F0: "", F1: v21, F2: 0}
        _ = v22
    case "":
        var v23 int32 = 41
        _ = v23
    case "foo bar":
        switch v20 {
        case 1:
            if false {
                var v24 uint64 = (v4 + v4)
                v25 := GenS3{F0: v1}
                _ = v24
                _ = v25
            }
        case 449:
            fmt.Println(v5, v1)
        }
    default:
        v21 = v21
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        if false {
            if ("" <= "hello") {
                var v27 uint = 221
                _ = v27
            }
        } else {
            v28 := []int{}
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 uint32 = 789
                fmt.Println(v21, v30, v30)
                v31 := GenS4{F0: v30, F1: true}
                _ = fi29
                _ = v30
                _ = v31
            }
            if false {
                v32 := GenS5{F0: -825}
                if ((^(v4)) <= v4) {
                    v33 := GenS6{F0: 42, F1: 1, F2: false}
                    for fi34 := 0; fi34 < 17; fi34++ {
                        v35 := GenS7{F0: v21}
                        _ = fi34
                        _ = v35
                    }
                    v36 := GenS6{F0: 0, F1: 32767, F2: (false || (v1 != v1))}
                    _ = v33
                    _ = v36
                } else {
                    v37 := GenS5{F0: 65}
                    if len(v28) > 0 {
                        v28[0] = fi26
                    }
                    fmt.Println(v4, v20)
                    _ = v37
                }
                _ = v32
            }
            _ = v28
        }
        _ = fi26
    }
    v38 := v3["hello"]
    v39 := GenS8{F0: (!(!(115 >= 2147483648))), F1: (v20 / (((-(v20))) | 1)), F2: -103}
    _ = len(v3)
    if (!true) {
        v40 := GenS1{F0: ("foo bar" + (((("foo bar" + "a") + "foo bar") + ("foo bar" + ("\n\t" + "\n\t"))) + "\n\t")), F1: 42}
        _ = v40
    }
    var v41 string = "\n\t"
    var v42 uint64 = v4
    var v43 int16 = int16(v21)
    switch v41 {
    case "":
        v20 -= v20
    case "\n\t":
        v44 := GenS9{F0: -667, F1: 42}
        _ = v44
    case "hello":
        v3[("a" + "\n\t")] = v20
    }
    delete(v3, ((v41 + ((v41 + v41) + ("\n\t" + "a"))) + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
