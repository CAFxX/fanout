package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int }
type GenS2 struct { F0 string; F1 uint32; F2 string }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint64, p1 string) int32 {
    p1 = ((p1 + p1) + p1)
    var v6 uint8 = 128
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 int8, p1 int64) int {
    fmt.Println(p1)
    if (18446744073709551615 == 2) {
        if false {
            v7 := GenS1{F0: (1 < -2147483648), F1: 1, F2: 1}
            if true {
                p1 = (p1 ^ p1)
                p1 -= 728
                v8 := GenS1{F0: (316 <= 0), F1: 255, F2: 9223372036854775807}
                _ = v8
            }
            _ = v7
        }
        var v9 int16 = 42
        _ = v9
    } else {
        var v10 string = ("" + ("a" + "foo bar"))
        _ = v10
    }
    v11 := GenS1{F0: true, F1: 1, F2: 381}
    if (true && true) {
        p0 -= p0
    } else {
        switch 4294967295 {
        case 32767:
            var v12 int16 = 458
            _ = v12
        }
    }
    var v13 int = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v13
    return -282
}

func main() {
    var v1 int = 4294967295
    var v2 uint8 = 255
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    if (402 == 0) {
        v2 = (v2 | 255)
        switch "hello" {
        case "foo bar":
            var v14 string = (("hello" + ("" + "hello")) + "")
            _ = v14
        case "hello":
            v1 = (v1 ^ v1)
        case "\n\t":
            if (42 > 192) {
                var v15 bool = false
                _ = v15
            }
        default:
            if ("a" != (("\n\t" + "foo bar") + ("hello" + ("" + "")))) {
                if true {
                    v2 += 163
                } else {
                    fmt.Println(v1, v2)
                    v16 := []int{(v1 * (v1 & (v1 | 32768)))}
                    _ = v16
                }
                if false {
                    v17 := GenS1{F0: ((v1 / (((v1 * 463)) | 1)) < v1), F1: 2, F2: genFn1(-86, -739)}
                    v1 -= (-(v1))
                    var v18 int32 = 938
                    _ = v17
                    _ = v18
                } else {
                    var v19 int16 = 42
                    v2 += (v2 / (((v2 & (v2 * (v2 - v2)))) | 1))
                    fmt.Println(v2, v19, v2)
                    _ = v19
                }
                switch "a" {
                case "":
                    var v20 uint16 = 611
                    _ = v20
                }
            }
        }
        var v21 uint = 4294967295
        _ = v21
    } else {
        v22 := GenS2{F0: "", F1: 2147483647, F2: "a"}
        fmt.Println(v1, v2)
        if (("" + "a") != (("foo bar" + ("" + (("\n\t" + "foo bar") + ("foo bar" + ("a" + ("" + "")))))) + ("\n\t" + "a"))) {
            var v23 uint64 = 962
            v24 := GenS3{F0: 212, F1: v23}
            _ = v23
            _ = v24
        }
        _ = v22
    }
    if (-1 != 0) {
        v1 += genFn1(-57, -492)
    } else {
        if true {
            var v25 int32 = -564
            _ = v25
        }
        v26 := []int{genFn1(-1, 9223372036854775807), 1}
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        fmt.Println(v2, fi27)
        v1 = fi27
        _ = fi27
    }
    var v28 string = "a"
    if true {
        v2 = v2
        v29 := GenS2{F0: v28, F1: 42, F2: ((v28 + "\n\t") + v28)}
        var v30 int = v1
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: v28, F1: 326, F2: (((v28 + v28) + v28) + (("hello" + (v28 + "hello")) + (("hello" + v28) + v28)))}
    var v32 int32 = 0
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 int8 = 70
        switch v28 {
        case "a":
            v35 := []int{(v1 ^ fi33), (v1 << ((2) & 63))}
            _ = v35
        case "foo bar":
            fmt.Println(v2)
        case "":
            var v36 int16 = 1000
            _ = v36
        default:
            for fi37 := 0; fi37 < 1; fi37++ {
                var v38 uint8 = (v2 << ((1) & 63))
                fmt.Println(v38)
                _ = fi37
                _ = v38
            }
        }
        _ = fi33
        _ = v34
    }
    v1 = v1
    var v39 uint64 = 0
    for fi40 := 0; fi40 < 19; fi40++ {
        for fi41 := 0; fi41 < 12; fi41++ {
            _ = len(v3)
            fmt.Println(v2, v32)
            _ = fi41
        }
        v2 -= (v2 >> ((314) & 63))
        _ = fi40
    }
    var v42 uint32 = 4294967295
    v43 := GenS2{F0: v28, F1: 2147483648, F2: v28}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
