package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint32; F2 uint }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }
type GenS4 struct { F0 string; F1 int64; F2 int16 }
func genFn0(p0 int64) uint32 {
    var v5 uint16 = 255
    switch (("\n\t" + "foo bar") + "hello") {
    case "a":
        v5 -= (v5 * 255)
    case "\n\t":
        v5 = (^(v5))
    }
    var v6 int16 = 0
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int, p1 int8, p2 uint64) uint8 {
    switch "hello" {
    case "foo bar":
        fmt.Println(p2, p1)
    case "a":
        v7 := []int{p0, -938, (p0 ^ p0)}
        _ = v7
    }
    fmt.Println(p1, p1)
    fmt.Println(p1, p2)
    v8 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 247
}

func main() {
    var v1 uint = 694
    var v2 uint8 = 56
    v3 := make(map[string]int)
    var v4 float64 = -2.25
    if (9223372036854775807 == 0) {
        fmt.Println(v1, v4, v4)
    } else {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: fi9}
            v11 := []int{fi9, fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
        v2 -= (v2 | v2)
    }
    var v12 int64 = -9223372036854775808
    v13 := GenS1{F0: 0}
    if (!(false && (v1 < (v1 * (^(v1)))))) {
        v4 = (v4 / (v4 + -2.25))
        if ((true && true) || true) {
            for fi14 := 0; fi14 < 5; fi14++ {
                _ = len(v3)
                v12 -= 0
                _ = fi14
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 int8 = -128
                var v17 float64 = (v4 - ((0.0 + v4) * 0.0))
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            v18 := GenS2{F0: 42, F1: 2, F2: (v1 ^ v1)}
            for fi19 := 0; fi19 < 5; fi19++ {
                v12 = v12
                fmt.Println(v1, v12, v1)
                _ = fi19
            }
            _ = v18
        }
    } else {
        var v20 uint = (v1 ^ (v1 ^ v1))
        _ = len(v3)
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        v1 += v1
        _ = fi21
    }
    if false {
        v22 := GenS2{F0: int64(v2), F1: 1, F2: v1}
        _ = v22
    }
    if true {
        _ = len(v3)
        switch -2147483648 {
        case -735:
            v23 := GenS1{F0: 127}
            _ = v23
        case 262:
            var v24 int8 = -1
            _ = v24
        }
        v4 = 1.5
    } else {
        v3[("foo bar" + "hello")] = -9223372036854775808
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(v12, fi25)
            _ = fi25
        }
        switch 350 {
        case -9223372036854775808:
            v2 += v2
        }
    }
    v26 := GenS3{F0: 2, F1: ("" + "foo bar"), F2: (v12 * 955)}
    if (!(!true)) {
        v4 += v4
        for fi27 := 0; fi27 < 11; fi27++ {
            v4 = 0.1
            delete(v3, "a")
            _ = fi27
        }
    } else {
        v28 := GenS2{F0: v12, F1: 2147483648, F2: v1}
        v12 = v12
        for fi29 := 0; fi29 < 9; fi29++ {
            switch (fi29 - 2147483647) {
            case 32767:
                v30 := GenS1{F0: (fi29 / ((fi29) | 1))}
                _ = v30
            case -195:
                for fi31 := 0; fi31 < 18; fi31++ {
                    v32 := []int{128, (^(fi29)), fi31}
                    v1 = v1
                    v3[""] = (^(fi31))
                    v4 -= (-2.25 / v4)
                    _ = fi31
                    _ = v32
                }
            case 1:
                fmt.Println(fi29, fi29)
            default:
                if (-38 < -42) {
                    v33 := GenS3{F0: 0, F1: ("" + "foo bar"), F2: -515}
                    _ = v33
                } else {
                    v2 = v2
                    v12 += v12
                }
            }
            _ = fi29
        }
        _ = v28
    }
    delete(v3, ("\n\t" + "foo bar"))
    v2 -= v2
    var v34 uint8 = 127
    for fi35 := 0; fi35 < 6; fi35++ {
        if (0.0 > 1e10) {
            v36 := []int{}
            var v37 bool = (!(true && (!false)))
            if false {
                var v38 uint8 = v2
                _ = v38
            }
            _ = v36
            _ = v37
        }
        _ = len(v3)
        _ = fi35
    }
    v4 = v4
    if ((830 < 1) && ((1000 != 32767) && true)) {
        v39 := GenS4{F0: "a", F1: v12, F2: 0}
        v40 := v3["foo bar"]
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
}
