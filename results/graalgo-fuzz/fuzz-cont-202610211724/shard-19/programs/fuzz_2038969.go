package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 string }
type GenS2 struct { F0 int16; F1 uint32; F2 uint }

func main() {
    v1 := []string{("" + "a"), "a", "\n\t", ("a" + "a")}
    var v2 float32 = -2.25
    v3 := make(map[int]uint32)
    v4 := []uint32{1, 18, 42, 86}
    var v5 int64 = -9223372036854775808
    _ = len(v3)
    fmt.Println(v2, v2)
    v5 += v5
    v2 += ((v2 - ((v2 + v2) * (((v2 / ((v2 * -2.25) / v2)) / v2) / -2.25))) * v2)
    fmt.Println(v2, v2)
    for fi6 := 0; fi6 < 3; fi6++ {
        v2 += v2
        v7 := GenS1{F0: 1, F1: ("\n\t" + "foo bar"), F2: "a"}
        _ = len(v3)
        _ = fi6
        _ = v7
    }
    if (1 != 0) {
        v2 = 0.1
    } else {
        var v8 uint = 18446744073709551615
        v9 := GenS2{F0: 1, F1: 65535, F2: v8}
        v5 = v5
        _ = v8
        _ = v9
    }
    switch (("hello" + ((("foo bar" + "\n\t") + ("" + (("hello" + ("hello" + "\n\t")) + "foo bar"))) + "a")) + ("hello" + "foo bar")) {
    case "":
        var v10 uint64 = 4294967295
        _ = v10
    case "a":
        var v11 float64 = float64(0.1)
        _ = v11
    default:
        v3[-950] = 65535
    }
    if (false || false) {
        switch ("" + ("hello" + "a")) {
        case "":
            for fi12 := 0; fi12 < 6; fi12++ {
                if (true || false) {
                    v2 = 1e10
                    var v13 int16 = 0
                    _ = v13
                }
                _ = fi12
            }
        case "\n\t":
            v14 := GenS2{F0: 1, F1: 2, F2: 18446744073709551615}
            _ = v14
        default:
            var v15 uint16 = 1
            _ = v15
        }
        delete(v3, 1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
