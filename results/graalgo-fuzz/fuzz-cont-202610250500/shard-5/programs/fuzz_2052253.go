package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 string; F2 int64 }
type GenS4 struct { F0 int32; F1 uint32; F2 string }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 uint16; F1 string; F2 int }
func genFn0(p0 uint64, p1 int64, p2 int64) int {
    if false {
        v5 := []int{}
        _ = v5
    }
    v6 := []int{9223372036854775807, 32768}
    switch "\n\t" {
    case "hello":
        fmt.Println(p0, p1, p2)
    }
    if false {
        v7 := GenS1{F0: 255, F1: (92 == int32(p0)), F2: 255}
        _ = v7
    } else {
        fmt.Println(p0, p1)
        if (p0 <= p0) {
            p0 = (p0 * 865)
            if len(v6) > 0 {
                v6[0] = 952
            }
        }
        fmt.Println(p0, p2, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 585
}
func genFn1() int32 {
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    switch -9223372036854775808 {
    case 128:
        var v10 float32 = ((1.5 / 1.0) / float32((3.14159 / ((((((float64(-2.25) * 1.5) - 1.5) / 0.1) * 0.1) * float64(1e10)) - 1e10))))
        _ = v10
    case 364:
        var v11 bool = false
        _ = v11
    }
    if ((649 == 2) || true) {
        v12 := GenS2{F0: 42}
        _ = v12
    } else {
        v13 := GenS3{F0: 42, F1: "\n\t", F2: -9223372036854775808}
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 bool = false
            if (0 < 295) {
                v16 := GenS4{F0: -696, F1: 2, F2: "hello"}
                var v17 int64 = 0
                _ = v16
                _ = v17
            }
            v18 := GenS5{F0: 42, F1: 2147483648}
            _ = fi14
            _ = v15
            _ = v18
        }
        if true {
            var v19 uint8 = 35
            _ = v19
        }
        _ = v13
    }
    return -2147483648
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v4 uint = 18446744073709551615
    for fi20 := 0; fi20 < 12; fi20++ {
        v4 = 279
        if true {
            var v21 int16 = -32768
            _ = v21
        } else {
            v22 := []int{(fi20 + genFn0(0, -1, 42))}
            _ = v22
        }
        v23 := []int{(fi20 / ((9223372036854775807) | 1)), (-(fi20)), (^(fi20)), -9223372036854775808}
        _ = fi20
        _ = v23
    }
    _ = len(v3)
    v24 := GenS6{F0: 0, F1: "a", F2: 1}
    fmt.Println(v1)
    fmt.Println(v4)
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 uint = (v4 >> ((v4) & 63))
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := []int{fi25, -2147483648, 183}
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        if true {
            var v30 uint32 = 1
            if true {
                for fi31 := 0; fi31 < 6; fi31++ {
                    v4 += v4
                    if false {
                        var v32 int64 = 364
                        var v33 float32 = -2.25
                        _ = v32
                        _ = v33
                    } else {
                        fmt.Println(v4)
                    }
                    v1 -= v1
                    _ = fi31
                }
                fmt.Println(v4, fi29, v4)
            } else {
                if (!(1e10 < 1.5)) {
                    v34 := []int{genFn0(0, 42, 0), (-(fi29)), 4294967295, fi29}
                    v35 := GenS3{F0: uint8(v30), F1: ((("foo bar" + ("hello" + "")) + "a") + "a"), F2: 1}
                    for fi36 := 0; fi36 < 17; fi36++ {
                        var v37 int = int(v30)
                        v37 = fi36
                        v4 = 18446744073709551615
                        _ = fi36
                        _ = v37
                    }
                    _ = v34
                    _ = v35
                }
            }
            _ = v30
        } else {
            switch "hello" {
            case "foo bar":
                v4 = (v4 + (v4 - v4))
            case "hello":
                if true {
                    v1 -= v1
                } else {
                    if (false || true) {
                        v38 := GenS4{F0: (v1 * v1), F1: 1, F2: "foo bar"}
                        if true {
                            v4 -= (^(v4))
                            v39 := GenS4{F0: v1, F1: 4294967295, F2: "a"}
                            _ = v39
                        }
                        _ = v38
                    } else {
                        var v40 uint32 = 65535
                        v41 := GenS3{F0: 2, F1: "foo bar", F2: -99}
                        _ = v40
                        _ = v41
                    }
                }
            case "a":
                v1 += genFn1()
            default:
                switch (fi29 % (((fi29 % ((-235) | 1))) | 1)) {
                case 32768:
                    var v42 uint = (v4 / ((v4) | 1))
                    _ = v42
                case 127:
                    v4 += (v4 * 4294967295)
                default:
                    var v43 uint16 = 66
                    _ = v43
                }
            }
        }
        if (v1 != v1) {
            var v44 int64 = -735
            v1 = 631
            _ = v44
        }
        _ = fi29
    }
    var v45 int64 = -1
    for fi46 := 0; fi46 < 4; fi46++ {
        for fi47 := 0; fi47 < 15; fi47++ {
            v45 -= (v45 * v45)
            v45 = 1
            var v48 uint32 = 885
            _ = fi47
            _ = v48
        }
        v49 := GenS6{F0: 28, F1: "", F2: fi46}
        _ = fi46
        _ = v49
    }
    fmt.Println(v4, v45, v4)
    if false {
        var v50 float32 = 3.14159
        _ = v50
    }
    fmt.Println(v1, v45)
    switch 32768 {
    case 523:
        if (!(0 == 4294967295)) {
            v4 = v4
        } else {
            var v51 uint64 = 0
            v4 += v4
            _ = v51
        }
    }
    switch "\n\t" {
    case "foo bar":
        if (486 == 127) {
            var v52 uint64 = uint64(v4)
            var v53 float32 = (3.14159 / float32((3.14159 * (3.14159 / 0.1))))
            _ = v52
            _ = v53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v45
}
