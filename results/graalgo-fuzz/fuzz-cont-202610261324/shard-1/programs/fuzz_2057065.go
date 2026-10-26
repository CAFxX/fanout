package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 bool }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 int8; F1 bool }
func genFn0() int64 {
    v7 := []int{}
    switch 667 {
    case -1:
        _ = len(v7)
    }
    _ = v7
    return 930
}

func main() {
    var v1 int64 = 0
    var v2 int32 = -1
    var v3 bool = (false || (1000 == 42))
    v4 := make(map[uint32]bool)
    var v5 float64 = -2.25
    var v6 string = "hello"
    v3 = (0 > 32767)
    if (v3 || (99 >= 100)) {
        var v8 int = 1
        _ = v8
    } else {
        var v9 uint16 = 65535
        _ = len(v4)
        _ = v9
    }
    fmt.Println(v2, v5, v6)
    fmt.Println(v6, v6, v6)
    switch 9223372036854775807 {
    case 2147483647:
        v6 = v6
    case 601:
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := []int{(^(fi10)), (fi10 >> ((1) & 63)), (fi10 << ((705) & 63)), (fi10 % (((fi10 / (((fi10 | fi10)) | 1))) | 1))}
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS1{F0: 868, F1: v2, F2: v3}
    _ = len(v4)
    v13 := GenS2{F0: 42, F1: 100}
    fmt.Println(v3)
    v4[248] = false
    v2 += v2
    fmt.Println(v1)
    if (v3 && false) {
        for fi14 := 0; fi14 < 14; fi14++ {
            switch "foo bar" {
            case "foo bar":
                switch fi14 {
                case -2147483648:
                    if v3 {
                        v6 = "a"
                        fmt.Println(fi14, fi14)
                    } else {
                        for fi15 := 0; fi15 < 15; fi15++ {
                            v2 = -2147483648
                            v1 = (v1 << ((1) & 63))
                            _ = fi15
                        }
                    }
                case -9223372036854775808:
                    if (!v3) {
                        for fi16 := 0; fi16 < 18; fi16++ {
                            delete(v4, 0)
                            _ = fi16
                        }
                        _ = len(v4)
                        var v17 int16 = 0
                        _ = v17
                    } else {
                        v18 := []int{(fi14 >> ((4294967295) & 63))}
                        fmt.Println(v1, v2, v2)
                        _ = v18
                    }
                case -961:
                    v19 := []int{}
                    _ = v19
                }
            default:
                delete(v4, 4294967295)
            }
            v2 += int32(fi14)
            _ = fi14
        }
    } else {
        v20 := GenS3{F0: (v1 & -9223372036854775808)}
        fmt.Println(v1, v1)
        switch -345 {
        case 32768:
            if v3 {
                for fi21 := 0; fi21 < 9; fi21++ {
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v23 := GenS4{F0: 42, F1: false, F2: 42}
                        v24 := GenS3{F0: v1}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    var v25 int = (^(fi21))
                    v26 := GenS5{F0: genFn0()}
                    _ = fi21
                    _ = v25
                    _ = v26
                }
                var v27 int16 = 42
                v28 := GenS3{F0: v1}
                _ = v27
                _ = v28
            } else {
                v5 = (v5 / 1e10)
                var v29 int32 = v2
                var v30 int = 221
                _ = v29
                _ = v30
            }
        }
        _ = v20
    }
    v31 := []int{}
    for fi32 := 0; fi32 < 11; fi32++ {
        switch 0 {
        case 683:
            _ = len(v31)
        default:
            switch "hello" {
            case "":
                switch (^(fi32)) {
                case 4294967295:
                    for fi33 := 0; fi33 < 20; fi33++ {
                        v34 := GenS2{F0: 255, F1: 42}
                        _ = fi33
                        _ = v34
                    }
                case 32768:
                    v35 := GenS6{F0: 32767, F1: 42}
                    _ = v35
                default:
                    v1 = v1
                }
            }
        }
        _ = fi32
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 string = v6
            v3 = true
            _ = fi37
            _ = v38
        }
        var v39 uint32 = 869
        switch (("a" + v6) + v6) {
        case "\n\t":
            if v3 {
                v1 += genFn0()
                fmt.Println(v3, fi36)
            }
        case "hello":
            var v40 uint16 = 42
            _ = v40
        case "":
            switch fi36 {
            case 1:
                var v41 uint64 = 18446744073709551615
                _ = v41
            case -1:
                if (42 > v2) {
                    if v3 {
                        v5 = (v5 - v5)
                    } else {
                        var v42 bool = (4294967295 < 18446744073709551615)
                        fmt.Println(v6, v42, v2)
                        v43 := GenS7{F0: -100, F1: true}
                        _ = v42
                        _ = v43
                    }
                    v1 = -1
                }
            case 128:
                v4[(v39 - (^(v39)))] = (int(v39) < fi36)
            }
        default:
            delete(v4, 879)
        }
        _ = fi36
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
