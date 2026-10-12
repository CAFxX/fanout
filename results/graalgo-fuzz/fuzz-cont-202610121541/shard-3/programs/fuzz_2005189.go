package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int64 = 0
    var v2 uint64 = 1
    var v3 int64 = (v1 + v1)
    v4 := []int{}
    var v5 uint = 2
    v6 := []string{"a", ("hello" + ("a" + ("hello" + "foo bar"))), "foo bar", ("foo bar" + "hello")}
    v4 = append(v4, 949)
    fmt.Println(v3, v2, v3)
    if (-17 != 0) {
        v7 := GenS1{F0: "", F1: 42}
        v1 = v1
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint64 = v2
        v9 -= (v2 | v2)
        _ = fi8
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 11; fi10++ {
            _ = len(v6)
            fmt.Println(v5, v1, v3)
            _ = fi10
        }
        v1 = v1
        if len(v6) > 0 {
            v6[0] = "\n\t"
        }
    }
    if (false && ((39 <= 1000) || ("" <= ""))) {
        _ = len(v4)
        fmt.Println(v5)
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        fmt.Println(fi11, v5)
        _ = fi11
    }
    v12 := GenS2{F0: 0}
    switch -9223372036854775808 {
    case 32767:
        v1 = -9223372036854775808
    case 42:
        var v13 int16 = -484
        _ = v13
    case 0:
        fmt.Println(v5)
    }
    switch -167 {
    case 128:
        fmt.Println(v3, v2)
    case 562:
        v14 := GenS2{F0: -2147483648}
        _ = v14
    default:
        var v15 int = -1
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        fmt.Println(v3)
        switch int(v3) {
        case 9223372036854775807:
            var v17 uint16 = 797
            _ = v17
        default:
            switch fi16 {
            case 9223372036854775807:
                v5 = v5
            default:
                v18 := GenS1{F0: "hello", F1: 1}
                _ = v18
            }
        }
        _ = fi16
    }
    v5 = 65535
    v3 = (v1 & v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
