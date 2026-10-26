package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    v1 := []int{1, 128, -436, 1}
    v2 := make(map[string]uint32)
    v3 := []uint32{}
    v4 := GenS1{F0: 589}
    v5 := GenS2{F0: "foo bar", F1: 218}
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int32 = -1
        switch (fi6 * (fi6 << ((243) & 63))) {
        case -589:
            v7 += -1
        default:
            fmt.Println(v7)
        }
        v7 += 1
        _ = fi6
        _ = v7
    }
    var v8 float32 = 0.1
    fmt.Println(v8)
    v8 = float32(3.14159)
    if (("\n\t" == "") || (!(false && false))) {
        v8 -= (v8 + v8)
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v8, fi9)
        v8 = 1e10
        if true {
            for fi10 := 0; fi10 < 11; fi10++ {
                v8 = ((v8 - 1e10) + v8)
                var v11 int16 = 536
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 14; fi12++ {
                if false {
                    switch fi9 {
                    case 4294967295:
                        v8 = v8
                    case 751:
                        var v13 int8 = 0
                        _ = v13
                    case -2147483648:
                        var v14 int32 = -76
                        _ = v14
                    default:
                        fmt.Println(fi9, fi9, v8)
                    }
                    switch (fi9 - (-(fi12))) {
                    case -568:
                        v3 = append(v3, 74)
                    case 239:
                        fmt.Println(fi9)
                    case 399:
                        switch "a" {
                        case "foo bar":
                            var v15 int64 = -9223372036854775808
                            _ = v15
                        case "a":
                            v3 = append(v3, 1)
                        }
                        v8 -= v8
                        switch "foo bar" {
                        case "a":
                            v8 = float32(3.14159)
                        default:
                            fmt.Println(v8, v8, fi9)
                        }
                    }
                    v2["a"] = 2147483647
                }
                v16 := GenS3{F0: 0, F1: false}
                v8 = (1e10 - (v8 + v8))
                _ = fi12
                _ = v16
            }
        } else {
            if false {
                v8 += v8
                v8 += ((v8 * 1.5) / v8)
            }
            switch "a" {
            case "\n\t":
                for fi17 := 0; fi17 < 16; fi17++ {
                    if (32767 > 938) {
                        fmt.Println(fi17)
                        v1 = append(v1, 9223372036854775807)
                    }
                    v8 -= ((((v8 + v8) * v8) + -2.25) / 1.5)
                    _ = fi17
                }
            default:
                v8 -= (3.14159 * (v8 * (-2.25 - (float32(1e10) - -2.25))))
            }
        }
        _ = fi9
    }
    switch ("a" + (("foo bar" + ((((((("hello" + "") + ("\n\t" + "a")) + "") + ("hello" + "\n\t")) + "\n\t") + "hello") + ((("\n\t" + "hello") + ("hello" + (("foo bar" + "a") + "foo bar"))) + ("" + "hello")))) + ("\n\t" + "\n\t"))) {
    case "":
        v18 := GenS3{F0: 9223372036854775807, F1: true}
        _ = v18
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = -197
        }
    case "hello":
        var v19 int = -2147483648
        _ = v19
    }
    fmt.Println(v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
