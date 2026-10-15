package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint32) uint {
    if true {
        var v7 int8 = -83
        switch "a" {
        case "":
            v8 := []int{128}
            _ = v8
        }
        _ = v7
    } else {
        p0 -= (p0 & p0)
        switch 9223372036854775807 {
        case 423:
            var v9 int32 = -2147483648
            _ = v9
        }
        p0 = p0
    }
    p0 = p0
    _ = p0
    _ = p1
    return 218
}

func main() {
    var v1 int32 = 42
    var v2 int32 = (v1 ^ (v1 + v1))
    v3 := make(map[uint32]int)
    var v4 float32 = 1e10
    v5 := []string{(("hello" + "\n\t") + ((("hello" + ("\n\t" + "a")) + "foo bar") + "a")), "foo bar", ""}
    var v6 int = 9223372036854775807
    if true {
        v2 -= 42
    }
    var v10 int16 = 1
    fmt.Println(v1)
    for fi11 := 0; fi11 < 13; fi11++ {
        _ = len(v5)
        v12 := GenS1{F0: (-(v6)), F1: (v6 * (-(v6))), F2: 225}
        _ = fi11
        _ = v12
    }
    if ((v1 & 2147483647) == v1) {
        var v13 int64 = 624
        v2 = (v2 % ((0) | 1))
        fmt.Println(v2, v10, v6)
        _ = v13
    } else {
        v14 := GenS1{F0: (-(v6)), F1: (v6 * 844), F2: -1}
        switch "a" {
        case "hello":
            fmt.Println(v2, v1)
        case "\n\t":
            fmt.Println(v1, v6)
        }
        v15 := v3[936]
        _ = v14
        _ = v15
    }
    if ((0 <= 42) || (!true)) {
        for fi16 := 0; fi16 < 20; fi16++ {
            v17 := GenS2{F0: ("a" + "hello")}
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int16 = v10
                var v20 float32 = v4
                var v21 float32 = v20
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 int16 = v10
                _ = fi22
                _ = v23
            }
            _ = fi16
            _ = v17
        }
    } else {
        var v24 float64 = 0.1
        v4 -= (0.1 / v4)
        v25 := GenS2{F0: ""}
        _ = v24
        _ = v25
    }
    switch (("foo bar" + ((("foo bar" + ("" + ("" + ("foo bar" + "a")))) + ("" + "foo bar")) + "")) + "foo bar") {
    case "a":
        var v26 int32 = v2
        _ = v26
    case "foo bar":
        _ = len(v5)
    case "hello":
        switch (v6 - v6) {
        case 32768:
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := GenS1{F0: 32768, F1: fi27, F2: 61}
                var v29 uint64 = 42
                _ = fi27
                _ = v28
                _ = v29
            }
        case 4294967295:
            v30 := GenS3{F0: "hello", F1: false}
            _ = v30
        case 953:
            if true {
                v3[21] = v6
                if false {
                    switch ("hello" + "a") {
                    case "":
                        fmt.Println(v1)
                    case "hello":
                        switch 32767 {
                        case -406:
                            var v31 int16 = (v10 % (((^(v10))) | 1))
                            _ = v31
                        }
                    }
                    var v32 float64 = 1e10
                    var v33 int32 = 1
                    _ = v32
                    _ = v33
                }
            }
        default:
            if true {
                v10 = -771
            } else {
                v6 += -9223372036854775808
            }
        }
    }
    var v34 int16 = (v10 * (v10 / (((v10 >> ((255) & 63))) | 1)))
    var v35 uint64 = 405
    v36 := GenS4{F0: 1}
    var v37 int8 = 0
    for fi38 := 0; fi38 < 6; fi38++ {
        for fi39 := 0; fi39 < 19; fi39++ {
            switch ("\n\t" + "hello") {
            case "\n\t":
                v35 = (v35 | 1)
            case "":
                var v40 uint64 = 2
                _ = v40
            }
            _ = fi39
        }
        _ = fi38
    }
    var v41 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
