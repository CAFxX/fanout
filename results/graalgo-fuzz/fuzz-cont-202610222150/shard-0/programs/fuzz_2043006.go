package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 1
    var v3 float64 = (float64(((0.0 + float32(((1e10 - -2.25) - -2.25))) - 1.5)) / 1.0)
    v4 := make(map[string]int)
    for fi5 := 0; fi5 < 13; fi5++ {
        switch (fi5 & fi5) {
        case -9223372036854775808:
            v1["\n\t"] = (v2 - v2)
        case -997:
            fmt.Println(v3, v2)
        case 128:
            var v6 int32 = 933
            _ = v6
        default:
            var v7 float64 = 0.0
            _ = v7
        }
        switch -627 {
        case 1:
            v8 := GenS1{F0: "foo bar", F1: 208}
            _ = v8
        case -1:
            v9 := GenS2{F0: 18446744073709551615}
            _ = v9
        case -2147483648:
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 string = (("" + "") + "")
                _ = len(v1)
                _ = fi10
                _ = v11
            }
        }
        _ = fi5
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v3 = 0.0
        if false {
            var v13 uint8 = 31
            _ = v13
        } else {
            v14 := []int{32768, (fi12 & fi12), fi12, (fi12 & 0)}
            _ = v14
        }
        _ = fi12
    }
    v2 += v2
    var v15 uint8 = 2
    switch ("" + "foo bar") {
    case "":
        switch "hello" {
        case "hello":
            for fi16 := 0; fi16 < 11; fi16++ {
                v15 += v15
                v15 -= v15
                _ = fi16
            }
        case "":
            fmt.Println(v15)
        case "foo bar":
            var v17 int64 = -1
            _ = v17
        }
    case "\n\t":
        v18 := GenS2{F0: 1}
        _ = v18
    }
    v15 -= v15
    var v19 string = "foo bar"
    fmt.Println(v19, v15, v3)
    fmt.Println(v15, v15, v19)
    fmt.Println(v2, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
