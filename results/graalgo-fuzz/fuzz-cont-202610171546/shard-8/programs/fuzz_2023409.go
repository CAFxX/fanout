package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int32; F2 uint }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int32 }
func genFn0(p0 bool, p1 uint16, p2 bool) int32 {
    p0 = false
    if p0 {
        v7 := []int{-1}
        var v8 uint32 = 2
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 4294967295}
    p1 += (p1 << ((2) & 63))
    v10 := GenS2{F0: 127, F1: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -198
}
func genFn1(p0 int16, p1 string, p2 string) int16 {
    if (p1 <= (p2 + (p2 + p1))) {
        p1 = ("foo bar" + "a")
        var v11 float32 = -2.25
        var v12 int32 = 2147483647
        _ = v11
        _ = v12
    }
    if (18446744073709551615 <= 801) {
        var v13 float64 = (float64(1.5) - 0.1)
        _ = v13
    } else {
        var v14 string = p2
        var v15 uint32 = 42
        if (24 != 42) {
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 bool = true
                v18 := GenS1{F0: 42}
                v17 = v17
                _ = fi16
                _ = v17
                _ = v18
            }
            if true {
                v19 := []int{-76}
                _ = v19
            } else {
                v20 := []int{-40}
                _ = v20
            }
            var v21 float64 = (1.5 * (((0.0 / -2.25) - -2.25) - 1.5))
            _ = v21
        }
        _ = v14
        _ = v15
    }
    fmt.Println(p0, p0)
    var v22 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return (p0 << ((42) & 63))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 18446744073709551615
    var v3 uint64 = v2
    var v4 int = 127
    var v5 uint32 = 676
    var v6 uint8 = 128
    v23 := []int{}
    v6 = 128
    for fi24 := 0; fi24 < 3; fi24++ {
        switch (v4 * 128) {
        case -303:
            v5 = (^(v5))
        default:
            var v25 uint = 65535
            _ = v25
        }
        v4 = fi24
        var v26 string = "hello"
        _ = fi24
        _ = v26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        delete(v1, v4)
        v28 := GenS3{F0: 1}
        var v29 uint8 = v6
        _ = fi27
        _ = v28
        _ = v29
    }
    delete(v1, (v4 ^ (v4 - v4)))
    var v30 int64 = -1
    switch v4 {
    case 0:
        v31 := GenS2{F0: 1, F1: ""}
        _ = v31
    case 2147483647:
        if ((v2 <= v3) && (3.14159 >= 0.0)) {
            var v32 int64 = v30
            for fi33 := 0; fi33 < 10; fi33++ {
                v6 += (v6 >> ((1) & 63))
                if false {
                    v34 := GenS2{F0: 91, F1: "\n\t"}
                    _ = v34
                } else {
                    v32 += v32
                    var v35 uint16 = 0
                    v2 += v2
                    _ = v35
                }
                var v36 uint64 = 42
                _ = fi33
                _ = v36
            }
            fmt.Println(v4, v2)
            _ = v32
        } else {
            var v37 float64 = (float64(float32(0.0)) - 0.0)
            var v38 uint32 = v5
            v5 -= v38
            _ = v37
            _ = v38
        }
    case 707:
        switch "" {
        case "a":
            v39 := GenS4{F0: 42, F1: -800, F2: 99}
            _ = v39
        case "hello":
            fmt.Println(v30)
        default:
            var v40 string = ("hello" + "a")
            _ = v40
        }
    default:
        if (65535 == (v5 >> ((1) & 63))) {
            if ((v4 % ((-730) | 1)) <= (v4 + v4)) {
                var v41 int64 = v30
                if (!false) {
                    for fi42 := 0; fi42 < 3; fi42++ {
                        var v43 int32 = 0
                        v44 := GenS3{F0: (v30 >> ((2) & 63))}
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    v45 := GenS5{F0: int8(v30), F1: 277}
                    var v46 float32 = (1e10 / 1.5)
                    _ = v45
                    _ = v46
                }
                v6 += (v6 & (v6 >> ((255) & 63)))
                _ = v41
            }
        }
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        switch fi47 {
        case -356:
            for fi48 := 0; fi48 < 7; fi48++ {
                if len(v23) > 0 {
                    v23[0] = 32768
                }
                _ = fi48
            }
        case 9223372036854775807:
            if true {
                for fi49 := 0; fi49 < 5; fi49++ {
                    fmt.Println(v6)
                    var v50 float64 = -2.25
                    _ = fi49
                    _ = v50
                }
            }
        default:
            v51 := GenS6{F0: 961, F1: 255}
            _ = v51
        }
        _ = fi47
    }
    fmt.Println(v4, v2, v4)
    v6 = uint8(v5)
    v30 = v30
    switch 127 {
    case -1:
        v5 -= (^(v5))
    case 128:
        var v52 int16 = 42
        _ = v52
    case 952:
        v53 := GenS5{F0: 1, F1: 2}
        _ = v53
    default:
        _ = len(v1)
    }
    var v54 float32 = 3.14159
    fmt.Println(v6, v2)
    v55 := GenS7{F0: 42}
    if false {
        var v56 float64 = (1.5 / (0.0 / (1e10 - (0.0 / (((-2.25 + (0.1 + 1.5)) / 1.0) - ((float64(v54) + (((((1.5 / 1.0) + (float64(-2.25) - 0.0)) / 1.0) * 1e10) / 0.1)) * (-2.25 - 0.1)))))))
        _ = v56
    } else {
        for fi57 := 0; fi57 < 14; fi57++ {
            if false {
                fmt.Println(v30, v3)
            }
            fmt.Println(fi57, v4, v3)
            v2 += (v3 + uint64(v5))
            _ = fi57
        }
        if (!(!((false || ((-2.25 == -2.25) && (true && true))) && (0 < 0)))) {
            var v58 string = (("" + (("" + "foo bar") + "foo bar")) + "")
            fmt.Println(v30, v3, v30)
            v58 = ("\n\t" + (v58 + v58))
            _ = v58
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
