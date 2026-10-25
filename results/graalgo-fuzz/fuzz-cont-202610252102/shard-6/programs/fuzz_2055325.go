package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }
func genFn0() int64 {
    var v4 int64 = -1
    v4 -= (v4 & v4)
    if (619 == 4294967295) {
        v4 = -310
        v5 := []int{}
        if true {
            var v6 int32 = 0
            var v7 float64 = 0.1
            switch 128 {
            case 127:
                v7 = v7
            default:
                for fi8 := 0; fi8 < 18; fi8++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    for fi9 := 0; fi9 < 4; fi9++ {
                        v10 := GenS1{F0: 0, F1: (fi8 | int(v6))}
                        v6 = 0
                        _ = fi9
                        _ = v10
                    }
                    var v11 uint8 = 42
                    _ = fi8
                    _ = v11
                }
            }
            _ = v6
            _ = v7
        }
        _ = v5
    } else {
        v4 = v4
    }
    _ = v4
    return (v4 + (-(v4)))
}
func genFn1() uint8 {
    if (82 <= 171) {
        var v12 uint64 = 0
        v13 := GenS1{F0: 1, F1: -217}
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 2147483648, F1: 42}
        v15 := []int{2147483647, 86}
        switch 4294967295 {
        case 0:
            v16 := GenS1{F0: 554, F1: 0}
            _ = v16
        case 361:
            var v17 uint64 = 961
            _ = v17
        default:
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        }
        _ = v14
        _ = v15
    }
    var v18 uint = 42
    v18 += 797
    v19 := []int{9223372036854775807, 128}
    _ = v18
    _ = v19
    return 2
}

func main() {
    v1 := []string{("hello" + ("hello" + "")), "a"}
    v2 := make(map[int]uint32)
    var v3 uint = 0
    delete(v2, 1)
    _ = len(v1)
    switch "hello" {
    case "\n\t":
        var v20 uint32 = 2
        _ = v20
    default:
        var v21 float64 = float64(3.14159)
        _ = v21
    }
    var v22 uint = v3
    for fi23 := 0; fi23 < 10; fi23++ {
        if ((!true) && ((9223372036854775807 > 0) || true)) {
            fmt.Println(fi23)
        }
        _ = fi23
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 bool = (!false)
        _ = fi24
        _ = v25
    }
    fmt.Println(v22, v3)
    var v26 float64 = 1.5
    var v27 string = ((("foo bar" + "hello") + (("" + ("hello" + "hello")) + "hello")) + "")
    for fi28 := 0; fi28 < 7; fi28++ {
        if (v22 >= (v3 & 65535)) {
            v29 := GenS1{F0: 2, F1: (-(fi28))}
            var v30 int8 = -1
            if false {
                v31 := GenS1{F0: 594, F1: (fi28 / (((-(fi28))) | 1))}
                v2[(-(fi28))] = 42
                _ = v31
            } else {
                fmt.Println(v22, v26)
                if (65535 > 2147483647) {
                    var v32 uint8 = 0
                    _ = v32
                }
            }
            _ = v29
            _ = v30
        }
        var v33 float64 = v26
        _ = fi28
        _ = v33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        v26 = (1.5 / ((v26 * v26) / float64((0.1 + (0.1 / 0.1)))))
        v35 := GenS2{F0: -2147483648, F1: (v22 + 456), F2: 32767}
        v22 = v22
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        for fi37 := 0; fi37 < 10; fi37++ {
            switch v27 {
            case "foo bar":
                v27 = (v27 + v27)
            case "\n\t":
                v26 = (0.0 - (v26 * (v26 * 1.5)))
            default:
                var v38 int8 = 1
                _ = v38
            }
            _ = fi37
        }
        v22 -= (^(v3))
        var v39 string = v27
        _ = fi36
        _ = v39
    }
    var v40 uint32 = 2147483648
    var v41 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v40
    _ = v41
}
