package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 bool; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint; F1 int8 }
func genFn0(p0 uint32, p1 uint) int64 {
    v7 := GenS1{F0: p1, F1: 0, F2: -295}
    p1 += p1
    var v8 uint = (p1 - 753)
    p0 = p0
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int16 = 42
    var v2 string = "\n\t"
    var v3 uint = 676
    v4 := []uint32{1, 303}
    var v5 uint16 = 937
    var v6 bool = true
    v9 := GenS1{F0: v3, F1: 255, F2: 4294967295}
    for fi10 := 0; fi10 < 7; fi10++ {
        v1 = (v1 * v1)
        var v11 float64 = -2.25
        if (-472 < 42) {
            if (v1 == v1) {
                v2 = ("\n\t" + v2)
                if len(v4) > 0 {
                    v4[0] = uint32(v5)
                }
            } else {
                v6 = v6
                if false {
                    v12 := GenS1{F0: 4294967295, F1: 44, F2: (fi10 >> ((v3) & 63))}
                    v2 = v2
                    _ = v12
                } else {
                    var v13 uint32 = 27
                    for fi14 := 0; fi14 < 1; fi14++ {
                        if true {
                            v1 = (v1 ^ 32767)
                            v15 := GenS2{F0: (2 <= 254), F1: (v1 & (-(v1))), F2: 200}
                            if ((v6 || (!(!v6))) || v6) {
                                var v16 string = v2
                                _ = v16
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = 0
                                }
                            }
                            _ = v15
                        } else {
                            var v17 uint16 = (v5 & v5)
                            var v18 int64 = 42
                            _ = v17
                            _ = v18
                        }
                        _ = fi14
                    }
                    _ = v13
                }
                fmt.Println(v5, v6, v2)
            }
            if v6 {
                for fi19 := 0; fi19 < 5; fi19++ {
                    v20 := GenS3{F0: 0, F1: 18446744073709551615}
                    if (!true) {
                        fmt.Println(v6, v2, v1)
                        if (!((-(v1)) >= v1)) {
                            v5 -= v5
                            v21 := GenS1{F0: 18446744073709551615, F1: 128, F2: (fi10 ^ fi19)}
                            _ = v21
                        }
                    }
                    _ = fi19
                    _ = v20
                }
                if v6 {
                    if false {
                        switch ("hello" + "foo bar") {
                        case "a":
                            if (false || (201 <= uint8(v5))) {
                                v22 := GenS4{F0: v3, F1: -41}
                                if v6 {
                                    var v23 uint64 = 0
                                    _ = v23
                                } else {
                                    v5 = v5
                                    var v24 uint16 = 65535
                                    _ = v24
                                }
                                var v25 uint64 = 4294967295
                                _ = v22
                                _ = v25
                            } else {
                                _ = len(v4)
                            }
                        case "hello":
                            var v26 bool = v6
                            _ = v26
                        }
                        var v27 int64 = -9223372036854775808
                        v2 = v2
                        _ = v27
                    }
                    var v28 float32 = 0.0
                    var v29 int8 = 42
                    _ = v28
                    _ = v29
                }
                var v30 int16 = v1
                _ = v30
            }
        } else {
            v11 = (float64(((1e10 - ((float32(float64(1e10)) * (0.1 * ((3.14159 * 1e10) + (0.0 / 1e10)))) * 0.1)) * 0.1)) - (0.1 / (v11 / float64(1e10))))
        }
        _ = fi10
        _ = v11
    }
    v5 -= 2
    var v31 int32 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v32 bool = (!v6)
    v31 += 42
    var v33 string = "foo bar"
    v5 = (v5 + (v5 * (v5 & (v5 * v5))))
    if v32 {
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        var v34 uint64 = 848
        for fi35 := 0; fi35 < 3; fi35++ {
            if ((((-2.25 / 1.5) * 1e10) < 0.1) && (((^(v34)) < 2) && ((v34 ^ 599) <= 18446744073709551615))) {
                v31 += (v31 << ((uint(v34)) & 63))
                if ((100 > -128) || (3.14159 >= (1e10 / -2.25))) {
                    v33 = v33
                    v36 := GenS2{F0: ((2 < 65535) || true), F1: (v1 | v1), F2: 128}
                    _ = v36
                } else {
                    if (127 < 0) {
                        var v37 uint8 = 105
                        for fi38 := 0; fi38 < 9; fi38++ {
                            var v39 uint = (v3 << ((v3) & 63))
                            _ = fi38
                            _ = v39
                        }
                        _ = v37
                    }
                }
                fmt.Println(v32, v33)
            }
            var v40 int16 = -1
            v34 = (v34 | v34)
            _ = fi35
            _ = v40
        }
        _ = v34
    } else {
        var v41 uint16 = (v5 << ((v3) & 63))
        _ = v41
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v31 = (v31 - v31)
    var v42 int32 = v31
    for fi43 := 0; fi43 < 12; fi43++ {
        for fi44 := 0; fi44 < 7; fi44++ {
            if len(v4) > 0 {
                v4[0] = 662
            }
            _ = fi44
        }
        _ = fi43
    }
    var v45 uint32 = 1
    var v46 float64 = float64((1e10 * float32(1e10)))
    for fi47 := 0; fi47 < 18; fi47++ {
        var v48 uint16 = v5
        _ = fi47
        _ = v48
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v9
}
