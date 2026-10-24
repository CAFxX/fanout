package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint; F2 int16 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 32767
    var v3 uint8 = 182
    var v4 uint8 = (v3 + (v3 << ((101) & 63)))
    var v5 bool = (!true)
    var v6 int8 = -128
    v5 = false
    v4 += 1
    if (v5 && v5) {
        fmt.Println(v3)
        var v7 string = "\n\t"
        _ = v7
    } else {
        if v5 {
            v8 := GenS1{F0: -9223372036854775808, F1: v5}
            _ = v8
        }
        v9 := v1[770]
        var v10 uint64 = uint64(v3)
        _ = v10
        _ = v9
    }
    v5 = (-9223372036854775808 >= -2147483648)
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := []int{fi11}
        var v13 int = fi11
        v14 := GenS2{F0: 0}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 uint64 = 18446744073709551615
    v16 := []int{}
    if true {
        var v17 uint16 = 2
        if (((v17 != 2) && true) || (-375 < 758)) {
            v2 -= 42
            v18 := v1[uint32(v3)]
            _ = v18
        } else {
            v15 = 488
            var v19 uint32 = 2147483648
            v15 -= v15
            _ = v19
        }
        _ = v17
    } else {
        v20 := GenS3{F0: 398}
        _ = v20
    }
    if len(v16) > 0 {
        v16[0] = 541
    }
    v1[2] = v5
    switch -635 {
    case 32768:
        if v5 {
            var v21 uint8 = 255
            if (-9223372036854775808 <= 0) {
                for fi22 := 0; fi22 < 8; fi22++ {
                    switch "\n\t" {
                    case "a":
                        v23 := v1[42]
                        _ = v23
                    }
                    _ = fi22
                }
                if v5 {
                    if (((1.5 / (-2.25 * (0.0 - 0.1))) + float32(((-2.25 / float64(-2.25)) + (3.14159 / 0.1)))) > 1e10) {
                        for fi24 := 0; fi24 < 20; fi24++ {
                            for fi25 := 0; fi25 < 18; fi25++ {
                                fmt.Println(v4)
                                if v5 {
                                    v3 = (v4 & v3)
                                    _ = len(v1)
                                }
                                _ = fi25
                            }
                            delete(v1, 0)
                            v5 = v5
                            _ = fi24
                        }
                        for fi26 := 0; fi26 < 15; fi26++ {
                            fmt.Println(v3, v15)
                            v27 := GenS4{F0: 4294967295, F1: 4294967295, F2: 42}
                            v28 := GenS2{F0: 42}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                } else {
                    if (v5 || (65535 >= 4294967295)) {
                        v29 := v1[1]
                        _ = v29
                    }
                }
            }
            var v30 uint64 = v15
            _ = v21
            _ = v30
        } else {
            var v31 string = "a"
            for fi32 := 0; fi32 < 6; fi32++ {
                switch "" {
                case "foo bar":
                    fmt.Println(v2, v15)
                }
                v15 -= (v15 + 1)
                _ = fi32
            }
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            _ = v31
        }
    case 2147483647:
        for fi33 := 0; fi33 < 7; fi33++ {
            if (!v5) {
                v15 += v15
                switch "foo bar" {
                case "hello":
                    var v34 int16 = 1
                    _ = v34
                }
            } else {
                v16 = append(v16, fi33)
                var v35 float32 = -2.25
                var v36 bool = false
                _ = v35
                _ = v36
            }
            var v37 string = "hello"
            for fi38 := 0; fi38 < 12; fi38++ {
                v39 := GenS5{F0: int16(v15)}
                v15 = v15
                switch "hello" {
                case "a":
                    fmt.Println(v6, v2, v37)
                case "hello":
                    v16 = append(v16, fi33)
                case "\n\t":
                    v3 = v3
                    var v40 int8 = (v6 % (((v6 << ((4294967295) & 63))) | 1))
                    v2 = -86
                    _ = v40
                }
                var v41 int32 = -634
                var v42 uint8 = 51
                _ = fi38
                _ = v39
                _ = v41
                _ = v42
            }
            _ = fi33
            _ = v37
        }
    case 4294967295:
        var v43 uint = 65535
        _ = v43
    }
    if v5 {
        var v44 string = "foo bar"
        var v45 int32 = 2147483647
        v45 -= v45
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 5; fi46++ {
        if false {
            var v47 int = (-(fi46))
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v2, v47, v4)
            case "hello":
                switch -1 {
                case 42:
                    fmt.Println(v2)
                default:
                    v48 := GenS3{F0: v15}
                    _ = v48
                }
            case "foo bar":
                for fi49 := 0; fi49 < 13; fi49++ {
                    var v50 uint16 = 2
                    _ = fi49
                    _ = v50
                }
            }
            _ = v47
        }
        var v51 uint64 = v15
        v6 += 127
        _ = fi46
        _ = v51
    }
    delete(v1, 2147483648)
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
