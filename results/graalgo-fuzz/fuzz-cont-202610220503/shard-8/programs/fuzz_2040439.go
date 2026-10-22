package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint32; F2 bool }
type GenS6 struct { F0 uint64; F1 uint64; F2 int16 }

func main() {
    var v1 bool = ("" >= "foo bar")
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 uint16 = 255
    v5 := make(map[string]int)
    v4 = v4
    var v6 int8 = 0
    var v7 uint8 = 1
    var v8 int8 = (-(v6))
    v9 := GenS1{F0: 100}
    for fi10 := 0; fi10 < 6; fi10++ {
        if v1 {
            var v11 float32 = (float32((1.5 + (0.1 * 1.5))) * (-2.25 - 1.5))
            if v1 {
                v11 = 0.0
            } else {
                for fi12 := 0; fi12 < 11; fi12++ {
                    if v1 {
                        var v13 uint32 = 42
                        v14 := GenS1{F0: v6}
                        fmt.Println(v11, v8, fi12)
                        if v1 {
                            var v15 uint = 42
                            v16 := GenS1{F0: (v8 * v8)}
                            v17 := GenS2{F0: v7, F1: (v4 / (((v4 + v4)) | 1))}
                            v15 = (v15 >> ((18446744073709551615) & 63))
                            _ = v15
                            _ = v16
                            _ = v17
                        } else {
                            switch 32767 {
                            case 32768:
                                switch -9223372036854775808 {
                                case -9223372036854775808:
                                    if true {
                                        v18 := v3[(v13 + uint32(v4))]
                                        if (-1 < 42) {
                                            v7 = v7
                                            var v19 float64 = (0.0 / 1.0)
                                            var v20 int = fi10
                                            _ = v19
                                            _ = v20
                                        } else {
                                            v13 = v13
                                            v8 = -10
                                            switch "hello" {
                                            case "foo bar":
                                                v21 := GenS2{F0: v7, F1: v4}
                                                _ = v21
                                            case "a":
                                                var v22 float32 = (v11 * v11)
                                                _ = v22
                                            default:
                                                fmt.Println(v18)
                                            }
                                            var v23 uint16 = 2
                                            _ = v23
                                        }
                                        switch (fi10 >> ((4294967295) & 63)) {
                                        case 2147483647:
                                            v24 := GenS3{F0: 0, F1: (v18 + "")}
                                            _ = v24
                                        case 4294967295:
                                            fmt.Println(fi12, v1)
                                            v25 := GenS4{F0: 1, F1: 561}
                                            _ = v25
                                        }
                                        _ = v18
                                    } else {
                                        v8 += (v6 | v6)
                                        v8 = v8
                                    }
                                case 42:
                                    v11 = 0.0
                                }
                            case 0:
                                v26 := []int{(fi10 + int(v4)), fi12}
                                _ = v26
                            default:
                                fmt.Println(v8, fi10)
                            }
                            var v27 int16 = int16(v7)
                            _ = v27
                        }
                        fmt.Println(v13, fi12, v8)
                        v28 := GenS5{F0: 970, F1: (v13 >> ((2) & 63)), F2: (!false)}
                        v29 := GenS1{F0: v8}
                        _ = v13
                        _ = v14
                        _ = v28
                        _ = v29
                    }
                    _ = fi12
                }
                v30 := GenS4{F0: 65535, F1: 42}
                _ = v30
            }
            _ = v11
        } else {
            switch ("a" + "hello") {
            case "a":
                v8 -= v6
            default:
                for fi31 := 0; fi31 < 14; fi31++ {
                    v7 = 49
                    _ = fi31
                }
            }
            v6 = (-(v6))
        }
        _ = fi10
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        if (!true) {
            v33 := GenS6{F0: 529, F1: 18446744073709551615, F2: -1}
            _ = v33
        } else {
            var v34 int16 = -281
            _ = v34
        }
        _ = fi32
    }
    fmt.Println(v4, v6, v1)
    var v35 int = 0
    v36 := v2["hello"]
    v37 := v5[("" + ("hello" + ("" + "\n\t")))]
    var v38 uint8 = uint8(v4)
    v39 := GenS4{F0: 4294967295, F1: 0}
    v40 := []int{v37, (v36 << ((615) & 63)), v36, (v37 & 9223372036854775807)}
    var v41 uint32 = 674
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
