package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }
func genFn0() uint64 {
    var v4 int64 = -1
    v4 -= v4
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 int8 = 42
    var v2 uint64 = 0
    var v3 int32 = 1
    var v5 uint8 = 0
    var v6 int = -874
    var v7 int16 = 718
    if true {
        v8 := GenS1{F0: (v2 % (((v2 % (((v2 * (v2 >> ((881) & 63)))) | 1))) | 1)), F1: false, F2: false}
        v9 := GenS1{F0: v2, F1: true, F2: true}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(v7, v1)
        _ = fi10
    }
    var v11 int32 = 351
    v12 := GenS1{F0: 0, F1: (false && false), F2: (921 >= 848)}
    fmt.Println(v11)
    var v13 int8 = v1
    for fi14 := 0; fi14 < 10; fi14++ {
        switch 9223372036854775807 {
        case -888:
            var v15 bool = (1e10 == 1e10)
            _ = v15
        case 42:
            if true {
                if (((1 < 2147483648) && (v5 <= v5)) && true) {
                    if (3.14159 > -2.25) {
                        v2 = (v2 - v2)
                        var v16 uint32 = 451
                        var v17 float64 = float64(-2.25)
                        _ = v16
                        _ = v17
                    }
                    if true {
                        switch "" {
                        case "foo bar":
                            for fi18 := 0; fi18 < 17; fi18++ {
                                if (((v7 & (v7 - v7)) == v7) || (!true)) {
                                    var v19 int = (fi14 << ((0) & 63))
                                    fmt.Println(v1)
                                    _ = v19
                                } else {
                                    var v20 bool = false
                                    _ = v20
                                }
                                _ = fi18
                            }
                        default:
                            v21 := []int{int(v5)}
                            _ = v21
                        }
                    }
                }
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 bool = true
                    v24 := GenS1{F0: 18446744073709551615, F1: ((v2 % (((^(v2))) | 1)) >= v2), F2: true}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v25 := []int{100, 1, v6}
                _ = v25
            }
        }
        v11 += v11
        _ = fi14
    }
    if (!(!false)) {
        v26 := []int{}
        _ = v26
    }
    if false {
        var v27 uint8 = 127
        var v28 int16 = int16(v1)
        _ = v27
        _ = v28
    } else {
        if (true && (-128 <= (v1 + 0))) {
            v7 -= 32767
        } else {
            var v29 uint64 = (v2 | genFn0())
            var v30 float64 = (((1e10 / 1e10) + 1.5) - 3.14159)
            v30 = ((1.5 + 0.0) / 3.14159)
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 17; fi31++ {
            if ((65535 == 42) && true) {
                fmt.Println(v2, v3)
                v32 := []int{fi31, 983}
                if (!(true && true)) {
                    if len(v32) >= 2 {
                        v32 = v32[1:len(v32)]
                    }
                    var v33 int8 = -21
                    _ = v33
                } else {
                    switch ("hello" + "") {
                    case "\n\t":
                        v11 -= (v3 % (((v3 * v11)) | 1))
                    default:
                        var v34 int = int(v11)
                        _ = v34
                    }
                    v11 = v3
                    var v35 uint = 18446744073709551615
                    _ = v35
                }
                _ = v32
            }
            _ = fi31
        }
        v1 -= (v13 ^ (v13 & (v13 | v1)))
    }
    v2 = 18446744073709551615
    v2 = v2
    var v36 int16 = v7
    var v37 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v5
    _ = v6
    _ = v7
}
