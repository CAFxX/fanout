package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint8; F1 int64 }
func genFn0() string {
    var v4 uint = 1
    if (false || false) {
        v5 := GenS1{F0: true, F1: (v4 ^ (v4 | (v4 >> (((^(v4))) & 63)))), F2: 1}
        _ = v5
    }
    var v6 float32 = ((3.14159 / ((1e10 - 3.14159) * 1.5)) * 3.14159)
    var v7 uint = (v4 * (^(v4)))
    _ = v4
    _ = v6
    _ = v7
    return "a"
}

func main() {
    var v1 bool = true
    var v2 string = "foo bar"
    var v3 int8 = 57
    v8 := GenS1{F0: false, F1: 18446744073709551615, F2: -11}
    var v9 uint64 = 687
    fmt.Println(v3, v9, v9)
    v2 = ("hello" + ("\n\t" + v2))
    var v10 uint = 539
    switch -516 {
    case -304:
        var v11 uint64 = (v9 + v9)
        _ = v11
    case -2147483648:
        v1 = v1
    }
    var v12 int16 = -582
    switch (v2 + (v2 + (("foo bar" + "hello") + ("" + v2)))) {
    case "\n\t":
        if ((1.5 / -2.25) < 1e10) {
            var v13 uint16 = 2
            _ = v13
        } else {
            var v14 bool = false
            _ = v14
        }
    case "a":
        v15 := GenS1{F0: ((!(!(4294967295 <= 0))) && (((3.14159 - 3.14159) * 1e10) < 1.5)), F1: (^(v10)), F2: (v3 ^ v3)}
        _ = v15
    default:
        if (2147483647 <= 61) {
            for fi16 := 0; fi16 < 4; fi16++ {
                if (!true) {
                    switch "a" {
                    case "hello":
                        v17 := GenS1{F0: (1e10 >= 3.14159), F1: 2, F2: 127}
                        _ = v17
                    case "a":
                        v12 = -1
                    default:
                        if v1 {
                            fmt.Println(v3, v3, v9)
                            switch (fi16 / ((1) | 1)) {
                            case -591:
                                v18 := []int{-452, 326}
                                _ = v18
                            default:
                                for fi19 := 0; fi19 < 16; fi19++ {
                                    v20 := GenS2{F0: v12, F1: 0}
                                    v21 := GenS2{F0: v12, F1: 688}
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                            }
                            v22 := GenS2{F0: v12, F1: -9223372036854775808}
                            _ = v22
                        } else {
                            if v1 {
                                v2 = v2
                                fmt.Println(v12, v12)
                            } else {
                                if (v1 && false) {
                                    for fi23 := 0; fi23 < 19; fi23++ {
                                        var v24 uint8 = 1
                                        _ = fi23
                                        _ = v24
                                    }
                                    v25 := []int{1, fi16}
                                    var v26 uint = (^(v10))
                                    _ = v25
                                    _ = v26
                                }
                                v27 := []int{fi16, fi16}
                                _ = v27
                            }
                        }
                    }
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 string = v2
                        var v30 float64 = (((-2.25 / float64((-2.25 + ((1.5 - -2.25) / (((0.1 / (float32((1e10 * (1e10 / 1.0))) * float32(1.5))) + (0.1 * 3.14159)) * ((-2.25 - (0.1 - 1e10)) - 3.14159)))))) + 1e10) * 1.5)
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
                v12 = 0
                _ = fi16
            }
            fmt.Println(v1, v3, v9)
            var v31 string = "foo bar"
            _ = v31
        } else {
            var v32 uint32 = 65535
            v33 := []int{}
            v1 = ((v12 >> ((v10) & 63)) <= 1)
            _ = v32
            _ = v33
        }
    }
    switch "hello" {
    case "\n\t":
        for fi34 := 0; fi34 < 14; fi34++ {
            v12 = (v12 / ((v12) | 1))
            _ = fi34
        }
    case "a":
        for fi35 := 0; fi35 < 11; fi35++ {
            var v36 int8 = (^(v3))
            fmt.Println(v36, v12)
            var v37 uint32 = 725
            _ = fi35
            _ = v36
            _ = v37
        }
    case "foo bar":
        v3 += (-(v3))
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v1)
    case "a":
        if (v1 && (true || ((v1 && (808 < 924)) || (v1 && (v1 || false))))) {
            fmt.Println(v12, v3)
        }
    case "\n\t":
        fmt.Println(v10, v2, v1)
    }
    v38 := GenS3{F0: 183, F1: -1}
    v9 -= (v9 | v9)
    v3 -= (-(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v38
    _ = v8
    _ = v9
}
