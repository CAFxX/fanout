package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
func genFn0() int8 {
    var v7 int = 1
    v8 := []int{}
    switch "foo bar" {
    case "hello":
        fmt.Println(v7)
    case "a":
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: (("" + "\n\t") + "foo bar"), F1: 117}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi9
            _ = v10
        }
    }
    v7 += (-(v7))
    _ = v7
    _ = v8
    return -74
}
func genFn1() int16 {
    v11 := GenS2{F0: 2, F1: -2147483648}
    if ((0 != -32768) || true) {
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(fi12)
            switch fi12 {
            case 42:
                var v13 int = fi12
                _ = v13
            case -865:
                switch fi12 {
                case 32767:
                    v14 := []int{259}
                    _ = v14
                case 223:
                    v15 := []int{(-(fi12)), fi12, fi12}
                    _ = v15
                }
            case 127:
                var v16 uint32 = 42
                _ = v16
            }
            _ = fi12
        }
        if (990 > 2) {
            for fi17 := 0; fi17 < 13; fi17++ {
                switch "a" {
                case "\n\t":
                    var v18 uint32 = 508
                    _ = v18
                default:
                    var v19 int = fi17
                    _ = v19
                }
                var v20 uint32 = 2
                switch ("" + "\n\t") {
                case "\n\t":
                    var v21 uint16 = 0
                    _ = v21
                default:
                    switch 230 {
                    case 32767:
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            var v22 string = "a"
                            _ = v22
                        case "\n\t":
                            v23 := []int{fi17, fi17, (fi17 ^ fi17)}
                            _ = v23
                        case "":
                            for fi24 := 0; fi24 < 5; fi24++ {
                                fmt.Println(fi17, fi24, fi24)
                                fmt.Println(v20, fi17)
                                if (736 != 401) {
                                    v25 := []int{(fi24 ^ (^(fi24))), (^(fi17)), fi24}
                                    var v26 uint8 = 0
                                    var v27 uint = 132
                                    v27 = 18446744073709551615
                                    if ((fi17 / ((fi24) | 1)) < fi24) {
                                        switch fi24 {
                                        case 4294967295:
                                            fmt.Println(v20)
                                        default:
                                            switch "hello" {
                                            case "":
                                                if (!true) {
                                                    fmt.Println(v27)
                                                    if len(v25) >= 2 {
                                                        v25 = v25[1:len(v25)]
                                                    }
                                                    var v28 uint8 = v26
                                                    v27 = (v27 / (((v27 << ((0) & 63))) | 1))
                                                    _ = v28
                                                } else {
                                                    v29 := GenS2{F0: 504, F1: fi24}
                                                    if (550 == 64) {
                                                        v26 -= v26
                                                    }
                                                    v30 := GenS3{F0: 2}
                                                    v27 += (v27 >> ((v27) & 63))
                                                    v26 += (v26 * 42)
                                                    _ = v29
                                                    _ = v30
                                                }
                                            }
                                        }
                                        var v31 float64 = ((0.0 + (1.5 / 1.0)) + 0.1)
                                        _ = v31
                                    } else {
                                        _ = len(v25)
                                    }
                                    v26 = 128
                                    var v32 int32 = -2147483648
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                    _ = v32
                                }
                                v20 = (v20 ^ 2147483648)
                                v20 = (v20 - 856)
                                v20 -= 1
                                _ = fi24
                            }
                        default:
                            v20 = v20
                        }
                    default:
                        fmt.Println(fi17, v20)
                    }
                }
                _ = fi17
                _ = v20
            }
        }
        var v33 bool = ((!(float64(0.0) >= 0.0)) || (true || (true || false)))
        _ = v33
    } else {
        var v34 string = "\n\t"
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 uint8 = 0
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    _ = v11
    return 42
}

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]uint32)
    v3 := []uint32{107}
    v4 := make(map[uint32]int)
    var v5 int16 = 32767
    var v6 int8 = 127
    v5 = (v5 / ((v5) | 1))
    v37 := GenS1{F0: "foo bar", F1: 255}
    v38 := GenS4{F0: -2147483648}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v39 uint32 = 65535
    var v40 int8 = -106
    v4[(^(v39))] = -2147483648
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v6)
    case "hello":
        if true {
            fmt.Println(v40, v5)
        } else {
            v1 = (^(v1))
        }
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        if true {
            var v42 uint = 65535
            delete(v4, (v39 * v39))
            _ = v42
        } else {
            fmt.Println(v5, fi41, v5)
            v40 -= v6
        }
        _ = fi41
    }
    v6 = v6
    v5 = 0
    var v43 uint32 = (v39 | 0)
    var v44 int16 = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
