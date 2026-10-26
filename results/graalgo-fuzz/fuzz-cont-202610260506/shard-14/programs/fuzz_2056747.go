package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int8, p1 uint8) string {
    p0 -= (p0 & 100)
    p1 = 127
    var v4 string = ("\n\t" + (("foo bar" + "hello") + "\n\t"))
    p1 = p1
    var v5 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "foo bar"
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]bool)
    var v3 float64 = 1.5
    v1 = 1
    if (0 == 1) {
        v3 += (((v3 - 0.0) - 0.0) - ((v3 - (v3 - v3)) / 1.0))
        v3 = -2.25
        if false {
            v6 := GenS1{F0: -9223372036854775808, F1: v1, F2: 42}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 float64 = v3
                if false {
                    var v9 bool = false
                    v3 = ((v3 * ((v3 + float64(3.14159)) / -2.25)) / 1.0)
                    _ = v9
                } else {
                    if (65535 <= 2) {
                        var v10 float32 = (-2.25 * (float32(v8) / -2.25))
                        for fi11 := 0; fi11 < 9; fi11++ {
                            var v12 int32 = 0
                            v13 := GenS2{F0: 127}
                            var v14 int64 = -909
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        v10 += (((((float32(v3) * 1e10) - v10) - v10) + v10) / float32(-2.25))
                        _ = v10
                    } else {
                        fmt.Println(v3, v3)
                        switch "a" {
                        case "foo bar":
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(fi7, v3, v8)
                            case "a":
                                switch "foo bar" {
                                case "foo bar":
                                    v3 = (((v3 * (((1e10 / v8) * (-2.25 - float64(((3.14159 * 1e10) / -2.25)))) + (v3 - 1.5))) * (((float64(float32(((-2.25 / float64((3.14159 + (0.0 / 0.1)))) + 1.5))) * ((0.0 + 1e10) * v8)) + float64(float32((float64((((1.5 * 1.5) + ((0.0 - ((0.0 - -2.25) / 1.5)) + 3.14159)) + float32(v3))) + -2.25)))) / (3.14159 / float64((0.1 * 3.14159))))) * v8)
                                default:
                                    if false {
                                        v3 -= float64(float32(float64((float32((v8 / v8)) + 3.14159))))
                                    } else {
                                        var v15 int8 = 42
                                        delete(v2, 350)
                                        v3 = (v8 - (v8 * (v3 * -2.25)))
                                        _ = v15
                                    }
                                }
                            default:
                                v8 -= v8
                            }
                        }
                        v16 := []int{(fi7 * fi7), -9223372036854775808}
                        _ = v16
                    }
                    v17 := GenS2{F0: 127}
                    v8 = v3
                    _ = v17
                }
                v8 = 1.5
                _ = fi7
                _ = v8
            }
        }
    } else {
        v1 -= 42
    }
    v18 := GenS1{F0: -9223372036854775808, F1: 2, F2: 43}
    for fi19 := 0; fi19 < 19; fi19++ {
        fmt.Println(v1, v3, fi19)
        for fi20 := 0; fi20 < 17; fi20++ {
            fmt.Println(fi19)
            switch "a" {
            case "foo bar":
                v2[65535] = true
            default:
                var v21 uint16 = 32767
                _ = v21
            }
            _ = fi20
        }
        _ = fi19
    }
    var v22 int32 = 2147483647
    fmt.Println(v1, v1, v22)
    v23 := GenS2{F0: -128}
    v1 = (^(v1))
    v22 -= v22
    v24 := GenS2{F0: 0}
    var v25 bool = (127 < -128)
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        var v26 int8 = 100
        _ = v26
    case "":
        fmt.Println(v25, v22)
    case "a":
        var v27 float32 = (0.0 - 0.1)
        _ = v27
    }
    fmt.Println(v22, v25, v3)
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 uint32 = 0
        var v30 uint = 2
        _ = fi28
        _ = v29
        _ = v30
    }
    if v25 {
        var v31 uint32 = 834
        _ = v31
    } else {
        v32 := GenS2{F0: 127}
        _ = v32
    }
    v33 := GenS2{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
}
