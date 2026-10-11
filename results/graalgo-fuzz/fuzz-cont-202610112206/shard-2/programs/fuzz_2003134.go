package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int; F1 int8 }

func main() {
    var v1 int32 = 0
    v2 := []uint32{1, 4294967295}
    var v3 uint8 = 0
    v4 := []int{415}
    var v5 int = -9223372036854775808
    var v6 float32 = float32((float64(0.0) + (1e10 * (0.0 * 0.0))))
    if true {
        switch 664 {
        case 978:
            v4 = append(v4, (v5 | v5))
        case -553:
            v7 := GenS1{F0: 65535, F1: 0}
            _ = v7
        case 4294967295:
            v4 = append(v4, 309)
        }
        v5 -= v5
    } else {
        switch (^(v5)) {
        case -724:
            var v8 bool = false
            _ = v8
        }
        if (int16(v5) >= -1) {
            if false {
                var v9 int8 = 1
                _ = v9
            } else {
                for fi10 := 0; fi10 < 1; fi10++ {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v12 := GenS1{F0: 1, F1: 0}
                        _ = fi11
                        _ = v12
                    }
                    v3 += 42
                    var v13 int64 = 42
                    _ = fi10
                    _ = v13
                }
            }
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int8 = 1
            _ = fi14
            _ = v15
        }
    }
    var v16 int32 = v1
    v17 := GenS1{F0: 0, F1: 4294967295}
    var v18 uint8 = (^(v3))
    if ((false || (false && false)) && false) {
        switch ("" + "") {
        case "":
            var v19 uint8 = v3
            _ = v19
        case "a":
            switch (v5 / ((v5) | 1)) {
            case 334:
                v20 := GenS1{F0: 65535, F1: 743}
                _ = v20
            case 42:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        case "foo bar":
            fmt.Println(v1, v3)
        default:
            v5 -= 127
        }
        if false {
            for fi21 := 0; fi21 < 2; fi21++ {
                v16 = (v1 * -2147483648)
                fmt.Println(v18, v1, v6)
                v22 := GenS1{F0: 1, F1: 681}
                _ = fi21
                _ = v22
            }
        }
    } else {
        if true {
            v23 := GenS2{F0: v5, F1: -14}
            _ = v23
        }
        var v24 int32 = (v16 | v1)
        v25 := GenS2{F0: 4294967295, F1: 127}
        _ = v24
        _ = v25
    }
    var v26 uint = 383
    var v27 int8 = 1
    v4 = append(v4, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
