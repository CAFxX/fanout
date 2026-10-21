package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint32 {
    v6 := GenS1{F0: 0}
    switch "" {
    case "":
        var v7 uint = 2
        _ = v7
    case "\n\t":
        var v8 uint8 = 0
        _ = v8
    default:
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 float64 = -2.25
            _ = fi9
            _ = v10
        }
    }
    var v11 uint = 0
    v11 -= 255
    _ = v11
    _ = v6
    return 550
}

func main() {
    var v1 int32 = 153
    v2 := make(map[int]bool)
    var v3 uint = 937
    v4 := make(map[int]int)
    var v5 uint64 = 4294967295
    if false {
        v1 = v1
    } else {
        switch "a" {
        case "a":
            var v12 uint16 = 1
            _ = v12
        case "hello":
            fmt.Println(v5)
        case "":
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 int = -1
                _ = fi13
                _ = v14
            }
        default:
            if (!(false && (-339 > -32768))) {
                var v15 uint32 = 0
                v16 := v2[42]
                v15 -= v15
                _ = v15
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            v1 = v1
            v18 := []int{(fi17 & (fi17 / ((4294967295) | 1))), (fi17 / ((fi17) | 1)), (fi17 | (fi17 & fi17))}
            if false {
                var v19 int64 = 0
                _ = v19
            } else {
                v20 := GenS1{F0: 41}
                _ = v20
            }
            _ = fi17
            _ = v18
        }
    }
    var v21 int32 = v1
    v22 := GenS1{F0: v5}
    v23 := []int{-15, 548, -2147483648}
    if (0 != 100) {
        if false {
            var v24 string = "\n\t"
            _ = v24
        } else {
            fmt.Println(v3, v5)
            fmt.Println(v3, v3, v21)
        }
        v25 := GenS1{F0: (v5 ^ (v5 - (v5 % ((314) | 1))))}
        switch ("hello" + ("" + "\n\t")) {
        case "foo bar":
            if false {
                var v26 float64 = 3.14159
                _ = v26
            }
        case "\n\t":
            fmt.Println(v5, v1, v5)
        default:
            switch "a" {
            case "\n\t":
                if false {
                    v21 -= v21
                }
            }
        }
        _ = v25
    } else {
        var v27 uint8 = 192
        var v28 uint64 = v5
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        if true {
            switch fi29 {
            case 163:
                v5 += 1
            case 295:
                var v30 bool = true
                _ = v30
            default:
                v23 = append(v23, -99)
            }
            v23 = append(v23, 389)
            fmt.Println(fi29, v21, v5)
        }
        _ = fi29
    }
    switch 2147483647 {
    case 351:
        v31 := GenS1{F0: v5}
        _ = v31
    default:
        var v32 int64 = -9223372036854775808
        _ = v32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        v5 -= (v5 * 242)
        var v34 uint = (v3 / ((1) | 1))
        _ = fi33
        _ = v34
    }
    if (2 < 72) {
        var v35 uint16 = 42
        v5 = 1
        _ = v35
    } else {
        var v36 int8 = -128
        if false {
            v5 += (^(v5))
        }
        _ = len(v2)
        _ = v36
    }
    _ = len(v4)
    var v37 uint64 = v5
    var v38 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
