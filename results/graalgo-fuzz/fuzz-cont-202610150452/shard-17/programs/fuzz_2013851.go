package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int32; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 42
    v3 := []uint32{2147483647, 2}
    v4 := []uint32{2147483647, 247, 42}
    var v5 uint64 = 494
    v6 := make(map[int]bool)
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v7 int8 = -1
    v7 += (v7 >> ((1) & 63))
    var v8 int = 0
    if true {
        v9 := GenS1{F0: 2147483647}
        switch -2147483648 {
        case 2147483647:
            v10 := GenS2{F0: 167}
            _ = v10
        }
        v5 = v5
        _ = v9
    } else {
        v11 := v6[(v8 ^ 0)]
        if (v11 && v11) {
            switch ("a" + "foo bar") {
            case "foo bar":
                var v12 int8 = (v7 % (((v7 & v7)) | 1))
                _ = v12
            case "":
                switch v8 {
                case 448:
                    if (!v11) {
                        if false {
                            if (v11 && ((v11 || ((v2 | v2) >= uint(v5))) || (int8(v5) <= (v7 / ((v7) | 1))))) {
                                var v13 float64 = 0.0
                                _ = v13
                            }
                            var v14 float64 = float64(3.14159)
                            for fi15 := 0; fi15 < 5; fi15++ {
                                switch (fi15 * 42) {
                                case 763:
                                    fmt.Println(v7, fi15)
                                case -9223372036854775808:
                                    for fi16 := 0; fi16 < 8; fi16++ {
                                        switch (("" + "\n\t") + "\n\t") {
                                        case "":
                                            fmt.Println(v7)
                                        case "a":
                                            v2 = v2
                                        case "hello":
                                            var v17 bool = (((((0.1 + v14) / -2.25) - v14) + (0.0 - v14)) == (v14 + v14))
                                            _ = v17
                                        }
                                        v8 -= fi16
                                        v18 := GenS1{F0: 65535}
                                        _ = fi16
                                        _ = v18
                                    }
                                case 4294967295:
                                    var v19 string = ""
                                    _ = v19
                                default:
                                    v5 -= (v5 * (v5 - 267))
                                }
                                var v20 uint32 = 4294967295
                                v1[""] = ("foo bar" + ("a" + (("foo bar" + ("\n\t" + "")) + "\n\t")))
                                _ = fi15
                                _ = v20
                            }
                            _ = v14
                        }
                    }
                }
            case "\n\t":
                var v21 int64 = 1
                _ = v21
            default:
                v22 := GenS1{F0: 2}
                _ = v22
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                fmt.Println(v2, v5, v2)
                switch (fi23 | (v8 / (((v8 & v8)) | 1))) {
                case -678:
                    var v24 int16 = -32768
                    _ = v24
                case 1:
                    fmt.Println(fi23, v7)
                }
                _ = fi23
            }
            v25 := GenS1{F0: 193}
            _ = v25
        }
        _ = v11
    }
    if false {
        if ((-(v8)) > v8) {
            v26 := GenS2{F0: -9223372036854775808}
            _ = v26
        }
        v7 += (^(v7))
        if (float64((-2.25 / -2.25)) >= (float64(0.0) - 1e10)) {
            switch v8 {
            case -1:
                v27 := GenS2{F0: 9223372036854775807}
                _ = v27
            case 32767:
                switch "hello" {
                case "\n\t":
                    v28 := GenS1{F0: 4294967295}
                    _ = v28
                case "":
                    var v29 uint64 = 0
                    _ = v29
                }
            case 1:
                var v30 uint32 = 65535
                _ = v30
            }
        }
    } else {
        fmt.Println(v8, v7, v8)
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        if (v5 != v5) {
            v32 := GenS1{F0: 1}
            switch fi31 {
            case 32768:
                if ((!true) && (true && false)) {
                    v5 = 395
                    fmt.Println(v2, v5)
                }
            case 4294967295:
                v8 = fi31
            case 32767:
                fmt.Println(fi31, v7)
            default:
                switch (("foo bar" + "a") + "a") {
                case "a":
                    for fi33 := 0; fi33 < 19; fi33++ {
                        v5 = 18446744073709551615
                        _ = fi33
                    }
                case "hello":
                    var v34 int16 = -32768
                    _ = v34
                case "":
                    var v35 uint32 = 253
                    _ = v35
                default:
                    if (false || false) {
                        v2 -= v2
                        v8 -= (v8 - v8)
                        v36 := GenS2{F0: 0}
                        _ = v36
                    }
                }
            }
            var v37 uint16 = 255
            _ = v32
            _ = v37
        } else {
            v2 = (v2 & 2)
        }
        _ = fi31
    }
    switch (v8 & v8) {
    case 224:
        v38 := GenS2{F0: 1}
        _ = v38
    case 32767:
        v39 := GenS3{F0: ((v8 >> ((v2) & 63)) > v8), F1: 42, F2: 196}
        _ = v39
    default:
        for fi40 := 0; fi40 < 9; fi40++ {
            v5 = v5
            fmt.Println(fi40)
            _ = fi40
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
