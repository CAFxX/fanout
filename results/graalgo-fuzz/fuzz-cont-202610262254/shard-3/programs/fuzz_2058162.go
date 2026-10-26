package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int64; F1 uint; F2 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = ((float64(1.5) + ((0.1 * 0.0) + float64((((1e10 / float32((1.5 / float64(3.14159)))) + float32(float64((((float32(3.14159) + 1.5) + float32(1.5)) * (1e10 / -2.25))))) - ((1.5 * ((float32(3.14159) / ((0.0 * ((0.0 + (1.5 + 1e10)) - 1.5)) - 1e10)) * ((3.14159 - (1.5 * float32(3.14159))) - -2.25))) / 0.1))))) + 3.14159)
    var v3 uint16 = 535
    v4 := []string{}
    var v5 int64 = -963
    var v6 int32 = -82
    if (!true) {
        var v7 int64 = 422
        var v8 float64 = v2
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 571, F1: 0}
        switch 9223372036854775807 {
        case 42:
            if true {
                v10 := GenS2{F0: (^(v3))}
                if false {
                    v6 -= v6
                }
                var v11 uint8 = 0
                _ = v10
                _ = v11
            } else {
                fmt.Println(v3)
                if (true || (18446744073709551615 > 2)) {
                    v2 += 0.1
                }
                if ((v6 + 0) > v6) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v2, v3)
                }
            }
        case 32768:
            v3 += v3
        case 68:
            switch "a" {
            case "hello":
                v12 := GenS3{F0: 98}
                _ = v12
            }
        default:
            if (42 >= 1) {
                if (1 != 100) {
                    if (!false) {
                        v3 += v3
                        v1[42] = -2147483648
                        for fi13 := 0; fi13 < 6; fi13++ {
                            var v14 int32 = v6
                            _ = fi13
                            _ = v14
                        }
                    } else {
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            fmt.Println(v5, v2, v2)
                        case "\n\t":
                            for fi15 := 0; fi15 < 20; fi15++ {
                                var v16 uint64 = 761
                                v1[9223372036854775807] = fi15
                                fmt.Println(v3, v5)
                                _ = fi15
                                _ = v16
                            }
                        default:
                            switch 42 {
                            case 4294967295:
                                v2 = float64(0.0)
                            case 128:
                                fmt.Println(v6)
                            case 253:
                                v17 := v1[32767]
                                _ = v17
                            }
                        }
                        var v18 uint = 42
                        _ = v18
                    }
                    v3 += v3
                    for fi19 := 0; fi19 < 3; fi19++ {
                        if true {
                            var v20 uint64 = 18446744073709551615
                            _ = v20
                        } else {
                            switch ("" + ("hello" + ((("\n\t" + "\n\t") + "\n\t") + ("hello" + (("foo bar" + (("\n\t" + (("hello" + ("\n\t" + "")) + "a")) + "hello")) + "hello"))))) {
                            case "":
                                if (!((v5 >> ((uint(v3)) & 63)) > (v5 / (((v5 / (((^(v5))) | 1))) | 1)))) {
                                    var v21 float32 = 1e10
                                    for fi22 := 0; fi22 < 4; fi22++ {
                                        v23 := GenS2{F0: v3}
                                        fmt.Println(v5)
                                        _ = fi22
                                        _ = v23
                                    }
                                    _ = v21
                                }
                            case "foo bar":
                                fmt.Println(v6, v6, v2)
                            case "hello":
                                v5 = v5
                            }
                            v4 = append(v4, "")
                        }
                        _ = fi19
                    }
                } else {
                    v2 -= float64(3.14159)
                    var v24 string = "foo bar"
                    v24 = ((v24 + (v24 + "a")) + ((v24 + v24) + v24))
                    _ = v24
                }
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 string = ""
                    for fi27 := 0; fi27 < 13; fi27++ {
                        fmt.Println(v26, fi27)
                        if true {
                            var v28 int32 = v6
                            v26 = v26
                            fmt.Println(v6, v2, v2)
                            _ = v28
                        }
                        switch (v26 + v26) {
                        case "":
                            if ((fi25 - fi25) >= (fi25 << ((0) & 63))) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            var v29 uint = 65535
                            _ = v29
                        case "\n\t":
                            switch 4294967295 {
                            case -2147483648:
                                v3 += 1
                            case 0:
                                fmt.Println(fi25)
                            case 778:
                                v30 := v1[32767]
                                _ = v30
                            }
                            var v31 int16 = 32767
                            _ = v31
                        case "hello":
                            fmt.Println(v5)
                        }
                        _ = fi27
                    }
                    v32 := GenS4{F0: (fi25 >> ((42) & 63)), F1: false}
                    _ = fi25
                    _ = v26
                    _ = v32
                }
                v6 -= (v6 % ((v6) | 1))
            } else {
                v5 = v5
                fmt.Println(v5, v6)
                fmt.Println(v3, v5)
            }
        }
        _ = v9
    }
    if (v6 < v6) {
        v33 := GenS5{F0: 42, F1: "hello"}
        v34 := GenS6{F0: -9223372036854775808, F1: 1, F2: 65535}
        _ = v33
        _ = v34
    } else {
        switch ("foo bar" + ("" + ("\n\t" + "foo bar"))) {
        case "\n\t":
            var v35 int = -2147483648
            _ = v35
        case "a":
            v3 -= v3
        case "":
            v36 := GenS7{F0: 0}
            _ = v36
        default:
            v4 = append(v4, "foo bar")
        }
    }
    switch 127 {
    case 127:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 958:
        var v37 int16 = 32767
        _ = v37
    default:
        v38 := GenS2{F0: v3}
        _ = v38
    }
    v6 += v6
    v39 := GenS8{F0: -708}
    var v40 int8 = 0
    var v41 uint64 = 0
    delete(v1, -9223372036854775808)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
