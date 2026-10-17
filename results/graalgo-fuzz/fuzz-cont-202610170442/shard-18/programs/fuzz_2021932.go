package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int }
type GenS2 struct { F0 int }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]bool)
    var v3 int = 485
    var v4 int64 = 9223372036854775807
    var v5 int8 = 64
    var v6 uint8 = 1
    v7 := GenS1{F0: v5, F1: "foo bar", F2: 1}
    switch "" {
    case "":
        var v8 int64 = (v4 << ((1) & 63))
        _ = v8
    case "a":
        var v9 uint = 4294967295
        _ = v9
    }
    var v10 string = ("" + "foo bar")
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 int16 = v1
            var v14 int = fi11
            v4 = v4
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    fmt.Println(v3)
    if false {
        if (!false) {
            v3 -= (-(v3))
        }
        if true {
            fmt.Println(v10, v1, v3)
        } else {
            v15 := []int{(-(v3))}
            _ = v15
        }
    }
    var v16 uint32 = 782
    v1 = -32768
    v17 := v2[v3]
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            if (false || ((!(255 == 1)) || v17)) {
                v20 := []int{(fi19 / ((128) | 1)), fi18, (fi18 / ((fi18) | 1))}
                v21 := v2[v3]
                _ = v20
                _ = v21
            }
            var v22 bool = (v16 == (v16 / (((v16 + v16)) | 1)))
            _ = fi19
            _ = v22
        }
        _ = fi18
    }
    v23 := GenS2{F0: v3}
    if (903 >= 42) {
        v24 := GenS2{F0: v3}
        var v25 string = ""
        _ = v24
        _ = v25
    } else {
        v5 -= (^(v5))
        v26 := GenS2{F0: int(v4)}
        _ = v26
    }
    v3 = (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
