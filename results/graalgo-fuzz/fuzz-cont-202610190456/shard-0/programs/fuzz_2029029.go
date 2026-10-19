package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() int64 {
    var v6 uint = 255
    v6 += v6
    _ = v6
    return 0
}

func main() {
    var v1 int = 0
    var v2 uint8 = 128
    var v3 int32 = 2147483647
    var v4 uint = 2
    v5 := []string{"a", ("hello" + "a"), "foo bar", "\n\t"}
    if (480 <= 65535) {
        var v7 int = 2147483647
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS1{F0: (fi8 | v7)}
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        if len(v5) > 0 {
            v5[0] = ("a" + "a")
        }
        switch fi10 {
        case -2147483648:
            v4 = 42
        }
        _ = fi10
    }
    var v11 bool = (42 != 248)
    var v12 uint16 = 491
    var v13 int = int(v3)
    var v14 uint64 = 433
    v13 -= v13
    v15 := GenS1{F0: 9223372036854775807}
    for fi16 := 0; fi16 < 11; fi16++ {
        v12 = (v12 ^ (v12 << ((190) & 63)))
        _ = fi16
    }
    v17 := GenS1{F0: (v13 + (^(v13)))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
