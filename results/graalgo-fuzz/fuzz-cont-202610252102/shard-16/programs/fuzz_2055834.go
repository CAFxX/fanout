package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 bool }
func genFn0(p0 int64) uint8 {
    if true {
        v7 := GenS1{F0: (false && false), F1: 255, F2: (false && true)}
        var v8 string = ("hello" + "a")
        _ = v7
        _ = v8
    } else {
        var v9 bool = true
        _ = v9
    }
    p0 = (p0 & p0)
    _ = p0
    return 0
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    var v6 float64 = 0.0
    fmt.Println(v1, v1)
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := []int{(fi10 - 9223372036854775807), fi10, fi10}
        var v12 int = fi10
        v4[(^(v12))] = (42 == 2)
        _ = fi10
        _ = v11
        _ = v12
    }
    switch "hello" {
    case "foo bar":
        v1 = 308
    case "":
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := []int{}
            fmt.Println(v6, v1)
            _ = fi13
            _ = v14
        }
    case "\n\t":
        switch 2147483647 {
        case 42:
            if (2 <= 952) {
                v1 = v1
            } else {
                var v15 int16 = int16(v1)
                v16 := v2[0]
                _ = v15
                _ = v16
            }
        case 4294967295:
            fmt.Println(v1, v1)
        default:
            var v17 uint64 = 0
            _ = v17
        }
    default:
        v6 += 1e10
    }
    if true {
        v18 := GenS1{F0: (((float32((float64(1e10) - v6)) * 0.1) - (0.1 - 1.5)) < 0.1), F1: 47, F2: (!(!((!true) || (float32(v6) <= float32(float64(3.14159))))))}
        _ = v18
    }
    v1 += (v1 / (((^(v1))) | 1))
    fmt.Println(v6, v1)
    for fi19 := 0; fi19 < 19; fi19++ {
        v1 = (v1 + 725)
        v20 := GenS1{F0: ((74 >= int64(fi19)) || (!true)), F1: genFn0(649), F2: (false && ((!(42 <= 18446744073709551615)) && (!(2 >= 2))))}
        _ = fi19
        _ = v20
    }
    var v21 uint16 = (v1 | (v1 / ((0) | 1)))
    if (105 == 1) {
        v22 := v5[int(v21)]
        _ = v22
    }
    v23 := GenS1{F0: (!(2 <= 0)), F1: 0, F2: (0 > -2147483648)}
    _ = len(v5)
    v24 := []int{}
    fmt.Println(v1, v6)
    v21 += (v1 * (v21 * 65535))
    if (true || ("" <= "\n\t")) {
        switch 47 {
        case -2147483648:
            v21 = 496
        default:
            fmt.Println(v1, v1)
        }
        v5[-302] = 42
        v1 -= 32767
    } else {
        var v25 float32 = (float32(-2.25) - float32(float64((0.1 / (-2.25 / 1.0)))))
        fmt.Println(v25, v6, v25)
        switch "hello" {
        case "":
            if false {
                var v26 uint8 = 0
                switch -402 {
                case 4294967295:
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v28 := GenS1{F0: false, F1: 30, F2: true}
                        for fi29 := 0; fi29 < 13; fi29++ {
                            var v30 int8 = 42
                            if (!true) {
                                var v31 float32 = (float32(float64((v25 + v25))) - 0.0)
                                var v32 uint8 = v26
                                _ = v31
                                _ = v32
                            } else {
                                var v33 bool = true
                                v30 = v30
                                _ = v33
                            }
                            delete(v2, fi27)
                            _ = fi29
                            _ = v30
                        }
                        v6 += (v6 - v6)
                        _ = fi27
                        _ = v28
                    }
                case 794:
                    v34 := GenS1{F0: true, F1: (v26 | v26), F2: ((42 < -2147483648) && true)}
                    _ = v34
                case 414:
                    v25 -= v25
                }
                _ = v26
            } else {
                var v35 uint32 = 65535
                if len(v24) > 0 {
                    v24[0] = 4294967295
                }
                _ = v35
            }
        default:
            switch (((("hello" + "a") + "foo bar") + "foo bar") + ("" + "a")) {
            case "hello":
                var v36 int8 = -1
                _ = v36
            case "foo bar":
                v25 -= v25
            case "a":
                v21 -= v21
            default:
                for fi37 := 0; fi37 < 9; fi37++ {
                    v38 := GenS1{F0: false, F1: 94, F2: false}
                    v25 = v25
                    _ = fi37
                    _ = v38
                }
            }
        }
        _ = v25
    }
    if ((-2147483648 <= 1) && true) {
        fmt.Println(v21, v6, v6)
        var v39 bool = false
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
