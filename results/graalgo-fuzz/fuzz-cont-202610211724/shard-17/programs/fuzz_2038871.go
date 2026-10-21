package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int; F1 int8; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint }

func main() {
    var v1 int64 = 477
    var v2 uint32 = 341
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    var v5 int = 564
    v5 = v5
    if false {
        var v6 string = "foo bar"
        _ = v6
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            switch (-(fi7)) {
            case 477:
                var v8 int16 = -594
                _ = v8
            }
            var v9 bool = true
            v10 := []int{32768, (v5 * (^(v5))), (v5 - (fi7 | -519)), (^(v5))}
            _ = fi7
            _ = v10
            _ = v9
        }
    }
    fmt.Println(v1)
    v11 := GenS1{F0: ("hello" + "\n\t"), F1: -128}
    v2 = 42
    var v12 uint32 = (v2 & v2)
    v13 := v3[v12]
    v13 = v2
    v1 = 42
    fmt.Println(v5, v12)
    if (255 >= 138) {
        if false {
            var v14 uint = uint(v2)
            _ = v14
        } else {
            for fi15 := 0; fi15 < 7; fi15++ {
                fmt.Println(v2)
                _ = fi15
            }
        }
    }
    var v16 int32 = 42
    fmt.Println(v1, v2)
    if true {
        v17 := []int{-2147483648, 128, (v5 | 42)}
        if (v12 == (^(v13))) {
            v5 = v5
            for fi18 := 0; fi18 < 6; fi18++ {
                v2 += 412
                var v19 int = -2147483648
                var v20 bool = false
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        _ = v17
    } else {
        if (!false) {
            var v21 float32 = 1e10
            if (!(false || (!true))) {
                v22 := GenS2{F0: -1, F1: 127, F2: (^(v1))}
                _ = v22
            }
            _ = v21
        } else {
            var v23 int = v5
            v24 := GenS2{F0: v23, F1: 100, F2: -9223372036854775808}
            _ = v23
            _ = v24
        }
        fmt.Println(v2, v5, v13)
        switch v5 {
        case 32768:
            v25 := []int{v5}
            _ = v25
        case -1:
            for fi26 := 0; fi26 < 17; fi26++ {
                v27 := GenS3{F0: 11}
                for fi28 := 0; fi28 < 17; fi28++ {
                    var v29 int16 = -32768
                    var v30 int64 = -9223372036854775808
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                switch (fi26 & 42) {
                case 128:
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        fmt.Println(v1, v13)
                    case "":
                        var v31 int16 = 0
                        _ = v31
                    case "hello":
                        v4["hello"] = (!false)
                    default:
                        for fi32 := 0; fi32 < 9; fi32++ {
                            v33 := GenS2{F0: -9223372036854775808, F1: 127, F2: v1}
                            fmt.Println(v1)
                            _ = fi32
                            _ = v33
                        }
                    }
                case -1:
                    v34 := GenS4{F0: 175, F1: 255, F2: 1}
                    _ = v34
                }
                _ = fi26
                _ = v27
            }
        case -2147483648:
            delete(v3, v12)
        default:
            var v35 float32 = 1.5
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
