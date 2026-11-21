package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }

func main() {
    var v1 uint8 = 2
    var v2 int = 168
    v3 := make(map[int]bool)
    var v4 int64 = 9223372036854775807
    if (18446744073709551615 > uint(v1)) {
        if true {
            for fi5 := 0; fi5 < 2; fi5++ {
                switch ("foo bar" + "hello") {
                case "\n\t":
                    v6 := GenS1{F0: (-322 <= -1), F1: -822}
                    _ = v6
                default:
                    v7 := []int{1, 1, -1, 32768}
                    _ = v7
                }
                v2 -= 32768
                _ = fi5
            }
        } else {
            if (-59 < 0) {
                switch "hello" {
                case "hello":
                    fmt.Println(v1)
                case "":
                    var v8 int32 = -2147483648
                    _ = v8
                default:
                    v4 = v4
                }
                v9 := []int{int(v4), (-(v2)), -1}
                _ = v9
            }
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 int = v2
                _ = fi10
                _ = v11
            }
            fmt.Println(v4, v2, v2)
        }
        var v12 int8 = 42
        _ = v12
    } else {
        fmt.Println(v4, v1, v1)
        for fi13 := 0; fi13 < 16; fi13++ {
            v3[(v2 / (((fi13 - 0)) | 1))] = true
            v14 := GenS1{F0: false, F1: -2147483648}
            _ = fi13
            _ = v14
        }
        v2 -= 911
    }
    v2 -= v2
    v15 := GenS1{F0: false, F1: 2147483647}
    switch "foo bar" {
    case "\n\t":
        v2 -= 9223372036854775807
    case "a":
        v2 -= (v2 * 0)
    }
    var v16 bool = false
    if (!true) {
        v17 := GenS1{F0: (v2 > v2), F1: -1}
        _ = v17
    }
    v1 = 1
    var v18 float32 = 1e10
    v19 := v3[(v2 ^ v2)]
    var v20 float64 = float64(1e10)
    v19 = false
    _ = len(v3)
    for fi21 := 0; fi21 < 3; fi21++ {
        v4 += (v4 ^ (-(v4)))
        fmt.Println(fi21)
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 int32 = -1
            fmt.Println(v2)
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    var v24 uint64 = 2
    switch ("\n\t" + "\n\t") {
    case "a":
        switch v2 {
        case 31:
            for fi25 := 0; fi25 < 12; fi25++ {
                switch v2 {
                case -729:
                    var v26 float32 = v18
                    _ = v26
                default:
                    v4 = -9223372036854775808
                }
                for fi27 := 0; fi27 < 14; fi27++ {
                    v28 := v3[v2]
                    v29 := GenS1{F0: false, F1: 42}
                    v30 := []int{(v2 << ((2) & 63)), fi27}
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                _ = fi25
            }
        }
    case "hello":
        fmt.Println(v16, v20, v24)
    case "\n\t":
        _ = len(v3)
    default:
        v31 := GenS1{F0: ((-(v2)) == v2), F1: 1}
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
