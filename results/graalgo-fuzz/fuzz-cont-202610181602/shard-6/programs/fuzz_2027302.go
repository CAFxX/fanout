package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 string; F1 int8; F2 uint }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint64 }
func genFn0() int64 {
    var v6 int8 = 0
    v7 := GenS1{F0: 119, F1: (v6 + (v6 - v6))}
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v6, fi8)
        if false {
            var v9 bool = (true || false)
            _ = v9
        } else {
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := []int{}
                _ = fi10
                _ = v11
            }
            if false {
                switch "" {
                case "foo bar":
                    switch ((("a" + "a") + "\n\t") + (("hello" + "\n\t") + "foo bar")) {
                    case "a":
                        var v12 uint16 = 32767
                        _ = v12
                    case "hello":
                        v13 := GenS2{F0: ("hello" + (("hello" + "hello") + "hello")), F1: 59, F2: 683}
                        _ = v13
                    case "foo bar":
                        v14 := []int{2147483647, fi8, fi8, fi8}
                        _ = v14
                    default:
                        if true {
                            for fi15 := 0; fi15 < 3; fi15++ {
                                v6 = int8(fi15)
                                _ = fi15
                            }
                        } else {
                            switch (fi8 ^ fi8) {
                            case 2147483647:
                                v16 := []int{(fi8 << ((2) & 63)), (^(fi8)), fi8, fi8}
                                _ = v16
                            case 900:
                                fmt.Println(v6)
                            case 127:
                                v6 = (v6 + 42)
                            }
                        }
                    }
                case "":
                    var v17 float64 = float64(3.14159)
                    _ = v17
                }
                v18 := GenS3{F0: (v6 % (((v6 >> ((191) & 63))) | 1)), F1: (!("\n\t" <= ("\n\t" + "\n\t")))}
                var v19 int8 = (v6 << ((255) & 63))
                _ = v18
                _ = v19
            }
            var v20 uint8 = 2
            _ = v20
        }
        _ = fi8
    }
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint64, p1 int16) uint32 {
    fmt.Println(p0, p1)
    p1 = 42
    _ = p0
    _ = p1
    return 356
}

func main() {
    var v1 uint8 = 128
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]uint32)
    var v4 int32 = -290
    v5 := make(map[uint32]string)
    v4 += 542
    v2 = (v2 + 42)
    _ = len(v3)
    var v21 float64 = 0.0
    var v22 int16 = 452
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 int = (fi23 + (fi23 << ((1) & 63)))
        for fi25 := 0; fi25 < 16; fi25++ {
            delete(v5, 4294967295)
            fmt.Println(v21)
            _ = fi25
        }
        var v26 uint = 549
        _ = fi23
        _ = v24
        _ = v26
    }
    v4 = 0
    v27 := GenS3{F0: -128, F1: ((v1 < v1) && true)}
    switch -1 {
    case 915:
        v21 = (float64(0.1) + 0.1)
    case -738:
        var v28 uint64 = 4294967295
        _ = v28
    case 2147483647:
        fmt.Println(v4, v2, v4)
    }
    v22 = v22
    switch ("foo bar" + "a") {
    case "a":
        v29 := GenS4{F0: (("foo bar" + "") + ""), F1: 0}
        _ = v29
    }
    if true {
        for fi30 := 0; fi30 < 16; fi30++ {
            switch "" {
            case "a":
                if ((v21 + float64(3.14159)) > (v21 * -2.25)) {
                    if (!(!false)) {
                        v2 = (v2 + (v2 % (((^(v2))) | 1)))
                    }
                }
            case "hello":
                v31 := []int{(fi30 << ((2) & 63)), (-(fi30))}
                _ = v31
            default:
                v32 := v3[655]
                _ = v32
            }
            _ = fi30
        }
        var v33 string = "hello"
        v4 = int32(v2)
        _ = v33
    }
    switch 2147483647 {
    case 32767:
        var v34 float64 = 0.1
        _ = v34
    case -782:
        v35 := []int{-9223372036854775808, -2147483648}
        _ = v35
    default:
        var v36 uint32 = 0
        _ = v36
    }
    v5[genFn1(0, (^(v22)))] = ("\n\t" + ("hello" + "\n\t"))
    v1 = v1
    v37 := GenS5{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
