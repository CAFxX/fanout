package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int; F1 int64; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 uint; F2 int }
func genFn0(p0 uint8) bool {
    var v7 float32 = (1e10 / 0.1)
    if (-1 == 42) {
        v8 := GenS1{F0: 9223372036854775807, F1: "a"}
        _ = v8
    }
    _ = p0
    _ = v7
    return false
}
func genFn1() uint16 {
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 1, F1: "a"}
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 string = ("" + ("" + ((("hello" + "foo bar") + "foo bar") + "")))
            v12 = (v12 + v12)
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    v13 := []int{-9223372036854775808, 42}
    var v14 bool = ((65535 >= 1) && true)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    _ = v13
    _ = v14
    return 586
}

func main() {
    var v1 int = 2147483647
    var v2 int8 = -128
    var v3 float32 = (1.5 - float32((float64((0.1 - float32((0.0 + 1.5)))) * 0.0)))
    var v4 int32 = 2147483647
    v5 := []int{v1, 208, v1, v1}
    var v6 float64 = 0.1
    if (42 <= 42) {
        var v15 int64 = 265
        v16 := GenS2{F0: 1, F1: 698}
        _ = v15
        _ = v16
    } else {
        v17 := GenS3{F0: (^(v1)), F1: 42, F2: 68}
        v18 := GenS2{F0: 4294967295, F1: 18446744073709551615}
        v4 += -2147483648
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v4 += (v4 * v4)
        fmt.Println(v4, v4, v3)
        var v20 int = (fi19 % (((fi19 / ((v1) | 1))) | 1))
        _ = fi19
        _ = v20
    }
    var v21 int = v1
    if (float32(v6) > v3) {
        if (0 >= 673) {
            var v22 uint = 0
            fmt.Println(v4, v3, v3)
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 float32 = (-2.25 / (0.0 + ((3.14159 / 1e10) * (v3 * float32(v6)))))
                switch 212 {
                case -1:
                    v6 += 3.14159
                case -810:
                    if false {
                        v22 = (v22 * (^(v22)))
                        var v25 uint16 = 42
                        var v26 float64 = v6
                        _ = v25
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 9; fi27++ {
                            v4 = -859
                            _ = fi27
                        }
                        switch "foo bar" {
                        case "\n\t":
                            if (!(0 > 127)) {
                                v2 = (v2 - v2)
                                fmt.Println(fi23, v3)
                                for fi28 := 0; fi28 < 20; fi28++ {
                                    var v29 bool = true
                                    _ = fi28
                                    _ = v29
                                }
                            } else {
                                if (0.1 <= -2.25) {
                                    for fi30 := 0; fi30 < 10; fi30++ {
                                        v3 = 0.0
                                        v24 = v3
                                        var v31 int16 = 1
                                        _ = fi30
                                        _ = v31
                                    }
                                } else {
                                    var v32 uint16 = genFn1()
                                    v24 -= (v3 + 3.14159)
                                    _ = v32
                                }
                                v22 = (v22 + 42)
                                v33 := GenS4{F0: v21, F1: 842, F2: 611}
                                _ = v33
                            }
                        case "":
                            v34 := GenS1{F0: -811, F1: "foo bar"}
                            _ = v34
                        }
                        if (!(42 < 1)) {
                            fmt.Println(v1, v6, v6)
                            var v35 uint = v22
                            _ = v35
                        }
                    }
                case 0:
                    v6 = v6
                }
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        var v36 uint64 = 975
        v3 = float32((((((v6 * 3.14159) - v6) + (v6 - (v6 * (0.1 * (v6 - v6))))) + v6) / 1.0))
        _ = v36
    }
    v21 = (v21 & 4294967295)
    var v37 bool = (42 > -121)
    if (v6 == float64((0.1 * v3))) {
        v38 := GenS1{F0: 9223372036854775807, F1: ("foo bar" + "foo bar")}
        v39 := GenS5{F0: 0}
        switch int(v2) {
        case 230:
            v37 = v37
        case 42:
            v4 += (v4 << ((1) & 63))
        case -1:
            v5 = append(v5, 0)
        }
        _ = v38
        _ = v39
    }
    v5 = append(v5, (v21 * 0))
    v21 = v1
    var v40 int16 = -987
    switch "" {
    case "foo bar":
        v4 = (v4 & (v4 * (v4 ^ (^(v4)))))
    case "a":
        v21 -= v21
    case "":
        if v37 {
            _ = len(v5)
        } else {
            fmt.Println(v2, v37)
        }
    default:
        switch "foo bar" {
        case "":
            for fi41 := 0; fi41 < 11; fi41++ {
                v42 := GenS1{F0: 0, F1: "foo bar"}
                if v37 {
                    v3 = v3
                }
                v5 = append(v5, v1)
                _ = fi41
                _ = v42
            }
        case "hello":
            _ = len(v5)
        case "\n\t":
            v43 := GenS6{F0: v40, F1: 255, F2: (v1 * -12)}
            _ = v43
        }
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
