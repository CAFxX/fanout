package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0(p0 int8, p1 int16, p2 uint16) bool {
    v5 := []int{-9223372036854775808, 42, 127}
    if true {
        switch "hello" {
        case "\n\t":
            var v6 bool = false
            _ = v6
        default:
            for fi7 := 0; fi7 < 17; fi7++ {
                switch -9223372036854775808 {
                case 0:
                    var v8 uint = 255
                    _ = v8
                }
                v9 := GenS1{F0: 4294967295, F1: int(p2)}
                _ = fi7
                _ = v9
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return false
}
func genFn1(p0 int8, p1 int, p2 uint64) int16 {
    var v10 uint16 = 641
    if (!((false && (42 == p2)) && false)) {
        var v11 uint8 = 255
        for fi12 := 0; fi12 < 3; fi12++ {
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 float64 = (((1e10 * 1.5) / ((1.5 / -2.25) + (1.5 * (0.0 + float64((3.14159 * float32((0.0 / 1.5)))))))) / 3.14159)
                var v15 string = ("" + ("a" + ("foo bar" + "hello")))
                var v16 float32 = 0.0
                var v17 float64 = 0.0
                p1 = (fi13 ^ (fi12 << ((uint(v10)) & 63)))
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi12
        }
        _ = v11
    } else {
        v18 := []int{}
        p2 -= p2
        var v19 int32 = -1
        _ = v18
        _ = v19
    }
    switch 1 {
    case 9223372036854775807:
        v10 = (v10 >> ((720) & 63))
    case -1:
        var v20 float32 = ((3.14159 * 1e10) + (-2.25 * (0.1 - float32(float64(0.1)))))
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 268
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "a"
    v3 := []int{}
    var v4 uint8 = 1
    if false {
        var v21 uint8 = v4
        if (-1 > 0) {
            var v22 uint32 = 42
            if ((!(-2.25 > float32(3.14159))) && (430 == 0)) {
                fmt.Println(v21, v22)
                if true {
                    var v23 int8 = 100
                    var v24 uint16 = 42
                    switch (((((((v2 + v2) + (v2 + (v2 + v2))) + v2) + v2) + v2) + v2) + ("\n\t" + "a")) {
                    case "hello":
                        for fi25 := 0; fi25 < 19; fi25++ {
                            v22 += 42
                            switch "" {
                            case "a":
                                var v26 uint16 = v24
                                _ = v26
                            }
                            _ = fi25
                        }
                    case "foo bar":
                        v3 = append(v3, 9223372036854775807)
                    default:
                        for fi27 := 0; fi27 < 1; fi27++ {
                            var v28 string = ""
                            var v29 uint16 = uint16(v21)
                            for fi30 := 0; fi30 < 2; fi30++ {
                                v22 -= 2
                                v1[fi27] = v22
                                _ = fi30
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = v23
                    _ = v24
                } else {
                    v31 := GenS1{F0: (v22 + v22), F1: 32767}
                    _ = v31
                }
            } else {
                if false {
                    var v32 float64 = float64(0.0)
                    var v33 string = (("a" + "") + "")
                    _ = v32
                    _ = v33
                }
            }
            _ = v22
        } else {
            switch 0 {
            case 9223372036854775807:
                v21 -= (^(v4))
            default:
                v34 := GenS1{F0: 878, F1: -443}
                _ = v34
            }
            var v35 bool = (!false)
            var v36 uint8 = 0
            _ = v35
            _ = v36
        }
        _ = v21
    }
    v4 = (v4 - (^(v4)))
    var v37 int = 32768
    v37 -= v37
    var v38 string = v2
    v4 = 1
    var v39 uint64 = 243
    var v40 int8 = -128
    var v41 string = (v2 + ("\n\t" + ("a" + "\n\t")))
    for fi42 := 0; fi42 < 1; fi42++ {
        for fi43 := 0; fi43 < 5; fi43++ {
            if len(v3) > 0 {
                v3[0] = fi43
            }
            _ = fi43
        }
        fmt.Println(v37, v40)
        _ = fi42
    }
    var v44 float64 = ((-2.25 / 3.14159) + 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v44
}
