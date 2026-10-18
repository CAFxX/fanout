package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }

func main() {
    v1 := []string{"foo bar"}
    var v2 int64 = 0
    v3 := []uint32{}
    var v4 int16 = -919
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 uint32 = 602
        fmt.Println(v2, v2)
        _ = fi5
        _ = v6
    }
    var v7 int16 = (v4 / ((v4) | 1))
    var v8 uint = 4294967295
    fmt.Println(v8, v2)
    v9 := GenS1{F0: (("\n\t" + "a") + ("hello" + "hello")), F1: -1}
    var v10 string = (("hello" + "\n\t") + "")
    if (!(-9223372036854775808 < -1)) {
        switch (v10 + v10) {
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v3 = append(v3, 328)
        v11 := GenS1{F0: v10, F1: 113}
        _ = v11
    }
    if ((v4 * v4) <= v7) {
        v10 = "\n\t"
        var v12 uint8 = 144
        _ = v12
    }
    var v13 uint16 = 0
    v2 = v2
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 int32 = 2147483647
            var v17 int8 = 127
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    v13 = (v13 * v13)
    var v18 uint = (^(v8))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
