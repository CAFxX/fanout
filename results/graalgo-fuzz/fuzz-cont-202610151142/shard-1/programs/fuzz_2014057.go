package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 int64 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint }
func genFn0() int8 {
    if true {
        var v5 uint = 65535
        if (9223372036854775807 == 1) {
            var v6 uint32 = 942
            if false {
                for fi7 := 0; fi7 < 9; fi7++ {
                    v8 := []int{(-(fi7)), 127}
                    if true {
                        v5 = 255
                        v6 -= v6
                        _ = len(v8)
                    } else {
                        v9 := GenS1{F0: (v6 >> ((v5) & 63))}
                        for fi10 := 0; fi10 < 2; fi10++ {
                            fmt.Println(fi7)
                            v6 -= 2147483648
                            _ = fi10
                        }
                        _ = v9
                    }
                    _ = fi7
                    _ = v8
                }
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 uint = (v5 << ((v5) & 63))
                    if true {
                        switch ("hello" + "foo bar") {
                        case "":
                            var v13 uint8 = 0
                            _ = v13
                        }
                        v12 = v5
                    } else {
                        if ((101 >= int16(v6)) || true) {
                            v12 += v5
                            for fi14 := 0; fi14 < 7; fi14++ {
                                v15 := []int{2147483647, fi11, fi14}
                                var v16 float64 = (0.0 + (((-2.25 / 1.5) / float64(0.1)) - 0.1))
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                        }
                    }
                    _ = fi11
                    _ = v12
                }
            }
            switch -761 {
            case 156:
                switch "\n\t" {
                case "":
                    fmt.Println(v5, v6)
                case "a":
                    for fi17 := 0; fi17 < 13; fi17++ {
                        for fi18 := 0; fi18 < 19; fi18++ {
                            v6 -= 60
                            _ = fi18
                        }
                        _ = fi17
                    }
                default:
                    fmt.Println(v6, v6)
                }
            }
            _ = v6
        } else {
            v19 := []int{32767, -91, 956, 32767}
            if true {
                if true {
                    switch "foo bar" {
                    case "":
                        v20 := GenS1{F0: 307}
                        _ = v20
                    case "a":
                        if true {
                            v5 -= (v5 >> ((v5) & 63))
                        }
                    case "\n\t":
                        v5 -= (^(v5))
                    }
                } else {
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    v5 = v5
                }
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 uint64 = 4294967295
                    _ = fi21
                    _ = v22
                }
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
            } else {
                switch -9223372036854775808 {
                case 42:
                    v23 := GenS1{F0: 728}
                    _ = v23
                default:
                    var v24 uint = v5
                    _ = v24
                }
                switch 9223372036854775807 {
                case 42:
                    var v25 uint = v5
                    _ = v25
                case -343:
                    if true {
                        if len(v19) > 0 {
                            v19[0] = -1
                        }
                        v19 = append(v19, 42)
                    }
                }
            }
            _ = v19
        }
        _ = v5
    }
    v26 := GenS1{F0: 1}
    var v27 float64 = ((1e10 / -2.25) / ((0.1 / (float64(1e10) / 0.1)) - float64((float32(1e10) / 1e10))))
    v28 := []int{2147483647, 32768, 42, 42}
    v27 -= float64(-2.25)
    _ = v26
    _ = v27
    _ = v28
    return -51
}
func genFn1(p0 uint32, p1 uint32) int32 {
    p0 = (^(p0))
    p0 += (p1 % ((p1) | 1))
    switch -1 {
    case 127:
        var v29 uint = 661
        _ = v29
    case 32767:
        p0 = (p0 >> ((255) & 63))
    default:
        p1 += p0
    }
    switch -375 {
    case -215:
        v30 := GenS2{F0: 42, F1: 1, F2: 9223372036854775807}
        _ = v30
    default:
        if true {
            v31 := GenS3{F0: (344 == 1), F1: 42}
            _ = v31
        }
    }
    switch 1 {
    case 991:
        var v32 float64 = 0.1
        _ = v32
    case -769:
        v33 := GenS4{F0: -1}
        _ = v33
    default:
        p1 -= 563
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 int = -624
    var v2 float64 = (-2.25 * 1e10)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    v1 = (-(v1))
    var v34 string = ""
    var v35 bool = true
    _ = len(v3)
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 bool = v35
        _ = fi36
        _ = v37
    }
    var v38 float32 = 1e10
    var v39 int16 = 1000
    v40 := v4[65535]
    var v41 float32 = v38
    v38 += (v41 - 1e10)
    switch "hello" {
    case "hello":
        v42 := GenS5{F0: 733, F1: 573}
        _ = v42
    default:
        v39 += v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
