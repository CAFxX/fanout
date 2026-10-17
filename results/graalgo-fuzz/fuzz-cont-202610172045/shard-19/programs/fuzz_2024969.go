package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 uint64, p1 int) uint16 {
    var v7 uint64 = p0
    v8 := []int{}
    p1 -= p1
    for fi9 := 0; fi9 < 13; fi9++ {
        p1 -= fi9
        switch 0 {
        case 932:
            var v10 int32 = 0
            _ = v10
        case 4294967295:
            v8 = append(v8, -177)
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(p0, fi9, fi9)
            _ = fi11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint64 = 0
            if len(v8) > 0 {
                v8[0] = (fi12 << ((18446744073709551615) & 63))
            }
            if len(v8) > 0 {
                v8[0] = p1
            }
            fmt.Println(p0, v7, fi13)
            fmt.Println(v7)
            if false {
                fmt.Println(v7, fi12)
            }
            v8 = append(v8, 4294967295)
            var v15 float32 = (1.5 - 1.5)
            fmt.Println(v7, v15)
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            switch (fi16 % (((p1 + (fi12 & fi12))) | 1)) {
            case 9223372036854775807:
                p1 -= (p1 % ((fi12) | 1))
            }
            if (!(!true)) {
                v7 = 42
                var v17 string = ("hello" + "hello")
                _ = v17
            }
            _ = fi16
        }
        v18 := GenS1{F0: 1, F1: 42, F2: 0}
        _ = fi12
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint = 970
    v2 := make(map[uint32]bool)
    var v3 uint8 = 255
    v4 := []int{}
    var v5 bool = ((2147483647 < 4294967295) && true)
    v6 := make(map[uint32]int)
    v1 += (v1 * 2)
    switch "hello" {
    case "a":
        v2[422] = v5
    case "foo bar":
        if (uint16(v1) != 32767) {
            v1 = (v1 % ((v1) | 1))
            var v19 string = "hello"
            _ = v19
        } else {
            var v20 uint64 = 0
            _ = v20
        }
    case "":
        var v21 bool = (!(373 != 255))
        _ = v21
    default:
        var v22 int = -387
        _ = v22
    }
    var v23 float32 = (3.14159 * 1.5)
    v23 = v23
    v3 = (v3 ^ (v3 * (v3 % ((v3) | 1))))
    var v24 int32 = 42
    v25 := GenS1{F0: 42, F1: 18446744073709551615, F2: v1}
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 uint16 = genFn0(755, (fi26 & (fi26 * fi26)))
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := GenS2{F0: (!v5), F1: (v1 ^ v1), F2: (!((v3 >= 25) || v5))}
        _ = fi28
        _ = v29
    }
    var v30 float64 = -2.25
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 int64 = -1
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 float64 = 3.14159
        _ = fi33
        _ = v34
    }
    if (!v5) {
        switch 4294967295 {
        case 2147483647:
            var v35 uint8 = (v3 * v3)
            _ = v35
        }
        v24 = (^(v24))
        var v36 float32 = (float32(v30) - (v23 / v23))
        _ = v36
    } else {
        switch -9223372036854775808 {
        case -2147483648:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 1:
            if ((("" + (("a" + "") + "a")) > ("\n\t" + "")) || (1000 < -195)) {
                var v37 uint = v1
                v38 := GenS1{F0: 4294967295, F1: 0, F2: v37}
                _ = v37
                _ = v38
            }
        case 127:
            for fi39 := 0; fi39 < 15; fi39++ {
                delete(v2, 1)
                if v5 {
                    _ = len(v6)
                    if len(v4) > 0 {
                        v4[0] = (fi39 % ((1) | 1))
                    }
                    if v5 {
                        if v5 {
                            var v40 int = fi39
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v5 = v5
                            _ = v40
                        }
                        var v41 uint = v1
                        var v42 float64 = v30
                        _ = v41
                        _ = v42
                    } else {
                        for fi43 := 0; fi43 < 19; fi43++ {
                            fmt.Println(v3, fi43, v3)
                            _ = fi43
                        }
                        if v5 {
                            var v44 uint16 = 32767
                            fmt.Println(v3)
                            _ = v44
                        }
                    }
                } else {
                    v23 = float32((1.5 * (v30 + v30)))
                    switch (fi39 - 42) {
                    case 0:
                        _ = len(v4)
                    case -9223372036854775808:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        if v5 {
                            switch ("\n\t" + "foo bar") {
                            case "foo bar":
                                for fi45 := 0; fi45 < 14; fi45++ {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v46 int16 = 1
                                    var v47 string = ""
                                    _ = fi45
                                    _ = v46
                                    _ = v47
                                }
                            case "\n\t":
                                if false {
                                    v3 -= (v3 << ((0) & 63))
                                    var v48 int64 = 0
                                    _ = v48
                                } else {
                                    v2[574] = v5
                                    for fi49 := 0; fi49 < 19; fi49++ {
                                        fmt.Println(v5, v23)
                                        _ = fi49
                                    }
                                    v5 = (!true)
                                }
                            }
                        } else {
                            if true {
                                for fi50 := 0; fi50 < 3; fi50++ {
                                    v1 -= (^(v1))
                                    v51 := GenS3{F0: "\n\t", F1: 651}
                                    var v52 uint16 = 255
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                }
                            } else {
                                if (true && ("foo bar" < "hello")) {
                                    var v53 uint = (v1 + v1)
                                    _ = v53
                                }
                                if len(v4) > 0 {
                                    v4[0] = 9223372036854775807
                                }
                                if false {
                                    v2[65535] = (32767 <= 255)
                                }
                            }
                            fmt.Println(v23)
                        }
                    }
                }
                _ = fi39
            }
        default:
            fmt.Println(v23, v23)
        }
        v5 = v5
    }
    var v54 int16 = int16(v24)
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
