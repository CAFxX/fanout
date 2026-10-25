package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]string)
    var v3 uint16 = 2
    var v4 float32 = 1.5
    v5 := make(map[int]bool)
    if (42 != 901) {
        switch 442 {
        case 9223372036854775807:
            fmt.Println(v1, v4, v4)
        case -2147483648:
            v4 -= (v4 - v4)
        case 819:
            v6 := GenS1{F0: 127, F1: 127, F2: 2147483648}
            _ = v6
        default:
            var v7 int8 = -128
            _ = v7
        }
    } else {
        if (1e10 > -2.25) {
            v8 := []int{}
            switch (v1 + "a") {
            case "foo bar":
                fmt.Println(v1, v4)
            }
            _ = v8
        } else {
            delete(v2, 280)
            if (!false) {
                var v9 float32 = (0.0 + 1.5)
                var v10 uint16 = v3
                v11 := GenS1{F0: 626, F1: 42, F2: 60}
                _ = v10
                _ = v11
                _ = v9
            }
        }
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        if (4294967295 > 0) {
            v4 = v4
            v13 := []int{25, (fi12 << ((2) & 63))}
            if true {
                v14 := v5[(fi12 << ((170) & 63))]
                if (32767 != int16(fi12)) {
                    var v15 uint32 = 65535
                    delete(v5, fi12)
                    _ = v15
                }
                v16 := GenS1{F0: (^(fi12)), F1: -128, F2: 1}
                _ = v14
                _ = v16
            }
            _ = v13
        }
        var v17 int = (fi12 + fi12)
        if (!((true || false) && ((2 == 4294967295) || false))) {
            var v18 uint = 42
            for fi19 := 0; fi19 < 11; fi19++ {
                v20 := GenS1{F0: v17, F1: 42, F2: 349}
                var v21 int32 = -2147483648
                _ = fi19
                _ = v20
                _ = v21
            }
            var v22 uint32 = 372
            _ = v18
            _ = v22
        }
        _ = fi12
        _ = v17
    }
    v3 -= 587
    v23 := GenS1{F0: 32768, F1: -128, F2: 2147483648}
    if (2 == 2147483648) {
        switch "hello" {
        case "foo bar":
            for fi24 := 0; fi24 < 9; fi24++ {
                v1 = v1
                _ = fi24
            }
        case "":
            if false {
                var v25 float32 = v4
                _ = v25
            }
        case "\n\t":
            v3 -= v3
        default:
            v5[-9223372036854775808] = false
        }
        v1 = "hello"
    } else {
        v3 -= v3
    }
    v26 := GenS2{F0: 4294967295}
    for fi27 := 0; fi27 < 13; fi27++ {
        if true {
            for fi28 := 0; fi28 < 1; fi28++ {
                v29 := v5[fi27]
                v30 := GenS1{F0: (fi28 >> ((255) & 63)), F1: -128, F2: 2147483648}
                switch "" {
                case "hello":
                    var v31 uint64 = 2
                    _ = v31
                case "\n\t":
                    v1 = "a"
                }
                _ = fi28
                _ = v29
                _ = v30
            }
            v4 = (v4 + 0.1)
        }
        _ = fi27
    }
    var v32 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
