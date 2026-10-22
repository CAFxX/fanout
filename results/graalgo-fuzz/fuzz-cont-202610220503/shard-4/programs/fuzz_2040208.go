package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint16 = 0
    var v2 int64 = 0
    var v3 float64 = 1e10
    v4 := make(map[string]int)
    var v5 uint = 18446744073709551615
    switch "" {
    case "\n\t":
        v2 += v2
    case "hello":
        var v6 uint64 = 42
        _ = v6
    case "":
        var v7 float32 = 1e10
        _ = v7
    default:
        if true {
            v8 := []int{4294967295, 9223372036854775807, 0}
            fmt.Println(v1, v3, v5)
            v1 = v1
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint64 = 0
        _ = fi9
        _ = v10
    }
    if false {
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := []int{}
            if true {
                if (!true) {
                    for fi13 := 0; fi13 < 18; fi13++ {
                        fmt.Println(v5, v1)
                        fmt.Println(v5)
                        _ = fi13
                    }
                }
                v14 := v4[("foo bar" + "\n\t")]
                _ = v14
            } else {
                v2 += (-(v2))
                switch "" {
                case "foo bar":
                    var v15 float32 = 0.1
                    _ = v15
                case "a":
                    v3 += v3
                }
                v1 = v1
            }
            _ = fi11
            _ = v12
        }
        v3 += v3
    } else {
        var v16 float32 = -2.25
        for fi17 := 0; fi17 < 13; fi17++ {
            if false {
                v18 := GenS1{F0: (v2 ^ (v2 >> (((v5 / ((v5) | 1))) & 63))), F1: (65535 != v1)}
                _ = v18
            }
            var v19 string = ("" + "hello")
            var v20 uint64 = 2
            _ = fi17
            _ = v19
            _ = v20
        }
        _ = v16
    }
    v5 -= 2
    if ((2 != 128) && false) {
        _ = len(v4)
        for fi21 := 0; fi21 < 13; fi21++ {
            fmt.Println(fi21, v3)
            _ = fi21
        }
    }
    if true {
        var v22 bool = (-995 > 2147483647)
        fmt.Println(v22)
        fmt.Println(v1)
        _ = v22
    } else {
        var v23 uint32 = 4294967295
        _ = v23
    }
    switch ("a" + ("a" + ("" + "a"))) {
    case "":
        if false {
            v2 -= 42
        } else {
            switch "" {
            case "a":
                var v24 int8 = -128
                _ = v24
            case "":
                v5 = (^(v5))
            case "hello":
                switch 177 {
                case 42:
                    v25 := GenS1{F0: (v2 ^ v2), F1: false}
                    _ = v25
                case -1:
                    v26 := v4[("hello" + "hello")]
                    _ = v26
                }
            default:
                for fi27 := 0; fi27 < 20; fi27++ {
                    for fi28 := 0; fi28 < 7; fi28++ {
                        fmt.Println(fi28)
                        var v29 bool = true
                        v4[""] = (fi27 << (((^(v5))) & 63))
                        var v30 uint8 = 255
                        if v29 {
                            v31 := GenS2{F0: (v2 | int64(v1))}
                            _ = v31
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    switch fi27 {
                    case 42:
                        fmt.Println(v1, v3)
                    case 4294967295:
                        v32 := []int{fi27}
                        _ = v32
                    }
                    v33 := GenS1{F0: (v2 ^ v2), F1: (true || (true && true))}
                    _ = fi27
                    _ = v33
                }
            }
            if (65535 > 2) {
                switch ((("" + ("foo bar" + "hello")) + ("a" + "hello")) + "foo bar") {
                case "foo bar":
                    var v34 uint8 = 255
                    _ = v34
                case "a":
                    for fi35 := 0; fi35 < 11; fi35++ {
                        if (-716 < 2147483647) {
                            v2 -= -1
                            v36 := GenS2{F0: (v2 + 9223372036854775807)}
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 13; fi37++ {
                            v38 := GenS2{F0: v2}
                            if (int64(v1) != 84) {
                                v2 += (v2 & (v2 << (((v5 << ((v5) & 63))) & 63)))
                                v5 -= (v5 * (v5 | (v5 * (v5 % (((v5 / ((917) | 1))) | 1)))))
                                if true {
                                    v39 := GenS3{F0: 128}
                                    var v40 uint8 = uint8(v1)
                                    v41 := []int{}
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                } else {
                                    fmt.Println(v5, fi37, v5)
                                    switch 128 {
                                    case 740:
                                        fmt.Println(v1)
                                        v42 := GenS1{F0: 450, F1: (!true)}
                                        _ = v42
                                    case -9223372036854775808:
                                        v5 += (v5 >> (((^(v5))) & 63))
                                    case -2147483648:
                                        v3 += -2.25
                                    }
                                    fmt.Println(fi35)
                                }
                            } else {
                                var v43 int = fi35
                                var v44 uint16 = (v1 >> (((v5 & (v5 | 1))) & 63))
                                var v45 int16 = 1000
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            _ = fi37
                            _ = v38
                        }
                        _ = fi35
                    }
                case "":
                    var v46 int16 = -296
                    _ = v46
                }
                v47 := GenS3{F0: 63}
                _ = v47
            } else {
                fmt.Println(v3, v5)
                if false {
                    if false {
                        var v48 uint8 = 42
                        v49 := []int{}
                        _ = v48
                        _ = v49
                    }
                    v2 += (v2 >> ((v5) & 63))
                    var v50 int16 = 135
                    _ = v50
                }
                var v51 int16 = -1
                _ = v51
            }
        }
    }
    fmt.Println(v2, v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
