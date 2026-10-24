package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 string; F2 uint }

func main() {
    var v1 int = -208
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 float32 = (1.5 * 0.0)
    switch v1 {
    case -827:
        if false {
            v5 := GenS1{F0: -41}
            _ = v5
        }
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 float64 = 3.14159
        for fi8 := 0; fi8 < 6; fi8++ {
            fmt.Println(v7, fi6, v4)
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(v4, v1, fi9)
        fmt.Println(v1, v4, v4)
        for fi10 := 0; fi10 < 9; fi10++ {
            fmt.Println(v4)
            var v11 bool = (621 <= 2)
            v12 := GenS2{F0: 334}
            if (1 <= -31) {
                if true {
                    fmt.Println(fi10, v11, fi9)
                    v1 = (-(v1))
                    var v13 uint32 = 4294967295
                    if (true && (32767 != 2)) {
                        v1 += fi10
                        v14 := []int{}
                        v15 := GenS1{F0: -128}
                        _ = v14
                        _ = v15
                    } else {
                        v13 = v13
                    }
                    _ = v13
                } else {
                    v16 := GenS1{F0: -1}
                    v4 = v4
                    v17 := GenS2{F0: 255}
                    _ = v16
                    _ = v17
                }
            } else {
                var v18 uint8 = 135
                if (431 > 2) {
                    fmt.Println(v1, fi9)
                } else {
                    var v19 uint8 = v18
                    fmt.Println(fi10, fi10)
                    if (1.5 != v4) {
                        v20 := GenS3{F0: 428}
                        v11 = (!false)
                        if v11 {
                            var v21 int16 = 1
                            v1 = fi10
                            var v22 int8 = 100
                            _ = v21
                            _ = v22
                        } else {
                            v18 = (v18 | (v19 | 49))
                            v23 := GenS3{F0: 2}
                            switch 0 {
                            case 481:
                                var v24 bool = (2 != 75)
                                if (!v11) {
                                    var v25 uint8 = (^(v18))
                                    var v26 uint8 = v18
                                    v19 = v19
                                    _ = v25
                                    _ = v26
                                } else {
                                    var v27 float64 = 1e10
                                    _ = v27
                                }
                                v28 := GenS3{F0: 1}
                                _ = v24
                                _ = v28
                            case -289:
                                v18 = 1
                            case 127:
                                v29 := []int{}
                                _ = v29
                            default:
                                v18 = v19
                            }
                            _ = v23
                        }
                        _ = v20
                    }
                    _ = v19
                }
                v30 := []int{fi9, v1, fi9}
                _ = v18
                _ = v30
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    v4 += ((0.0 - 3.14159) / 1e10)
    v1 = (v1 * (v1 - 2147483647))
    var v31 bool = (-1 != -213)
    var v32 uint64 = 327
    fmt.Println(v32, v31)
    for fi33 := 0; fi33 < 14; fi33++ {
        for fi34 := 0; fi34 < 3; fi34++ {
            if (!v31) {
                if true {
                    switch "a" {
                    case "hello":
                        _ = len(v3)
                    }
                    if (255 > uint(v32)) {
                        v35 := GenS1{F0: 42}
                        var v36 int32 = 0
                        if v31 {
                            if false {
                                var v37 int = (fi34 << ((1) & 63))
                                if (2147483647 <= 0) {
                                    var v38 bool = ((v36 / ((v36) | 1)) < (v36 / (((v36 % ((v36) | 1))) | 1)))
                                    fmt.Println(fi34, v38, v32)
                                    _ = v38
                                }
                                v39 := GenS4{F0: v36, F1: (("\n\t" + (("\n\t" + "\n\t") + ((("\n\t" + "\n\t") + ("foo bar" + "a")) + "foo bar"))) + ("" + ("hello" + ("hello" + "")))), F2: 1}
                                fmt.Println(v36, fi33)
                                _ = v37
                                _ = v39
                            } else {
                                v31 = true
                                fmt.Println(v1, v36, fi33)
                            }
                        }
                        if (!(42 != 1)) {
                            switch 32767 {
                            case 4294967295:
                                fmt.Println(v1)
                                v40 := GenS1{F0: 1}
                                _ = v40
                            case -677:
                                v4 += (-2.25 * (-2.25 * float32(((float64((v4 + v4)) * ((((3.14159 - 3.14159) + 1e10) * ((1.5 * 0.0) * 0.0)) - -2.25)) / 0.1))))
                            default:
                                v31 = v31
                                var v41 string = ("a" + "foo bar")
                                _ = v41
                            }
                            var v42 float32 = v4
                            v3[(("a" + "hello") + "a")] = false
                            var v43 uint64 = v32
                            _ = v42
                            _ = v43
                        } else {
                            fmt.Println(v32)
                            fmt.Println(v4)
                        }
                        fmt.Println(v4, v1)
                        _ = v35
                        _ = v36
                    }
                }
                var v44 uint64 = (v32 * v32)
                v45 := GenS1{F0: 127}
                _ = v44
                _ = v45
            }
            v46 := GenS1{F0: 42}
            v32 -= (^(v32))
            var v47 bool = ((^(v32)) == v32)
            _ = fi34
            _ = v46
            _ = v47
        }
        var v48 uint = 18446744073709551615
        v2[(fi33 ^ -9223372036854775808)] = v31
        _ = fi33
        _ = v48
    }
    v49 := []int{398, v1, (^(v1)), v1}
    fmt.Println(v31, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v49
}
