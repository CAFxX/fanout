package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int32; F2 bool }
type GenS4 struct { F0 uint8; F1 string; F2 int16 }
func genFn0() uint32 {
    switch -724 {
    case 275:
        switch -2147483648 {
        case -339:
            if ((true && (!(true || true))) && ((-2.25 + -2.25) <= ((0.1 - (0.1 / ((0.0 - -2.25) / ((0.0 * (float64(float32((((3.14159 / 3.14159) / 1.0) * 0.0))) / 3.14159)) - ((0.0 + (0.1 - (float64(1.5) + (float64(0.0) / (0.1 * 1.5))))) - 1.5))))) + (0.0 + 0.1)))) {
                v6 := GenS1{F0: 127}
                v7 := []int{42, 0}
                _ = v6
                _ = v7
            } else {
                var v8 uint16 = 32767
                v8 += 502
                for fi9 := 0; fi9 < 6; fi9++ {
                    var v10 float32 = 0.1
                    _ = fi9
                    _ = v10
                }
                _ = v8
            }
        case 1:
            v11 := GenS1{F0: 100}
            _ = v11
        }
    default:
        var v12 uint = 340
        _ = v12
    }
    v13 := GenS2{F0: -76}
    _ = v13
    return 0
}
func genFn1(p0 int64) int64 {
    var v14 int = -9223372036854775808
    v15 := GenS3{F0: p0, F1: -1, F2: true}
    var v16 uint8 = 155
    var v17 int32 = -542
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return p0
}

func main() {
    var v1 int32 = -552
    var v2 int64 = 42
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    fmt.Println(v2, v1, v2)
    switch 61 {
    case -315:
        v1 = 42
    case 550:
        fmt.Println(v1)
    default:
        v18 := GenS2{F0: -33}
        _ = v18
    }
    var v19 string = "\n\t"
    v2 = v2
    fmt.Println(v1)
    fmt.Println(v1, v2)
    v2 = (v2 + v2)
    fmt.Println(v1, v1, v19)
    v20 := GenS3{F0: genFn1(genFn1((-(v2)))), F1: (^(v1)), F2: v3}
    switch 128 {
    case 918:
        v19 = v19
    case 102:
        var v21 bool = (!(v3 && v3))
        _ = v21
    case 32767:
        v22 := GenS4{F0: 127, F1: "a", F2: 42}
        _ = v22
    default:
        switch "" {
        case "foo bar":
            v23 := GenS4{F0: 2, F1: v19, F2: 1}
            _ = v23
        case "\n\t":
            var v24 bool = (874 >= 0)
            _ = v24
        case "a":
            v25 := []int{int(v1), 9223372036854775807, -506, -653}
            _ = v25
        default:
            switch 0 {
            case 128:
                v26 := GenS4{F0: 1, F1: v19, F2: -32768}
                _ = v26
            case 32767:
                fmt.Println(v19)
            case 2147483647:
                for fi27 := 0; fi27 < 20; fi27++ {
                    _ = len(v4)
                    for fi28 := 0; fi28 < 19; fi28++ {
                        if (42 < 42) {
                            fmt.Println(fi28, v19)
                            switch v19 {
                            case "a":
                                v29 := []int{(fi28 * (fi27 / ((fi27) | 1))), fi27}
                                _ = v29
                            }
                        }
                        _ = fi28
                    }
                    _ = fi27
                }
            }
        }
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v1)
    case "":
        fmt.Println(v2, v1)
    case "\n\t":
        v1 += (v1 + v1)
    default:
        delete(v5, -1)
    }
    v5[128] = 0
    for fi30 := 0; fi30 < 13; fi30++ {
        v19 = v19
        switch fi30 {
        case 42:
            for fi31 := 0; fi31 < 11; fi31++ {
                v32 := []int{}
                if v3 {
                    var v33 uint32 = 2
                    v34 := GenS4{F0: 127, F1: ("" + "a"), F2: 151}
                    _ = v33
                    _ = v34
                }
                v35 := GenS4{F0: 2, F1: v19, F2: 42}
                _ = fi31
                _ = v32
                _ = v35
            }
        case 131:
            for fi36 := 0; fi36 < 6; fi36++ {
                v2 -= genFn1(42)
                _ = fi36
            }
        }
        var v37 int8 = 54
        _ = fi30
        _ = v37
    }
    var v38 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
