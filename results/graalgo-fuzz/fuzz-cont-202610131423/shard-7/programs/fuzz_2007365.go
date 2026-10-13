package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 bool }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int) int32 {
    v6 := GenS1{F0: 0, F1: 1}
    var v7 float64 = 1.5
    var v8 uint16 = 771
    v7 = v7
    v9 := []int{p0, (p0 & p0), 32767, (-(p0))}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() bool {
    switch ((("\n\t" + "foo bar") + "\n\t") + "hello") {
    case "hello":
        if true {
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := []int{}
                v12 := GenS1{F0: 304, F1: -9223372036854775808}
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 uint8 = 128
            _ = v13
        }
    }
    var v14 uint64 = 576
    v14 += 107
    _ = v14
    return ((2 < 698) || false)
}

func main() {
    var v1 uint = 255
    var v2 int32 = 42
    var v3 uint16 = 103
    var v4 uint = v1
    var v5 bool = true
    var v15 uint8 = 2
    fmt.Println(v2)
    fmt.Println(v3, v4)
    for fi16 := 0; fi16 < 16; fi16++ {
        v15 = v15
        switch -1 {
        case 2147483647:
            switch fi16 {
            case 367:
                var v17 uint32 = 2147483648
                _ = v17
            case -815:
                var v18 int16 = 1
                _ = v18
            }
        }
        _ = fi16
    }
    if (v5 && (!v5)) {
        if false {
            v15 += v15
            v19 := GenS1{F0: (v3 ^ 0), F1: 9223372036854775807}
            _ = v19
        }
    } else {
        v4 = (v4 & v4)
    }
    if false {
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 int16 = 42
            if (-9223372036854775808 <= -773) {
                v21 = (v21 / ((42) | 1))
                switch fi20 {
                case 1:
                    v22 := []int{(fi20 % (((fi20 ^ (fi20 + (fi20 | fi20)))) | 1))}
                    _ = v22
                default:
                    var v23 int32 = v2
                    _ = v23
                }
            } else {
                v2 = 0
            }
            fmt.Println(v3)
            _ = fi20
            _ = v21
        }
        for fi24 := 0; fi24 < 19; fi24++ {
            var v25 string = (("\n\t" + "a") + "hello")
            v3 -= 135
            var v26 int64 = -768
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 int8 = -1
        _ = v27
    } else {
        switch "hello" {
        case "a":
            if v5 {
                var v28 float64 = ((0.1 * 3.14159) * 1.5)
                for fi29 := 0; fi29 < 9; fi29++ {
                    v30 := GenS1{F0: (v3 % (((v3 << (((v1 ^ uint(v15))) & 63))) | 1)), F1: 9223372036854775807}
                    switch (fi29 >> ((v1) & 63)) {
                    case 127:
                        switch (-(fi29)) {
                        case 413:
                            v31 := []int{}
                            _ = v31
                        case 127:
                            if v5 {
                                for fi32 := 0; fi32 < 13; fi32++ {
                                    var v33 uint = v1
                                    _ = fi32
                                    _ = v33
                                }
                                var v34 int = fi29
                                _ = v34
                            }
                        default:
                            v35 := GenS1{F0: v3, F1: 9223372036854775807}
                            _ = v35
                        }
                    case -954:
                        fmt.Println(v2, v5, fi29)
                    default:
                        v36 := []int{(fi29 + (^(fi29))), fi29}
                        _ = v36
                    }
                    _ = fi29
                    _ = v30
                }
                if v5 {
                    v37 := GenS1{F0: (v3 / (((v3 + (^(v3)))) | 1)), F1: 42}
                    _ = v37
                }
                _ = v28
            }
        case "":
            for fi38 := 0; fi38 < 13; fi38++ {
                fmt.Println(v3, v4, v15)
                v5 = v5
                _ = fi38
            }
        case "foo bar":
            fmt.Println(v3, v15)
        }
        v2 = (v2 - v2)
    }
    v1 = (v1 * v1)
    if v5 {
        v1 -= (v1 + v1)
    }
    var v39 int = 128
    var v40 bool = false
    v41 := GenS2{F0: 42, F1: (^(v3)), F2: ((!(-108 <= 0)) || v40)}
    if v40 {
        for fi42 := 0; fi42 < 5; fi42++ {
            var v43 float64 = ((float64((1.5 + (((float32((0.1 / 1.0)) * 3.14159) + 0.1) - 0.0))) * ((0.0 + -2.25) - 0.1)) / 3.14159)
            fmt.Println(v1)
            _ = fi42
            _ = v43
        }
        v44 := GenS3{F0: 2}
        v45 := []int{(v39 & -2147483648), v39, v39, v39}
        _ = v44
        _ = v45
    }
    var v46 uint = (v1 + 42)
    switch 32767 {
    case 180:
        for fi47 := 0; fi47 < 8; fi47++ {
            var v48 float32 = 1e10
            _ = fi47
            _ = v48
        }
    }
    var v49 uint8 = v15
    switch int(v3) {
    case -491:
        v50 := GenS2{F0: 0, F1: 42, F2: (!(v3 <= 1))}
        _ = v50
    case 127:
        var v51 int8 = -89
        _ = v51
    case -9223372036854775808:
        v52 := []int{}
        _ = v52
    }
    v49 -= v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v46
    _ = v49
    _ = v5
}
