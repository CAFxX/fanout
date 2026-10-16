package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int }
func genFn0(p0 uint, p1 int) bool {
    v5 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: false}
    fmt.Println(p1)
    var v6 int16 = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return false
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    v4 := []int{-305, 127, 2147483647, -1}
    _ = len(v1)
    _ = len(v4)
    for fi7 := 0; fi7 < 1; fi7++ {
        switch 391 {
        case 42:
            if (((3.14159 * 1e10) - -2.25) != (1e10 - float64((((1.5 / 1.0) / 1.0) - (0.1 - 3.14159))))) {
                v1["hello"] = (fi7 << ((1) & 63))
                v4 = append(v4, fi7)
            }
        case -301:
            var v8 bool = (883 <= 0)
            _ = v8
        case 282:
            var v9 uint8 = 1
            _ = v9
        default:
            var v10 uint64 = 338
            _ = v10
        }
        var v11 int8 = 1
        if (42 != 42) {
            fmt.Println(fi7, v11)
            delete(v3, "a")
            switch "" {
            case "hello":
                if (v11 < (^(v11))) {
                    var v12 int8 = v11
                    _ = v12
                } else {
                    if (710 > 627) {
                        fmt.Println(v11)
                        var v13 uint8 = 1
                        fmt.Println(v11)
                        _ = v13
                    }
                }
            default:
                for fi14 := 0; fi14 < 12; fi14++ {
                    v11 = (^(v11))
                    v11 += v11
                    v15 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: (65535 == 65535)}
                    _ = fi14
                    _ = v15
                }
            }
        } else {
            _ = len(v4)
        }
        _ = fi7
        _ = v11
    }
    var v16 uint16 = 255
    var v17 uint16 = (v16 & (v16 << ((0) & 63)))
    v18 := GenS2{F0: 325}
    v16 -= v16
    if false {
        var v19 float64 = -2.25
        var v20 int64 = 342
        v20 -= (^(v20))
        _ = v19
        _ = v20
    } else {
        switch "\n\t" {
        case "\n\t":
            switch 42 {
            case 1:
                v21 := GenS1{F0: 255, F1: 272, F2: false}
                _ = v21
            case 32768:
                if false {
                    v17 = v17
                    v22 := GenS1{F0: 57, F1: 18446744073709551615, F2: false}
                    _ = len(v3)
                    _ = v22
                }
            case 0:
                var v23 int32 = 0
                _ = v23
            default:
                var v24 uint64 = 42
                _ = v24
            }
        }
        var v25 int = 127
        _ = v25
    }
    var v26 string = ("hello" + "foo bar")
    v16 = v17
    if false {
        v27 := GenS3{F0: 843, F1: -9223372036854775808}
        var v28 uint32 = 666
        switch ((v26 + "hello") + (v26 + "")) {
        case "hello":
            var v29 int16 = 0
            _ = v29
        case "a":
            var v30 uint32 = (v28 << ((255) & 63))
            _ = v30
        default:
            v16 += 872
        }
        _ = v27
        _ = v28
    } else {
        if false {
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v31 uint8 = 11
                _ = v31
            }
            delete(v2, "foo bar")
        }
        switch 127 {
        case -354:
            fmt.Println(v26)
        case -9223372036854775808:
            var v32 int64 = -9223372036854775808
            _ = v32
        }
        v16 -= 42
    }
    v33 := GenS2{F0: 196}
    for fi34 := 0; fi34 < 9; fi34++ {
        if true {
            fmt.Println(v17)
            if false {
                for fi35 := 0; fi35 < 7; fi35++ {
                    v36 := GenS1{F0: 0, F1: 2, F2: false}
                    _ = fi35
                    _ = v36
                }
                fmt.Println(v17)
            }
        } else {
            fmt.Println(v26)
            if true {
                if (false && false) {
                    var v37 string = v26
                    _ = v37
                } else {
                    v38 := GenS1{F0: 2, F1: 4294967295, F2: (726 < 0)}
                    var v39 int16 = -1
                    _ = v38
                    _ = v39
                }
                var v40 uint8 = 127
                _ = v40
            } else {
                v2[v26] = (v26 + v26)
                var v41 string = v26
                _ = v41
            }
        }
        _ = fi34
    }
    fmt.Println(v17, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v4
}
