package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 810
    v3 := make(map[string]string)
    var v4 uint32 = 2147483648
    v5 := make(map[string]int)
    var v6 uint32 = v4
    if (((1.5 / (float64(3.14159) - float64(0.0))) >= 0.1) || true) {
        v7 := GenS1{F0: 45}
        fmt.Println(v2)
        _ = v7
    }
    var v8 float64 = (-2.25 / 1.5)
    v9 := GenS2{F0: 127, F1: -398}
    switch "" {
    case "foo bar":
        fmt.Println(v8, v8, v2)
    case "hello":
        fmt.Println(v2, v4, v6)
    case "\n\t":
        if (-32768 > 166) {
            for fi10 := 0; fi10 < 20; fi10++ {
                v6 -= v6
                var v11 int8 = 0
                if (4294967295 >= 42) {
                    var v12 float32 = (0.0 - -2.25)
                    var v13 uint16 = 32767
                    var v14 int8 = -65
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
        } else {
            var v15 uint16 = 2
            if (!false) {
                if false {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        var v17 bool = true
                        for fi18 := 0; fi18 < 5; fi18++ {
                            fmt.Println(v2, fi16)
                            v19 := GenS2{F0: int(v4), F1: -151}
                            _ = fi18
                            _ = v19
                        }
                        _ = fi16
                        _ = v17
                    }
                    if (0 <= -72) {
                        for fi20 := 0; fi20 < 15; fi20++ {
                            var v21 float64 = v8
                            _ = fi20
                            _ = v21
                        }
                    } else {
                        if (!false) {
                            if (42 < 1) {
                                for fi22 := 0; fi22 < 10; fi22++ {
                                    v15 = (v15 << ((v2) & 63))
                                    v15 += (^(v15))
                                    _ = fi22
                                }
                                var v23 string = ""
                                v15 -= 42
                                _ = v23
                            } else {
                                v24 := v3["a"]
                                v25 := []int{-652}
                                v26 := GenS3{F0: 9223372036854775807, F1: 49, F2: 84}
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                            if true {
                                v15 -= v15
                                v27 := v1[42]
                                _ = v27
                            } else {
                                var v28 float64 = (v8 + v8)
                                v29 := []int{}
                                _ = v28
                                _ = v29
                            }
                        } else {
                            fmt.Println(v8, v4, v8)
                            fmt.Println(v8)
                        }
                        fmt.Println(v6, v8, v2)
                        var v30 uint = 152
                        _ = v30
                    }
                    v31 := GenS2{F0: 32767, F1: 42}
                    _ = v31
                } else {
                    if false {
                        v2 = (v2 * (v2 ^ (v2 * v2)))
                    } else {
                        var v32 int64 = 1
                        _ = v32
                    }
                }
            } else {
                v33 := GenS3{F0: -745, F1: 0, F2: 0}
                fmt.Println(v6)
                _ = v33
            }
            var v34 int64 = 42
            _ = v15
            _ = v34
        }
    default:
        v35 := []int{}
        _ = v35
    }
    var v36 bool = ((false || true) || false)
    var v37 int32 = -50
    switch -9223372036854775808 {
    case 32768:
        delete(v1, 128)
    case 38:
        v38 := v5["\n\t"]
        _ = v38
    }
    var v39 uint8 = 0
    v40 := GenS2{F0: 32768, F1: 2147483647}
    v37 += v37
    v36 = false
    v36 = (!false)
    v2 -= (^(v2))
    var v41 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
