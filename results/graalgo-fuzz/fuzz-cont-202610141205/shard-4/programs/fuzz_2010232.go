package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int64; F2 string }
type GenS3 struct { F0 uint64; F1 int64; F2 uint64 }

func main() {
    var v1 int32 = 2147483647
    var v2 uint16 = 1
    var v3 int16 = 1
    v4 := make(map[uint32]int)
    var v5 uint = 976
    v6 := make(map[string]int)
    switch 180 {
    case 1:
        delete(v6, "a")
    }
    switch 127 {
    case -234:
        var v7 uint16 = (v2 >> ((v5) & 63))
        _ = v7
    case 1:
        fmt.Println(v3, v1, v1)
    case 91:
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                _ = len(v6)
                if true {
                    v3 = v3
                    var v10 uint16 = 778
                    if (128 <= 32767) {
                        var v11 uint32 = 2147483647
                        v12 := GenS1{F0: 0}
                        if ((float32(1.5) / 3.14159) != 0.0) {
                            var v13 int8 = 42
                            if (v2 >= (^(v2))) {
                                fmt.Println(fi8)
                                if false {
                                    if ("\n\t" >= ("\n\t" + "a")) {
                                        v13 -= v13
                                        v14 := []int{fi9}
                                        v1 += int32(v5)
                                        _ = v14
                                    }
                                    v2 = (v2 | v2)
                                    v6[""] = (fi9 | (fi9 - (fi8 << (((v5 >> (((v5 / (((^(v5))) | 1))) & 63))) & 63))))
                                    v15 := GenS2{F0: v11, F1: -1, F2: ((("" + "hello") + "a") + "foo bar")}
                                    _ = v15
                                } else {
                                    v3 = 887
                                    fmt.Println(v5, fi8)
                                    var v16 string = ("\n\t" + "a")
                                    _ = v16
                                }
                                var v17 uint32 = (v11 ^ (v11 % ((v11) | 1)))
                                _ = v17
                            }
                            _ = v13
                        }
                        _ = v11
                        _ = v12
                    } else {
                        v18 := []int{int(v1), -2147483648, (fi9 + fi8), fi9}
                        _ = v18
                    }
                    _ = v10
                } else {
                    v2 = v2
                }
                _ = fi9
            }
            if true {
                var v19 bool = (fi8 > fi8)
                var v20 int64 = 0
                _ = v19
                _ = v20
            }
            _ = fi8
        }
    }
    fmt.Println(v3, v3, v1)
    var v21 string = (("\n\t" + (("hello" + "foo bar") + ("" + "\n\t"))) + "foo bar")
    v22 := GenS2{F0: 65535, F1: 0, F2: v21}
    var v23 bool = (!false)
    _ = len(v6)
    v21 = "a"
    switch 794 {
    case -343:
        if (!true) {
            var v24 uint8 = uint8(v5)
            v25 := []int{42, -1, 754}
            _ = v24
            _ = v25
        } else {
            fmt.Println(v5, v23)
            fmt.Println(v3, v1)
        }
    case 127:
        if (v23 && (true && (0 < 18446744073709551615))) {
            if v23 {
                for fi26 := 0; fi26 < 2; fi26++ {
                    for fi27 := 0; fi27 < 3; fi27++ {
                        v28 := GenS1{F0: 9223372036854775807}
                        fmt.Println(v1, v2, fi27)
                        var v29 float64 = ((float64(-2.25) * -2.25) / 3.14159)
                        if v23 {
                            if (fi26 == fi26) {
                                v30 := []int{(fi27 & (fi27 >> (((v5 ^ (^(v5)))) & 63))), 128, fi26, fi27}
                                v1 = 2147483647
                                _ = v30
                            }
                            var v31 uint16 = (v2 * (^(v2)))
                            v29 = (3.14159 - (v29 + v29))
                            if (!(3.14159 == (float32((-2.25 / v29)) * 3.14159))) {
                                var v32 int64 = 9223372036854775807
                                _ = v32
                            }
                            _ = v31
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    _ = fi26
                }
            } else {
                switch v21 {
                case "foo bar":
                    var v33 float32 = 0.1
                    _ = v33
                default:
                    for fi34 := 0; fi34 < 17; fi34++ {
                        var v35 bool = (2 >= 4294967295)
                        v36 := []int{int(v1), fi34, fi34}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
            }
        }
    case 776:
        if (-60 == 32767) {
            var v37 bool = ((!(-2.25 != 0.0)) || (!v23))
            if v23 {
                var v38 int64 = -910
                var v39 uint64 = 878
                _ = v38
                _ = v39
            }
            _ = v37
        } else {
            v40 := GenS1{F0: 9223372036854775807}
            _ = v40
        }
    default:
        v41 := GenS3{F0: 4294967295, F1: 42, F2: 82}
        _ = v41
    }
    if (0.0 >= 3.14159) {
        switch -125 {
        case 5:
            if v23 {
                v3 = (-(v3))
            }
        }
    }
    var v42 uint8 = 0
    for fi43 := 0; fi43 < 8; fi43++ {
        var v44 uint32 = 4294967295
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
