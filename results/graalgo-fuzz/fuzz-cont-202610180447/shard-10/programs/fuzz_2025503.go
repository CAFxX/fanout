package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 int }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 int64; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 709
    var v3 uint = 586
    v4 := GenS1{F0: -268}
    if (true && (-1 <= -261)) {
        switch ("a" + "a") {
        case "a":
            switch "foo bar" {
            case "a":
                v5 := GenS2{F0: 2}
                _ = v5
            case "hello":
                if false {
                    fmt.Println(v3, v3)
                } else {
                    v6 := GenS2{F0: 42}
                    _ = v6
                }
            }
        default:
            fmt.Println(v3, v2)
        }
        fmt.Println(v3, v3)
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 uint = (v3 + (^(v3)))
            for fi9 := 0; fi9 < 19; fi9++ {
                v10 := GenS3{F0: (fi7 + fi9), F1: 32767, F2: (fi9 >> (((v3 / ((v3) | 1))) & 63))}
                _ = fi9
                _ = v10
            }
            v8 = v3
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(v3)
    fmt.Println(v3, v2)
    v11 := []int{-159, 1, 42, -103}
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v3)
        _ = fi12
    }
    var v13 bool = ((255 < 2) && true)
    var v14 int32 = -1
    var v15 uint64 = 18446744073709551615
    v16 := GenS3{F0: 42, F1: 32767, F2: 128}
    var v17 bool = v13
    if (-2147483648 < int(v15)) {
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int8 = -128
            for fi20 := 0; fi20 < 19; fi20++ {
                fmt.Println(v19, fi18, v14)
                v13 = (((v2 == uint32(v15)) && v17) || v13)
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
    }
    if v13 {
        for fi21 := 0; fi21 < 7; fi21++ {
            switch 127 {
            case 0:
                for fi22 := 0; fi22 < 7; fi22++ {
                    v23 := GenS3{F0: fi21, F1: 32767, F2: -908}
                    v24 := v1[fi21]
                    v3 += (v3 * 255)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            case -439:
                for fi25 := 0; fi25 < 19; fi25++ {
                    fmt.Println(v17, v15)
                    v3 += 783
                    fmt.Println(v14, v13, v3)
                    v13 = v17
                    var v26 string = "\n\t"
                    if v13 {
                        if (!v17) {
                            v27 := GenS4{F0: -245, F1: v15}
                            _ = v27
                        } else {
                            v26 = "foo bar"
                            if false {
                                v13 = (float64((3.14159 + 3.14159)) < 3.14159)
                                if (!((v3 >> (((v3 ^ v3)) & 63)) < v3)) {
                                    v28 := GenS5{F0: uint32(v3)}
                                    var v29 uint16 = uint16(v2)
                                    _ = v28
                                    _ = v29
                                } else {
                                    var v30 uint16 = 167
                                    _ = v30
                                }
                                v13 = (9223372036854775807 >= -927)
                                var v31 string = v26
                                _ = v31
                            }
                            v1[fi25] = (v26 + "hello")
                        }
                        v32 := GenS2{F0: 1}
                        var v33 float32 = 0.0
                        _ = v32
                        _ = v33
                    } else {
                        var v34 int64 = 8
                        v35 := GenS1{F0: (v34 | (v34 + (v34 - (^(v34)))))}
                        _ = v34
                        _ = v35
                    }
                    _ = fi25
                    _ = v26
                }
            case 2147483647:
                v36 := GenS2{F0: 255}
                _ = v36
            }
            v37 := GenS6{F0: (v14 % ((v14) | 1)), F1: -1, F2: -121}
            _ = fi21
            _ = v37
        }
        v2 += (v2 % ((521) | 1))
        fmt.Println(v14)
    } else {
        if v13 {
            _ = len(v1)
            var v38 uint32 = (^(v2))
            var v39 bool = ((((3.14159 / 1e10) / (1.5 + 3.14159)) <= float64((-2.25 / 1.0))) || (42 >= uint8(v15)))
            _ = v38
            _ = v39
        }
    }
    v40 := GenS2{F0: 0}
    var v41 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
