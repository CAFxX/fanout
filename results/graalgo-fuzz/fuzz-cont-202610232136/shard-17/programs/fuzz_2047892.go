package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{4294967295}
    v3 := []uint32{2, 1}
    v4 := []int{2147483647, -699, 1, 2147483647}
    v5 := make(map[int]bool)
    var v6 int = -9223372036854775808
    v7 := GenS1{F0: "hello"}
    v8 := GenS2{F0: -1}
    if (true && (v6 < (v6 << ((555) & 63)))) {
        v6 += (-(v6))
        if (32767 >= 0) {
            if true {
                if false {
                    if (1.5 <= (0.0 + 1.5)) {
                        v6 = (v6 & -2147483648)
                        if (-128 <= -46) {
                            if (0 != -304) {
                                if false {
                                    v9 := GenS3{F0: 312, F1: 2}
                                    v10 := v1[(v6 * (v6 & 32768))]
                                    _ = v10
                                    _ = v9
                                }
                                if true {
                                    v11 := GenS4{F0: 2}
                                    if false {
                                        var v12 int8 = 127
                                        for fi13 := 0; fi13 < 13; fi13++ {
                                            var v14 int64 = 0
                                            if (2147483648 < 47) {
                                                v12 = v12
                                                fmt.Println(v6)
                                                for fi15 := 0; fi15 < 14; fi15++ {
                                                    v16 := GenS2{F0: 1}
                                                    var v17 uint64 = 2
                                                    _ = fi15
                                                    _ = v16
                                                    _ = v17
                                                }
                                            }
                                            v18 := GenS5{F0: -826, F1: v14}
                                            _ = fi13
                                            _ = v14
                                            _ = v18
                                        }
                                        _ = v12
                                    } else {
                                        _ = len(v3)
                                    }
                                    _ = v11
                                }
                                v6 = (v6 >> ((1) & 63))
                            }
                            var v19 int64 = -9223372036854775808
                            var v20 int32 = 2147483647
                            _ = v19
                            _ = v20
                        }
                    }
                }
                v6 = v6
            } else {
                v21 := GenS2{F0: -853}
                _ = len(v1)
                _ = v21
            }
        } else {
            if false {
                var v22 int16 = 42
                var v23 int = 935
                for fi24 := 0; fi24 < 9; fi24++ {
                    for fi25 := 0; fi25 < 16; fi25++ {
                        var v26 int32 = 2147483647
                        _ = len(v1)
                        v2 = append(v2, 2147483647)
                        _ = fi25
                        _ = v26
                    }
                    _ = fi24
                }
                _ = v22
                _ = v23
            } else {
                var v27 int16 = 1000
                v6 = (v6 * 972)
                _ = v27
            }
            var v28 int16 = 1
            _ = v28
        }
        delete(v5, (v6 | (^(v6))))
    }
    var v29 string = ("\n\t" + "\n\t")
    for fi30 := 0; fi30 < 13; fi30++ {
        v6 -= fi30
        v31 := GenS4{F0: 255}
        _ = fi30
        _ = v31
    }
    if (true && (false && (-1 > -72))) {
        var v32 int = (^(v6))
        switch (v32 + (v32 * (^(v6)))) {
        case 9223372036854775807:
            v6 = (v32 % ((v6) | 1))
        default:
            var v33 uint = 875
            _ = v33
        }
        _ = v32
    }
    fmt.Println(v6)
    switch v29 {
    case "\n\t":
        if true {
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 int16 = 385
                var v36 uint32 = 0
                var v37 float64 = 0.1
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
        }
    case "hello":
        var v38 uint64 = 0
        _ = v38
    default:
        if false {
            switch (v29 + "\n\t") {
            case "hello":
                switch v29 {
                case "foo bar":
                    var v39 int = v6
                    _ = v39
                case "\n\t":
                    for fi40 := 0; fi40 < 11; fi40++ {
                        var v41 uint = 4294967295
                        _ = fi40
                        _ = v41
                    }
                case "hello":
                    v42 := GenS6{F0: 0}
                    _ = v42
                }
            }
        }
    }
    _ = len(v1)
    if ((-(v6)) >= (v6 & (v6 % (((v6 ^ (v6 << ((637) & 63)))) | 1)))) {
        v43 := v1[(v6 & v6)]
        _ = v43
    }
    v6 += (v6 % (((v6 - (v6 / (((v6 * v6)) | 1)))) | 1))
    switch ("\n\t" + v29) {
    case "a":
        if true {
            for fi44 := 0; fi44 < 2; fi44++ {
                var v45 uint32 = 2147483648
                _ = fi44
                _ = v45
            }
            v29 = ((("" + (v29 + v29)) + "foo bar") + v29)
        }
    default:
        if (true || true) {
            var v46 int = -514
            _ = v46
        } else {
            if (!(!true)) {
                v6 -= 42
                v6 += v6
            }
        }
    }
    var v47 uint32 = 509
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
