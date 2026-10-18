package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int8; F1 int32; F2 int32 }
type GenS4 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    var v3 uint32 = 4294967295
    v4 := make(map[int]bool)
    var v5 float64 = ((-2.25 + 0.1) + -2.25)
    var v6 float64 = v5
    v5 += 0.1
    if ((!false) || (v3 != v3)) {
        v7 := GenS1{F0: 1}
        fmt.Println(v5, v6)
        v8 := []int{-2147483648, -9223372036854775808, -9223372036854775808, -491}
        _ = v7
        _ = v8
    }
    var v9 float64 = 1.5
    for fi10 := 0; fi10 < 16; fi10++ {
        v9 -= (v9 + (1e10 + v9))
        _ = fi10
    }
    if true {
        v9 = v5
        if true {
            v11 := GenS2{F0: -341, F1: 1, F2: uint64(v3)}
            switch "a" {
            case "hello":
                if (!true) {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        if (0 >= 0) {
                            v13 := []int{(fi12 & (fi12 - (fi12 / (((^(fi12))) | 1))))}
                            _ = v13
                        } else {
                            var v14 uint = 255
                            v15 := GenS3{F0: 42, F1: 0, F2: -565}
                            _ = v14
                            _ = v15
                        }
                        var v16 int = fi12
                        var v17 string = "hello"
                        _ = fi12
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 14; fi18++ {
                        delete(v2, v3)
                        _ = fi18
                    }
                } else {
                    v19 := GenS2{F0: -285, F1: 0, F2: 462}
                    _ = v19
                }
            case "":
                for fi20 := 0; fi20 < 2; fi20++ {
                    v6 += v5
                    _ = fi20
                }
            default:
                for fi21 := 0; fi21 < 3; fi21++ {
                    v22 := GenS3{F0: -128, F1: 42, F2: 2147483647}
                    _ = fi21
                    _ = v22
                }
            }
            v23 := []int{0, 0, 32767, -9223372036854775808}
            _ = v11
            _ = v23
        } else {
            var v24 uint32 = v3
            fmt.Println(v6)
            _ = v24
        }
    }
    var v25 uint32 = (v3 << ((654) & 63))
    for fi26 := 0; fi26 < 2; fi26++ {
        if (!true) {
            switch (("foo bar" + "") + "") {
            case "\n\t":
                var v27 uint16 = 50
                _ = v27
            case "hello":
                fmt.Println(v6, v3, v25)
            case "a":
                v28 := GenS1{F0: 2}
                _ = v28
            default:
                v25 += v3
            }
            if (v6 > ((((float64(-2.25) / 1.0) + (v5 * (v6 + (v5 * -2.25)))) - v5) * float64(-2.25))) {
                v1[(^(v25))] = 2
                fmt.Println(v25)
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 uint = 515
                if true {
                    var v31 int32 = -2147483648
                    v32 := []int{(fi29 ^ (-(fi26))), fi26, -14, fi29}
                    v3 = v25
                    _ = v31
                    _ = v32
                } else {
                    fmt.Println(v25)
                }
                _ = fi29
                _ = v30
            }
        } else {
            var v33 int64 = -9223372036854775808
            _ = v33
        }
        _ = fi26
    }
    fmt.Println(v3, v6)
    v5 -= v6
    var v34 int8 = -16
    fmt.Println(v6, v6, v6)
    if true {
        v25 -= v25
        _ = len(v4)
    }
    fmt.Println(v6, v9, v6)
    if ((423 <= 1) && false) {
        switch 127 {
        case 32768:
            if true {
                var v35 int32 = 574
                var v36 uint64 = 42
                _ = v35
                _ = v36
            } else {
                for fi37 := 0; fi37 < 10; fi37++ {
                    fmt.Println(v5, v25)
                    v38 := GenS4{F0: (fi37 & -2147483648), F1: 2}
                    var v39 uint = 607
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                for fi40 := 0; fi40 < 4; fi40++ {
                    v41 := []int{}
                    for fi42 := 0; fi42 < 3; fi42++ {
                        _ = len(v2)
                        _ = fi42
                    }
                    _ = fi40
                    _ = v41
                }
            }
        case 296:
            var v43 float32 = 0.0
            _ = v43
        default:
            v9 -= 1e10
        }
    } else {
        fmt.Println(v3, v5, v25)
    }
    var v44 uint64 = 0
    v34 = v34
    fmt.Println(v9)
    var v45 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v9
}
