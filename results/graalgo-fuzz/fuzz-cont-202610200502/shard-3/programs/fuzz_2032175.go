package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 int64, p1 int) int64 {
    if (!(-83 != 0)) {
        v4 := GenS1{F0: -128}
        p0 += p0
        _ = v4
    } else {
        v5 := []int{}
        _ = v5
    }
    var v6 uint = 65535
    switch "" {
    case "hello":
        v6 = (^(v6))
    }
    v6 = (v6 * 42)
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 string = ("a" + "\n\t")
    v2 := []string{((v1 + v1) + v1), (v1 + ((v1 + v1) + v1)), (v1 + v1)}
    var v3 float64 = 0.1
    fmt.Println(v1, v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v3)
    switch (v1 + (("" + v1) + ("foo bar" + (v1 + ("" + ("a" + v1)))))) {
    case "a":
        var v7 uint64 = 186
        _ = v7
    case "":
        v8 := GenS1{F0: 42}
        _ = v8
    default:
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int8 = 127
            _ = fi9
            _ = v10
        }
    }
    v3 = v3
    var v11 bool = (0.0 >= -2.25)
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS2{F0: -979, F1: 1000}
        _ = fi12
        _ = v13
    }
    fmt.Println(v3, v3, v11)
    for fi14 := 0; fi14 < 16; fi14++ {
        if (v11 && (!false)) {
            switch "" {
            case "":
                if false {
                    switch (^(fi14)) {
                    case 4294967295:
                        _ = len(v2)
                    case -9223372036854775808:
                        v11 = false
                    default:
                        v1 = v1
                    }
                    var v15 uint32 = 829
                    _ = v15
                } else {
                    v16 := GenS2{F0: -301, F1: -32768}
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v1 = v1
                        _ = fi17
                    }
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + v1)
                    }
                    _ = v16
                }
            default:
                v18 := GenS2{F0: 1, F1: 0}
                _ = v18
            }
            v11 = (!v11)
            switch 128 {
            case -982:
                v11 = true
            case 2147483647:
                var v19 int8 = 127
                _ = v19
            case 4294967295:
                var v20 int32 = 42
                _ = v20
            }
        } else {
            switch ((v1 + v1) + "a") {
            case "hello":
                v1 = v1
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        var v21 uint8 = 1
        _ = fi14
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
