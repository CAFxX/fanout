package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int32, p1 int, p2 int16) bool {
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: p2, F1: 17, F2: 255}
            if (!(!true)) {
                var v9 float64 = (0.1 / 1.0)
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    } else {
        var v10 uint16 = 52
        _ = v10
    }
    p1 = p1
    for fi11 := 0; fi11 < 18; fi11++ {
        p1 -= int(p2)
        var v12 int8 = 100
        _ = fi11
        _ = v12
    }
    if (!(false && true)) {
        var v13 uint64 = 2
        _ = v13
    } else {
        v14 := GenS1{F0: 1000, F1: 261, F2: 255}
        for fi15 := 0; fi15 < 1; fi15++ {
            if false {
                p1 = 42
                if true {
                    p2 += -148
                }
            }
            p0 = 2147483647
            _ = fi15
        }
        switch (-(p1)) {
        case -603:
            if (!(!true)) {
                v16 := []int{}
                _ = v16
            }
        case -1:
            p1 = p1
        case 9223372036854775807:
            var v17 string = ("a" + ("" + "a"))
            _ = v17
        }
        _ = v14
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        if false {
            switch (p1 | (^(p1))) {
            case -601:
                for fi19 := 0; fi19 < 7; fi19++ {
                    if true {
                        p1 -= (-(p1))
                        if (0.1 > 0.1) {
                            v20 := []int{}
                            _ = v20
                        } else {
                            v21 := []int{}
                            if len(v21) > 0 {
                                v21[0] = (-(fi18))
                            }
                            v22 := GenS2{F0: 0, F1: 42, F2: 0}
                            _ = v21
                            _ = v22
                        }
                        var v23 int32 = p0
                        p2 = 0
                        v24 := []int{-1, (fi19 >> ((42) & 63))}
                        var v25 int32 = v23
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = fi19
                }
            }
        }
        v26 := GenS3{F0: 0, F1: 817}
        _ = fi18
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    return (0.1 >= 1e10)
}
func genFn1(p0 int) int8 {
    if false {
        p0 = p0
        switch (^(p0)) {
        case -2147483648:
            switch (p0 ^ (p0 * (p0 % ((p0) | 1)))) {
            case 4294967295:
                v27 := []int{(p0 ^ p0), (p0 % (((p0 % ((-1) | 1))) | 1))}
                _ = v27
            default:
                var v28 int64 = 9223372036854775807
                _ = v28
            }
        default:
            switch 0 {
            case -9223372036854775808:
                for fi29 := 0; fi29 < 1; fi29++ {
                    var v30 string = "a"
                    if (true || true) {
                        p0 -= p0
                        for fi31 := 0; fi31 < 20; fi31++ {
                            v32 := []int{(fi29 >> ((18446744073709551615) & 63)), (p0 + (^(fi31))), fi31, -2147483648}
                            p0 = 128
                            _ = fi31
                            _ = v32
                        }
                    } else {
                        v33 := GenS4{F0: 287}
                        _ = v33
                    }
                    _ = fi29
                    _ = v30
                }
            }
        }
    }
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 481
    v3 := make(map[string]bool)
    var v4 int16 = 0
    v5 := make(map[int]uint32)
    var v6 int8 = 0
    switch -9 {
    case 1:
        switch 319 {
        case -440:
            fmt.Println(v4, v4)
        case 281:
            for fi34 := 0; fi34 < 9; fi34++ {
                for fi35 := 0; fi35 < 12; fi35++ {
                    v6 += -1
                    _ = fi35
                }
                _ = fi34
            }
        case -794:
            var v36 uint64 = 0
            _ = v36
        }
    case -9223372036854775808:
        switch 2147483647 {
        case 1:
            delete(v1, 2147483648)
        default:
            for fi37 := 0; fi37 < 16; fi37++ {
                v38 := GenS5{F0: 2}
                _ = fi37
                _ = v38
            }
        }
    }
    v6 = (v6 | -38)
    delete(v1, 662)
    fmt.Println(v6, v4, v4)
    var v39 bool = true
    var v40 int16 = (v4 << ((v2) & 63))
    fmt.Println(v6)
    var v41 uint32 = 42
    if v39 {
        if v39 {
            var v42 int64 = 1
            _ = v42
        }
        v43 := GenS2{F0: 42, F1: -1, F2: 2}
        _ = v43
    }
    var v44 int = 128
    var v45 int64 = 473
    var v46 float64 = float64(0.0)
    var v47 int = v44
    v48 := []int{(v44 + (^(v44))), -708, v47}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
