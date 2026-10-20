package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 string; F1 string; F2 uint32 }
func genFn0(p0 bool, p1 uint16, p2 int) uint64 {
    p0 = p0
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := []int{-489, p2}
        p1 -= (^(p1))
        var v7 uint16 = (p1 ^ p1)
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}
func genFn1(p0 uint32) uint32 {
    v8 := []int{-297}
    p0 += p0
    fmt.Println(p0)
    for fi9 := 0; fi9 < 4; fi9++ {
        switch "a" {
        case "":
            var v10 uint8 = 1
            _ = v10
        case "\n\t":
            switch fi9 {
            case 32767:
                if ("a" == (("a" + "") + "a")) {
                    for fi11 := 0; fi11 < 14; fi11++ {
                        switch (fi9 >> ((4294967295) & 63)) {
                        case 2147483647:
                            v8 = append(v8, (fi11 + fi11))
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            var v12 int = fi11
                            _ = v12
                        case -2147483648:
                            fmt.Println(fi11, fi11)
                            _ = len(v8)
                        default:
                            p0 -= (p0 | 0)
                        }
                        _ = fi11
                    }
                } else {
                    var v13 float64 = ((float64(((1.5 * -2.25) / (1.5 - -2.25))) * 0.1) + (3.14159 / 1.0))
                    var v14 string = "a"
                    switch ("" + (v14 + (v14 + ""))) {
                    case "\n\t":
                        for fi15 := 0; fi15 < 10; fi15++ {
                            var v16 int32 = 1
                            _ = fi15
                            _ = v16
                        }
                    }
                    _ = v13
                    _ = v14
                }
            case 2147483647:
                var v17 uint32 = (p0 % (((^(p0))) | 1))
                _ = v17
            default:
                v18 := GenS1{F0: fi9, F1: 741, F2: 71}
                _ = v18
            }
        }
        switch (fi9 + (-(fi9))) {
        case 1:
            p0 = (p0 / ((p0) | 1))
        default:
            if len(v8) > 0 {
                v8[0] = 2147483647
            }
        }
        if true {
            fmt.Println(fi9, p0)
            p0 = p0
            v8 = append(v8, fi9)
        } else {
            p0 = 190
            _ = len(v8)
        }
        _ = fi9
    }
    _ = p0
    _ = v8
    return 65535
}

func main() {
    v1 := []string{""}
    var v2 uint8 = 2
    var v3 uint32 = 42
    var v4 int64 = 1
    v1 = append(v1, ("a" + "a"))
    if (1.5 >= ((float32(1.5) * 0.0) * 1e10)) {
        var v19 uint = 435
        _ = v19
    } else {
        var v20 string = "a"
        _ = v20
    }
    v3 = v3
    v21 := GenS1{F0: -243, F1: 65535, F2: 128}
    v1 = append(v1, "a")
    v22 := GenS2{F0: 205, F1: ("\n\t" + "a")}
    var v23 int8 = 100
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v24 uint16 = 32767
    v25 := GenS2{F0: -460, F1: "foo bar"}
    for fi26 := 0; fi26 < 4; fi26++ {
        v4 += v4
        _ = fi26
    }
    v27 := GenS1{F0: 4294967295, F1: 1, F2: v2}
    v2 = v2
    var v28 uint8 = 127
    var v29 int32 = 42
    for fi30 := 0; fi30 < 9; fi30++ {
        if len(v1) > 0 {
            v1[0] = ("a" + "\n\t")
        }
        if true {
            for fi31 := 0; fi31 < 19; fi31++ {
                switch fi30 {
                case 4294967295:
                    v4 = v4
                case -109:
                    switch "\n\t" {
                    case "":
                        var v32 uint64 = 4294967295
                        _ = v32
                    case "hello":
                        v33 := GenS3{F0: "hello", F1: ("foo bar" + "hello"), F2: v3}
                        _ = v33
                    default:
                        switch "a" {
                        case "":
                            v1 = append(v1, "a")
                        case "\n\t":
                            switch "a" {
                            case "hello":
                                v29 = (-(v29))
                            case "a":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case "foo bar":
                            var v34 uint64 = 42
                            _ = v34
                        }
                    }
                default:
                    v35 := GenS3{F0: "foo bar", F1: "", F2: genFn1(v3)}
                    _ = v35
                }
                _ = fi31
            }
        } else {
            fmt.Println(v4, v3)
        }
        _ = fi30
    }
    switch -9223372036854775808 {
    case 2147483647:
        v36 := GenS3{F0: ("\n\t" + "\n\t"), F1: "\n\t", F2: 1}
        _ = v36
    case 517:
        v3 = genFn1(553)
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        v1 = append(v1, (("" + "\n\t") + ("foo bar" + "foo bar")))
        _ = fi37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
