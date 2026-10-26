package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 bool; F1 int16; F2 int64 }

func main() {
    var v1 uint16 = 2
    var v2 float64 = 1e10
    v3 := make(map[int]uint32)
    if false {
        v1 += 715
    }
    if true {
        switch 32767 {
        case 127:
            v2 = (v2 / v2)
        case -9223372036854775808:
            var v4 int32 = 1
            _ = v4
        case 4294967295:
            fmt.Println(v1, v2)
        }
        if false {
            fmt.Println(v2)
            fmt.Println(v2, v1)
        }
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := GenS1{F0: 2}
            if (64 <= 37) {
                var v7 int32 = -2147483648
                _ = v7
            } else {
                var v8 string = ("a" + "hello")
                _ = v8
            }
            v2 = 1e10
            _ = fi5
            _ = v6
        }
    } else {
        if ((135 > 2147483647) && false) {
            var v9 uint16 = (v1 >> ((502) & 63))
            var v10 uint32 = 1
            v11 := GenS2{F0: 32767}
            _ = v10
            _ = v11
            _ = v9
        }
    }
    var v12 int64 = -267
    v13 := GenS3{F0: -32768, F1: 716}
    if (1 <= 2) {
        v14 := GenS3{F0: 32767, F1: 2}
        v1 += 2
        v15 := GenS4{F0: 42, F1: 0}
        _ = v14
        _ = v15
    }
    v16 := GenS5{F0: false, F1: -32768, F2: v12}
    for fi17 := 0; fi17 < 2; fi17++ {
        if (false || false) {
            if false {
                var v18 uint16 = v1
                _ = v18
            }
            fmt.Println(fi17)
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v12 += (v12 << ((4294967295) & 63))
        _ = fi19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        switch ("a" + ("hello" + "hello")) {
        case "foo bar":
            if (2 < 2) {
                var v21 uint16 = (v1 >> ((1) & 63))
                if true {
                    v1 += v21
                }
                v21 = 593
                _ = v21
            } else {
                delete(v3, 0)
                switch 998 {
                case 592:
                    switch fi20 {
                    case 2147483647:
                        for fi22 := 0; fi22 < 4; fi22++ {
                            switch "a" {
                            case "a":
                                var v23 string = "hello"
                                _ = v23
                            case "foo bar":
                                switch 32768 {
                                case 41:
                                    var v24 uint16 = (v1 & (v1 ^ 1))
                                    _ = v24
                                default:
                                    v1 += v1
                                }
                            }
                            _ = fi22
                        }
                    }
                default:
                    if (18446744073709551615 != 42) {
                        fmt.Println(v1)
                        v3[fi20] = 2
                    }
                }
                switch fi20 {
                case 32767:
                    var v25 string = "foo bar"
                    _ = v25
                }
            }
        case "hello":
            switch ("foo bar" + "a") {
            case "foo bar":
                if false {
                    for fi26 := 0; fi26 < 11; fi26++ {
                        if true {
                            v1 += v1
                            v27 := []int{(^(fi26))}
                            _ = v27
                        } else {
                            var v28 int64 = v12
                            _ = v28
                        }
                        fmt.Println(v1, fi20)
                        var v29 string = ("hello" + "\n\t")
                        _ = fi26
                        _ = v29
                    }
                }
            }
        case "a":
            var v30 string = "\n\t"
            _ = v30
        }
        _ = fi20
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 int64 = (v12 >> ((255) & 63))
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
}
