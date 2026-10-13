package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint32; F1 bool; F2 int32 }

func main() {
    var v1 float64 = (3.14159 / 1e10)
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    v4 := []string{"\n\t"}
    v5 := []uint32{2147483647, 393}
    var v6 uint16 = 42
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 bool = false
        if ((v6 << ((255) & 63)) == v6) {
            v9 := GenS1{F0: v6}
            for fi10 := 0; fi10 < 19; fi10++ {
                fmt.Println(fi10, v6, v6)
                if true {
                    if false {
                        var v11 uint = 214
                        switch "foo bar" {
                        case "\n\t":
                            var v12 uint8 = 128
                            _ = v12
                        }
                        _ = v11
                    } else {
                        v1 = 3.14159
                        var v13 int64 = 1
                        fmt.Println(fi7)
                        _ = v13
                    }
                }
                v4 = append(v4, "\n\t")
                if ("hello" >= "a") {
                    fmt.Println(v8, fi7, v8)
                    v1 += v1
                    v6 += 874
                    var v14 float64 = v1
                    _ = len(v4)
                    fmt.Println(v6)
                    v1 = v1
                    if false {
                        fmt.Println(fi10)
                        delete(v3, (fi10 >> ((uint(v6)) & 63)))
                    }
                    v15 := GenS1{F0: (v6 / (((v6 % ((v6) | 1))) | 1))}
                    var v16 uint16 = v6
                    fmt.Println(v6, v8, v6)
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = fi10
            }
            if (!v8) {
                var v17 float32 = 1.5
                v18 := GenS2{F0: 0, F1: -2147483648, F2: 670}
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 int32 = -2147483648
                    v20 = (v20 + -2147483648)
                    _ = fi19
                    _ = v20
                }
                _ = v17
                _ = v18
            }
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v21 bool = true
    v21 = false
    v22 := GenS1{F0: v6}
    v2["foo bar"] = 9223372036854775807
    for fi23 := 0; fi23 < 7; fi23++ {
        v21 = (897 != 4294967295)
        _ = fi23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            if true {
                var v26 bool = v21
                if v26 {
                    if ((v6 ^ (v6 << ((2) & 63))) >= 42) {
                        if (-32768 <= -32768) {
                            var v27 bool = (100 == 42)
                            _ = v27
                        } else {
                            v2[("foo bar" + ("foo bar" + "\n\t"))] = fi25
                        }
                        v28 := GenS3{F0: 42, F1: 430}
                        var v29 uint16 = 0
                        _ = v28
                        _ = v29
                    }
                } else {
                    if v21 {
                        var v30 uint = 255
                        _ = v30
                    }
                }
                v31 := GenS2{F0: 156, F1: fi24, F2: 2}
                _ = v26
                _ = v31
            } else {
                _ = len(v2)
                var v32 uint64 = 42
                fmt.Println(fi25, v32, v32)
                _ = v32
            }
            var v33 int8 = 127
            _ = fi25
            _ = v33
        }
        for fi34 := 0; fi34 < 17; fi34++ {
            var v35 float32 = (1.5 + 0.0)
            _ = fi34
            _ = v35
        }
        _ = fi24
    }
    v6 += (v6 & 32767)
    for fi36 := 0; fi36 < 6; fi36++ {
        v1 += 0.0
        switch "" {
        case "":
            _ = len(v2)
        default:
            switch (fi36 - (fi36 & (-(fi36)))) {
            case 2147483647:
                for fi37 := 0; fi37 < 9; fi37++ {
                    var v38 float32 = float32(v1)
                    fmt.Println(v21, v6, fi36)
                    switch fi37 {
                    case 32767:
                        if true {
                            var v39 float64 = (v1 / float64(v38))
                            if ((fi36 ^ (fi36 / ((fi37) | 1))) == fi36) {
                                var v40 float64 = (float64(1e10) - -2.25)
                                var v41 uint32 = 2147483648
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v40
                                _ = v41
                            } else {
                                v42 := GenS3{F0: 217, F1: uint32(v6)}
                                v1 = v39
                                _ = v42
                            }
                            switch fi37 {
                            case 9223372036854775807:
                                var v43 float64 = v1
                                _ = v43
                            case -423:
                                var v44 string = ""
                                _ = v44
                            case 1:
                                if v21 {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v45 string = "\n\t"
                                    var v46 uint = 144
                                    fmt.Println(v1, v21)
                                    _ = v45
                                    _ = v46
                                }
                            }
                            switch 127 {
                            case -2147483648:
                                var v47 bool = v21
                                _ = v47
                            case 2147483647:
                                if false {
                                    var v48 int16 = -1
                                    if (!false) {
                                        v6 -= v6
                                    }
                                    _ = v48
                                }
                            }
                            fmt.Println(fi37, v38)
                            _ = v39
                        }
                    }
                    _ = fi37
                    _ = v38
                }
            case 9223372036854775807:
                _ = len(v3)
            case 1:
                v49 := GenS4{F0: 732, F1: v21, F2: -847}
                _ = v49
            default:
                var v50 uint = 4294967295
                _ = v50
            }
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
