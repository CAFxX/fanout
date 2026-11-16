package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int32; F2 uint }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int32; F1 bool; F2 uint64 }
func genFn0() int64 {
    var v7 float64 = 1.5
    v8 := GenS1{F0: 128, F1: "hello"}
    var v9 float64 = ((v7 - v7) * (v7 + (v7 + v7)))
    _ = v7
    _ = v8
    _ = v9
    return 610
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 128
    v3 := make(map[uint32]uint32)
    var v4 int = 2147483647
    var v5 int32 = -2147483648
    v6 := make(map[int]bool)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int64 = -9223372036854775808
        _ = fi10
        _ = v11
    }
    fmt.Println(v2, v4, v2)
    var v12 uint8 = 0
    var v13 bool = (false && (false || true))
    var v14 float32 = 1e10
    v2 -= 1
    for fi15 := 0; fi15 < 2; fi15++ {
        v2 -= v12
        for fi16 := 0; fi16 < 13; fi16++ {
            v5 = v5
            v17 := []int{2147483647, fi16}
            fmt.Println(fi15, v2)
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := GenS2{F0: ("\n\t" + ("" + ("" + ("\n\t" + "hello")))), F1: 42}
        _ = fi18
        _ = v19
    }
    v5 -= -2147483648
    fmt.Println(v4, v12)
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := []int{fi20, fi20, (v4 ^ v4), fi20}
        v22 := GenS1{F0: (fi20 >> ((1) & 63)), F1: "a"}
        _ = fi20
        _ = v21
        _ = v22
    }
    if v13 {
        var v23 uint64 = 622
        switch (^(v4)) {
        case 32767:
            v24 := GenS1{F0: (v4 + v4), F1: "foo bar"}
            _ = v24
        case 301:
            switch v4 {
            case 2147483647:
                v25 := GenS1{F0: v4, F1: (("" + "foo bar") + "\n\t")}
                _ = v25
            case 128:
                switch 0 {
                case 42:
                    fmt.Println(v5, v12)
                case 160:
                    v23 = 877
                }
            case 1:
                if (!(((int64(v23) != -9223372036854775808) && (91 > 127)) && (!(true && (-2147483648 != 0))))) {
                    if v13 {
                        v26 := GenS3{F0: "hello"}
                        v4 = v4
                        v27 := GenS4{F0: (v4 % (((-(v4))) | 1)), F1: (^(v5)), F2: 1}
                        _ = v26
                        _ = v27
                    }
                    fmt.Println(v13, v14, v4)
                    switch ("a" + (("hello" + "hello") + (("foo bar" + "hello") + "a"))) {
                    case "a":
                        switch (("" + "\n\t") + "foo bar") {
                        case "a":
                            _ = len(v1)
                        case "":
                            v14 = v14
                        default:
                            v28 := []int{v4}
                            _ = v28
                        }
                    case "":
                        if v13 {
                            for fi29 := 0; fi29 < 6; fi29++ {
                                v30 := GenS4{F0: (v4 << ((155) & 63)), F1: v5, F2: 42}
                                v31 := GenS5{F0: 9223372036854775807, F1: v23}
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                        }
                    case "foo bar":
                        v32 := GenS6{F0: (true || true), F1: 797}
                        _ = v32
                    }
                }
            }
        }
        var v33 string = "hello"
        _ = v23
        _ = v33
    }
    switch (v4 - -2147483648) {
    case 651:
        if (!false) {
            v34 := GenS3{F0: "a"}
            switch v4 {
            case 32768:
                v35 := GenS6{F0: (1e10 <= 0.1), F1: 42}
                _ = v35
            case 127:
                v36 := GenS7{F0: v5, F1: (-1 != 1), F2: 2}
                _ = v36
            }
            v14 += (((3.14159 - 0.1) * v14) / v14)
            _ = v34
        } else {
            v2 = (v12 * (v2 ^ v12))
        }
    case -1:
        if (2 == 303) {
            v13 = v13
            v12 -= 2
        } else {
            var v37 int32 = 1
            var v38 int32 = -639
            _ = v37
            _ = v38
        }
    case 0:
        v13 = (((-2.25 - ((v14 * 3.14159) + (-2.25 * v14))) - v14) == ((v14 + (v14 / v14)) / v14))
    default:
        v39 := []int{(v4 / ((0) | 1))}
        _ = v39
    }
    v5 = -498
    var v40 float32 = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
