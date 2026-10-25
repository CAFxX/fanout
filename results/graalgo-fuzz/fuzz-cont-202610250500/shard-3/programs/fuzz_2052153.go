package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int64; F1 bool }

func main() {
    var v1 int8 = 79
    var v2 uint8 = 2
    var v3 int = 128
    v4 := make(map[int]bool)
    var v5 bool = false
    v6 := make(map[string]bool)
    if (false && ((3.14159 * 0.1) < float32((1e10 * (((0.1 + 1e10) + 3.14159) - 3.14159))))) {
        fmt.Println(v1, v1)
        v7 := []int{(v3 / ((-369) | 1)), (v3 | v3)}
        _ = v7
    } else {
        var v8 uint = 255
        v9 := GenS1{F0: false, F1: (v1 + v1)}
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: (!(0 >= -712)), F1: (^(v1))}
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    }
    var v12 uint32 = 1
    var v13 float64 = ((1.5 - ((((0.0 - 3.14159) * -2.25) * 0.1) + (0.1 * 3.14159))) - float64(1.5))
    var v14 int16 = 1000
    if v5 {
        v12 -= v12
        v3 = v3
    } else {
        v15 := GenS2{F0: true, F1: 2, F2: 1}
        var v16 uint64 = 1
        for fi17 := 0; fi17 < 6; fi17++ {
            fmt.Println(v14, v14, v16)
            switch 4294967295 {
            case -453:
                var v18 bool = v5
                _ = v18
            case -882:
                switch "foo bar" {
                case "":
                    var v19 float32 = -2.25
                    _ = v19
                case "hello":
                    var v20 bool = ((^(v16)) < v16)
                    _ = v20
                }
            default:
                if false {
                    var v21 uint8 = 188
                    v22 := []int{fi17, fi17, (v3 | -2147483648), (v3 / ((4294967295) | 1))}
                    delete(v4, 128)
                    _ = v21
                    _ = v22
                }
            }
            v1 = -1
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    var v23 int8 = (v1 << ((665) & 63))
    v2 = (v2 % ((v2) | 1))
    fmt.Println(v5, v1)
    switch (v3 / (((v3 + (-(v3)))) | 1)) {
    case 1:
        switch 0 {
        case -1:
            if false {
                if v5 {
                    var v24 uint32 = v12
                    switch v3 {
                    case 128:
                        if (215 < 4294967295) {
                            v25 := GenS1{F0: (!v5), F1: v23}
                            v12 -= 563
                            _ = v25
                        } else {
                            var v26 int = v3
                            _ = v26
                        }
                    }
                    for fi27 := 0; fi27 < 11; fi27++ {
                        v28 := v6["\n\t"]
                        _ = fi27
                        _ = v28
                    }
                    _ = v24
                }
            } else {
                for fi29 := 0; fi29 < 9; fi29++ {
                    if false {
                        for fi30 := 0; fi30 < 17; fi30++ {
                            v31 := GenS2{F0: true, F1: 65535, F2: 65535}
                            _ = fi30
                            _ = v31
                        }
                    } else {
                        if v5 {
                            if (v14 > v14) {
                                fmt.Println(v23, fi29, fi29)
                                var v32 int = v3
                                _ = v32
                            }
                        } else {
                            var v33 float32 = (0.0 * float32((v13 - v13)))
                            for fi34 := 0; fi34 < 13; fi34++ {
                                v35 := []int{926, (fi29 >> ((18446744073709551615) & 63))}
                                _ = fi34
                                _ = v35
                            }
                            _ = v33
                        }
                    }
                    fmt.Println(v13, v5)
                    var v36 bool = v5
                    _ = fi29
                    _ = v36
                }
                for fi37 := 0; fi37 < 7; fi37++ {
                    v1 = (v1 + 42)
                    _ = fi37
                }
            }
        default:
            v38 := []int{}
            _ = v38
        }
    case 0:
        v23 = (v1 ^ 127)
    default:
        v39 := []int{(v3 | v3)}
        _ = v39
    }
    if true {
        for fi40 := 0; fi40 < 20; fi40++ {
            delete(v6, ("" + ""))
            _ = fi40
        }
    }
    switch (v3 + v3) {
    case 127:
        v41 := GenS1{F0: (0.0 <= v13), F1: v23}
        _ = v41
    case 32768:
        v42 := GenS3{F0: v2, F1: 65535}
        _ = v42
    case 1:
        switch (("a" + "foo bar") + "a") {
        case "foo bar":
            v43 := []int{v3, (v3 << ((uint(v12)) & 63)), v3}
            _ = v43
        }
    default:
        v44 := GenS4{F0: 0, F1: (!v5)}
        _ = v44
    }
    for fi45 := 0; fi45 < 4; fi45++ {
        fmt.Println(v12, v13, v1)
        v5 = (!((v23 <= (v1 | v1)) && (42 >= 255)))
        v46 := []int{}
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
