package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int16 }

func main() {
    var v1 float64 = float64(float32((0.0 * 1e10)))
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v5 float32 = 1e10
    var v6 uint8 = 127
    var v7 int = -748
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: 0}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v6 = 99
        _ = fi10
    }
    delete(v4, 1)
    switch v7 {
    case 0:
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 6; fi12++ {
                switch ("hello" + "") {
                case "foo bar":
                    if ((65535 <= 909) && (true || true)) {
                        if (42 >= 100) {
                            fmt.Println(v7)
                            v7 = (v7 / ((v7) | 1))
                            v13 := GenS2{F0: 0}
                            _ = v13
                        }
                        v7 += v7
                        v14 := []int{1, 127, (fi11 ^ 734)}
                        v15 := GenS1{F0: "foo bar", F1: 255, F2: 100}
                        var v16 uint16 = 65535
                        var v17 string = ("" + "\n\t")
                        if ((1 <= 255) || true) {
                            fmt.Println(v7)
                        } else {
                            if false {
                                switch (fi11 ^ 931) {
                                case -87:
                                    v4[241] = fi12
                                default:
                                    v5 -= 3.14159
                                }
                            } else {
                                v5 -= (v5 - (float32((float64(1e10) / 1.5)) / 0.1))
                            }
                            if true {
                                v14 = append(v14, 9223372036854775807)
                            } else {
                                if false {
                                    switch (v7 / (((fi11 % (((fi12 - fi12)) | 1))) | 1)) {
                                    case -541:
                                        v1 = (v1 - (v1 * v1))
                                    default:
                                        v6 = v6
                                    }
                                    v18 := GenS3{F0: -88, F1: 1000}
                                    _ = v18
                                } else {
                                    var v19 int = v7
                                    v20 := GenS2{F0: 923}
                                    _ = v19
                                    _ = v20
                                }
                            }
                        }
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        fmt.Println(v5)
                        v6 = v6
                        v7 = (fi12 & 32767)
                    }
                    fmt.Println(v5, fi11, fi12)
                    fmt.Println(fi12)
                case "":
                    _ = len(v4)
                }
                v1 = 3.14159
                v5 = 1.5
                _ = fi12
            }
            v3[""] = (v7 >> ((488) & 63))
            _ = fi11
        }
    case -446:
        fmt.Println(v7, v7, v6)
    default:
        var v21 string = ("" + "hello")
        _ = v21
    }
    v22 := GenS3{F0: -1, F1: -32768}
    if false {
        var v23 int16 = -766
        _ = v23
    } else {
        var v24 uint64 = 18446744073709551615
        v6 -= (v6 - v6)
        _ = v24
    }
    fmt.Println(v1)
    v25 := []int{}
    var v26 string = "hello"
    for fi27 := 0; fi27 < 2; fi27++ {
        v28 := GenS1{F0: v26, F1: 166, F2: 80}
        switch (v26 + v26) {
        case "foo bar":
            v1 = ((((v1 / v1) - v1) + v1) * (1.5 / 1.5))
        case "a":
            var v29 int = 128
            _ = v29
        case "\n\t":
            var v30 uint32 = 928
            _ = v30
        }
        fmt.Println(v1)
        _ = fi27
        _ = v28
    }
    v31 := GenS1{F0: "hello", F1: 1, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
