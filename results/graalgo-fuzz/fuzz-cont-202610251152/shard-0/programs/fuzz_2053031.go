package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
func genFn0(p0 int32, p1 uint16) int32 {
    var v7 uint32 = 2147483648
    var v8 uint = 2
    p0 = (p0 & p0)
    switch "foo bar" {
    case "a":
        var v9 uint32 = (^(v7))
        _ = v9
    default:
        var v10 uint16 = (p1 >> ((2) & 63))
        _ = v10
    }
    var v11 bool = false
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 42
    var v3 uint8 = (^(v2))
    var v4 uint64 = 876
    var v5 float64 = (float64((float32(0.1) * 1.5)) - 1.5)
    v6 := make(map[string]int)
    v12 := []int{}
    if false {
        if false {
            v4 = v4
            v13 := GenS1{F0: 1, F1: -128, F2: true}
            v14 := GenS1{F0: 0, F1: -42, F2: false}
            _ = v13
            _ = v14
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := GenS2{F0: -335}
        v2 = (v3 ^ v3)
        v3 += 42
        _ = fi15
        _ = v16
    }
    switch "\n\t" {
    case "\n\t":
        v17 := GenS1{F0: 255, F1: -75, F2: (!true)}
        _ = v17
    case "hello":
        v3 += v3
    case "a":
        fmt.Println(v3, v2, v4)
    default:
        v3 -= v3
    }
    var v18 float32 = (0.1 * (1.5 + 1e10))
    switch "hello" {
    case "foo bar":
        v19 := GenS2{F0: 9223372036854775807}
        _ = v19
    case "a":
        v20 := GenS1{F0: 305, F1: 1, F2: false}
        _ = v20
    default:
        var v21 bool = (255 <= 265)
        _ = v21
    }
    var v22 uint32 = 2
    for fi23 := 0; fi23 < 7; fi23++ {
        v24 := GenS1{F0: 535, F1: 127, F2: true}
        fmt.Println(fi23, v22)
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 uint16 = 65535
            v2 = (^(v2))
            if len(v12) > 0 {
                v12[0] = fi23
            }
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        switch fi27 {
        case 32767:
            v18 = v18
        case -676:
            switch (fi27 - fi27) {
            case -388:
                for fi28 := 0; fi28 < 15; fi28++ {
                    v4 = v4
                    _ = fi28
                }
            case 127:
                var v29 string = "hello"
                _ = v29
            default:
                if (970 < uint16(v2)) {
                    switch (^(fi27)) {
                    case -42:
                        var v30 uint16 = 149
                        _ = v30
                    default:
                        switch "" {
                        case "foo bar":
                            var v31 uint8 = v3
                            _ = v31
                        default:
                            v5 = 0.0
                        }
                    }
                    _ = len(v12)
                }
            }
        case 9223372036854775807:
            var v32 uint32 = (v22 | 2)
            _ = v32
        }
        fmt.Println(v2, v5)
        _ = fi27
    }
    v2 = v2
    fmt.Println(v5)
    v33 := GenS3{F0: "\n\t"}
    fmt.Println(v4, v3, v2)
    if (v22 <= (v22 >> ((79) & 63))) {
        v4 -= v4
        if false {
            if false {
                fmt.Println(v4, v3)
            }
            v34 := v6[("" + "a")]
            _ = v34
        } else {
            v35 := GenS1{F0: 65535, F1: -123, F2: (!true)}
            for fi36 := 0; fi36 < 2; fi36++ {
                var v37 float64 = v5
                var v38 float32 = v18
                switch "\n\t" {
                case "foo bar":
                    for fi39 := 0; fi39 < 10; fi39++ {
                        v22 = (v22 ^ (v22 * uint32(v4)))
                        _ = fi39
                    }
                case "":
                    fmt.Println(v18, v18, v37)
                case "\n\t":
                    if false {
                        if false {
                            var v40 float64 = (float64(v38) - -2.25)
                            for fi41 := 0; fi41 < 9; fi41++ {
                                var v42 uint8 = v2
                                v37 = (1e10 * v5)
                                v42 += 128
                                _ = fi41
                                _ = v42
                            }
                            v2 += 196
                            _ = v40
                        }
                        var v43 int32 = genFn0(int32(v2), 65535)
                        v2 += 255
                        _ = v43
                    }
                default:
                    switch "hello" {
                    case "":
                        v44 := GenS3{F0: ((("hello" + "a") + "hello") + ("a" + "hello"))}
                        _ = v44
                    case "hello":
                        v45 := v6[("\n\t" + "\n\t")]
                        _ = v45
                    case "a":
                        v38 += -2.25
                    default:
                        if false {
                            v46 := GenS1{F0: 2, F1: 127, F2: true}
                            var v47 uint = 4294967295
                            v48 := GenS1{F0: 42, F1: 127, F2: ((("" + "foo bar") + ("" + (("" + (("foo bar" + "\n\t") + (("foo bar" + "\n\t") + "hello"))) + "a"))) >= ("foo bar" + "hello"))}
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                    }
                }
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = v35
        }
    } else {
        fmt.Println(v4)
        fmt.Println(v3, v3, v18)
    }
    v6[("a" + "a")] = -447
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
