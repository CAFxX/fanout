package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 bool }
func genFn0(p0 int8, p1 uint64, p2 int32) int8 {
    p0 += p0
    var v4 uint64 = (^(p1))
    for fi5 := 0; fi5 < 2; fi5++ {
        if (0.0 > 0.0) {
            p1 += v4
            v6 := []int{(fi5 / (((fi5 & (fi5 | -9223372036854775808))) | 1)), fi5, (^(fi5)), fi5}
            var v7 int8 = p0
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    if (!(!(!true))) {
        if true {
            v4 = 4294967295
            var v8 uint16 = 32767
            _ = v8
        } else {
            fmt.Println(v4)
        }
        var v9 uint32 = 153
        fmt.Println(p0)
        _ = v9
    } else {
        var v10 uint64 = (^(v4))
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}

func main() {
    v1 := []int{-1, -9223372036854775808, -669, 9223372036854775807}
    var v2 float64 = 1e10
    var v3 uint32 = 4294967295
    for fi11 := 0; fi11 < 12; fi11++ {
        _ = len(v1)
        v12 := GenS1{F0: (!(!true))}
        v13 := GenS2{F0: 32767, F1: 127}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float32 = 0.0
    var v15 int = 9223372036854775807
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(v15, v15)
        v15 += fi16
        _ = fi16
    }
    v17 := GenS2{F0: 301, F1: 43}
    v18 := GenS3{F0: 4294967295}
    fmt.Println(v14, v15, v3)
    v1 = append(v1, (v15 & v15))
    if true {
        v3 += 0
    } else {
        v3 -= v3
        v3 += 42
    }
    if true {
        if ((("a" + "") < ("" + "foo bar")) || true) {
            v15 = -1
            v19 := GenS2{F0: 0, F1: 46}
            v20 := GenS1{F0: false}
            _ = v19
            _ = v20
        }
        if true {
            var v21 float64 = v2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v21
        }
        if false {
            var v22 uint32 = v3
            _ = v22
        } else {
            v14 = (v14 - (v14 / 1.0))
            if (255 > 1) {
                v23 := GenS4{F0: 21}
                fmt.Println(v3)
                _ = v23
            } else {
                if false {
                    v15 += (^(v15))
                    v2 = 0.0
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v24 := GenS5{F0: "a", F1: 0}
                    var v25 float64 = v2
                    var v26 uint32 = 2147483648
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
        }
    } else {
        if false {
            var v27 uint = 1
            if (!false) {
                var v28 uint16 = uint16(v27)
                _ = v28
            } else {
                switch (v15 << ((v27) & 63)) {
                case 4294967295:
                    var v29 uint = 2
                    _ = v29
                case -257:
                    if len(v1) > 0 {
                        v1[0] = (-(v15))
                    }
                }
                if ("foo bar" <= ("hello" + (((("\n\t" + ((((("\n\t" + "\n\t") + "\n\t") + ("foo bar" + ("hello" + ("foo bar" + "hello")))) + "\n\t") + "\n\t")) + ("hello" + "foo bar")) + "\n\t") + "\n\t"))) {
                    for fi30 := 0; fi30 < 5; fi30++ {
                        v31 := GenS6{F0: 0, F1: 42}
                        v27 = v27
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi30
                        _ = v31
                    }
                    _ = len(v1)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v32 := GenS2{F0: 0, F1: 142}
                _ = v32
            }
            var v33 int = (v15 - v15)
            _ = v27
            _ = v33
        } else {
            v1 = append(v1, v15)
            var v34 uint8 = 0
            _ = v34
        }
        switch v15 {
        case 32767:
            var v35 int32 = -1
            _ = v35
        case -9223372036854775808:
            var v36 uint8 = 127
            _ = v36
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            if (42 == 1) {
                if true {
                    v14 -= (v14 / 0.1)
                }
                switch -490 {
                case 4294967295:
                    var v37 string = "foo bar"
                    _ = v37
                }
            }
        }
    }
    var v38 int = (v15 % (((v15 / ((127) | 1))) | 1))
    v39 := GenS6{F0: 0, F1: genFn0(genFn0(42, 978, 1), 2, -756)}
    switch -2147483648 {
    case 42:
        v3 = v3
    case -9223372036854775808:
        switch "\n\t" {
        case "\n\t":
            _ = len(v1)
        }
    case -729:
        switch "hello" {
        case "":
            v40 := GenS7{F0: (9223372036854775807 > (v15 % ((-61) | 1)))}
            _ = v40
        }
    default:
        v1 = append(v1, v38)
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v39
}
