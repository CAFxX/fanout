package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 bool; F1 int64; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 uint8; F2 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v5 uint8 = 0
    for fi6 := 0; fi6 < 10; fi6++ {
        if true {
            for fi7 := 0; fi7 < 19; fi7++ {
                if (!(false && (false && ((fi7 >= -9223372036854775808) && false)))) {
                    v8 := v3[732]
                    delete(v4, fi7)
                    _ = v8
                } else {
                    fmt.Println(fi7, fi7)
                    fmt.Println(v5)
                    var v9 float32 = ((0.0 - float32(1.5)) * (3.14159 + (1.5 / float32(1.5))))
                    v10 := []int{(fi7 << ((196) & 63)), 2147483647, -134, -2147483648}
                    fmt.Println(v9, fi7, fi6)
                    _ = v10
                    _ = v9
                }
                _ = fi7
            }
        } else {
            v5 = (^(v5))
            var v11 uint16 = 753
            if false {
                v12 := GenS1{F0: "a", F1: 2}
                _ = v12
            } else {
                for fi13 := 0; fi13 < 5; fi13++ {
                    _ = len(v4)
                    v14 := GenS2{F0: 1, F1: -36}
                    _ = len(v2)
                    _ = fi13
                    _ = v14
                }
            }
            _ = v11
        }
        if false {
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := GenS1{F0: "hello", F1: 676}
                var v17 int64 = -703
                _ = fi15
                _ = v16
                _ = v17
            }
            v18 := GenS2{F0: 2147483647, F1: 1}
            v19 := GenS2{F0: 2, F1: 127}
            _ = v18
            _ = v19
        } else {
            v5 = (v5 | 1)
        }
        v5 += v5
        _ = fi6
    }
    v2[32768] = (282 <= 2)
    if (32767 >= 280) {
        v20 := v1[(("a" + "a") + ("hello" + ("\n\t" + "\n\t")))]
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 float64 = 3.14159
            v23 := []int{32768, -680}
            _ = len(v4)
            _ = fi21
            _ = v22
            _ = v23
        }
        if v20 {
            switch 4294967295 {
            case 32768:
                for fi24 := 0; fi24 < 16; fi24++ {
                    if v20 {
                        for fi25 := 0; fi25 < 9; fi25++ {
                            v5 -= 128
                            v20 = false
                            v26 := GenS3{F0: v20, F1: -9223372036854775808, F2: fi25}
                            v27 := v2[202]
                            v28 := []int{fi25, (fi24 - fi24), fi25, 2147483647}
                            _ = fi25
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        v29 := GenS4{F0: (!v20)}
                        _ = v29
                    } else {
                        v5 += v5
                        v5 -= (^(v5))
                        v20 = false
                    }
                    _ = fi24
                }
            }
            for fi30 := 0; fi30 < 2; fi30++ {
                for fi31 := 0; fi31 < 11; fi31++ {
                    v32 := []int{(fi31 - (-(fi30))), (^(fi31))}
                    _ = len(v32)
                    _ = fi31
                    _ = v32
                }
                v5 -= (^(v5))
                fmt.Println(v20)
                _ = fi30
            }
            if (v20 && (42 == 18446744073709551615)) {
                if (65535 < 61) {
                    v20 = v20
                }
            }
        } else {
            v33 := v4[2147483647]
            switch 42 {
            case -9223372036854775808:
                for fi34 := 0; fi34 < 18; fi34++ {
                    v3[32768] = (fi34 * (^(fi34)))
                    v20 = v20
                    if (42 > -38) {
                        for fi35 := 0; fi35 < 8; fi35++ {
                            v36 := []int{(fi35 + fi35)}
                            var v37 float32 = ((float32(1e10) - 0.1) - 0.1)
                            v37 = 0.0
                            var v38 int64 = -9223372036854775808
                            _ = fi35
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                        var v39 uint32 = 1
                        _ = v39
                    } else {
                        fmt.Println(fi34, v33)
                    }
                    _ = fi34
                }
            case 32767:
                v40 := []int{}
                _ = v40
            default:
                fmt.Println(v5)
            }
            v5 = 127
            _ = v33
        }
        _ = v20
    } else {
        fmt.Println(v5)
    }
    v41 := []int{}
    v5 += v5
    v42 := GenS5{F0: true, F1: v5, F2: v5}
    v43 := v3[32768]
    v5 = 155
    _ = len(v2)
    fmt.Println(v43)
    var v44 uint64 = 4294967295
    var v45 int = 602
    switch "" {
    case "a":
        if false {
            var v46 uint16 = 71
            _ = v46
        }
    }
    v5 = (v5 >> ((18446744073709551615) & 63))
    var v47 int64 = -65
    v43 += v43
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v47
    _ = v5
}
