package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 int64 = -1
    var v2 uint32 = 471
    var v3 uint = 226
    v4 := []int{4294967295, 32767, 127, -860}
    v5 := GenS1{F0: (v1 - v1), F1: (v3 * v3), F2: v3}
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 12; fi7++ {
        _ = len(v4)
        v2 -= 925
        _ = fi7
    }
    v1 += v1
    if true {
        v8 := GenS2{F0: ("" + "foo bar"), F1: 42}
        _ = v8
    } else {
        var v9 bool = true
        v10 := GenS1{F0: v1, F1: (^(v3)), F2: v3}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        switch ("\n\t" + "foo bar") {
        case "hello":
            fmt.Println(v3, v6, v3)
        case "a":
            var v12 int32 = 314
            _ = v12
        default:
            if false {
                fmt.Println(fi11, v1, v1)
                v13 := GenS2{F0: ("" + "a"), F1: 1}
                v1 -= (v1 << ((742) & 63))
                _ = v13
            } else {
                var v14 int64 = (v1 >> (((v3 * v3)) & 63))
                for fi15 := 0; fi15 < 15; fi15++ {
                    switch (("" + ("foo bar" + (("hello" + "hello") + ("foo bar" + ("\n\t" + "\n\t"))))) + ("a" + ((("\n\t" + "hello") + "") + (("foo bar" + "foo bar") + "foo bar")))) {
                    case "":
                        v1 += (v1 << (((v3 % ((v3) | 1))) & 63))
                    }
                    switch (fi15 ^ fi11) {
                    case 32767:
                        var v16 int8 = 0
                        fmt.Println(fi15, fi15)
                        var v17 string = "a"
                        _ = v16
                        _ = v17
                    case -906:
                        if false {
                            v18 := GenS3{F0: 42}
                            _ = v18
                        }
                    case 1:
                        v19 := GenS2{F0: ("foo bar" + "\n\t"), F1: 4294967295}
                        _ = v19
                    }
                    _ = fi15
                }
                switch 685 {
                case -964:
                    v2 -= v2
                default:
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                }
                _ = v14
            }
        }
        fmt.Println(v2, v6)
        _ = fi11
    }
    var v20 string = "\n\t"
    v3 -= v3
    var v21 float32 = v6
    if (false && (v2 == (^(v2)))) {
        var v22 float32 = 1e10
        v3 = 350
        v2 += (v2 / ((v2) | 1))
        _ = v22
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v23 int32 = -2147483648
    v4 = append(v4, 2147483647)
    if (v1 != (v1 + -1)) {
        v24 := GenS3{F0: 0}
        if ((true || (!false)) && true) {
            if false {
                switch "foo bar" {
                case "":
                    v6 = 0.0
                case "a":
                    var v25 uint32 = (v2 / ((42) | 1))
                    _ = v25
                case "foo bar":
                    v2 -= v2
                }
                var v26 uint64 = 4294967295
                _ = v26
            } else {
                switch v20 {
                case "a":
                    v2 = v2
                case "hello":
                    v27 := GenS4{F0: (v23 << ((0) & 63))}
                    _ = v27
                case "":
                    fmt.Println(v23)
                }
                var v28 uint16 = 42
                _ = v28
            }
        }
        v21 -= v6
        _ = v24
    } else {
        for fi29 := 0; fi29 < 19; fi29++ {
            v23 = 42
            v6 = (1e10 / v6)
            _ = fi29
        }
        var v30 int8 = 127
        _ = v30
    }
    fmt.Println(v3, v3)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
