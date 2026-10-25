package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]bool)
    var v3 int16 = -1
    var v4 uint = 18446744073709551615
    v5 := make(map[string]string)
    v6 := GenS1{F0: 1, F1: 255}
    fmt.Println(v4, v1, v1)
    for fi7 := 0; fi7 < 11; fi7++ {
        switch (fi7 << ((4294967295) & 63)) {
        case 653:
            v8 := GenS2{F0: 1}
            _ = v8
        default:
            switch "hello" {
            case "foo bar":
                var v9 float64 = 0.1
                _ = v9
            case "\n\t":
                var v10 bool = (v4 == 1)
                _ = v10
            case "":
                if (42 > 0) {
                    var v11 float64 = (float64(1e10) * 1e10)
                    if (true || (!((0 < 1) || ((-2.25 + (1.5 - ((0.0 - float32(3.14159)) - (1.5 / 3.14159)))) <= 0.0)))) {
                        switch (v1 + (v1 + v1)) {
                        case "foo bar":
                            if (!true) {
                                var v12 int8 = 1
                                v4 = (v4 + (v4 ^ v4))
                                if (true || (42 != 824)) {
                                    var v13 float64 = (v11 - (v11 * v11))
                                    v14 := []int{}
                                    _ = v13
                                    _ = v14
                                }
                                _ = v12
                            } else {
                                v15 := GenS1{F0: 255, F1: 0}
                                var v16 int64 = -519
                                var v17 uint32 = 527
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        }
                        for fi18 := 0; fi18 < 13; fi18++ {
                            v19 := v2[42]
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        for fi20 := 0; fi20 < 12; fi20++ {
                            switch v1 {
                            case "a":
                                fmt.Println(v3, fi7)
                            case "hello":
                                v5[(v1 + ("" + "foo bar"))] = (v1 + (("\n\t" + "hello") + "foo bar"))
                            }
                            v21 := GenS2{F0: 1}
                            _ = fi20
                            _ = v21
                        }
                        v22 := []int{}
                        var v23 float32 = (1e10 + 1e10)
                        _ = v22
                        _ = v23
                    }
                    var v24 int8 = 1
                    _ = v11
                    _ = v24
                } else {
                    var v25 uint64 = 2
                    _ = v25
                }
            }
        }
        fmt.Println(fi7, v3)
        _ = fi7
    }
    v3 += -790
    fmt.Println(v4, v3, v4)
    v26 := GenS2{F0: 2}
    if (!true) {
        switch 128 {
        case -159:
            v3 = 1000
        default:
            v5[v1] = v1
        }
    } else {
        fmt.Println(v4, v1)
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := GenS2{F0: 42}
            _ = len(v2)
            _ = fi27
            _ = v28
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            for fi30 := 0; fi30 < 19; fi30++ {
                if false {
                    var v31 uint32 = 2147483648
                    _ = v31
                } else {
                    v32 := []int{(fi30 / ((-1) | 1)), (fi29 * (^(fi30))), (fi29 << ((v4) & 63))}
                    switch fi29 {
                    case 530:
                        v4 = 18446744073709551615
                    }
                    if (!false) {
                        v1 = v1
                        var v33 int32 = 404
                        var v34 int8 = 0
                        v3 = (v3 % ((v3) | 1))
                        _ = v33
                        _ = v34
                    }
                    _ = v32
                }
                v1 = v1
                if (255 >= 128) {
                    if false {
                        var v35 uint32 = 597
                        if (2147483647 > -1) {
                            v36 := v2[v35]
                            v36 = (true && (v36 && v36))
                            v36 = v36
                            v37 := v5[v1]
                            _ = v36
                            _ = v37
                        } else {
                            v38 := GenS3{F0: 32767}
                            v39 := []int{fi29}
                            var v40 uint16 = 690
                            var v41 int = (fi30 ^ (fi29 + (fi30 | fi30)))
                            _ = v38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        _ = v35
                    } else {
                        var v42 int16 = 895
                        switch (fi30 >> ((v4) & 63)) {
                        case 9223372036854775807:
                            fmt.Println(v1)
                        default:
                            v43 := []int{(-(fi29)), (fi29 / ((fi29) | 1))}
                            _ = v43
                        }
                        _ = v42
                    }
                } else {
                    v44 := []int{}
                    _ = v44
                }
                _ = fi30
            }
            _ = len(v2)
            for fi45 := 0; fi45 < 9; fi45++ {
                var v46 float32 = -2.25
                var v47 uint8 = 127
                v48 := GenS1{F0: 304, F1: 65535}
                _ = fi45
                _ = v46
                _ = v47
                _ = v48
            }
            _ = fi29
        }
    }
    var v49 bool = true
    v50 := []int{-9223372036854775808, -57, -30, 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
