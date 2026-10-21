package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint; F1 int; F2 uint8 }
func genFn0(p0 string, p1 int64, p2 int16) int32 {
    p2 += (p2 << ((255) & 63))
    v6 := GenS1{F0: 2, F1: (p2 >> ((683) & 63))}
    var v7 int32 = 21
    var v8 bool = (100 < 100)
    var v9 uint32 = 267
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -30
}
func genFn1() uint32 {
    var v10 int = 9223372036854775807
    v11 := GenS2{F0: "hello", F1: 878}
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 bool = (9223372036854775807 == 9223372036854775807)
        _ = fi12
        _ = v13
    }
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    var v1 int16 = -1
    v2 := []string{("\n\t" + ("a" + "a")), "hello"}
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]string)
    v5 := []uint32{65535, 65535, 2}
    if len(v5) > 0 {
        v5[0] = 121
    }
    if (!(1 == 291)) {
        v1 += v1
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS3{F0: 1, F1: true}
            switch "\n\t" {
            case "hello":
                v16 := GenS4{F0: 4294967295, F1: (fi14 >> ((4294967295) & 63)), F2: 1}
                _ = v16
            default:
                for fi17 := 0; fi17 < 7; fi17++ {
                    fmt.Println(fi14, v3)
                    var v18 string = ("hello" + "foo bar")
                    v3 = v3
                    _ = fi17
                    _ = v18
                }
            }
            switch (("foo bar" + "") + "foo bar") {
            case "foo bar":
                for fi19 := 0; fi19 < 7; fi19++ {
                    switch 2147483647 {
                    case 453:
                        var v20 int16 = (v1 << ((961) & 63))
                        _ = v20
                    default:
                        switch ("hello" + ("hello" + "")) {
                        case "foo bar":
                            var v21 uint = 2
                            if true {
                                if true {
                                    switch (fi19 - fi14) {
                                    case 42:
                                        var v22 uint16 = 65535
                                        _ = v22
                                    }
                                    v3 = v3
                                } else {
                                    _ = len(v4)
                                }
                                v1 = (v1 % (((v1 * (v1 * 42))) | 1))
                            } else {
                                fmt.Println(fi19, v3)
                            }
                            switch "foo bar" {
                            case "foo bar":
                                var v23 uint = (^(v21))
                                _ = v23
                            }
                            _ = v21
                        case "\n\t":
                            fmt.Println(v1, fi19)
                        case "":
                            v24 := GenS3{F0: v3, F1: true}
                            _ = v24
                        }
                        var v25 string = ""
                        if true {
                            v1 += 42
                            if true {
                                if (true || (false || ((0 == 1) && false))) {
                                    v1 = -32768
                                    v26 := GenS1{F0: 42, F1: (v1 % ((42) | 1))}
                                    _ = v26
                                } else {
                                    var v27 uint = 30
                                    fmt.Println(v1, v27)
                                    if false {
                                        var v28 uint8 = 2
                                        fmt.Println(v25, v1)
                                        _ = v28
                                    }
                                    _ = v27
                                }
                                var v29 float64 = 1e10
                                _ = v29
                            } else {
                                v5 = append(v5, 0)
                                var v30 string = ""
                                v30 = ((v25 + v30) + "foo bar")
                                fmt.Println(fi19, fi14, v25)
                                _ = v30
                            }
                            var v31 float32 = ((3.14159 - 3.14159) / 1.5)
                            _ = v31
                        }
                        var v32 float64 = 0.0
                        _ = v25
                        _ = v32
                    }
                    v33 := GenS3{F0: (v3 % ((42) | 1)), F1: (1 != 42)}
                    _ = fi19
                    _ = v33
                }
            case "":
                if (42 == 42) {
                    v3 = (v3 | v3)
                    switch fi14 {
                    case 407:
                        switch (fi14 * int(v1)) {
                        case -2147483648:
                            var v34 string = "hello"
                            _ = v34
                        default:
                            v35 := GenS2{F0: ("\n\t" + "a"), F1: 761}
                            _ = v35
                        }
                    case -1:
                        var v36 int16 = (-(v1))
                        _ = v36
                    case 557:
                        var v37 float32 = (-2.25 - 0.1)
                        _ = v37
                    default:
                        if true {
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(v1, v1)
                            case "\n\t":
                                var v38 int16 = (v1 - (v1 - -32768))
                                _ = v38
                            default:
                                v1 -= (v1 >> ((634) & 63))
                            }
                        } else {
                            var v39 uint16 = 0
                            v39 = v39
                            _ = v39
                        }
                    }
                    var v40 uint = 314
                    _ = v40
                } else {
                    v41 := GenS2{F0: ("\n\t" + ("hello" + "\n\t")), F1: 1}
                    _ = v41
                }
            case "a":
                v1 -= (v1 ^ v1)
            }
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v1, v1)
    v3 = v3
    v5 = append(v5, 1)
    v1 = (v1 + (v1 & v1))
    v42 := GenS3{F0: (^(v3)), F1: (127 <= 1)}
    v1 = (^(v1))
    var v43 float32 = 0.1
    v1 = v1
    v5 = append(v5, 2)
    if true {
        v1 -= -1
    } else {
        fmt.Println(v43)
        v43 = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
