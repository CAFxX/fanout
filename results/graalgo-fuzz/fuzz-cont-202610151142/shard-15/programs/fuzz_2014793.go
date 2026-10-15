package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int8 }
func genFn0() string {
    if true {
        switch -861 {
        case 127:
            var v7 uint8 = 0
            _ = v7
        case 32767:
            var v8 string = "foo bar"
            _ = v8
        case 9223372036854775807:
            if (434 <= 87) {
                for fi9 := 0; fi9 < 16; fi9++ {
                    v10 := GenS1{F0: -32768, F1: 42, F2: 120}
                    _ = fi9
                    _ = v10
                }
                var v11 int16 = 1
                _ = v11
            }
        default:
            var v12 int64 = 0
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        fmt.Println(fi13)
        _ = fi13
    }
    v14 := GenS1{F0: 1000, F1: 42, F2: -1}
    _ = v14
    return ""
}
func genFn1(p0 uint8, p1 uint8, p2 string) uint32 {
    var v15 uint = 51
    v16 := GenS1{F0: 32767, F1: 1000, F2: -1}
    switch -9223372036854775808 {
    case 9223372036854775807:
        v15 -= 42
    case -562:
        v17 := GenS1{F0: 1, F1: -1, F2: -123}
        _ = v17
    case -757:
        var v18 int16 = 1000
        _ = v18
    }
    var v19 uint = (v15 & v15)
    v19 = v19
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v19
    return 2147483647
}

func main() {
    v1 := []int{32768, -2147483648, 32767, 128}
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    v4 := []string{("\n\t" + ""), "a", ((("\n\t" + "") + ("hello" + ("foo bar" + "hello"))) + "foo bar"), "foo bar"}
    var v5 uint32 = 581
    var v6 int8 = -128
    v5 += v5
    fmt.Println(v5, v5)
    if false {
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 float32 = 1.5
            switch "foo bar" {
            case "":
                switch (fi20 & 19) {
                case 128:
                    if false {
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v23 := GenS1{F0: -811, F1: 1, F2: v6}
                            fmt.Println(v6)
                            v6 -= (-(v6))
                            _ = fi22
                            _ = v23
                        }
                        if ((false || false) || (404 > 855)) {
                            if true {
                                if true {
                                    v24 := GenS1{F0: 32767, F1: 32767, F2: (-(v6))}
                                    _ = v24
                                }
                            } else {
                                v6 -= (v6 * (^(v6)))
                            }
                            var v25 int = (fi20 - 9223372036854775807)
                            _ = v25
                        } else {
                            v6 = v6
                        }
                        v5 -= (^(v5))
                    }
                case -2147483648:
                    var v26 int8 = v6
                    _ = v26
                case 1:
                    var v27 uint8 = 0
                    _ = v27
                }
            case "hello":
                for fi28 := 0; fi28 < 10; fi28++ {
                    fmt.Println(v21, v21, v5)
                    if true {
                        var v29 string = "hello"
                        _ = v29
                    }
                    _ = fi28
                }
            default:
                for fi30 := 0; fi30 < 2; fi30++ {
                    v21 = (v21 - (v21 + 0.0))
                    _ = fi30
                }
            }
            v21 = v21
            _ = fi20
            _ = v21
        }
    } else {
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 bool = true
            fmt.Println(v6, v32, v32)
            _ = fi31
            _ = v32
        }
        v5 -= 602
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        fmt.Println(v6)
        if false {
            var v34 int16 = 32767
            _ = v34
        } else {
            fmt.Println(v5, fi33)
            for fi35 := 0; fi35 < 16; fi35++ {
                v5 = v5
                _ = fi35
            }
        }
        for fi36 := 0; fi36 < 4; fi36++ {
            v37 := GenS1{F0: 32767, F1: -676, F2: v6}
            _ = fi36
            _ = v37
        }
        _ = fi33
    }
    fmt.Println(v5)
    v6 = (v6 + (v6 * (v6 + (v6 >> ((589) & 63)))))
    delete(v3, "hello")
    var v38 int = 876
    if len(v1) > 0 {
        v1[0] = v38
    }
    delete(v2, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
