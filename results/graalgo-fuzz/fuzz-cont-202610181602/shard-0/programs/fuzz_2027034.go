package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 string; F2 bool }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int; F1 uint8; F2 int32 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 int16 }
func genFn0(p0 string) int32 {
    var v7 int16 = -204
    var v8 bool = (42 >= 32767)
    var v9 int8 = 120
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 int8, p1 int64, p2 uint) uint64 {
    var v10 uint8 = 188
    v11 := []int{127}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    p2 -= (p2 | (p2 | p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 uint = 255
    var v2 int = -239
    var v3 uint64 = 2
    v4 := make(map[string]uint32)
    var v5 int32 = 1
    var v6 int8 = 1
    fmt.Println(v5, v1)
    var v12 string = "hello"
    if false {
        for fi13 := 0; fi13 < 8; fi13++ {
            v2 = v2
            v12 = v12
            v14 := GenS1{F0: 203, F1: -1, F2: v6}
            _ = fi13
            _ = v14
        }
    }
    if ((v5 * v5) < (v5 % ((v5) | 1))) {
        var v15 string = v12
        v5 += (v5 / ((int32(v2)) | 1))
        _ = v15
    }
    v2 -= 599
    v16 := GenS1{F0: 0, F1: -1, F2: v6}
    if true {
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := GenS1{F0: 127, F1: (v6 << ((v1) & 63)), F2: v6}
            if false {
                fmt.Println(v5, v3)
                v1 = 65535
                if ((v5 >> (((v1 << (((^(v1))) & 63))) & 63)) < (v5 >> ((4294967295) & 63))) {
                    if true {
                        if false {
                            switch (v2 | (v2 + (v2 & fi17))) {
                            case -122:
                                v5 += v5
                            case 127:
                                for fi19 := 0; fi19 < 6; fi19++ {
                                    v5 = (v5 + v5)
                                    _ = fi19
                                }
                            }
                            if false {
                                var v20 int32 = (v5 ^ int32(v1))
                                v20 = (^(v5))
                                _ = v20
                            } else {
                                v5 += 1
                                v21 := []int{-970, -137, 0, v2}
                                _ = v21
                            }
                        } else {
                            fmt.Println(v1)
                            var v22 int32 = genFn0(((v12 + v12) + (v12 + v12)))
                            for fi23 := 0; fi23 < 17; fi23++ {
                                delete(v4, v12)
                                switch ("\n\t" + v12) {
                                case "":
                                    var v24 int8 = (v6 / (((v6 % (((-(v6))) | 1))) | 1))
                                    _ = v24
                                case "hello":
                                    var v25 string = "foo bar"
                                    _ = v25
                                case "\n\t":
                                    var v26 float32 = float32(0.1)
                                    _ = v26
                                default:
                                    if false {
                                        switch (fi17 + fi17) {
                                        case -9223372036854775808:
                                            var v27 string = (((("" + "foo bar") + (("\n\t" + ("hello" + "a")) + v12)) + v12) + ((v12 + v12) + v12))
                                            _ = v27
                                        case 127:
                                            v2 += int(v22)
                                            fmt.Println(fi23, v6)
                                            var v28 float32 = 3.14159
                                            _ = v28
                                        case 2147483647:
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v29 bool = ((v12 + v12) == "hello")
                                                _ = v29
                                            case "a":
                                                switch fi17 {
                                                case 42:
                                                    if (((((0.0 - float32(1e10)) - (-2.25 * float32(1e10))) - (3.14159 + (((1e10 + 1.5) - 1.5) + float32(0.0)))) / (1e10 + 3.14159)) > 3.14159) {
                                                        v6 += (v6 & (v6 ^ v6))
                                                    } else {
                                                        if (v6 <= (v6 + (^(v6)))) {
                                                            v6 = 100
                                                            v30 := []int{}
                                                            if true {
                                                                var v31 bool = false
                                                                _ = v31
                                                            } else {
                                                                var v32 int8 = v6
                                                                v32 += v6
                                                                _ = v32
                                                            }
                                                            _ = v30
                                                        }
                                                        var v33 uint = (v1 + 1)
                                                        _ = len(v4)
                                                        _ = v33
                                                    }
                                                case 32768:
                                                    v34 := []int{fi23, int(v3), fi23, (fi23 - v2)}
                                                    _ = v34
                                                case -2147483648:
                                                    switch 1 {
                                                    case 127:
                                                        v1 = (^(v1))
                                                    }
                                                }
                                            }
                                        }
                                        v4["hello"] = 961
                                        if ((1.5 / 1.0) <= (1.5 * (1e10 + 1.5))) {
                                            var v35 float64 = (3.14159 + 1.5)
                                            _ = v35
                                        } else {
                                            var v36 uint = 432
                                            fmt.Println(v36)
                                            _ = len(v4)
                                            fmt.Println(v36, v3, v5)
                                            v22 = v22
                                            _ = v36
                                        }
                                        v37 := GenS2{F0: 255, F1: v1, F2: 128}
                                        _ = v37
                                    } else {
                                        v38 := GenS2{F0: 1, F1: 2, F2: 42}
                                        v39 := []int{(^(fi23)), fi23}
                                        v40 := GenS3{F0: 0}
                                        v41 := GenS4{F0: false, F1: ("foo bar" + ""), F2: (1.5 <= 1.5)}
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                }
                                _ = fi23
                            }
                            _ = v22
                        }
                    }
                    v42 := GenS3{F0: 42}
                    v43 := GenS5{F0: v1, F1: 0}
                    _ = v42
                    _ = v43
                }
            }
            fmt.Println(v12)
            _ = fi17
            _ = v18
        }
        switch "" {
        case "\n\t":
            if (0.0 != 0.0) {
                if true {
                    var v44 float64 = 1.5
                    switch 267 {
                    case 127:
                        v45 := GenS6{F0: v2, F1: 128, F2: v5}
                        _ = v45
                    default:
                        v46 := GenS7{F0: -128, F1: (v6 & v6)}
                        _ = v46
                    }
                    switch 128 {
                    case 1:
                        v3 -= (v3 * (^(v3)))
                    case 0:
                        var v47 int = v2
                        _ = v47
                    case 42:
                        fmt.Println(v44, v6, v44)
                    }
                    _ = v44
                } else {
                    v48 := GenS8{F0: 1000}
                    v2 = v2
                    _ = v48
                }
            } else {
                var v49 uint8 = 127
                _ = v49
            }
        case "hello":
            v50 := []int{9223372036854775807}
            _ = v50
        }
        for fi51 := 0; fi51 < 5; fi51++ {
            v2 += (-(fi51))
            _ = fi51
        }
    }
    v52 := []int{}
    for fi53 := 0; fi53 < 2; fi53++ {
        if true {
            var v54 uint16 = 32767
            _ = v54
        } else {
            var v55 uint8 = 128
            v56 := GenS8{F0: 42}
            _ = len(v52)
            _ = v55
            _ = v56
        }
        fmt.Println(v2, v3)
        fmt.Println(v12)
        _ = fi53
    }
    fmt.Println(v1)
    for fi57 := 0; fi57 < 15; fi57++ {
        switch (v12 + v12) {
        case "":
            switch v2 {
            case 963:
                var v58 bool = false
                _ = v58
            case 9223372036854775807:
                fmt.Println(v5, v5, v12)
            default:
                if len(v52) >= 2 {
                    v52 = v52[1:len(v52)]
                }
            }
        case "hello":
            var v59 int8 = (v6 << ((1) & 63))
            _ = v59
        }
        var v60 uint64 = (^(v3))
        v61 := v4[((v12 + v12) + "")]
        _ = fi57
        _ = v60
        _ = v61
    }
    switch (v2 & v2) {
    case 1:
        v4[v12] = 42
    case -732:
        var v62 uint32 = 4294967295
        _ = v62
    }
    v1 = 22
    if (127 == 204) {
        if (877 >= 65535) {
            fmt.Println(v6, v1)
        } else {
            v63 := GenS3{F0: -9223372036854775808}
            _ = v63
        }
        v4[v12] = 968
    } else {
        var v64 int32 = (-(v5))
        if ((v5 >> (((v1 & (v1 - v1))) & 63)) != v5) {
            v1 -= v1
            v6 = v6
        }
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
