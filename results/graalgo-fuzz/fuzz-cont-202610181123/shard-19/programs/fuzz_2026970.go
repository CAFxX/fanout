package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 uint16; F1 string; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0(p0 int16, p1 int64) uint8 {
    var v7 uint = 0
    p1 += (p1 >> ((1) & 63))
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 0
    var v3 int32 = -237
    var v4 string = "\n\t"
    var v5 uint32 = 2
    v6 := make(map[string]bool)
    switch 935 {
    case 128:
        var v8 uint32 = v5
        _ = v8
    case 2147483647:
        var v9 uint = 255
        _ = v9
    case 32767:
        v10 := GenS1{F0: "foo bar", F1: v4}
        _ = v10
    default:
        v2 = v2
    }
    switch 1 {
    case 32767:
        if (270 < 1) {
            var v11 int8 = 31
            if false {
                v12 := GenS1{F0: "a", F1: v4}
                switch 42 {
                case 127:
                    fmt.Println(v5)
                case -926:
                    var v13 uint16 = 1
                    _ = v13
                default:
                    v2 = v2
                }
                fmt.Println(v2, v2, v2)
                _ = v12
            }
            _ = v11
        } else {
            var v14 uint16 = 677
            var v15 uint8 = genFn0(v2, -1)
            var v16 int8 = 1
            _ = v14
            _ = v15
            _ = v16
        }
    case -241:
        var v17 uint = 263
        _ = v17
    default:
        if ((1.5 - 1.5) < 1e10) {
            v3 += -872
            fmt.Println(v3)
        } else {
            v18 := GenS1{F0: "a", F1: ""}
            v19 := []int{2147483647}
            v3 += v3
            _ = v18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        fmt.Println(v5, v3)
        _ = fi20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        switch fi21 {
        case 1:
            switch 32768 {
            case 32767:
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 uint16 = 32767
                    v24 := []int{(fi21 % ((32768) | 1)), 127}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            default:
                fmt.Println(v4)
            }
        case 0:
            var v25 int8 = int8(v2)
            _ = v25
        default:
            var v26 bool = (1 <= 0)
            _ = v26
        }
        _ = fi21
    }
    v27 := GenS1{F0: (((("hello" + "foo bar") + v4) + v4) + (v4 + "hello")), F1: "foo bar"}
    var v28 string = ("\n\t" + v4)
    v29 := GenS2{F0: 4294967295, F1: v4, F2: (("a" + ((v4 + "\n\t") + "hello")) + v28)}
    switch 32767 {
    case -139:
        if false {
            v30 := GenS1{F0: v4, F1: "a"}
            v31 := GenS3{F0: 32767, F1: "a", F2: (v5 % ((1) | 1))}
            _ = v30
            _ = v31
        }
    case 128:
        if true {
            fmt.Println(v2, v28)
            v4 = v4
            if ((!(!false)) && (0 != 42)) {
                switch v28 {
                case "a":
                    v28 = ((v4 + "hello") + "foo bar")
                case "":
                    var v32 float32 = 1.5
                    _ = v32
                }
                v5 = v5
            } else {
                delete(v1, v4)
            }
        } else {
            v5 -= (^(v5))
        }
    case 0:
        v2 = (v2 % ((v2) | 1))
    default:
        for fi33 := 0; fi33 < 20; fi33++ {
            fmt.Println(v5, v4, v2)
            for fi34 := 0; fi34 < 10; fi34++ {
                _ = len(v6)
                switch fi33 {
                case 456:
                    if ((false && false) && false) {
                        if (3.14159 < 3.14159) {
                            v35 := GenS4{F0: -513}
                            v36 := []int{(fi33 << ((255) & 63))}
                            _ = v35
                            _ = v36
                        } else {
                            v37 := GenS5{F0: 0}
                            var v38 int16 = -32768
                            _ = v37
                            _ = v38
                        }
                        fmt.Println(fi34, v3)
                        v39 := []int{}
                        v40 := GenS4{F0: (v2 ^ 1000)}
                        _ = v39
                        _ = v40
                    }
                case 42:
                    if ("hello" == "hello") {
                        switch (v28 + ("foo bar" + v28)) {
                        case "hello":
                            v41 := v6[v4]
                            _ = v41
                        case "a":
                            var v42 int64 = 607
                            _ = v42
                        case "foo bar":
                            var v43 int32 = v3
                            _ = v43
                        default:
                            v2 = 1
                        }
                        v3 += v3
                        switch fi33 {
                        case 127:
                            var v44 uint16 = 65535
                            _ = v44
                        case -9223372036854775808:
                            v28 = ((("\n\t" + (v28 + v4)) + v4) + "hello")
                        }
                    }
                default:
                    v45 := GenS6{F0: v3, F1: 8}
                    _ = v45
                }
                v5 -= (v5 << ((4294967295) & 63))
                var v46 int64 = 42
                _ = fi34
                _ = v46
            }
            _ = fi33
        }
    }
    var v47 string = ((((v4 + (v4 + "foo bar")) + v28) + "foo bar") + ((((v28 + "foo bar") + (v28 + (((("a" + ((((("hello" + v4) + v28) + v4) + v4) + (((v28 + "a") + "foo bar") + "\n\t"))) + "\n\t") + (v4 + v28)) + ""))) + v28) + v28))
    delete(v6, v47)
    fmt.Println(v4)
    var v48 int64 = -517
    delete(v1, "a")
    for fi49 := 0; fi49 < 5; fi49++ {
        switch "" {
        case "":
            delete(v6, v4)
        case "a":
            if false {
                v5 += (v5 & v5)
                v50 := []int{127}
                _ = v50
            } else {
                for fi51 := 0; fi51 < 1; fi51++ {
                    if (!((3.14159 / (1e10 - ((((float32((0.0 - 1.5)) * (0.0 - ((1e10 - 1.5) / (3.14159 * 0.1)))) - 1e10) / 1.0) * 3.14159))) < 1e10)) {
                        var v52 uint8 = 2
                        _ = v52
                    }
                    switch (v47 + (v28 + "hello")) {
                    case "hello":
                        v28 = ((v4 + v28) + v28)
                    case "\n\t":
                        if (false || (!(0.0 != 0.0))) {
                            switch fi51 {
                            case 1:
                                v53 := []int{(fi51 << ((255) & 63))}
                                _ = v53
                            default:
                                v54 := GenS6{F0: (v3 % ((2147483647) | 1)), F1: -128}
                                _ = v54
                            }
                        } else {
                            fmt.Println(v48)
                        }
                    case "a":
                        var v55 uint16 = 32767
                        _ = v55
                    }
                    v56 := GenS2{F0: 42, F1: v47, F2: "\n\t"}
                    _ = fi51
                    _ = v56
                }
                var v57 int8 = 127
                var v58 float64 = float64(-2.25)
                _ = v57
                _ = v58
            }
        case "\n\t":
            v59 := GenS7{F0: 2, F1: v3, F2: v48}
            _ = v59
        }
        fmt.Println(v2, v2)
        _ = len(v1)
        _ = fi49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
