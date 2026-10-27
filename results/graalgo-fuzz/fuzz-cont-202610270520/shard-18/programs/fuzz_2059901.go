package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []string{("hello" + (("" + "hello") + "hello")), "\n\t"}
    var v2 int = 9223372036854775807
    var v3 uint32 = 0
    v4 := GenS1{F0: "\n\t", F1: 1}
    v3 -= (^(v3))
    v5 := GenS1{F0: ("" + (("foo bar" + ("a" + (("foo bar" + "") + "\n\t"))) + ((("\n\t" + "\n\t") + ((("\n\t" + "foo bar") + (("" + "\n\t") + "")) + "")) + "hello"))), F1: 221}
    v6 := GenS1{F0: (("a" + ("a" + ("" + ""))) + ("\n\t" + "")), F1: 18446744073709551615}
    for fi7 := 0; fi7 < 15; fi7++ {
        v3 += 275
        fmt.Println(v3)
        _ = len(v1)
        _ = fi7
    }
    var v8 uint = 255
    _ = len(v1)
    v3 += 2147483647
    v9 := GenS1{F0: ("\n\t" + ("a" + "hello")), F1: v8}
    v2 = (-(v2))
    v10 := GenS1{F0: "foo bar", F1: 357}
    if (true && true) {
        switch (v2 | (-(v2))) {
        case 32767:
            for fi11 := 0; fi11 < 12; fi11++ {
                v12 := GenS2{F0: 4294967295, F1: 32767}
                var v13 int32 = 559
                if ((^(v8)) != v8) {
                    switch ("" + "foo bar") {
                    case "hello":
                        v13 += v13
                    case "foo bar":
                        fmt.Println(fi11, v3, fi11)
                    case "a":
                        var v14 uint = v8
                        _ = v14
                    }
                    if (1e10 == 0.0) {
                        var v15 int = (v2 / ((fi11) | 1))
                        _ = v15
                    }
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        case 1:
            v3 += (v3 & (v3 | uint32(v8)))
        case -548:
            fmt.Println(v3, v8, v8)
        default:
            v2 = 32767
        }
    } else {
        v16 := GenS3{F0: v3, F1: (-(v2))}
        for fi17 := 0; fi17 < 6; fi17++ {
            v3 -= (^(v3))
            var v18 uint8 = 2
            v19 := GenS2{F0: 4294967295, F1: 287}
            _ = fi17
            _ = v18
            _ = v19
        }
        v1 = append(v1, ("\n\t" + ((("\n\t" + ("a" + "")) + "a") + ("a" + "a"))))
        _ = v16
    }
    var v20 int64 = 42
    _ = len(v1)
    var v21 bool = (true || (false || (float64((-2.25 * 0.0)) == 3.14159)))
    v22 := GenS4{F0: -436}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
