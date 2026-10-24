package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]int)
    var v3 bool = (238 != 127)
    for fi4 := 0; fi4 < 12; fi4++ {
        v3 = (-1 > 2147483647)
        for fi5 := 0; fi5 < 10; fi5++ {
            v3 = (0 >= 2)
            var v6 uint8 = 1
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v3 = v3
    v1 -= (v1 << ((18446744073709551615) & 63))
    v1 += (v1 / (((v1 % (((^(v1))) | 1))) | 1))
    v7 := GenS1{F0: -1, F1: 25, F2: -1}
    fmt.Println(v3, v3)
    v8 := GenS1{F0: -32768, F1: 100, F2: -2147483648}
    v3 = true
    v3 = v3
    var v9 float32 = float32((0.0 + 0.1))
    for fi10 := 0; fi10 < 7; fi10++ {
        v9 = (float32(-2.25) * v9)
        if true {
            var v11 float64 = float64((v9 + 0.0))
            _ = v11
        } else {
            if (!(true && false)) {
                var v12 int8 = 100
                _ = v12
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v14 := GenS1{F0: -1, F1: 1, F2: -1}
                    _ = fi13
                    _ = v14
                }
                switch "foo bar" {
                case "":
                    v9 = 1e10
                case "\n\t":
                    v15 := GenS1{F0: -279, F1: -42, F2: 1}
                    _ = v15
                default:
                    if (false || true) {
                        _ = len(v2)
                    }
                }
                v16 := []int{fi10}
                _ = v16
            }
            v9 -= (v9 + 0.1)
        }
        v3 = (42 != 0)
        _ = fi10
    }
    v9 -= (v9 - v9)
    if ("" != "\n\t") {
        switch -645 {
        case 637:
            var v17 int64 = 0
            _ = v17
        case -957:
            v1 = (v1 | v1)
        case 32767:
            switch (("a" + "\n\t") + (("foo bar" + "hello") + "hello")) {
            case "foo bar":
                for fi18 := 0; fi18 < 6; fi18++ {
                    _ = len(v2)
                    _ = fi18
                }
            }
        }
    } else {
        v19 := []int{218, 32768}
        v1 = (^(v1))
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
