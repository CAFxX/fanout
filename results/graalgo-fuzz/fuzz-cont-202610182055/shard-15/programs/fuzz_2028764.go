package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 bool; F1 int32; F2 uint64 }
type GenS5 struct { F0 string; F1 uint16; F2 int32 }
type GenS6 struct { F0 uint; F1 int8; F2 int }
func genFn0(p0 uint64, p1 int8) int64 {
    v6 := GenS1{F0: 1, F1: 2147483648, F2: 122}
    if (42 < 128) {
        p1 += (-(p1))
        var v7 bool = true
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return -1
}
func genFn1(p0 int64) int {
    p0 -= p0
    if false {
        var v8 uint64 = 42
        var v9 int16 = -1
        _ = v8
        _ = v9
    }
    if (182 <= 896) {
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := []int{fi10, fi10}
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    return 32767
}

func main() {
    var v1 uint16 = 143
    var v2 bool = (-47 <= 108)
    var v3 float64 = (1e10 - 3.14159)
    var v4 int = 625
    var v5 float64 = (-2.25 - v3)
    v12 := []int{42, v4, (v4 & v4)}
    v4 -= int(v1)
    if true {
        v13 := GenS1{F0: 639, F1: 65535, F2: v1}
        _ = v13
    } else {
        v14 := GenS2{F0: 127}
        if v2 {
            var v15 uint16 = (v1 >> ((0) & 63))
            v16 := GenS3{F0: 18446744073709551615, F1: 42}
            v15 = 0
            _ = v15
            _ = v16
        }
        _ = v14
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            if v2 {
                if (1.5 < 0.1) {
                    if true {
                        _ = len(v12)
                        var v19 uint = 65535
                        _ = v19
                    }
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
            }
            if v2 {
                fmt.Println(v1, fi17)
                switch (fi17 | (^(fi18))) {
                case 273:
                    v5 = -2.25
                case -1:
                    switch (fi17 ^ (^(v4))) {
                    case 1:
                        if (!v2) {
                            v4 -= fi18
                        } else {
                            v1 = (v1 - (v1 / (((^(v1))) | 1)))
                            var v20 uint16 = (v1 | v1)
                            _ = v20
                        }
                    default:
                        v21 := GenS1{F0: 9223372036854775807, F1: 757, F2: 81}
                        _ = v21
                    }
                case -368:
                    v22 := GenS4{F0: false, F1: -1, F2: 42}
                    _ = v22
                }
            } else {
                switch (-(fi17)) {
                case 446:
                    fmt.Println(v1, v3)
                case 1:
                    v3 -= v5
                case -189:
                    v23 := GenS5{F0: "foo bar", F1: 0, F2: -1}
                    _ = v23
                }
                v2 = (!false)
                v24 := GenS2{F0: 1}
                _ = v24
            }
            _ = fi18
        }
        if (v1 <= 0) {
            if (false || false) {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                switch fi17 {
                case 32767:
                    v25 := GenS3{F0: 803, F1: 0}
                    _ = v25
                }
                v5 += (0.1 - v3)
            }
        }
        _ = fi17
    }
    v3 = v5
    v4 += (v4 | (v4 ^ (v4 << ((0) & 63))))
    var v26 string = "a"
    fmt.Println(v5)
    for fi27 := 0; fi27 < 20; fi27++ {
        if v2 {
            _ = len(v12)
            for fi28 := 0; fi28 < 19; fi28++ {
                v29 := GenS6{F0: 427, F1: -1, F2: -2147483648}
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 2; fi30++ {
                v31 := GenS1{F0: 42, F1: 42, F2: (^(v1))}
                if (585 > 845) {
                    switch v26 {
                    case "hello":
                        v4 = int(v1)
                    case "\n\t":
                        var v32 int64 = 9223372036854775807
                        _ = v32
                    }
                    switch 2147483647 {
                    case -888:
                        var v33 int64 = -9223372036854775808
                        _ = v33
                    case 0:
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    case -875:
                        v34 := GenS3{F0: 508, F1: 544}
                        _ = v34
                    default:
                        var v35 uint64 = 18446744073709551615
                        _ = v35
                    }
                } else {
                    var v36 uint = 2
                    _ = v36
                }
                _ = fi30
                _ = v31
            }
        }
        var v37 uint = 2
        _ = fi27
        _ = v37
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        for fi39 := 0; fi39 < 8; fi39++ {
            var v40 int32 = 42
            if (771 != 0) {
                v4 = (fi38 * fi38)
            } else {
                fmt.Println(v3, v4, v40)
                if ((!(!v2)) && ((18446744073709551615 <= 416) && ((v1 / (((v1 + v1)) | 1)) < v1))) {
                    var v41 uint16 = v1
                    if (v26 >= "hello") {
                        v12 = append(v12, fi39)
                    } else {
                        var v42 uint16 = (v1 * v41)
                        if (!false) {
                            if ((-2.25 <= 0.0) || (((1 == 1) && (v2 && false)) || true)) {
                                var v43 uint32 = 1
                                _ = v43
                            } else {
                                v44 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: 207}
                                if ((!(!true)) || false) {
                                    _ = len(v12)
                                    v42 -= (v1 % ((0) | 1))
                                    fmt.Println(v26)
                                } else {
                                    v2 = false
                                }
                                if len(v12) > 0 {
                                    v12[0] = (fi39 << ((255) & 63))
                                }
                                var v45 uint16 = v1
                                _ = v44
                                _ = v45
                            }
                        } else {
                            fmt.Println(v2, v42, v4)
                        }
                        _ = v42
                    }
                    var v46 uint = 1
                    _ = v41
                    _ = v46
                }
            }
            _ = fi39
            _ = v40
        }
        switch v4 {
        case 127:
            var v47 string = (v26 + (("hello" + ("foo bar" + v26)) + "foo bar"))
            _ = v47
        case 539:
            _ = len(v12)
        case 581:
            switch genFn1(42) {
            case 4294967295:
                var v48 int = v4
                _ = v48
            default:
                fmt.Println(fi38, v4)
            }
        }
        switch "hello" {
        case "\n\t":
            var v49 uint32 = 131
            _ = v49
        case "foo bar":
            var v50 uint = 0
            _ = v50
        }
        _ = fi38
    }
    fmt.Println(v2, v5, v3)
    _ = len(v12)
    var v51 string = v26
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
