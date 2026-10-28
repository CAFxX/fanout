package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 string; F2 int64 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int16; F1 bool; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string }
func genFn0() string {
    var v6 int16 = 0
    if true {
        if true {
            v7 := GenS1{F0: 1}
            _ = v7
        }
        var v8 string = ("" + ("foo bar" + ("" + (("hello" + "") + ("foo bar" + "hello")))))
        _ = v8
    } else {
        var v9 int = 32767
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v6, fi10)
        v11 := []int{(fi10 ^ fi10), fi10, -312, -926}
        v12 := GenS2{F0: 1, F1: "hello", F2: int64(v6)}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch "foo bar" {
    case "foo bar":
        v13 := GenS1{F0: -81}
        _ = v13
    default:
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 int64 = 227
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v6)
    _ = v6
    return ("hello" + "")
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 255
    v3 := make(map[uint32]uint32)
    v4 := []int{4294967295, -300}
    v5 := make(map[int]string)
    v16 := GenS3{F0: 940, F1: 32768}
    fmt.Println(v2)
    switch 9223372036854775807 {
    case -778:
        var v17 int64 = 842
        _ = v17
    default:
        for fi18 := 0; fi18 < 9; fi18++ {
            if len(v4) > 0 {
                v4[0] = fi18
            }
            if ((1 < -1) || false) {
                v2 = v2
            }
            _ = fi18
        }
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        if true {
            for fi20 := 0; fi20 < 14; fi20++ {
                fmt.Println(v2, fi19, fi20)
                if (!false) {
                    v21 := v3[2]
                    _ = v21
                } else {
                    v2 += (v2 | v2)
                    if (!((fi19 + (fi19 & int(v2))) <= (fi19 * fi20))) {
                        var v22 uint8 = 162
                        _ = v22
                    }
                }
                v23 := v3[65535]
                _ = fi20
                _ = v23
            }
            fmt.Println(fi19, fi19)
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 uint32 = 42
                var v26 bool = true
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        _ = len(v4)
        if (18446744073709551615 >= 42) {
            for fi27 := 0; fi27 < 12; fi27++ {
                v2 += (v2 & v2)
                _ = fi27
            }
            var v28 float64 = 1e10
            _ = len(v4)
            _ = v28
        }
        _ = fi19
    }
    v29 := GenS4{F0: 0, F1: false, F2: 0}
    var v30 int = -326
    v4 = append(v4, (v30 & (^(v30))))
    switch (v30 / ((2147483647) | 1)) {
    case 42:
        switch "foo bar" {
        case "hello":
            v31 := GenS5{F0: -32768}
            _ = v31
        case "a":
            var v32 int8 = 0
            _ = v32
        case "":
            fmt.Println(v2, v30)
        default:
            if true {
                for fi33 := 0; fi33 < 6; fi33++ {
                    fmt.Println(fi33, fi33)
                    _ = fi33
                }
                var v34 uint64 = 2
                _ = v34
            } else {
                var v35 uint64 = 573
                _ = v35
            }
        }
    }
    v36 := GenS6{F0: ("" + (("a" + "foo bar") + "a"))}
    var v37 int32 = -428
    var v38 int8 = 1
    var v39 float32 = 1e10
    if false {
        fmt.Println(v30)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
