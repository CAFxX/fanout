package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 int = 42
    v2 := make(map[string]string)
    var v3 float32 = 1.5
    var v4 float32 = -2.25
    v5 := make(map[string]bool)
    v3 = v4
    if (!(4294967295 < 2)) {
        v6 := GenS1{F0: 1}
        switch v1 {
        case 2147483647:
            var v7 uint32 = 2147483648
            _ = v7
        case 1:
            v1 -= -162
        case 9223372036854775807:
            v1 = v1
        }
        var v8 float32 = v3
        _ = v6
        _ = v8
    }
    fmt.Println(v4, v3)
    v9 := []int{v1, 42, v1}
    fmt.Println(v4, v4)
    if false {
        v10 := GenS2{F0: 18446744073709551615}
        switch "hello" {
        case "":
            for fi11 := 0; fi11 < 6; fi11++ {
                if (((!(((-2.25 <= (3.14159 * (0.0 * (3.14159 - -2.25)))) && (!(((v3 + v4) > 0.1) || ((fi11 >> ((18446744073709551615) & 63)) >= (-(fi11)))))) || false)) && true) || (((("foo bar" + ("\n\t" + "hello")) + "") + "a") >= "hello")) {
                    v12 := GenS1{F0: 424}
                    var v13 int8 = 127
                    _ = v12
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        _ = len(v9)
                        _ = fi14
                    }
                }
                v15 := GenS1{F0: -32}
                var v16 uint16 = 160
                _ = fi11
                _ = v15
                _ = v16
            }
        }
        _ = v10
    } else {
        v17 := GenS3{F0: (true || true), F1: 0, F2: 2}
        var v18 uint32 = 515
        v19 := GenS1{F0: 9223372036854775807}
        _ = v17
        _ = v18
        _ = v19
    }
    switch 4294967295 {
    case -2147483648:
        v20 := v5[""]
        _ = v20
    case -844:
        var v21 uint8 = 127
        _ = v21
    case 32768:
        var v22 string = "foo bar"
        _ = v22
    default:
        v23 := GenS3{F0: false, F1: 726, F2: 42}
        _ = v23
    }
    v24 := GenS4{F0: 42}
    v25 := GenS2{F0: 1}
    var v26 uint = 165
    var v27 int = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
