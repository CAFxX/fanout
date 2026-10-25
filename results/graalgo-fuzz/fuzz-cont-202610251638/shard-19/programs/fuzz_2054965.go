package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 bool; F2 uint }
type GenS5 struct { F0 int8 }
func genFn0(p0 int32) int16 {
    for fi4 := 0; fi4 < 2; fi4++ {
        p0 = p0
        fmt.Println(fi4, p0)
        v5 := GenS1{F0: -1}
        _ = fi4
        _ = v5
    }
    var v6 bool = true
    _ = p0
    _ = v6
    return 32767
}
func genFn1(p0 int16, p1 string) int8 {
    p1 = "\n\t"
    var v7 uint64 = 2
    p0 -= genFn0(-1)
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 bool = true
    var v3 int64 = (v1 % (((v1 / (((v1 - (v1 | (v1 / ((v1) | 1))))) | 1))) | 1))
    var v8 int64 = v3
    switch 9223372036854775807 {
    case 128:
        var v9 uint = 57
        _ = v9
    case -698:
        v10 := []int{32767, 129}
        _ = v10
    }
    v11 := GenS2{F0: 32768, F1: 255, F2: (!v2)}
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(v3, v1, fi12)
        v13 := []int{(fi12 * (fi12 % (((fi12 * (fi12 & fi12))) | 1))), fi12, fi12, (^(fi12))}
        _ = fi12
        _ = v13
    }
    fmt.Println(v8, v1)
    if true {
        var v14 bool = v2
        _ = v14
    }
    var v15 int32 = 42
    if false {
        v16 := GenS2{F0: 42, F1: 873, F2: (0 > -9223372036854775808)}
        v3 = (v3 << ((1) & 63))
        var v17 uint16 = 32767
        _ = v16
        _ = v17
    } else {
        var v18 uint64 = 18446744073709551615
        _ = v18
    }
    v19 := GenS3{F0: false}
    switch -1 {
    case 128:
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := []int{(fi20 * -533), fi20, fi20, fi20}
            if len(v21) > 0 {
                v21[0] = int(v15)
            }
            _ = fi20
            _ = v21
        }
    case 0:
        for fi22 := 0; fi22 < 19; fi22++ {
            for fi23 := 0; fi23 < 12; fi23++ {
                fmt.Println(v2, fi23, v15)
                v24 := []int{}
                v25 := GenS4{F0: 2, F1: (42 >= 222), F2: 2}
                _ = fi23
                _ = v24
                _ = v25
            }
            fmt.Println(v2, fi22)
            _ = fi22
        }
    default:
        for fi26 := 0; fi26 < 8; fi26++ {
            v1 += (^(v1))
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 int16 = -802
                var v29 uint = 2
                v15 += v15
                var v30 uint32 = 4294967295
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            _ = fi26
        }
    }
    v31 := GenS4{F0: 747, F1: false, F2: 18446744073709551615}
    v32 := []int{32768, 316, -90, 4294967295}
    if (1 >= 0) {
        if (v2 && true) {
            switch 32767 {
            case -2147483648:
                v8 = (v3 >> ((65535) & 63))
            default:
                switch -9223372036854775808 {
                case 4294967295:
                    for fi33 := 0; fi33 < 4; fi33++ {
                        for fi34 := 0; fi34 < 18; fi34++ {
                            fmt.Println(v1, v15, fi33)
                            v32 = append(v32, fi34)
                            var v35 float32 = 1.5
                            v32 = append(v32, -2147483648)
                            _ = fi34
                            _ = v35
                        }
                        _ = fi33
                    }
                case 683:
                    v1 += (v3 - (v3 - v1))
                case 0:
                    var v36 string = ("hello" + "hello")
                    _ = v36
                default:
                    if v2 {
                        _ = len(v32)
                        if len(v32) >= 2 {
                            v32 = v32[1:len(v32)]
                        }
                    } else {
                        if v2 {
                            var v37 uint32 = 4294967295
                            _ = v37
                        }
                        v8 = v8
                        var v38 float64 = (3.14159 - 0.1)
                        _ = v38
                    }
                }
            }
            fmt.Println(v3)
            for fi39 := 0; fi39 < 14; fi39++ {
                v40 := GenS5{F0: int8(v15)}
                _ = fi39
                _ = v40
            }
        } else {
            v2 = v2
            var v41 bool = true
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 5; fi42++ {
        if true {
            v43 := GenS3{F0: false}
            for fi44 := 0; fi44 < 13; fi44++ {
                v2 = v2
                _ = fi44
            }
            switch "foo bar" {
            case "\n\t":
                for fi45 := 0; fi45 < 8; fi45++ {
                    fmt.Println(fi45)
                    var v46 string = ("a" + "foo bar")
                    _ = fi45
                    _ = v46
                }
            case "hello":
                var v47 uint16 = 995
                _ = v47
            default:
                fmt.Println(v1, v1, v1)
            }
            _ = v43
        } else {
            v1 = v1
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v8
}
