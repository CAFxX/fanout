package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint8; F2 bool }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 string, p1 string) uint64 {
    var v7 uint16 = 0
    var v8 uint8 = 42
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 953
}

func main() {
    var v1 int64 = -797
    v2 := make(map[int]string)
    var v3 int32 = 2147483647
    var v4 bool = (1e10 < float32(3.14159))
    var v5 int = 9223372036854775807
    var v6 bool = v4
    if ((1 == 42) && v6) {
        var v9 int32 = (v3 & v3)
        var v10 uint64 = 18446744073709551615
        fmt.Println(v4, v10, v1)
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            v6 = v6
            _ = fi11
        }
    }
    var v12 float32 = (float32((1e10 / (1.5 - 1e10))) + ((((3.14159 / 3.14159) / 1.0) + (3.14159 * -2.25)) * (((-2.25 * -2.25) / (1.5 + ((float32(((float64(1e10) / 1.5) * -2.25)) + (0.1 * 0.1)) * 3.14159))) / 1.5)))
    fmt.Println(v4, v3, v4)
    var v13 uint8 = 127
    var v14 int16 = 1000
    for fi15 := 0; fi15 < 17; fi15++ {
        switch ((("hello" + ("hello" + "\n\t")) + "foo bar") + (("hello" + "foo bar") + "a")) {
        case "a":
            v16 := GenS1{F0: 2147483647, F1: 2}
            _ = v16
        case "hello":
            switch "hello" {
            case "hello":
                for fi17 := 0; fi17 < 12; fi17++ {
                    v18 := GenS2{F0: -2, F1: 255, F2: ((fi15 % ((219) | 1)) == 32767)}
                    v19 := []int{(fi17 - v5)}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case "a":
                switch (-(v5)) {
                case 128:
                    var v20 int32 = 1
                    _ = v20
                case -9223372036854775808:
                    v13 = v13
                }
            case "":
                if false {
                    var v21 float32 = (v12 - v12)
                    var v22 uint8 = v13
                    var v23 uint64 = 4294967295
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    if false {
                        var v24 int64 = (v1 << ((332) & 63))
                        var v25 bool = false
                        _ = v24
                        _ = v25
                    } else {
                        v26 := []int{(v5 >> ((255) & 63))}
                        _ = v26
                    }
                    v5 += fi15
                }
            default:
                fmt.Println(v5, v4, v3)
            }
        }
        _ = fi15
    }
    v13 += (^(v13))
    if (!v6) {
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := GenS3{F0: v5, F1: (("a" + "hello") + ("hello" + ("\n\t" + "")))}
            _ = fi27
            _ = v28
        }
    } else {
        for fi29 := 0; fi29 < 8; fi29++ {
            switch "hello" {
            case "":
                if v4 {
                    v3 -= (v3 >> ((18446744073709551615) & 63))
                    v30 := GenS1{F0: 4294967295, F1: 2}
                    if ((3.14159 == v12) || (255 != 2)) {
                        var v31 uint16 = 1
                        var v32 uint8 = (^(v13))
                        v33 := GenS2{F0: 0, F1: (^(v13)), F2: v6}
                        _ = v31
                        _ = v32
                        _ = v33
                    } else {
                        v34 := GenS4{F0: (("" + "a") + "foo bar"), F1: (int64(v13) <= v1)}
                        fmt.Println(fi29, v6, fi29)
                        for fi35 := 0; fi35 < 20; fi35++ {
                            if (-514 != (v1 ^ v1)) {
                                var v36 float32 = float32(3.14159)
                                var v37 uint16 = 0
                                var v38 int64 = v1
                                _ = v36
                                _ = v37
                                _ = v38
                            } else {
                                switch (fi29 / ((fi29) | 1)) {
                                case -348:
                                    if (2147483647 != 2147483647) {
                                        v39 := GenS1{F0: 0, F1: 2147483647}
                                        v1 -= v1
                                        v5 = fi35
                                        _ = v39
                                    }
                                default:
                                    v4 = false
                                    _ = len(v2)
                                }
                                fmt.Println(v4)
                            }
                            delete(v2, fi35)
                            var v40 int16 = v14
                            _ = fi35
                            _ = v40
                        }
                        _ = v34
                    }
                    _ = v30
                }
            }
            var v41 int16 = v14
            v42 := v2[fi29]
            _ = fi29
            _ = v41
            _ = v42
        }
        var v43 uint32 = 569
        _ = v43
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        var v45 int = (v5 % ((fi44) | 1))
        var v46 int = (^(v5))
        if (true && true) {
            var v47 string = ("a" + "a")
            v48 := []int{(^(v45)), v45}
            _ = v47
            _ = v48
        } else {
            v14 = v14
            v49 := []int{9223372036854775807, int(v3)}
            var v50 int32 = -1
            _ = v49
            _ = v50
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    for fi51 := 0; fi51 < 4; fi51++ {
        if (2 <= 981) {
            var v52 bool = false
            _ = v52
        } else {
            switch "hello" {
            case "\n\t":
                v1 = v1
            case "foo bar":
                var v53 float32 = 3.14159
                _ = v53
            case "hello":
                fmt.Println(v12, v5)
            }
            var v54 uint16 = 569
            _ = v54
        }
        _ = fi51
    }
    v3 = -816
    v55 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
