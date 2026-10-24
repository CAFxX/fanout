package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint16 }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 42
    var v3 float32 = 1e10
    v4 := make(map[int]uint32)
    var v5 int16 = 827
    var v6 uint = 42
    var v7 int8 = 42
    if false {
        var v8 float32 = v3
        _ = v8
    } else {
        v9 := GenS1{F0: (("a" + "a") + ""), F1: v5, F2: 1}
        _ = v9
    }
    var v10 uint32 = 2147483648
    var v11 int64 = -1
    v12 := GenS1{F0: "", F1: (v5 % (((^(v5))) | 1)), F2: 65535}
    v13 := GenS2{F0: (false && ((((v5 / ((-1) | 1)) >= (v5 / ((v5) | 1))) && ("\n\t" <= ("" + "foo bar"))) || (false || false))), F1: (v5 << ((65535) & 63))}
    fmt.Println(v2, v6)
    switch ((("foo bar" + ("\n\t" + "a")) + "a") + "\n\t") {
    case "a":
        v7 -= 127
    default:
        if true {
            fmt.Println(v3)
        } else {
            fmt.Println(v3, v5, v6)
            v5 = (-(v5))
            if ((false && ((239 > 4294967295) && ((0 <= 2) || (2 == 71)))) && true) {
                var v14 uint64 = 2
                var v15 uint32 = uint32(v14)
                _ = v14
                _ = v15
            }
        }
    }
    var v16 bool = true
    fmt.Println(v6)
    v5 += int16(v7)
    v2 += (^(v2))
    switch "foo bar" {
    case "\n\t":
        v2 = 214
    }
    var v17 float64 = (0.0 - (float64(v3) - 0.1))
    v7 += (v7 | v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
