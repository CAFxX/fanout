package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int; F1 int; F2 int32 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint) int32 {
    var v4 int32 = 1
    switch "foo bar" {
    case "\n\t":
        var v5 int8 = -46
        _ = v5
    case "a":
        fmt.Println(p0, p0)
    default:
        switch "a" {
        case "hello":
            if true {
                switch (("hello" + "a") + "a") {
                case "\n\t":
                    for fi6 := 0; fi6 < 2; fi6++ {
                        if (!true) {
                            var v7 int64 = -9223372036854775808
                            var v8 string = "\n\t"
                            _ = v7
                            _ = v8
                        }
                        var v9 int = fi6
                        switch (v9 & (v9 % (((v9 >> ((p0) & 63))) | 1))) {
                        case 180:
                            v10 := []int{-252, (v9 % ((fi6) | 1)), (fi6 - fi6)}
                            _ = v10
                        case 4294967295:
                            if false {
                                v11 := GenS1{F0: 1, F1: 0, F2: 438}
                                _ = v11
                            } else {
                                for fi12 := 0; fi12 < 17; fi12++ {
                                    var v13 bool = false
                                    _ = fi12
                                    _ = v13
                                }
                                v4 += 42
                                v14 := []int{(fi6 & (fi6 >> ((255) & 63))), (fi6 ^ fi6), v9, 4294967295}
                                _ = v14
                            }
                        default:
                            v15 := []int{v9, (fi6 % (((v9 % ((32768) | 1))) | 1)), v9}
                            _ = v15
                        }
                        _ = fi6
                        _ = v9
                    }
                case "hello":
                    v16 := GenS2{F0: 127, F1: 174, F2: (v4 / (((v4 | v4)) | 1))}
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 18; fi17++ {
                        fmt.Println(p0, p0, v4)
                        for fi18 := 0; fi18 < 13; fi18++ {
                            v4 -= -365
                            _ = fi18
                        }
                        v19 := GenS1{F0: 42, F1: 0, F2: 1}
                        _ = fi17
                        _ = v19
                    }
                }
            } else {
                for fi20 := 0; fi20 < 4; fi20++ {
                    v4 -= v4
                    var v21 float64 = 1.5
                    v22 := []int{}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                p0 = 4294967295
                v4 += (v4 >> ((p0) & 63))
            }
        case "foo bar":
            v23 := []int{2147483647, 2147483647}
            _ = v23
        case "a":
            var v24 bool = true
            _ = v24
        }
    }
    _ = p0
    _ = v4
    return v4
}
func genFn1() uint {
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 uint8 = 74
        switch 113 {
        case -9223372036854775808:
            var v27 int8 = -1
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 uint16 = 1
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: 9223372036854775807, F1: -763, F2: -1}
    v31 := GenS2{F0: -619, F1: -1, F2: -1}
    var v32 string = "a"
    _ = v30
    _ = v31
    _ = v32
    return 146
}

func main() {
    var v1 int32 = 715
    var v2 int32 = (v1 ^ (v1 & v1))
    var v3 string = ("foo bar" + "\n\t")
    for fi33 := 0; fi33 < 18; fi33++ {
        fmt.Println(v2, fi33)
        _ = fi33
    }
    v34 := GenS3{F0: 49}
    var v35 int8 = -122
    var v36 uint16 = 32767
    var v37 float32 = 1e10
    v38 := GenS2{F0: 2147483647, F1: 169, F2: (v2 << ((1) & 63))}
    var v39 int32 = v1
    v40 := []int{int(v1), 0, 32768, -134}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
