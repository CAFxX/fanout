package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint32; F2 int }
type GenS4 struct { F0 int; F1 int64; F2 uint8 }
type GenS5 struct { F0 uint64; F1 string; F2 uint64 }
type GenS6 struct { F0 int32; F1 uint64 }
func genFn0(p0 int16, p1 int16) uint8 {
    var v5 uint = 18446744073709551615
    var v6 int16 = p1
    for fi7 := 0; fi7 < 11; fi7++ {
        if true {
            if true {
                if true {
                    v8 := []int{}
                    fmt.Println(fi7, v5)
                    _ = v8
                } else {
                    v9 := []int{(fi7 | (fi7 & (fi7 % ((32767) | 1))))}
                    _ = v9
                }
                fmt.Println(v6, fi7)
            } else {
                for fi10 := 0; fi10 < 11; fi10++ {
                    if false {
                        var v11 uint32 = 1
                        var v12 float64 = (3.14159 - 1e10)
                        v13 := []int{fi7, (-(fi10)), 128}
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        fmt.Println(p1, fi7, fi10)
                    }
                    p0 -= p0
                    if true {
                        if true {
                            var v14 int32 = -874
                            fmt.Println(fi10, v5)
                            p0 -= (^(v6))
                            _ = v14
                        }
                        v15 := GenS1{F0: 65535}
                        _ = v15
                    } else {
                        if ((32767 == 2) && false) {
                            switch "\n\t" {
                            case "a":
                                var v16 int64 = 365
                                _ = v16
                            }
                        } else {
                            var v17 uint8 = 1
                            if (42 <= 563) {
                                v18 := []int{fi7, (fi7 % ((int(v6)) | 1))}
                                switch "\n\t" {
                                case "foo bar":
                                    var v19 bool = (1 <= 1)
                                    _ = v19
                                case "a":
                                    v20 := GenS1{F0: 32767}
                                    _ = v20
                                case "":
                                    v21 := GenS1{F0: 1}
                                    _ = v21
                                default:
                                    var v22 float32 = (((1.5 - 0.0) / (((1e10 - -2.25) + -2.25) + -2.25)) / 0.1)
                                    _ = v22
                                }
                                _ = v18
                            } else {
                                if (!(v17 == 127)) {
                                    fmt.Println(v6)
                                    v23 := GenS2{F0: v5, F1: 42}
                                    v24 := GenS1{F0: 0}
                                    var v25 int16 = v6
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    fmt.Println(v5)
                                    var v26 uint64 = 405
                                    v26 += (^(v26))
                                    var v27 string = "a"
                                    _ = v26
                                    _ = v27
                                }
                                if ((v17 * (v17 - 72)) <= v17) {
                                    v28 := []int{(fi7 % ((fi7) | 1)), -868, fi10, (fi10 - fi7)}
                                    _ = v28
                                }
                                v17 = 42
                            }
                            _ = v17
                        }
                    }
                    fmt.Println(p1, p0)
                    _ = fi10
                }
            }
            v6 -= (p1 / ((int16(fi7)) | 1))
        } else {
            v5 += v5
            var v29 int32 = -608
            _ = v29
        }
        _ = fi7
    }
    var v30 int16 = p0
    if false {
        switch (("hello" + "a") + "a") {
        case "a":
            fmt.Println(v5, v5)
        }
    } else {
        var v31 int8 = -108
        for fi32 := 0; fi32 < 3; fi32++ {
            v33 := []int{(fi32 << ((v5) & 63)), fi32, (^(fi32))}
            _ = fi32
            _ = v33
        }
        if ((18446744073709551615 == 42) && false) {
            fmt.Println(v30)
        } else {
            var v34 uint64 = 442
            _ = v34
        }
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v30
    _ = v5
    _ = v6
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 856
    var v3 uint = 0
    var v4 uint8 = 199
    var v35 uint64 = 705
    if false {
        if false {
            fmt.Println(v4, v2, v4)
        }
        switch ("hello" + "a") {
        case "foo bar":
            var v36 uint8 = (v4 | genFn0(1000, -1))
            _ = v36
        case "\n\t":
            v37 := GenS1{F0: uint16(v35)}
            _ = v37
        default:
            v3 -= (v3 << (((v3 >> ((v3) & 63))) & 63))
        }
    }
    v38 := GenS3{F0: -1, F1: 2147483647, F2: 4294967295}
    if false {
        for fi39 := 0; fi39 < 4; fi39++ {
            fmt.Println(v35)
            delete(v1, 0)
            _ = fi39
        }
    } else {
        var v40 uint8 = (v4 >> ((389) & 63))
        var v41 int64 = -495
        _ = v40
        _ = v41
    }
    v42 := GenS4{F0: 1, F1: 1, F2: (^(v4))}
    v43 := GenS5{F0: (^(v35)), F1: ("hello" + ("a" + "\n\t")), F2: (v35 & (v35 * (v35 + (v35 << (((^(v3))) & 63)))))}
    if true {
        var v44 int64 = 57
        for fi45 := 0; fi45 < 1; fi45++ {
            v3 += v3
            var v46 int64 = (v44 % ((v44) | 1))
            if (((786 == -2147483648) || (v4 == genFn0(771, -32768))) || (-9223372036854775808 <= (^(v44)))) {
                for fi47 := 0; fi47 < 19; fi47++ {
                    var v48 int16 = -503
                    _ = fi47
                    _ = v48
                }
            } else {
                var v49 uint = v3
                for fi50 := 0; fi50 < 2; fi50++ {
                    var v51 uint = v49
                    _ = fi50
                    _ = v51
                }
                _ = v49
            }
            _ = fi45
            _ = v46
        }
        switch ("foo bar" + "") {
        case "foo bar":
            v52 := GenS4{F0: -75, F1: v44, F2: (v4 >> (((v3 ^ v3)) & 63))}
            _ = v52
        case "":
            for fi53 := 0; fi53 < 11; fi53++ {
                v54 := GenS6{F0: 850, F1: 527}
                fmt.Println(v3)
                _ = fi53
                _ = v54
            }
        case "a":
            if (!false) {
                _ = len(v1)
            } else {
                v35 = (v35 % ((4294967295) | 1))
            }
        default:
            v55 := v1[4294967295]
            _ = v55
        }
        _ = v44
    }
    v56 := []int{9223372036854775807}
    v3 = (v3 * v3)
    v1[959] = int(v4)
    fmt.Println(v4, v3, v4)
    var v57 bool = ("" < "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v56
    _ = v57
}
