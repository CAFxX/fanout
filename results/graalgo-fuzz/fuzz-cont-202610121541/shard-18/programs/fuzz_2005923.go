package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int16; F1 int16; F2 uint8 }
func genFn0(p0 uint32) uint64 {
    var v5 uint16 = 0
    v6 := []int{1, -544}
    if len(v6) > 0 {
        v6[0] = -268
    }
    switch (("foo bar" + "\n\t") + "\n\t") {
    case "\n\t":
        for fi7 := 0; fi7 < 11; fi7++ {
            if (100 != -128) {
                v5 += v5
                v6 = append(v6, (fi7 - (fi7 & fi7)))
                var v8 uint = 464
                _ = v8
            } else {
                if true {
                    v6 = append(v6, (fi7 | (^(fi7))))
                }
            }
            for fi9 := 0; fi9 < 1; fi9++ {
                v5 -= (v5 & 255)
                p0 = (p0 & (p0 % (((^(p0))) | 1)))
                _ = fi9
            }
            if (2147483647 < 2147483647) {
                for fi10 := 0; fi10 < 17; fi10++ {
                    if ((fi10 ^ fi10) >= fi10) {
                        v11 := GenS1{F0: 328}
                        v12 := GenS2{F0: 1, F1: (fi10 & 32768), F2: 4294967295}
                        var v13 bool = false
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    v6 = append(v6, 1)
                    _ = fi10
                }
                var v14 int64 = 744
                _ = len(v6)
                _ = v14
            } else {
                var v15 uint = 777
                _ = v15
            }
            _ = fi7
        }
    case "":
        fmt.Println(p0)
    case "a":
        if false {
            var v16 float64 = 3.14159
            _ = v16
        } else {
            v17 := GenS1{F0: 585}
            _ = v17
        }
    default:
        if false {
            v18 := GenS3{F0: 2147483647}
            _ = v18
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return uint64(p0)
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    var v3 int32 = 1
    v4 := []int{127, 42, -2147483648}
    if false {
        switch 128 {
        case 9223372036854775807:
            v19 := GenS4{F0: 1}
            _ = v19
        case -9223372036854775808:
            _ = len(v2)
        default:
            var v20 uint8 = 244
            _ = v20
        }
    }
    v3 += (v3 * (-(v3)))
    if (100 >= 42) {
        delete(v1, "\n\t")
        v3 = 2147483647
        var v21 uint16 = 658
        _ = v21
    }
    fmt.Println(v3)
    v22 := GenS4{F0: int64(v3)}
    v23 := GenS3{F0: 4294967295}
    v2 = append(v2, 42)
    var v24 uint64 = 2
    var v25 uint32 = 42
    v26 := GenS1{F0: -9223372036854775808}
    if false {
        switch "hello" {
        case "foo bar":
            v25 = v25
        case "hello":
            fmt.Println(v3)
        case "\n\t":
            var v27 int = 127
            _ = v27
        }
        fmt.Println(v24, v3, v24)
        v28 := GenS5{F0: v25, F1: 1}
        _ = v28
    } else {
        v25 = uint32(v24)
        delete(v1, ("hello" + (("hello" + "hello") + (("\n\t" + ("foo bar" + "a")) + "foo bar"))))
        v3 = v3
    }
    v29 := GenS6{F0: -1, F1: -1, F2: 42}
    v25 = v25
    v3 += 1
    var v30 int64 = -957
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
