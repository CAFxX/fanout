package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 bool; F2 string }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint32 }
func genFn0(p0 int64, p1 string, p2 uint64) int64 {
    switch -676 {
    case 127:
        var v6 uint64 = (p2 * p2)
        _ = v6
    case 42:
        v7 := []int{-9223372036854775808}
        _ = v7
    case -761:
        switch 42 {
        case -112:
            if true {
                switch -1 {
                case 32768:
                    for fi8 := 0; fi8 < 15; fi8++ {
                        for fi9 := 0; fi9 < 15; fi9++ {
                            v10 := []int{(^(fi8)), fi9}
                            p0 = -1
                            if len(v10) > 0 {
                                v10[0] = 128
                            }
                            var v11 uint16 = 255
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                        p0 = (^(p0))
                        _ = fi8
                    }
                case -948:
                    var v12 float64 = 1e10
                    _ = v12
                }
                var v13 bool = (false || false)
                v14 := []int{-1, int(p0)}
                _ = v13
                _ = v14
            } else {
                fmt.Println(p1, p1, p1)
                p2 += p2
            }
        case 0:
            var v15 int = 24
            _ = v15
        case 42:
            v16 := GenS1{F0: (^(p2))}
            _ = v16
        default:
            for fi17 := 0; fi17 < 19; fi17++ {
                if true {
                    v18 := GenS1{F0: (p2 ^ (p2 * p2))}
                    _ = v18
                } else {
                    v19 := GenS1{F0: (p2 ^ p2)}
                    _ = v19
                }
                _ = fi17
            }
        }
    }
    var v20 float32 = 0.1
    fmt.Println(v20, p0, v20)
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 791
}
func genFn1(p0 uint32) int8 {
    if true {
        var v21 bool = true
        _ = v21
    }
    fmt.Println(p0)
    if (100 > -65) {
        var v22 bool = (!true)
        _ = v22
    } else {
        switch -2147483648 {
        case 32767:
            v23 := GenS2{F0: 42}
            _ = v23
        case -9223372036854775808:
            p0 += p0
        case 0:
            v24 := []int{}
            _ = v24
        default:
            p0 = p0
        }
        p0 += (p0 << ((0) & 63))
        var v25 uint64 = 2
        _ = v25
    }
    _ = p0
    return 0
}

func main() {
    var v1 float32 = float32(0.1)
    var v2 string = ((((("\n\t" + "") + "") + "") + "\n\t") + "a")
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    var v26 int16 = 42
    for fi27 := 0; fi27 < 13; fi27++ {
        v1 += (1.5 / (v1 / v1))
        var v28 int8 = 127
        _ = fi27
        _ = v28
    }
    v3[v2] = -906
    var v29 float32 = v1
    if true {
        switch 42 {
        case 112:
            v29 = -2.25
        default:
            v29 += 0.0
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            v31 := []int{4294967295}
            var v32 uint16 = 1
            v2 = ((v2 + v2) + v2)
            _ = fi30
            _ = v31
            _ = v32
        }
        fmt.Println(v1, v2)
    } else {
        for fi33 := 0; fi33 < 8; fi33++ {
            switch (("\n\t" + ("hello" + "a")) + "") {
            case "\n\t":
                v2 = ((("\n\t" + v2) + v2) + v2)
            }
            _ = fi33
        }
        var v34 int16 = (^(v26))
        var v35 int8 = -126
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        if (true && (!(0 != -1))) {
            v37 := GenS3{F0: fi36, F1: true, F2: "foo bar"}
            _ = v37
        } else {
            var v38 uint16 = 2
            _ = v38
        }
        _ = fi36
    }
    var v39 int32 = -1
    for fi40 := 0; fi40 < 5; fi40++ {
        v1 += 1e10
        var v41 string = v2
        switch (fi40 | fi40) {
        case 584:
            v42 := GenS4{F0: (fi40 / ((fi40) | 1)), F1: 392}
            _ = v42
        case -1:
            fmt.Println(fi40, v29, v2)
        default:
            switch fi40 {
            case -2147483648:
                for fi43 := 0; fi43 < 7; fi43++ {
                    v39 = (v39 << ((777) & 63))
                    _ = fi43
                }
            case -9223372036854775808:
                for fi44 := 0; fi44 < 19; fi44++ {
                    if ((!true) && true) {
                        v45 := GenS5{F0: -1, F1: 4294967295}
                        _ = v45
                    }
                    var v46 float64 = -2.25
                    _ = fi44
                    _ = v46
                }
            case 2147483647:
                v29 = 0.0
            }
        }
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
