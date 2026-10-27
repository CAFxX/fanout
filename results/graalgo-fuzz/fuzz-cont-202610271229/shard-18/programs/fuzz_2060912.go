package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int8 = -1
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    v5 := GenS1{F0: 217}
    fmt.Println(v1)
    v6 := GenS2{F0: 4294967295, F1: v1}
    var v7 int64 = -9223372036854775808
    v7 += (v7 * 0)
    var v8 uint = 1
    switch "hello" {
    case "":
        v9 := v4["hello"]
        _ = v9
    case "foo bar":
        if true {
            v8 = 255
            if (65535 > 397) {
                switch ("a" + "foo bar") {
                case "":
                    for fi10 := 0; fi10 < 8; fi10++ {
                        var v11 float64 = 1e10
                        _ = fi10
                        _ = v11
                    }
                case "a":
                    if (false && (1e10 <= (((((((0.1 / 1.5) / 1.5) / (0.1 + 1e10)) + 0.1) - 1e10) - (1e10 + ((1e10 - ((1e10 - 3.14159) * -2.25)) / 3.14159))) / 1.0))) {
                        if true {
                            v8 = (^(v8))
                            v12 := v4[(("" + "\n\t") + "\n\t")]
                            _ = v12
                        }
                        v1 = v1
                        v8 -= 144
                    } else {
                        var v13 bool = (((0 == 2) && false) || true)
                        v1 = 1
                        v14 := []int{166}
                        _ = v13
                        _ = v14
                    }
                case "hello":
                    var v15 int8 = (v1 / (((v1 << ((v8) & 63))) | 1))
                    _ = v15
                }
                fmt.Println(v7, v8, v8)
            }
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 uint8 = 1
                delete(v3, "a")
                _ = fi16
                _ = v17
            }
        } else {
            fmt.Println(v1, v1, v7)
        }
    default:
        switch "hello" {
        case "foo bar":
            fmt.Println(v8)
        case "":
            for fi18 := 0; fi18 < 4; fi18++ {
                switch (("hello" + "") + ("a" + ("foo bar" + ("" + "\n\t")))) {
                case "hello":
                    v8 -= v8
                case "":
                    v1 += v1
                case "\n\t":
                    v19 := []int{(fi18 << ((65535) & 63))}
                    _ = v19
                }
                _ = fi18
            }
        }
    }
    switch "a" {
    case "foo bar":
        v20 := GenS2{F0: 42, F1: v1}
        _ = v20
    case "a":
        switch ("\n\t" + (("hello" + "foo bar") + ("hello" + "a"))) {
        case "":
            v1 = v1
        case "a":
            var v21 float32 = (((1e10 + 0.1) - -2.25) + (0.0 * 0.0))
            _ = v21
        }
    default:
        v22 := GenS1{F0: 0}
        _ = v22
    }
    if (!false) {
        fmt.Println(v1, v8, v7)
        v7 += 1
    }
    v7 = int64(v1)
    v23 := GenS2{F0: 4294967295, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
