package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 bool }
func genFn0(p0 int64) bool {
    v7 := []int{-391, -513, 2147483647}
    for fi8 := 0; fi8 < 5; fi8++ {
        p0 = 1
        v9 := GenS1{F0: 1, F1: 810, F2: true}
        _ = fi8
        _ = v9
    }
    v7 = append(v7, 259)
    fmt.Println(p0)
    if false {
        var v10 uint16 = 876
        fmt.Println(v10, p0)
        if (!false) {
            var v11 uint32 = 969
            v12 := GenS2{F0: 247, F1: 0}
            _ = v11
            _ = v12
        } else {
            if (true || (!true)) {
                var v13 int32 = -2147483648
                _ = v13
            } else {
                v10 = (v10 + 340)
                p0 = p0
                fmt.Println(v10, p0)
            }
            if (943 != -1) {
                switch "a" {
                case "a":
                    var v14 uint32 = 0
                    _ = v14
                case "foo bar":
                    v10 = v10
                default:
                    var v15 float32 = -2.25
                    _ = v15
                }
                if ("\n\t" <= ((("foo bar" + ("hello" + "foo bar")) + "\n\t") + "a")) {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        if true {
                            p0 -= (-(p0))
                            if false {
                                switch ("foo bar" + "\n\t") {
                                case "foo bar":
                                    _ = len(v7)
                                case "":
                                    if ((!(2147483648 == 2147483647)) || true) {
                                        var v17 uint8 = 1
                                        _ = v17
                                    }
                                case "hello":
                                    if len(v7) > 0 {
                                        v7[0] = fi16
                                    }
                                }
                                v18 := GenS1{F0: (^(v10)), F1: 965, F2: true}
                                _ = v18
                            } else {
                                var v19 uint64 = 18446744073709551615
                                for fi20 := 0; fi20 < 1; fi20++ {
                                    if true {
                                        if len(v7) > 0 {
                                            v7[0] = 32767
                                        }
                                        var v21 uint8 = 33
                                        if false {
                                            var v22 uint8 = v21
                                            var v23 uint16 = 2
                                            _ = v22
                                            _ = v23
                                        } else {
                                            v21 += (v21 | (^(v21)))
                                        }
                                        _ = v21
                                    } else {
                                        fmt.Println(fi16)
                                        if false {
                                            v7 = append(v7, 32767)
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                        } else {
                                            v19 += 500
                                        }
                                    }
                                    var v24 int8 = 100
                                    if len(v7) > 0 {
                                        v7[0] = -1
                                    }
                                    var v25 uint = 919
                                    _ = len(v7)
                                    _ = fi20
                                    _ = v24
                                    _ = v25
                                }
                                _ = v19
                            }
                            v26 := GenS3{F0: 127}
                            _ = v26
                        }
                        if false {
                            v27 := GenS3{F0: 0}
                            if true {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            for fi28 := 0; fi28 < 14; fi28++ {
                                v10 -= 619
                                _ = fi28
                            }
                            _ = v27
                        }
                        _ = fi16
                    }
                    if ("foo bar" < "a") {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        for fi29 := 0; fi29 < 3; fi29++ {
                            p0 = int64(v10)
                            var v30 int16 = 42
                            _ = fi29
                            _ = v30
                        }
                    }
                    fmt.Println(p0, v10)
                } else {
                    v31 := GenS4{F0: 2147483647, F1: (!(false && (!true)))}
                    _ = v31
                }
                var v32 string = "hello"
                _ = v32
            }
        }
        _ = v10
    }
    _ = p0
    _ = v7
    return false
}
func genFn1(p0 int8, p1 uint16, p2 uint64) uint16 {
    v33 := []int{32767}
    fmt.Println(p1, p0)
    for fi34 := 0; fi34 < 16; fi34++ {
        p1 = (p1 >> ((2) & 63))
        var v35 int8 = 42
        _ = fi34
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    return (p1 * p1)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 int64 = 0
    var v4 string = "\n\t"
    v5 := []int{}
    v6 := make(map[uint32]int)
    var v36 int32 = -487
    if (!(!(551 <= 0))) {
        delete(v1, 42)
        var v37 int8 = 0
        _ = v37
    } else {
        v36 = v36
        fmt.Println(v4, v3, v4)
        v5 = append(v5, -458)
    }
    v4 = (v4 + "\n\t")
    v4 = v4
    var v38 uint = 4294967295
    v39 := GenS1{F0: uint16(v38), F1: 4294967295, F2: ((32767 == 32767) && false)}
    if ((3.14159 * 1e10) > (-2.25 - (0.0 - ((0.0 + (1e10 + (-2.25 + (float64(float32((1.5 - (((0.1 - 0.1) / -2.25) + 0.1)))) * 0.1)))) / (float64(0.1) * 0.1))))) {
        if (true || (-286 == 127)) {
            var v40 float64 = float64((3.14159 / 1.0))
            var v41 int = -525
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v40
            _ = v41
        }
    }
    v42 := GenS3{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
