package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int8; F2 bool }
type GenS4 struct { F0 string; F1 uint64; F2 string }
func genFn0(p0 string, p1 uint) int {
    p0 = ("" + "foo bar")
    if (p1 == 743) {
        if false {
            p0 = "a"
            var v5 int16 = 0
            v6 := GenS1{F0: 842}
            _ = v5
            _ = v6
        } else {
            v7 := GenS1{F0: 2147483647}
            var v8 int8 = 38
            _ = v7
            _ = v8
        }
        p0 = (((("a" + p0) + "a") + "") + "foo bar")
        v9 := []int{4294967295, 128}
        _ = v9
    }
    if false {
        if true {
            v10 := GenS1{F0: 672}
            if (("" + p0) < p0) {
                for fi11 := 0; fi11 < 9; fi11++ {
                    v12 := GenS1{F0: 2147483647}
                    p1 -= (p1 << ((1) & 63))
                    switch fi11 {
                    case 127:
                        switch (fi11 << ((p1) & 63)) {
                        case -9223372036854775808:
                            var v13 float32 = 1e10
                            _ = v13
                        default:
                            v14 := []int{4294967295, 2147483647, (fi11 * fi11)}
                            _ = v14
                        }
                    case 9223372036854775807:
                        v15 := []int{}
                        _ = v15
                    case -2147483648:
                        switch 894 {
                        case 128:
                            switch "" {
                            case "":
                                for fi16 := 0; fi16 < 1; fi16++ {
                                    v17 := GenS2{F0: ((((0.0 * (1e10 - -2.25)) * 0.0) != 0.1) || (!(1 < fi11)))}
                                    _ = fi16
                                    _ = v17
                                }
                            default:
                                v18 := GenS2{F0: false}
                                _ = v18
                            }
                        case 4294967295:
                            p0 = "\n\t"
                        default:
                            for fi19 := 0; fi19 < 5; fi19++ {
                                p0 = "a"
                                _ = fi19
                            }
                        }
                    default:
                        v20 := GenS2{F0: true}
                        _ = v20
                    }
                    _ = fi11
                    _ = v12
                }
                v21 := GenS1{F0: -1}
                p1 -= 2
                _ = v21
            } else {
                v22 := GenS1{F0: -90}
                v23 := []int{1}
                _ = v22
                _ = v23
            }
            _ = v10
        }
    } else {
        p0 = "\n\t"
        fmt.Println(p0, p1)
    }
    p1 += 42
    var v24 uint16 = 2
    _ = p0
    _ = p1
    _ = v24
    return 32768
}
func genFn1(p0 uint16, p1 int8) bool {
    var v25 uint16 = (p0 << ((255) & 63))
    if ((p1 ^ p1) > (p1 * (p1 & (p1 + -1)))) {
        p1 = 127
    } else {
        var v26 int = -1
        if (42 < p0) {
            var v27 uint64 = 60
            _ = v27
        } else {
            fmt.Println(p0)
            v28 := GenS3{F0: 2147483647, F1: p1, F2: false}
            _ = v28
        }
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v25
    return false
}

func main() {
    var v1 uint = 277
    v2 := make(map[string]uint32)
    var v3 bool = (!(true || false))
    v4 := make(map[uint32]string)
    v1 -= 65535
    v29 := []int{32767, 32768, 42}
    v30 := GenS2{F0: v3}
    v31 := GenS4{F0: "a", F1: 4294967295, F2: "\n\t"}
    var v32 int8 = -72
    v3 = (0 < 1)
    if (-674 <= -1) {
        v32 = (v32 ^ v32)
        v1 -= v1
    } else {
        fmt.Println(v32, v32)
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        for fi34 := 0; fi34 < 1; fi34++ {
            v35 := GenS3{F0: 0, F1: (v32 / (((v32 - v32)) | 1)), F2: v3}
            if (341 < 1) {
                v32 += -16
            } else {
                if ((0 < 0) || (-708 >= -688)) {
                    v36 := GenS2{F0: v3}
                    _ = v36
                } else {
                    if v3 {
                        v32 = v32
                    } else {
                        v32 -= v32
                    }
                    delete(v2, "hello")
                    v3 = (255 <= 42)
                }
                switch (("a" + "") + "foo bar") {
                case "a":
                    var v37 uint32 = 2147483648
                    _ = v37
                }
                v4[686] = ""
            }
            _ = fi34
            _ = v35
        }
        var v38 int32 = -739
        _ = fi33
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
