package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int32 }

func main() {
    var v1 uint64 = 2
    var v2 string = "foo bar"
    var v3 int = 128
    v4 := make(map[int]string)
    v5 := []int{1, (v3 << ((255) & 63)), (v3 - (v3 >> ((979) & 63))), (v3 * (v3 ^ 548))}
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int8 = 127
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(v2, v2, v3)
            _ = fi10
        }
        _ = fi7
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if ((3.14159 != (-2.25 * 1.5)) && false) {
        if (true && (!true)) {
            for fi11 := 0; fi11 < 12; fi11++ {
                _ = len(v5)
                v12 := GenS1{F0: 42, F1: -116, F2: -475}
                var v13 uint8 = 128
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                v1 = v1
                v5 = append(v5, 32768)
                _ = fi14
            }
            v15 := GenS1{F0: 1, F1: 0, F2: -903}
            _ = v15
        }
    } else {
        fmt.Println(v1)
        var v16 bool = false
        _ = v16
    }
    var v17 uint64 = v1
    v1 = (v17 - (^(v17)))
    switch v3 {
    case -176:
        v6[1] = (v3 ^ (-(v3)))
    }
    switch (v2 + ("a" + v2)) {
    case "foo bar":
        _ = len(v6)
    }
    var v18 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
