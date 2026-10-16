package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
type GenS2 struct { F0 bool }

func main() {
    var v1 int = 4294967295
    var v2 uint = 4294967295
    var v3 float64 = (-2.25 - float64(((0.0 * 0.0) * -2.25)))
    var v4 uint32 = 720
    var v5 uint = v2
    for fi6 := 0; fi6 < 18; fi6++ {
        v3 = float64(-2.25)
        if false {
            if (true && ((v1 % ((fi6) | 1)) != (v1 ^ fi6))) {
                if false {
                    v1 -= (v1 << ((v5) & 63))
                } else {
                    for fi7 := 0; fi7 < 5; fi7++ {
                        v8 := []int{}
                        v2 = v5
                        _ = fi7
                        _ = v8
                    }
                    var v9 uint64 = 4294967295
                    var v10 uint8 = 0
                    _ = v10
                    _ = v9
                }
            } else {
                var v11 uint16 = 11
                v12 := []int{(fi6 | fi6), fi6, (v1 | 128)}
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := []int{fi6}
                if (v4 > (v4 >> (((v2 - (v5 - (^(v5))))) & 63))) {
                    fmt.Println(v2, v2)
                } else {
                    fmt.Println(fi6)
                }
                v15 := GenS1{F0: -128, F1: (fi6 - v1), F2: (true || false)}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        _ = fi6
    }
    if (42 == -1) {
        v4 += 2147483648
    } else {
        v2 = v5
        if true {
            var v16 int16 = -34
            v17 := GenS1{F0: 18, F1: 1, F2: true}
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 16; fi18++ {
                fmt.Println(v2, v4, v1)
                _ = fi18
            }
        }
    }
    fmt.Println(v1, v1)
    v4 = (v4 << ((v2) & 63))
    if (!(true && true)) {
        v2 = v5
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            switch "" {
            case "a":
                v3 = float64((1.5 / -2.25))
                var v21 bool = (42 >= uint16(v5))
                var v22 float64 = 0.0
                _ = v21
                _ = v22
            case "":
                v23 := []int{fi19, 127}
                _ = v23
            case "foo bar":
                fmt.Println(v3, v4, fi20)
            default:
                switch fi20 {
                case -582:
                    fmt.Println(v4, v1)
                case 127:
                    var v24 string = ("" + "foo bar")
                    _ = v24
                case 0:
                    fmt.Println(v3, v1, v2)
                }
            }
            var v25 int8 = 127
            v26 := GenS1{F0: (^(v25)), F1: fi19, F2: false}
            v5 += v5
            _ = fi20
            _ = v25
            _ = v26
        }
        var v27 int16 = 42
        var v28 uint8 = 1
        _ = fi19
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        if false {
            v30 := []int{v1}
            _ = v30
        } else {
            for fi31 := 0; fi31 < 13; fi31++ {
                v32 := []int{(fi31 & fi31), (v1 * (v1 << ((778) & 63)))}
                fmt.Println(v4, v2)
                if (0.0 == 1e10) {
                    _ = len(v32)
                    v33 := GenS2{F0: (false || (!((true && (fi29 == fi31)) && false)))}
                    v1 -= fi29
                    _ = v33
                }
                _ = fi31
                _ = v32
            }
        }
        v34 := []int{42, v1}
        for fi35 := 0; fi35 < 9; fi35++ {
            v36 := GenS1{F0: -1, F1: fi29, F2: false}
            _ = fi35
            _ = v36
        }
        _ = fi29
        _ = v34
    }
    v37 := []int{(v1 + v1), 42, (-(v1))}
    v4 += 2147483648
    v2 = v5
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
