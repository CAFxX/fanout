package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint32 }
func genFn0(p0 int64, p1 uint64) int16 {
    v5 := []int{}
    for fi6 := 0; fi6 < 5; fi6++ {
        switch (-(fi6)) {
        case -190:
            for fi7 := 0; fi7 < 8; fi7++ {
                if (-1 < (-(p0))) {
                    if len(v5) > 0 {
                        v5[0] = fi7
                    }
                    v8 := GenS1{F0: 0, F1: 951, F2: 2147483647}
                    v9 := GenS1{F0: 2, F1: -340, F2: 446}
                    _ = v8
                    _ = v9
                }
                _ = fi7
            }
        default:
            p1 = (^(p1))
        }
        fmt.Println(p0, p0)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi6
    }
    p1 -= p1
    v10 := GenS1{F0: 255, F1: 575, F2: 623}
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return 42
}

func main() {
    var v1 int16 = -32768
    v2 := []uint32{0}
    var v3 int16 = v1
    var v4 uint8 = 255
    switch "" {
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v2 = append(v2, 2147483647)
    }
    fmt.Println(v3, v4, v4)
    v4 -= 128
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 bool = true
        v13 := GenS1{F0: 2, F1: v3, F2: 1}
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        switch (fi14 * fi14) {
        case 42:
            v15 := GenS1{F0: uint16(v4), F1: genFn0(-1, 42), F2: 2147483648}
            _ = v15
        case 9223372036854775807:
            switch 0 {
            case 2147483647:
                var v16 uint8 = 1
                _ = v16
            default:
                v17 := GenS1{F0: 42, F1: genFn0(0, 2), F2: 1}
                _ = v17
            }
        default:
            var v18 int = (fi14 / (((fi14 ^ (-(fi14)))) | 1))
            _ = v18
        }
        var v19 float32 = 0.1
        _ = fi14
        _ = v19
    }
    switch 32767 {
    case -2147483648:
        v4 = v4
    case 128:
        var v20 int32 = -52
        _ = v20
    case -1:
        var v21 int16 = v3
        _ = v21
    }
    switch "a" {
    case "a":
        switch 9223372036854775807 {
        case 128:
            if (true || ("\n\t" > "a")) {
                v1 = (v1 << ((19) & 63))
                v1 -= (v3 & v1)
            } else {
                if ((47 <= 0) || true) {
                    if false {
                        v3 -= (v1 << ((2) & 63))
                        var v22 uint32 = 65535
                        _ = v22
                    }
                } else {
                    var v23 int8 = -122
                    _ = v23
                }
                v4 += (v4 * (v4 * (v4 & 92)))
            }
        case -2147483648:
            _ = len(v2)
        case 0:
            v3 -= v1
        default:
            if true {
                var v24 uint64 = 0
                _ = len(v2)
                var v25 uint = 65535
                _ = v24
                _ = v25
            } else {
                var v26 uint16 = 32767
                v3 = v3
                _ = v26
            }
        }
    case "":
        fmt.Println(v3, v4, v3)
    case "hello":
        v27 := GenS1{F0: 32767, F1: v3, F2: 2}
        _ = v27
    }
    var v28 int = -9223372036854775808
    switch ("foo bar" + "") {
    case "\n\t":
        v4 += 127
    case "a":
        for fi29 := 0; fi29 < 12; fi29++ {
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 int8 = 59
                switch "hello" {
                case "a":
                    var v32 int32 = 0
                    _ = v32
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = 2147483647
                    }
                case "":
                    fmt.Println(v28, v28)
                }
                _ = fi30
                _ = v31
            }
            var v33 uint = 4294967295
            var v34 int32 = -2147483648
            _ = fi29
            _ = v33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        v36 := GenS1{F0: 42, F1: genFn0(0, 1), F2: 2147483647}
        _ = fi35
        _ = v36
    }
    var v37 string = "foo bar"
    fmt.Println(v1, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v4
}
