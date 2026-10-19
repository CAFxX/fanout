package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int16; F1 string; F2 uint8 }
func genFn0(p0 int, p1 uint64) int16 {
    if (-2147483648 == 0) {
        p1 = p1
        var v5 int16 = -833
        v6 := []int{(p0 & p0), (^(p0)), -1, (p0 % ((1) | 1))}
        _ = v5
        _ = v6
    }
    switch 0 {
    case 236:
        p0 -= 9223372036854775807
    default:
        if (289 < 2) {
            switch (p0 & p0) {
            case 32768:
                for fi7 := 0; fi7 < 17; fi7++ {
                    var v8 int8 = 52
                    v9 := []int{(fi7 / (((p0 << ((427) & 63))) | 1)), (p0 >> ((255) & 63)), p0}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case 2147483647:
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 uint = 255
                    p0 = (fi10 | p0)
                    var v12 int16 = 1000
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            default:
                if (!false) {
                    var v13 uint16 = 662
                    var v14 int = p0
                    v15 := GenS1{F0: (p1 >> ((699) & 63)), F1: true}
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    v16 := []int{p0, (^(p0))}
                    _ = v16
                }
            }
            var v17 int64 = -9223372036854775808
            _ = v17
        }
    }
    if ((0.1 >= 0.1) && false) {
        fmt.Println(p0, p0)
        for fi18 := 0; fi18 < 9; fi18++ {
            p0 -= 42
            _ = fi18
        }
    } else {
        switch p0 {
        case -9223372036854775808:
            v19 := []int{(p0 | p0), 127, (^(p0))}
            _ = v19
        case -2147483648:
            switch p0 {
            case -9223372036854775808:
                var v20 int32 = 825
                _ = v20
            case 36:
                if (!(!true)) {
                    v21 := []int{(p0 ^ p0)}
                    _ = v21
                }
            }
        }
    }
    v22 := GenS1{F0: 42, F1: (!true)}
    if (32767 <= 128) {
        v23 := GenS1{F0: (^(p1)), F1: false}
        fmt.Println(p1, p0)
        _ = v23
    } else {
        var v24 float64 = 3.14159
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v22
    return -792
}
func genFn1(p0 uint64) uint8 {
    fmt.Println(p0)
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 uint32 = 368
        for fi27 := 0; fi27 < 12; fi27++ {
            switch 4294967295 {
            case -1:
                fmt.Println(v26, v26)
            }
            fmt.Println(v26)
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    _ = p0
    return 1
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[int]bool)
    var v3 bool = false
    var v4 uint16 = v1
    for fi28 := 0; fi28 < 4; fi28++ {
        if v3 {
            var v29 int16 = 1000
            _ = v29
        }
        _ = fi28
    }
    v30 := GenS2{F0: 0, F1: "hello", F2: 167}
    v3 = (1 < -1)
    var v31 uint32 = 1
    fmt.Println(v3, v1)
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 int8 = 100
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        if false {
            v35 := GenS2{F0: 643, F1: "hello", F2: 42}
            _ = v35
        } else {
            v31 += v31
        }
        v36 := []int{(fi34 % ((2147483647) | 1))}
        switch (fi34 >> ((4) & 63)) {
        case 591:
            v4 -= 255
        case 381:
            var v37 string = "\n\t"
            _ = v37
        case -835:
            _ = len(v36)
        }
        _ = fi34
        _ = v36
    }
    v38 := []int{}
    v1 = v1
    if len(v38) >= 2 {
        v38 = v38[1:len(v38)]
    }
    v31 += 211
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
}
