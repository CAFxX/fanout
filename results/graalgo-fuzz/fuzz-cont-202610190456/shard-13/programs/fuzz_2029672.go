package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint64; F1 uint; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int; F2 uint16 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 uint16; F2 int }
func genFn0(p0 bool, p1 bool) uint64 {
    p0 = true
    if (0 > 863) {
        p0 = p1
    } else {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := []int{(-(fi6))}
            var v8 uint = 65535
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    return 862
}
func genFn1(p0 uint8, p1 uint16) uint32 {
    var v9 uint16 = p1
    var v10 bool = (!((p1 - (p1 ^ (v9 & p1))) != v9))
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int8 = 127
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 int = 127
    v2 := []string{"hello"}
    v3 := []uint32{}
    var v4 int64 = -9223372036854775808
    var v5 bool = true
    v4 += 21
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 int16 = 0
    if false {
        v4 -= v4
    } else {
        v14 := GenS1{F0: -75}
        v15 := GenS1{F0: 1}
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: v1, F1: 0}
    fmt.Println(v5)
    var v17 int8 = 1
    var v18 uint16 = 1
    if (!v5) {
        for fi19 := 0; fi19 < 11; fi19++ {
            v18 -= v18
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 int64 = v4
                if v5 {
                    fmt.Println(v5, v21)
                }
                if true {
                    var v22 float32 = 1.5
                    _ = len(v3)
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 uint16 = v18
            _ = fi23
            _ = v24
        }
        if len(v2) > 0 {
            v2[0] = ("\n\t" + "")
        }
    }
    v25 := GenS3{F0: 4294967295, F1: 718, F2: (v1 << ((42) & 63))}
    for fi26 := 0; fi26 < 9; fi26++ {
        v1 -= fi26
        _ = fi26
    }
    if v5 {
        var v27 int8 = v17
        _ = v27
    } else {
        v17 += v17
    }
    if v5 {
        if ((!true) && (v5 || true)) {
            var v28 int16 = 568
            if true {
                v4 = (v4 >> ((583) & 63))
                var v29 uint32 = 1
                v30 := GenS4{F0: (v17 >> ((2) & 63))}
                _ = v29
                _ = v30
            } else {
                fmt.Println(v4, v5, v4)
                var v31 bool = (((613 < 42) && v5) || v5)
                _ = v31
            }
            if (1.5 < 3.14159) {
                fmt.Println(v5)
            }
            _ = v28
        }
        v3 = append(v3, 4294967295)
    } else {
        var v32 float32 = (0.0 * ((1e10 + 0.0) - -2.25))
        _ = v32
    }
    if (v5 || false) {
        if (492 <= 0) {
            switch "foo bar" {
            case "foo bar":
                var v33 int64 = (v4 >> ((18446744073709551615) & 63))
                _ = v33
            case "\n\t":
                v34 := GenS4{F0: v17}
                _ = v34
            case "hello":
                if ((!((!true) && (2 >= 2))) && v5) {
                    for fi35 := 0; fi35 < 15; fi35++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(v4, v5)
                        v1 -= (^(v1))
                        _ = fi35
                    }
                    if v5 {
                        var v36 uint16 = (v18 * v18)
                        _ = v36
                    } else {
                        if v5 {
                            for fi37 := 0; fi37 < 14; fi37++ {
                                for fi38 := 0; fi38 < 8; fi38++ {
                                    var v39 int32 = 42
                                    _ = fi38
                                    _ = v39
                                }
                                v40 := GenS5{F0: 2, F1: int(v17), F2: (v18 & v18)}
                                for fi41 := 0; fi41 < 18; fi41++ {
                                    v5 = (-2.25 <= 0.1)
                                    _ = fi41
                                }
                                _ = fi37
                                _ = v40
                            }
                            var v42 uint16 = (^(v18))
                            _ = v42
                        } else {
                            switch (-(v1)) {
                            case -1:
                                _ = len(v2)
                            case 950:
                                if len(v2) > 0 {
                                    v2[0] = ((("\n\t" + "foo bar") + "a") + (("hello" + "hello") + "a"))
                                }
                            case -469:
                                _ = len(v3)
                            }
                            v43 := GenS6{F0: false, F1: 488}
                            _ = v43
                        }
                    }
                } else {
                    switch "a" {
                    case "":
                        fmt.Println(v18, v4)
                    case "foo bar":
                        v18 = v18
                    }
                    var v44 uint32 = 117
                    _ = v44
                }
            }
            v17 -= v17
            fmt.Println(v4, v18)
        } else {
            v45 := GenS7{F0: 128}
            var v46 int8 = (-(v17))
            _ = v45
            _ = v46
        }
    } else {
        v47 := GenS8{F0: -2147483648}
        v3 = append(v3, 771)
        _ = v47
    }
    var v48 int64 = int64(v17)
    if false {
        v49 := GenS9{F0: (v17 ^ int8(v1)), F1: v18, F2: v1}
        fmt.Println(v13)
        var v50 string = ""
        _ = v49
        _ = v50
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v48
    _ = v5
}
