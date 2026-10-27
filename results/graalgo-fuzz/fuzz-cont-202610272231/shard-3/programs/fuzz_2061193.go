package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 bool; F2 int64 }
type GenS5 struct { F0 int }
func genFn0() uint32 {
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 float32 = (((-2.25 / float32(-2.25)) * 1.5) + (float32(3.14159) * 1.5))
            v7 = (v7 - (-2.25 / ((v7 + (v7 * ((v7 * -2.25) * v7))) / v7)))
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v8 string = "a"
    for fi9 := 0; fi9 < 6; fi9++ {
        switch 32768 {
        case 0:
            switch fi9 {
            case 1:
                if false {
                    if true {
                        var v10 bool = false
                        v11 := []int{fi9}
                        v12 := GenS1{F0: 32767}
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                }
            case 32767:
                switch 4294967295 {
                case 127:
                    v13 := GenS1{F0: 65535}
                    _ = v13
                case 128:
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v8 = v8
                        v15 := GenS2{F0: 0}
                        _ = fi14
                        _ = v15
                    }
                case 32768:
                    v16 := GenS3{F0: (fi9 >> ((18446744073709551615) & 63))}
                    _ = v16
                default:
                    fmt.Println(fi9, fi9)
                }
            case 128:
                fmt.Println(fi9, fi9)
            }
        case -44:
            var v17 uint16 = 267
            _ = v17
        case 9223372036854775807:
            if (860 > -1) {
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 uint16 = 0
                    _ = fi18
                    _ = v19
                }
                if ((!(v8 >= v8)) && (32767 < 515)) {
                    var v20 uint32 = 0
                    _ = v20
                }
                var v21 uint32 = 511
                _ = v21
            } else {
                v8 = v8
                var v22 int32 = -1
                _ = v22
            }
        default:
            for fi23 := 0; fi23 < 6; fi23++ {
                v8 = "foo bar"
                _ = fi23
            }
        }
        _ = fi9
    }
    var v24 bool = (42 <= 96)
    v24 = ((3.14159 * 1e10) <= 3.14159)
    _ = v24
    _ = v8
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 18446744073709551615
    var v3 float32 = 1.5
    v4 := []string{}
    v1[1] = 1
    var v25 float64 = (0.1 * -2.25)
    var v26 int = 128
    v25 += (float64(v3) - v25)
    if false {
        switch v26 {
        case -155:
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 uint32 = genFn0()
                fmt.Println(v26, v25, v28)
                _ = fi27
                _ = v28
            }
        case 506:
            v4 = append(v4, ((("a" + "a") + ("\n\t" + "")) + "foo bar"))
        case -9223372036854775808:
            v29 := GenS1{F0: 2}
            _ = v29
        }
        if (!true) {
            var v30 uint8 = 1
            v31 := GenS3{F0: v26}
            v32 := GenS4{F0: 0, F1: false, F2: -247}
            _ = v30
            _ = v31
            _ = v32
        } else {
            v26 += v26
            var v33 uint64 = (v2 & 4294967295)
            if (-2147483648 < 2147483647) {
                if ((v33 & v2) < v33) {
                    var v34 int64 = 42
                    if true {
                        var v35 int32 = -723
                        delete(v1, genFn0())
                        var v36 uint = 0
                        _ = v35
                        _ = v36
                    } else {
                        if (-112 < 1000) {
                            var v37 string = "a"
                            _ = v37
                        }
                    }
                    fmt.Println(v33, v33, v3)
                    _ = v34
                }
            }
            _ = v33
        }
        if true {
            fmt.Println(v3)
        }
    }
    var v38 float64 = ((float64(v3) - 1.5) / ((1.5 / v25) - (v25 - float64(v3))))
    v3 -= v3
    fmt.Println(v26)
    v39 := v1[genFn0()]
    v40 := GenS3{F0: (-(v26))}
    v25 -= v25
    for fi41 := 0; fi41 < 12; fi41++ {
        v25 = v25
        if len(v4) > 0 {
            v4[0] = ((((("\n\t" + (("hello" + "foo bar") + ("hello" + "foo bar"))) + "hello") + "foo bar") + "foo bar") + "")
        }
        _ = fi41
    }
    v42 := GenS5{F0: v26}
    v43 := GenS1{F0: 1}
    if false {
        v2 += (v2 - (^(v2)))
        fmt.Println(v38)
    } else {
        var v44 int64 = 9223372036854775807
        var v45 uint = 18446744073709551615
        v39 = (v39 << (((v45 - 1)) & 63))
        _ = v44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v43
}
