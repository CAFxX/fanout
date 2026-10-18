package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 int16; F1 uint8; F2 uint }
type GenS3 struct { F0 uint8; F1 string; F2 uint8 }

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{487, 456}
    v3 := make(map[string]string)
    var v4 int8 = 127
    switch "foo bar" {
    case "hello":
        var v5 int = -89
        _ = v5
    }
    if (36 > 42) {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 float32 = 1e10
            for fi8 := 0; fi8 < 15; fi8++ {
                v9 := GenS1{F0: (v4 >> ((45) & 63)), F1: -1, F2: (v4 == int8(fi8))}
                _ = fi8
                _ = v9
            }
            fmt.Println(v4, v1, v7)
            _ = fi6
            _ = v7
        }
        v1 = v1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        switch "foo bar" {
        case "hello":
            for fi11 := 0; fi11 < 4; fi11++ {
                if true {
                    v1 = v1
                } else {
                    var v12 uint16 = v1
                    v13 := GenS1{F0: v4, F1: 42, F2: (-1 <= -9223372036854775808)}
                    v3["hello"] = "a"
                    switch "" {
                    case "a":
                        if true {
                            delete(v3, "foo bar")
                            if false {
                                var v14 float32 = float32(3.14159)
                                v3[("" + ("foo bar" + ("a" + "\n\t")))] = "a"
                                _ = v14
                            } else {
                                v4 = 0
                                v15 := GenS2{F0: 1000, F1: 163, F2: 242}
                                _ = v15
                            }
                        }
                    case "hello":
                        switch (("foo bar" + (("a" + "") + "foo bar")) + "") {
                        case "":
                            fmt.Println(v4, fi10, fi11)
                        case "a":
                            var v16 uint64 = 1
                            _ = v16
                        case "\n\t":
                            v17 := v3["\n\t"]
                            _ = v17
                        default:
                            _ = len(v3)
                        }
                    case "":
                        v1 -= (v1 * v12)
                    default:
                        var v18 uint64 = 0
                        fmt.Println(v12, v4, v18)
                        _ = v18
                    }
                    delete(v3, "a")
                    switch (("foo bar" + "\n\t") + "\n\t") {
                    case "a":
                        v19 := GenS1{F0: 127, F1: -1, F2: false}
                        _ = v19
                    default:
                        v20 := GenS1{F0: v4, F1: 42, F2: true}
                        _ = v20
                    }
                    delete(v3, "\n\t")
                    if false {
                        v21 := GenS1{F0: v4, F1: 1, F2: (!(true && false))}
                        fmt.Println(fi11, fi11, fi10)
                        switch "hello" {
                        case "a":
                            v22 := GenS1{F0: 0, F1: 1, F2: false}
                            _ = v22
                        case "\n\t":
                            switch "foo bar" {
                            case "hello":
                                v4 = 100
                            case "a":
                                v12 = 338
                            }
                            if ("hello" != ((("" + "hello") + "") + "\n\t")) {
                                var v23 float64 = 1e10
                                _ = v23
                            }
                            var v24 bool = ((705 >= 42) || true)
                            _ = v24
                        default:
                            v12 -= (v12 | 443)
                        }
                        _ = v21
                    }
                    _ = v12
                    _ = v13
                }
                _ = fi11
            }
        case "":
            for fi25 := 0; fi25 < 2; fi25++ {
                switch (("\n\t" + "hello") + "\n\t") {
                case "a":
                    var v26 uint64 = 0
                    if (87 > 127) {
                        v1 = (^(v1))
                    }
                    _ = len(v2)
                    _ = v26
                case "":
                    var v27 uint16 = 255
                    _ = v27
                case "foo bar":
                    var v28 uint64 = 4294967295
                    _ = v28
                }
                _ = fi25
            }
        default:
            var v29 float32 = (1e10 * 0.0)
            _ = v29
        }
        switch (^(fi10)) {
        case 32767:
            switch (-(fi10)) {
            case 32768:
                for fi30 := 0; fi30 < 11; fi30++ {
                    if (!(fi30 >= -2147483648)) {
                        var v31 uint8 = 177
                        var v32 int = fi30
                        var v33 int32 = -2147483648
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    var v34 uint32 = 65535
                    if (true && true) {
                        v35 := GenS3{F0: 201, F1: ("foo bar" + "a"), F2: 0}
                        _ = v35
                    }
                    _ = fi30
                    _ = v34
                }
            case 909:
                for fi36 := 0; fi36 < 19; fi36++ {
                    v1 = (v1 >> ((4294967295) & 63))
                    v4 += (v4 % (((v4 * (v4 + v4))) | 1))
                    fmt.Println(fi36, v4)
                    _ = fi36
                }
            }
        case -1:
            v4 += v4
        case 833:
            v37 := v3["\n\t"]
            _ = v37
        }
        for fi38 := 0; fi38 < 14; fi38++ {
            var v39 bool = ((v4 ^ int8(v1)) > v4)
            _ = fi38
            _ = v39
        }
        _ = fi10
    }
    var v40 uint16 = v1
    var v41 float32 = 1.5
    var v42 uint64 = 42
    var v43 uint64 = 638
    for fi44 := 0; fi44 < 1; fi44++ {
        switch "" {
        case "foo bar":
            v3[("a" + "\n\t")] = ((("hello" + ("\n\t" + "foo bar")) + ("hello" + "a")) + "hello")
        }
        _ = fi44
    }
    var v45 float32 = v41
    delete(v3, "hello")
    v46 := GenS1{F0: v4, F1: 205, F2: true}
    if true {
        var v47 uint8 = 128
        v2 = append(v2, 0)
        _ = v47
    }
    if (float64(v41) == 1.5) {
        v48 := GenS2{F0: -32768, F1: 42, F2: 4294967295}
        v49 := GenS1{F0: v4, F1: 9223372036854775807, F2: (false && (false && true))}
        v43 = v42
        _ = v48
        _ = v49
    } else {
        var v50 uint32 = 65535
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v45
    _ = v46
}
