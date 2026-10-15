package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint8; F2 string }
func genFn0(p0 int64) int32 {
    fmt.Println(p0)
    p0 = (p0 ^ 1)
    _ = p0
    return 950
}
func genFn1(p0 uint32, p1 int64, p2 int64) int64 {
    fmt.Println(p0, p1)
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: "a"}
        var v9 uint8 = 127
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (p1 ^ (p2 % ((0) | 1)))
}

func main() {
    var v1 int = 95
    v2 := make(map[int]string)
    v3 := []uint32{751, 65535, 274}
    v4 := make(map[uint32]bool)
    var v5 int = (v1 / (((v1 << ((0) & 63))) | 1))
    var v6 uint8 = 116
    for fi11 := 0; fi11 < 15; fi11++ {
        v1 -= 1
        fmt.Println(v6, v1)
        switch (fi11 - (v1 + (-(fi11)))) {
        case -742:
            _ = len(v2)
        default:
            if (!true) {
                var v12 float32 = float32(3.14159)
                _ = v12
            } else {
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 int16 = 1
                    switch (fi13 * (fi11 & fi13)) {
                    case 891:
                        var v15 bool = true
                        _ = v15
                    case 42:
                        v16 := GenS2{F0: ("a" + "foo bar")}
                        _ = v16
                    }
                    v17 := v2[1]
                    _ = fi13
                    _ = v14
                    _ = v17
                }
                v5 = fi11
            }
        }
        _ = fi11
    }
    var v18 int16 = 0
    var v19 uint8 = 186
    switch "foo bar" {
    case "\n\t":
        v5 = 582
    case "":
        if (true || true) {
            v20 := GenS2{F0: ("a" + (("" + ("\n\t" + "a")) + ("" + "a")))}
            switch 0 {
            case -2147483648:
                if (!(("\n\t" + "\n\t") >= "a")) {
                    _ = len(v4)
                    switch "" {
                    case "foo bar":
                        v21 := GenS1{F0: ""}
                        _ = v21
                    case "":
                        if false {
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                        }
                    case "hello":
                        var v22 int16 = v18
                        _ = v22
                    default:
                        v3 = append(v3, 514)
                    }
                    switch "a" {
                    case "hello":
                        v6 = v19
                    case "a":
                        var v23 float32 = 0.1
                        _ = v23
                    default:
                        v19 -= v6
                    }
                }
            case 4294967295:
                fmt.Println(v6, v5, v18)
            default:
                fmt.Println(v19, v1, v1)
            }
            for fi24 := 0; fi24 < 17; fi24++ {
                v25 := GenS2{F0: ""}
                v19 += (v19 ^ (v6 / ((172) | 1)))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi24
                _ = v25
            }
            _ = v20
        } else {
            switch (-(v5)) {
            case -863:
                switch (-(v1)) {
                case 0:
                    for fi26 := 0; fi26 < 10; fi26++ {
                        v5 -= (fi26 % ((v1) | 1))
                        if (922 < 1) {
                            fmt.Println(fi26)
                            v27 := GenS3{F0: -1}
                            var v28 int32 = -586
                            _ = v27
                            _ = v28
                        } else {
                            v5 += v5
                            switch "\n\t" {
                            case "hello":
                                v2[(-(v1))] = "foo bar"
                            case "foo bar":
                                v29 := GenS1{F0: (("foo bar" + "foo bar") + ("hello" + "a"))}
                                _ = v29
                            case "\n\t":
                                v30 := GenS1{F0: ((("a" + "a") + ("a" + ("hello" + "foo bar"))) + ("foo bar" + "hello"))}
                                _ = v30
                            }
                        }
                        _ = fi26
                    }
                case -1:
                    if (-113 == 42) {
                        var v31 float32 = 0.0
                        if false {
                            v19 -= 156
                            v1 += 128
                        } else {
                            var v32 int8 = -1
                            _ = v32
                        }
                        _ = v31
                    } else {
                        if (!true) {
                            v3 = append(v3, 1)
                        } else {
                            v33 := GenS4{F0: 1, F1: v19, F2: "a"}
                            for fi34 := 0; fi34 < 7; fi34++ {
                                v18 -= 1000
                                for fi35 := 0; fi35 < 4; fi35++ {
                                    var v36 uint = 4294967295
                                    _ = fi35
                                    _ = v36
                                }
                                _ = fi34
                            }
                            _ = v33
                        }
                        v37 := GenS3{F0: 127}
                        _ = v37
                    }
                case 4294967295:
                    switch "" {
                    case "foo bar":
                        for fi38 := 0; fi38 < 5; fi38++ {
                            v39 := GenS3{F0: -128}
                            _ = fi38
                            _ = v39
                        }
                    case "":
                        v18 -= (v18 - v18)
                    case "\n\t":
                        delete(v4, 4294967295)
                    }
                }
            default:
                v18 = -1
            }
            v6 += v6
            for fi40 := 0; fi40 < 9; fi40++ {
                v41 := GenS1{F0: "a"}
                var v42 uint64 = 18446744073709551615
                _ = fi40
                _ = v41
                _ = v42
            }
        }
    case "foo bar":
        fmt.Println(v18, v6, v5)
    default:
        switch "a" {
        case "foo bar":
            v6 += v19
        case "":
            v18 = -1
        default:
            var v43 uint32 = 4294967295
            _ = v43
        }
    }
    var v44 uint16 = 499
    var v45 uint16 = v44
    var v46 float32 = ((-2.25 + 3.14159) + 3.14159)
    for fi47 := 0; fi47 < 12; fi47++ {
        var v48 int = (v5 - -1)
        _ = fi47
        _ = v48
    }
    v49 := GenS4{F0: 253, F1: v19, F2: (("\n\t" + (("\n\t" + "\n\t") + "hello")) + "\n\t")}
    v45 = (^(v45))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v49
    _ = v5
    _ = v6
}
