package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int16 }
func genFn0() uint8 {
    var v4 uint32 = 2
    v5 := GenS1{F0: 0, F1: false, F2: 858}
    fmt.Println(v4)
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 0, F1: ((fi6 & (^(fi6))) > (^(fi6))), F2: 1000}
        if false {
            var v8 float32 = 1e10
            var v9 int8 = -59
            var v10 uint32 = 747
            _ = v10
            _ = v8
            _ = v9
        }
        var v11 int = fi6
        _ = fi6
        _ = v11
        _ = v7
    }
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint32, p1 int32) uint64 {
    p0 = (p0 * p0)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 939
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    var v3 int = 3
    if (0 >= 990) {
        for fi12 := 0; fi12 < 4; fi12++ {
            v2 = (419 != 0)
            switch v3 {
            case 1:
                v2 = v2
            }
            _ = fi12
        }
        v3 = 564
    }
    v1[32767] = 2147483648
    var v13 int64 = 9223372036854775807
    switch -1 {
    case 554:
        var v14 uint64 = 12
        _ = v14
    case -1:
        v15 := GenS1{F0: 42, F1: v2, F2: -105}
        _ = v15
    case 2147483647:
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int32 = -932
            switch ((("" + ("hello" + ("\n\t" + "\n\t"))) + "\n\t") + "hello") {
            case "":
                var v18 uint64 = 98
                _ = v18
            case "\n\t":
                var v19 uint8 = 128
                _ = v19
            default:
                v17 += -80
            }
            _ = fi16
            _ = v17
        }
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 3; fi21++ {
            if v2 {
                switch (-(fi20)) {
                case -1:
                    v22 := []int{(fi21 | 42)}
                    _ = v22
                case 0:
                    v2 = (77 > 2)
                case 9223372036854775807:
                    v2 = ("\n\t" <= "hello")
                default:
                    var v23 int16 = 0
                    _ = v23
                }
                v24 := GenS1{F0: 2, F1: false, F2: -1}
                v3 += v3
                _ = v24
            }
            v25 := GenS1{F0: 42, F1: v2, F2: 831}
            _ = fi21
            _ = v25
        }
        _ = fi20
    }
    v2 = (0.0 != 1.5)
    for fi26 := 0; fi26 < 20; fi26++ {
        if false {
            var v27 bool = ((4294967295 <= 1) || v2)
            v13 += 373
            v3 += fi26
            _ = v27
        }
        var v28 uint16 = 79
        var v29 uint = 255
        _ = fi26
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: 87, F1: (-1 >= 0), F2: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
}
