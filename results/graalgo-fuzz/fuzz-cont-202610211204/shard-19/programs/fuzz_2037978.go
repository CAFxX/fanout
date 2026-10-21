package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 string; F1 uint64; F2 uint }
type GenS4 struct { F0 uint; F1 uint }
func genFn0(p0 bool, p1 string, p2 int32) uint {
    v5 := GenS1{F0: 437, F1: 4294967295, F2: int16(p2)}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint = 18446744073709551615
        v8 := GenS1{F0: (^(v7)), F1: 18446744073709551615, F2: 1000}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        p1 = ((("\n\t" + ("foo bar" + p1)) + p1) + "\n\t")
        fmt.Println(p2, p2)
        switch p1 {
        case "":
            if (!(-77 == -1)) {
                for fi10 := 0; fi10 < 17; fi10++ {
                    p2 -= p2
                    p0 = (!(float64(0.0) == (0.0 / 1.0)))
                    _ = fi10
                }
                var v11 int32 = p2
                switch 128 {
                case -1:
                    var v12 uint8 = 42
                    _ = v12
                }
                _ = v11
            } else {
                p1 = "\n\t"
            }
        case "\n\t":
            v13 := GenS1{F0: 65535, F1: 1, F2: 0}
            _ = v13
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 871
}
func genFn1(p0 int8, p1 uint32, p2 uint64) uint32 {
    v14 := []int{-9223372036854775808, 38, 1}
    fmt.Println(p2, p1, p0)
    switch "hello" {
    case "a":
        fmt.Println(p0, p2, p2)
    case "foo bar":
        if false {
            switch 42 {
            case 32767:
                var v15 bool = (((^(p1)) >= (p1 / ((p1) | 1))) && (-1 != (p0 & p0)))
                _ = v15
            case 2147483647:
                switch 486 {
                case 0:
                    if false {
                        for fi16 := 0; fi16 < 11; fi16++ {
                            p2 -= p2
                            p1 -= (^(p1))
                            _ = fi16
                        }
                    }
                }
            default:
                p2 -= 880
            }
        } else {
            if false {
                if true {
                    p0 = 42
                    p2 = (p2 << ((81) & 63))
                    switch -400 {
                    case 0:
                        var v17 uint16 = uint16(p2)
                        _ = v17
                    case 42:
                        if (-1 < 32767) {
                            _ = len(v14)
                            var v18 string = ("foo bar" + "\n\t")
                            _ = v18
                        } else {
                            fmt.Println(p0)
                            if len(v14) > 0 {
                                v14[0] = -212
                            }
                        }
                    }
                } else {
                    if len(v14) > 0 {
                        v14[0] = 32768
                    }
                }
                if (65535 != 42) {
                    var v19 float32 = 3.14159
                    _ = v19
                }
                p2 += p2
            }
        }
    case "hello":
        var v20 uint64 = p2
        _ = v20
    default:
        if ((p2 > (p2 - p2)) && true) {
            var v21 uint = 295
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 string = "a"
        var v24 float32 = 0.0
        fmt.Println(p0)
        _ = fi22
        _ = v23
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (p1 + p1)
}

func main() {
    var v1 uint32 = 997
    v2 := []uint32{804, (v1 + v1), (v1 | (^(v1)))}
    v3 := make(map[int]int)
    var v4 uint = 2
    v4 -= v4
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := v3[fi25]
        var v27 string = ""
        _ = fi25
        _ = v26
        _ = v27
    }
    v1 = v1
    v1 += v1
    v28 := GenS2{F0: 381, F1: 874}
    var v29 int32 = 0
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 10; fi31++ {
            switch "hello" {
            case "a":
                if (42 >= (^(fi31))) {
                    var v32 bool = true
                    switch (fi30 ^ fi30) {
                    case -2147483648:
                        v33 := GenS2{F0: 18446744073709551615, F1: 2}
                        _ = v33
                    }
                    _ = v32
                } else {
                    delete(v3, fi31)
                }
            case "foo bar":
                fmt.Println(v4, v29)
            case "hello":
                switch fi30 {
                case 127:
                    var v34 bool = (((-62 > 0) && ((v1 & (v1 | genFn1(-128, uint32(v4), 293))) == v1)) || false)
                    _ = v34
                case -476:
                    var v35 uint8 = 2
                    _ = v35
                case 1:
                    switch "a" {
                    case "foo bar":
                        v2 = append(v2, v1)
                    }
                default:
                    switch (fi31 + (fi31 << ((0) & 63))) {
                    case 474:
                        if ("\n\t" == ("" + ((("hello" + "foo bar") + "hello") + "foo bar"))) {
                            var v36 int16 = -515
                            v37 := GenS2{F0: 42, F1: 0}
                            var v38 uint8 = 250
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                    }
                }
            default:
                v4 = v4
            }
            v39 := GenS2{F0: 4294967295, F1: 42}
            _ = fi31
            _ = v39
        }
        if (!false) {
            v40 := GenS3{F0: "", F1: 1, F2: 1}
            v4 -= v4
            v4 = 2
            _ = v40
        } else {
            for fi41 := 0; fi41 < 2; fi41++ {
                var v42 int32 = 0
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi41
                _ = v42
            }
            var v43 uint8 = 47
            _ = v43
        }
        _ = fi30
    }
    var v44 uint16 = 318
    var v45 int64 = 1
    var v46 uint8 = 48
    v44 -= 2
    v45 = -535
    v2 = append(v2, v1)
    var v47 int32 = -1
    var v48 string = ("" + ("foo bar" + "foo bar"))
    if ((^(v46)) > v46) {
        v29 = -2147483648
        v49 := GenS3{F0: ((v48 + (v48 + "foo bar")) + v48), F1: 4294967295, F2: 42}
        _ = v49
    } else {
        v50 := GenS4{F0: (^(v4)), F1: v4}
        v48 = v48
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
