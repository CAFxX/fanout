package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint8 }

func main() {
    var v1 string = ("hello" + "\n\t")
    var v2 uint64 = 292
    var v3 uint32 = 51
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 65535
    var v6 uint8 = 0
    if false {
        fmt.Println(v2, v5, v1)
        v2 = v2
        v7 := GenS1{F0: ((("a" + v1) + "\n\t") + ("a" + (v1 + ((v1 + "foo bar") + "foo bar")))), F1: (v6 << ((255) & 63))}
        _ = v7
    } else {
        if (!((-2.25 * -2.25) > (0.1 + (-2.25 - 0.1)))) {
            fmt.Println(v1, v1)
            for fi8 := 0; fi8 < 18; fi8++ {
                v1 = v1
                _ = fi8
            }
            v9 := []int{127, 9223372036854775807, 32768, 42}
            _ = v9
        }
        var v10 uint = 883
        _ = v10
    }
    v3 += (v3 | v3)
    var v11 uint = 45
    if true {
        switch v1 {
        case "foo bar":
            if (2147483647 <= -598) {
                v1 = ("hello" + (("hello" + v1) + v1))
                for fi12 := 0; fi12 < 11; fi12++ {
                    v13 := GenS1{F0: v1, F1: uint8(v3)}
                    v3 = (v3 ^ (v3 / ((v3) | 1)))
                    switch v1 {
                    case "":
                        v2 += (v2 / ((v2) | 1))
                    case "foo bar":
                        var v14 int8 = 0
                        _ = v14
                    case "a":
                        switch "foo bar" {
                        case "":
                            var v15 int = (fi12 / (((-(fi12))) | 1))
                            _ = v15
                        case "hello":
                            v6 = 75
                        case "a":
                            if (!true) {
                                var v16 int8 = int8(v11)
                                v4[v3] = v3
                                for fi17 := 0; fi17 < 8; fi17++ {
                                    if (true && true) {
                                        v2 = 2
                                        var v18 uint = (^(v11))
                                        fmt.Println(v3)
                                        _ = v18
                                    } else {
                                        var v19 int8 = (v16 >> (((^(v11))) & 63))
                                        fmt.Println(v6)
                                        if false {
                                            v20 := []int{}
                                            v3 = v3
                                            _ = v20
                                        } else {
                                            v5 -= (v5 ^ (v5 - (v5 + (v5 & 0))))
                                            v6 = (v6 << (((^(v11))) & 63))
                                        }
                                        _ = v19
                                    }
                                    v11 -= (v11 | (v11 ^ v11))
                                    _ = fi17
                                }
                                _ = v16
                            }
                        }
                    }
                    _ = fi12
                    _ = v13
                }
                v21 := GenS1{F0: (v1 + v1), F1: 127}
                _ = v21
            } else {
                var v22 int8 = 42
                if ((42 < 0) && true) {
                    for fi23 := 0; fi23 < 13; fi23++ {
                        v24 := []int{(fi23 << ((0) & 63)), fi23}
                        if true {
                            fmt.Println(v5, v22)
                        }
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        _ = fi23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 4; fi25++ {
                        for fi26 := 0; fi26 < 4; fi26++ {
                            var v27 bool = true
                            switch (v1 + ("a" + v1)) {
                            case "a":
                                v2 += v2
                                var v28 uint16 = (^(v5))
                                _ = v28
                            case "hello":
                                if v27 {
                                    var v29 bool = (int8(v6) < v22)
                                    var v30 uint = 476
                                    _ = v29
                                    _ = v30
                                } else {
                                    if v27 {
                                        v31 := GenS1{F0: v1, F1: (v6 + (v6 & (v6 | (v6 + 128))))}
                                        v27 = true
                                        if v27 {
                                            var v32 float64 = 1e10
                                            v33 := GenS1{F0: "\n\t", F1: v6}
                                            v2 += v2
                                            delete(v4, v3)
                                            v3 = 4294967295
                                            _ = v32
                                            _ = v33
                                        } else {
                                            switch fi25 {
                                            case 32768:
                                                if false {
                                                    v34 := []int{(fi26 / (((-(fi26))) | 1)), fi26}
                                                    _ = v34
                                                }
                                            case 202:
                                                var v35 bool = v27
                                                v36 := GenS1{F0: "hello", F1: (^(v6))}
                                                _ = v35
                                                _ = v36
                                            case 127:
                                                v3 = (v3 << ((v11) & 63))
                                            }
                                        }
                                        _ = v31
                                    } else {
                                        var v37 float32 = (1.5 + 3.14159)
                                        _ = v37
                                    }
                                    v38 := []int{(fi25 * -434), 32768}
                                    _ = v38
                                }
                                v39 := GenS2{F0: (fi25 + fi25), F1: 1, F2: v22}
                                var v40 uint8 = (^(v6))
                                v1 = (("foo bar" + "a") + v1)
                                var v41 int32 = 42
                                var v42 int = fi26
                                _ = v39
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            v27 = true
                            _ = fi26
                            _ = v27
                        }
                        _ = fi25
                    }
                    var v43 int8 = -38
                    _ = v43
                }
                switch (v1 + v1) {
                case "\n\t":
                    var v44 int = -2147483648
                    _ = v44
                case "":
                    _ = len(v4)
                }
                _ = v22
            }
        case "a":
            v45 := GenS3{F0: false, F1: uint64(v3), F2: v6}
            _ = v45
        case "\n\t":
            v46 := GenS3{F0: (3.14159 == (3.14159 / 0.1)), F1: v2, F2: (v6 & (v6 + 255))}
            _ = v46
        default:
            switch (v1 + ((v1 + v1) + v1)) {
            case "\n\t":
                if true {
                    v47 := GenS2{F0: 278, F1: (^(v2)), F2: 58}
                    _ = v47
                }
            case "hello":
                var v48 int8 = 82
                _ = v48
            case "":
                if (2147483647 >= 1) {
                    var v49 uint16 = 2
                    _ = v49
                }
            default:
                switch v1 {
                case "a":
                    for fi50 := 0; fi50 < 14; fi50++ {
                        fmt.Println(v1, v3, v6)
                        _ = fi50
                    }
                case "hello":
                    v5 = 1
                default:
                    fmt.Println(v2, v2, v5)
                }
            }
        }
        var v51 uint8 = 0
        delete(v4, 611)
        _ = v51
    }
    v1 = (v1 + v1)
    var v52 string = v1
    fmt.Println(v3, v2)
    v53 := []int{}
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
}
