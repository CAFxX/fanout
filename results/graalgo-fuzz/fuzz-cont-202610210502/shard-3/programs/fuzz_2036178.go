package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{}
    v3 := make(map[uint32]bool)
    var v4 int64 = -661
    var v5 int32 = 2147483647
    v6 := GenS1{F0: 1, F1: "hello", F2: -1}
    v5 += (v5 + 42)
    var v7 int32 = (-(v5))
    v8 := GenS1{F0: 18446744073709551615, F1: "foo bar", F2: v4}
    switch -1 {
    case -2147483648:
        _ = len(v2)
    }
    if len(v2) > 0 {
        v2[0] = -226
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 uint32 = 1
            _ = fi10
            _ = v11
        }
        v5 -= (v5 / ((-2147483648) | 1))
        fmt.Println(v5)
        _ = fi9
    }
    var v12 int8 = 127
    for fi13 := 0; fi13 < 10; fi13++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v14 float64 = float64(1.5)
        var v15 float32 = 0.0
        _ = fi13
        _ = v14
        _ = v15
    }
    v5 -= -53
    var v16 uint16 = 65535
    v17 := GenS1{F0: 237, F1: "", F2: (v4 >> ((1) & 63))}
    v5 = (v7 + (v7 << ((2) & 63)))
    var v18 string = "a"
    switch (v18 + v18) {
    case "":
        for fi19 := 0; fi19 < 2; fi19++ {
            v4 += v4
            v12 += v12
            switch v18 {
            case "a":
                var v20 bool = true
                _ = v20
            }
            _ = fi19
        }
    case "\n\t":
        for fi21 := 0; fi21 < 4; fi21++ {
            if (!true) {
                v3[956] = (!((1.5 / 1e10) != 0.1))
                switch (v18 + "hello") {
                case "":
                    var v22 int32 = v5
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 8; fi23++ {
                v18 = ("foo bar" + ("\n\t" + ""))
                var v24 uint16 = 0
                _ = fi23
                _ = v24
            }
            var v25 int8 = v12
            _ = fi21
            _ = v25
        }
    case "foo bar":
        var v26 int64 = (^(v4))
        _ = v26
    }
    fmt.Println(v5, v16)
    v5 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
