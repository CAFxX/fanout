package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int32; F2 string }
func genFn0(p0 int32) uint32 {
    switch (("hello" + "\n\t") + ("a" + "a")) {
    case "\n\t":
        v4 := GenS1{F0: 442, F1: 0, F2: 9223372036854775807}
        _ = v4
    default:
        var v5 int16 = 0
        _ = v5
    }
    if false {
        v6 := []int{-809, 128}
        p0 = (p0 << ((1) & 63))
        _ = v6
    } else {
        fmt.Println(p0)
        p0 += 1
    }
    p0 = p0
    _ = p0
    return 4294967295
}
func genFn1() int64 {
    var v7 float32 = 0.0
    var v8 float64 = (1e10 * -2.25)
    v9 := GenS1{F0: 554, F1: -690, F2: 1}
    if (866 <= 550) {
        for fi10 := 0; fi10 < 15; fi10++ {
            v7 -= (3.14159 + v7)
            _ = fi10
        }
        v11 := []int{-1}
        _ = v11
    } else {
        v12 := []int{42, -2147483648, 2147483647, 32767}
        for fi13 := 0; fi13 < 19; fi13++ {
            if len(v12) > 0 {
                v12[0] = fi13
            }
            v8 += v8
            if true {
                v14 := GenS2{F0: 42}
                _ = v14
            }
            _ = fi13
        }
        fmt.Println(v8, v8)
        _ = v12
    }
    _ = v7
    _ = v8
    _ = v9
    return 891
}

func main() {
    var v1 int = 128
    v2 := make(map[int]string)
    var v3 string = ""
    v15 := v2[(v1 | -1)]
    var v16 uint64 = 2
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 bool = false
        var v19 uint64 = v16
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: genFn0(-357), F1: genFn1(), F2: -9223372036854775808}
    switch "foo bar" {
    case "hello":
        switch "" {
        case "\n\t":
            v21 := []int{(v1 >> ((2) & 63)), (v1 % (((v1 >> ((65535) & 63))) | 1)), v1}
            _ = v21
        default:
            var v22 bool = true
            _ = v22
        }
    case "foo bar":
        var v23 int32 = -1
        _ = v23
    }
    if false {
        v24 := GenS3{F0: 933, F1: 1, F2: ((v15 + v3) + v3)}
        _ = v24
    } else {
        fmt.Println(v3)
    }
    var v25 float64 = 0.1
    var v26 float64 = v25
    switch v3 {
    case "hello":
        v1 += v1
    case "foo bar":
        var v27 uint16 = 42
        _ = v27
    case "a":
        var v28 string = ("foo bar" + (v15 + v15))
        _ = v28
    default:
        for fi29 := 0; fi29 < 2; fi29++ {
            v30 := GenS1{F0: 2147483648, F1: -1, F2: -9223372036854775808}
            _ = fi29
            _ = v30
        }
    }
    switch (v1 / ((v1) | 1)) {
    case -762:
        v31 := GenS3{F0: 53, F1: 0, F2: ("a" + "foo bar")}
        _ = v31
    case -1:
        if (-1 >= 42) {
            var v32 uint8 = 2
            for fi33 := 0; fi33 < 6; fi33++ {
                v34 := v2[v1]
                var v35 int = -338
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = v32
        }
    default:
        fmt.Println(v25, v1, v3)
    }
    switch v1 {
    case 811:
        var v36 uint = 4294967295
        _ = v36
    }
    if (!true) {
        v15 = ("a" + "\n\t")
    }
    v37 := GenS3{F0: 776, F1: -744, F2: "hello"}
    v38 := GenS3{F0: 65535, F1: 1, F2: v15}
    v39 := []int{v1, (v1 >> ((178) & 63)), v1, 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v39
}
