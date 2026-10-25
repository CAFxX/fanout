package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{("a" + ""), "hello", "hello", ((("a" + ((("foo bar" + ("hello" + "")) + "\n\t") + ("a" + (("\n\t" + "hello") + "\n\t")))) + "a") + ((("hello" + "\n\t") + "a") + ""))}
    var v3 int16 = 110
    var v4 int8 = -70
    v5 := GenS1{F0: 294, F1: 1, F2: 42}
    switch "hello" {
    case "a":
        v6 := GenS1{F0: 1, F1: v4, F2: 65535}
        _ = v6
    case "foo bar":
        switch "a" {
        case "":
            var v7 int64 = -9223372036854775808
            _ = v7
        case "a":
            v2 = append(v2, "")
        default:
            var v8 int = -207
            _ = v8
        }
    case "":
        var v9 uint8 = 2
        _ = v9
    }
    v4 = (-(v4))
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS2{F0: 0}
        var v12 uint64 = 2
        v3 = v3
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int8 = v4
    if (true || (!(!false))) {
        for fi14 := 0; fi14 < 13; fi14++ {
            switch "a" {
            case "foo bar":
                if (((0.0 * ((float32((0.0 * 0.1)) / 1e10) * 0.1)) + 0.0) == 1e10) {
                    for fi15 := 0; fi15 < 6; fi15++ {
                        fmt.Println(v3, fi14, v13)
                        _ = fi15
                    }
                    for fi16 := 0; fi16 < 3; fi16++ {
                        switch (fi16 / ((1) | 1)) {
                        case 9223372036854775807:
                            if true {
                                v17 := v1[42]
                                _ = v17
                            }
                        }
                        v3 -= v3
                        _ = fi16
                    }
                    v3 -= (v3 ^ v3)
                }
            default:
                var v18 uint64 = 651
                _ = v18
            }
            switch ("" + "") {
            case "\n\t":
                fmt.Println(v3, fi14)
            }
            v4 -= 127
            _ = fi14
        }
    }
    if (62 < 2147483647) {
        var v19 int64 = -172
        if (2147483647 >= -725) {
            v13 += 1
            var v20 string = ("hello" + "\n\t")
            v4 -= (v4 % ((v13) | 1))
            _ = v20
        } else {
            switch -1 {
            case -728:
                v19 -= v19
            case 32767:
                v21 := GenS1{F0: 0, F1: 0, F2: 4294967295}
                _ = v21
            case 2147483647:
                var v22 int16 = v3
                _ = v22
            }
            var v23 float32 = 1.5
            v1[2147483647] = 398
            _ = v23
        }
        v19 -= 0
        _ = v19
    } else {
        var v24 uint16 = 608
        _ = v24
    }
    if false {
        v1[797] = 668
        if true {
            var v25 float32 = ((3.14159 / 0.1) + 1e10)
            v2 = append(v2, "")
            fmt.Println(v13, v25)
            _ = v25
        } else {
            var v26 int8 = (^(v13))
            for fi27 := 0; fi27 < 9; fi27++ {
                if (v26 < (v4 % (((v4 >> ((0) & 63))) | 1))) {
                    delete(v1, 0)
                    if false {
                        fmt.Println(fi27, fi27, fi27)
                    }
                }
                v26 = (v26 * -128)
                _ = fi27
            }
            fmt.Println(v3)
            _ = v26
        }
    }
    var v28 uint16 = 415
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
