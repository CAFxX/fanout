package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    var v6 float64 = float64(3.14159)
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int64 = 42
        _ = fi7
        _ = v8
    }
    v6 -= v6
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(fi9)
        v4[("\n\t" + "hello")] = (65535 == 0)
        if false {
            v6 = 0.0
        }
        _ = fi9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        switch (fi10 - fi10) {
        case 42:
            var v11 int8 = 127
            _ = v11
        default:
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := []int{}
                v6 = v6
                _ = fi12
                _ = v13
            }
        }
        var v14 int64 = 42
        switch "hello" {
        case "hello":
            if (((((10 >= 255) || true) || true) && false) && (4294967295 == 2)) {
                var v15 int = 32767
                v1[-1] = (!(!((("\n\t" + "a") >= "") && (42 == 1000))))
                _ = v15
            } else {
                v16 := []int{758, fi10, (fi10 * (fi10 * fi10))}
                for fi17 := 0; fi17 < 18; fi17++ {
                    switch fi17 {
                    case 9223372036854775807:
                        _ = len(v16)
                        if false {
                            var v18 float64 = float64(1e10)
                            if len(v16) > 0 {
                                v16[0] = (fi10 & (fi17 & 0))
                            }
                            _ = v18
                        } else {
                            v14 = 1
                        }
                    case 759:
                        if (!(0.1 >= (1.5 * 1e10))) {
                            v6 = -2.25
                            v19 := GenS1{F0: v14, F1: "a"}
                            _ = v19
                        } else {
                            v20 := GenS1{F0: v14, F1: "\n\t"}
                            _ = v20
                        }
                    }
                    _ = fi17
                }
                _ = v16
            }
        }
        _ = fi10
        _ = v14
    }
    fmt.Println(v6)
    var v21 int32 = -190
    if true {
        var v22 float64 = (v6 / 3.14159)
        var v23 bool = true
        if v23 {
            fmt.Println(v21, v21, v22)
            var v24 int32 = (v21 % ((v21) | 1))
            _ = v24
        } else {
            delete(v4, "foo bar")
            var v25 uint8 = 7
            _ = v25
        }
        _ = v22
        _ = v23
    }
    var v26 uint16 = 814
    switch "hello" {
    case "a":
        switch "hello" {
        case "hello":
            var v27 uint = 2
            _ = v27
        case "foo bar":
            var v28 bool = false
            _ = v28
        }
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v26, v21, v21)
    case "foo bar":
        for fi29 := 0; fi29 < 14; fi29++ {
            v6 = (v6 - 0.0)
            fmt.Println(v6)
            v6 = 0.0
            _ = fi29
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
