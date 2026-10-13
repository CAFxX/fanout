package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int8; F1 int16; F2 int64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint16) uint32 {
    var v7 uint64 = 689
    v8 := GenS1{F0: "foo bar", F1: -759}
    switch -839 {
    case 2147483647:
        v9 := GenS2{F0: 59, F1: p0}
        _ = v9
    case -2147483648:
        var v10 int16 = 1000
        _ = v10
    case 0:
        v11 := []int{-1, 1, -940}
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint8 = 127
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}
func genFn1() int64 {
    switch -2147483648 {
    case 127:
        var v14 int = -9223372036854775808
        _ = v14
    case 9223372036854775807:
        var v15 int32 = 42
        _ = v15
    case -579:
        var v16 int64 = 469
        _ = v16
    }
    if (-1 > -1) {
        var v17 uint = 418
        var v18 float32 = (0.1 * float32((float64(float32(((1.5 * 0.1) * -2.25))) + 1.5)))
        _ = v17
        _ = v18
    }
    v19 := []int{581, 830, -2147483648}
    if false {
        var v20 uint32 = 1
        v20 -= v20
        v21 := GenS2{F0: -128, F1: false}
        _ = v20
        _ = v21
    }
    var v22 uint = 652
    _ = v19
    _ = v22
    return 504
}

func main() {
    v1 := []uint32{65535}
    v2 := make(map[string]uint32)
    v3 := []int{32767}
    v4 := []string{}
    var v5 int16 = -1
    var v6 uint8 = 127
    for fi23 := 0; fi23 < 5; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            if (false && true) {
                fmt.Println(fi23)
                var v25 string = "\n\t"
                var v26 float64 = 1.5
                _ = v25
                _ = v26
            } else {
                if ((^(v6)) > v6) {
                    v27 := GenS3{F0: -128, F1: v5, F2: -1}
                    switch (("a" + "a") + "\n\t") {
                    case "":
                        if (((1e10 * 3.14159) > (3.14159 * 1.5)) || (!false)) {
                            v28 := v2[("foo bar" + ((("foo bar" + ("" + "hello")) + "a") + "foo bar"))]
                            v29 := GenS4{F0: 191}
                            var v30 float32 = (1.5 * ((0.0 / 1.0) - 3.14159))
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            var v31 int8 = -128
                            if ("\n\t" == ("a" + "a")) {
                                v6 = 0
                                delete(v2, ("" + "foo bar"))
                            } else {
                                v5 -= (v5 * 32767)
                            }
                            _ = v31
                        }
                    case "hello":
                        var v32 bool = true
                        _ = v32
                    case "\n\t":
                        var v33 uint16 = 255
                        _ = v33
                    default:
                        if false {
                            if false {
                                if (!false) {
                                    v6 = 131
                                    if ("foo bar" > "\n\t") {
                                        fmt.Println(fi23)
                                        var v34 uint16 = 565
                                        _ = v34
                                    } else {
                                        switch (-(fi23)) {
                                        case 525:
                                            v5 = (-(v5))
                                        case 127:
                                            v35 := GenS1{F0: ("foo bar" + ""), F1: 42}
                                            _ = v35
                                        }
                                        switch (fi24 | 1) {
                                        case 127:
                                            var v36 float32 = (3.14159 + 0.0)
                                            v37 := GenS4{F0: 704}
                                            _ = v36
                                            _ = v37
                                        case 32768:
                                            var v38 int = fi24
                                            _ = v38
                                        }
                                    }
                                } else {
                                    v5 += v5
                                }
                            }
                            _ = len(v4)
                            var v39 int = (fi23 >> ((255) & 63))
                            _ = v39
                        }
                    }
                    v4 = append(v4, "hello")
                    switch ("" + "\n\t") {
                    case "hello":
                        var v40 uint64 = 0
                        _ = v40
                    }
                    _ = v27
                }
                v41 := GenS3{F0: -95, F1: v5, F2: -9223372036854775808}
                var v42 int64 = 1
                _ = v41
                _ = v42
            }
            var v43 uint64 = 444
            _ = fi24
            _ = v43
        }
        v44 := GenS4{F0: 65535}
        _ = fi23
        _ = v44
    }
    v5 = (v5 / ((42) | 1))
    v5 += v5
    var v45 uint8 = v6
    fmt.Println(v6, v6, v6)
    v5 = (-(v5))
    var v46 int = 2147483647
    v45 = (v45 ^ v6)
    var v47 int16 = 1
    v46 = 32768
    var v48 int16 = v5
    v5 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
