package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint; F1 uint32; F2 int64 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint8, p1 uint64, p2 int16) uint8 {
    for fi4 := 0; fi4 < 4; fi4++ {
        fmt.Println(p1)
        if ((((-9223372036854775808 > -9223372036854775808) || (42 <= 42)) || ("a" != "hello")) || ("a" >= "")) {
            if false {
                if (float64(1.5) < (0.0 + (float64(0.0) + (((float64(1.5) * 1.5) + -2.25) * float64(0.0))))) {
                    switch "foo bar" {
                    case "foo bar":
                        var v5 int64 = 201
                        _ = v5
                    case "":
                        var v6 uint32 = 2
                        _ = v6
                    }
                    if (!true) {
                        p2 = 1000
                    } else {
                        p1 = 18446744073709551615
                    }
                    v7 := []int{fi4, (fi4 & 32768), (fi4 - int(p2))}
                    _ = v7
                }
                v8 := GenS1{F0: true}
                v9 := []int{fi4, -9223372036854775808}
                _ = v8
                _ = v9
            }
            fmt.Println(fi4, p1)
            v10 := GenS2{F0: "\n\t"}
            _ = v10
        }
        _ = fi4
    }
    v11 := GenS2{F0: "foo bar"}
    if ((65535 != 881) && (!false)) {
        for fi12 := 0; fi12 < 11; fi12++ {
            if (!((p1 & 952) <= p1)) {
                v13 := []int{fi12, fi12, fi12}
                for fi14 := 0; fi14 < 16; fi14++ {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    if (0.0 == 1.5) {
                        p2 -= -800
                        p2 = 1000
                        if true {
                            v13 = append(v13, (-(fi14)))
                        }
                        if len(v13) > 0 {
                            v13[0] = 743
                        }
                        v15 := GenS2{F0: "\n\t"}
                        _ = v15
                    } else {
                        var v16 float64 = (0.0 - 0.0)
                        _ = v16
                    }
                    switch fi12 {
                    case 496:
                        v17 := GenS3{F0: 127, F1: -1}
                        _ = v17
                    default:
                        _ = len(v13)
                    }
                    if false {
                        if true {
                            p2 = (p2 - (-(p2)))
                            v18 := GenS4{F0: 1, F1: 2, F2: 1}
                            v19 := GenS1{F0: true}
                            _ = v18
                            _ = v19
                        } else {
                            var v20 float64 = 0.1
                            _ = v20
                        }
                        v13 = append(v13, fi12)
                        var v21 uint = 255
                        if len(v13) > 0 {
                            v13[0] = fi12
                        }
                        _ = v21
                    }
                    _ = fi14
                }
                if (!(p2 != -1)) {
                    p0 = p0
                    v22 := GenS5{F0: (-9223372036854775808 > 9223372036854775807), F1: p1}
                    p0 = p0
                    _ = v22
                } else {
                    if false {
                        var v23 uint = 1
                        fmt.Println(p2, p1)
                        _ = v23
                    } else {
                        fmt.Println(fi12)
                        fmt.Println(p1)
                        p2 -= (p2 & 42)
                    }
                    if (1.5 != (float32((0.0 / 3.14159)) * 0.0)) {
                        var v24 float64 = (0.1 + 0.0)
                        var v25 uint8 = (p0 - (p0 | p0))
                        _ = v24
                        _ = v25
                    } else {
                        for fi26 := 0; fi26 < 3; fi26++ {
                            var v27 uint16 = 2
                            var v28 bool = false
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        var v29 int64 = 9223372036854775807
                        for fi30 := 0; fi30 < 19; fi30++ {
                            var v31 int32 = 1
                            _ = fi30
                            _ = v31
                        }
                        _ = v29
                    }
                    var v32 int = (^(fi12))
                    _ = v32
                }
                _ = v13
            }
            _ = fi12
        }
        p2 = (p2 >> ((255) & 63))
        p2 = p2
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        switch (fi33 << ((255) & 63)) {
        case 2147483647:
            for fi34 := 0; fi34 < 18; fi34++ {
                p0 -= 2
                _ = fi34
            }
        }
        var v35 uint64 = (p1 | (p1 * 535))
        p2 = p2
        _ = fi33
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (p0 >> ((2) & 63))
}

func main() {
    var v1 int = 0
    var v2 int8 = 46
    var v3 uint16 = 42
    if (0 > 0) {
        switch "foo bar" {
        case "":
            var v36 float64 = (-2.25 + 0.1)
            _ = v36
        default:
            v37 := GenS6{F0: 0, F1: 1}
            _ = v37
        }
        if false {
            v1 += (-(v1))
            v3 = v3
            v38 := GenS4{F0: 4294967295, F1: 844, F2: -257}
            _ = v38
        }
        fmt.Println(v3, v1)
    }
    v39 := GenS4{F0: 65535, F1: 65535, F2: 454}
    v40 := []int{}
    v41 := GenS7{F0: "foo bar", F1: ((0.0 * 3.14159) < 3.14159)}
    var v42 float32 = 1e10
    var v43 int64 = 841
    if len(v40) >= 2 {
        v40 = v40[1:len(v40)]
    }
    v3 -= (v3 >> ((255) & 63))
    v42 = v42
    v2 += v2
    v42 += (((((v42 / v42) + 3.14159) / v42) - -2.25) / v42)
    v44 := GenS1{F0: false}
    switch 9223372036854775807 {
    case 9223372036854775807:
        for fi45 := 0; fi45 < 20; fi45++ {
            v46 := GenS8{F0: (v3 | 388), F1: (v43 + v43)}
            if (true || ((v43 + 1) > v43)) {
                v47 := GenS2{F0: ("\n\t" + "foo bar")}
                _ = v47
            }
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
            _ = fi45
            _ = v46
        }
    case 4294967295:
        for fi48 := 0; fi48 < 3; fi48++ {
            v3 = (v3 - (v3 ^ 255))
            var v49 int8 = (v2 & -128)
            var v50 uint8 = 137
            _ = fi48
            _ = v49
            _ = v50
        }
    default:
        for fi51 := 0; fi51 < 17; fi51++ {
            v1 += fi51
            for fi52 := 0; fi52 < 2; fi52++ {
                if true {
                    if (947 >= fi52) {
                        fmt.Println(v3, v2)
                    }
                }
                var v53 int16 = 161
                var v54 int16 = (^(v53))
                _ = fi52
                _ = v53
                _ = v54
            }
            if len(v40) > 0 {
                v40[0] = 1
            }
            _ = fi51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
