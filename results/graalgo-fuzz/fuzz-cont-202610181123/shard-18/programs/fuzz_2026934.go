package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int }
func genFn0(p0 bool) uint8 {
    switch ("hello" + "") {
    case "\n\t":
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{fi7}
            fmt.Println(p0, fi7)
            _ = fi7
            _ = v8
        }
    default:
        fmt.Println(p0)
    }
    if p0 {
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 1000, F1: 32767}
            _ = fi9
            _ = v10
        }
    } else {
        p0 = true
    }
    if true {
        var v11 int32 = -1
        var v12 uint64 = 417
        _ = v11
        _ = v12
    } else {
        var v13 uint = 0
        if ((v13 <= (v13 | (v13 | v13))) && p0) {
            switch 2147483647 {
            case 128:
                for fi14 := 0; fi14 < 14; fi14++ {
                    v13 = v13
                    _ = fi14
                }
            case 32767:
                p0 = false
            case 734:
                fmt.Println(p0, v13)
            default:
                for fi15 := 0; fi15 < 2; fi15++ {
                    fmt.Println(v13, p0, fi15)
                    _ = fi15
                }
            }
            v13 = (v13 - (^(v13)))
            if (!(1 == int32(v13))) {
                switch 1 {
                case -546:
                    v16 := []int{-2147483648, 128, 0}
                    _ = v16
                default:
                    fmt.Println(v13, p0)
                }
                for fi17 := 0; fi17 < 18; fi17++ {
                    for fi18 := 0; fi18 < 8; fi18++ {
                        switch ("hello" + ("\n\t" + "hello")) {
                        case "foo bar":
                            p0 = (1.5 != 3.14159)
                        }
                        v19 := GenS2{F0: (fi18 / (((fi17 + (fi18 * (fi18 + 293)))) | 1))}
                        _ = fi18
                        _ = v19
                    }
                    _ = fi17
                }
                if p0 {
                    v20 := []int{32768}
                    _ = v20
                }
            } else {
                var v21 int8 = 0
                var v22 int32 = 0
                if ((v21 < v21) && (4294967295 > 2147483647)) {
                    var v23 uint32 = 65535
                    _ = v23
                }
                _ = v21
                _ = v22
            }
        }
        _ = v13
    }
    v24 := []int{}
    _ = p0
    _ = v24
    return 5
}

func main() {
    var v1 int32 = 826
    v2 := []uint32{182, 2147483648, 4294967295}
    var v3 float32 = ((0.1 - 1e10) / 1.5)
    var v4 bool = true
    v5 := []int{-2147483648, -965}
    var v6 uint8 = 0
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v5 = append(v5, 4294967295)
    if len(v5) > 0 {
        v5[0] = 42
    }
    var v25 int = 42
    v26 := GenS1{F0: -1, F1: 1000}
    v6 = (^(v6))
    var v27 uint32 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = 282
    }
    if false {
        for fi28 := 0; fi28 < 13; fi28++ {
            switch ("foo bar" + ((("hello" + (("hello" + "a") + "")) + "") + "")) {
            case "a":
                var v29 int8 = 1
                _ = v29
            default:
                var v30 string = "\n\t"
                _ = v30
            }
            v31 := GenS2{F0: (fi28 % (((fi28 ^ int(v1))) | 1))}
            _ = fi28
            _ = v31
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v6 += 128
    }
    var v32 uint16 = 0
    var v33 int8 = 3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
