package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int8; F2 int16 }
type GenS4 struct { F0 uint64; F1 int16; F2 int }
type GenS5 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 int32 }
func genFn0(p0 uint) int8 {
    v7 := GenS1{F0: 2}
    var v8 uint8 = 107
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1() string {
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint64 = 42
            if (("hello" + ("hello" + "a")) >= ((((("" + "hello") + ("hello" + "\n\t")) + ("" + ("foo bar" + ("\n\t" + "foo bar")))) + (("\n\t" + ("a" + "foo bar")) + "\n\t")) + "foo bar")) {
                var v12 uint16 = 168
                _ = v12
            } else {
                v11 = v11
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    switch 2147483647 {
    case 924:
        var v13 uint16 = 255
        _ = v13
    default:
        var v14 float32 = 1e10
        _ = v14
    }
    if true {
        var v15 uint32 = 117
        _ = v15
    } else {
        var v16 uint16 = 328
        v17 := GenS2{F0: 2}
        var v18 uint = 1
        _ = v16
        _ = v17
        _ = v18
    }
    return ((("" + (("" + "hello") + "a")) + "\n\t") + "")
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]int)
    var v3 uint32 = 0
    v4 := make(map[uint32]bool)
    var v5 string = (((v1 + "foo bar") + (((v1 + (v1 + v1)) + "") + "foo bar")) + v1)
    var v6 uint32 = v3
    switch (v1 + ((v1 + "a") + ((("" + ("" + (v5 + (v1 + "hello")))) + "") + ""))) {
    case "a":
        v1 = v5
    case "foo bar":
        switch "foo bar" {
        case "":
            v6 = 805
        case "foo bar":
            var v19 uint64 = 599
            _ = v19
        case "hello":
            switch int(v6) {
            case -1:
                var v20 uint32 = v6
                _ = v20
            case -9223372036854775808:
                var v21 uint32 = 382
                _ = v21
            }
        }
    }
    v22 := v2[-2147483648]
    v23 := GenS2{F0: 1}
    v3 = v6
    if (!((^(v6)) >= (v3 - (v6 >> ((0) & 63))))) {
        switch (v1 + "a") {
        case "a":
            v6 = (v6 | (v3 * 2))
        case "":
            var v24 int = 32767
            _ = v24
        case "foo bar":
            v25 := GenS1{F0: (v3 % (((v3 % ((v3) | 1))) | 1))}
            _ = v25
        }
        if true {
            v26 := GenS3{F0: 2, F1: 42, F2: 958}
            v27 := GenS4{F0: 309, F1: -1, F2: (v22 / ((v22) | 1))}
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 float64 = (-2.25 / 1.0)
                if false {
                    fmt.Println(v6, v29)
                } else {
                    var v30 uint64 = 42
                    switch (fi28 - (v22 - v22)) {
                    case -790:
                        var v31 int32 = 2147483647
                        _ = v31
                    case 0:
                        var v32 bool = ((491 > 255) || false)
                        _ = v32
                    default:
                        for fi33 := 0; fi33 < 2; fi33++ {
                            fmt.Println(fi33)
                            v34 := GenS4{F0: v30, F1: 1, F2: fi28}
                            _ = fi33
                            _ = v34
                        }
                    }
                    _ = v30
                }
                for fi35 := 0; fi35 < 11; fi35++ {
                    v2[(fi35 % (((v22 / ((fi28) | 1))) | 1))] = fi35
                    v29 += v29
                    _ = fi35
                }
                _ = fi28
                _ = v29
            }
            _ = v26
            _ = v27
        }
        var v36 uint64 = 304
        _ = v36
    } else {
        v37 := GenS5{F0: -1, F1: v6, F2: 18446744073709551615}
        fmt.Println(v22, v3)
        _ = v37
    }
    var v38 int8 = 100
    v39 := []int{(^(v22))}
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 float64 = ((3.14159 * 0.1) - ((0.0 * (-2.25 / 0.1)) * 1.5))
        _ = len(v39)
        _ = fi40
        _ = v41
    }
    if len(v39) > 0 {
        v39[0] = v22
    }
    v38 -= int8(v22)
    var v42 int8 = (v38 + v38)
    switch 107 {
    case 1:
        v43 := GenS6{F0: 390}
        _ = v43
    default:
        v5 = v1
    }
    switch v5 {
    case "\n\t":
        var v44 float64 = float64(float32(float64((0.0 - (float32(3.14159) / 1.5)))))
        _ = v44
    case "":
        for fi45 := 0; fi45 < 9; fi45++ {
            v46 := GenS7{F0: (("hello" + v1) + ("a" + ("\n\t" + ("\n\t" + v1)))), F1: -2147483648}
            _ = fi45
            _ = v46
        }
    case "a":
        _ = len(v39)
    }
    if (-2.25 >= ((3.14159 / 0.1) * 0.1)) {
        v1 = v5
    } else {
        v42 = v38
        v5 = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
