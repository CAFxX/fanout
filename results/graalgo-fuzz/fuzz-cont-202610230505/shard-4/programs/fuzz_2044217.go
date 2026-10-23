package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int16; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint16; F1 int; F2 int }
type GenS6 struct { F0 uint16; F1 string; F2 int }
func genFn0(p0 int) uint8 {
    p0 -= (p0 % (((p0 ^ p0)) | 1))
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 -= -2147483648
        var v7 string = ""
        var v8 string = v7
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: int16(p0)}
    v10 := GenS2{F0: 1, F1: "foo bar"}
    _ = p0
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 127
    var v2 uint = 538
    var v3 uint16 = 585
    var v4 uint64 = 42
    v5 := make(map[int]uint32)
    v11 := []int{42}
    v1 = 120
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint = (v2 >> ((4294967295) & 63))
        v2 = (^(v2))
        var v14 uint = 18446744073709551615
        _ = fi12
        _ = v13
        _ = v14
    }
    v11 = append(v11, -1)
    var v15 uint32 = 903
    if (!false) {
        v16 := GenS3{F0: 42, F1: 100, F2: (^(v1))}
        for fi17 := 0; fi17 < 19; fi17++ {
            if false {
                for fi18 := 0; fi18 < 19; fi18++ {
                    var v19 uint8 = v1
                    if (9223372036854775807 != 642) {
                        if (true && (!((!((uint8(v15) != v19) || true)) || false))) {
                            fmt.Println(v3, v19, fi17)
                            delete(v5, (fi17 | fi17))
                            if (672 > uint64(v2)) {
                                v1 = v1
                                fmt.Println(v4, v19)
                                fmt.Println(fi17)
                                v20 := GenS1{F0: -258}
                                _ = v20
                            } else {
                                v15 += (^(v15))
                                if false {
                                    fmt.Println(v19, v2, v15)
                                    v3 = v3
                                } else {
                                    v3 -= v3
                                    switch "hello" {
                                    case "foo bar":
                                        v21 := GenS4{F0: uint64(v2), F1: ""}
                                        _ = v21
                                    default:
                                        if len(v11) > 0 {
                                            v11[0] = fi17
                                        }
                                    }
                                }
                                v22 := GenS5{F0: v3, F1: fi17, F2: 42}
                                _ = v22
                            }
                        }
                        var v23 float32 = 1.5
                        _ = v23
                    }
                    if (v3 == (v3 % ((2) | 1))) {
                        if false {
                            if true {
                                var v24 float32 = ((-2.25 / (3.14159 - -2.25)) - 1.5)
                                _ = v24
                            } else {
                                if true {
                                    var v25 uint16 = 863
                                    _ = v25
                                }
                                var v26 float32 = (float32(0.0) * 3.14159)
                                _ = v26
                            }
                            var v27 uint32 = 1
                            var v28 uint64 = 973
                            _ = v27
                            _ = v28
                        }
                    } else {
                        if (3.14159 < 3.14159) {
                            var v29 int = -1
                            switch "hello" {
                            case "a":
                                if false {
                                    var v30 string = ("\n\t" + "")
                                    v4 = (v4 ^ (v4 | uint64(v1)))
                                    var v31 int16 = 1
                                    _ = v30
                                    _ = v31
                                } else {
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    v19 -= (v19 << (((v2 % (((v2 / ((430) | 1))) | 1))) & 63))
                                    v29 = (fi17 * fi17)
                                }
                                var v32 bool = ((("hello" + ("a" + "foo bar")) <= "foo bar") && ((false && (v19 > (v1 % ((genFn0((fi18 & fi18))) | 1)))) && false))
                                v2 = v2
                                _ = v32
                            default:
                                v1 = v1
                            }
                            var v33 uint = v2
                            _ = v29
                            _ = v33
                        }
                        v4 += 2
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    }
                    _ = fi18
                    _ = v19
                }
                v34 := GenS2{F0: (^(v3)), F1: ((("" + "\n\t") + "a") + ("" + "\n\t"))}
                _ = v34
            } else {
                for fi35 := 0; fi35 < 12; fi35++ {
                    var v36 float64 = (-2.25 - 1.5)
                    _ = fi35
                    _ = v36
                }
                v37 := GenS6{F0: 1, F1: (("a" + "") + "foo bar"), F2: -9223372036854775808}
                _ = v37
            }
            _ = fi17
        }
        if ("a" > "") {
            v3 = v3
            var v38 bool = (-2147483648 <= 42)
            _ = v38
        } else {
            fmt.Println(v1)
        }
        _ = v16
    } else {
        var v39 uint16 = 2
        v3 += (v39 >> ((906) & 63))
        for fi40 := 0; fi40 < 1; fi40++ {
            v4 -= 2
            var v41 int32 = 42
            switch (fi40 - (fi40 % ((fi40) | 1))) {
            case -180:
                var v42 int8 = 100
                _ = v42
            case 899:
                var v43 int = (-(fi40))
                _ = v43
            case 42:
                fmt.Println(fi40, v2)
            }
            _ = fi40
            _ = v41
        }
        _ = v39
    }
    switch 42 {
    case 557:
        fmt.Println(v2, v1, v15)
    case -1:
        for fi44 := 0; fi44 < 1; fi44++ {
            for fi45 := 0; fi45 < 19; fi45++ {
                v2 = (v2 | (v2 % (((v2 << ((578) & 63))) | 1)))
                _ = fi45
            }
            var v46 uint64 = v4
            _ = fi44
            _ = v46
        }
    default:
        fmt.Println(v4)
    }
    fmt.Println(v2)
    fmt.Println(v3)
    var v47 uint64 = (v4 & (v4 % (((^(v4))) | 1)))
    var v48 uint16 = uint16(v1)
    v5[-552] = (v15 + v15)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
