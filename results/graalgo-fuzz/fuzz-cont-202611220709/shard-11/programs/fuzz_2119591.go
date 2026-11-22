package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 string; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 bool }
func genFn0() int16 {
    v6 := GenS1{F0: -2147483648}
    var v7 float32 = 1e10
    v8 := []int{999, 580, -1}
    _ = len(v8)
    _ = v6
    _ = v7
    _ = v8
    return -132
}
func genFn1(p0 uint) int16 {
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 bool = true
        var v11 uint32 = 2147483647
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = 32767
    fmt.Println(v12, p0)
    _ = p0
    _ = v12
    return -879
}

func main() {
    var v1 uint16 = 42
    var v2 int32 = -2147483648
    var v3 bool = false
    v4 := make(map[string]int)
    var v5 float64 = 0.1
    v13 := GenS1{F0: v2}
    var v14 float64 = (((v5 * 1.5) - (-2.25 * v5)) + (((v5 - 0.0) * (0.0 + v5)) / (float64(1.5) / (v5 - v5))))
    v15 := v4["foo bar"]
    var v16 int = 9223372036854775807
    v17 := GenS1{F0: v2}
    var v18 float64 = float64(-2.25)
    v19 := v4["a"]
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS2{F0: 255, F1: v3}
        _ = len(v4)
        _ = fi20
        _ = v21
    }
    if v3 {
        var v22 int8 = 42
        if v3 {
            v23 := GenS2{F0: 18446744073709551615, F1: ((!v3) || (int64(v1) <= 9223372036854775807))}
            _ = v23
        } else {
            if false {
                v1 = 255
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 bool = v3
                    for fi26 := 0; fi26 < 8; fi26++ {
                        v1 = 1
                        fmt.Println(v14)
                        v1 = (v1 & v1)
                        _ = fi26
                    }
                    if ((v1 >> ((2) & 63)) > (v1 << ((2) & 63))) {
                        if (4294967295 >= 4294967295) {
                            if v25 {
                                v27 := GenS2{F0: 2, F1: (!(132 == 65535))}
                                v2 -= (v2 >> ((18446744073709551615) & 63))
                                _ = v27
                            }
                            v28 := GenS1{F0: v2}
                            _ = v28
                        } else {
                            v29 := []int{v19}
                            if v25 {
                                if len(v29) > 0 {
                                    v29[0] = v15
                                }
                            }
                            var v30 uint32 = uint32(v1)
                            _ = v29
                            _ = v30
                        }
                        v14 = (-2.25 / v14)
                        switch v15 {
                        case 1:
                            var v31 int32 = v2
                            _ = v31
                        case 4294967295:
                            v32 := []int{-707}
                            _ = v32
                        default:
                            v14 -= float64(1e10)
                        }
                    }
                    _ = fi24
                    _ = v25
                }
                v4["hello"] = (^(v19))
            } else {
                switch 127 {
                case -2147483648:
                    v16 -= v15
                default:
                    switch "a" {
                    case "a":
                        v5 += v5
                    case "foo bar":
                        v33 := []int{(v19 + v15)}
                        _ = v33
                    }
                }
                var v34 float64 = float64((3.14159 * -2.25))
                v19 += (v19 >> ((316) & 63))
                _ = v34
            }
            for fi35 := 0; fi35 < 14; fi35++ {
                v36 := GenS3{F0: v2, F1: ("\n\t" + ""), F2: genFn1(589)}
                v37 := []int{(v16 ^ fi35), (v15 & (v16 >> ((54) & 63))), -396}
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        _ = v22
    }
    v3 = (-404 >= v19)
    if ((v16 + -9223372036854775808) > v16) {
        v38 := GenS4{F0: v3}
        for fi39 := 0; fi39 < 20; fi39++ {
            v40 := GenS5{F0: (("foo bar" + ("hello" + "foo bar")) + "a"), F1: v3}
            if false {
                v19 = fi39
                delete(v4, ("hello" + ""))
                var v41 uint8 = 42
                _ = v41
            } else {
                var v42 uint8 = uint8(v1)
                var v43 uint = 224
                if true {
                    v3 = (true && v3)
                    v44 := GenS4{F0: v3}
                    var v45 int32 = (^(v2))
                    _ = v44
                    _ = v45
                } else {
                    v46 := v4["hello"]
                    v2 = 685
                    switch v16 {
                    case 815:
                        var v47 float64 = v5
                        _ = v47
                    default:
                        v48 := []int{}
                        _ = v48
                    }
                    _ = v46
                }
                _ = v42
                _ = v43
            }
            _ = fi39
            _ = v40
        }
        var v49 uint16 = 826
        _ = v38
        _ = v49
    } else {
        v50 := GenS5{F0: "a", F1: true}
        _ = v50
    }
    v51 := GenS3{F0: (v2 + 914), F1: ((("hello" + ("\n\t" + "\n\t")) + "\n\t") + ""), F2: -572}
    var v52 uint8 = 94
    v2 = v2
    fmt.Println(v16, v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
