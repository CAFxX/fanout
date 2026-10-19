package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
func genFn0() bool {
    var v5 uint32 = 493
    switch "foo bar" {
    case "":
        v5 -= (^(v5))
    default:
        fmt.Println(v5)
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{fi6, (fi6 * 0), fi7}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int = (fi6 % (((^(fi9))) | 1))
            var v11 float64 = float64(float32(float64(0.1)))
            if false {
                v12 := []int{(fi6 >> ((934) & 63)), 1, (fi9 >> ((4294967295) & 63)), fi6}
                _ = v12
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        if true {
            var v13 bool = (!true)
            var v14 uint16 = 896
            for fi15 := 0; fi15 < 17; fi15++ {
                v16 := []int{(fi6 * fi6), fi15, 32767}
                _ = fi15
                _ = v16
            }
            _ = v13
            _ = v14
        }
        _ = fi6
    }
    _ = v5
    return false
}

func main() {
    var v1 int8 = 127
    var v2 float32 = 3.14159
    var v3 int8 = (v1 << ((1) & 63))
    var v4 float32 = v2
    v1 += -128
    var v17 uint32 = 2147483647
    v17 = (^(v17))
    var v18 float32 = v4
    var v19 float64 = -2.25
    var v20 int = 32768
    var v21 int16 = -413
    v22 := GenS1{F0: 42}
    switch 765 {
    case 2147483647:
        fmt.Println(v17)
    default:
        var v23 uint16 = 338
        _ = v23
    }
    var v24 float64 = 0.0
    fmt.Println(v17)
    var v25 int64 = 42
    v26 := GenS1{F0: 0}
    v20 += v20
    switch "foo bar" {
    case "\n\t":
        var v27 bool = (((!(!(0 != (v3 - v1)))) || false) && true)
        _ = v27
    case "foo bar":
        for fi28 := 0; fi28 < 10; fi28++ {
            v29 := GenS2{F0: (true || ((-(v25)) >= (-(v25))))}
            v30 := []int{(v20 & fi28), v20}
            var v31 int8 = (-(v1))
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
    case "":
        v32 := GenS1{F0: 700}
        _ = v32
    }
    v33 := []int{(v20 >> ((2) & 63))}
    for fi34 := 0; fi34 < 2; fi34++ {
        for fi35 := 0; fi35 < 7; fi35++ {
            v36 := GenS1{F0: 145}
            if false {
                v24 += ((float64((0.0 - v4)) - ((v19 + v19) * v19)) * (((0.1 * 1e10) - ((v24 / (v24 + v24)) - -2.25)) + 1.5))
            } else {
                v37 := GenS1{F0: -2147483648}
                if true {
                    var v38 int16 = (v21 & v21)
                    _ = v38
                } else {
                    _ = len(v33)
                }
                var v39 uint64 = 493
                _ = v37
                _ = v39
            }
            v2 = 0.0
            var v40 uint16 = 32767
            if (!((((-2.25 / v19) * (v24 - -2.25)) > 1e10) && (!(!false)))) {
                var v41 bool = (0.0 == v24)
                var v42 int64 = (v25 >> ((18446744073709551615) & 63))
                v24 = 1.5
                _ = v41
                _ = v42
            }
            var v43 int64 = -1
            if false {
                v44 := GenS1{F0: -222}
                _ = v44
            }
            _ = fi35
            _ = v36
            _ = v40
            _ = v43
        }
        _ = fi34
    }
    switch 0 {
    case 2147483647:
        for fi45 := 0; fi45 < 11; fi45++ {
            if true {
                if len(v33) > 0 {
                    v33[0] = fi45
                }
            }
            var v46 int16 = 969
            if (42 >= 255) {
                for fi47 := 0; fi47 < 19; fi47++ {
                    fmt.Println(v2, v19)
                    v24 += -2.25
                    fmt.Println(v1, v17)
                    _ = fi47
                }
            }
            _ = fi45
            _ = v46
        }
    default:
        switch (v20 * v20) {
        case 694:
            _ = len(v33)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
}
