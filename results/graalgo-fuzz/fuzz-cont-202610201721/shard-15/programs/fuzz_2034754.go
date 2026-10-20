package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint) int8 {
    v6 := []int{}
    if (896 <= -40) {
        p0 -= (p0 & p0)
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 2, F1: "foo bar", F2: 212}
        _ = fi7
        _ = v8
    }
    p0 -= 4294967295
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = v6
    return 42
}

func main() {
    v1 := []int{-725, 127, 1}
    v2 := make(map[string]string)
    var v3 uint = 1
    v4 := []string{(("\n\t" + "a") + ("\n\t" + "\n\t")), ("" + "\n\t"), "hello", (("a" + "foo bar") + ("a" + "a"))}
    v5 := []int{128, -521, 128}
    var v9 uint32 = 1
    if false {
        var v10 uint16 = 42
        if (960 <= int16(v3)) {
            fmt.Println(v3, v9, v10)
            var v11 uint16 = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            v3 += uint(v9)
            v3 = v3
            fmt.Println(v10, fi12, v9)
            _ = fi12
        }
        _ = v10
    }
    switch "foo bar" {
    case "hello":
        var v13 int64 = -9223372036854775808
        _ = v13
    case "\n\t":
        switch 42 {
        case 1:
            for fi14 := 0; fi14 < 4; fi14++ {
                v3 = (v3 % ((v3) | 1))
                v15 := GenS1{F0: 190, F1: "", F2: 0}
                v16 := GenS1{F0: 91, F1: "", F2: 81}
                _ = fi14
                _ = v15
                _ = v16
            }
        case -9223372036854775808:
            switch ("" + "foo bar") {
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = 32768
                }
            case "hello":
                var v17 uint = 4294967295
                _ = v17
            case "a":
                v18 := GenS1{F0: 42, F1: ("a" + "a"), F2: 65535}
                _ = v18
            default:
                var v19 bool = (806 > 0)
                _ = v19
            }
        }
    case "a":
        v20 := GenS1{F0: 0, F1: "hello", F2: 42}
        _ = v20
    }
    v3 += (v3 / (((v3 ^ v3)) | 1))
    if (143 >= 2) {
        if (!(0 > 32767)) {
            switch 4294967295 {
            case 32767:
                v9 = (v9 * 0)
            case 4294967295:
                v4 = append(v4, (("a" + "a") + "\n\t"))
            case 127:
                if true {
                    var v21 uint64 = 42
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 1; fi22++ {
                        if true {
                            var v23 uint16 = 32767
                            _ = v23
                        } else {
                            v24 := GenS2{F0: -589}
                            _ = v24
                        }
                        _ = fi22
                    }
                    switch "hello" {
                    case "a":
                        if false {
                            v3 += v3
                        }
                    case "hello":
                        v25 := GenS2{F0: -32768}
                        _ = v25
                    case "\n\t":
                        fmt.Println(v9, v9)
                    default:
                        fmt.Println(v3)
                    }
                }
            }
        }
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        fmt.Println(v3, fi26)
        _ = fi26
    }
    v27 := GenS1{F0: 42, F1: (("hello" + ("a" + "a")) + "\n\t"), F2: 531}
    _ = len(v1)
    var v28 int8 = -51
    for fi29 := 0; fi29 < 15; fi29++ {
        v28 = -97
        _ = fi29
    }
    v30 := v2["\n\t"]
    v3 = (v3 - (^(v3)))
    v9 += (^(v9))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 uint16 = 0
        fmt.Println(v32, v28, fi31)
        _ = fi31
        _ = v32
    }
    v3 = v3
    fmt.Println(v9, v28, v30)
    var v33 uint8 = 94
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v9
}
