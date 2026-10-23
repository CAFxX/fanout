package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int32; F2 int8 }
type GenS7 struct { F0 int8 }

func main() {
    var v1 bool = false
    var v2 int16 = 42
    var v3 float32 = 0.0
    var v4 int32 = 543
    var v5 int16 = 32767
    var v6 int = 128
    v4 += v4
    v3 += v3
    fmt.Println(v3, v3, v6)
    v5 = 337
    fmt.Println(v6)
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: 127, F1: (^(v7))}
    var v9 int8 = 127
    v10 := GenS2{F0: "", F1: 32767}
    v9 -= v9
    if (((!true) || v1) && (18446744073709551615 < 255)) {
        for fi11 := 0; fi11 < 1; fi11++ {
            switch (^(fi11)) {
            case -9223372036854775808:
                for fi12 := 0; fi12 < 4; fi12++ {
                    fmt.Println(v3)
                    _ = fi12
                }
            case 1:
                if (42 > 2147483648) {
                    if v1 {
                        if (1 > 18446744073709551615) {
                            for fi13 := 0; fi13 < 20; fi13++ {
                                if v1 {
                                    var v14 string = ("hello" + ("\n\t" + ""))
                                    v15 := []int{}
                                    _ = v14
                                    _ = v15
                                }
                                var v16 float32 = 3.14159
                                _ = fi13
                                _ = v16
                            }
                            if true {
                                fmt.Println(v4)
                                if (!v1) {
                                    for fi17 := 0; fi17 < 15; fi17++ {
                                        v3 = v3
                                        if v1 {
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v18 bool = false
                                                _ = v18
                                            case "foo bar":
                                                v19 := []int{fi11}
                                                _ = v19
                                            }
                                            if (!(false || (v1 || ((42 > 42) && (((false && ((0 != v7) || (346 >= 255))) || v1) && ((true && false) && v1)))))) {
                                                fmt.Println(v7)
                                                v20 := GenS3{F0: "\n\t"}
                                                var v21 string = "\n\t"
                                                _ = v20
                                                _ = v21
                                            } else {
                                                v7 = (-(v7))
                                                switch -418 {
                                                case 4294967295:
                                                    if ((167 > 65535) || ((v6 / ((42) | 1)) <= 42)) {
                                                        var v22 uint32 = 436
                                                        _ = v22
                                                    } else {
                                                        v23 := []int{0, (v6 | (v6 * v6))}
                                                        _ = v23
                                                    }
                                                default:
                                                    v24 := []int{(fi17 >> ((255) & 63)), (fi11 | v6), (v6 ^ int(v7))}
                                                    _ = v24
                                                }
                                                v6 += fi17
                                            }
                                        }
                                        if (v1 || v1) {
                                            var v25 float32 = float32(-2.25)
                                            var v26 int64 = (v7 + v7)
                                            _ = v25
                                            _ = v26
                                        } else {
                                            switch (v6 << ((80) & 63)) {
                                            case -820:
                                                v27 := GenS4{F0: -9223372036854775808}
                                                _ = v27
                                            default:
                                                v28 := GenS5{F0: 32767}
                                                _ = v28
                                            }
                                            if (2 != 970) {
                                                v7 -= (v7 << ((987) & 63))
                                            }
                                        }
                                        _ = fi17
                                    }
                                }
                                v5 += 0
                            } else {
                                v6 += fi11
                                v9 = (v9 / (((-(v9))) | 1))
                            }
                        } else {
                            var v29 float32 = v3
                            for fi30 := 0; fi30 < 20; fi30++ {
                                v1 = true
                                if (v1 && v1) {
                                    v4 = v4
                                } else {
                                    v31 := GenS1{F0: 75, F1: (v7 << ((4294967295) & 63))}
                                    v32 := []int{(fi30 % ((int(v7)) | 1)), fi30, v6, (fi11 - v6)}
                                    v1 = false
                                    v33 := GenS4{F0: v7}
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                fmt.Println(fi30, v4, v5)
                                _ = fi30
                            }
                            var v34 bool = v1
                            _ = v29
                            _ = v34
                        }
                        v5 = (-(v5))
                    } else {
                        v7 += -1
                        if (255 != 127) {
                            for fi35 := 0; fi35 < 8; fi35++ {
                                v36 := GenS6{F0: 42, F1: (v4 ^ -1), F2: (v9 >> ((42) & 63))}
                                fmt.Println(v1, v9, fi11)
                                v37 := []int{fi11, v6}
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            v3 -= (float32(-2.25) * (1e10 + v3))
                        } else {
                            v38 := GenS2{F0: "foo bar", F1: 1}
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 1; fi39++ {
                            var v40 uint64 = 1
                            v41 := GenS5{F0: 65535}
                            v42 := []int{v6, int(v40), (v6 * fi39)}
                            _ = fi39
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                    }
                    var v43 int16 = 429
                    fmt.Println(v9, v2)
                    _ = v43
                } else {
                    for fi44 := 0; fi44 < 19; fi44++ {
                        v45 := GenS7{F0: (v9 >> ((42) & 63))}
                        _ = fi44
                        _ = v45
                    }
                }
            case 2147483647:
                var v46 uint64 = 643
                _ = v46
            }
            var v47 bool = (v1 && (true && false))
            _ = fi11
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 19; fi48++ {
        var v49 float32 = v3
        v3 += 0.0
        v3 -= v49
        _ = fi48
        _ = v49
    }
    var v50 int = v6
    fmt.Println(v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
