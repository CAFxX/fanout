package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    var v1 int32 = 839
    v2 := make(map[uint32]string)
    var v3 int64 = 0
    var v4 int8 = -128
    var v5 int8 = 0
    var v6 uint16 = 512
    fmt.Println(v6, v6, v1)
    switch 2147483647 {
    case 128:
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := v2[968]
            v9 := []int{(fi7 / (((fi7 + fi7)) | 1))}
            _ = fi7
            _ = v8
            _ = v9
        }
    case -9223372036854775808:
        switch "" {
        case "a":
            v10 := GenS1{F0: v3}
            _ = v10
        case "\n\t":
            _ = len(v2)
        default:
            for fi11 := 0; fi11 < 20; fi11++ {
                v1 -= v1
                _ = fi11
            }
        }
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            if (747 <= 18446744073709551615) {
                switch (fi12 - (fi12 ^ fi12)) {
                case 0:
                    for fi13 := 0; fi13 < 14; fi13++ {
                        v14 := GenS1{F0: (v3 | v3)}
                        v5 -= (v5 / ((v5) | 1))
                        v3 += 1
                        _ = fi13
                        _ = v14
                    }
                case -405:
                    v4 = v4
                case 961:
                    v5 -= (v4 ^ -115)
                }
            }
            _ = fi12
        }
    }
    if (32768 <= 42) {
        if true {
            for fi15 := 0; fi15 < 5; fi15++ {
                switch 32768 {
                case -149:
                    if true {
                        v16 := []int{(^(fi15)), int(v3), (fi15 >> ((0) & 63))}
                        if (2 == 2) {
                            var v17 float32 = (3.14159 - 0.0)
                            _ = v17
                        }
                        _ = v16
                    }
                case 1:
                    var v18 int = (-(fi15))
                    _ = v18
                case 4294967295:
                    if (fi15 <= (fi15 - (fi15 + (fi15 ^ fi15)))) {
                        for fi19 := 0; fi19 < 10; fi19++ {
                            var v20 uint64 = 681
                            fmt.Println(fi19)
                            v21 := GenS1{F0: v3}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    } else {
                        switch "hello" {
                        case "hello":
                            var v22 bool = true
                            _ = v22
                        case "foo bar":
                            v23 := GenS2{F0: 2, F1: fi15}
                            _ = v23
                        default:
                            fmt.Println(v6)
                        }
                        if (v6 < (v6 & v6)) {
                            var v24 int8 = -1
                            for fi25 := 0; fi25 < 1; fi25++ {
                                var v26 string = "foo bar"
                                switch (fi25 / (((fi25 - int(v24))) | 1)) {
                                case 127:
                                    v27 := GenS1{F0: 0}
                                    fmt.Println(v3, v1, v3)
                                    _ = v27
                                case 941:
                                    v6 = v6
                                default:
                                    v28 := v2[65535]
                                    if (1.5 <= 1.5) {
                                        v29 := GenS1{F0: (v3 - v3)}
                                        _ = v29
                                    }
                                    _ = v28
                                }
                                v5 = (v4 >> ((550) & 63))
                                v5 = (v5 * (-(v4)))
                                _ = fi25
                                _ = v26
                            }
                            v30 := v2[946]
                            _ = v24
                            _ = v30
                        } else {
                            switch 32768 {
                            case 820:
                                fmt.Println(v3, fi15, v6)
                            default:
                                v4 += (v4 + -128)
                            }
                            var v31 uint = 18446744073709551615
                            v4 = (v5 * v5)
                            _ = v31
                        }
                        v32 := []int{1, 624}
                        _ = v32
                    }
                }
                v5 += -128
                var v33 int64 = (v3 - v3)
                _ = fi15
                _ = v33
            }
            if false {
                var v34 int64 = (v3 % ((v3) | 1))
                fmt.Println(v5)
                _ = v34
            }
            var v35 uint16 = (^(v6))
            _ = v35
        } else {
            var v36 int32 = (v1 << ((2) & 63))
            v1 -= -1
            var v37 uint32 = 2
            _ = v36
            _ = v37
        }
        var v38 int8 = 0
        var v39 bool = true
        _ = v38
        _ = v39
    }
    var v40 int8 = (^(v4))
    var v41 uint32 = 4294967295
    for fi42 := 0; fi42 < 20; fi42++ {
        var v43 bool = (127 < 164)
        var v44 float32 = float32(((float64(1e10) / (3.14159 + 1.5)) - 0.0))
        _ = fi42
        _ = v43
        _ = v44
    }
    v2[v41] = ("hello" + "hello")
    v40 = int8(v1)
    v45 := GenS2{F0: v41, F1: -1}
    var v46 int64 = int64(v1)
    switch 2147483647 {
    case -1:
        v47 := v2[(v41 | v41)]
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
