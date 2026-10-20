package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 string }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int8; F1 uint; F2 string }
func genFn0(p0 bool) int {
    var v4 uint = 18446744073709551615
    v4 = (^(v4))
    _ = p0
    _ = v4
    return 127
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 65535
    var v3 int64 = -9223372036854775808
    v1 -= v1
    var v5 int32 = 1
    v6 := GenS1{F0: (v5 >> ((2) & 63)), F1: "hello", F2: ("foo bar" + "hello")}
    v7 := GenS1{F0: -1, F1: "hello", F2: (("\n\t" + "\n\t") + "")}
    v8 := GenS1{F0: (v5 & v5), F1: "", F2: ""}
    v5 = (v5 ^ 2147483647)
    var v9 int = 219
    v10 := GenS1{F0: (v5 % (((v5 & (v5 % ((2147483647) | 1)))) | 1)), F1: ("\n\t" + ((("\n\t" + ("hello" + "")) + "hello") + "foo bar")), F2: "\n\t"}
    v3 = int64(v9)
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v2, v9, v3)
        v5 = (^(v5))
        _ = fi11
    }
    v3 = 878
    v12 := []int{v9, 1}
    v1 += (v1 - (v1 * 42))
    var v13 uint8 = 161
    fmt.Println(v1)
    v14 := GenS2{F0: (v5 + (v5 * 42)), F1: -642}
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS2{F0: 2147483647, F1: -1}
        switch "\n\t" {
        case "a":
            fmt.Println(v2)
        case "foo bar":
            if (!(810 != 124)) {
                v1 = -128
            } else {
                v3 = (v3 | -911)
                var v17 int8 = v1
                _ = v17
            }
        default:
            v18 := GenS3{F0: v1, F1: 0, F2: ("foo bar" + ("foo bar" + "hello"))}
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
