package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint16; F2 uint16 }
func genFn0(p0 int, p1 uint8) uint64 {
    var v5 int16 = -26
    var v6 uint = 65535
    v7 := GenS1{F0: 0, F1: 127, F2: -1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 350
}
func genFn1(p0 uint32) uint32 {
    fmt.Println(p0)
    p0 = (^(p0))
    _ = p0
    return (p0 << ((65535) & 63))
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    v3 := []int{}
    v4 := make(map[uint32]bool)
    v1 -= v1
    for fi8 := 0; fi8 < 13; fi8++ {
        if ("a" > ("foo bar" + ("foo bar" + ("foo bar" + ("" + ""))))) {
            for fi9 := 0; fi9 < 15; fi9++ {
                v1 = v1
                v10 := GenS1{F0: v1, F1: 42, F2: 9223372036854775807}
                fmt.Println(fi9)
                _ = fi9
                _ = v10
            }
            v1 += int32(fi8)
            v3 = append(v3, (fi8 | (fi8 >> ((2) & 63))))
        } else {
            v11 := GenS1{F0: 1, F1: 42, F2: -1}
            fmt.Println(v1, fi8)
            v1 += v1
            _ = v11
        }
        _ = fi8
    }
    var v12 uint64 = genFn0(-190, 42)
    if (!false) {
        if (v1 < -939) {
            v13 := GenS1{F0: 1, F1: -113, F2: 308}
            _ = v13
        } else {
            v1 -= v1
            var v14 string = ("\n\t" + "hello")
            switch -440 {
            case 42:
                v14 = v14
            case -143:
                for fi15 := 0; fi15 < 8; fi15++ {
                    switch "\n\t" {
                    case "\n\t":
                        switch fi15 {
                        case -2147483648:
                            v14 = v14
                        }
                    case "foo bar":
                        if (!false) {
                            v16 := GenS2{F0: 128}
                            if false {
                                for fi17 := 0; fi17 < 18; fi17++ {
                                    v18 := GenS3{F0: 0}
                                    var v19 string = "hello"
                                    if true {
                                        var v20 int64 = 1
                                        _ = v20
                                    }
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            }
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = v16
                        }
                    default:
                        v12 -= (^(v12))
                    }
                    _ = fi15
                }
            }
            _ = v14
        }
        var v21 int = -934
        _ = v21
    } else {
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS4{F0: 9223372036854775807, F1: false}
            v24 := GenS1{F0: (^(v1)), F1: -18, F2: 42}
            var v25 uint8 = 42
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        delete(v4, 2147483647)
    }
    v12 -= genFn0(127, 128)
    switch (("a" + "") + ("" + (("hello" + ((("\n\t" + ("\n\t" + ("a" + "a"))) + ("\n\t" + ("\n\t" + ("hello" + "hello")))) + "hello")) + ""))) {
    case "hello":
        switch "hello" {
        case "hello":
            for fi26 := 0; fi26 < 1; fi26++ {
                v1 = -137
                var v27 uint16 = 1
                _ = fi26
                _ = v27
            }
        case "foo bar":
            v28 := GenS2{F0: 17}
            _ = v28
        case "a":
            if (3.14159 != 1.5) {
                v12 = 4294967295
                fmt.Println(v1, v1)
            }
        default:
            if true {
                if false {
                    var v29 int = 1
                    _ = v29
                }
                var v30 uint8 = 255
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 string = "\n\t"
                    _ = fi31
                    _ = v32
                }
                _ = v30
            }
        }
    default:
        var v33 float64 = (((0.1 * float64(0.1)) * float64(-2.25)) - ((-2.25 / 0.1) * 1.5))
        _ = v33
    }
    v1 -= v1
    var v34 int = 0
    var v35 int = v34
    switch v35 {
    case -916:
        v12 = v12
    case -232:
        v12 += (v12 >> ((21) & 63))
    }
    v36 := GenS5{F0: -128}
    v35 -= v34
    for fi37 := 0; fi37 < 10; fi37++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v35)
        _ = fi37
    }
    v2[genFn1(4294967295)] = (v34 - (v34 & 32767))
    var v38 int16 = -1
    v39 := GenS6{F0: (v34 >> ((0) & 63)), F1: 65535, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v39
    _ = v4
}
