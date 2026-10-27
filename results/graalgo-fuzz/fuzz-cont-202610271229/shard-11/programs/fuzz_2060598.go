package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS4 struct { F0 int; F1 int16 }

func main() {
    var v1 uint = 1
    v2 := make(map[string]string)
    var v3 int8 = -1
    var v4 int8 = (v3 * v3)
    var v5 int8 = (v3 % ((v4) | 1))
    v6 := GenS1{F0: -101}
    fmt.Println(v1, v4, v4)
    var v7 int64 = 0
    v7 -= 694
    v4 += (v4 - (-(v5)))
    for fi8 := 0; fi8 < 17; fi8++ {
        delete(v2, "a")
        var v9 float32 = 1.5
        if false {
            v10 := GenS2{F0: v7, F1: "a"}
            _ = v10
        } else {
            if (("" + "\n\t") >= "hello") {
                switch ("foo bar" + ("\n\t" + ("hello" + "\n\t"))) {
                case "\n\t":
                    v5 += (^(v4))
                default:
                    var v11 uint32 = 0
                    _ = v11
                }
                if true {
                    v12 := GenS2{F0: 9223372036854775807, F1: (("a" + "hello") + "")}
                    switch ("\n\t" + "") {
                    case "foo bar":
                        if false {
                            v4 += 44
                            for fi13 := 0; fi13 < 14; fi13++ {
                                switch (fi8 - 2147483647) {
                                case -881:
                                    v9 += v9
                                case 128:
                                    v14 := GenS1{F0: v3}
                                    _ = v14
                                case 930:
                                    var v15 uint = 4294967295
                                    _ = v15
                                default:
                                    v5 = (v4 + (v5 & -64))
                                }
                                if (((42 <= -1) && (1e10 == (v9 * v9))) || (!false)) {
                                    if ((1.5 - (1.5 - ((-2.25 - (3.14159 * 0.1)) * ((0.0 - float64((((((v9 - float32(0.1)) - 0.0) + v9) - v9) - ((v9 / 1.0) + v9)))) * float64((1e10 * ((1.5 - float32(3.14159)) + ((0.1 * 0.0) * 1.5)))))))) != 0.1) {
                                        v16 := []int{(^(fi8)), fi8, -2147483648}
                                        var v17 int64 = int64(v3)
                                        _ = v16
                                        _ = v17
                                    } else {
                                        v3 += (v5 + (v5 / ((v5) | 1)))
                                        switch "" {
                                        case "":
                                            switch (fi13 & fi13) {
                                            case 42:
                                                var v18 int32 = -1
                                                _ = v18
                                            }
                                        default:
                                            switch (-(fi13)) {
                                            case -1:
                                                v4 = v4
                                            }
                                        }
                                        var v19 int64 = 1
                                        v19 += 1
                                        fmt.Println(v9, v4, v1)
                                        v4 -= (v5 | 42)
                                        if false {
                                            if false {
                                                v20 := GenS3{F0: 127, F1: 2, F2: -447}
                                                _ = v20
                                            }
                                            v5 += (-(v5))
                                            var v21 int = fi13
                                            _ = v21
                                        }
                                        _ = v19
                                    }
                                    v9 -= (v9 - ((v9 * v9) - (1.5 - (0.0 - (float32(-2.25) - v9)))))
                                }
                                v1 += 18446744073709551615
                                _ = fi13
                            }
                            fmt.Println(v7, v3)
                        }
                    case "\n\t":
                        var v22 int = (fi8 ^ (fi8 << ((v1) & 63)))
                        _ = v22
                    case "":
                        switch fi8 {
                        case -400:
                            var v23 int64 = (v7 / ((int64(v4)) | 1))
                            _ = v23
                        case 736:
                            for fi24 := 0; fi24 < 2; fi24++ {
                                var v25 uint32 = 42
                                if false {
                                    v26 := []int{(fi8 * 682)}
                                    delete(v2, "")
                                    v27 := GenS3{F0: 128, F1: v25, F2: -1}
                                    _ = v26
                                    _ = v27
                                }
                                _ = fi24
                                _ = v25
                            }
                        }
                    }
                    if false {
                        switch "\n\t" {
                        case "a":
                            if false {
                                v3 -= v5
                            }
                        }
                    } else {
                        var v28 int8 = v4
                        _ = v28
                    }
                    _ = v12
                }
                if true {
                    fmt.Println(fi8, v7)
                    var v29 string = "foo bar"
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        var v31 int64 = 756
                        _ = len(v2)
                        switch 128 {
                        case -545:
                            v32 := GenS4{F0: fi8, F1: -1}
                            _ = v32
                        case 242:
                            var v33 float64 = -2.25
                            _ = v33
                        default:
                            fmt.Println(v7, v7, v31)
                        }
                        v31 = 42
                        _ = fi30
                        _ = v31
                    }
                    v5 = 100
                }
            } else {
                var v34 int = fi8
                _ = v34
            }
            v7 = (^(v7))
        }
        _ = fi8
        _ = v9
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        v2["hello"] = (("foo bar" + "\n\t") + "\n\t")
        v36 := GenS3{F0: 1, F1: 65535, F2: -1}
        var v37 int = 514
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 int32 = 330
    for fi39 := 0; fi39 < 20; fi39++ {
        for fi40 := 0; fi40 < 14; fi40++ {
            switch fi40 {
            case 32768:
                var v41 float64 = 3.14159
                _ = v41
            case 32767:
                v42 := []int{(fi40 * -2147483648), 42, 127, (fi39 << (((v1 ^ (v1 ^ (v1 + v1)))) & 63))}
                _ = v42
            case 969:
                if (42 < 4294967295) {
                    var v43 int16 = -32768
                    _ = v43
                }
            }
            _ = fi40
        }
        switch fi39 {
        case -2147483648:
            v44 := GenS1{F0: (v4 * (v5 % ((v5) | 1)))}
            _ = v44
        default:
            for fi45 := 0; fi45 < 7; fi45++ {
                var v46 bool = (!true)
                _ = fi45
                _ = v46
            }
        }
        fmt.Println(fi39, v4, v3)
        _ = fi39
    }
    var v47 uint = 255
    v48 := []int{487, 1, -775, 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
