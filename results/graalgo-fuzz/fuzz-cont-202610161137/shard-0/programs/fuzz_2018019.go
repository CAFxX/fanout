package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int; F2 int16 }
func genFn0(p0 int, p1 uint16) uint8 {
    switch p0 {
    case 32767:
        v7 := GenS1{F0: true}
        _ = v7
    case 893:
        if (false || (((0.0 - -2.25) * 3.14159) == (3.14159 + 1.5))) {
            p1 -= 370
            var v8 float64 = -2.25
            _ = v8
        }
    }
    var v9 int64 = -9223372036854775808
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 116
}
func genFn1() uint64 {
    var v11 float64 = (((3.14159 - -2.25) * -2.25) / 1e10)
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint = 2
        _ = fi12
        _ = v13
    }
    v14 := []int{-650, 0}
    for fi15 := 0; fi15 < 9; fi15++ {
        v11 += 3.14159
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi15
    }
    var v16 int32 = -2147483648
    _ = v11
    _ = v14
    _ = v16
    return 1
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []string{"\n\t"}
    var v3 string = ""
    var v4 int64 = 1
    var v5 bool = true
    var v6 string = ""
    var v17 int16 = -146
    switch 128 {
    case 32767:
        if false {
            if v5 {
                var v18 float64 = 1e10
                if v5 {
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + v6)
                    }
                    var v19 uint32 = 359
                    _ = v19
                }
                _ = v18
            }
            switch -1 {
            case 1:
                var v20 int64 = 1
                _ = v20
            default:
                var v21 uint16 = 1
                _ = v21
            }
            var v22 int8 = -1
            _ = v22
        }
    case -2147483648:
        v5 = v5
    default:
        v5 = (127 > 0)
    }
    v4 = v4
    if v5 {
        if (!v5) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            switch 127 {
            case 0:
                if (((!false) && (uint16(v1) <= 1)) || v5) {
                    v3 = ("hello" + v3)
                    if ((v17 % (((^(v17))) | 1)) != (-(v17))) {
                        for fi23 := 0; fi23 < 12; fi23++ {
                            v17 = (v17 - (v17 * (v17 | v17)))
                            _ = fi23
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v4 = (^(v4))
                    } else {
                        var v24 int64 = 42
                        v1 -= v1
                        _ = v24
                    }
                    v1 -= (v1 - (^(v1)))
                } else {
                    var v25 string = ("foo bar" + "hello")
                    if false {
                        var v26 uint64 = 90
                        v4 -= (v4 - v4)
                        _ = v26
                    }
                    _ = v25
                }
            case -901:
                for fi27 := 0; fi27 < 19; fi27++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi27
                }
            }
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 int32 = 1
                _ = fi28
                _ = v29
            }
        }
        var v30 int = 1
        _ = v30
    }
    var v31 float64 = 1.5
    if v5 {
        var v32 int32 = -503
        _ = v32
    } else {
        var v33 float32 = (1.5 * -2.25)
        if v5 {
            if ((v1 | 58) > 1) {
                v34 := GenS1{F0: (0 != 128)}
                _ = v34
            }
        } else {
            var v35 uint16 = 249
            _ = len(v2)
            v5 = v5
            _ = v35
        }
        _ = v33
    }
    v36 := GenS1{F0: v5}
    switch -388 {
    case 278:
        v5 = true
    }
    v4 += v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v5 {
        if v5 {
            v37 := GenS2{F0: 128, F1: 492, F2: 0}
            _ = v37
        }
    }
    switch "" {
    case "foo bar":
        fmt.Println(v3, v6)
    default:
        if v5 {
            fmt.Println(v5, v5, v4)
            var v38 string = ""
            _ = v38
        }
    }
    var v39 uint8 = 152
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
