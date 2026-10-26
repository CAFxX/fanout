package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint; F2 int16 }
type GenS4 struct { F0 int32; F1 uint64; F2 string }
type GenS5 struct { F0 int32; F1 int8; F2 uint16 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{}
    var v3 int = 0
    var v4 float32 = 0.0
    v5 := make(map[int]bool)
    var v6 bool = (!(!(!true)))
    v4 = ((v4 - 1e10) - v4)
    var v7 uint = 1
    var v8 bool = (65535 == 2)
    fmt.Println(v7, v4)
    fmt.Println(v3, v8)
    v9 := v5[(v3 % (((v3 << ((2) & 63))) | 1))]
    if v9 {
        fmt.Println(v3, v8)
        v10 := GenS1{F0: 0, F1: (!false)}
        fmt.Println(v9, v4)
        _ = v10
    }
    v2 = append(v2, 324)
    if true {
        var v11 float32 = v4
        var v12 uint8 = 185
        _ = v11
        _ = v12
    }
    if v8 {
        if ((255 == 2) && (0 <= 127)) {
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 int = v3
                _ = fi13
                _ = v14
            }
            var v15 uint = (v7 | (^(v7)))
            _ = v15
        } else {
            for fi16 := 0; fi16 < 10; fi16++ {
                v17 := GenS2{F0: 716, F1: 2147483647}
                for fi18 := 0; fi18 < 7; fi18++ {
                    v6 = (42 < 1)
                    v19 := GenS3{F0: v1, F1: 544, F2: -32768}
                    var v20 int16 = 1000
                    v4 = v4
                    var v21 uint8 = 128
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = fi16
                _ = v17
            }
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            v5[4294967295] = (!v8)
        }
        v2 = append(v2, 2147483647)
        v4 = v4
    } else {
        v22 := GenS4{F0: -1, F1: v1, F2: "a"}
        v4 = (1.5 / ((v4 - v4) * (float32(1e10) + -2.25)))
        var v23 float32 = 0.1
        _ = v22
        _ = v23
    }
    if false {
        for fi24 := 0; fi24 < 17; fi24++ {
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 uint = (^(v7))
                _ = fi25
                _ = v26
            }
            var v27 string = "\n\t"
            if v9 {
                for fi28 := 0; fi28 < 12; fi28++ {
                    v9 = (v8 && v9)
                    v2 = append(v2, 42)
                    if v6 {
                        var v29 uint16 = 1
                        v4 = -2.25
                        if v9 {
                            var v30 int = fi24
                            v29 += 65535
                            v3 += v30
                            v31 := v5[fi24]
                            var v32 bool = (false && (1 != 665))
                            if v32 {
                                v32 = (1e10 <= (0.1 / -2.25))
                            } else {
                                v33 := GenS5{F0: -2147483648, F1: 127, F2: v29}
                                _ = v33
                            }
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        _ = v29
                    } else {
                        delete(v5, (fi24 ^ (fi28 & v3)))
                        switch fi28 {
                        case 32768:
                            fmt.Println(v6, v7, v8)
                        case -700:
                            v34 := GenS2{F0: 255, F1: 318}
                            _ = v34
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi28
                }
            } else {
                if v9 {
                    for fi35 := 0; fi35 < 14; fi35++ {
                        fmt.Println(fi24, v1, v1)
                        _ = fi35
                    }
                }
                if (42 < -71) {
                    var v36 int32 = 1
                    v1 -= 4294967295
                    _ = v36
                }
                var v37 bool = false
                _ = v37
            }
            _ = fi24
            _ = v27
        }
    }
    var v38 string = (((("foo bar" + (("foo bar" + "\n\t") + "")) + (("" + "hello") + "hello")) + "foo bar") + "\n\t")
    fmt.Println(v6, v3)
    v8 = v9
    v38 = ("a" + "\n\t")
    var v39 uint = v7
    for fi40 := 0; fi40 < 14; fi40++ {
        v9 = (v6 && v9)
        _ = fi40
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
