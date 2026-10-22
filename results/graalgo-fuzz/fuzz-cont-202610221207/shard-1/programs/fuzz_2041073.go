package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0(p0 int8) int {
    v6 := []int{9223372036854775807, 123, -828, 9223372036854775807}
    if (-667 < -957) {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint64 = 4294967295
            _ = fi7
            _ = v8
        }
    } else {
        v9 := GenS1{F0: p0}
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int32 = 0
            _ = fi10
            _ = v11
        }
        if (0.1 > 1.5) {
            var v12 uint = 629
            _ = v12
        } else {
            var v13 uint = 333
            if len(v6) > 0 {
                v6[0] = 9223372036854775807
            }
            var v14 uint = v13
            _ = v13
            _ = v14
        }
        _ = v9
    }
    var v15 float64 = (1e10 - (1e10 / ((((-2.25 - float64(0.0)) / (((1.5 * 1.5) / (0.0 + 1.5)) - ((0.1 / (1e10 / 1.0)) / 1e10))) / 0.1) / 1.0)))
    var v16 uint32 = 2147483647
    switch "" {
    case "":
        var v17 string = ("" + "a")
        _ = v17
    default:
        fmt.Println(v16)
    }
    _ = p0
    _ = v15
    _ = v16
    _ = v6
    return -932
}
func genFn1(p0 uint16, p1 uint64, p2 int8) int16 {
    for fi18 := 0; fi18 < 17; fi18++ {
        switch fi18 {
        case 2147483647:
            for fi19 := 0; fi19 < 15; fi19++ {
                p2 += p2
                fmt.Println(p1, p1)
                _ = fi19
            }
        case 1:
            var v20 uint16 = p0
            _ = v20
        case 0:
            var v21 string = "\n\t"
            _ = v21
        }
        v22 := []int{(fi18 >> ((1) & 63))}
        var v23 uint8 = 2
        _ = fi18
        _ = v22
        _ = v23
    }
    v24 := []int{}
    v25 := GenS2{F0: p1}
    if (!(!(127 == 42))) {
        fmt.Println(p1)
        v26 := GenS2{F0: 1}
        var v27 uint16 = (^(p0))
        _ = v26
        _ = v27
    } else {
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        var v28 int8 = p2
        _ = v28
    }
    var v29 int8 = (p2 | 127)
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v29
    return -945
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[uint32]bool)
    var v3 bool = (42 <= 586)
    v4 := []string{("\n\t" + ("hello" + ""))}
    var v5 int32 = 42
    v30 := GenS2{F0: 18446744073709551615}
    fmt.Println(v5, v5)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        v5 += 42
        v3 = v3
        if v3 {
            v5 += v5
            var v32 int64 = -1
            fmt.Println(v5, v3)
            _ = v32
        } else {
            v3 = v3
            v3 = v3
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
        _ = fi31
    }
    v33 := GenS2{F0: 42}
    v3 = false
    if len(v4) > 0 {
        v4[0] = ("\n\t" + "a")
    }
    switch "hello" {
    case "a":
        if false {
            var v34 int8 = -128
            var v35 uint16 = 42
            if v3 {
                fmt.Println(v5, v3)
                v36 := GenS3{F0: (v34 + 100), F1: 0}
                var v37 uint = 4294967295
                _ = v36
                _ = v37
            }
            _ = v34
            _ = v35
        } else {
            var v38 uint8 = 42
            _ = v38
        }
    case "hello":
        fmt.Println(v3, v5)
    case "\n\t":
        for fi39 := 0; fi39 < 15; fi39++ {
            var v40 uint8 = 127
            for fi41 := 0; fi41 < 1; fi41++ {
                v42 := GenS1{F0: 0}
                v40 = 1
                _ = fi41
                _ = v42
            }
            v5 = v5
            _ = fi39
            _ = v40
        }
    default:
        var v43 int64 = 0
        _ = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
}
