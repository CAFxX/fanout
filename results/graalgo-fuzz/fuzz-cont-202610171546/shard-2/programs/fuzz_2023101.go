package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 string; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }
func genFn0() uint64 {
    var v6 int32 = 151
    for fi7 := 0; fi7 < 18; fi7++ {
        switch "hello" {
        case "foo bar":
            v8 := []int{fi7, (fi7 >> ((18446744073709551615) & 63)), -795, (fi7 << ((42) & 63))}
            _ = v8
        default:
            fmt.Println(v6, fi7)
        }
        _ = fi7
    }
    v9 := []int{32767, 1, -1, 9223372036854775807}
    if (true || false) {
        var v10 string = ""
        var v11 uint = 42
        _ = v10
        _ = v11
    }
    _ = v6
    _ = v9
    return 609
}

func main() {
    var v1 uint64 = 0
    var v2 float64 = 1.5
    var v3 uint = 362
    v4 := []int{}
    var v5 int = -9223372036854775808
    fmt.Println(v1)
    if true {
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(fi12, v2)
            _ = fi12
        }
    } else {
        if ((v5 >> ((0) & 63)) < (^(v5))) {
            switch "hello" {
            case "":
                for fi13 := 0; fi13 < 18; fi13++ {
                    if (float64((((1.5 * 1.5) / 3.14159) / 1.0)) >= v2) {
                        if (208 != 4) {
                            if (-1 == -2147483648) {
                                fmt.Println(v5)
                                v5 += fi13
                                var v14 bool = (("a" <= "") || false)
                                _ = v14
                            } else {
                                if false {
                                    v3 = (v3 | (v3 | v3))
                                    var v15 float64 = (v2 * (0.0 * v2))
                                    _ = v15
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = (^(fi13))
                                    }
                                }
                                for fi16 := 0; fi16 < 12; fi16++ {
                                    if (-1 == -1) {
                                        v3 = v3
                                    } else {
                                        fmt.Println(fi16, fi13, fi16)
                                        v2 -= v2
                                        if true {
                                            fmt.Println(fi13, v5, v2)
                                        } else {
                                            v17 := GenS1{F0: int(v1)}
                                            fmt.Println(fi16, v1)
                                            _ = v17
                                        }
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = (fi16 * fi13)
                                    }
                                    _ = fi16
                                }
                            }
                        }
                        if (2 >= 1) {
                            for fi18 := 0; fi18 < 18; fi18++ {
                                var v19 uint64 = v1
                                _ = fi18
                                _ = v19
                            }
                            fmt.Println(v3)
                        } else {
                            v4 = append(v4, -1)
                        }
                        v2 = v2
                    }
                    _ = fi13
                }
            case "hello":
                var v20 float32 = ((float32(0.0) - 3.14159) / (1.5 * 1.5))
                _ = v20
            case "foo bar":
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 string = (("foo bar" + "") + ("a" + "\n\t"))
                    var v23 uint16 = 255
                    var v24 uint32 = 42
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            default:
                if (-9223372036854775808 <= (v5 - (v5 << ((1) & 63)))) {
                    if false {
                        var v25 int = -49
                        _ = v25
                    }
                    fmt.Println(v2, v1, v1)
                } else {
                    v26 := GenS1{F0: v5}
                    v27 := GenS2{F0: (v5 + -2147483648)}
                    _ = v26
                    _ = v27
                }
            }
            for fi28 := 0; fi28 < 3; fi28++ {
                fmt.Println(v2)
                _ = fi28
            }
        }
        if false {
            switch "a" {
            case "a":
                switch (v5 + (v5 & v5)) {
                case 127:
                    v1 = v1
                case -9223372036854775808:
                    var v29 float64 = v2
                    _ = v29
                default:
                    var v30 int8 = -1
                    _ = v30
                }
            default:
                for fi31 := 0; fi31 < 16; fi31++ {
                    for fi32 := 0; fi32 < 9; fi32++ {
                        v1 = v1
                        var v33 float32 = (0.0 / ((3.14159 + ((((0.1 * 3.14159) * (3.14159 - 0.1)) - 3.14159) * 0.0)) / 1.0))
                        _ = fi32
                        _ = v33
                    }
                    v34 := GenS3{F0: -1, F1: "hello", F2: 258}
                    _ = fi31
                    _ = v34
                }
            }
        } else {
            var v35 uint16 = 110
            v36 := GenS1{F0: 4294967295}
            _ = v35
            _ = v36
        }
        var v37 uint8 = 4
        _ = v37
    }
    var v38 int64 = -1
    v3 = 128
    v39 := GenS4{F0: -1}
    fmt.Println(v2)
    v40 := GenS4{F0: 106}
    var v41 string = "\n\t"
    fmt.Println(v1, v3)
    fmt.Println(v38, v1)
    switch v5 {
    case -142:
        v42 := GenS5{F0: -1}
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
