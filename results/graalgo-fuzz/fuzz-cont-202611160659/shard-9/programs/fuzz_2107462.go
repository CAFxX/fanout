package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint; F1 int16; F2 string }
type GenS4 struct { F0 int; F1 bool; F2 string }

func main() {
    var v1 uint32 = 20
    v2 := make(map[string]bool)
    var v3 uint64 = 504
    var v4 uint8 = uint8(v3)
    v5 := []int{-372, -1}
    v4 = v4
    fmt.Println(v3, v1)
    v4 -= v4
    var v6 uint16 = 4
    v1 += (v1 / (((v1 | 1)) | 1))
    fmt.Println(v1)
    var v7 uint = 65535
    var v8 float64 = 1e10
    if ((v6 << ((987) & 63)) >= v6) {
        switch ("foo bar" + "foo bar") {
        case "hello":
            fmt.Println(v6)
        case "a":
            for fi9 := 0; fi9 < 4; fi9++ {
                switch ("" + "a") {
                case "hello":
                    var v10 int32 = -1
                    _ = v10
                case "":
                    if ((v3 & 18446744073709551615) > 2) {
                        fmt.Println(v6, v1, fi9)
                    } else {
                        for fi11 := 0; fi11 < 3; fi11++ {
                            var v12 uint16 = 2
                            fmt.Println(v4, fi9)
                            var v13 float64 = 0.1
                            switch fi11 {
                            case 32767:
                                v13 += v13
                            case -9223372036854775808:
                                v5 = append(v5, (fi9 * int(v7)))
                            case -268:
                                switch "" {
                                case "a":
                                    v14 := GenS1{F0: (^(v1)), F1: (^(fi11))}
                                    _ = v14
                                }
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    switch (fi11 / ((fi11) | 1)) {
                                    case 127:
                                        v15 := v2[""]
                                        _ = v15
                                    case 32767:
                                        if (!(!true)) {
                                            if true {
                                                v6 = v6
                                                if ((^(v4)) == v4) {
                                                    switch "hello" {
                                                    case "a":
                                                        var v16 int16 = 289
                                                        _ = v16
                                                    }
                                                    var v17 int = fi9
                                                    if (true && ((!false) || true)) {
                                                        switch fi9 {
                                                        case 32767:
                                                            v18 := GenS2{F0: (^(fi11)), F1: v7}
                                                            _ = v18
                                                        default:
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        }
                                                    } else {
                                                        var v19 uint8 = (v4 + (^(v4)))
                                                        _ = len(v2)
                                                        v13 += 0.1
                                                        _ = v19
                                                    }
                                                    _ = v17
                                                }
                                                var v20 float32 = (1e10 + 1e10)
                                                v21 := GenS2{F0: fi9, F1: v7}
                                                _ = v20
                                                _ = v21
                                            } else {
                                                if true {
                                                    var v22 uint = (v7 & v7)
                                                    _ = v22
                                                } else {
                                                    fmt.Println(v4)
                                                    var v23 uint = uint(v1)
                                                    _ = v23
                                                }
                                                v4 = v4
                                            }
                                            if (v6 == (^(v12))) {
                                                v13 -= 0.1
                                            }
                                        }
                                    case -255:
                                        var v24 float32 = -2.25
                                        _ = v24
                                    default:
                                        var v25 uint64 = uint64(v4)
                                        _ = v25
                                    }
                                case "":
                                    var v26 int = fi11
                                    _ = v26
                                case "\n\t":
                                    v27 := v2["a"]
                                    _ = v27
                                default:
                                    v28 := GenS3{F0: 443, F1: 793, F2: (("foo bar" + "\n\t") + "a")}
                                    _ = v28
                                }
                            }
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        _ = len(v2)
                        var v29 bool = false
                        _ = v29
                    }
                default:
                    v30 := GenS4{F0: (fi9 >> ((v7) & 63)), F1: false, F2: "\n\t"}
                    _ = v30
                }
                var v31 uint = v7
                _ = fi9
                _ = v31
            }
        case "foo bar":
            switch 210 {
            case 732:
                var v32 int = 9223372036854775807
                _ = v32
            case 4294967295:
                var v33 uint16 = (v6 | 255)
                _ = v33
            case -2147483648:
                v34 := GenS3{F0: v7, F1: 32767, F2: ("foo bar" + "\n\t")}
                _ = v34
            default:
                v2["hello"] = false
            }
        }
        v8 = v8
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
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
