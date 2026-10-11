package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
func genFn0() string {
    v6 := []int{890, 32767, 4294967295, 128}
    if ((((("" + ("\n\t" + ("\n\t" + ""))) + "") + "foo bar") + ("\n\t" + ("hello" + "\n\t"))) < "") {
        var v7 int64 = -1
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                v6 = append(v6, fi9)
                v10 := GenS1{F0: 1, F1: 42}
                _ = fi9
                _ = v10
            }
            _ = len(v6)
            _ = fi8
        }
        var v11 bool = (!false)
        _ = v11
        _ = v7
    }
    _ = v6
    return (("hello" + "") + "\n\t")
}

func main() {
    var v1 int16 = 1000
    v2 := []uint32{2147483648}
    var v3 uint64 = 4294967295
    v4 := make(map[string]int)
    v5 := []int{-164}
    var v12 int = 127
    var v13 string = "\n\t"
    var v14 float64 = ((-2.25 * 3.14159) - 0.1)
    if (!false) {
        v15 := GenS2{F0: (v13 + "a")}
        if false {
            for fi16 := 0; fi16 < 10; fi16++ {
                v1 = 32767
                _ = fi16
            }
        }
        _ = v15
    } else {
        if (v14 != v14) {
            v5 = append(v5, -556)
            switch v13 {
            case "hello":
                var v17 uint32 = 65535
                _ = v17
            case "foo bar":
                switch v12 {
                case 9223372036854775807:
                    v1 = (v1 ^ v1)
                case 32767:
                    switch (v13 + v13) {
                    case "hello":
                        _ = len(v5)
                    case "\n\t":
                        for fi18 := 0; fi18 < 5; fi18++ {
                            for fi19 := 0; fi19 < 6; fi19++ {
                                delete(v4, v13)
                                _ = fi19
                            }
                            for fi20 := 0; fi20 < 3; fi20++ {
                                if true {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v21 bool = (1 < 32767)
                                    _ = v21
                                } else {
                                    fmt.Println(v1)
                                    _ = len(v4)
                                }
                                _ = fi20
                            }
                            _ = fi18
                        }
                    }
                case -821:
                    for fi22 := 0; fi22 < 13; fi22++ {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            var v24 string = v13
                            v5 = append(v5, fi23)
                            _ = fi23
                            _ = v24
                        }
                        if ((^(v12)) == -9223372036854775808) {
                            for fi25 := 0; fi25 < 14; fi25++ {
                                var v26 uint = 572
                                _ = fi25
                                _ = v26
                            }
                            fmt.Println(v3, v3)
                        }
                        var v27 uint8 = 127
                        _ = fi22
                        _ = v27
                    }
                }
            default:
                if (!false) {
                    v14 = 3.14159
                } else {
                    switch v12 {
                    case 2147483647:
                        v14 = ((v14 + (v14 * (v14 * 0.1))) * ((float64(-2.25) / ((((v14 - -2.25) / 3.14159) / float64(((1.5 + ((1e10 + 1e10) / 1e10)) + 0.0))) - -2.25)) * ((float64(((((((float32(v14) / 3.14159) / 1e10) + (0.0 / 1.0)) * (0.0 / float32(v14))) - float32(v14)) * 3.14159)) + -2.25) / (v14 + float64(1.5)))))
                    case 231:
                        var v28 int8 = 28
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 12; fi29++ {
                        var v30 uint16 = 1
                        var v31 uint16 = uint16(v3)
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    var v32 int64 = 9223372036854775807
                    _ = v32
                }
            }
            fmt.Println(v13)
        }
    }
    var v33 uint16 = 0
    fmt.Println(v33, v1)
    for fi34 := 0; fi34 < 8; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            v36 := GenS3{F0: true}
            _ = fi35
            _ = v36
        }
        var v37 int = (fi34 ^ fi34)
        _ = fi34
        _ = v37
    }
    var v38 uint16 = v33
    v39 := GenS2{F0: ("foo bar" + "\n\t")}
    v13 = v13
    if false {
        var v40 int = 0
        var v41 int32 = -498
        _ = v40
        _ = v41
    }
    var v42 float64 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
}
