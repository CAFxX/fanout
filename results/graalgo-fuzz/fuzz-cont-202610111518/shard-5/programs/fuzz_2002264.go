package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }

func main() {
    var v1 string = (((("hello" + "a") + "a") + "") + "\n\t")
    var v2 int = -341
    var v3 int32 = -2147483648
    v4 := make(map[uint32]uint32)
    var v5 float64 = (1.5 * 0.0)
    if false {
        v6 := GenS1{F0: 172, F1: "a"}
        switch "" {
        case "foo bar":
            var v7 bool = true
            _ = v7
        case "hello":
            switch ((v1 + v1) + "a") {
            case "a":
                v3 -= v3
            case "hello":
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 int64 = 533
                    var v10 float64 = v5
                    var v11 int8 = 0
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
            }
        case "\n\t":
            var v12 int32 = -1
            _ = v12
        default:
            v13 := GenS1{F0: 127, F1: ("foo bar" + v1)}
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int16 = -111
            _ = len(v4)
            _ = fi14
            _ = v15
        }
        _ = v6
    }
    var v16 uint32 = 2147483647
    v3 -= (v3 * v3)
    v17 := GenS1{F0: 127, F1: "hello"}
    if false {
        for fi18 := 0; fi18 < 7; fi18++ {
            v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
            v4[(^(v16))] = v16
            _ = fi18
        }
        v5 += ((1.5 * ((1e10 / v5) * 1.5)) - v5)
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        delete(v4, (v16 << ((407) & 63)))
        var v20 uint = 255
        var v21 uint32 = (v16 * (v16 * (v16 << ((v20) & 63))))
        _ = fi19
        _ = v20
        _ = v21
    }
    v5 -= ((3.14159 - v5) - v5)
    switch (v1 + (("foo bar" + ("" + v1)) + ("foo bar" + ""))) {
    case "a":
        v1 = (v1 + v1)
    default:
        switch v1 {
        case "\n\t":
            v22 := []int{127, 0, (^(v2))}
            _ = v22
        case "hello":
            if false {
                fmt.Println(v3, v1)
                v2 -= v2
                v2 -= (^(v2))
            }
        case "a":
            var v23 uint = 342
            _ = v23
        default:
            if (v2 <= (-(v2))) {
                if false {
                    v4[v16] = v16
                }
            } else {
                v4[(v16 >> ((2) & 63))] = (v16 ^ v16)
                for fi24 := 0; fi24 < 11; fi24++ {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        switch 2147483647 {
                        case 32768:
                            var v26 int64 = -937
                            _ = v26
                        case 128:
                            fmt.Println(v16, v16)
                            var v27 uint32 = v16
                            _ = v27
                        default:
                            v3 = v3
                        }
                        var v28 int16 = 0
                        _ = fi25
                        _ = v28
                    }
                    fmt.Println(v1, fi24)
                    _ = fi24
                }
                v29 := []int{127}
                _ = v29
            }
        }
    }
    switch ((v1 + v1) + v1) {
    case "hello":
        for fi30 := 0; fi30 < 19; fi30++ {
            if (332 == 531) {
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 bool = false
                    var v33 uint16 = 0
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                var v34 uint8 = 225
                _ = v34
            } else {
                v35 := []int{v2, 2147483647, 42}
                _ = v35
            }
            for fi36 := 0; fi36 < 10; fi36++ {
                v3 = (v3 % ((v3) | 1))
                var v37 int = 127
                _ = fi36
                _ = v37
            }
            _ = fi30
        }
    case "foo bar":
        v3 += v3
    default:
        v16 = 227
    }
    fmt.Println(v2, v16, v5)
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 uint = 2
        for fi40 := 0; fi40 < 20; fi40++ {
            v41 := GenS1{F0: 175, F1: v1}
            fmt.Println(v5, fi38)
            var v42 uint8 = 95
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = fi38
        _ = v39
    }
    var v43 string = v1
    v16 = 0
    var v44 float32 = ((0.1 / 3.14159) * (1e10 / (float32((v5 / (1e10 * v5))) * 0.1)))
    if ((18 <= 127) && true) {
        var v45 float32 = float32((v5 * v5))
        var v46 uint8 = 2
        for fi47 := 0; fi47 < 9; fi47++ {
            fmt.Println(v2, v43, v2)
            if false {
                var v48 string = (v1 + v1)
                v49 := []int{v2, (v2 << ((1) & 63)), 536}
                var v50 float32 = v45
                _ = v48
                _ = v49
                _ = v50
            }
            var v51 uint16 = 65535
            _ = fi47
            _ = v51
        }
        _ = v45
        _ = v46
    } else {
        v52 := GenS1{F0: 127, F1: v43}
        _ = v52
    }
    v53 := GenS1{F0: 128, F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v53
}
