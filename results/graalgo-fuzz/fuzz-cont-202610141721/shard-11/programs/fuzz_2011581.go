package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int16 }

func main() {
    var v1 float32 = ((0.1 * 1.5) / (1e10 - 3.14159))
    var v2 uint16 = 161
    var v3 int32 = -1
    var v4 uint = 0
    v5 := []string{"foo bar", "hello", "hello", "foo bar"}
    v6 := GenS1{F0: "foo bar"}
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: ("hello" + (("hello" + "a") + ((("\n\t" + "") + (("foo bar" + "") + "foo bar")) + "a")))}
        _ = fi7
        _ = v8
    }
    fmt.Println(v3)
    if (!(4294967295 == -200)) {
        v2 = 32767
    }
    if true {
        v9 := GenS2{F0: 808, F1: (true && ((100 <= 0) && (!false)))}
        _ = v9
    }
    v4 = (^(v4))
    v5 = append(v5, ((("hello" + "") + (("" + "") + "a")) + ""))
    v10 := GenS2{F0: 866, F1: (uint8(v2) < 170)}
    var v11 int32 = (v3 % ((v3) | 1))
    var v12 int8 = -128
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: ""}
        _ = len(v5)
        for fi15 := 0; fi15 < 8; fi15++ {
            if ((v12 <= v12) && (-1 <= v3)) {
                _ = len(v5)
                var v16 bool = ((1.5 <= 0.0) || false)
                _ = v16
            }
            var v17 uint16 = (v2 ^ v2)
            _ = fi15
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    if ((805 > 18446744073709551615) || false) {
        v11 -= v11
        v18 := GenS3{F0: -1}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
