package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 string }
type GenS2 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint = 0
    v2 := []int{183, -9223372036854775808, 32768, 9223372036854775807}
    var v3 uint8 = 42
    v4 := []uint32{uint32(v3), 4294967295, 42, 4294967295}
    var v5 int32 = -2147483648
    _ = len(v2)
    v1 += (^(v1))
    v5 += 195
    if (!(!(true && false))) {
        switch 32768 {
        case 32768:
            v3 += (v3 / ((2) | 1))
        case -9223372036854775808:
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 int = (fi6 - fi6)
                for fi8 := 0; fi8 < 20; fi8++ {
                    _ = len(v4)
                    _ = fi8
                }
                _ = fi6
                _ = v7
            }
        default:
            v3 += v3
        }
        var v9 string = ("\n\t" + "hello")
        _ = v9
    }
    v5 += (v5 / ((996) | 1))
    v1 = 1
    if true {
        var v10 int = -2147483648
        v11 := GenS1{F0: v10, F1: ("a" + ""), F2: ((("\n\t" + "foo bar") + "a") + ("a" + (("hello" + "foo bar") + ("" + "\n\t"))))}
        _ = len(v2)
        _ = v10
        _ = v11
    }
    switch "hello" {
    case "a":
        if ("a" <= "") {
            var v12 int8 = int8(v1)
            _ = v12
        } else {
            var v13 int = 9223372036854775807
            switch v13 {
            case 866:
                fmt.Println(v1)
            case 616:
                var v14 uint64 = 18446744073709551615
                _ = v14
            default:
                v15 := GenS1{F0: (v13 + 9223372036854775807), F1: "\n\t", F2: "hello"}
                _ = v15
            }
            v16 := GenS2{F0: (v3 >> ((v1) & 63)), F1: 0}
            _ = v13
            _ = v16
        }
    case "":
        var v17 int = -2147483648
        _ = v17
    }
    var v18 uint = (v1 - 0)
    v19 := GenS1{F0: 0, F1: "foo bar", F2: ""}
    var v20 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
