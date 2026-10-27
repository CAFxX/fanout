package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 string; F2 int32 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]int)
    var v3 bool = false
    var v4 uint64 = 159
    var v5 float32 = 3.14159
    if ((v4 == v4) && (!false)) {
        v6 := []int{}
        var v7 uint16 = 1
        if (v3 || (false || ((v1 << ((255) & 63)) >= v1))) {
            var v8 int8 = 1
            _ = v8
        }
        _ = v6
        _ = v7
    }
    if (true && (((-9223372036854775808 <= -1) && ((false || (!(!(42 <= 128)))) && (!false))) || (0 <= 18446744073709551615))) {
        v9 := GenS1{F0: 375, F1: 65535, F2: (!v3)}
        switch "hello" {
        case "":
            var v10 int64 = -9223372036854775808
            _ = v10
        case "foo bar":
            var v11 uint16 = 154
            _ = v11
        }
        _ = v9
    } else {
        switch "a" {
        case "foo bar":
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := []int{fi12, (fi12 + fi12), fi12, (fi12 * fi12)}
                switch "" {
                case "a":
                    for fi14 := 0; fi14 < 8; fi14++ {
                        v3 = v3
                        v5 = v5
                        _ = fi14
                    }
                case "\n\t":
                    delete(v2, "foo bar")
                default:
                    v4 -= v4
                }
                _ = fi12
                _ = v13
            }
        case "hello":
            v15 := GenS2{F0: v1}
            _ = v15
        case "a":
            v3 = v3
        }
        var v16 int = int(v1)
        _ = v16
    }
    fmt.Println(v1)
    switch "\n\t" {
    case "":
        v1 = (^(v1))
    default:
        v3 = v3
    }
    v17 := GenS3{F0: v4, F1: "a", F2: -450}
    v18 := GenS4{F0: -8}
    if (v3 && false) {
        v19 := GenS4{F0: -580}
        v20 := v2["foo bar"]
        _ = v19
        _ = v20
    } else {
        v4 = v4
    }
    switch ("hello" + (("" + "") + "a")) {
    case "\n\t":
        v5 = 0.0
    case "foo bar":
        delete(v2, (("foo bar" + "foo bar") + ""))
    case "hello":
        switch 1 {
        case 9223372036854775807:
            _ = len(v2)
        }
    default:
        v21 := []int{127}
        _ = v21
    }
    if (v3 || (855 == -1)) {
        var v22 string = ("\n\t" + "")
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 int32 = -2147483648
            _ = len(v2)
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    fmt.Println(v1)
    var v25 uint64 = v4
    var v26 int32 = -852
    v1 += v1
    fmt.Println(v4, v26)
    v27 := GenS4{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
