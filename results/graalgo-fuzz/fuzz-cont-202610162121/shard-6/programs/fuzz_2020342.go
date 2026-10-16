package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int32 }
type GenS2 struct { F0 int32 }
func genFn0() int64 {
    for fi6 := 0; fi6 < 1; fi6++ {
        if (694 == 2) {
            switch 478 {
            case 9223372036854775807:
                var v7 uint = 4294967295
                _ = v7
            }
            v8 := []int{}
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 bool = true
                v10 = v10
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        _ = fi6
    }
    v11 := []int{-9223372036854775808}
    _ = v11
    return -295
}
func genFn1() uint16 {
    var v12 bool = (315 >= 62)
    v12 = v12
    v13 := GenS1{F0: 397, F1: 521, F2: 42}
    var v14 uint = 0
    v12 = (!v12)
    _ = v12
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 uint32 = 162
    v2 := make(map[int]uint32)
    var v3 uint8 = 1
    var v4 int = 4294967295
    var v5 int64 = -581
    switch v4 {
    case 9223372036854775807:
        for fi15 := 0; fi15 < 12; fi15++ {
            v2[(fi15 >> ((0) & 63))] = (v1 << ((224) & 63))
            _ = fi15
        }
    case -9223372036854775808:
        switch "foo bar" {
        case "\n\t":
            var v16 uint = 2
            _ = v16
        }
    case 1:
        if (!false) {
            v17 := GenS2{F0: -477}
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 int16 = int16(v5)
                _ = fi18
                _ = v19
            }
            switch (v4 * (v4 & 1)) {
            case -120:
                v5 = (-(v5))
            case -2147483648:
                var v20 uint64 = 18446744073709551615
                _ = v20
            case 128:
                var v21 string = (("" + ("" + "a")) + "foo bar")
                _ = v21
            }
            _ = v17
        } else {
            for fi22 := 0; fi22 < 13; fi22++ {
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 uint8 = (v3 / (((^(v3))) | 1))
                    if true {
                        v24 -= v24
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 uint64 = 0
                _ = fi25
                _ = v26
            }
        }
    default:
        delete(v2, (v4 + (^(v4))))
    }
    v3 = 127
    var v27 int8 = -26
    for fi28 := 0; fi28 < 7; fi28++ {
        if false {
            var v29 float64 = 0.0
            if true {
                v30 := []int{0, (fi28 * fi28), 9223372036854775807, 42}
                fmt.Println(v1)
                if (true && (994 <= v1)) {
                    _ = len(v2)
                }
                _ = v30
            }
            _ = v29
        }
        _ = fi28
    }
    var v31 float32 = 1.5
    v3 += (v3 << ((42) & 63))
    _ = len(v2)
    fmt.Println(v27, v4)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
