package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int16; F1 string; F2 uint8 }
func genFn0(p0 string, p1 uint64) string {
    fmt.Println(p1, p0)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    fmt.Println(p1)
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: -1, F1: 65535}
        switch p0 {
        case "":
            v9 := []int{fi7, (^(fi7)), (^(fi7)), fi7}
            _ = v9
        default:
            p0 = p0
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            fmt.Println(p1, p0)
            _ = fi10
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return ""
}

func main() {
    var v1 uint32 = 618
    var v2 int8 = 127
    v3 := make(map[int]uint32)
    var v4 int16 = 0
    var v11 float64 = (0.1 * (1.5 + 0.0))
    fmt.Println(v2, v1, v2)
    switch ("hello" + "") {
    case "":
        v12 := GenS1{F0: v2, F1: v1}
        _ = v12
    case "foo bar":
        var v13 uint16 = 65535
        _ = v13
    }
    if false {
        if true {
            switch ("a" + "\n\t") {
            case "a":
                if (v4 >= v4) {
                    var v14 string = "\n\t"
                    v2 -= v2
                    v15 := []int{128, int(v1), 572}
                    _ = v14
                    _ = v15
                } else {
                    var v16 string = ("hello" + ("foo bar" + "a"))
                    var v17 int = 337
                    _ = v16
                    _ = v17
                }
            case "hello":
                for fi18 := 0; fi18 < 2; fi18++ {
                    if true {
                        if (int32(v1) >= 2147483647) {
                            for fi19 := 0; fi19 < 18; fi19++ {
                                v1 = v1
                                v20 := GenS1{F0: (v2 | v2), F1: v1}
                                if true {
                                    if (128 < 150) {
                                        v21 := GenS1{F0: 1, F1: v1}
                                        if (v2 <= v2) {
                                            v22 := []int{-190, 900, (fi18 >> ((372) & 63))}
                                            if true {
                                                if (v1 == (^(v1))) {
                                                    if true {
                                                        v23 := GenS1{F0: v2, F1: v1}
                                                        _ = v23
                                                    } else {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            _ = len(v22)
                                                        default:
                                                            _ = len(v3)
                                                            switch -425 {
                                                            case -1:
                                                                v24 := GenS2{F0: v2, F1: 42, F2: 870}
                                                                _ = v24
                                                            case -771:
                                                                fmt.Println(v1)
                                                            default:
                                                                if (true && true) {
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v25 := GenS2{F0: (v2 ^ v2), F1: 144, F2: 322}
                                                                        _ = v25
                                                                    case "":
                                                                        v26 := GenS3{F0: v4, F1: 285}
                                                                        v27 := GenS3{F0: v4, F1: 2}
                                                                        var v28 float64 = v11
                                                                        _ = v26
                                                                        _ = v27
                                                                        _ = v28
                                                                    }
                                                                    v1 = (v1 / ((65535) | 1))
                                                                    var v29 float32 = (3.14159 - 3.14159)
                                                                    _ = v29
                                                                } else {
                                                                    v30 := GenS4{F0: false, F1: 320}
                                                                    fmt.Println(v11, v1, fi19)
                                                                    _ = v30
                                                                }
                                                            }
                                                            fmt.Println(v11)
                                                        }
                                                    }
                                                }
                                            }
                                            if len(v22) > 0 {
                                                v22[0] = 0
                                            }
                                            _ = v22
                                        }
                                        _ = len(v3)
                                        v11 += float64(0.1)
                                        _ = v21
                                    } else {
                                        v31 := []int{(fi18 << ((587) & 63)), (fi18 / ((fi19) | 1)), (fi18 << ((1) & 63))}
                                        fmt.Println(v11, fi18, v2)
                                        v2 -= v2
                                        _ = v31
                                    }
                                }
                                var v32 int16 = (v4 | v4)
                                var v33 float64 = v11
                                _ = fi19
                                _ = v20
                                _ = v32
                                _ = v33
                            }
                            v1 = (v1 | v1)
                        }
                        v1 = (v1 - (v1 & (v1 & (v1 * 4294967295))))
                    }
                    _ = fi18
                }
            default:
                var v34 uint = 65535
                _ = v34
            }
        } else {
            v4 -= 1
            v4 = 1
        }
        v35 := GenS1{F0: -1, F1: (v1 ^ (v1 ^ (v1 + 127)))}
        if true {
            var v36 float64 = 0.0
            v37 := GenS5{F0: 0, F1: 0}
            _ = v36
            _ = v37
        }
        _ = v35
    } else {
        v11 = v11
        fmt.Println(v2)
    }
    var v38 uint = 255
    if ((("foo bar" + ("a" + "\n\t")) + "\n\t") < "\n\t") {
        v39 := GenS6{F0: v4, F1: (("a" + "foo bar") + ("" + ("foo bar" + "hello"))), F2: 61}
        _ = v39
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        if false {
            var v41 int32 = 0
            _ = v41
        } else {
            var v42 float64 = ((((v11 * (0.1 - float64(float32(v11)))) - 1e10) * v11) * v11)
            _ = v42
        }
        v1 -= v1
        _ = fi40
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
