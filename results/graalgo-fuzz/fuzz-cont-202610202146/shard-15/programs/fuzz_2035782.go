package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint16; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int8 = -1
    var v2 string = ((("" + (("\n\t" + "\n\t") + "hello")) + "a") + "hello")
    v3 := make(map[int]string)
    v4 := []int{}
    v5 := GenS1{F0: 882}
    switch -260 {
    case -9223372036854775808:
        v6 := v3[-1]
        _ = v6
    case -116:
        var v7 string = (v2 + ("a" + (v2 + ("\n\t" + (v2 + v2)))))
        _ = v7
    case 42:
        for fi8 := 0; fi8 < 8; fi8++ {
            if true {
                fmt.Println(fi8)
                v9 := GenS2{F0: 728}
                _ = v9
            }
            v10 := GenS3{F0: fi8, F1: 32767, F2: 0}
            var v11 int32 = -1
            _ = fi8
            _ = v10
            _ = v11
        }
    }
    if true {
        switch ("hello" + "") {
        case "foo bar":
            _ = len(v4)
        default:
            var v12 int32 = -268
            _ = v12
        }
        v2 = "foo bar"
        v13 := GenS4{F0: 1}
        _ = v13
    }
    v1 = -128
    if true {
        var v14 uint16 = 255
        _ = v14
    }
    v4 = append(v4, 128)
    v15 := GenS5{F0: -248}
    var v16 uint = 1
    if false {
        v17 := v3[42]
        for fi18 := 0; fi18 < 3; fi18++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v19 uint32 = 2147483647
            fmt.Println(fi18, fi18, v2)
            _ = fi18
            _ = v19
        }
        if len(v4) > 0 {
            v4[0] = -873
        }
        _ = v17
    } else {
        if (!false) {
            var v20 uint = v16
            for fi21 := 0; fi21 < 10; fi21++ {
                if (-267 == -1) {
                    v22 := GenS1{F0: 939}
                    _ = v22
                }
                if false {
                    v20 -= v20
                }
                fmt.Println(v16)
                _ = fi21
            }
            for fi23 := 0; fi23 < 8; fi23++ {
                if false {
                    var v24 int8 = v1
                    _ = v24
                } else {
                    _ = len(v4)
                    if true {
                        fmt.Println(fi23)
                        v4 = append(v4, fi23)
                        for fi25 := 0; fi25 < 19; fi25++ {
                            delete(v3, -231)
                            v20 = (v16 + v20)
                            _ = fi25
                        }
                    } else {
                        if false {
                            var v26 uint32 = 0
                            if true {
                                var v27 float64 = 0.1
                                var v28 uint64 = 18446744073709551615
                                _ = v27
                                _ = v28
                            } else {
                                v16 -= (v16 + (v16 * 11))
                            }
                            if (true && ((v2 + "foo bar") <= v2)) {
                                v20 += v20
                                for fi29 := 0; fi29 < 17; fi29++ {
                                    _ = len(v4)
                                    if true {
                                        fmt.Println(v20)
                                    } else {
                                        v26 -= (v26 % ((v26) | 1))
                                        if (v2 != v2) {
                                            var v30 uint64 = 1
                                            switch (fi23 << ((v16) & 63)) {
                                            case 9223372036854775807:
                                                v16 = v20
                                            default:
                                                v16 += (v16 * 833)
                                            }
                                            if false {
                                                v26 += v26
                                                v20 += v16
                                            }
                                            _ = v30
                                        }
                                    }
                                    _ = fi29
                                }
                            }
                            _ = v26
                        }
                        switch (fi23 >> (((v20 >> (((v20 * 0)) & 63))) & 63)) {
                        case -2147483648:
                            var v31 uint64 = 2
                            _ = v31
                        default:
                            fmt.Println(v2)
                        }
                    }
                    for fi32 := 0; fi32 < 12; fi32++ {
                        v16 += 42
                        _ = fi32
                    }
                }
                switch (v2 + ("foo bar" + v2)) {
                case "\n\t":
                    fmt.Println(fi23, v2, v1)
                case "hello":
                    if len(v4) > 0 {
                        v4[0] = fi23
                    }
                }
                if true {
                    fmt.Println(v16)
                    if true {
                        var v33 float64 = 3.14159
                        v33 = (v33 / v33)
                        _ = v33
                    }
                    v16 = 511
                }
                _ = fi23
            }
            _ = v20
        }
        switch -319 {
        case 2147483647:
            if (!(1 >= 926)) {
                v34 := GenS5{F0: int16(v1)}
                v35 := GenS5{F0: -1}
                var v36 string = "foo bar"
                _ = v34
                _ = v35
                _ = v36
            } else {
                var v37 int = 32767
                _ = v37
            }
        case -824:
            v4 = append(v4, 32767)
        case 32767:
            var v38 float64 = 0.1
            _ = v38
        default:
            if len(v4) > 0 {
                v4[0] = 62
            }
        }
    }
    for fi39 := 0; fi39 < 10; fi39++ {
        v1 = (v1 | 100)
        for fi40 := 0; fi40 < 9; fi40++ {
            var v41 int64 = -1
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    var v42 int = 2147483647
    var v43 uint8 = 42
    var v44 int = (v42 / ((2147483647) | 1))
    var v45 uint32 = 2147483647
    v2 = v2
    var v46 uint64 = uint64(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
