package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int64; F2 int16 }

func main() {
    v1 := []int{128, -934, 32767}
    var v2 float64 = 1.5
    var v3 int = 403
    v3 -= v3
    v2 = -2.25
    if false {
        for fi4 := 0; fi4 < 6; fi4++ {
            var v5 uint8 = 0
            var v6 int8 = 100
            _ = fi4
            _ = v5
            _ = v6
        }
        if (0.1 > 1e10) {
            v1 = append(v1, v3)
            var v7 int8 = 0
            _ = v7
        }
        var v8 int16 = 1000
        _ = v8
    }
    if false {
        var v9 string = (("a" + "foo bar") + "")
        switch ("" + "hello") {
        case "\n\t":
            _ = len(v1)
        case "":
            if ((v9 + "hello") != v9) {
                v10 := GenS1{F0: (true || ((!(!false)) || false))}
                for fi11 := 0; fi11 < 14; fi11++ {
                    if true {
                        for fi12 := 0; fi12 < 18; fi12++ {
                            v9 = ((v9 + (v9 + v9)) + ((v9 + v9) + (v9 + ("hello" + ""))))
                            _ = fi12
                        }
                    } else {
                        v2 = v2
                    }
                    v1 = append(v1, -6)
                    _ = fi11
                }
                if (!(0 == 255)) {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        switch "hello" {
                        case "foo bar":
                            v3 += (fi13 % ((42) | 1))
                        default:
                            v3 += fi13
                        }
                        _ = fi13
                    }
                } else {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        fmt.Println(fi14, v3, fi14)
                        v15 := GenS1{F0: true}
                        _ = fi14
                        _ = v15
                    }
                }
                _ = v10
            }
        default:
            for fi16 := 0; fi16 < 19; fi16++ {
                for fi17 := 0; fi17 < 17; fi17++ {
                    switch 127 {
                    case -9223372036854775808:
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v9, v2)
                        }
                    default:
                        if true {
                            v18 := GenS2{F0: (v9 != "a"), F1: -877, F2: 0}
                            v3 += (-(v3))
                            _ = v18
                        } else {
                            v19 := GenS1{F0: (42 < 52)}
                            v20 := GenS1{F0: (!(!true))}
                            _ = v19
                            _ = v20
                        }
                        v1 = append(v1, -865)
                    }
                    var v21 int64 = 9223372036854775807
                    var v22 float64 = v2
                    _ = fi17
                    _ = v21
                    _ = v22
                }
                v2 = (v2 - 1.5)
                _ = fi16
            }
        }
        _ = v9
    } else {
        v3 = 2147483647
    }
    v23 := GenS2{F0: (-428 > 0), F1: 9223372036854775807, F2: 1}
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v1 = append(v1, 2147483647)
    fmt.Println(v2, v3)
    var v24 string = "\n\t"
    v2 = float64(float32(v2))
    switch 949 {
    case 9223372036854775807:
        v3 = v3
    case 0:
        v3 = (-(v3))
    case 662:
        v3 -= v3
    }
    if (true && false) {
        for fi25 := 0; fi25 < 16; fi25++ {
            switch "" {
            case "\n\t":
                v26 := GenS1{F0: true}
                _ = v26
            }
            v27 := GenS2{F0: false, F1: int64(fi25), F2: -32768}
            v1 = append(v1, (-(fi25)))
            _ = fi25
            _ = v27
        }
        var v28 uint8 = 127
        _ = v28
    }
    _ = len(v1)
    fmt.Println(v2)
    if false {
        var v29 int = 127
        var v30 string = v24
        if false {
            var v31 int8 = -1
            var v32 int = (v3 - (v29 % ((v29) | 1)))
            fmt.Println(v29)
            _ = v31
            _ = v32
        } else {
            for fi33 := 0; fi33 < 5; fi33++ {
                switch (v3 | (^(v29))) {
                case 128:
                    v24 = v24
                default:
                    v24 = v24
                }
                var v34 int32 = -1
                var v35 uint64 = 2
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        _ = v29
        _ = v30
    } else {
        var v36 int8 = 42
        switch -177 {
        case -228:
            if (728 != 42) {
                switch (v3 >> ((1) & 63)) {
                case -682:
                    switch ("foo bar" + ("a" + (v24 + ((v24 + v24) + v24)))) {
                    case "\n\t":
                        if false {
                            var v37 int = (-(v3))
                            fmt.Println(v2, v37)
                            _ = v37
                        }
                    case "a":
                        v3 = v3
                    }
                default:
                    fmt.Println(v3, v2)
                }
                v2 += float64(0.0)
            }
        case 32768:
            if ((v36 << ((0) & 63)) == (^(v36))) {
                v2 += (v2 * float64(1e10))
            }
        case -1:
            v3 = 4294967295
        }
        var v38 string = (v24 + v24)
        _ = v36
        _ = v38
    }
    v39 := GenS1{F0: true}
    fmt.Println(v24)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v39
}
