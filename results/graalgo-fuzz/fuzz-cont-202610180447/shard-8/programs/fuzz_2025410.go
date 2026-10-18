package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 string; F1 int16; F2 bool }
type GenS3 struct { F0 int32; F1 uint; F2 int64 }
func genFn0(p0 int16) uint {
    fmt.Println(p0)
    var v4 uint16 = 255
    v5 := []int{32767, 9223372036854775807}
    var v6 int32 = 0
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1() uint64 {
    if (true || true) {
        if (169 == 235) {
            v7 := GenS1{F0: "foo bar", F1: 1, F2: 9223372036854775807}
            _ = v7
        } else {
            var v8 bool = true
            if v8 {
                if false {
                    if v8 {
                        v8 = v8
                    }
                } else {
                    switch 32767 {
                    case 2147483647:
                        fmt.Println(v8)
                    case -1:
                        for fi9 := 0; fi9 < 7; fi9++ {
                            var v10 bool = v8
                            _ = fi9
                            _ = v10
                        }
                    case 389:
                        v11 := []int{4294967295, 32768}
                        _ = v11
                    default:
                        var v12 uint = 349
                        _ = v12
                    }
                    if (42 != 1) {
                        var v13 uint32 = 42
                        fmt.Println(v8, v8)
                        fmt.Println(v13)
                        _ = v13
                    }
                }
            } else {
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS1{F0: "a", F1: 719, F2: 0}
                    _ = fi14
                    _ = v15
                }
            }
            v8 = (520 == 308)
            _ = v8
        }
        if (42 <= 2) {
            var v16 uint64 = 0
            _ = v16
        }
    } else {
        var v17 uint32 = 2147483647
        switch 2147483647 {
        case -144:
            var v18 int8 = 51
            _ = v18
        case 2147483647:
            fmt.Println(v17)
        case 727:
            if ((32767 != 42) || (((0 < 42) || ((255 <= 129) && (!(!true)))) || (((-2.25 * (-2.25 + ((-2.25 + float64(((3.14159 - 3.14159) - (0.0 + ((1.5 - 0.0) / 1.0))))) - ((-2.25 - -2.25) / 0.1)))) + 3.14159) > ((0.0 + (3.14159 / (((1.5 + 3.14159) + 0.1) / (1e10 / (1e10 + -2.25))))) + (float64(0.1) - 0.1))))) {
                v17 = (v17 & v17)
            } else {
                var v19 uint = 326
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 int8 = 127
                    var v22 uint = (v19 + uint(v17))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = v19
            }
        }
        _ = v17
    }
    v23 := []int{-1, -450, 184, 9223372036854775807}
    if true {
        var v24 string = "a"
        _ = v24
    }
    _ = v23
    return 1
}

func main() {
    var v1 uint8 = 255
    var v2 uint32 = 2147483647
    var v3 string = "\n\t"
    var v25 float64 = float64(3.14159)
    for fi26 := 0; fi26 < 11; fi26++ {
        v3 = ("foo bar" + v3)
        if (!((!(v1 == 127)) || false)) {
            v2 = (v2 / ((v2) | 1))
            fmt.Println(v1)
            var v27 int = fi26
            _ = v27
        }
        if false {
            v2 = 173
        }
        _ = fi26
    }
    var v28 uint32 = v2
    v29 := GenS2{F0: v3, F1: 32767, F2: (99 < 0)}
    switch v3 {
    case "":
        if (1 >= 32767) {
            var v30 uint32 = 865
            v30 += v2
            fmt.Println(v2, v2)
            _ = v30
        } else {
            if true {
                for fi31 := 0; fi31 < 10; fi31++ {
                    v28 += (v2 % ((v2) | 1))
                    v1 -= v1
                    _ = fi31
                }
                if (!(false || true)) {
                    v28 += 0
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 string = "a"
                        _ = fi32
                        _ = v33
                    }
                    var v34 uint16 = 0
                    _ = v34
                }
                fmt.Println(v28)
            } else {
                v35 := GenS2{F0: v3, F1: 234, F2: false}
                v36 := GenS2{F0: (v3 + "foo bar"), F1: 833, F2: (42 < 1)}
                v37 := GenS1{F0: "a", F1: 601, F2: 1}
                _ = v35
                _ = v36
                _ = v37
            }
            for fi38 := 0; fi38 < 16; fi38++ {
                v2 += (^(v2))
                _ = fi38
            }
            v2 = v28
        }
    case "foo bar":
        v25 += v25
    }
    v3 = v3
    var v39 float32 = 0.1
    var v40 int16 = 0
    v41 := GenS1{F0: v3, F1: 841, F2: 0}
    var v42 float64 = (-2.25 * float64(((v39 / (v39 / v39)) + (float32(float64((1.5 * v39))) - (v39 - ((v39 * 0.0) * (float32((-2.25 - 1.5)) + 1.5)))))))
    v43 := GenS3{F0: 1, F1: genFn0(1000), F2: 70}
    v3 = ("a" + (v3 + v3))
    v25 = 3.14159
    for fi44 := 0; fi44 < 2; fi44++ {
        v45 := []int{-401, -2147483648, fi44}
        for fi46 := 0; fi46 < 11; fi46++ {
            v47 := GenS1{F0: v3, F1: 1, F2: -727}
            var v48 uint8 = v1
            _ = fi46
            _ = v47
            _ = v48
        }
        _ = fi44
        _ = v45
    }
    fmt.Println(v1, v1)
    if (!false) {
        var v49 string = v3
        _ = v49
    }
    v50 := GenS2{F0: v3, F1: v40, F2: (2147483647 >= -191)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v50
}
