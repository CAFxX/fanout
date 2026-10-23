package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 bool; F2 int }

func main() {
    var v1 uint32 = 956
    var v2 int16 = -281
    var v3 bool = (false && false)
    v4 := make(map[string]bool)
    var v5 float64 = float64(((3.14159 + 1.5) * 3.14159))
    v6 := GenS1{F0: 4294967295, F1: 1}
    var v7 float32 = 0.1
    if (("a" + "\n\t") >= "\n\t") {
        if (3.14159 > float64((0.1 + (0.1 / 1.5)))) {
            v1 -= v1
        } else {
            v8 := GenS1{F0: 0, F1: 42}
            v2 = v2
            v9 := GenS1{F0: 42, F1: -837}
            _ = v8
            _ = v9
        }
        if v3 {
            switch "foo bar" {
            case "\n\t":
                var v10 float32 = v7
                _ = v10
            case "hello":
                var v11 uint16 = 524
                _ = v11
            default:
                v12 := GenS2{F0: v1, F1: "hello"}
                _ = v12
            }
        } else {
            fmt.Println(v7, v1)
            fmt.Println(v3, v7)
        }
        switch "hello" {
        case "\n\t":
            v4["a"] = v3
        }
    }
    var v13 int16 = (v2 & 0)
    v14 := v4["\n\t"]
    switch 127 {
    case 1:
        var v15 uint8 = 2
        _ = v15
    }
    v16 := GenS3{F0: 42}
    v13 = (v2 * v2)
    switch 324 {
    case -2147483648:
        fmt.Println(v2, v14)
    case 127:
        var v17 int64 = 42
        _ = v17
    case 2147483647:
        if (((!true) && v3) && false) {
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 uint64 = 4294967295
                v19 = v19
                if (65535 < 299) {
                    var v20 int32 = 1
                    if ((v19 % (((v19 >> ((0) & 63))) | 1)) >= (^(v19))) {
                        var v21 float64 = v5
                        for fi22 := 0; fi22 < 14; fi22++ {
                            switch fi22 {
                            case 603:
                                v5 = -2.25
                            }
                            switch ("hello" + "a") {
                            case "":
                                v21 = v5
                            case "a":
                                v13 -= v13
                            }
                            v5 = (((v5 * 1e10) * v5) * float64(0.0))
                            _ = fi22
                        }
                        v7 = v7
                        _ = v21
                    } else {
                        v13 += v13
                    }
                    for fi23 := 0; fi23 < 10; fi23++ {
                        var v24 uint64 = v19
                        var v25 int8 = -1
                        if (!(360 < -1)) {
                            v14 = v3
                            v1 += 823
                            v1 = 633
                        } else {
                            var v26 int8 = v25
                            _ = v26
                        }
                        v27 := []int{fi18}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v27
                    }
                    _ = v20
                }
                _ = fi18
                _ = v19
            }
            v28 := []int{564}
            _ = v28
        }
    }
    switch "a" {
    case "":
        var v29 float32 = (float32(v5) / 0.1)
        _ = v29
    case "\n\t":
        v30 := GenS4{F0: "a", F1: (846 != 252), F2: -435}
        _ = v30
    default:
        fmt.Println(v7, v3)
    }
    fmt.Println(v14, v7)
    v31 := v4["foo bar"]
    switch 9223372036854775807 {
    case 525:
        if (((^(v1)) != v1) || false) {
            var v32 int8 = 100
            v1 = 0
            _ = v32
        } else {
            v33 := GenS2{F0: v1, F1: "hello"}
            _ = v33
        }
    case 128:
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 bool = false
            _ = fi34
            _ = v35
        }
    case 4294967295:
        fmt.Println(v14, v13, v5)
    default:
        for fi36 := 0; fi36 < 15; fi36++ {
            var v37 uint16 = 1
            _ = fi36
            _ = v37
        }
    }
    v38 := []int{32768, -1}
    v5 = (v5 - v5)
    var v39 uint = 42
    var v40 float64 = v5
    var v41 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
