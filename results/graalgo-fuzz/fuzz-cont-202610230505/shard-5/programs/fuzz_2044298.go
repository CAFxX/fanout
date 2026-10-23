package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint64; F2 int }
type GenS7 struct { F0 int64; F1 uint16; F2 uint32 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]uint32)
    v4 := []uint32{42, (^(v1)), v1, (v1 << ((1) & 63))}
    v5 := []uint32{0, v1}
    v6 := make(map[uint32]string)
    if (("" == ("" + "a")) || false) {
        for fi7 := 0; fi7 < 3; fi7++ {
            v1 += v1
            v2 = v2
            _ = fi7
        }
        delete(v3, 1)
    } else {
        v2 = (^(v2))
        fmt.Println(v2, v1)
        for fi8 := 0; fi8 < 11; fi8++ {
            switch "\n\t" {
            case "foo bar":
                v9 := v3[(fi8 - fi8)]
                _ = v9
            case "":
                switch (^(fi8)) {
                case -2147483648:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case "\n\t":
                var v10 string = ("foo bar" + ("hello" + "a"))
                _ = v10
            default:
                var v11 int32 = -895
                _ = v11
            }
            var v12 int16 = 0
            for fi13 := 0; fi13 < 17; fi13++ {
                switch fi13 {
                case -1:
                    v14 := GenS1{F0: (v12 - (v12 << ((1) & 63))), F1: 72}
                    _ = v14
                }
                var v15 int32 = int32(fi13)
                _ = fi13
                _ = v15
            }
            _ = fi8
            _ = v12
        }
    }
    fmt.Println(v2, v2)
    var v16 uint = 844
    v17 := GenS1{F0: 32767, F1: 1}
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int8 = -1
        var v20 float64 = 1.5
        _ = fi18
        _ = v19
        _ = v20
    }
    if (2147483647 > -1) {
        v16 = v16
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 uint64 = 42
            v23 := GenS1{F0: 42, F1: 42}
            _ = fi21
            _ = v22
            _ = v23
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        if false {
            var v25 uint64 = 1
            _ = v25
        } else {
            var v26 int = fi24
            var v27 uint = v16
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    if true {
        v28 := GenS2{F0: v16}
        for fi29 := 0; fi29 < 1; fi29++ {
            fmt.Println(v2)
            var v30 int16 = -32768
            for fi31 := 0; fi31 < 2; fi31++ {
                v32 := GenS1{F0: (v30 & (^(v30))), F1: 145}
                v16 += (v16 & 42)
                if (false && (((false && (!true)) || true) && true)) {
                    v1 += v1
                    v1 = (v1 & (v1 | (^(v1))))
                    switch fi29 {
                    case -9223372036854775808:
                        v33 := GenS1{F0: v30, F1: 42}
                        _ = v33
                    case 0:
                        v30 += v30
                    default:
                        if true {
                            v34 := GenS3{F0: (v2 & 496)}
                            var v35 uint16 = 42
                            _ = v34
                            _ = v35
                        }
                    }
                } else {
                    v30 -= v30
                    v2 = -1
                }
                _ = fi31
                _ = v32
            }
            _ = fi29
            _ = v30
        }
        _ = v28
    } else {
        v36 := GenS4{F0: 1, F1: 70}
        for fi37 := 0; fi37 < 19; fi37++ {
            v1 = 2
            _ = fi37
        }
        var v38 string = ("a" + ("hello" + (((((("a" + ("a" + "foo bar")) + "a") + ("hello" + "hello")) + "hello") + ("a" + ("\n\t" + "\n\t"))) + ("foo bar" + (("hello" + "a") + ("\n\t" + ("foo bar" + "hello")))))))
        _ = v36
        _ = v38
    }
    v39 := v6[(v1 >> (((^(v16))) & 63))]
    _ = len(v6)
    v39 = "\n\t"
    for fi40 := 0; fi40 < 15; fi40++ {
        v39 = v39
        switch v39 {
        case "foo bar":
            if false {
                v16 = (v16 ^ 1)
                v41 := GenS5{F0: -822, F1: v2}
                _ = v41
            } else {
                v42 := GenS5{F0: 42, F1: v2}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v42
            }
        case "a":
            v43 := GenS6{F0: 42, F1: 1, F2: 0}
            _ = v43
        case "\n\t":
            var v44 uint = v16
            _ = v44
        }
        if false {
            for fi45 := 0; fi45 < 13; fi45++ {
                switch "\n\t" {
                case "a":
                    var v46 int64 = (v2 << (((v16 ^ (^(v16)))) & 63))
                    _ = v46
                case "hello":
                    v39 = "\n\t"
                }
                _ = fi45
            }
        }
        _ = fi40
    }
    v1 = (v1 * (^(v1)))
    if len(v5) > 0 {
        v5[0] = v1
    }
    v47 := GenS4{F0: 0, F1: 42}
    v48 := GenS7{F0: v2, F1: 42, F2: (^(v1))}
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
