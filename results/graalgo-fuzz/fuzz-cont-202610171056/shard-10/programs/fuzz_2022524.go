package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    var v1 bool = (43 == 0)
    var v2 bool = true
    var v3 bool = (4294967295 < 1)
    var v4 uint = 255
    if false {
        v5 := GenS1{F0: 4294967295, F1: -1}
        var v6 uint8 = 2
        var v7 uint = v4
        _ = v5
        _ = v6
        _ = v7
    } else {
        v2 = (1e10 <= (3.14159 - 0.1))
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint16 = 42
        _ = fi8
        _ = v9
    }
    var v10 bool = v1
    v3 = v1
    switch 168 {
    case -1:
        switch 675 {
        case -2147483648:
            if v10 {
                if (false || (0 < 0)) {
                    var v11 int8 = -1
                    _ = v11
                }
                v12 := GenS1{F0: 571, F1: 55}
                _ = v12
            }
        }
    case 9223372036854775807:
        var v13 uint64 = 1
        _ = v13
    case 32767:
        for fi14 := 0; fi14 < 15; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 uint64 = 42
                v3 = false
                _ = fi15
                _ = v16
            }
            v2 = (int8(fi14) <= -1)
            v2 = true
            _ = fi14
        }
    }
    if v1 {
        var v17 float32 = 3.14159
        _ = v17
    }
    v18 := []int{-9223372036854775808}
    v19 := GenS1{F0: 127, F1: -1}
    if v10 {
        v4 = (v4 - v4)
        var v20 uint = (v4 + v4)
        for fi21 := 0; fi21 < 17; fi21++ {
            switch "\n\t" {
            case "foo bar":
                v10 = (v2 || v3)
            case "\n\t":
                var v22 int = fi21
                _ = v22
            case "hello":
                v20 += v4
            default:
                switch "a" {
                case "hello":
                    var v23 float64 = 0.1
                    _ = v23
                case "foo bar":
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                case "\n\t":
                    fmt.Println(v3)
                default:
                    switch "hello" {
                    case "foo bar":
                        var v24 int16 = 42
                        _ = v24
                    }
                }
            }
            for fi25 := 0; fi25 < 17; fi25++ {
                v20 = v20
                var v26 int64 = -268
                var v27 int8 = -46
                _ = fi25
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 3; fi28++ {
                var v29 uint8 = 42
                fmt.Println(v1, v29, v29)
                v3 = v10
                var v30 int = 2147483647
                v31 := GenS1{F0: (fi28 | 50), F1: 100}
                v1 = (!(v3 || (!v2)))
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi21
        }
        _ = v20
    } else {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        fmt.Println(v4, v2, v1)
    }
    if v2 {
        var v32 uint64 = 1
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 uint8 = 128
            _ = fi33
            _ = v34
        }
        _ = v32
    } else {
        var v35 string = "a"
        _ = v35
    }
    v36 := GenS1{F0: 42, F1: 127}
    if (128 != 2) {
        for fi37 := 0; fi37 < 7; fi37++ {
            if (!false) {
                var v38 int = (^(fi37))
                v18 = append(v18, fi37)
                _ = v38
            } else {
                if (v3 && (42 > 2)) {
                    for fi39 := 0; fi39 < 7; fi39++ {
                        var v40 uint16 = 255
                        v3 = false
                        _ = fi39
                        _ = v40
                    }
                    var v41 int32 = -1
                    _ = v41
                }
                v3 = (1e10 != 0.0)
            }
            _ = fi37
        }
        v4 += (v4 >> ((18446744073709551615) & 63))
        var v42 int = 0
        _ = v42
    }
    var v43 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v43
}
