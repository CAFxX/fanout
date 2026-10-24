package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 int }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 uint8; F2 int }
type GenS7 struct { F0 string }
func genFn0() uint64 {
    var v6 string = ("" + "foo bar")
    var v7 float32 = (3.14159 - (1.5 / -2.25))
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 string, p1 bool, p2 uint32) int64 {
    fmt.Println(p2, p2)
    p0 = "hello"
    fmt.Println(p2, p1)
    switch "" {
    case "hello":
        v8 := GenS1{F0: p2}
        _ = v8
    default:
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := []int{fi9, (fi9 >> ((908) & 63)), (-(fi9)), (fi9 - (fi9 + fi9))}
            var v11 uint8 = 218
            p1 = p1
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}

func main() {
    var v1 int32 = -82
    var v2 float64 = 0.1
    var v3 int8 = -12
    v4 := make(map[uint32]int)
    v5 := make(map[string]int)
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int32 = 860
        v14 := GenS2{F0: (v3 | -105), F1: 32767, F2: (fi12 + fi12)}
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v2)
        var v16 float32 = -2.25
        _ = fi15
        _ = v16
    }
    switch (("a" + "hello") + "a") {
    case "hello":
        if ((^(v1)) > (v1 - v1)) {
            if true {
                var v17 int16 = 1000
                switch -47 {
                case -2147483648:
                    v18 := GenS3{F0: 0, F1: true}
                    _ = v18
                case 9223372036854775807:
                    v17 = v17
                }
                _ = v17
            }
        } else {
            if false {
                var v19 uint16 = 417
                v20 := GenS3{F0: 466, F1: false}
                _ = v19
                _ = v20
            } else {
                switch 4294967295 {
                case 42:
                    v21 := v4[227]
                    _ = v21
                case -2147483648:
                    var v22 int64 = 749
                    _ = v22
                case 1:
                    if true {
                        v1 = (v1 * -2147483648)
                        switch "hello" {
                        case "a":
                            v23 := []int{}
                            _ = v23
                        case "":
                            for fi24 := 0; fi24 < 15; fi24++ {
                                for fi25 := 0; fi25 < 7; fi25++ {
                                    v2 -= v2
                                    if ((v1 << ((41) & 63)) != v1) {
                                        v26 := v5["\n\t"]
                                        _ = v26
                                    } else {
                                        v1 += (v1 << ((18446744073709551615) & 63))
                                        fmt.Println(v3)
                                    }
                                    _ = fi25
                                }
                                v3 += v3
                                var v27 uint32 = 2147483648
                                _ = fi24
                                _ = v27
                            }
                        default:
                            for fi28 := 0; fi28 < 17; fi28++ {
                                v1 -= -2147483648
                                _ = fi28
                            }
                        }
                        v3 -= (v3 / ((v3) | 1))
                    } else {
                        v2 += (1.5 + 0.1)
                        v3 += -23
                        if (!(-2147483648 >= -340)) {
                            var v29 bool = ((0.0 < (v2 - v2)) && false)
                            v30 := GenS2{F0: (v3 % (((v3 & v3)) | 1)), F1: 0, F2: 128}
                            var v31 float32 = 3.14159
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    }
                default:
                    v1 += -1
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    v33 := GenS4{F0: genFn0()}
                    v2 = 0.1
                    v3 = 1
                    _ = fi32
                    _ = v33
                }
                v3 = (v3 / ((v3) | 1))
            }
            v2 = v2
        }
    default:
        v34 := GenS5{F0: 42}
        _ = v34
    }
    if true {
        v3 -= -1
        var v35 int64 = genFn1("a", false, 100)
        for fi36 := 0; fi36 < 6; fi36++ {
            fmt.Println(v35)
            _ = fi36
        }
        _ = v35
    } else {
        for fi37 := 0; fi37 < 11; fi37++ {
            var v38 float64 = 0.0
            var v39 uint64 = 18446744073709551615
            var v40 uint = 1
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
        v41 := GenS1{F0: 1}
        _ = v41
    }
    var v42 int8 = (v3 - (v3 * v3))
    v2 += (0.0 / (float64((0.1 / 1e10)) / v2))
    for fi43 := 0; fi43 < 15; fi43++ {
        for fi44 := 0; fi44 < 8; fi44++ {
            v45 := GenS6{F0: -9223372036854775808, F1: 128, F2: -2147483648}
            var v46 int32 = 2147483647
            _ = fi44
            _ = v45
            _ = v46
        }
        var v47 bool = (("" + ("hello" + "a")) < "foo bar")
        _ = fi43
        _ = v47
    }
    if true {
        var v48 int = 993
        _ = v48
    }
    v49 := GenS7{F0: ""}
    var v50 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v49
    _ = v5
    _ = v50
}
