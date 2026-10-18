package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8; F1 uint32; F2 uint }
func genFn0(p0 uint32, p1 int, p2 int64) uint32 {
    p0 = (p0 % (((^(p0))) | 1))
    var v4 uint16 = 70
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 / ((p0) | 1))
}
func genFn1(p0 uint8, p1 int16, p2 int8) int16 {
    p0 += 129
    if true {
        v5 := GenS1{F0: "", F1: 65535}
        v6 := GenS1{F0: "\n\t", F1: 65535}
        p0 += (p0 / ((p0) | 1))
        _ = v5
        _ = v6
    } else {
        switch ("hello" + "") {
        case "a":
            v7 := GenS1{F0: ("a" + ("\n\t" + ("a" + "\n\t"))), F1: 32767}
            _ = v7
        case "hello":
            fmt.Println(p0, p1)
        default:
            p2 += (p2 + p2)
        }
        var v8 int16 = (p1 / (((p1 | p1)) | 1))
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := []int{fi9}
        fmt.Println(p0, p1)
        var v11 int32 = -143
        _ = fi9
        _ = v10
        _ = v11
    }
    p2 += (p2 & (^(p2)))
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint16 = 42
    var v3 int8 = 100
    var v12 uint8 = 189
    v12 = 250
    if false {
        v1 += (v1 >> ((18446744073709551615) & 63))
    } else {
        if false {
            fmt.Println(v3, v3, v2)
            v3 = int8(v1)
            v13 := GenS1{F0: "a", F1: v2}
            _ = v13
        } else {
            var v14 int32 = v1
            _ = v14
        }
        v15 := GenS2{F0: -9223372036854775808, F1: genFn1(v12, 0, v3), F2: (v3 >> ((369) & 63))}
        if false {
            if (42 == 42) {
                v1 += v1
                v16 := GenS3{F0: 2147483648}
                _ = v16
            } else {
                var v17 uint8 = v12
                _ = v17
            }
        } else {
            var v18 int32 = v1
            _ = v18
        }
        _ = v15
    }
    v3 = v3
    v19 := []int{}
    if (2 != 0) {
        if false {
            if (false || false) {
                if false {
                    _ = len(v19)
                    v20 := GenS4{F0: -233}
                    _ = v20
                } else {
                    v2 += (v2 | 258)
                    if false {
                        fmt.Println(v12, v1)
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v2)
                        case "hello":
                            if (!false) {
                                v12 = (v12 ^ (v12 % ((4) | 1)))
                                v3 += -128
                                v21 := GenS5{F0: (v2 == (v2 / ((v2) | 1))), F1: -571}
                                _ = v21
                            } else {
                                v1 = (v1 / ((v1) | 1))
                            }
                        default:
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        }
                        switch 586 {
                        case -2147483648:
                            _ = len(v19)
                        case -73:
                            fmt.Println(v12, v12)
                        }
                    } else {
                        fmt.Println(v2)
                    }
                    switch 32767 {
                    case -1:
                        fmt.Println(v12)
                    }
                }
                var v22 uint64 = 392
                v19 = append(v19, int(v3))
                _ = v22
            } else {
                for fi23 := 0; fi23 < 2; fi23++ {
                    v19 = append(v19, (fi23 + (fi23 | fi23)))
                    _ = fi23
                }
                v3 = (v3 / (((^(v3))) | 1))
                v24 := GenS6{F0: false}
                _ = v24
            }
            v2 = 1
        } else {
            for fi25 := 0; fi25 < 8; fi25++ {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = fi25
            }
        }
        var v26 float64 = 0.1
        v27 := GenS7{F0: v2}
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        for fi29 := 0; fi29 < 17; fi29++ {
            var v30 uint = 0
            v31 := GenS8{F0: v12, F1: 1, F2: (v30 * (^(v30)))}
            _ = fi29
            _ = v30
            _ = v31
        }
        fmt.Println(v12, v1, fi28)
        v19 = append(v19, 32767)
        _ = fi28
    }
    if true {
        switch ((("a" + ("a" + "hello")) + ("" + "\n\t")) + "a") {
        case "hello":
            fmt.Println(v1, v1, v1)
        case "":
            var v32 int = 1
            _ = v32
        }
        switch 892 {
        case 9223372036854775807:
            var v33 uint8 = 128
            _ = v33
        case 1:
            var v34 int8 = v3
            _ = v34
        case 269:
            v12 -= 2
        }
        var v35 bool = ((547 < 717) && (!false))
        _ = v35
    }
    v3 -= v3
    if (true && (false && false)) {
        v12 = (v12 + (v12 * v12))
        if len(v19) > 0 {
            v19[0] = -2147483648
        }
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    } else {
        var v36 uint16 = 42
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
}
