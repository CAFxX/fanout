package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint16, p1 int32) int16 {
    v5 := GenS1{F0: 42, F1: 42}
    p1 = (p1 + (p1 >> ((2) & 63)))
    switch (("hello" + ("hello" + "hello")) + "hello") {
    case "a":
        p1 = p1
    case "\n\t":
        if (false && (153 > 832)) {
            var v6 bool = (!(1e10 == float64(0.0)))
            v7 := GenS1{F0: -367, F1: 1}
            _ = v6
            _ = v7
        }
    default:
        var v8 int16 = 1
        _ = v8
    }
    p1 -= 1
    var v9 int = -9223372036854775808
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 32767
}
func genFn1(p0 uint32, p1 int64, p2 int16) int8 {
    switch "" {
    case "":
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 string = "a"
            var v12 int8 = 14
            if (!true) {
                v13 := GenS2{F0: v11, F1: p0}
                _ = v13
            } else {
                p2 -= genFn0(42, 2147483647)
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    default:
        v14 := GenS2{F0: "hello", F1: p0}
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        if true {
            p1 -= p1
            for fi16 := 0; fi16 < 17; fi16++ {
                fmt.Println(fi15)
                var v17 int8 = 1
                _ = fi16
                _ = v17
            }
        } else {
            switch "hello" {
            case "foo bar":
                v18 := []int{fi15}
                _ = v18
            case "hello":
                for fi19 := 0; fi19 < 11; fi19++ {
                    switch fi15 {
                    case 32767:
                        var v20 int64 = (p1 * (p1 ^ -9223372036854775808))
                        _ = v20
                    case -9223372036854775808:
                        fmt.Println(fi15, p2)
                        if true {
                            v21 := []int{(fi15 - int(p1)), fi15, fi15, 128}
                            if false {
                                var v22 uint8 = 2
                                if false {
                                    if false {
                                        p2 += p2
                                    }
                                } else {
                                    var v23 float32 = float32(float64(((1.5 / ((-2.25 * (0.1 - -2.25)) + -2.25)) * 1e10)))
                                    var v24 bool = (-2.25 == float64((1.5 / (v23 / v23))))
                                    _ = len(v21)
                                    _ = v23
                                    _ = v24
                                }
                                var v25 uint16 = 851
                                _ = v22
                                _ = v25
                            } else {
                                v21 = append(v21, (fi15 ^ 4294967295))
                                p1 += 765
                                _ = len(v21)
                                v26 := GenS1{F0: int16(p0), F1: (p2 << ((uint(p0)) & 63))}
                                _ = v26
                            }
                            v27 := GenS3{F0: 873, F1: 127}
                            switch (("foo bar" + "hello") + "foo bar") {
                            case "hello":
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                            case "a":
                                if false {
                                    p1 += 42
                                }
                            case "foo bar":
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                            default:
                                if true {
                                    var v28 uint16 = 65535
                                    var v29 int16 = genFn0(v28, -809)
                                    _ = v28
                                    _ = v29
                                } else {
                                    var v30 bool = false
                                    _ = v30
                                }
                            }
                            fmt.Println(p1, fi15)
                            v31 := GenS3{F0: 467, F1: 255}
                            _ = v21
                            _ = v27
                            _ = v31
                        }
                    case 127:
                        v32 := []int{(fi15 * (fi19 - fi15)), (-(fi19)), fi19, fi19}
                        _ = v32
                    }
                    _ = fi19
                }
            default:
                p2 -= (p2 + 0)
            }
        }
        v33 := []int{}
        _ = fi15
        _ = v33
    }
    var v34 uint8 = 135
    v35 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    _ = v35
    return 127
}

func main() {
    var v1 int64 = 42
    var v2 bool = (56 > 115)
    v3 := make(map[string]bool)
    var v4 int8 = 102
    v36 := []int{32768, 0, 32768, 242}
    delete(v3, "hello")
    var v37 float32 = float32(float64(0.1))
    var v38 float32 = (3.14159 / v37)
    _ = len(v3)
    fmt.Println(v38, v4)
    var v39 uint32 = 324
    var v40 float32 = v38
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
