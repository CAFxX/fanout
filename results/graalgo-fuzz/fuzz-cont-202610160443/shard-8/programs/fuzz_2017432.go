package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
type GenS2 struct { F0 string; F1 uint8 }
func genFn0(p0 int8, p1 int64, p2 uint16) int {
    if (42 != 18446744073709551615) {
        p1 = (p1 * p1)
        var v5 uint = 42
        _ = v5
    }
    var v6 uint64 = 2
    if (false && true) {
        v7 := GenS1{F0: 2, F1: 918, F2: (!false)}
        for fi8 := 0; fi8 < 1; fi8++ {
            switch fi8 {
            case 140:
                var v9 uint = 966
                _ = v9
            }
            _ = fi8
        }
        var v10 int32 = -1
        _ = v10
        _ = v7
    } else {
        var v11 int8 = p0
        _ = v11
    }
    if (!(-943 >= 42)) {
        for fi12 := 0; fi12 < 12; fi12++ {
            fmt.Println(p2, p0)
            switch (fi12 % ((fi12) | 1)) {
            case 127:
                v13 := GenS1{F0: 42, F1: 560, F2: (1 == -32768)}
                _ = v13
            case -337:
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 uint16 = 163
                    _ = fi14
                    _ = v15
                }
            }
            _ = fi12
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            fmt.Println(p2, p1, v6)
            _ = fi16
        }
        switch "foo bar" {
        case "\n\t":
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := GenS2{F0: "", F1: 102}
                _ = fi17
                _ = v18
            }
        case "":
            v6 = 2
        case "hello":
            v19 := []int{792, -1, int(p1)}
            _ = v19
        default:
            switch "\n\t" {
            case "hello":
                var v20 uint32 = 42
                _ = v20
            case "":
                v21 := GenS2{F0: ("foo bar" + "foo bar"), F1: 3}
                _ = v21
            case "foo bar":
                if (false && ((p1 % ((p1) | 1)) < p1)) {
                    var v22 uint = 4294967295
                    _ = v22
                }
            default:
                switch "a" {
                case "a":
                    fmt.Println(p0, p0)
                case "\n\t":
                    v23 := []int{0, -377, int(p0)}
                    _ = v23
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}

func main() {
    v1 := []uint32{2, 65535, 0, 1}
    v2 := make(map[string]int)
    var v3 uint64 = 2
    v4 := []string{("foo bar" + "foo bar"), "a", "\n\t", "foo bar"}
    switch 1 {
    case 42:
        var v24 int64 = -9223372036854775808
        _ = v24
    case 127:
        v3 = v3
    }
    v3 = (v3 + v3)
    v3 = 767
    v3 += v3
    delete(v2, "\n\t")
    fmt.Println(v3)
    fmt.Println(v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3)
    var v25 uint32 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
