package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 uint16 = 65535
    v2 := []string{"\n\t", "hello", "\n\t", (("foo bar" + (((("foo bar" + "\n\t") + "hello") + "hello") + (("foo bar" + "foo bar") + "hello"))) + "hello")}
    var v3 float32 = 0.0
    v4 := make(map[uint32]uint32)
    var v5 int64 = 1
    var v6 bool = true
    v1 = v1
    var v7 uint64 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v5, v3, v5)
    fmt.Println(v5)
    var v8 uint16 = 610
    switch "foo bar" {
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        var v9 float32 = 1.5
        _ = v9
    case "hello":
        for fi10 := 0; fi10 < 16; fi10++ {
            if (2147483647 == fi10) {
                var v11 bool = v6
                _ = v11
            }
            v7 = (^(v7))
            _ = fi10
        }
    }
    if true {
        if v6 {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v7 = v7
            fmt.Println(v7)
        }
        var v12 int64 = (v5 | (v5 >> ((18446744073709551615) & 63)))
        var v13 uint = 255
        _ = v12
        _ = v13
    } else {
        var v14 uint = 255
        for fi15 := 0; fi15 < 10; fi15++ {
            switch (fi15 + fi15) {
            case 0:
                _ = len(v4)
            case -9223372036854775808:
                v16 := GenS1{F0: 42, F1: (v7 * v7), F2: 334}
                _ = v16
            }
            if v6 {
                if (((v5 | v5) < (v5 | (v5 * (v5 * (-(v5)))))) && true) {
                    v17 := GenS1{F0: -1, F1: 2, F2: 255}
                    fmt.Println(v6, v8, v8)
                    _ = v17
                }
                v5 = v5
            } else {
                for fi18 := 0; fi18 < 15; fi18++ {
                    v5 -= 1
                    var v19 uint16 = 255
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 9; fi20++ {
                    v21 := GenS1{F0: 0, F1: v7, F2: v1}
                    if v6 {
                        var v22 int64 = (v5 & v5)
                        _ = v22
                    }
                    v1 = 32767
                    switch 9223372036854775807 {
                    case 1:
                        var v23 uint = v14
                        _ = v23
                    case 216:
                        v7 += (v7 + v7)
                    }
                    fmt.Println(v8, v1, v3)
                    v7 = v7
                    v5 -= v5
                    _ = fi20
                    _ = v21
                }
                if v6 {
                    if (127 > uint8(v7)) {
                        for fi24 := 0; fi24 < 20; fi24++ {
                            switch fi24 {
                            case 4294967295:
                                var v25 int64 = 246
                                _ = v25
                            case 867:
                                if len(v2) > 0 {
                                    v2[0] = ("a" + "foo bar")
                                }
                            }
                            _ = fi24
                        }
                        var v26 float64 = 0.1
                        var v27 float64 = v26
                        _ = v26
                        _ = v27
                    }
                    if (!((2 < 1) && true)) {
                        if (!v6) {
                            v5 -= (v5 | v5)
                            if (((32767 == 42) && true) || true) {
                                switch ("" + ("a" + "foo bar")) {
                                case "\n\t":
                                    fmt.Println(v5)
                                case "foo bar":
                                    v28 := GenS2{F0: v6, F1: v5}
                                    _ = v28
                                }
                                v5 = (v5 * v5)
                                v3 = float32((0.1 * 0.1))
                            } else {
                                v29 := GenS2{F0: (v5 == (v5 / (((v5 - v5)) | 1))), F1: -9223372036854775808}
                                _ = v29
                            }
                            if true {
                                var v30 int16 = -1
                                _ = v30
                            }
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v31 uint32 = 4294967295
                        _ = v31
                    }
                    v32 := GenS2{F0: ((v14 % (((v14 & uint(v7))) | 1)) < v14), F1: int64(fi15)}
                    _ = v32
                }
            }
            for fi33 := 0; fi33 < 16; fi33++ {
                v34 := GenS1{F0: 127, F1: v7, F2: (v8 - (^(v1)))}
                v6 = v6
                var v35 int16 = 32767
                fmt.Println(fi15, v1, fi15)
                v36 := GenS1{F0: -1, F1: (v7 / (((v7 % ((18446744073709551615) | 1))) | 1)), F2: 255}
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = fi15
        }
        v37 := GenS3{F0: 388}
        _ = v14
        _ = v37
    }
    v5 = (v5 >> ((1) & 63))
    v38 := GenS1{F0: 42, F1: v7, F2: v8}
    v7 = 2
    var v39 uint16 = v1
    for fi40 := 0; fi40 < 5; fi40++ {
        var v41 uint32 = 788
        v2 = append(v2, "foo bar")
        _ = fi40
        _ = v41
    }
    v7 = v7
    _ = len(v4)
    for fi42 := 0; fi42 < 3; fi42++ {
        v6 = v6
        fmt.Println(v5, v8, v5)
        v5 += (v5 << ((455) & 63))
        _ = fi42
    }
    v7 = (v7 + (v7 - (v7 % ((0) | 1))))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
}
