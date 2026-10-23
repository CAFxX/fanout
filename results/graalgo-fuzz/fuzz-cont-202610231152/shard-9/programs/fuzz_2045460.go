package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 bool; F1 int8; F2 int }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 int64 }

func main() {
    var v1 int64 = -1
    var v2 int32 = 556
    var v3 uint32 = 42
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    v6 := v5["hello"]
    v7 := GenS1{F0: 1, F1: 127}
    var v8 uint8 = 162
    v6 = (("hello" + (v6 + ("hello" + "a"))) + "")
    var v9 string = v6
    v8 += (v8 + (v8 | (v8 ^ v8)))
    v10 := GenS2{F0: false, F1: 42, F2: 438}
    v6 = (v6 + ("foo bar" + v6))
    if false {
        v11 := GenS3{F0: int16(v3), F1: 42}
        _ = v11
    }
    v12 := []int{1, 32768, -1, 127}
    if true {
        v4[814] = true
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            if true {
                var v14 uint16 = 0
                for fi15 := 0; fi15 < 10; fi15++ {
                    if ((v9 != v6) && true) {
                        var v16 float64 = 0.0
                        var v17 int32 = v2
                        v18 := GenS3{F0: 42, F1: -1}
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        v12 = append(v12, (fi15 % ((32768) | 1)))
                        var v19 int64 = (v1 * (v1 & v1))
                        if (!(1 <= -128)) {
                            var v20 bool = ((3.14159 - (-2.25 / 1.0)) < 0.0)
                            _ = v20
                        } else {
                            v21 := GenS4{F0: ((fi15 - (fi13 >> ((42) & 63))) < fi13), F1: 38}
                            fmt.Println(fi13, v1, v8)
                            _ = v21
                        }
                        _ = v19
                    }
                    delete(v5, ((v6 + (v9 + v9)) + "hello"))
                    _ = fi15
                }
                var v22 float32 = 0.1
                _ = v14
                _ = v22
            } else {
                if false {
                    switch int(v3) {
                    case 42:
                        for fi23 := 0; fi23 < 7; fi23++ {
                            v24 := GenS3{F0: 32767, F1: int16(fi23)}
                            _ = fi23
                            _ = v24
                        }
                    }
                    if false {
                        v3 -= v3
                    }
                }
            }
            _ = fi13
        }
    }
    if ((v2 + (v2 % ((1) | 1))) == -1) {
        v25 := GenS4{F0: ((!(43 == 0)) && true), F1: 0}
        for fi26 := 0; fi26 < 2; fi26++ {
            switch ((((v9 + v6) + "foo bar") + v9) + v9) {
            case "a":
                if true {
                    if true {
                        var v27 int8 = -1
                        fmt.Println(v8)
                        v6 = v9
                        _ = v27
                    } else {
                        v3 += v3
                        if (42 == 32767) {
                            fmt.Println(v3, v3, v1)
                            var v28 bool = (!(-2.25 >= (1e10 + (float32((1.5 * float64(float32((1e10 + ((-2.25 / 3.14159) / (((1e10 / (0.1 / 1.5)) + 0.0) * 1e10))))))) * float32(float64(0.0))))))
                            _ = v28
                        }
                    }
                    v29 := GenS2{F0: (739 <= 4294967295), F1: 100, F2: -409}
                    v2 += v2
                    _ = v29
                } else {
                    v8 = v8
                    v1 += v1
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
            case "\n\t":
                if false {
                    v9 = "\n\t"
                }
            case "foo bar":
                fmt.Println(v9)
            default:
                _ = len(v4)
            }
            v6 = v6
            switch int(v3) {
            case 1:
                var v30 int16 = -935
                _ = v30
            case 127:
                for fi31 := 0; fi31 < 17; fi31++ {
                    var v32 uint8 = 0
                    _ = fi31
                    _ = v32
                }
            case 2147483647:
                v33 := GenS2{F0: ((-438 <= 32767) && ((1e10 / 1.0) != (((1e10 * 1e10) * (1.5 * ((float32(1.5) * (float32((-2.25 - float64(0.0))) * float32(float64(0.0)))) - ((1e10 * -2.25) + (0.0 - ((float32(0.0) + 3.14159) / (3.14159 / 3.14159))))))) - (1e10 + 0.0)))), F1: 56, F2: fi26}
                _ = v33
            }
            _ = fi26
        }
        var v34 int8 = -106
        _ = v25
        _ = v34
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v9)
    case "hello":
        for fi35 := 0; fi35 < 15; fi35++ {
            _ = len(v12)
            switch "hello" {
            case "foo bar":
                switch v6 {
                case "hello":
                    v9 = "\n\t"
                case "a":
                    v36 := GenS5{F0: v3}
                    _ = v36
                default:
                    v1 += v1
                }
            }
            if (!(159 != 42)) {
                v37 := GenS6{F0: 609, F1: (v1 | v1)}
                var v38 uint64 = 2
                _ = v37
                _ = v38
            } else {
                for fi39 := 0; fi39 < 20; fi39++ {
                    fmt.Println(v6)
                    v8 = (v8 * (v8 % (((v8 >> ((4294967295) & 63))) | 1)))
                    fmt.Println(v6, v8, v2)
                    _ = fi39
                }
                v1 -= 831
                var v40 int = (fi35 / (((fi35 - fi35)) | 1))
                _ = v40
            }
            _ = fi35
        }
    case "":
        var v41 uint = 912
        _ = v41
    }
    var v42 uint16 = 198
    if (!(false || true)) {
        for fi43 := 0; fi43 < 1; fi43++ {
            if (false || (3.14159 != (-2.25 + ((1.5 - 0.0) + (1e10 - ((1e10 - (1e10 - 0.0)) / 0.1)))))) {
                if len(v12) > 0 {
                    v12[0] = fi43
                }
                v3 -= (v3 - v3)
                var v44 int = 32768
                _ = v44
            }
            v9 = v6
            for fi45 := 0; fi45 < 18; fi45++ {
                switch (^(fi43)) {
                case 32767:
                    switch "a" {
                    case "\n\t":
                        v46 := v5[v6]
                        _ = v46
                    case "":
                        delete(v4, (fi45 / (((fi45 / (((^(fi45))) | 1))) | 1)))
                    default:
                        if false {
                            var v47 int64 = (v1 >> ((65535) & 63))
                            fmt.Println(fi43, v2, v6)
                            switch v6 {
                            case "hello":
                                var v48 int8 = -48
                                _ = v48
                            }
                            var v49 int64 = (v47 & (v47 ^ v1))
                            if len(v12) > 0 {
                                v12[0] = 42
                            }
                            _ = v47
                            _ = v49
                        } else {
                            v3 = 2
                            var v50 string = v6
                            var v51 bool = true
                            _ = v50
                            _ = v51
                        }
                    }
                default:
                    var v52 float32 = float32((1e10 * -2.25))
                    _ = v52
                }
                var v53 uint8 = 44
                var v54 bool = (!true)
                _ = fi45
                _ = v53
                _ = v54
            }
            _ = fi43
        }
    } else {
        v6 = ""
        var v55 int32 = v2
        _ = v55
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
