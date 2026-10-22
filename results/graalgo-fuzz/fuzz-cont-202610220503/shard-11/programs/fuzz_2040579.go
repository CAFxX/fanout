package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int32; F2 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 int32; F2 uint64 }
func genFn0(p0 uint64, p1 uint16, p2 int8) string {
    v7 := GenS1{F0: 32768}
    v8 := GenS1{F0: 333}
    v9 := GenS1{F0: -1}
    for fi10 := 0; fi10 < 13; fi10++ {
        if (fi10 < (-(fi10))) {
            p2 -= (-(p2))
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 float32 = 1.5
                switch (((("a" + "") + ("foo bar" + "foo bar")) + "a") + "") {
                case "hello":
                    v12 += v12
                }
                var v13 int8 = -1
                _ = fi11
                _ = v12
                _ = v13
            }
            switch fi10 {
            case 665:
                switch ("foo bar" + "a") {
                case "":
                    var v14 uint = 2
                    _ = v14
                default:
                    p1 = (p1 << ((263) & 63))
                }
            default:
                v15 := GenS2{F0: -2147483648}
                _ = v15
            }
        }
        v16 := GenS3{F0: 793}
        _ = fi10
        _ = v16
    }
    switch 4294967295 {
    case 127:
        fmt.Println(p2, p0, p2)
    default:
        if (!(!(!false))) {
            v17 := []int{}
            for fi18 := 0; fi18 < 18; fi18++ {
                for fi19 := 0; fi19 < 12; fi19++ {
                    p0 = (p0 << ((279) & 63))
                    _ = fi19
                }
                if false {
                    var v20 uint = uint(p1)
                    if (!(uint16(p0) != (p1 % (((p1 - uint16(p0))) | 1)))) {
                        for fi21 := 0; fi21 < 6; fi21++ {
                            fmt.Println(fi21)
                            v20 = v20
                            _ = fi21
                        }
                    } else {
                        fmt.Println(p2)
                        fmt.Println(v20, fi18, fi18)
                        v22 := GenS2{F0: int32(p1)}
                        _ = v22
                    }
                    _ = v20
                } else {
                    fmt.Println(p0)
                    fmt.Println(p1, fi18, p0)
                }
                for fi23 := 0; fi23 < 2; fi23++ {
                    fmt.Println(p1)
                    _ = fi23
                }
                _ = fi18
            }
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 int16 = -405
    var v2 uint = 1
    var v3 uint32 = 2147483647
    var v4 int32 = 42
    v5 := make(map[uint32]int)
    var v6 int8 = 127
    if true {
        v24 := GenS4{F0: "hello", F1: -675, F2: (!((1.5 / 3.14159) > (3.14159 - 1.5)))}
        v25 := []int{1, 4294967295}
        v5[v3] = 2147483647
        _ = v24
        _ = v25
    }
    var v26 int = 9223372036854775807
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 bool = true
        _ = fi27
        _ = v28
    }
    var v29 uint32 = 2147483647
    var v30 int = (v26 << (((v2 | v2)) & 63))
    var v31 bool = (!true)
    var v32 uint8 = 1
    v33 := []int{v26}
    if (!v31) {
        v29 += 823
    } else {
        var v34 uint = (v2 * (v2 + v2))
        if (!(v26 != v26)) {
            for fi35 := 0; fi35 < 9; fi35++ {
                var v36 float64 = ((((float64(0.1) * float64(0.0)) * 3.14159) + 0.0) * ((((float64(float32(0.1)) - float64(0.1)) - 3.14159) - ((float64(0.1) / 1.0) / (1.5 * 1.5))) / 1e10))
                for fi37 := 0; fi37 < 9; fi37++ {
                    v38 := GenS4{F0: "a", F1: v4, F2: (fi37 == 128)}
                    if ((!v31) && v31) {
                        v34 = 255
                        delete(v5, (v3 + 2147483648))
                        var v39 int32 = v4
                        _ = v39
                    }
                    var v40 uint64 = 18446744073709551615
                    _ = fi37
                    _ = v38
                    _ = v40
                }
                _ = fi35
                _ = v36
            }
            v31 = ((1e10 / -2.25) == float32((-2.25 - 1e10)))
            switch v26 {
            case 1:
                switch int(v32) {
                case 2147483647:
                    for fi41 := 0; fi41 < 20; fi41++ {
                        if (false && (128 != 56)) {
                            if (v6 != v6) {
                                v42 := GenS5{F0: 1}
                                fmt.Println(v26, v6)
                                var v43 int8 = (v6 % ((v6) | 1))
                                _ = v42
                                _ = v43
                            } else {
                                for fi44 := 0; fi44 < 17; fi44++ {
                                    var v45 int16 = (v1 << (((^(v34))) & 63))
                                    _ = fi44
                                    _ = v45
                                }
                            }
                            v46 := GenS4{F0: "hello", F1: 2147483647, F2: (v30 != fi41)}
                            _ = v46
                        }
                        _ = fi41
                    }
                case 32768:
                    var v47 int8 = v6
                    _ = v47
                case 254:
                    v48 := GenS5{F0: int32(v29)}
                    _ = v48
                }
            case -141:
                v30 -= v26
            default:
                for fi49 := 0; fi49 < 13; fi49++ {
                    v30 += 4294967295
                    _ = fi49
                }
            }
        } else {
            fmt.Println(v1, v2)
            switch (v30 * -9223372036854775808) {
            case -1:
                var v50 int16 = v1
                _ = v50
            case 360:
                v2 += v2
            }
        }
        var v51 int32 = -2147483648
        _ = v34
        _ = v51
    }
    v33 = append(v33, (^(v26)))
    fmt.Println(v3, v1, v31)
    switch v30 {
    case 1:
        if true {
            var v52 float32 = -2.25
            _ = v52
        } else {
            fmt.Println(v2, v1, v3)
            if (1e10 < 3.14159) {
                if len(v33) >= 2 {
                    v33 = v33[1:len(v33)]
                }
                v53 := GenS4{F0: "\n\t", F1: (v4 ^ (v4 & (v4 / ((v4) | 1)))), F2: false}
                if v31 {
                    switch (^(v30)) {
                    case -316:
                        v30 = (v26 | (^(v30)))
                    case 68:
                        var v54 int64 = 627
                        _ = v54
                    case 861:
                        var v55 uint8 = (v32 >> (((v2 << (((v2 - 0)) & 63))) & 63))
                        _ = v55
                    }
                } else {
                    v5[1] = -1
                    var v56 int = v26
                    v32 = v32
                    _ = v56
                }
                _ = v53
            }
        }
    case 998:
        for fi57 := 0; fi57 < 16; fi57++ {
            var v58 int32 = (v4 << ((v2) & 63))
            var v59 uint16 = 516
            _ = fi57
            _ = v58
            _ = v59
        }
    case 9223372036854775807:
        switch "foo bar" {
        case "\n\t":
            v3 = (^(v29))
        }
    default:
        v60 := GenS6{F0: v2, F1: (v4 + -57), F2: 18446744073709551615}
        _ = v60
    }
    var v61 uint64 = 2
    for fi62 := 0; fi62 < 15; fi62++ {
        v63 := GenS4{F0: "", F1: (v4 << (((v2 ^ v2)) & 63)), F2: (int(v2) != v30)}
        v3 = v29
        fmt.Println(v29)
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
