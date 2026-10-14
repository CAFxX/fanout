package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int }

func main() {
    var v1 uint32 = 65535
    var v2 bool = ((0.0 / 1e10) > 0.1)
    var v3 int64 = 719
    v4 := make(map[uint32]bool)
    var v5 int64 = -1
    v3 = (v5 ^ (v5 * (v5 * (v3 >> ((1) & 63)))))
    v2 = v2
    fmt.Println(v3, v3)
    v4[(v1 >> ((65535) & 63))] = ((v1 >> ((18446744073709551615) & 63)) < v1)
    v6 := GenS1{F0: 66}
    v7 := []int{4294967295}
    v2 = v2
    v8 := GenS1{F0: 2}
    v5 = v3
    v9 := GenS1{F0: v1}
    switch 0 {
    case 9223372036854775807:
        switch 2147483647 {
        case 9223372036854775807:
            v10 := GenS1{F0: (v1 - (v1 >> ((515) & 63)))}
            _ = v10
        default:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
    case -2147483648:
        var v11 uint8 = 174
        _ = v11
    case 4294967295:
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 int = (fi12 | (-(fi12)))
            v14 := GenS2{F0: ((("\n\t" + ("" + "\n\t")) + ("" + "\n\t")) + "\n\t"), F1: fi12}
            _ = fi12
            _ = v13
            _ = v14
        }
    default:
        switch ("hello" + ("foo bar" + "\n\t")) {
        case "foo bar":
            switch ("hello" + "foo bar") {
            case "foo bar":
                v1 = 2147483647
            case "\n\t":
                v2 = (0.1 >= 1.5)
            }
        case "a":
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := v4[2147483647]
                _ = fi15
                _ = v16
            }
        case "hello":
            var v17 int8 = -86
            _ = v17
        }
    }
    fmt.Println(v1, v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
