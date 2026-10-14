package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int8 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]string)
    var v3 string = ""
    if ((4294967295 < 783) || true) {
        switch -2147483648 {
        case 128:
            for fi4 := 0; fi4 < 1; fi4++ {
                fmt.Println(fi4)
                _ = fi4
            }
        case 260:
            var v5 uint8 = 42
            _ = v5
        case -721:
            if true {
                v3 = ((v3 + v3) + ("\n\t" + "a"))
                if false {
                    if (!(true || (!(0 != 0)))) {
                        var v6 int8 = 127
                        _ = v6
                    }
                } else {
                    var v7 int = 9223372036854775807
                    v7 += v7
                    for fi8 := 0; fi8 < 1; fi8++ {
                        v7 -= (fi8 * (v7 | -501))
                        _ = fi8
                    }
                    _ = v7
                }
            }
        }
        if true {
            if false {
                v9 := []int{-291}
                _ = v9
            }
            var v10 uint16 = 255
            _ = v10
        }
    }
    switch v3 {
    case "\n\t":
        v11 := []int{}
        _ = v11
    case "foo bar":
        switch 1 {
        case -9223372036854775808:
            var v12 int = -168
            _ = v12
        case 4294967295:
            var v13 uint64 = 4294967295
            _ = v13
        case 549:
            if (!(1 < 255)) {
                for fi14 := 0; fi14 < 18; fi14++ {
                    v3 = v3
                    switch (fi14 & fi14) {
                    case -2147483648:
                        v15 := v2[127]
                        _ = v15
                    }
                    _ = fi14
                }
                switch v3 {
                case "foo bar":
                    fmt.Println(v1, v1)
                case "\n\t":
                    fmt.Println(v1, v1)
                case "hello":
                    v1 += 1000
                }
                var v16 int8 = -100
                _ = v16
            }
        default:
            v3 = "a"
        }
    default:
        var v17 float64 = 1e10
        _ = v17
    }
    fmt.Println(v1, v1)
    v1 = (v1 + -32768)
    var v18 uint32 = 1
    var v19 int = 42
    if (-2147483648 != 2147483647) {
        if (false || true) {
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := GenS1{F0: v19, F1: 85, F2: 127}
                if ((v19 != 656) && false) {
                    fmt.Println(fi20, v18)
                } else {
                    if false {
                        var v22 bool = false
                        _ = v22
                    } else {
                        v23 := GenS1{F0: (^(fi20)), F1: 255, F2: -1}
                        v3 = "foo bar"
                        var v24 uint64 = 327
                        _ = v23
                        _ = v24
                    }
                    var v25 uint16 = 982
                    _ = v25
                }
                v3 = v3
                _ = fi20
                _ = v21
            }
            var v26 bool = (v19 >= v19)
            _ = v26
        } else {
            v1 -= v1
            switch (v3 + "foo bar") {
            case "\n\t":
                if false {
                    switch v19 {
                    case 564:
                        for fi27 := 0; fi27 < 15; fi27++ {
                            v1 = (v1 * v1)
                            v28 := GenS1{F0: (v19 | (fi27 / (((^(v19))) | 1))), F1: 2, F2: -8}
                            var v29 bool = (472 < 0)
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    case 94:
                        for fi30 := 0; fi30 < 20; fi30++ {
                            v1 -= (v1 % ((v1) | 1))
                            _ = fi30
                        }
                    case 231:
                        fmt.Println(v3, v19)
                    default:
                        v31 := GenS1{F0: v19, F1: 255, F2: -100}
                        _ = v31
                    }
                    v3 = v3
                } else {
                    switch v3 {
                    case "foo bar":
                        v32 := GenS1{F0: (^(v19)), F1: 42, F2: 0}
                        _ = v32
                    }
                    var v33 float64 = 1.5
                    _ = v33
                }
            default:
                v18 += (v18 + (v18 >> ((4294967295) & 63)))
            }
            var v34 uint16 = 42
            _ = v34
        }
        v2[v18] = v3
        v2[(v18 * 42)] = ((v3 + v3) + ((v3 + (v3 + "a")) + v3))
    } else {
        v35 := GenS1{F0: (v19 - v19), F1: 255, F2: 127}
        _ = v35
    }
    v18 += 2147483647
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v19, v1)
    default:
        var v36 uint64 = 60
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
