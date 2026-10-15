package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 236
    var v3 string = ("foo bar" + "a")
    v4 := make(map[string]string)
    var v5 bool = ((0.1 / 1.0) > 1.5)
    var v6 uint = 735
    for fi7 := 0; fi7 < 16; fi7++ {
        if (1.5 != 1.5) {
            var v8 string = (("a" + v3) + (v3 + v3))
            _ = v8
        }
        fmt.Println(v3)
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := []int{(fi9 >> ((v6) & 63)), fi9, (-(fi9))}
            var v11 int8 = 0
            v3 = v3
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
    }
    v3 = ""
    fmt.Println(v5, v5)
    var v12 int32 = 0
    if (v5 || true) {
        var v13 float32 = 0.0
        _ = v13
    } else {
        var v14 float32 = 1.5
        _ = v14
    }
    var v15 uint16 = v2
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint32 = 872
        v18 := GenS1{F0: 93, F1: 0, F2: 874}
        switch (fi16 % ((fi16) | 1)) {
        case 473:
            v19 := v4["foo bar"]
            _ = v19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    if v5 {
        v5 = false
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 uint16 = 199
            if true {
                v22 := []int{}
                var v23 float32 = float32(1.5)
                fmt.Println(v12, v3, v3)
                _ = v22
                _ = v23
            }
            _ = fi20
            _ = v21
        }
    } else {
        var v24 uint8 = 255
        v12 -= (v12 / ((v12) | 1))
        _ = v24
    }
    fmt.Println(v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
