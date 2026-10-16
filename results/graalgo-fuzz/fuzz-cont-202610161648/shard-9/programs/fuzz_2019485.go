package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint64; F2 bool }

func main() {
    v1 := []uint32{1}
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 int8 = -128
    var v5 float32 = 0.0
    v4 = v4
    var v6 uint16 = 65535
    v4 += v4
    v6 = 32767
    if false {
        v4 = -128
        if true {
            var v7 int16 = 1000
            if (false || (-1 != 32767)) {
                fmt.Println(v4)
                switch 4294967295 {
                case 1:
                    var v8 uint64 = 1
                    _ = v8
                }
            } else {
                v5 = (0.0 - v5)
            }
            _ = v7
        }
    } else {
        v5 = (v5 * v5)
        v3[65535] = 0
        v9 := GenS1{F0: (^(v6))}
        _ = v9
    }
    switch (("foo bar" + "a") + (("foo bar" + ("" + "hello")) + "a")) {
    case "":
        for fi10 := 0; fi10 < 11; fi10++ {
            delete(v2, 2147483648)
            if false {
                v11 := v3[182]
                _ = v11
            } else {
                var v12 int64 = -1
                _ = v12
            }
            v6 = v6
            _ = fi10
        }
    case "foo bar":
        v13 := GenS2{F0: 1, F1: 149, F2: true}
        _ = v13
    default:
        fmt.Println(v4, v6)
    }
    fmt.Println(v4, v4, v5)
    if false {
        v6 += v6
    } else {
        if (!("\n\t" == (("foo bar" + ("" + "")) + ""))) {
            var v14 int16 = 465
            switch 2147483647 {
            case -1:
                v6 = v6
            default:
                v6 = 509
            }
            _ = v14
        } else {
            var v15 float64 = (1.5 - -2.25)
            var v16 int8 = (v4 * v4)
            _ = v15
            _ = v16
        }
    }
    if (!false) {
        v1 = append(v1, 0)
        fmt.Println(v4, v4, v4)
    }
    if ((-508 <= -2147483648) && ((!(!true)) && (145 < -1))) {
        if true {
            var v17 int16 = 857
            fmt.Println(v6)
            for fi18 := 0; fi18 < 9; fi18++ {
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                var v19 uint64 = 0
                v4 = int8(v6)
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            v6 += v6
            var v21 uint = 18446744073709551615
            _ = fi20
            _ = v21
        }
    } else {
        v4 = v4
    }
    var v22 uint = 2
    v4 = (-(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
