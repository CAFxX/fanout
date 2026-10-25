package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    v1 := []string{("" + "foo bar"), ("\n\t" + ("hello" + "hello")), ("a" + (("a" + "foo bar") + ""))}
    var v2 int8 = 127
    var v3 string = "foo bar"
    var v4 int32 = 42
    var v5 int64 = 1
    v6 := make(map[uint32]bool)
    switch v3 {
    case "\n\t":
        _ = len(v1)
    }
    var v7 int32 = -2147483648
    v8 := GenS1{F0: 110}
    switch v3 {
    case "\n\t":
        switch 430 {
        case 938:
            var v9 uint16 = 65535
            _ = v9
        }
    case "foo bar":
        v10 := GenS2{F0: "\n\t", F1: (v5 % ((42) | 1))}
        _ = v10
    case "":
        if (!true) {
            v11 := GenS3{F0: 2, F1: 0}
            _ = v11
        } else {
            var v12 int = 4294967295
            for fi13 := 0; fi13 < 8; fi13++ {
                if (1 <= 42) {
                    v5 = v5
                } else {
                    v2 -= 66
                }
                v5 = (v5 / ((v5) | 1))
                if false {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        if (!(255 == 255)) {
                            v2 = (v2 * v2)
                            fmt.Println(v12)
                        }
                        _ = fi14
                    }
                }
                _ = fi13
            }
            _ = v12
        }
    }
    delete(v6, 2147483647)
    var v15 uint = 1
    if (((float64(0.1) + 0.1) <= ((-2.25 - (float64(1e10) / 1.0)) * 3.14159)) || (0 != 0)) {
        for fi16 := 0; fi16 < 1; fi16++ {
            v5 -= (v5 / ((-9223372036854775808) | 1))
            _ = fi16
        }
        if true {
            switch 32767 {
            case 1:
                var v17 string = "foo bar"
                _ = v17
            }
            v4 = (v4 | v4)
        } else {
            fmt.Println(v3, v15)
            for fi18 := 0; fi18 < 1; fi18++ {
                if true {
                    v2 += int8(v15)
                }
                v19 := GenS1{F0: 32767}
                var v20 float64 = -2.25
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            delete(v6, 2147483647)
            _ = fi21
        }
    }
    v5 -= (v5 | (-(v5)))
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 uint16 = 291
        switch (fi22 | fi22) {
        case 0:
            var v24 int32 = 1
            _ = v24
        default:
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 float32 = (-2.25 * ((0.1 / (float32(-2.25) * float32(3.14159))) - (0.1 + -2.25)))
                _ = fi25
                _ = v26
            }
        }
        _ = fi22
        _ = v23
    }
    v15 += v15
    v6[674] = true
    if true {
        var v27 int8 = 48
        fmt.Println(v2, v2, v7)
        var v28 int16 = -32768
        _ = v27
        _ = v28
    } else {
        var v29 bool = (769 == -723)
        _ = v29
    }
    v3 = v3
    var v30 float32 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
