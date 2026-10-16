package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int64; F2 uint64 }
type GenS3 struct { F0 string; F1 uint; F2 uint }
type GenS4 struct { F0 int }
func genFn0(p0 string, p1 uint64) uint32 {
    var v4 int64 = 905
    switch p0 {
    case "":
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := GenS1{F0: 0}
            v4 += (v4 + (v4 >> ((321) & 63)))
            _ = fi5
            _ = v6
        }
    case "\n\t":
        v4 = (v4 | v4)
    case "hello":
        if true {
            v7 := []int{0}
            fmt.Println(p0, v4, v4)
            v4 = v4
            _ = v7
        } else {
            v8 := GenS2{F0: 1, F1: (v4 << ((204) & 63)), F2: (p1 - p1)}
            if true {
                for fi9 := 0; fi9 < 10; fi9++ {
                    v10 := []int{fi9, (fi9 + fi9)}
                    v10 = append(v10, (fi9 ^ fi9))
                    _ = fi9
                    _ = v10
                }
                switch 32767 {
                case -415:
                    v11 := []int{}
                    _ = v11
                case 9223372036854775807:
                    if ((1 != p1) || false) {
                        v12 := []int{-2147483648, 42, -178}
                        if true {
                            for fi13 := 0; fi13 < 19; fi13++ {
                                var v14 float32 = 0.1
                                switch p0 {
                                case "foo bar":
                                    var v15 int8 = -24
                                    _ = v15
                                }
                                var v16 uint16 = 32767
                                _ = fi13
                                _ = v14
                                _ = v16
                            }
                            var v17 uint8 = 255
                            _ = v17
                        } else {
                            p0 = "\n\t"
                            if (1 > 255) {
                                for fi18 := 0; fi18 < 3; fi18++ {
                                    var v19 string = (p0 + (p0 + p0))
                                    _ = fi18
                                    _ = v19
                                }
                                v20 := GenS2{F0: -305, F1: v4, F2: (p1 + (^(p1)))}
                                switch "a" {
                                case "":
                                    v21 := GenS2{F0: 42, F1: (v4 + v4), F2: 42}
                                    _ = v21
                                case "a":
                                    if (true && true) {
                                        for fi22 := 0; fi22 < 20; fi22++ {
                                            p0 = (p0 + p0)
                                            for fi23 := 0; fi23 < 10; fi23++ {
                                                if false {
                                                    var v24 uint = 255
                                                    _ = v24
                                                }
                                                _ = fi23
                                            }
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            _ = fi22
                                        }
                                        var v25 int8 = 100
                                        _ = v25
                                    } else {
                                        v26 := GenS2{F0: -198, F1: (v4 % ((-588) | 1)), F2: p1}
                                        _ = v26
                                    }
                                case "hello":
                                    var v27 float32 = float32(0.0)
                                    _ = v27
                                default:
                                    v4 = (v4 - -1)
                                }
                                _ = v20
                            }
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        }
                        _ = v12
                    } else {
                        v28 := []int{2147483647, 774, 32767, -2147483648}
                        p0 = p0
                        _ = v28
                    }
                }
            }
            _ = v8
        }
    }
    fmt.Println(p0, p0, p1)
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint8 = 128
        var v31 uint = 18446744073709551615
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := []int{-635}
    _ = p0
    _ = p1
    _ = v32
    _ = v4
    return 2
}

func main() {
    var v1 uint64 = 725
    var v2 uint8 = 139
    v3 := make(map[string]bool)
    if false {
        switch "foo bar" {
        case "hello":
            v33 := GenS3{F0: "a", F1: 0, F2: 18446744073709551615}
            _ = v33
        case "a":
            v34 := GenS4{F0: -2147483648}
            _ = v34
        case "foo bar":
            var v35 float32 = 3.14159
            _ = v35
        }
    }
    v36 := v3[""]
    var v37 int64 = 460
    v38 := GenS4{F0: 32768}
    v39 := GenS1{F0: 332}
    if (18446744073709551615 != 4294967295) {
        v40 := GenS1{F0: 4294967295}
        _ = v40
    } else {
        var v41 uint16 = 0
        _ = v41
    }
    var v42 float64 = -2.25
    v36 = (v36 || ((2 <= 2) || (-9223372036854775808 <= -493)))
    v1 -= v1
    var v43 int32 = 0
    v44 := []int{490, -9223372036854775808, -1}
    var v45 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
