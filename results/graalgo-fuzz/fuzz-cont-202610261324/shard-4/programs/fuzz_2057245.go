package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 873
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    v5 := []int{9223372036854775807}
    fmt.Println(v2, v2)
    v1 += v1
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 bool = (777 >= 9223372036854775807)
        _ = fi6
        _ = v7
    }
    delete(v4, 117)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int16 = int16(v2)
        for fi10 := 0; fi10 < 13; fi10++ {
            v2 += v2
            var v11 int16 = v9
            switch "foo bar" {
            case "hello":
                _ = len(v3)
            case "foo bar":
                var v12 uint64 = 574
                _ = v12
            case "a":
                if len(v5) > 0 {
                    v5[0] = fi10
                }
            }
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v13 int32 = -362
    switch ("a" + ("hello" + "foo bar")) {
    case "a":
        var v14 bool = (639 == 289)
        _ = v14
    case "\n\t":
        v13 = v13
    case "hello":
        var v15 int16 = 101
        _ = v15
    }
    switch "" {
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "":
        var v16 float32 = (1e10 - (0.0 * 1e10))
        _ = v16
    default:
        var v17 int8 = 0
        _ = v17
    }
    if false {
        v3[739] = 310
        v5 = append(v5, 128)
    }
    v5 = append(v5, 0)
    fmt.Println(v13, v1)
    for fi18 := 0; fi18 < 13; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 float32 = 1.5
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            v2 = (v2 - (v2 * (v2 % (((v2 + (v2 % ((696) | 1)))) | 1))))
            _ = fi21
        }
        var v22 int64 = 0
        _ = fi18
        _ = v22
    }
    switch "\n\t" {
    case "foo bar":
        for fi23 := 0; fi23 < 1; fi23++ {
            for fi24 := 0; fi24 < 7; fi24++ {
                if (int32(v1) > (v13 << ((1) & 63))) {
                    var v25 bool = true
                    v4[2147483648] = 1
                    _ = v25
                }
                _ = fi24
            }
            if (v13 <= 52) {
                v5 = append(v5, fi23)
                fmt.Println(v2)
                for fi26 := 0; fi26 < 19; fi26++ {
                    fmt.Println(fi23, v13, fi23)
                    if (-32768 >= -1) {
                        var v27 uint64 = 42
                        switch "hello" {
                        case "hello":
                            var v28 int = 127
                            _ = v28
                        default:
                            switch (fi26 % ((fi23) | 1)) {
                            case 127:
                                v1 += 127
                                var v29 float32 = 3.14159
                                _ = v29
                            case -349:
                                var v30 float64 = 1e10
                                _ = v30
                            case 352:
                                var v31 uint64 = 243
                                switch (fi26 & (fi26 * (fi26 >> ((255) & 63)))) {
                                case 0:
                                    fmt.Println(v27, v13)
                                default:
                                    var v32 uint16 = v2
                                    _ = v32
                                }
                                if (v31 >= 292) {
                                    v27 += 42
                                    v13 -= (v13 - (v13 << ((647) & 63)))
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        if false {
                                            var v33 string = ""
                                            v27 += 0
                                            v34 := GenS1{F0: 42}
                                            _ = v33
                                            _ = v34
                                        }
                                    }
                                    var v35 uint32 = 2
                                    _ = v35
                                }
                                _ = v31
                            default:
                                var v36 int8 = 0
                                _ = v36
                            }
                        }
                        var v37 bool = true
                        _ = v27
                        _ = v37
                    }
                    fmt.Println(fi26, v2)
                    _ = fi26
                }
            } else {
                var v38 uint16 = (^(v2))
                var v39 uint32 = 2
                _ = v38
                _ = v39
            }
            for fi40 := 0; fi40 < 2; fi40++ {
                switch "foo bar" {
                case "a":
                    v41 := GenS1{F0: 833}
                    _ = v41
                }
                _ = fi40
            }
            _ = fi23
        }
    case "hello":
        var v42 uint64 = 847
        _ = v42
    }
    if (v13 == v13) {
        delete(v3, 9223372036854775807)
        v2 += v2
    } else {
        v2 = 300
        var v43 int8 = 42
        var v44 uint = 30
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 17; fi45++ {
        var v46 uint32 = 2147483647
        var v47 float64 = (1e10 * 3.14159)
        _ = fi45
        _ = v46
        _ = v47
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
