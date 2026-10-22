package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 string }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 0
    var v3 string = ("hello" + "\n\t")
    v4 := GenS1{F0: 2, F1: 42, F2: ((v3 + "a") + v3)}
    if false {
        v3 = ("" + ((v3 + v3) + v3))
        for fi5 := 0; fi5 < 14; fi5++ {
            var v6 bool = ((v2 + (v2 * 0)) > (v2 * -645))
            switch fi5 {
            case 2147483647:
                var v7 int64 = 9223372036854775807
                _ = v7
            case 581:
                var v8 uint8 = 171
                _ = v8
            case 1:
                v9 := GenS2{F0: (^(fi5))}
                _ = v9
            }
            v6 = (576 < 0)
            _ = fi5
            _ = v6
        }
        v3 = ((((v3 + v3) + "\n\t") + v3) + "a")
    }
    v10 := GenS2{F0: -783}
    var v11 int32 = v2
    if (("foo bar" + v3) < "\n\t") {
        var v12 float32 = -2.25
        _ = v12
    }
    v2 = (v11 ^ (v2 - v2))
    var v13 int8 = 1
    if false {
        for fi14 := 0; fi14 < 7; fi14++ {
            fmt.Println(v11, v11, v2)
            _ = fi14
        }
        if false {
            v15 := v1[460]
            var v16 int16 = -749
            if false {
                var v17 uint8 = 2
                _ = v17
            } else {
                v13 = (v13 / ((v13) | 1))
            }
            _ = v15
            _ = v16
        } else {
            if true {
                var v18 bool = false
                v19 := GenS2{F0: -9223372036854775808}
                _ = v18
                _ = v19
            } else {
                v20 := GenS2{F0: 4294967295}
                v21 := []int{4294967295, -2147483648}
                _ = len(v1)
                _ = v20
                _ = v21
            }
        }
        fmt.Println(v3)
    }
    var v22 float32 = ((float32(((0.1 / (0.1 - (0.0 + (1e10 * 0.1)))) + 1.5)) - (0.0 - 0.0)) / 3.14159)
    if (!(!false)) {
        if (4294967295 != 870) {
            var v23 bool = (false && (!false))
            _ = v23
        }
        var v24 int = -725
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 int8 = 1
        _ = fi25
        _ = v26
    }
    fmt.Println(v11, v3, v13)
    var v27 bool = false
    var v28 int32 = (v11 >> ((1) & 63))
    v29 := GenS1{F0: 278, F1: 552, F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
