package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
func genFn0(p0 int16, p1 int16, p2 int64) string {
    p2 += (-(p2))
    fmt.Println(p1)
    if true {
        v7 := []int{}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 20; fi8++ {
            p0 = (^(p1))
            _ = fi8
        }
    }
    p1 -= p0
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}
func genFn1(p0 uint8) uint64 {
    v9 := []int{}
    fmt.Println(p0)
    switch 2147483647 {
    case -2147483648:
        if (true && false) {
            var v10 int8 = 0
            _ = v10
        } else {
            p0 += 255
            var v11 float32 = (3.14159 + float32((3.14159 + (0.0 * (1.5 * 0.0)))))
            var v12 bool = (-128 >= -23)
            _ = v11
            _ = v12
        }
    default:
        var v13 int32 = 1
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 string = ""
        v15 = v15
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (0.1 + ((float64(-2.25) / 1.0) - (float64((1.5 * (1e10 - 1e10))) - 3.14159)))
    var v3 int16 = -812
    var v4 int32 = -2147483648
    var v5 int = 127
    var v6 string = "foo bar"
    if (1 < 1) {
        v16 := GenS1{F0: genFn1(42), F1: 128}
        _ = v16
    }
    switch v6 {
    case "foo bar":
        v6 = "hello"
    default:
        var v17 uint = 671
        _ = v17
    }
    v3 = v3
    _ = len(v1)
    var v18 bool = false
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 uint = 623
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    fmt.Println(v18, v6)
    v22 := GenS1{F0: 1, F1: (v5 | v5)}
    if (((v3 | (v3 << ((65535) & 63))) == v3) && v18) {
        if v18 {
            switch v6 {
            case "hello":
                var v23 int32 = (v4 >> ((664) & 63))
                _ = v23
            }
            if (548 != 1) {
                if (-58 >= 9223372036854775807) {
                    switch ((v6 + (v6 + v6)) + ("hello" + ((v6 + (v6 + ("a" + "foo bar"))) + v6))) {
                    case "hello":
                        var v24 int8 = -15
                        _ = v24
                    }
                    var v25 string = (v6 + "a")
                    _ = v25
                }
            }
            delete(v1, (^(v5)))
        } else {
            if v18 {
                if v18 {
                    switch v6 {
                    case "\n\t":
                        for fi26 := 0; fi26 < 10; fi26++ {
                            var v27 float64 = 3.14159
                            _ = fi26
                            _ = v27
                        }
                    case "hello":
                        v5 = v5
                    case "a":
                        for fi28 := 0; fi28 < 11; fi28++ {
                            for fi29 := 0; fi29 < 17; fi29++ {
                                _ = len(v1)
                                v30 := []int{(fi28 + (fi29 % ((-532) | 1))), (^(fi29))}
                                _ = fi29
                                _ = v30
                            }
                            _ = fi28
                        }
                    }
                    var v31 uint64 = 1
                    _ = v31
                } else {
                    v32 := GenS1{F0: 1, F1: 0}
                    v33 := []int{(^(v5)), v5, (v5 >> ((4294967295) & 63))}
                    var v34 uint = 776
                    _ = v32
                    _ = v33
                    _ = v34
                }
            }
            if v18 {
                var v35 uint16 = 255
                fmt.Println(v6)
                var v36 float32 = float32(v2)
                _ = v35
                _ = v36
            } else {
                v37 := []int{v5, (v5 ^ v5)}
                v2 = (((v2 * ((0.0 * v2) / float64((3.14159 * 3.14159)))) / v2) + (((1.5 + 0.1) * v2) + v2))
                _ = v37
            }
            switch v5 {
            case 219:
                var v38 int32 = 1
                _ = v38
            case 42:
                switch 553 {
                case -9223372036854775808:
                    fmt.Println(v3, v18)
                }
            case 1:
                v3 = v3
            default:
                v18 = v18
            }
        }
    } else {
        _ = len(v1)
    }
    var v39 uint16 = 1
    if v18 {
        v1[(v5 + v5)] = (((("" + v6) + ((v6 + (v6 + ("\n\t" + "hello"))) + "a")) + "a") + ("a" + v6))
        var v40 int = (v5 % ((-2147483648) | 1))
        _ = v40
    } else {
        switch (v5 - v5) {
        case -1:
            v41 := v1[(v5 ^ v5)]
            _ = v41
        default:
            v42 := []int{(v5 % ((int(v3)) | 1)), v5, v5, 398}
            _ = v42
        }
        v5 += v5
        v43 := []int{}
        _ = v43
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        v45 := []int{v5}
        _ = fi44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
