package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int) uint32 {
    v5 := GenS1{F0: 609, F1: 2147483647, F2: 0}
    v6 := []int{(p0 ^ (^(p0))), (p0 * 0), (p0 >> ((247) & 63)), 656}
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = ""
    var v3 bool = false
    v4 := []int{128, 462, -9223372036854775808, 2147483647}
    var v7 int16 = 42
    var v8 int8 = 0
    var v9 bool = v3
    for fi10 := 0; fi10 < 7; fi10++ {
        if true {
            var v11 uint = 255
            _ = v11
        } else {
            v9 = (-1 == 9223372036854775807)
        }
        switch fi10 {
        case -2147483648:
            v12 := GenS2{F0: 405}
            _ = v12
        }
        var v13 uint = 255
        _ = fi10
        _ = v13
    }
    if (!(!(1 < 42))) {
        v3 = (1 <= 2)
        var v14 uint64 = 71
        _ = v14
    } else {
        v9 = (!(18446744073709551615 <= 18446744073709551615))
        switch 9223372036854775807 {
        case -396:
            fmt.Println(v9)
        case -9223372036854775808:
            var v15 int = 32768
            _ = v15
        case 128:
            v16 := GenS1{F0: 42, F1: 2147483647, F2: v8}
            _ = v16
        default:
            v9 = (!((^(v8)) >= (v8 & v8)))
        }
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        fmt.Println(v8, v2)
        _ = fi17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        v19 := v1[v2]
        _ = fi18
        _ = v19
    }
    v9 = (v7 == v7)
    _ = len(v4)
    var v20 uint = 0
    v21 := GenS2{F0: uint32(v20)}
    for fi22 := 0; fi22 < 4; fi22++ {
        fmt.Println(v3, fi22)
        var v23 uint8 = 42
        _ = fi22
        _ = v23
    }
    if ((42 != 42) && (v9 && v9)) {
        v9 = (480 < 1000)
    }
    fmt.Println(v9, v8)
    if v9 {
        for fi24 := 0; fi24 < 5; fi24++ {
            v9 = v3
            var v25 int = (fi24 ^ (fi24 & fi24))
            var v26 uint = (v20 & 961)
            _ = fi24
            _ = v25
            _ = v26
        }
        if ((v20 / (((^(v20))) | 1)) <= (v20 ^ v20)) {
            v27 := GenS2{F0: 42}
            v3 = false
            fmt.Println(v2, v8)
            _ = v27
        } else {
            v3 = ((^(v20)) > v20)
            v28 := GenS1{F0: v20, F1: 65535, F2: 127}
            v9 = (v7 > (v7 & v7))
            _ = v28
        }
    }
    switch -9223372036854775808 {
    case -741:
        switch v2 {
        case "\n\t":
            v8 = (v8 | (v8 / ((v8) | 1)))
        case "a":
            v2 = "foo bar"
        case "hello":
            v2 = (v2 + ("hello" + ""))
        }
    case -1:
        var v29 uint64 = 2
        _ = v29
    case 9223372036854775807:
        if v3 {
            var v30 int = -459
            _ = v30
        }
    default:
        var v31 uint32 = 355
        _ = v31
    }
    if len(v4) > 0 {
        v4[0] = 128
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
