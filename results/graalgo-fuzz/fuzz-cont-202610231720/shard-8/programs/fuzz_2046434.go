package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint, p1 uint, p2 string) int64 {
    v7 := GenS1{F0: (p0 - p0)}
    if true {
        for fi8 := 0; fi8 < 10; fi8++ {
            p2 = ("" + (("\n\t" + p2) + ((p2 + (p2 + p2)) + "hello")))
            switch (fi8 >> ((p0) & 63)) {
            case 220:
                if (-1 > 0) {
                    v9 := []int{650}
                    v10 := GenS1{F0: 1}
                    for fi11 := 0; fi11 < 1; fi11++ {
                        fmt.Println(p1, p2)
                        var v12 uint16 = 32767
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                    _ = v9
                }
            case 32767:
                for fi13 := 0; fi13 < 1; fi13++ {
                    v14 := []int{1, (fi13 * fi8), fi8, 1}
                    var v15 uint64 = 1
                    p1 += 255
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case -1:
                for fi16 := 0; fi16 < 11; fi16++ {
                    if true {
                        v17 := GenS1{F0: (^(p0))}
                        _ = v17
                    }
                    _ = fi16
                }
            default:
                p2 = ("\n\t" + "")
            }
            _ = fi8
        }
        var v18 uint64 = 2
        if true {
            if false {
                v19 := []int{}
                p2 = p2
                _ = v19
            } else {
                p0 -= (p1 - (p0 / (((p1 & p0)) | 1)))
                var v20 string = p2
                _ = v20
            }
            p1 = (p1 - (p1 << ((p0) & 63)))
        } else {
            fmt.Println(p1, p1)
            switch 42 {
            case 4294967295:
                v21 := GenS1{F0: 255}
                _ = v21
            case 118:
                var v22 string = p2
                _ = v22
            case -457:
                fmt.Println(p2, v18, v18)
            }
            var v23 float32 = 1e10
            _ = v23
        }
        _ = v18
    } else {
        switch 288 {
        case 2147483647:
            p0 -= p0
        case 546:
            var v24 string = ((((p2 + ((p2 + p2) + ("\n\t" + ""))) + (p2 + p2)) + (p2 + ((p2 + p2) + p2))) + "\n\t")
            _ = v24
        case 32768:
            if false {
                v25 := GenS2{F0: 0, F1: 693, F2: 127}
                _ = v25
            }
        }
        p2 = p2
    }
    v26 := GenS2{F0: 18446744073709551615, F1: 749, F2: 120}
    fmt.Println(p0, p2)
    var v27 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    _ = v7
    return 42
}

func main() {
    var v1 int8 = 127
    var v2 bool = false
    var v3 bool = ((float32(((0.0 - ((0.0 - 3.14159) - (-2.25 + 1e10))) - -2.25)) / (0.1 - -2.25)) <= (1e10 / (0.1 + 0.1)))
    var v4 int8 = (v1 ^ v1)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]bool)
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := GenS3{F0: (v4 & v1)}
        if (v2 && (genFn0(65535, 1, "\n\t") != genFn0(18446744073709551615, 18446744073709551615, "\n\t"))) {
            v1 = v1
            var v30 int8 = (-(v1))
            fmt.Println(v30)
            _ = v30
        }
        v31 := GenS3{F0: (-(v1))}
        _ = fi28
        _ = v29
        _ = v31
    }
    var v32 string = "hello"
    var v33 bool = ("\n\t" >= (v32 + (v32 + v32)))
    if false {
        for fi34 := 0; fi34 < 16; fi34++ {
            v35 := []int{(fi34 % (((fi34 << ((42) & 63))) | 1)), fi34, 42, (fi34 & fi34)}
            v4 += v4
            for fi36 := 0; fi36 < 18; fi36++ {
                if (!(!((!false) && (675 <= 798)))) {
                    if false {
                        v35 = append(v35, fi34)
                    }
                    if (478 < 0) {
                        var v37 int32 = -816
                        _ = v37
                    } else {
                        v32 = ((v32 + v32) + v32)
                    }
                }
                _ = fi36
            }
            _ = fi34
            _ = v35
        }
        v32 = v32
    }
    fmt.Println(v33)
    v1 += 0
    switch -455 {
    case 9223372036854775807:
        var v38 bool = true
        _ = v38
    case 0:
        v39 := GenS2{F0: 0, F1: 255, F2: 127}
        _ = v39
    default:
        if v33 {
            v4 += v4
        } else {
            var v40 string = v32
            _ = v40
        }
    }
    switch -169 {
    case 53:
        fmt.Println(v32, v33, v1)
    }
    v41 := GenS1{F0: 18446744073709551615}
    switch v32 {
    case "foo bar":
        v4 -= v1
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        v3 = (1000 == -1)
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
