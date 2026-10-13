package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 int16 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint }

func main() {
    var v1 int64 = -1
    v2 := []uint32{1, 0, 1, 1}
    var v3 int16 = 42
    for fi4 := 0; fi4 < 16; fi4++ {
        if true {
            v5 := GenS1{F0: 323}
            _ = v5
        }
        if true {
            switch "a" {
            case "a":
                v3 += int16(v1)
            }
            if true {
                if (true && false) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v3 -= -546
                    var v6 uint16 = 255
                    _ = v6
                } else {
                    v3 += 745
                    var v7 uint8 = 176
                    _ = v7
                }
                var v8 uint16 = 1
                for fi9 := 0; fi9 < 18; fi9++ {
                    if true {
                        v3 -= v3
                        if true {
                            var v10 uint8 = 130
                            v11 := GenS1{F0: 0}
                            var v12 int64 = v1
                            _ = v10
                            _ = v11
                            _ = v12
                        } else {
                            v3 -= v3
                            var v13 int32 = 2147483647
                            switch "foo bar" {
                            case "":
                                var v14 string = "foo bar"
                                _ = v14
                            default:
                                if (("\n\t" + "") >= "\n\t") {
                                    var v15 int16 = v3
                                    var v16 uint16 = v8
                                    v17 := GenS1{F0: v13}
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                            }
                            _ = v13
                        }
                        v18 := GenS1{F0: 1}
                        _ = v18
                    }
                    if false {
                        v2 = append(v2, 2)
                        v1 = int64(v3)
                    }
                    var v19 int64 = -564
                    _ = fi9
                    _ = v19
                }
                _ = v8
            } else {
                if false {
                    fmt.Println(fi4)
                    v3 += (v3 << ((175) & 63))
                    var v20 int32 = -2147483648
                    _ = v20
                }
                switch (("hello" + "foo bar") + ("\n\t" + "\n\t")) {
                case "\n\t":
                    fmt.Println(fi4)
                default:
                    var v21 uint64 = 42
                    _ = v21
                }
                v22 := GenS2{F0: 1}
                _ = v22
            }
        }
        v3 = 1
        _ = fi4
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        switch ("a" + (("hello" + ((("foo bar" + "") + (((("a" + "a") + "a") + "hello") + (("a" + (("a" + "hello") + "foo bar")) + ("" + "foo bar")))) + "a")) + ("a" + "hello"))) {
        case "":
            v1 = (v1 ^ v1)
        case "\n\t":
            fmt.Println(v3, fi23)
        case "hello":
            var v24 uint8 = 128
            _ = v24
        default:
            if true {
                var v25 int8 = 0
                _ = v25
            } else {
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 uint8 = 17
                    v3 = (-(v3))
                    _ = len(v2)
                    _ = fi26
                    _ = v27
                }
            }
        }
        switch (fi23 % (((fi23 - fi23)) | 1)) {
        case 32767:
            var v28 uint = 295
            _ = v28
        }
        if ((0.1 / 3.14159) < 3.14159) {
            var v29 int = 9223372036854775807
            if len(v2) > 0 {
                v2[0] = 8
            }
            _ = v29
        }
        _ = fi23
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = v1
    switch -512 {
    case 1:
        if false {
            fmt.Println(v1, v3)
            v1 -= v1
        }
    default:
        var v30 uint8 = 127
        _ = v30
    }
    var v31 string = "foo bar"
    v32 := GenS3{F0: 1, F1: -128, F2: v3}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = (v1 - v1)
    var v33 int32 = -84
    if (127 <= 331) {
        var v34 int32 = -2147483648
        _ = v34
    } else {
        v2 = append(v2, 553)
    }
    v35 := GenS4{F0: -128, F1: 4294967295, F2: 928}
    var v36 uint16 = 65535
    v36 += 2
    if (((true || (4294967295 > 0)) || true) && true) {
        if len(v2) > 0 {
            v2[0] = 2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v36
}
