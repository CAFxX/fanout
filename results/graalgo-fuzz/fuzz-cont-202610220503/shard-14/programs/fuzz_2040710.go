package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 string; F2 uint8 }
type GenS4 struct { F0 uint64; F1 int32; F2 int }
func genFn0(p0 int16) uint64 {
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := []int{(fi6 << ((4294967295) & 63))}
        var v8 uint32 = 269
        var v9 float64 = -2.25
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    p0 += p0
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint = 614
        _ = fi10
        _ = v11
    }
    _ = p0
    return 42
}

func main() {
    var v1 uint32 = 0
    var v2 uint = 65535
    v3 := []uint32{v1, v1}
    var v4 uint32 = (v1 * v1)
    var v5 uint64 = 921
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint32 = 2147483648
        switch "a" {
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = (v13 + (^(v13)))
            }
        case "hello":
            fmt.Println(v2, v2)
        default:
            v1 += v13
        }
        v1 += v4
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 255}
    _ = len(v3)
    fmt.Println(v5, v5)
    var v15 int = -782
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v16 uint = 218
    fmt.Println(v1, v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "" {
    case "foo bar":
        var v17 float64 = -2.25
        _ = v17
    default:
        if false {
            _ = len(v3)
            var v18 uint16 = 32767
            fmt.Println(v2)
            _ = v18
        } else {
            if false {
                fmt.Println(v4, v16)
            }
            fmt.Println(v15)
        }
    }
    if (!false) {
        if false {
            for fi19 := 0; fi19 < 6; fi19++ {
                if true {
                    var v20 bool = (false || true)
                    switch "\n\t" {
                    case "hello":
                        if true {
                            for fi21 := 0; fi21 < 9; fi21++ {
                                var v22 string = (("" + "") + "foo bar")
                                v2 -= 1
                                _ = fi21
                                _ = v22
                            }
                        }
                    case "\n\t":
                        for fi23 := 0; fi23 < 10; fi23++ {
                            v24 := GenS2{F0: uint8(v5)}
                            var v25 int32 = 2147483647
                            var v26 bool = ((127 >= 128) || (v20 || v20))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v16 += v2
                            v3 = append(v3, (v1 % ((676) | 1)))
                            _ = len(v3)
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    default:
                        v20 = v20
                    }
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                    _ = v20
                } else {
                    v27 := GenS3{F0: 65535, F1: ("\n\t" + (("\n\t" + "\n\t") + ("hello" + "foo bar"))), F2: 2}
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    v1 -= 1
                    _ = v27
                }
                if (false || (1 == 255)) {
                    switch 2147483647 {
                    case 717:
                        var v28 uint16 = 42
                        _ = v28
                    default:
                        for fi29 := 0; fi29 < 8; fi29++ {
                            var v30 int8 = 42
                            _ = fi29
                            _ = v30
                        }
                    }
                    fmt.Println(fi19, v16, v15)
                }
                _ = fi19
            }
        }
        if true {
            fmt.Println(v2, v15)
            for fi31 := 0; fi31 < 11; fi31++ {
                v2 -= (^(v2))
                if false {
                    v5 += (v5 >> (((v2 % ((v2) | 1))) & 63))
                    var v32 int32 = 401
                    _ = v32
                }
                v15 += (v15 + v15)
                _ = fi31
            }
            var v33 bool = true
            _ = v33
        } else {
            _ = len(v3)
        }
    } else {
        v34 := GenS3{F0: 65535, F1: "foo bar", F2: 2}
        for fi35 := 0; fi35 < 16; fi35++ {
            v36 := GenS1{F0: 127}
            _ = fi35
            _ = v36
        }
        switch "hello" {
        case "":
            v3 = append(v3, (v1 % ((2147483647) | 1)))
        case "hello":
            v3 = append(v3, 2147483647)
        case "\n\t":
            v1 -= (v1 * v4)
        default:
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 bool = true
                if ((^(v4)) != 2147483647) {
                    if (0.0 >= 1e10) {
                        v3 = append(v3, (v1 & v1))
                        v3 = append(v3, v1)
                        switch "a" {
                        case "":
                            var v39 bool = (false || false)
                            _ = v39
                        case "foo bar":
                            if v38 {
                                fmt.Println(v1)
                            } else {
                                v2 -= 0
                                if v38 {
                                    for fi40 := 0; fi40 < 3; fi40++ {
                                        v16 += 42
                                        _ = fi40
                                    }
                                    v3 = append(v3, (^(v1)))
                                    v3 = append(v3, 2147483647)
                                }
                                fmt.Println(v1)
                            }
                        case "hello":
                            v41 := GenS3{F0: 32767, F1: "hello", F2: uint8(v16)}
                            _ = v41
                        default:
                            var v42 int16 = 32767
                            _ = v42
                        }
                    } else {
                        for fi43 := 0; fi43 < 18; fi43++ {
                            fmt.Println(fi43, v16)
                            _ = fi43
                        }
                        for fi44 := 0; fi44 < 20; fi44++ {
                            if false {
                                v38 = (1 >= 1)
                                v45 := GenS4{F0: v5, F1: 0, F2: fi37}
                                _ = v45
                            }
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                            v38 = (!v38)
                            _ = fi44
                        }
                    }
                    v46 := GenS4{F0: v5, F1: 42, F2: fi37}
                    v2 -= v2
                    _ = v46
                } else {
                    v4 = (v1 & (v1 + (v1 / ((v4) | 1))))
                    _ = len(v3)
                    var v47 int16 = 1000
                    _ = v47
                }
                fmt.Println(v16, v16, v1)
                _ = fi37
                _ = v38
            }
        }
        _ = v34
    }
    for fi48 := 0; fi48 < 13; fi48++ {
        var v49 uint32 = v1
        v5 += v5
        _ = fi48
        _ = v49
    }
    v5 -= v5
    fmt.Println(v2, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
