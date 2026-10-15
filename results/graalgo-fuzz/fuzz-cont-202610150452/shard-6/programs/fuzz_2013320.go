package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 string }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64, p1 uint16) bool {
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(fi5, p1, p1)
        var v6 uint = 42
        _ = fi5
        _ = v6
    }
    if true {
        p0 += (p0 & p0)
        v7 := []int{-656, 127}
        _ = v7
    } else {
        if (0 == -2147483648) {
            var v8 uint64 = 618
            switch "a" {
            case "a":
                v9 := []int{32768, -1, 42, -951}
                _ = v9
            case "hello":
                var v10 uint64 = (v8 / ((18446744073709551615) | 1))
                _ = v10
            default:
                v8 -= 2
            }
            _ = v8
        }
        v11 := []int{4294967295, -769, -460}
        _ = v11
    }
    p1 -= 0
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return (!true)
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 42
    var v3 int = v2
    var v4 string = "a"
    v4 = "hello"
    v12 := GenS1{F0: -1, F1: -1, F2: "foo bar"}
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 string = "a"
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 1, F1: 100, F2: "foo bar"}
    var v16 uint8 = 0
    if (false || true) {
        for fi17 := 0; fi17 < 15; fi17++ {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 int16 = -760
                v19 -= v19
                v20 := GenS1{F0: 42, F1: 1, F2: v4}
                v21 := []int{}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            switch (v3 >> ((42) & 63)) {
            case 2147483647:
                v22 := GenS2{F0: 42}
                _ = v22
            case 9223372036854775807:
                var v23 int32 = -1
                _ = v23
            }
            _ = fi17
        }
    }
    v16 -= v16
    v24 := GenS2{F0: 202}
    v25 := v1[((v4 + ((("\n\t" + v4) + "\n\t") + v4)) + ("hello" + v4))]
    switch "foo bar" {
    case "\n\t":
        v2 = (v2 << ((2) & 63))
    case "a":
        var v26 int32 = 2147483647
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
