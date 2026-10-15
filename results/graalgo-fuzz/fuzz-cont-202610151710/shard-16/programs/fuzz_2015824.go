package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 string }
type GenS2 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint32 = 891
    v2 := make(map[uint32]int)
    var v3 uint8 = uint8(v1)
    var v4 int64 = 9223372036854775807
    v5 := []string{"", "\n\t", "a", "a"}
    var v6 int16 = 1000
    v4 = -1
    _ = len(v5)
    var v7 string = "hello"
    var v8 string = v7
    fmt.Println(v4, v1, v1)
    v4 = (-(v4))
    v7 = "a"
    if (!true) {
        var v9 uint = 2
        _ = v9
    }
    var v10 bool = true
    if len(v5) > 0 {
        v5[0] = (v8 + v8)
    }
    v11 := GenS1{F0: uint8(v1), F1: (((v7 + "hello") + "hello") != v8), F2: v8}
    v12 := GenS1{F0: (v3 - v3), F1: v10, F2: "hello"}
    if (!v10) {
        v10 = (1 >= 1)
        v13 := GenS1{F0: v3, F1: (false || (!v10)), F2: v7}
        switch 127 {
        case 9:
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 uint8 = 42
                _ = fi14
                _ = v15
            }
        }
        _ = v13
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        if (true && v10) {
            if v10 {
                var v17 uint = 4294967295
                _ = v17
            }
            v3 -= (^(v3))
        } else {
            for fi18 := 0; fi18 < 3; fi18++ {
                v3 -= 255
                v19 := GenS2{F0: (106 >= 1), F1: v3}
                _ = len(v5)
                _ = fi18
                _ = v19
            }
        }
        fmt.Println(v1, v10, fi16)
        _ = fi16
    }
    if len(v5) > 0 {
        v5[0] = ((v8 + "") + "hello")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
