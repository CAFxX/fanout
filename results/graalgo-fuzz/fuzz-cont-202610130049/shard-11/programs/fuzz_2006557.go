package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int16 = 512
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := GenS1{F0: ((9223372036854775807 == 9223372036854775807) || true)}
    delete(v2, 432)
    var v5 bool = false
    v1 = v1
    for fi6 := 0; fi6 < 7; fi6++ {
        v1 -= 1000
        var v7 int64 = -1
        var v8 int32 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        fmt.Println(v5)
        switch 127 {
        case 32768:
            var v9 int16 = (v1 - v1)
            _ = v9
        case 2147483647:
            if v5 {
                v5 = true
                v1 += (v1 / ((v1) | 1))
            }
        case 1:
            var v10 float32 = 1e10
            _ = v10
        }
        switch "\n\t" {
        case "foo bar":
            for fi11 := 0; fi11 < 16; fi11++ {
                delete(v2, 0)
                v5 = (true || (1 <= 2))
                _ = fi11
            }
        case "":
            v12 := []int{}
            _ = v12
        }
    } else {
        if ((0 > 65535) || true) {
            v5 = true
        } else {
            v13 := v3[""]
            v5 = (((v13 - (v13 * v13)) < 2) || (1e10 >= 1.5))
            var v14 int8 = -1
            _ = v13
            _ = v14
        }
        var v15 uint8 = 2
        v16 := []int{}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 float32 = -2.25
            v19 = (0.0 * 3.14159)
            var v20 float64 = 1.5
            _ = fi18
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := GenS1{F0: (42 < 129)}
            _ = fi21
            _ = v22
        }
        v5 = v5
        _ = fi17
    }
    if (377 != 543) {
        var v23 float32 = (3.14159 / 0.1)
        _ = v23
    } else {
        if v5 {
            v24 := GenS1{F0: v5}
            v1 = v1
            _ = v24
        } else {
            fmt.Println(v5, v1)
            for fi25 := 0; fi25 < 13; fi25++ {
                fmt.Println(fi25, v5, fi25)
                _ = fi25
            }
        }
    }
    if (!(v5 || true)) {
        var v26 float32 = 1.5
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
