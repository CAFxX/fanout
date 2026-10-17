package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int16, p1 int16) uint8 {
    var v7 float32 = (1e10 - 1e10)
    var v8 int64 = -371
    if (("" + ("\n\t" + "hello")) <= "a") {
        v9 := []int{42, 127}
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: (p0 - p0)}
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 float64 = (0.0 * 0.0)
                var v14 uint64 = 754
                if (-2147483648 > 0) {
                    var v15 int8 = -24
                    v7 = v7
                    if (false && true) {
                        v8 -= int64(p0)
                    }
                    _ = v15
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
        if false {
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 uint16 = 32767
                _ = fi16
                _ = v17
            }
            switch 127 {
            case -496:
                v18 := GenS1{F0: 32767}
                _ = v18
            default:
                v19 := GenS1{F0: p1}
                _ = v19
            }
        } else {
            var v20 int16 = p1
            v8 = int64(v20)
            _ = v20
        }
        _ = v9
    } else {
        v21 := GenS1{F0: 42}
        var v22 uint = 0
        v22 = v22
        _ = v21
        _ = v22
    }
    var v23 int16 = (-(p0))
    if ((42 >= -1) || ((((1 <= 0) || (true || false)) && false) && ((p1 / ((735) | 1)) < v23))) {
        v24 := GenS1{F0: 1}
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v23
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 string = ("foo bar" + "hello")
    var v2 float64 = 3.14159
    var v3 int32 = 377
    var v4 string = v1
    v5 := make(map[string]string)
    var v6 int8 = 127
    if (911 > 0) {
        var v25 uint64 = 0
        v26 := v5[v4]
        fmt.Println(v26)
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 7; fi27++ {
            for fi28 := 0; fi28 < 16; fi28++ {
                fmt.Println(fi27, v1, v1)
                _ = fi28
            }
            v6 += (v6 & (^(v6)))
            _ = fi27
        }
    }
    var v29 uint = 42
    var v30 int64 = 1
    v2 = (float64(1e10) - 1.5)
    v31 := GenS2{F0: 0, F1: 42}
    v1 = (v1 + ("a" + "a"))
    var v32 string = v1
    v33 := []int{-83, -1, 42, -2147483648}
    if (v29 >= 18446744073709551615) {
        switch 673 {
        case -480:
            switch -1 {
            case -313:
                var v34 uint8 = 2
                _ = v34
            default:
                var v35 int8 = v6
                _ = v35
            }
        case 128:
            if false {
                for fi36 := 0; fi36 < 13; fi36++ {
                    for fi37 := 0; fi37 < 5; fi37++ {
                        v32 = v4
                        _ = fi37
                    }
                    var v38 float64 = v2
                    var v39 bool = true
                    _ = fi36
                    _ = v38
                    _ = v39
                }
                var v40 int64 = (v30 % (((v30 + v30)) | 1))
                _ = v40
            } else {
                v4 = "a"
                v6 = -128
            }
        }
        var v41 uint = v29
        _ = v41
    } else {
        if (498 > -965) {
            for fi42 := 0; fi42 < 5; fi42++ {
                var v43 uint = (^(v29))
                _ = fi42
                _ = v43
            }
        }
        v44 := GenS2{F0: 2147483648, F1: 2}
        v30 = (v30 / ((364) | 1))
        _ = v44
    }
    v4 = v1
    v6 -= v6
    for fi45 := 0; fi45 < 7; fi45++ {
        v2 = v2
        if true {
            v46 := GenS2{F0: 2147483647, F1: 255}
            _ = v46
        }
        _ = fi45
    }
    fmt.Println(v6, v6)
    switch 42 {
    case 32768:
        v3 = v3
    case 42:
        if len(v33) > 0 {
            v33[0] = 128
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
