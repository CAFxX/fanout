package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -9223372036854775808
    var v3 int16 = 349
    var v4 int = 32768
    var v5 uint = 0
    var v6 int8 = -1
    var v7 uint64 = 4294967295
    var v8 uint = (v5 / ((v5) | 1))
    if (2147483647 <= -957) {
        if true {
            delete(v1, "\n\t")
        } else {
            var v9 int32 = 1
            _ = v9
        }
        v4 -= (v4 - (v4 << (((^(v5))) & 63)))
        var v10 int8 = (v6 - (-(v6)))
        _ = v10
    } else {
        v5 = v5
        v11 := v1["a"]
        _ = len(v1)
        _ = v11
    }
    v6 -= (v6 + (v6 + v6))
    switch (("\n\t" + "foo bar") + "hello") {
    case "":
        v8 = v5
    case "\n\t":
        v12 := GenS1{F0: 515, F1: 42, F2: (v3 | 980)}
        _ = v12
    default:
        var v13 uint64 = 404
        _ = v13
    }
    v5 += (v5 ^ 255)
    if (1.5 == -2.25) {
        var v14 bool = true
        var v15 bool = v14
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int64 = -9223372036854775808
        var v18 int64 = v17
        var v19 string = ""
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    if (v8 != 18446744073709551615) {
        var v20 int8 = v6
        v21 := []int{}
        v6 = 100
        _ = v20
        _ = v21
    } else {
        var v22 float32 = 0.1
        for fi23 := 0; fi23 < 11; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 uint32 = 4294967295
                if false {
                    v26 := GenS2{F0: 9223372036854775807}
                    _ = v26
                }
                _ = fi24
                _ = v25
            }
            v22 = (1.5 + (v22 - -2.25))
            _ = fi23
        }
        for fi27 := 0; fi27 < 6; fi27++ {
            v28 := v1[("\n\t" + "")]
            _ = fi27
            _ = v28
        }
        _ = v22
    }
    var v29 bool = false
    v30 := GenS2{F0: (v4 + v4)}
    var v31 int64 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
