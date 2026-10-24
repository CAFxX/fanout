package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 bool; F2 bool }
type GenS3 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := []int{4294967295, -9223372036854775808, 555}
    v3 := make(map[int]int)
    var v4 int = -948
    v5 := v3[v4]
    v4 += (v4 >> ((2) & 63))
    switch v5 {
    case -2147483648:
        fmt.Println(v5)
    case 939:
        v4 = v5
    }
    var v6 int = v5
    var v7 uint = 1
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint64 = 926
        v10 := GenS1{F0: 1}
        fmt.Println(fi8, v6)
        _ = fi8
        _ = v10
        _ = v9
    }
    switch "hello" {
    case "a":
        if (88 == 100) {
            v5 += v4
            var v11 uint16 = 1
            _ = len(v1)
            _ = v11
        } else {
            var v12 uint8 = 1
            var v13 uint = (v7 * v7)
            var v14 int16 = 680
            _ = v12
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v4 += v6
        var v16 float32 = -2.25
        _ = fi15
        _ = v16
    }
    if false {
        _ = len(v2)
        _ = len(v1)
    } else {
        if false {
            switch "foo bar" {
            case "hello":
                if (2 > 2) {
                    var v17 int32 = 1
                    _ = v17
                } else {
                    v7 -= v7
                    if true {
                        var v18 int64 = -456
                        v19 := GenS1{F0: -1}
                        _ = v18
                        _ = v19
                    } else {
                        v7 = 2
                    }
                    var v20 int32 = 42
                    _ = v20
                }
            case "a":
                v21 := GenS1{F0: 1}
                _ = v21
            case "\n\t":
                switch "hello" {
                case "\n\t":
                    var v22 int = 1
                    _ = v22
                case "a":
                    switch "foo bar" {
                    case "\n\t":
                        for fi23 := 0; fi23 < 7; fi23++ {
                            var v24 uint16 = 32767
                            for fi25 := 0; fi25 < 20; fi25++ {
                                fmt.Println(v5, v6, fi23)
                                v2 = append(v2, -2147483648)
                                _ = fi25
                            }
                            _ = fi23
                            _ = v24
                        }
                    case "":
                        switch ("\n\t" + ("\n\t" + ((("\n\t" + "") + ("a" + ("foo bar" + "a"))) + (("\n\t" + "hello") + "hello")))) {
                        case "\n\t":
                            for fi26 := 0; fi26 < 14; fi26++ {
                                for fi27 := 0; fi27 < 13; fi27++ {
                                    v28 := v3[(fi27 * (^(v5)))]
                                    _ = fi27
                                    _ = v28
                                }
                                _ = fi26
                            }
                        default:
                            if true {
                                var v29 uint8 = 255
                                fmt.Println(v7, v6, v7)
                                v30 := GenS1{F0: 10}
                                _ = v29
                                _ = v30
                            } else {
                                var v31 int16 = -623
                                v4 += (v5 & v4)
                                var v32 uint16 = 662
                                _ = v31
                                _ = v32
                            }
                        }
                    case "a":
                        v2 = append(v2, v4)
                    }
                default:
                    fmt.Println(v4, v7)
                }
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                var v34 uint16 = 255
                var v35 uint8 = 91
                _ = fi33
                _ = v34
                _ = v35
            }
            v4 = (^(v5))
        } else {
            if (1.5 <= 0.1) {
                for fi36 := 0; fi36 < 12; fi36++ {
                    var v37 uint32 = 42
                    var v38 uint64 = 662
                    v39 := v3[v5]
                    _ = fi36
                    _ = v37
                    _ = v38
                    _ = v39
                }
                if true {
                    var v40 uint64 = 18446744073709551615
                    for fi41 := 0; fi41 < 16; fi41++ {
                        switch int(v7) {
                        case 925:
                            fmt.Println(v5, v7)
                        case 42:
                            fmt.Println(v5, v40, v6)
                        case 2147483647:
                            if (((0.0 + 0.0) / -2.25) == 0.1) {
                                delete(v1, (("hello" + "a") + "foo bar"))
                            } else {
                                switch (-(fi41)) {
                                case 9223372036854775807:
                                    var v42 int64 = -94
                                    _ = v42
                                case 111:
                                    v43 := GenS1{F0: -1}
                                    _ = v43
                                case -885:
                                    for fi44 := 0; fi44 < 7; fi44++ {
                                        v2 = append(v2, (fi44 << ((0) & 63)))
                                        _ = fi44
                                    }
                                }
                                v5 = v6
                                fmt.Println(v5, v40)
                            }
                        }
                        v45 := GenS1{F0: 30}
                        _ = fi41
                        _ = v45
                    }
                    v6 = (v4 & (v5 | v5))
                    _ = v40
                }
            } else {
                var v46 uint8 = 255
                _ = v46
            }
            fmt.Println(v5, v4)
        }
        var v47 float32 = 1e10
        for fi48 := 0; fi48 < 4; fi48++ {
            var v49 int64 = -665
            v50 := GenS1{F0: 1}
            switch v5 {
            case 1:
                if false {
                    v51 := GenS1{F0: 118}
                    var v52 int16 = -429
                    var v53 int32 = -763
                    _ = v51
                    _ = v52
                    _ = v53
                }
            case 693:
                if true {
                    v54 := GenS2{F0: (v7 << (((v7 / (((v7 + 1)) | 1))) & 63)), F1: true, F2: ((v49 & 9223372036854775807) >= (v49 % ((v49) | 1)))}
                    _ = v54
                } else {
                    var v55 int32 = -203
                    v6 -= (v5 ^ v4)
                    v1["a"] = (v4 - 127)
                    _ = v55
                }
            case 431:
                var v56 uint16 = 1
                _ = v56
            }
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = v47
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi57 := 0; fi57 < 14; fi57++ {
        switch fi57 {
        case -1:
            for fi58 := 0; fi58 < 3; fi58++ {
                fmt.Println(v6, fi57)
                if (!(((1 >= 0) && false) || true)) {
                    v59 := v3[v4]
                    fmt.Println(fi57, v5, v4)
                    _ = v59
                }
                var v60 int8 = 0
                _ = fi58
                _ = v60
            }
        case 185:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v61 := GenS3{F0: 42, F1: -1}
        v62 := GenS1{F0: -1}
        _ = fi57
        _ = v61
        _ = v62
    }
    switch (^(v4)) {
    case 9223372036854775807:
        var v63 int32 = 436
        _ = v63
    default:
        fmt.Println(v4, v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
