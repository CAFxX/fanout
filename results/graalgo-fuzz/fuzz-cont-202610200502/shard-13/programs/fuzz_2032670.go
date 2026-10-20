package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = (("\n\t" > "\n\t") && false)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 0
    fmt.Println(v4, v4)
    _ = len(v3)
    if (v2 && (0 <= 0)) {
        v2 = (690 != 255)
        if v2 {
            fmt.Println(v4)
            v4 = v4
        } else {
            if (1 <= 42) {
                v4 -= v4
            }
            fmt.Println(v4, v2)
        }
    }
    if true {
        v4 = (v4 % ((v4) | 1))
        for fi5 := 0; fi5 < 13; fi5++ {
            switch (fi5 ^ 42) {
            case -9223372036854775808:
                v4 += (v4 ^ 0)
            case 4294967295:
                v2 = ((0.1 >= 0.0) && true)
            case 9223372036854775807:
                var v6 int16 = -1
                _ = v6
            }
            switch -289 {
            case 1:
                switch (fi5 * fi5) {
                case 4294967295:
                    var v7 float64 = ((((float64((1e10 + float32(0.1))) * float64((0.0 - 1.5))) + 0.1) / 1e10) + float64(0.1))
                    _ = v7
                case 32767:
                    var v8 uint16 = 933
                    _ = v8
                case 9223372036854775807:
                    v4 = (v4 ^ 100)
                }
            case 291:
                v9 := GenS1{F0: 0}
                _ = v9
            case -2147483648:
                fmt.Println(v2, fi5)
            }
            _ = fi5
        }
    } else {
        v2 = v2
    }
    v10 := []int{2147483647, 108}
    fmt.Println(v4)
    for fi11 := 0; fi11 < 10; fi11++ {
        switch "\n\t" {
        case "hello":
            var v12 int32 = -2147483648
            _ = v12
        case "":
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 bool = ((65535 == 32767) && (1 >= 42))
                _ = fi13
                _ = v14
            }
        }
        switch (^(fi11)) {
        case 42:
            switch ("hello" + "\n\t") {
            case "hello":
                switch ("" + ("\n\t" + ("" + "foo bar"))) {
                case "hello":
                    v15 := v1[(fi11 >> ((0) & 63))]
                    _ = v15
                case "":
                    if (-2147483648 != 42) {
                        if ((v4 * (v4 % ((v4) | 1))) < (v4 + -128)) {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            v2 = (((true || v2) && v2) && v2)
                            if v2 {
                                _ = len(v3)
                            }
                        }
                        fmt.Println(fi11, fi11, v4)
                        fmt.Println(v4, v4)
                    } else {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        if (!(0.0 != 3.14159)) {
                            var v16 int = 128
                            _ = v16
                        } else {
                            if v2 {
                                var v17 int32 = -1
                                if (!(v2 && (false || (4294967295 > 809)))) {
                                    v4 = (v4 << ((1) & 63))
                                    for fi18 := 0; fi18 < 5; fi18++ {
                                        _ = len(v3)
                                        _ = fi18
                                    }
                                    if v2 {
                                        var v19 float64 = -2.25
                                        _ = v19
                                    }
                                }
                                _ = v17
                            } else {
                                v20 := GenS2{F0: 857}
                                var v21 float64 = float64(0.1)
                                if (!false) {
                                    fmt.Println(v4)
                                    _ = len(v10)
                                    v22 := GenS2{F0: 2}
                                    _ = v22
                                } else {
                                    if (255 <= 2) {
                                        var v23 int64 = 0
                                        v3[778] = 1
                                        _ = v23
                                    }
                                    v24 := GenS3{F0: 65535, F1: -2147483648}
                                    for fi25 := 0; fi25 < 4; fi25++ {
                                        var v26 int = -2147483648
                                        v4 = 0
                                        v3[458] = 2147483648
                                        _ = fi25
                                        _ = v26
                                    }
                                    _ = v24
                                }
                                _ = v20
                                _ = v21
                            }
                        }
                    }
                default:
                    var v27 uint64 = 4294967295
                    _ = v27
                }
            case "foo bar":
                v2 = v2
            }
        case 0:
            if v2 {
                var v28 float32 = float32(1.5)
                var v29 float64 = 1e10
                v4 = v4
                _ = v28
                _ = v29
            }
        }
        _ = fi11
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v2 = (!((1 <= 1) || ((!(9223372036854775807 >= 42)) || v2)))
    v2 = ((v2 || (v4 != (-(v4)))) && (!((v4 != v4) && (((32768 <= 127) || false) || v2))))
    v30 := GenS3{F0: 245, F1: -920}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
