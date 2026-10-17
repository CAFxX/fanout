package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
func genFn0(p0 int16, p1 uint, p2 uint16) int32 {
    var v7 uint = (p1 * (p1 << (((p1 + (p1 * p1))) & 63)))
    fmt.Println(v7)
    var v8 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 684
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]bool)
    var v3 uint = 2
    var v4 int64 = -1
    var v5 uint64 = 1
    var v6 float64 = 0.1
    fmt.Println(v5, v5)
    var v9 uint64 = v5
    if (!(!(1 > 55))) {
        var v10 string = ("a" + (((("hello" + "\n\t") + "\n\t") + "") + ("" + "\n\t")))
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v9 += (^(v5))
        if false {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int32 = 2147483647
                _ = fi12
                _ = v13
            }
        }
        v1 = append(v1, 4294967295)
        _ = fi11
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int16 = -138
        if false {
            var v16 uint8 = 127
            _ = v16
        } else {
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 int16 = 1000
                v9 += 160
                v5 = (^(v9))
                _ = fi17
                _ = v18
            }
            v19 := GenS1{F0: (v15 % (((v15 << ((v3) & 63))) | 1)), F1: ("a" + ("foo bar" + "")), F2: v15}
            fmt.Println(v5)
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    v20 := GenS1{F0: 1, F1: (((("hello" + ("" + "")) + "foo bar") + "") + ""), F2: 0}
    if true {
        v21 := GenS1{F0: 32767, F1: ("" + ((("\n\t" + "foo bar") + "foo bar") + "foo bar")), F2: 758}
        fmt.Println(v6, v5)
        v4 = (v4 << (((v3 - v3)) & 63))
        _ = v21
    }
    var v22 int64 = int64(v3)
    v23 := v2["foo bar"]
    fmt.Println(v4, v23)
    v24 := GenS1{F0: -32768, F1: ("\n\t" + "a"), F2: 42}
    v25 := GenS1{F0: 32767, F1: "hello", F2: -1}
    v2[""] = true
    var v26 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
