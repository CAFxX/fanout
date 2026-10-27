package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
func genFn0() uint {
    switch ("a" + ("foo bar" + "a")) {
    case "a":
        var v6 uint = 65535
        _ = v6
    case "":
        var v7 uint16 = 729
        _ = v7
    case "foo bar":
        var v8 uint = 255
        _ = v8
    }
    var v9 int32 = -1
    v9 += (v9 - v9)
    v10 := []int{943, 32768, 2147483647, -1}
    _ = len(v10)
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 int8 = -128
    var v2 int = 0
    v3 := make(map[string]int)
    var v4 uint32 = 113
    v5 := []uint32{65535}
    _ = len(v5)
    var v11 uint64 = 42
    v3["foo bar"] = (v2 - 0)
    v12 := GenS1{F0: 9223372036854775807, F1: 42}
    var v13 int16 = -597
    fmt.Println(v4, v4)
    v2 -= (v2 | (v2 + (v2 << ((65535) & 63))))
    for fi14 := 0; fi14 < 5; fi14++ {
        if len(v5) > 0 {
            v5[0] = v4
        }
        fmt.Println(fi14)
        fmt.Println(fi14, v2)
        _ = fi14
    }
    fmt.Println(v1, v2)
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int16 = (^(v13))
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: -109, F1: 255}
    if true {
        if false {
            var v18 uint32 = (v4 - 867)
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 string = "foo bar"
        _ = fi19
        _ = v20
    }
    var v21 int64 = 1
    if (1e10 > 0.0) {
        v5 = append(v5, (v4 * v4))
    } else {
        var v22 uint64 = 18446744073709551615
        switch (((("\n\t" + ("foo bar" + "foo bar")) + "\n\t") + "hello") + "a") {
        case "foo bar":
            v13 += (v13 & (v13 - (^(v13))))
        case "\n\t":
            _ = len(v3)
        }
        _ = v22
    }
    fmt.Println(v21, v4, v21)
    v13 = v13
    for fi23 := 0; fi23 < 20; fi23++ {
        v24 := GenS1{F0: (v21 - 671), F1: 296}
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint8 = 2
            switch "" {
            case "foo bar":
                var v27 int32 = 155
                _ = v27
            case "\n\t":
                fmt.Println(fi23, fi23, v21)
            case "hello":
                v28 := GenS1{F0: (v21 ^ -1), F1: 867}
                _ = v28
            }
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
