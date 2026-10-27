package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32 }
func genFn0() bool {
    var v6 uint64 = 4294967295
    v7 := GenS1{F0: 42}
    v8 := GenS2{F0: 879}
    _ = v6
    _ = v7
    _ = v8
    return (0.1 >= 0.0)
}
func genFn1(p0 uint64, p1 uint) int {
    v9 := GenS3{F0: -1, F1: 0, F2: 1}
    p0 -= (p0 / ((p0) | 1))
    v10 := GenS3{F0: 32767, F1: 625, F2: 1}
    p1 -= (p1 & 18446744073709551615)
    v11 := GenS4{F0: p0}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return -859
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (float64(1e10) * ((-2.25 / float64((3.14159 + 1e10))) - (1e10 * (3.14159 / 0.1))))
    var v3 int64 = -903
    v4 := make(map[string]int)
    var v5 uint8 = 42
    v12 := GenS3{F0: 14, F1: 2, F2: v5}
    var v13 float32 = ((float32(1.5) / 3.14159) - (-2.25 - ((1.5 * 0.0) * float32(float64(1.5)))))
    if true {
        v14 := GenS2{F0: 42}
        _ = v14
    }
    var v15 float64 = 0.0
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 int = fi16
        var v18 uint = 65535
        v19 := GenS5{F0: -128}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    switch "\n\t" {
    case "":
        v20 := GenS6{F0: -889}
        _ = v20
    case "foo bar":
        switch (("\n\t" + "") + "\n\t") {
        case "a":
            for fi21 := 0; fi21 < 14; fi21++ {
                for fi22 := 0; fi22 < 12; fi22++ {
                    v13 += 0.1
                    _ = fi22
                }
                _ = fi21
            }
        case "hello":
            fmt.Println(v15, v5)
        default:
            v3 += (v3 - 1)
        }
    case "hello":
        switch "" {
        case "":
            var v23 int64 = v3
            _ = v23
        case "foo bar":
            v24 := []int{genFn1(222, 211), 128, 42}
            _ = v24
        default:
            if (4294967295 == 1) {
                v5 += v5
            } else {
                v25 := GenS4{F0: 0}
                switch 127 {
                case -1:
                    for fi26 := 0; fi26 < 18; fi26++ {
                        if false {
                            v27 := []int{}
                            v27 = append(v27, (-(fi26)))
                            v5 = v5
                            _ = v27
                        } else {
                            if true {
                                _ = len(v1)
                                if true {
                                    for fi28 := 0; fi28 < 14; fi28++ {
                                        v5 = v5
                                        v29 := v1["\n\t"]
                                        _ = fi28
                                        _ = v29
                                    }
                                    fmt.Println(fi26, v13, v15)
                                    v13 = (float32((0.0 + v15)) - float32(float64(0.0)))
                                }
                            } else {
                                var v30 uint16 = 1
                                v1["\n\t"] = 2147483647
                                _ = v30
                            }
                        }
                        _ = fi26
                    }
                case 1:
                    if ((v13 - v13) > v13) {
                        for fi31 := 0; fi31 < 8; fi31++ {
                            fmt.Println(fi31)
                            var v32 int16 = 556
                            _ = fi31
                            _ = v32
                        }
                        delete(v1, ((("hello" + ((("foo bar" + "a") + "\n\t") + "\n\t")) + "") + "a"))
                    }
                case -101:
                    for fi33 := 0; fi33 < 18; fi33++ {
                        v5 = 127
                        var v34 int8 = 100
                        for fi35 := 0; fi35 < 18; fi35++ {
                            var v36 uint32 = 178
                            _ = fi35
                            _ = v36
                        }
                        _ = fi33
                        _ = v34
                    }
                default:
                    v37 := v4[("hello" + "hello")]
                    _ = v37
                }
                var v38 float32 = v13
                _ = v25
                _ = v38
            }
        }
    }
    v15 = (v15 + 1.5)
    switch 4294967295 {
    case 128:
        fmt.Println(v15, v15, v13)
    }
    fmt.Println(v15, v13)
    var v39 string = "a"
    var v40 float64 = v2
    v15 -= v40
    fmt.Println(v15, v40)
    if false {
        v41 := []int{-402, -476}
        for fi42 := 0; fi42 < 2; fi42++ {
            var v43 uint64 = 0
            var v44 float32 = -2.25
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = v41
    } else {
        var v45 uint64 = 4294967295
        v46 := []int{int(v5)}
        var v47 uint64 = v45
        _ = v45
        _ = v46
        _ = v47
    }
    var v48 string = "hello"
    var v49 uint8 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v48
    _ = v49
    _ = v5
}
