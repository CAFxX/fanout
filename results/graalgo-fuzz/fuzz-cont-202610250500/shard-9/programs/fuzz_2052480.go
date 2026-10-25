package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }

func main() {
    v1 := []int{-1, 0, -707, 127}
    v2 := []string{("\n\t" + ("hello" + "foo bar")), ("hello" + "foo bar")}
    var v3 bool = (("\n\t" != ("foo bar" + (("hello" + "hello") + ("" + "a")))) && ("foo bar" >= ""))
    var v4 uint64 = 0
    var v5 bool = v3
    v5 = true
    fmt.Println(v5, v3, v3)
    var v6 uint16 = 65535
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint16 = (^(v6))
        for fi9 := 0; fi9 < 7; fi9++ {
            v2 = append(v2, (("foo bar" + "hello") + ("hello" + (("foo bar" + ("\n\t" + "\n\t")) + "hello"))))
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    var v10 uint8 = 73
    var v11 uint = 2
    var v12 float64 = 1.5
    if ((v4 ^ 18446744073709551615) != (^(v4))) {
        fmt.Println(v6, v4)
        var v13 int64 = 1
        v14 := GenS1{F0: 1000, F1: (v11 << ((v11) & 63))}
        _ = v13
        _ = v14
    } else {
        var v15 uint32 = 241
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v6 = (v6 + v6)
        v11 += (v11 >> (((^(v11))) & 63))
        _ = fi16
    }
    var v17 uint8 = (v10 & v10)
    if true {
        var v18 float64 = v12
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
        _ = v18
    }
    if v3 {
        v19 := GenS1{F0: uint16(v10), F1: 896}
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
