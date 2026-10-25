package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS3 struct { F0 int8; F1 bool; F2 uint64 }
func genFn0(p0 uint8, p1 int8) uint32 {
    v7 := []int{int(p1)}
    var v8 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 692
}

func main() {
    var v1 uint64 = 0
    var v2 int32 = 2147483647
    var v3 int = -101
    var v4 int32 = (v2 >> ((0) & 63))
    var v5 uint32 = 703
    var v6 uint16 = 874
    v9 := []int{(v3 | (-(v3))), 298, v3}
    if len(v9) > 0 {
        v9[0] = (-(v3))
    }
    var v10 int64 = int64(v3)
    var v11 int32 = (-(v4))
    for fi12 := 0; fi12 < 20; fi12++ {
        switch "\n\t" {
        case "foo bar":
            var v13 uint64 = 2
            _ = v13
        case "hello":
            _ = len(v9)
        }
        v5 += 65535
        if len(v9) > 0 {
            v9[0] = v3
        }
        _ = fi12
    }
    if false {
        v14 := GenS1{F0: v3}
        var v15 float64 = 1e10
        _ = v14
        _ = v15
    }
    _ = len(v9)
    if (1.5 > float64((0.0 + 0.0))) {
        switch v3 {
        case 9223372036854775807:
            v10 = 0
        default:
            fmt.Println(v2, v2)
        }
    } else {
        var v16 uint16 = 255
        if (!false) {
            _ = len(v9)
            v17 := GenS1{F0: (v3 + v3)}
            _ = v17
        }
        _ = v16
    }
    v18 := GenS1{F0: int(v10)}
    v11 -= v2
    switch ("hello" + "foo bar") {
    case "hello":
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 6; fi20++ {
                fmt.Println(v1, v10, fi19)
                fmt.Println(v3, v11, v6)
                if (v6 == v6) {
                    if (720 != 1) {
                        var v21 int32 = (v11 % (((v11 >> ((65535) & 63))) | 1))
                        _ = v21
                    } else {
                        v2 = -29
                    }
                    v22 := GenS1{F0: v3}
                    _ = v22
                } else {
                    var v23 float64 = ((float64(-2.25) / -2.25) * float64((((1e10 / 1e10) * float32(0.1)) + 0.1)))
                    _ = v23
                }
                _ = fi20
            }
            if false {
                v1 -= v1
                v9 = append(v9, int(v5))
            } else {
                v9 = append(v9, fi19)
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := GenS1{F0: -1}
                    _ = fi24
                    _ = v25
                }
                v26 := GenS1{F0: (-(fi19))}
                _ = v26
            }
            _ = fi19
        }
    default:
        for fi27 := 0; fi27 < 6; fi27++ {
            for fi28 := 0; fi28 < 4; fi28++ {
                if (-2.25 != 1.5) {
                    var v29 int = (fi27 * (^(fi28)))
                    v1 -= v1
                    _ = v29
                }
                if true {
                    if (((false || (2147483647 == (v5 + (^(v5))))) || false) || (!(812 >= 1))) {
                        var v30 int64 = (v10 * (v10 | (-(v10))))
                        var v31 int = 9223372036854775807
                        _ = len(v9)
                        _ = v30
                        _ = v31
                    } else {
                        v10 = v10
                        v32 := GenS2{F0: v5, F1: 42, F2: v2}
                        v1 = v1
                        _ = v32
                    }
                    var v33 float64 = (0.1 / 1.0)
                    _ = v33
                } else {
                    var v34 uint8 = 0
                    _ = v34
                }
                if false {
                    if (false && (!false)) {
                        if true {
                            if false {
                                v10 = (v10 + -558)
                                var v35 int32 = v11
                                _ = len(v9)
                                _ = v35
                            }
                            v36 := GenS1{F0: (v3 ^ 32767)}
                            _ = v36
                        } else {
                            switch "\n\t" {
                            case "":
                                var v37 float32 = (((((1.5 + ((1.5 * 1e10) * 0.0)) * 1e10) / 3.14159) + (1e10 - 1e10)) * 0.0)
                                _ = v37
                            }
                            switch (v3 - v3) {
                            case 2147483647:
                                fmt.Println(v5)
                            case -943:
                                v38 := GenS2{F0: (^(v5)), F1: 0, F2: (v2 ^ 605)}
                                v9 = append(v9, (fi27 - v3))
                                switch (fi27 + v3) {
                                case 487:
                                    v6 = v6
                                case 2147483647:
                                    if false {
                                        v5 = v5
                                    }
                                case 654:
                                    v10 -= v10
                                }
                                _ = v38
                            case 32767:
                                var v39 float64 = 0.1
                                _ = v39
                            default:
                                v40 := GenS2{F0: (v5 % ((357) | 1)), F1: 73, F2: (^(v2))}
                                _ = v40
                            }
                        }
                        var v41 int = v3
                        _ = v41
                    } else {
                        _ = len(v9)
                        v42 := GenS1{F0: (-(fi27))}
                        _ = v42
                    }
                }
                _ = fi28
            }
            v43 := GenS3{F0: 84, F1: false, F2: (v1 * (v1 >> ((2) & 63)))}
            var v44 int64 = v10
            _ = fi27
            _ = v43
            _ = v44
        }
    }
    fmt.Println(v2, v1, v5)
    for fi45 := 0; fi45 < 1; fi45++ {
        v10 -= 0
        var v46 uint8 = 128
        _ = fi45
        _ = v46
    }
    _ = len(v9)
    var v47 int64 = (v10 >> ((42) & 63))
    v1 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v9
}
