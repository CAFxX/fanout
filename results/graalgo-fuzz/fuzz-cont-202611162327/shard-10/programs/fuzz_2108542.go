package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int8, p1 int32) uint8 {
    if (false && (-2.25 == -2.25)) {
        v7 := []int{}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 1; fi8++ {
            p1 -= (p1 - 2147483647)
            _ = fi8
        }
        switch "hello" {
        case "":
            for fi9 := 0; fi9 < 3; fi9++ {
                fmt.Println(p0, p0, p1)
                fmt.Println(p0)
                _ = fi9
            }
        case "\n\t":
            fmt.Println(p0)
        }
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 bool = true
        fmt.Println(v11, p0, fi10)
        _ = fi10
        _ = v11
    }
    if (!true) {
        var v12 uint8 = 2
        v13 := GenS1{F0: 18446744073709551615, F1: true, F2: p0}
        var v14 int64 = 849
        _ = v12
        _ = v13
        _ = v14
    } else {
        p1 += p1
        v15 := []int{385, 4294967295}
        _ = v15
    }
    if (566 <= 2) {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 uint = 773
            _ = fi16
            _ = v17
        }
    } else {
        p0 += 68
        p1 = p1
        p1 += (p1 | -1)
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 bool, p1 int, p2 int) int {
    p0 = p0
    var v18 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return p1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 2
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    var v6 float64 = 0.0
    v6 = v6
    var v19 uint64 = 4294967295
    fmt.Println(v2)
    for fi20 := 0; fi20 < 5; fi20++ {
        v19 = 811
        _ = fi20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 int64 = 694
        v23 := []int{}
        _ = fi21
        _ = v22
        _ = v23
    }
    v19 -= (v19 >> ((42) & 63))
    v24 := v3[656]
    if false {
        v25 := GenS1{F0: (v19 | v19), F1: (42 != 127), F2: 42}
        _ = v25
    }
    if (false || false) {
        fmt.Println(v24, v6)
        if (v24 == (v24 >> ((255) & 63))) {
            v26 := GenS2{F0: 2}
            v27 := v1[("\n\t" + ("a" + "hello"))]
            for fi28 := 0; fi28 < 8; fi28++ {
                for fi29 := 0; fi29 < 8; fi29++ {
                    v30 := []int{fi29, -1, 32767, v24}
                    v19 -= v19
                    switch fi28 {
                    case 9223372036854775807:
                        var v31 int16 = -775
                        _ = v31
                    case 903:
                        _ = len(v5)
                    }
                    if (0 < 42) {
                        if false {
                            var v32 uint = 2
                            v33 := GenS2{F0: (v19 - v19)}
                            fmt.Println(fi28, v32, v32)
                            _ = v32
                            _ = v33
                        } else {
                            v6 = -2.25
                            fmt.Println(v27)
                        }
                        v24 += fi29
                    } else {
                        if (-1 == 0) {
                            var v34 uint32 = 4294967295
                            var v35 float64 = (1.5 * float64(0.0))
                            if (!((false && ((127 <= -35) && (1000 != 0))) || v27)) {
                                switch genFn1((float64(0.1) != (float64(((3.14159 + (0.0 * (float32((v6 * 0.1)) + float32(v35)))) / (-2.25 + (((1e10 / float32(v6)) / 1.0) * 1.5)))) * v35)), (fi29 * fi28), (v24 - (v24 * fi28))) {
                                case 233:
                                    var v36 uint16 = 0
                                    _ = v36
                                case 127:
                                    v35 = (v6 * float64(float32(v35)))
                                case -1:
                                    switch "" {
                                    case "a":
                                        if false {
                                            var v37 int64 = -1
                                            _ = v37
                                        }
                                    case "\n\t":
                                        var v38 uint64 = v19
                                        _ = v38
                                    default:
                                        var v39 uint32 = 2147483648
                                        _ = v39
                                    }
                                }
                                if len(v30) >= 2 {
                                    v30 = v30[1:len(v30)]
                                }
                                var v40 int64 = -918
                                _ = len(v30)
                                _ = v40
                            }
                            _ = v34
                            _ = v35
                        }
                        fmt.Println(fi28)
                    }
                    _ = fi29
                    _ = v30
                }
                v41 := []int{}
                fmt.Println(v2, v19)
                _ = fi28
                _ = v41
            }
            _ = v26
            _ = v27
        } else {
            var v42 float64 = 3.14159
            var v43 uint16 = 969
            v42 = 0.0
            _ = v42
            _ = v43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
