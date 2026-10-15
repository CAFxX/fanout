package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int, p1 uint16) int32 {
    var v6 int16 = 0
    if false {
        fmt.Println(p0, v6)
    } else {
        if false {
            v7 := GenS1{F0: -128, F1: 1, F2: -751}
            var v8 uint16 = p1
            _ = v7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 uint32 = 1
        v11 := GenS1{F0: 88, F1: 4294967295, F2: -644}
        p1 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{1}
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return -651
}
func genFn1(p0 int, p1 uint, p2 uint32) uint16 {
    for fi13 := 0; fi13 < 7; fi13++ {
        p0 += fi13
        fmt.Println(p2, p0)
        _ = fi13
    }
    v14 := []int{(p0 & -9223372036854775808)}
    p0 = p0
    p2 += (p2 ^ (^(p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 2
}

func main() {
    var v1 int64 = 84
    v2 := make(map[uint32]string)
    var v3 float64 = 1e10
    var v4 int8 = -1
    var v5 uint32 = 676
    var v15 int64 = v1
    switch "hello" {
    case "\n\t":
        fmt.Println(v15, v4, v15)
    case "":
        var v16 int64 = v15
        _ = v16
    case "a":
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 2; fi18++ {
                switch ("\n\t" + "a") {
                case "hello":
                    v1 = 42
                case "":
                    switch ("\n\t" + "hello") {
                    case "a":
                        if false {
                            switch ("hello" + "\n\t") {
                            case "foo bar":
                                var v19 uint8 = 42
                                _ = v19
                            case "\n\t":
                                switch (fi17 % ((fi17) | 1)) {
                                case 127:
                                    v20 := []int{(^(fi18)), 9223372036854775807, (^(fi17)), (fi18 - (^(fi18)))}
                                    _ = v20
                                case 128:
                                    v21 := GenS1{F0: v4, F1: 60, F2: 2147483647}
                                    _ = v21
                                case 329:
                                    v15 = v15
                                default:
                                    if true {
                                        v2[(v5 + (v5 / ((v5) | 1)))] = "\n\t"
                                        if (v1 > v15) {
                                            switch (fi17 / (((fi17 + 524)) | 1)) {
                                            case 128:
                                                v1 -= 292
                                                v22 := GenS1{F0: v4, F1: 484, F2: 42}
                                                v23 := v2[v5]
                                                _ = v22
                                                _ = v23
                                            case 127:
                                                var v24 int = fi18
                                                _ = v24
                                            case -1:
                                                v1 = (v15 >> ((65535) & 63))
                                            default:
                                                var v25 uint16 = 255
                                                v4 = (v4 | int8(v5))
                                                v26 := GenS2{F0: 603}
                                                _ = v25
                                                _ = v26
                                            }
                                        }
                                        v2[(^(v5))] = "foo bar"
                                    } else {
                                        var v27 float64 = float64(0.1)
                                        v28 := []int{fi17, int(v1), 4294967295, fi18}
                                        delete(v2, (v5 * 658))
                                        _ = v27
                                        _ = v28
                                    }
                                }
                            default:
                                v15 -= v15
                            }
                        }
                    default:
                        v29 := GenS3{F0: 198, F1: 2}
                        _ = v29
                    }
                case "\n\t":
                    v30 := GenS2{F0: 664}
                    _ = v30
                }
                var v31 bool = false
                switch "\n\t" {
                case "hello":
                    switch "hello" {
                    case "hello":
                        v3 += v3
                    case "":
                        v2[v5] = "hello"
                    case "a":
                        v32 := GenS3{F0: 65535, F1: 0}
                        _ = v32
                    }
                case "a":
                    var v33 uint32 = (v5 ^ v5)
                    _ = v33
                case "foo bar":
                    var v34 uint = 21
                    _ = v34
                }
                _ = fi18
                _ = v31
            }
            for fi35 := 0; fi35 < 19; fi35++ {
                if (18446744073709551615 > 1) {
                    v3 += float64((float32((v3 * float64(-2.25))) + 1.5))
                    v36 := GenS2{F0: 326}
                    var v37 uint32 = (v5 + 4294967295)
                    _ = v36
                    _ = v37
                }
                _ = fi35
            }
            _ = fi17
        }
    default:
        v38 := GenS4{F0: -634}
        _ = v38
    }
    fmt.Println(v5)
    fmt.Println(v4, v3)
    var v39 int16 = -32768
    fmt.Println(v15, v3, v5)
    var v40 uint = 650
    var v41 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
