package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int32; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint32 }
func genFn0(p0 int8, p1 int8, p2 uint32) string {
    v7 := []int{}
    p1 -= 1
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}

func main() {
    var v1 int = 128
    var v2 int = (v1 | v1)
    v3 := make(map[uint32]bool)
    var v4 int = 4294967295
    var v5 string = "foo bar"
    var v6 int16 = 0
    switch v5 {
    case "a":
        switch (v5 + v5) {
        case "\n\t":
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 int16 = (v6 % (((v6 * (v6 - (v6 | (^(v6)))))) | 1))
                for fi10 := 0; fi10 < 4; fi10++ {
                    v11 := GenS1{F0: (v6 & v6), F1: 0, F2: "\n\t"}
                    _ = fi10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        case "a":
            v6 = (v6 - v6)
        }
    case "":
        v12 := GenS2{F0: false}
        _ = v12
    }
    var v13 int64 = -1
    fmt.Println(v2, v4, v1)
    fmt.Println(v4, v13, v2)
    v3[501] = true
    var v14 uint32 = 2
    fmt.Println(v14, v1, v6)
    v13 = int64(v14)
    var v15 float64 = -2.25
    fmt.Println(v2, v5)
    if (-128 != 0) {
        if (false || (0.0 >= 1.5)) {
            var v16 int64 = -547
            _ = v16
        } else {
            if (v15 == 1e10) {
                var v17 string = v5
                var v18 uint = 0
                switch v5 {
                case "":
                    for fi19 := 0; fi19 < 11; fi19++ {
                        fmt.Println(v13, v2, v14)
                        var v20 bool = false
                        for fi21 := 0; fi21 < 8; fi21++ {
                            if false {
                                if true {
                                    delete(v3, v14)
                                    fmt.Println(v6, v15, v6)
                                } else {
                                    v22 := GenS3{F0: 84}
                                    _ = v22
                                }
                                if v20 {
                                    if (v20 || v20) {
                                        var v23 float64 = v15
                                        _ = v23
                                    }
                                    var v24 int = fi21
                                    fmt.Println(v13, v5)
                                    _ = v24
                                } else {
                                    var v25 string = v17
                                    fmt.Println(v20, v18, v13)
                                    v26 := GenS4{F0: false, F1: 42, F2: -1}
                                    _ = v25
                                    _ = v26
                                }
                                v13 = v13
                            } else {
                                var v27 float32 = float32(((v15 + v15) / 1e10))
                                if v20 {
                                    v28 := GenS1{F0: (v6 >> ((v18) & 63)), F1: 255, F2: v17}
                                    v29 := GenS4{F0: (1e10 != (((3.14159 - v15) / (v15 / 1.5)) / v15)), F1: 42, F2: -2147483648}
                                    v3[(v14 ^ (v14 * uint32(v18)))] = v20
                                    _ = v28
                                    _ = v29
                                } else {
                                    v14 = (v14 + v14)
                                    if v20 {
                                        var v30 float32 = ((3.14159 - (v27 + 3.14159)) * (0.0 + v27))
                                        var v31 int64 = (v13 - 1)
                                        v14 -= v14
                                        _ = v30
                                        _ = v31
                                    } else {
                                        switch v17 {
                                        case "hello":
                                            _ = len(v3)
                                        default:
                                            v32 := GenS5{F0: "hello"}
                                            _ = v32
                                        }
                                        v14 = v14
                                    }
                                    switch (fi19 & (v4 - 0)) {
                                    case 274:
                                        v27 = (v27 / (v27 + (v27 * (v27 / float32(v15)))))
                                    case 42:
                                        fmt.Println(v2, v5, fi21)
                                    case 4294967295:
                                        if ((!true) || (!(v20 && true))) {
                                            switch v5 {
                                            case "foo bar":
                                                var v33 float32 = v27
                                                _ = v33
                                            case "hello":
                                                fmt.Println(v4, v14, v5)
                                            case "\n\t":
                                                if v20 {
                                                    var v34 int32 = -1
                                                    switch v5 {
                                                    case "a":
                                                        v35 := v3[(^(v14))]
                                                        _ = v35
                                                    case "\n\t":
                                                        var v36 int32 = (^(v34))
                                                        _ = v36
                                                    case "hello":
                                                        v37 := GenS2{F0: v20}
                                                        _ = v37
                                                    default:
                                                        if (!v20) {
                                                            var v38 float32 = 1e10
                                                            if (-1 == 100) {
                                                                v14 -= (v14 >> ((v18) & 63))
                                                                var v39 float32 = float32(-2.25)
                                                                switch v1 {
                                                                case 32768:
                                                                    v27 -= float32(3.14159)
                                                                case 2147483647:
                                                                    v2 += 252
                                                                default:
                                                                    v40 := GenS4{F0: v20, F1: int32(v14), F2: -1}
                                                                    _ = v40
                                                                }
                                                                _ = v39
                                                            } else {
                                                                var v41 int64 = v13
                                                                switch (fi21 ^ (v2 + int(v14))) {
                                                                case 32768:
                                                                    var v42 uint16 = 1
                                                                    _ = v42
                                                                case 0:
                                                                    v43 := GenS6{F0: (fi21 + (-(v4))), F1: (v14 - (v14 << ((v18) & 63)))}
                                                                    _ = v43
                                                                default:
                                                                    if v20 {
                                                                        v44 := GenS6{F0: (-(fi21)), F1: 194}
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            v45 := []int{}
                                                                            _ = v45
                                                                        case "":
                                                                            if (v34 > v34) {
                                                                                v46 := GenS4{F0: false, F1: (v34 % ((v34) | 1)), F2: v34}
                                                                                _ = v46
                                                                            }
                                                                        default:
                                                                            fmt.Println(v2, v1, v41)
                                                                        }
                                                                        _ = v44
                                                                    } else {
                                                                        var v47 uint32 = 4294967295
                                                                        _ = v47
                                                                    }
                                                                }
                                                                _ = v41
                                                            }
                                                            var v48 uint8 = 42
                                                            var v49 float64 = v15
                                                            _ = v38
                                                            _ = v48
                                                            _ = v49
                                                        } else {
                                                            v15 = 1.5
                                                        }
                                                    }
                                                    _ = v34
                                                } else {
                                                    var v50 int64 = (v13 | -1)
                                                    _ = v50
                                                }
                                            }
                                        } else {
                                            fmt.Println(v18, v15, v4)
                                            var v51 uint32 = (v14 + (v14 & v14))
                                            var v52 int = -922
                                            v3[v51] = (0.1 == v15)
                                            _ = v51
                                            _ = v52
                                        }
                                    default:
                                        fmt.Println(v6, fi19)
                                    }
                                }
                                var v53 int = 32768
                                _ = v27
                                _ = v53
                            }
                            var v54 int16 = (v6 ^ (v6 * v6))
                            _ = fi21
                            _ = v54
                        }
                        _ = fi19
                        _ = v20
                    }
                default:
                    switch (v4 - (v4 | v2)) {
                    case -9223372036854775808:
                        fmt.Println(v2)
                    case 582:
                        var v55 uint64 = 9
                        _ = v55
                    }
                }
                _ = v17
                _ = v18
            }
        }
        var v56 uint16 = 461
        _ = v56
    }
    var v57 string = "a"
    var v58 float64 = v15
    var v59 uint64 = 570
    v60 := GenS5{F0: ""}
    v61 := GenS7{F0: 0, F1: v14}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
