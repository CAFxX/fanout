package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 42
    var v3 string = ("hello" + "\n\t")
    switch 485 {
    case 980:
        v4 := []int{}
        _ = v4
    case 270:
        v3 = v3
    case -144:
        v2 -= (v2 ^ 42)
    }
    if false {
        v5 := GenS1{F0: (-1 <= 2147483647), F1: 4294967295, F2: 798}
        _ = v5
    } else {
        v2 = (^(v2))
        var v6 uint16 = 2
        if true {
            v7 := GenS1{F0: ((!true) && (329 != -9223372036854775808)), F1: 227, F2: 255}
            _ = v7
        } else {
            var v8 uint16 = (v6 | (v6 + v6))
            switch (v3 + "a") {
            case "\n\t":
                var v9 int16 = 332
                _ = v9
            }
            v3 = "hello"
            _ = v8
        }
        _ = v6
    }
    if (false && false) {
        if (!false) {
            v2 += v2
            fmt.Println(v2, v3)
            v3 = v3
        } else {
            var v10 int64 = -393
            v3 = "foo bar"
            _ = v10
        }
    }
    var v11 bool = true
    delete(v1, 42)
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 float32 = (((float32(1e10) * 1.5) - (0.0 * float32((0.0 - 0.1)))) - 1.5)
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: v11, F1: 18446744073709551615, F2: 420}
    for fi15 := 0; fi15 < 10; fi15++ {
        v16 := []int{-2147483648, (fi15 | 32768), fi15, (fi15 & fi15)}
        _ = fi15
        _ = v16
    }
    var v17 uint64 = 42
    v17 = (v17 >> ((382) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v3
}
