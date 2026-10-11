package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint; F2 string }
func genFn0(p0 uint32, p1 int32) int8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6, p1, fi6)
        fmt.Println(fi6, p0, p1)
        _ = fi6
    }
    var v7 int = 9223372036854775807
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int64 = 42
        if true {
            switch "\n\t" {
            case "foo bar":
                if false {
                    v7 = v7
                    var v10 float64 = (1.5 / (3.14159 - 0.1))
                    v11 := []int{}
                    _ = v10
                    _ = v11
                }
            case "":
                v9 = v9
            }
            switch fi8 {
            case 0:
                var v12 int32 = -1
                _ = v12
            case 1:
                v9 -= (v9 / ((v9) | 1))
            case -9223372036854775808:
                v7 += (fi8 >> ((18446744073709551615) & 63))
            default:
                v13 := GenS1{F0: 2}
                _ = v13
            }
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 127
}

func main() {
    var v1 int16 = 14
    var v2 bool = (42 == 18446744073709551615)
    v3 := []uint32{0}
    var v4 string = "a"
    v5 := make(map[int]string)
    switch v4 {
    case "\n\t":
        var v14 int16 = v1
        _ = v14
    case "hello":
        switch -9223372036854775808 {
        case 4294967295:
            v15 := GenS1{F0: 984}
            _ = v15
        case 967:
            v16 := GenS1{F0: 2147483647}
            _ = v16
        case 32768:
            if (0 < 2147483647) {
                v17 := GenS1{F0: 1}
                _ = v17
            } else {
                v18 := GenS1{F0: 1}
                _ = v18
            }
        }
    default:
        switch int(v1) {
        case -478:
            v1 -= -634
        default:
            for fi19 := 0; fi19 < 20; fi19++ {
                v20 := GenS1{F0: 0}
                fmt.Println(v4, v4, v2)
                _ = fi19
                _ = v20
            }
        }
    }
    var v21 string = "a"
    var v22 bool = ((3.14159 / 1.0) != (1e10 / 1.0))
    v4 = "\n\t"
    v23 := GenS1{F0: 594}
    switch -211 {
    case -185:
        v1 -= v1
    case -84:
        if (!v2) {
            v24 := GenS1{F0: 2}
            _ = v24
        } else {
            v25 := GenS1{F0: 2147483647}
            _ = v25
        }
    default:
        v26 := GenS2{F0: 42}
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 uint64 = 803
        var v29 int8 = 42
        for fi30 := 0; fi30 < 1; fi30++ {
            var v31 int = (fi27 / ((fi27) | 1))
            v21 = v21
            _ = fi30
            _ = v31
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := v5[fi32]
        _ = fi32
        _ = v33
    }
    delete(v5, 586)
    if true {
        var v34 uint8 = 0
        v2 = v2
        var v35 int = 127
        _ = v34
        _ = v35
    } else {
        v36 := GenS3{F0: (true || v2), F1: 746, F2: v21}
        _ = v36
    }
    if (false && v2) {
        var v37 uint32 = 762
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
