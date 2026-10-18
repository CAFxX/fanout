package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 bool }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int }

func main() {
    var v1 int32 = -2147483648
    v2 := []int{-847, -592}
    var v3 uint = 2
    v4 := []int{}
    for fi5 := 0; fi5 < 1; fi5++ {
        switch "a" {
        case "hello":
            if (false || true) {
                var v6 int64 = -912
                var v7 uint8 = 255
                switch fi5 {
                case 494:
                    var v8 int64 = (v6 & (v6 & v6))
                    _ = v8
                case 1:
                    v9 := GenS1{F0: 2147483647, F1: "hello", F2: (true || false)}
                    _ = v9
                case 0:
                    if true {
                        if (42 >= 255) {
                            var v10 uint = (v3 ^ (v3 * v3))
                            if (543 >= 291) {
                                if ((fi5 | (^(fi5))) >= fi5) {
                                    v1 = v1
                                    v7 += v7
                                    v6 = v6
                                }
                                var v11 int32 = v1
                                fmt.Println(v7, v10, v11)
                                _ = v11
                            }
                            fmt.Println(v3, v3)
                            _ = v10
                        } else {
                            if (!false) {
                                if (!(!false)) {
                                    for fi12 := 0; fi12 < 20; fi12++ {
                                        _ = len(v2)
                                        _ = fi12
                                    }
                                    _ = len(v4)
                                }
                                v3 -= 4294967295
                                if (127 < 100) {
                                    switch (-(fi5)) {
                                    case 128:
                                        var v13 string = "a"
                                        _ = v13
                                    case -48:
                                        var v14 uint64 = 4294967295
                                        _ = v14
                                    }
                                    switch (fi5 % ((fi5) | 1)) {
                                    case 4294967295:
                                        if false {
                                            var v15 uint16 = 0
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v15
                                        }
                                    case -509:
                                        var v16 string = "foo bar"
                                        _ = v16
                                    }
                                }
                            } else {
                                var v17 uint16 = 32767
                                if len(v2) > 0 {
                                    v2[0] = -9223372036854775808
                                }
                                _ = v17
                            }
                            switch "a" {
                            case "":
                                fmt.Println(v3, v7, v6)
                            case "hello":
                                var v18 uint = v3
                                _ = v18
                            case "foo bar":
                                if true {
                                    v19 := GenS2{F0: 42, F1: -128}
                                    var v20 int64 = v6
                                    for fi21 := 0; fi21 < 20; fi21++ {
                                        var v22 float32 = float32(1e10)
                                        _ = fi21
                                        _ = v22
                                    }
                                    _ = v19
                                    _ = v20
                                }
                            }
                        }
                        _ = len(v2)
                        v6 = v6
                    } else {
                        v3 = 719
                        v2 = append(v2, 127)
                    }
                }
                _ = v6
                _ = v7
            } else {
                for fi23 := 0; fi23 < 2; fi23++ {
                    v24 := GenS2{F0: 1, F1: -30}
                    if true {
                        var v25 float32 = 3.14159
                        v26 := GenS3{F0: 346}
                        _ = v25
                        _ = v26
                    } else {
                        var v27 int32 = (v1 % ((v1) | 1))
                        fmt.Println(v3)
                        v28 := GenS4{F0: -1, F1: (-(fi5))}
                        _ = v27
                        _ = v28
                    }
                    _ = fi23
                    _ = v24
                }
            }
        case "\n\t":
            fmt.Println(v3, v3)
        case "":
            var v29 bool = true
            _ = v29
        default:
            var v30 float64 = (0.0 - (1.5 - float64(float32(((((1e10 / 1.5) * 1e10) - 1.5) * -2.25)))))
            _ = v30
        }
        for fi31 := 0; fi31 < 17; fi31++ {
            v32 := GenS2{F0: 127, F1: 100}
            _ = fi31
            _ = v32
        }
        var v33 int8 = 1
        _ = fi5
        _ = v33
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    if (-928 >= 32767) {
        v2 = append(v2, -498)
    } else {
        if false {
            var v34 int16 = 41
            _ = v34
        }
        fmt.Println(v3)
    }
    v1 += 2147483647
    v4 = append(v4, -1)
    var v35 uint = (v3 << ((v3) & 63))
    var v36 uint = v35
    var v37 uint8 = 128
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
