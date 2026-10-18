package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int64; F1 int16 }

func main() {
    var v1 int64 = -1
    var v2 bool = (-35 < 0)
    var v3 string = "foo bar"
    v4 := []int{14, 32768, int(v1)}
    if (-951 != 32768) {
        if false {
            if (!v2) {
                v5 := GenS1{F0: 42}
                fmt.Println(v2)
                var v6 int32 = -672
                _ = v5
                _ = v6
            } else {
                var v7 int8 = -128
                _ = v7
            }
            v8 := GenS1{F0: v1}
            _ = v8
        }
        v9 := GenS1{F0: 86}
        switch -387 {
        case 128:
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 bool = true
                switch v3 {
                case "":
                    if (1 > 601) {
                        var v12 uint = 0
                        _ = len(v4)
                        switch int(v1) {
                        case 4294967295:
                            for fi13 := 0; fi13 < 18; fi13++ {
                                _ = len(v4)
                                var v14 uint = (v12 & v12)
                                switch fi10 {
                                case -285:
                                    v15 := GenS1{F0: v1}
                                    _ = v15
                                case 1:
                                    v2 = ((fi13 & (fi10 << ((v12) & 63))) <= 9223372036854775807)
                                case 383:
                                    v16 := GenS1{F0: -1}
                                    _ = v16
                                }
                                _ = fi13
                                _ = v14
                            }
                        default:
                            if v2 {
                                var v17 bool = false
                                _ = v17
                            } else {
                                switch "a" {
                                case "":
                                    var v18 float32 = (1.5 * (0.1 + float32((3.14159 - (3.14159 / -2.25)))))
                                    _ = v18
                                default:
                                    fmt.Println(fi10)
                                }
                            }
                        }
                        _ = v12
                    }
                case "a":
                    fmt.Println(fi10)
                case "foo bar":
                    v19 := GenS1{F0: (-(v1))}
                    _ = v19
                default:
                    _ = len(v4)
                }
                if ((("foo bar" + v3) != ("" + ((("a" + (("a" + v3) + "")) + ((("\n\t" + (v3 + v3)) + (v3 + "\n\t")) + v3)) + v3))) && v2) {
                    _ = len(v4)
                    if (-9223372036854775808 != (^(v1))) {
                        var v20 string = "hello"
                        _ = v20
                    } else {
                        var v21 float32 = 1e10
                        var v22 uint = 65535
                        _ = v21
                        _ = v22
                    }
                }
                _ = fi10
                _ = v11
            }
        case 4294967295:
            v2 = v2
        case -9223372036854775808:
            var v23 string = (v3 + v3)
            _ = v23
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = v9
    }
    if (!((127 != 0) || v2)) {
        var v24 uint8 = 0
        for fi25 := 0; fi25 < 7; fi25++ {
            v26 := GenS2{F0: v3}
            _ = fi25
            _ = v26
        }
        v24 += v24
        _ = v24
    } else {
        v2 = ((false || (!v2)) || true)
        var v27 int32 = 1
        _ = v27
    }
    if v2 {
        var v28 int8 = 1
        v2 = v2
        fmt.Println(v3, v1)
        _ = v28
    }
    var v29 float32 = (0.1 / (float32(0.1) * 1.5))
    for fi30 := 0; fi30 < 3; fi30++ {
        if v2 {
            v29 = float32((-2.25 / 1.5))
            var v31 uint32 = 65535
            var v32 int8 = -30
            _ = v31
            _ = v32
        }
        _ = fi30
    }
    var v33 float32 = v29
    if len(v4) > 0 {
        v4[0] = 217
    }
    v34 := GenS3{F0: 42, F1: 643}
    v35 := GenS4{F0: (v1 % (((^(v1))) | 1)), F1: 1000}
    var v36 bool = v2
    for fi37 := 0; fi37 < 17; fi37++ {
        var v38 bool = (!(458 <= 65535))
        var v39 uint32 = 358
        switch (fi37 >> ((4294967295) & 63)) {
        case -9223372036854775808:
            var v40 float64 = 3.14159
            _ = v40
        case -1:
            v41 := GenS3{F0: (fi37 ^ -388), F1: fi37}
            _ = v41
        case 1:
            var v42 string = v3
            _ = v42
        default:
            if ((v36 && v36) && v38) {
                v4 = append(v4, fi37)
            }
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
