package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0() int16 {
    if false {
        v6 := []int{}
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint = 901
            _ = fi7
            _ = v8
        }
        if (0.0 < float32(1e10)) {
            if true {
                if len(v6) > 0 {
                    v6[0] = 127
                }
                for fi9 := 0; fi9 < 14; fi9++ {
                    if false {
                        var v10 uint32 = 918
                        _ = v10
                    } else {
                        for fi11 := 0; fi11 < 15; fi11++ {
                            var v12 int16 = 32767
                            _ = fi11
                            _ = v12
                        }
                    }
                    _ = fi9
                }
                var v13 int8 = 0
                _ = v13
            } else {
                var v14 uint32 = 2147483648
                _ = v14
            }
        }
        _ = v6
    } else {
        var v15 float32 = 1e10
        _ = v15
    }
    if false {
        if false {
            var v16 uint8 = 127
            v16 = 0
            _ = v16
        } else {
            v17 := GenS1{F0: 0, F1: 627}
            v18 := []int{0}
            switch 53 {
            case 379:
                var v19 uint16 = 65535
                _ = v19
            case 32768:
                if len(v18) > 0 {
                    v18[0] = 871
                }
            default:
                var v20 uint64 = 4294967295
                _ = v20
            }
            _ = v17
            _ = v18
        }
        if (255 == 2) {
            var v21 int32 = -735
            if false {
                v21 -= -1
            } else {
                var v22 bool = true
                if v22 {
                    v21 = v21
                    v22 = false
                } else {
                    v22 = (((!false) && (65535 <= 2147483648)) || ((v22 || (!false)) && v22))
                    fmt.Println(v22)
                }
                _ = v22
            }
            var v23 uint8 = 235
            _ = v21
            _ = v23
        }
        var v24 int32 = 2147483647
        _ = v24
    } else {
        var v25 float32 = float32(0.1)
        _ = v25
    }
    v26 := []int{1, -9223372036854775808, 1}
    v26 = append(v26, 32768)
    _ = v26
    return -1
}
func genFn1(p0 uint8) uint32 {
    for fi27 := 0; fi27 < 11; fi27++ {
        p0 += p0
        _ = fi27
    }
    var v28 bool = true
    fmt.Println(v28)
    _ = p0
    _ = v28
    return 2147483647
}

func main() {
    var v1 uint8 = 42
    var v2 string = "\n\t"
    var v3 uint64 = 2
    var v4 int8 = -63
    v5 := []int{479, 42, 0, 32768}
    var v29 uint16 = 991
    v1 -= v1
    for fi30 := 0; fi30 < 18; fi30++ {
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 int64 = -727
            if (!(true || (-181 != (v32 >> ((4294967295) & 63))))) {
                v33 := GenS1{F0: (v1 | (^(v1))), F1: 811}
                v29 = v29
                switch (v2 + ("foo bar" + ("\n\t" + (v2 + v2)))) {
                case "":
                    fmt.Println(v2)
                default:
                    v34 := GenS1{F0: (v1 | (v1 ^ v1)), F1: uint(v1)}
                    _ = v34
                }
                _ = v33
            }
            _ = fi31
            _ = v32
        }
        var v35 int32 = -1
        if (genFn0() == 0) {
            if (!false) {
                v4 -= (v4 & v4)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v2 = v2
            } else {
                var v36 int8 = v4
                for fi37 := 0; fi37 < 4; fi37++ {
                    var v38 uint8 = v1
                    if (835 >= 0) {
                        fmt.Println(v38, v3)
                    } else {
                        var v39 string = v2
                        fmt.Println(v4, v39, fi37)
                        _ = v39
                    }
                    var v40 float64 = (0.1 / 1.0)
                    _ = fi37
                    _ = v38
                    _ = v40
                }
                v5 = append(v5, 32768)
                _ = v36
            }
        }
        _ = fi30
        _ = v35
    }
    v4 = -1
    var v41 uint32 = 0
    var v42 float32 = (((-2.25 + 3.14159) / (-2.25 / (0.1 * 1.5))) + 0.1)
    _ = len(v5)
    v43 := GenS2{F0: 32767, F1: (v29 ^ 42), F2: false}
    v4 -= (v4 | (v4 | -8))
    v3 += v3
    if ((42 == -462) && true) {
        switch 2147483647 {
        case 128:
            switch 765 {
            case -2147483648:
                if ((v2 + ((v2 + v2) + v2)) >= "hello") {
                    v44 := GenS2{F0: 912, F1: 1, F2: true}
                    _ = v44
                }
            case 128:
                v42 = -2.25
            case 32767:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                for fi45 := 0; fi45 < 7; fi45++ {
                    v46 := GenS1{F0: (v1 % ((220) | 1)), F1: 65535}
                    _ = fi45
                    _ = v46
                }
            }
        case -278:
            switch "hello" {
            case "\n\t":
                v47 := GenS3{F0: 98, F1: 0}
                _ = v47
            case "foo bar":
                _ = len(v5)
            case "":
                _ = len(v5)
            }
        }
        if (false || (-2147483648 != -484)) {
            v2 = (v2 + (((("a" + (v2 + v2)) + "\n\t") + v2) + v2))
            v2 = (("foo bar" + "a") + (v2 + v2))
            var v48 uint16 = (^(v29))
            _ = v48
        }
        var v49 uint8 = (v1 & v1)
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
