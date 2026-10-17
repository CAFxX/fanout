package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() string {
    var v6 float64 = (1e10 * 0.1)
    var v7 int = 0
    if true {
        var v8 int = (v7 << ((42) & 63))
        if (!(-474 != -189)) {
            var v9 int = -848
            v10 := []int{v8, (^(v7)), 653, v8}
            _ = v10
            _ = v9
        } else {
            v11 := []int{(v7 >> ((97) & 63)), (^(v8)), (v8 << ((65535) & 63)), (v7 ^ v8)}
            for fi12 := 0; fi12 < 13; fi12++ {
                v11 = append(v11, 32767)
                for fi13 := 0; fi13 < 10; fi13++ {
                    switch v7 {
                    case 167:
                        if len(v11) > 0 {
                            v11[0] = fi13
                        }
                    }
                    _ = fi13
                }
                var v14 bool = (127 < 127)
                _ = fi12
                _ = v14
            }
            _ = v11
        }
        _ = v8
    }
    var v15 int32 = -28
    v16 := GenS1{F0: 46}
    _ = v15
    _ = v16
    _ = v6
    _ = v7
    return ""
}

func main() {
    var v1 uint16 = 32767
    var v2 float32 = (0.0 - (0.0 + 0.1))
    var v3 uint8 = 218
    var v4 bool = (32767 < -1)
    v5 := make(map[uint32]bool)
    v17 := v5[126]
    if v4 {
        var v18 int32 = 2147483647
        var v19 uint64 = 0
        _ = v18
        _ = v19
    } else {
        v2 = v2
        switch ("" + (("" + "foo bar") + "")) {
        case "hello":
            if (float64(0.0) != 3.14159) {
                for fi20 := 0; fi20 < 9; fi20++ {
                    v17 = true
                    _ = fi20
                }
                if v4 {
                    v21 := []int{-2147483648, 833}
                    _ = v21
                } else {
                    var v22 float64 = 0.1
                    v17 = (-2.25 != ((v22 - ((((float64(v2) - float64((3.14159 / v2))) * 0.0) + (-2.25 * v22)) * -2.25)) / (1e10 * v22)))
                    var v23 int16 = -32768
                    _ = v22
                    _ = v23
                }
                fmt.Println(v2)
            }
        default:
            v17 = true
        }
        var v24 string = "a"
        _ = v24
    }
    fmt.Println(v4, v17)
    var v25 uint32 = 0
    v1 = v1
    if true {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 int8 = 0
            v17 = (!v4)
            v5[(v25 % ((2147483648) | 1))] = v17
            _ = fi26
            _ = v27
        }
    } else {
        v28 := GenS1{F0: 42}
        if ((v25 * uint32(v3)) <= v25) {
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 float32 = v2
                _ = fi29
                _ = v30
            }
        }
        _ = v28
    }
    if false {
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 bool = (v17 || true)
            var v33 float64 = ((0.1 * 0.1) * 1e10)
            if v32 {
                v34 := GenS1{F0: -122}
                v35 := GenS1{F0: 32}
                switch "" {
                case "\n\t":
                    delete(v5, (v25 | (^(v25))))
                }
                _ = v34
                _ = v35
            }
            _ = fi31
            _ = v32
            _ = v33
        }
        if v17 {
            if (!v4) {
                v4 = ((v3 | uint8(v1)) <= 127)
                v3 -= (v3 * (^(v3)))
            }
        }
        var v36 float64 = -2.25
        _ = v36
    } else {
        v2 = (float32(((0.1 + (0.0 / -2.25)) * float64(v2))) - (v2 - 0.0))
        v3 = (^(v3))
        if (((((!(!(2147483647 == 0))) || ("" == "\n\t")) || (1 < 0)) || v17) || (!(0 >= 337))) {
            v37 := []int{9223372036854775807, 128, -191, -283}
            _ = v37
        }
    }
    v25 = 2147483648
    var v38 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
