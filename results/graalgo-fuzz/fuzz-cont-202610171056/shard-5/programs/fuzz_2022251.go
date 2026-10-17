package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int8) int16 {
    switch ("foo bar" + "hello") {
    case "":
        if ((!false) && false) {
            var v5 float32 = -2.25
            _ = v5
        }
    }
    p0 += p0
    _ = p0
    return int16(p0)
}

func main() {
    var v1 int = -2147483648
    v2 := []int{(v1 / ((v1) | 1)), v1, (v1 | (^(v1)))}
    v3 := make(map[int]uint32)
    v4 := []int{(v1 * v1)}
    var v6 uint32 = 220
    switch "" {
    case "":
        v1 += -906
    case "hello":
        var v7 int16 = 32767
        _ = v7
    case "a":
        var v8 uint8 = 132
        _ = v8
    default:
        if (42 > genFn0(42)) {
            for fi9 := 0; fi9 < 17; fi9++ {
                v1 = (fi9 * v1)
                _ = fi9
            }
            v6 = v6
        }
    }
    var v10 int16 = 42
    v1 -= (v1 & -9223372036854775808)
    v10 = (v10 | (v10 >> ((42) & 63)))
    for fi11 := 0; fi11 < 15; fi11++ {
        fmt.Println(v1, v10)
        _ = fi11
    }
    if false {
        fmt.Println(v1, v6, v1)
    } else {
        var v12 int64 = 870
        var v13 int16 = genFn0(66)
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 string = "\n\t"
            switch (v15 + (v15 + "")) {
            case "\n\t":
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 int8 = 0
                    v6 -= 165
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi14
            _ = v15
        }
        _ = v12
        _ = v13
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        v6 = v6
        _ = fi18
    }
    fmt.Println(v10, v6, v6)
    for fi19 := 0; fi19 < 19; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            v1 = 32767
            var v21 uint16 = 32767
            v21 -= (v21 - v21)
            fmt.Println(v1, v6)
            var v22 float64 = float64((((3.14159 * -2.25) - -2.25) - float32(1.5)))
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 int16 = v10
        _ = fi19
        _ = v23
    }
    v6 += v6
    var v24 float32 = 1e10
    v25 := GenS1{F0: -1}
    v4 = append(v4, (v1 - (v1 * (v1 + (^(v1))))))
    if (((v24 / (((v24 - (v24 / (0.1 + v24))) * (((v24 / 1.0) / 0.1) + (-2.25 * v24))) + v24)) * 3.14159) < float32(3.14159)) {
        v1 = (v1 * v1)
    }
    if false {
        for fi26 := 0; fi26 < 7; fi26++ {
            if true {
                v24 = (v24 - 1e10)
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                v28 := GenS1{F0: 1}
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        fmt.Println(v6, v1)
        var v29 bool = (-9223372036854775808 > -9223372036854775808)
        _ = v29
    } else {
        v3[(v1 << ((42) & 63))] = 2
        var v30 float32 = v24
        for fi31 := 0; fi31 < 15; fi31++ {
            delete(v3, (fi31 + (fi31 & fi31)))
            for fi32 := 0; fi32 < 7; fi32++ {
                v10 = v10
                v33 := GenS1{F0: 1}
                v34 := GenS1{F0: 0}
                if true {
                    v1 = (fi32 | (fi32 % ((1) | 1)))
                    if false {
                        var v35 bool = false
                        v3[-967] = v6
                        v36 := GenS1{F0: -1}
                        _ = v35
                        _ = v36
                    } else {
                        if (65535 > 4294967295) {
                            v1 = fi32
                        } else {
                            v37 := v3[32767]
                            _ = v37
                        }
                    }
                } else {
                    var v38 uint8 = 127
                    _ = v38
                }
                _ = fi32
                _ = v33
                _ = v34
            }
            fmt.Println(v24, v1)
            _ = fi31
        }
        _ = v30
    }
    fmt.Println(v6, v10, v6)
    v24 -= v24
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
}
