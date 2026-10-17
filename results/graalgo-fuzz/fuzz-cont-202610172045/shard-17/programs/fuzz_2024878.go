package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 65535
    var v3 string = "foo bar"
    v4 := []string{v3, "foo bar", ((v3 + v3) + v3), "\n\t"}
    switch v3 {
    case "foo bar":
        v5 := GenS1{F0: -9223372036854775808, F1: 1, F2: 652}
        _ = v5
    case "\n\t":
        fmt.Println(v2)
    }
    v6 := GenS1{F0: -9223372036854775808, F1: 255, F2: 0}
    if (!false) {
        _ = len(v1)
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        if false {
            switch fi7 {
            case -524:
                v8 := GenS1{F0: 0, F1: 42, F2: 42}
                _ = v8
            case -919:
                if false {
                    v4 = append(v4, (v3 + ("a" + v3)))
                }
            }
            v9 := v1[267]
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint = 42
        v3 = v3
        var v12 float32 = (1.5 * (-2.25 / (1e10 / ((1e10 / 1e10) - (((1e10 - -2.25) - 1e10) * -2.25)))))
        _ = fi10
        _ = v11
        _ = v12
    }
    switch 397 {
    case 0:
        if (true || (false || false)) {
            var v13 uint16 = (v2 >> ((65535) & 63))
            var v14 string = ""
            _ = v13
            _ = v14
        } else {
            v15 := GenS1{F0: 97, F1: 69, F2: 65535}
            var v16 uint32 = 42
            v17 := GenS1{F0: 0, F1: 255, F2: 4294967295}
            _ = v15
            _ = v16
            _ = v17
        }
    case 9223372036854775807:
        fmt.Println(v3)
    case 2147483647:
        if ((823 < 348) || ((true && (683 < 383)) && (!false))) {
            v2 -= (v2 % ((v2) | 1))
        }
    default:
        v18 := GenS1{F0: 42, F1: 75, F2: 18446744073709551615}
        _ = v18
    }
    v19 := GenS1{F0: -254, F1: 123, F2: 4294967295}
    v20 := GenS1{F0: 42, F1: 0, F2: 255}
    v2 -= v2
    if false {
        v21 := GenS1{F0: -29, F1: 0, F2: 1}
        _ = v21
    }
    v2 = (v2 - 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
}
