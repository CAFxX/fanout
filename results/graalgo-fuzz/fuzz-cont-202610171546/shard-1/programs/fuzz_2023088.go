package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint = 2
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    var v4 bool = (false && (!false))
    var v5 uint = 42
    switch -1 {
    case 32768:
        v4 = (!v4)
    case -801:
        switch 127 {
        case 263:
            var v6 uint = 164
            _ = v6
        case -2147483648:
            fmt.Println(v5, v4, v4)
        case 128:
            var v7 int8 = 1
            _ = v7
        }
    case -168:
        switch 323 {
        case 32767:
            v1 = 1
        default:
            fmt.Println(v1, v5)
        }
    default:
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(v5, fi8)
            v9 := []int{fi8}
            _ = fi8
            _ = v9
        }
    }
    var v10 uint64 = 4294967295
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 float64 = (1e10 / (-2.25 / ((1.5 - ((1.5 * 0.0) / 3.14159)) - float64(3.14159))))
            if (1e10 > -2.25) {
                v14 := GenS1{F0: v4, F1: (v10 << ((42) & 63))}
                if (42 >= 1) {
                    if (4294967295 >= 42) {
                        v15 := []int{(fi11 / ((128) | 1)), fi11, (fi12 ^ fi11)}
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        var v16 int32 = 42
                        v3[9223372036854775807] = "foo bar"
                        fmt.Println(v5, v13, v16)
                        _ = v15
                        _ = v16
                    }
                } else {
                    v1 = v5
                    switch "foo bar" {
                    case "\n\t":
                        if (!false) {
                            var v17 int = (^(fi12))
                            var v18 int64 = -199
                            switch "hello" {
                            case "\n\t":
                                var v19 uint = 947
                                _ = v19
                            case "foo bar":
                                fmt.Println(v17, v4)
                            }
                            _ = v17
                            _ = v18
                        } else {
                            var v20 float64 = v13
                            _ = v20
                        }
                        var v21 float32 = (((float32(0.1) - -2.25) * 0.0) + 3.14159)
                        _ = v21
                    case "":
                        v22 := v3[(fi12 % (((fi11 - (-(fi12)))) | 1))]
                        _ = v22
                    case "a":
                        fmt.Println(v10)
                    default:
                        var v23 int = fi12
                        _ = v23
                    }
                }
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        if v4 {
            v24 := []int{fi11, (fi11 % (((-(fi11))) | 1))}
            _ = v24
        } else {
            var v25 bool = true
            _ = v25
        }
        var v26 int16 = 1
        _ = fi11
        _ = v26
    }
    v27 := GenS1{F0: (!(31 < 0)), F1: (v10 % ((v10) | 1))}
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 int8 = 127
        _ = fi28
        _ = v29
    }
    var v30 int64 = 0
    var v31 uint16 = 1
    v32 := v3[4294967295]
    v33 := GenS2{F0: -2147483648}
    for fi34 := 0; fi34 < 3; fi34++ {
        fmt.Println(v1, v30)
        if (((false || true) || ((667 <= v31) || (!true))) || (2147483647 <= 463)) {
            v10 = v10
        } else {
            if (v31 == (v31 * 42)) {
                _ = len(v3)
                for fi35 := 0; fi35 < 19; fi35++ {
                    var v36 uint = (^(v1))
                    fmt.Println(v30, v4)
                    var v37 uint = v1
                    var v38 string = "foo bar"
                    v36 = v1
                    v36 -= (v36 * v5)
                    _ = fi35
                    _ = v36
                    _ = v37
                    _ = v38
                }
            }
            if (!(((3.14159 * (3.14159 / 1.0)) + (((-2.25 + float64(1.5)) * (3.14159 / 1e10)) - (1.5 + (1.5 - float64(float32(-2.25)))))) >= (1e10 * float64(0.1)))) {
                v30 = v30
            }
            switch (fi34 << (((v1 * 637)) & 63)) {
            case -9223372036854775808:
                var v39 uint = v5
                _ = v39
            case 2147483647:
                _ = len(v3)
            default:
                for fi40 := 0; fi40 < 16; fi40++ {
                    v41 := GenS2{F0: (^(fi40))}
                    v42 := []int{fi34, fi40, fi40}
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
        }
        _ = fi34
    }
    v30 += v30
    v30 = 688
    for fi43 := 0; fi43 < 7; fi43++ {
        v44 := []int{}
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 12; fi45++ {
        for fi46 := 0; fi46 < 9; fi46++ {
            var v47 int = fi46
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    var v48 float64 = 1e10
    v49 := GenS1{F0: ((false || (2 != 127)) && (!((1e10 * 0.0) <= (0.0 - (0.0 - (1.5 * ((1e10 * (0.1 - 1.5)) / (0.1 - 0.0)))))))), F1: v10}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
