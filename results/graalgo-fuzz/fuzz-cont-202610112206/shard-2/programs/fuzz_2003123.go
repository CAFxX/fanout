package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int32; F1 int; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int8; F2 int }
type GenS5 struct { F0 int64; F1 int64; F2 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    var v1 string = ("" + "")
    v2 := []string{v1, v1, v1}
    var v3 bool = true
    v4 := make(map[uint32]bool)
    var v5 int32 = 42
    if true {
        delete(v4, 4294967295)
        if (v3 && v3) {
            if len(v2) > 0 {
                v2[0] = (("foo bar" + (v1 + ("" + ((((((v1 + ("" + "a")) + (v1 + (("\n\t" + "foo bar") + "foo bar"))) + ("foo bar" + "hello")) + "\n\t") + (v1 + "hello")) + v1)))) + (v1 + "foo bar"))
            }
            for fi6 := 0; fi6 < 16; fi6++ {
                for fi7 := 0; fi7 < 16; fi7++ {
                    var v8 uint32 = 2
                    if v3 {
                        v5 = v5
                        var v9 float32 = -2.25
                        var v10 bool = false
                        if v3 {
                            if v10 {
                                if (float32(1.5) > v9) {
                                    fmt.Println(v1)
                                    v3 = (98 >= 2)
                                }
                            }
                        } else {
                            switch (fi6 - int(v5)) {
                            case 32768:
                                v11 := GenS1{F0: 0, F1: 121}
                                _ = v11
                            case 42:
                                v2 = append(v2, "foo bar")
                            default:
                                switch v1 {
                                case "foo bar":
                                    if (v3 && false) {
                                        fmt.Println(v1, v3)
                                        var v12 int64 = 0
                                        var v13 float64 = 1.5
                                        v10 = (1000 != 1)
                                        v14 := GenS2{F0: (v5 - 0), F1: fi7, F2: v12}
                                        _ = v12
                                        _ = v13
                                        _ = v14
                                    } else {
                                        v10 = (1 > 157)
                                        fmt.Println(v10, v8, v9)
                                    }
                                case "":
                                    v8 = v8
                                case "\n\t":
                                    switch 32767 {
                                    case -521:
                                        v15 := GenS3{F0: 575}
                                        _ = v15
                                    default:
                                        if false {
                                            v8 = v8
                                        }
                                    }
                                default:
                                    v1 = ("\n\t" + v1)
                                }
                            }
                            var v16 int64 = 9223372036854775807
                            _ = v16
                        }
                        _ = v10
                        _ = v9
                    }
                    _ = fi7
                    _ = v8
                }
                _ = fi6
            }
        }
        switch "hello" {
        case "foo bar":
            if true {
                var v17 uint64 = 4294967295
                _ = len(v2)
                delete(v4, 2147483648)
                _ = v17
            }
        case "\n\t":
            fmt.Println(v1, v5)
        case "a":
            var v18 uint8 = 128
            _ = v18
        }
    }
    v19 := GenS3{F0: 1}
    var v20 uint8 = 128
    if (!(100 < 100)) {
        var v21 bool = v3
        v5 += v5
        _ = v21
    }
    var v22 uint = 296
    var v23 int = 4294967295
    v24 := GenS4{F0: 42, F1: 1, F2: (v23 | v23)}
    if ((v23 >> ((4294967295) & 63)) < 9223372036854775807) {
        var v25 float32 = ((0.1 * 3.14159) - 3.14159)
        _ = v25
    } else {
        var v26 int64 = 1
        v1 = ("a" + "\n\t")
        _ = len(v4)
        _ = v26
    }
    if (476 <= -32768) {
        _ = len(v2)
        for fi27 := 0; fi27 < 20; fi27++ {
            for fi28 := 0; fi28 < 6; fi28++ {
                switch (v23 ^ 0) {
                case 128:
                    v29 := GenS3{F0: 2}
                    if v3 {
                        v1 = "a"
                        fmt.Println(fi27, v1)
                    } else {
                        var v30 int8 = -30
                        var v31 string = (v1 + (v1 + ""))
                        _ = v30
                        _ = v31
                    }
                    _ = v29
                case 32767:
                    v32 := GenS1{F0: 42, F1: 100}
                    _ = v32
                }
                v33 := GenS5{F0: -1, F1: -9223372036854775808, F2: v1}
                switch (fi28 << ((v22) & 63)) {
                case -9223372036854775808:
                    v34 := GenS6{F0: 2147483647}
                    _ = v34
                case 0:
                    v23 += v23
                case 197:
                    _ = len(v2)
                default:
                    switch (fi27 | (v23 * fi27)) {
                    case 128:
                        v2 = append(v2, "a")
                    default:
                        switch v1 {
                        case "\n\t":
                            var v35 int = 127
                            _ = v35
                        }
                    }
                }
                _ = fi28
                _ = v33
            }
            _ = len(v2)
            _ = fi27
        }
    }
    v20 = v20
    v36 := GenS6{F0: v5}
    var v37 uint16 = 32767
    var v38 uint = v22
    if (v20 <= v20) {
        delete(v4, 26)
        if (32767 == 526) {
            var v39 bool = false
            v4[65535] = (18446744073709551615 != 4294967295)
            _ = v39
        } else {
            var v40 uint = (v38 | v22)
            for fi41 := 0; fi41 < 17; fi41++ {
                var v42 string = v1
                v20 = (v20 << (((v40 << (((^(v38))) & 63))) & 63))
                _ = fi41
                _ = v42
            }
            var v43 uint32 = 42
            _ = v40
            _ = v43
        }
    } else {
        var v44 int8 = -85
        v2 = append(v2, "\n\t")
        switch int(v44) {
        case 127:
            switch v1 {
            case "hello":
                v4[0] = ((v44 | v44) == (^(v44)))
            case "":
                v45 := GenS7{F0: v44, F1: 0, F2: (!v3)}
                _ = v45
            case "\n\t":
                var v46 int64 = 9223372036854775807
                _ = v46
            }
        }
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
