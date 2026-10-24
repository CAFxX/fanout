package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS3 struct { F0 int; F1 uint8; F2 string }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 1
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    if false {
        v4[826] = "a"
        v3 = v3
    } else {
        _ = len(v5)
    }
    fmt.Println(v2, v3)
    if false {
        for fi6 := 0; fi6 < 3; fi6++ {
            v2 -= int8(fi6)
            v2 -= v2
            _ = fi6
        }
    } else {
        var v7 int8 = v2
        if (!(!((42 >= 32767) || false))) {
            for fi8 := 0; fi8 < 14; fi8++ {
                v7 = v2
                if ((65535 > 65535) || (!true)) {
                    v7 = (v2 >> ((0) & 63))
                    v3 -= v3
                } else {
                    var v9 float32 = -2.25
                    switch "" {
                    case "a":
                        v10 := GenS1{F0: 583, F1: v2, F2: 18446744073709551615}
                        _ = v10
                    case "hello":
                        for fi11 := 0; fi11 < 6; fi11++ {
                            fmt.Println(v3, fi11, v3)
                            _ = fi11
                        }
                    case "\n\t":
                        v7 += v7
                    }
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v13 := GenS2{F0: 65535, F1: 834, F2: 1}
                        _ = fi12
                        _ = v13
                    }
                    _ = v9
                }
                _ = fi8
            }
            v14 := GenS2{F0: 65535, F1: 255, F2: (v3 | (v3 * 18446744073709551615))}
            _ = v14
        } else {
            v2 = v2
        }
        fmt.Println(v7, v2)
        _ = v7
    }
    switch -939 {
    case 32767:
        var v15 int8 = (v2 >> ((18446744073709551615) & 63))
        _ = v15
    default:
        delete(v5, "")
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 string = "hello"
        _ = fi16
        _ = v17
    }
    if (false || (false && (-1 < 674))) {
        if false {
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS2{F0: 2147483648, F1: 0, F2: v3}
                v3 += v3
                v20 := v4[90]
                _ = fi18
                _ = v19
                _ = v20
            }
            v21 := []int{1}
            switch 32768 {
            case 42:
                var v22 int8 = (v2 + (v2 & v2))
                _ = v22
            case -2147483648:
                fmt.Println(v3, v3)
            case -1:
                v3 -= (v3 * v3)
            default:
                for fi23 := 0; fi23 < 20; fi23++ {
                    v2 = (-(v2))
                    v24 := GenS1{F0: fi23, F1: -100, F2: 274}
                    _ = fi23
                    _ = v24
                }
            }
            _ = v21
        } else {
            var v25 float64 = 0.0
            v26 := GenS2{F0: 2147483648, F1: 18, F2: v3}
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 int8 = (v2 + (v2 & (v2 * (v2 ^ 0))))
            v29 := GenS3{F0: fi27, F1: 60, F2: "hello"}
            if (!(false || false)) {
                var v30 uint8 = 42
                v31 := GenS3{F0: fi27, F1: 127, F2: ((("\n\t" + ("\n\t" + (((("\n\t" + "foo bar") + (("\n\t" + ("hello" + "hello")) + "a")) + "a") + "\n\t"))) + "\n\t") + (("" + "a") + "hello"))}
                _ = v30
                _ = v31
            } else {
                if ((v28 + v28) == (v28 + v28)) {
                    v3 = v3
                    v3 = (v3 + (v3 / (((v3 % ((v3) | 1))) | 1)))
                    var v32 int = (fi27 << ((65535) & 63))
                    _ = v32
                }
                fmt.Println(v2, v3)
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        v33 := []int{}
        _ = v33
    }
    var v34 int16 = -1
    v34 = -32768
    for fi35 := 0; fi35 < 16; fi35++ {
        if false {
            v3 = v3
            switch fi35 {
            case -2147483648:
                fmt.Println(v34, fi35, fi35)
            case 32767:
                v34 += v34
            default:
                var v36 float32 = 1.5
                _ = v36
            }
        } else {
            v37 := GenS1{F0: fi35, F1: v2, F2: 18446744073709551615}
            var v38 float64 = -2.25
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    fmt.Println(v2, v34)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
