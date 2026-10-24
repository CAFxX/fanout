package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{}
    var v3 int32 = 0
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 19; fi5++ {
        if len(v2) > 0 {
            v2[0] = 652
        }
        v3 -= v3
        _ = fi5
    }
    _ = len(v4)
    var v6 bool = (!(((^(v3)) == (^(v3))) || (4294967295 < 65535)))
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int32 = 2147483647
        var v9 int8 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v1 = 42
        _ = fi10
    }
    v11 := GenS1{F0: "\n\t"}
    _ = len(v4)
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(fi12)
        _ = len(v2)
        _ = fi12
    }
    v2 = append(v2, 1)
    v13 := GenS1{F0: "\n\t"}
    var v14 int64 = v1
    v3 -= (v3 % (((v3 ^ (v3 - (-(v3))))) | 1))
    if (3.14159 >= float64(1e10)) {
        if v6 {
            v14 -= v1
            v15 := GenS2{F0: 1000}
            _ = v15
        } else {
            v16 := GenS1{F0: ("a" + (("a" + ("foo bar" + "foo bar")) + ("foo bar" + "\n\t")))}
            var v17 uint = 654
            _ = v16
            _ = v17
        }
    } else {
        var v18 int8 = 1
        fmt.Println(v3)
        _ = v18
    }
    v6 = (v6 && (v6 || (-1 < 21)))
    var v19 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
