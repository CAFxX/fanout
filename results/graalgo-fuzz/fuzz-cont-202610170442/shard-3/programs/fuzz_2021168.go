package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int64; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint32; F1 int32 }
func genFn0(p0 int8, p1 int32) uint64 {
    v5 := []int{4294967295, -896}
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if len(v5) > 0 {
        v5[0] = -1
    }
    if (int64(p0) >= 42) {
        _ = len(v5)
        var v6 int64 = 702
        if (false || true) {
            switch "foo bar" {
            case "":
                v7 := GenS1{F0: -196}
                _ = v7
            case "a":
                var v8 float64 = 1.5
                _ = v8
            case "foo bar":
                for fi9 := 0; fi9 < 13; fi9++ {
                    fmt.Println(fi9)
                    if false {
                        for fi10 := 0; fi10 < 6; fi10++ {
                            v6 -= (-(v6))
                            var v11 int = 780
                            var v12 float64 = (-2.25 * 1.5)
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        fmt.Println(fi9, v6)
                    } else {
                        v6 = (v6 / ((v6) | 1))
                    }
                    _ = fi9
                }
            default:
                v13 := GenS2{F0: 32768}
                _ = v13
            }
            var v14 uint32 = 2147483648
            _ = v14
        }
        _ = v6
    } else {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int = fi15
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 int16 = 1000
                p0 += 100
                _ = fi17
                _ = v18
            }
            fmt.Println(p1)
            _ = fi15
            _ = v16
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            _ = len(v5)
            for fi20 := 0; fi20 < 3; fi20++ {
                p1 += p1
                if true {
                    var v21 int64 = 0
                    var v22 int64 = -9223372036854775808
                    var v23 int64 = (^(v21))
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = fi20
            }
            if true {
                v24 := GenS3{F0: p1, F1: int64(p0), F2: (!false)}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v25 float64 = -2.25
                _ = v24
                _ = v25
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi19
        }
        var v26 int16 = 44
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v5
    return 290
}
func genFn1(p0 uint64, p1 bool) int32 {
    p0 = p0
    if ((p0 * genFn0(-37, 867)) > p0) {
        v27 := []int{-822, 969, 4294967295, 9223372036854775807}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 16; fi28++ {
            for fi29 := 0; fi29 < 6; fi29++ {
                p0 += p0
                v30 := GenS2{F0: (fi28 >> ((931) & 63))}
                p1 = p1
                _ = fi29
                _ = v30
            }
            switch "a" {
            case "hello":
                v31 := GenS4{F0: 0}
                _ = v31
            case "a":
                switch "foo bar" {
                case "hello":
                    var v32 uint16 = 42
                    _ = v32
                case "a":
                    v33 := GenS2{F0: fi28}
                    _ = v33
                }
            }
            var v34 int8 = -81
            _ = fi28
            _ = v34
        }
    }
    fmt.Println(p1, p1)
    if (false || false) {
        p1 = true
    } else {
        v35 := GenS1{F0: 42}
        p1 = p1
        var v36 bool = (0 < -128)
        _ = v35
        _ = v36
    }
    _ = p0
    _ = p1
    return 761
}

func main() {
    var v1 uint16 = 736
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 int64 = -1
    if false {
        fmt.Println(v4)
    }
    if true {
        var v37 uint8 = 255
        v4 = (v4 << ((65535) & 63))
        _ = v37
    }
    var v38 string = ("a" + "hello")
    for fi39 := 0; fi39 < 3; fi39++ {
        v1 += 0
        _ = fi39
    }
    if false {
        var v40 int32 = genFn1(42, (!true))
        _ = v40
    } else {
        v4 = (v4 * v4)
        v41 := GenS2{F0: 200}
        _ = v41
    }
    delete(v2, 4294967295)
    v38 = v38
    var v42 uint8 = 128
    for fi43 := 0; fi43 < 19; fi43++ {
        fmt.Println(v38, v38, v42)
        _ = fi43
    }
    if false {
        v44 := GenS5{F0: 299, F1: (v42 % (((^(v42))) | 1)), F2: 0}
        v1 -= (v1 / (((^(v1))) | 1))
        _ = v44
    } else {
        for fi45 := 0; fi45 < 14; fi45++ {
            var v46 int8 = -115
            for fi47 := 0; fi47 < 18; fi47++ {
                var v48 int = (^(fi47))
                _ = fi47
                _ = v48
            }
            _ = fi45
            _ = v46
        }
        if (!false) {
            var v49 float32 = (1.5 * 1e10)
            v1 = (v1 * 255)
            var v50 uint16 = (v1 + v1)
            _ = v49
            _ = v50
        }
        for fi51 := 0; fi51 < 13; fi51++ {
            fmt.Println(v4)
            v38 = v38
            if true {
                fmt.Println(v42, v38, v1)
            }
            _ = fi51
        }
    }
    for fi52 := 0; fi52 < 15; fi52++ {
        v53 := GenS6{F0: 230, F1: genFn1(2, true)}
        _ = fi52
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
}
