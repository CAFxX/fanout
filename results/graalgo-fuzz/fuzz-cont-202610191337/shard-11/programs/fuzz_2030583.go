package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }

func main() {
    var v1 int64 = 320
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    delete(v2, 0)
    var v4 int = -9223372036854775808
    v1 += (v1 % ((v1) | 1))
    delete(v2, 421)
    var v5 int8 = 0
    switch (-(v4)) {
    case -9223372036854775808:
        var v6 int64 = (v1 << ((494) & 63))
        _ = v6
    case 9:
        v7 := GenS1{F0: 32767, F1: 0, F2: ""}
        _ = v7
    }
    if (!((v5 > (^(v5))) && true)) {
        for fi8 := 0; fi8 < 14; fi8++ {
            fmt.Println(v5, v1)
            _ = fi8
        }
        var v9 uint = 255
        _ = v9
    }
    switch v4 {
    case 56:
        if false {
            fmt.Println(v5)
            if false {
                for fi10 := 0; fi10 < 6; fi10++ {
                    if (((v5 << ((193) & 63)) == (v5 * (v5 % ((v5) | 1)))) && true) {
                        switch (("\n\t" + "foo bar") + "foo bar") {
                        case "foo bar":
                            if false {
                                fmt.Println(v4)
                                if true {
                                    if (true && (526 > 0)) {
                                        if true {
                                            v11 := GenS2{F0: -2147483648}
                                            _ = v11
                                        }
                                        var v12 float32 = 3.14159
                                        _ = v12
                                    }
                                    v13 := GenS1{F0: 42, F1: 1, F2: "foo bar"}
                                    if ((32767 <= 1) && (65 == 255)) {
                                        fmt.Println(v4, fi10)
                                        var v14 bool = false
                                        _ = v14
                                    }
                                    _ = v13
                                }
                            } else {
                                for fi15 := 0; fi15 < 6; fi15++ {
                                    v16 := GenS3{F0: 1}
                                    var v17 int16 = 562
                                    v17 = v17
                                    v18 := GenS2{F0: 42}
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                            }
                        case "\n\t":
                            var v19 uint = 255
                            _ = v19
                        case "":
                            v5 = (v5 + -85)
                        }
                        fmt.Println(v5, fi10)
                    }
                    var v20 uint = 4294967295
                    v21 := GenS3{F0: -945}
                    _ = fi10
                    _ = v20
                    _ = v21
                }
            } else {
                v22 := GenS2{F0: 1}
                v23 := []int{(v4 >> ((4294967295) & 63))}
                _ = v22
                _ = v23
            }
        }
    }
    if false {
        for fi24 := 0; fi24 < 10; fi24++ {
            if (!true) {
                fmt.Println(fi24, v5)
                v4 += 128
            } else {
                fmt.Println(v1)
            }
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 bool = (!(false || false))
                fmt.Println(fi25)
                _ = fi25
                _ = v26
            }
            var v27 bool = false
            _ = fi24
            _ = v27
        }
        v28 := GenS3{F0: 918}
        if (!((-2.25 + -2.25) >= 0.1)) {
            v1 -= (v1 >> ((594) & 63))
            for fi29 := 0; fi29 < 14; fi29++ {
                fmt.Println(v4, v1, fi29)
                _ = fi29
            }
        }
        _ = v28
    } else {
        for fi30 := 0; fi30 < 1; fi30++ {
            v31 := GenS4{F0: 1}
            _ = fi30
            _ = v31
        }
        if false {
            var v32 uint32 = 65535
            v33 := GenS2{F0: 1}
            var v34 int8 = (-(v5))
            _ = v32
            _ = v33
            _ = v34
        }
        v35 := []int{127, 32767, v4, (v4 * v4)}
        _ = v35
    }
    if (!true) {
        v36 := GenS5{F0: (v4 / (((v4 & 969)) | 1))}
        _ = v36
    } else {
        switch 32767 {
        case 127:
            var v37 uint = 18446744073709551615
            _ = v37
        case -2147483648:
            for fi38 := 0; fi38 < 4; fi38++ {
                v4 -= (v4 * v4)
                _ = fi38
            }
        default:
            v4 -= (v4 - -1)
        }
        var v39 int16 = 1000
        _ = v39
    }
    v4 = (v4 * (v4 | (-(v4))))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
