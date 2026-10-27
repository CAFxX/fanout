package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int16 = 0
    v2 := []int{0, 4294967295}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    var v5 bool = (false || true)
    var v6 uint16 = 595
    if (2147483647 >= -9223372036854775808) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (!v5) {
            var v7 float32 = 1e10
            _ = v7
        }
    }
    switch "foo bar" {
    case "":
        v8 := GenS1{F0: v6, F1: "a", F2: 181}
        _ = v8
    }
    var v9 uint32 = 1
    v10 := GenS2{F0: -749}
    var v11 uint = 551
    v12 := GenS1{F0: (^(v6)), F1: ("hello" + ("foo bar" + "hello")), F2: 42}
    var v13 string = "foo bar"
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 int64 = -9223372036854775808
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            v11 -= v11
            _ = fi17
        }
        v11 = (^(v11))
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
