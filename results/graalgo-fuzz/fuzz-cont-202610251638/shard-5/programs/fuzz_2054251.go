package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int; F2 bool }
type GenS5 struct { F0 int8; F1 int; F2 int16 }
type GenS6 struct { F0 int64; F1 string; F2 uint64 }
func genFn0(p0 uint, p1 int8) int64 {
    var v6 uint32 = 82
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(p0)
        _ = fi7
    }
    var v8 int8 = p1
    if (2 == 200) {
        if (0 >= 1) {
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := []int{-9223372036854775808}
                var v11 float32 = 0.1
                switch "hello" {
                case "a":
                    fmt.Println(v6)
                case "":
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v13 := GenS1{F0: 0}
                        var v14 string = "a"
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case "foo bar":
                    for fi15 := 0; fi15 < 8; fi15++ {
                        if (false || (false && false)) {
                            _ = len(v10)
                            if false {
                                p0 -= p0
                                var v16 int32 = -52
                                _ = v16
                            } else {
                                var v17 bool = true
                                if len(v10) > 0 {
                                    v10[0] = fi9
                                }
                                _ = v17
                            }
                            v8 = v8
                            v6 += v6
                        }
                        switch (("foo bar" + "hello") + "hello") {
                        case "a":
                            switch 421 {
                            case -2147483648:
                                fmt.Println(p1)
                            case 656:
                                fmt.Println(v6, fi9, p0)
                            case 32767:
                                v6 -= (v6 & v6)
                            }
                        case "foo bar":
                            var v18 int8 = (^(p1))
                            _ = v18
                        case "":
                            switch ("hello" + "\n\t") {
                            case "":
                                fmt.Println(fi15, p1, fi9)
                            case "a":
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            }
                        }
                        _ = fi15
                    }
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            v8 -= 108
            v8 += (v8 >> ((581) & 63))
        }
        v6 = v6
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 42
}

func main() {
    v1 := []int{756, 636}
    v2 := []int{}
    var v3 uint64 = 1
    v4 := make(map[int]string)
    var v5 bool = (true || true)
    v1 = append(v1, 0)
    fmt.Println(v3, v3)
    v19 := GenS2{F0: (-128 > 100)}
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            switch "hello" {
            case "hello":
                var v22 int64 = 9223372036854775807
                _ = v22
            }
            v3 += v3
            v5 = v5
            _ = fi21
        }
        v23 := GenS3{F0: v3}
        var v24 int16 = -872
        _ = fi20
        _ = v23
        _ = v24
    }
    v3 = v3
    switch "foo bar" {
    case "":
        v25 := GenS4{F0: 42, F1: 42, F2: v5}
        _ = v25
    case "a":
        var v26 int16 = 518
        _ = v26
    }
    var v27 uint = 889
    v28 := GenS5{F0: -128, F1: 2147483647, F2: 1}
    v5 = false
    switch ("a" + ("foo bar" + ("" + "\n\t"))) {
    case "":
        if (((1.5 + (float32(3.14159) + (-2.25 - 0.1))) - (float32(3.14159) - (1.5 / 1.5))) > 0.0) {
            fmt.Println(v3, v5, v5)
        }
    case "foo bar":
        switch ("" + "foo bar") {
        case "hello":
            var v29 uint64 = 2
            _ = v29
        case "foo bar":
            var v30 int = -2147483648
            _ = v30
        default:
            v31 := GenS6{F0: -9223372036854775808, F1: "", F2: 125}
            _ = v31
        }
    case "\n\t":
        var v32 int8 = 42
        _ = v32
    default:
        var v33 uint64 = v3
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
