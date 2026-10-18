package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int; F2 int }
type GenS4 struct { F0 uint8; F1 uint32; F2 int32 }
func genFn0(p0 int64, p1 uint8) string {
    switch ((("foo bar" + "foo bar") + "") + ((("\n\t" + "hello") + "hello") + "\n\t")) {
    case "foo bar":
        v7 := GenS1{F0: 65535}
        _ = v7
    case "":
        v8 := []int{190, 1}
        _ = v8
    case "hello":
        fmt.Println(p1, p0)
    }
    v9 := []int{-45, -781, 966}
    for fi10 := 0; fi10 < 7; fi10++ {
        if (2 <= 729) {
            p0 = p0
        }
        var v11 int64 = (p0 ^ 0)
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(fi10, v11)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    var v13 float32 = 1e10
    var v14 int32 = -2147483648
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v9
    return "hello"
}
func genFn1() int16 {
    v15 := GenS1{F0: 2}
    if false {
        var v16 uint8 = 128
        v17 := GenS1{F0: 1}
        _ = v16
        _ = v17
    } else {
        if ("" > ("foo bar" + ("\n\t" + (("" + ((("foo bar" + (("" + "") + "\n\t")) + "a") + "hello")) + "a")))) {
            var v18 float64 = (0.0 / (-2.25 / (-2.25 / (-2.25 / 0.1))))
            _ = v18
        }
        v19 := GenS1{F0: 860}
        var v20 float32 = (0.0 * -2.25)
        _ = v19
        _ = v20
    }
    var v21 int = 713
    _ = v15
    _ = v21
    return 42
}

func main() {
    var v1 uint = 806
    v2 := []int{32768, 1}
    var v3 int16 = -32768
    v4 := make(map[string]bool)
    var v5 uint16 = 1
    var v6 bool = (-1 < 1)
    for fi22 := 0; fi22 < 5; fi22++ {
        if v6 {
            var v23 uint8 = 210
            v3 += (v3 | -787)
            fmt.Println(v1, v1, v5)
            _ = v23
        } else {
            for fi24 := 0; fi24 < 18; fi24++ {
                v3 -= -1
                v25 := GenS2{F0: v5}
                _ = fi24
                _ = v25
            }
            v6 = v6
            var v26 int32 = 767
            _ = v26
        }
        var v27 int = (fi22 >> (((^(v1))) & 63))
        _ = fi22
        _ = v27
    }
    var v28 uint8 = 127
    v29 := GenS1{F0: (^(v1))}
    for fi30 := 0; fi30 < 1; fi30++ {
        fmt.Println(v28)
        v31 := GenS3{F0: 100, F1: (fi30 >> ((uint(v5)) & 63)), F2: fi30}
        _ = fi30
        _ = v31
    }
    v5 += 0
    if (3.14159 < 1.5) {
        v1 = (^(v1))
    } else {
        v5 = 1
    }
    var v32 bool = (v1 == v1)
    var v33 bool = ((-2.25 + ((float32(3.14159) - 0.1) * 1.5)) >= 0.1)
    for fi34 := 0; fi34 < 18; fi34++ {
        var v35 uint = 910
        if v33 {
            if false {
                v36 := v4[("" + "a")]
                _ = v36
            }
            v37 := GenS4{F0: 128, F1: 4294967295, F2: 1}
            var v38 uint8 = 243
            _ = v37
            _ = v38
        }
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
