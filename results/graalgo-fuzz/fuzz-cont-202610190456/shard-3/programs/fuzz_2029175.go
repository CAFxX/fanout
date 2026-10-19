package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint32; F2 string }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint32, p1 bool) uint {
    if p1 {
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: 0}
            var v9 bool = ((float32(-2.25) + 1e10) == (0.1 / 1.5))
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{9223372036854775807, 0, 2147483647}
        if p1 {
            for fi11 := 0; fi11 < 8; fi11++ {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                if false {
                    fmt.Println(p1, p1)
                }
                _ = fi11
            }
            p1 = false
            for fi12 := 0; fi12 < 11; fi12++ {
                switch fi12 {
                case 127:
                    p0 += 1
                default:
                    for fi13 := 0; fi13 < 7; fi13++ {
                        p0 -= p0
                        var v14 int8 = -128
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        v15 := GenS2{F0: -333}
                        var v16 float32 = float32(3.14159)
                        v10 = append(v10, fi13)
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                }
                _ = fi12
            }
        }
        _ = v10
    } else {
        p1 = true
        var v17 int8 = 127
        _ = v17
    }
    var v18 uint = 4294967295
    if (42 == -2147483648) {
        v19 := []int{}
        for fi20 := 0; fi20 < 7; fi20++ {
            if (((p0 == 42) && ((p1 || ("hello" >= "a")) || false)) || true) {
                v18 -= (v18 | (^(v18)))
                var v21 int32 = 160
                if p1 {
                    if true {
                        if (!false) {
                            var v22 uint = 255
                            _ = v22
                        } else {
                            var v23 int16 = 691
                            if p1 {
                                for fi24 := 0; fi24 < 14; fi24++ {
                                    var v25 int32 = (-(v21))
                                    v19 = append(v19, fi24)
                                    var v26 int32 = (v21 + (^(v21)))
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                                v27 := GenS2{F0: (-(v23))}
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 11; fi28++ {
                                v21 -= v21
                                v29 := GenS3{F0: 42, F1: p0, F2: "foo bar"}
                                _ = fi28
                                _ = v29
                            }
                            _ = v23
                        }
                        v21 -= -2147483648
                        if (p1 || (uint8(v18) != 121)) {
                            v30 := GenS3{F0: 100, F1: p0, F2: "\n\t"}
                            switch ("\n\t" + "\n\t") {
                            case "":
                                var v31 int16 = 0
                                _ = v31
                            case "\n\t":
                                v19 = append(v19, (fi20 * (^(fi20))))
                            case "hello":
                                v32 := GenS4{F0: 9223372036854775807, F1: "hello"}
                                _ = v32
                            }
                            fmt.Println(v21, v18, p1)
                            _ = v30
                        } else {
                            v21 -= v21
                        }
                    }
                    fmt.Println(v18, v21, p0)
                    var v33 int = fi20
                    _ = v33
                } else {
                    if (p1 && (p0 == p0)) {
                        if len(v19) > 0 {
                            v19[0] = fi20
                        }
                    } else {
                        for fi34 := 0; fi34 < 20; fi34++ {
                            switch fi34 {
                            case 658:
                                var v35 float32 = float32(0.1)
                                _ = v35
                            default:
                                v19 = append(v19, fi20)
                            }
                            _ = fi34
                        }
                        v21 += v21
                        v21 += v21
                    }
                    var v36 uint16 = 404
                    _ = v36
                }
                _ = v21
            }
            _ = fi20
        }
        _ = v19
    } else {
        fmt.Println(p1, v18)
        for fi37 := 0; fi37 < 5; fi37++ {
            if (245 != -1) {
                if (true && false) {
                    v38 := []int{fi37}
                    _ = v38
                }
            } else {
                var v39 int16 = 134
                var v40 uint = (^(v18))
                _ = v39
                _ = v40
            }
            if (true || p1) {
                v41 := GenS3{F0: -128, F1: (p0 + p0), F2: "a"}
                v42 := []int{fi37, fi37, (^(fi37)), 27}
                for fi43 := 0; fi43 < 1; fi43++ {
                    if (fi37 > fi43) {
                        if (!((0.0 <= 3.14159) && p1)) {
                            p1 = (p1 || (p0 > p0))
                        }
                        var v44 float64 = 1.5
                        p1 = true
                        p0 += (p0 >> (((v18 % (((v18 * (v18 >> (((v18 % (((^(v18))) | 1))) & 63)))) | 1))) & 63))
                        _ = v44
                    }
                    p1 = (!(275 <= -342))
                    _ = fi43
                }
                _ = v41
                _ = v42
            }
            _ = fi37
        }
        var v45 int32 = 42
        _ = v45
    }
    v46 := GenS1{F0: uint64(p0)}
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v18, v18)
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v46
    return (v18 & v18)
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = float64((float32(3.14159) * -2.25))
    var v3 string = "a"
    var v4 uint = 888
    v5 := make(map[int]int)
    v6 := make(map[int]uint32)
    delete(v5, 4294967295)
    delete(v1, 1)
    v4 += (v4 | (v4 % ((genFn0(0, (true && (!true)))) | 1)))
    for fi47 := 0; fi47 < 16; fi47++ {
        v48 := []int{(^(fi47))}
        v49 := v6[fi47]
        v49 += 173
        _ = fi47
        _ = v48
        _ = v49
    }
    if true {
        v50 := GenS3{F0: 0, F1: 460, F2: v3}
        _ = v50
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        var v52 int32 = 0
        _ = fi51
        _ = v52
    }
    if true {
        fmt.Println(v4, v4)
        var v53 uint8 = 127
        v53 = 42
        _ = v53
    }
    var v54 uint = 926
    for fi55 := 0; fi55 < 3; fi55++ {
        switch -44 {
        case 127:
            v56 := GenS5{F0: 0}
            _ = v56
        case -2147483648:
            for fi57 := 0; fi57 < 6; fi57++ {
                v58 := v6[(fi55 * (fi57 + fi57))]
                if false {
                    v4 = (v4 << ((v4) & 63))
                    if false {
                        v59 := []int{}
                        _ = v59
                    }
                }
                v1[fi57] = -16
                _ = fi57
                _ = v58
            }
        case -955:
            var v60 int64 = -593
            _ = v60
        default:
            v4 -= (v54 & 326)
        }
        fmt.Println(v4)
        fmt.Println(fi55)
        _ = fi55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
