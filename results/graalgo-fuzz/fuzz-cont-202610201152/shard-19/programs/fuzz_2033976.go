package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint; F1 int; F2 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 uint8; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := []int{772}
    v4 := make(map[string]string)
    var v5 bool = ("hello" != ("a" + "a"))
    var v6 int8 = 118
    _ = len(v3)
    v6 += (v6 >> ((42) & 63))
    v6 -= (^(v6))
    v1[4294967295] = ""
    switch "hello" {
    case "foo bar":
        v5 = (!false)
    default:
        for fi7 := 0; fi7 < 18; fi7++ {
            if len(v3) > 0 {
                v3[0] = (fi7 % (((fi7 * (fi7 % ((fi7) | 1)))) | 1))
            }
            switch (fi7 * fi7) {
            case -624:
                switch fi7 {
                case -2147483648:
                    fmt.Println(fi7, fi7, v6)
                default:
                    v6 += v6
                }
            case 9223372036854775807:
                if true {
                    v3 = append(v3, -9223372036854775808)
                } else {
                    if len(v3) > 0 {
                        v3[0] = (^(fi7))
                    }
                    var v8 int = 127
                    if (true && false) {
                        switch "foo bar" {
                        case "foo bar":
                            for fi9 := 0; fi9 < 17; fi9++ {
                                var v10 int = (v8 << ((512) & 63))
                                var v11 bool = v5
                                if (!((42 == 179) || v11)) {
                                    fmt.Println(fi9, v5)
                                    var v12 uint8 = 42
                                    v13 := GenS1{F0: 957, F1: 65535, F2: (v6 / ((v6) | 1))}
                                    v11 = false
                                    _ = v12
                                    _ = v13
                                }
                                _ = fi9
                                _ = v10
                                _ = v11
                            }
                        default:
                            v14 := GenS1{F0: 736, F1: 2, F2: 127}
                            _ = v14
                        }
                        delete(v4, "")
                    }
                    _ = v8
                }
            default:
                switch "a" {
                case "foo bar":
                    switch ("\n\t" + "foo bar") {
                    case "\n\t":
                        v6 = v6
                    case "hello":
                        switch ("hello" + ((("" + "hello") + "\n\t") + "foo bar")) {
                        case "a":
                            v5 = (v6 <= 127)
                        case "":
                            var v15 float32 = 3.14159
                            _ = v15
                        default:
                            for fi16 := 0; fi16 < 15; fi16++ {
                                if (32768 > fi16) {
                                    v5 = true
                                    v6 -= 29
                                    v5 = v5
                                }
                                _ = fi16
                            }
                        }
                    case "a":
                        var v17 float64 = 0.0
                        _ = v17
                    }
                case "a":
                    fmt.Println(v5, v5, v5)
                default:
                    var v18 string = "foo bar"
                    _ = v18
                }
            }
            v19 := GenS2{F0: 0, F1: 42}
            _ = fi7
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int64 = 0
        v5 = ((float32(float64((-2.25 / 1e10))) == 3.14159) || false)
        _ = fi20
        _ = v21
    }
    if v5 {
        for fi22 := 0; fi22 < 20; fi22++ {
            if (!(!v5)) {
                if (-1 == (v6 & (v6 * v6))) {
                    if (-2147483648 > -873) {
                        var v23 int64 = 9223372036854775807
                        v24 := GenS1{F0: 332, F1: 255, F2: int8(v23)}
                        _ = v23
                        _ = v24
                    } else {
                        if v5 {
                            var v25 int16 = -32768
                            _ = v25
                        }
                    }
                }
                if v5 {
                    fmt.Println(v5)
                    fmt.Println(v6, fi22, v5)
                }
                switch (fi22 ^ 32767) {
                case -961:
                    if false {
                        v26 := GenS3{F0: 1, F1: fi22}
                        var v27 float64 = 3.14159
                        v28 := GenS1{F0: -895, F1: 995, F2: (v6 % ((v6) | 1))}
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                default:
                    v6 -= int8(fi22)
                }
            }
            v29 := GenS1{F0: -825, F1: 1, F2: (v6 / ((59) | 1))}
            v6 += 1
            _ = fi22
            _ = v29
        }
        var v30 uint16 = 32767
        _ = v30
    } else {
        v5 = (false || v5)
        if v5 {
            fmt.Println(v6, v6)
        } else {
            v5 = ((4294967295 < 798) && false)
        }
        var v31 uint = 18446744073709551615
        _ = v31
    }
    v32 := v1[699]
    v32 = "\n\t"
    v32 = "\n\t"
    switch 2147483647 {
    case -763:
        v33 := GenS4{F0: 4294967295, F1: -885, F2: v5}
        _ = v33
    default:
        var v34 uint64 = 4294967295
        _ = v34
    }
    switch ((v32 + v32) + v32) {
    case "\n\t":
        v35 := GenS5{F0: 0}
        _ = v35
    case "hello":
        fmt.Println(v32, v6)
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        fmt.Println(v5, v32)
        for fi37 := 0; fi37 < 5; fi37++ {
            v6 = (v6 ^ 1)
            var v38 int8 = 89
            _ = fi37
            _ = v38
        }
        switch "a" {
        case "foo bar":
            switch -2147483648 {
            case 4294967295:
                if v5 {
                    var v39 uint16 = 2
                    _ = v39
                } else {
                    v40 := v2[1]
                    _ = v40
                }
            case -2147483648:
                var v41 string = "a"
                _ = v41
            default:
                for fi42 := 0; fi42 < 18; fi42++ {
                    switch "a" {
                    case "a":
                        _ = len(v3)
                    case "hello":
                        v43 := GenS6{F0: 4, F1: 161, F2: 1}
                        _ = v43
                    case "\n\t":
                        var v44 bool = (83 >= 139)
                        _ = v44
                    default:
                        v45 := GenS4{F0: 1, F1: 42, F2: v5}
                        _ = v45
                    }
                    _ = fi42
                }
            }
        case "a":
            var v46 uint = 65535
            _ = v46
        case "":
            v47 := GenS5{F0: 2}
            _ = v47
        }
        _ = fi36
    }
    var v48 float32 = 1e10
    if (!v5) {
        if true {
            v5 = v5
        }
        v4[v32] = (((v32 + "\n\t") + (((v32 + v32) + v32) + "")) + v32)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
