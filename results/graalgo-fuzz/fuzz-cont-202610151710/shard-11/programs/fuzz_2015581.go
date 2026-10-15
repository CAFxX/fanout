package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint }
type GenS5 struct { F0 uint }

func main() {
    var v1 uint8 = 7
    v2 := []uint32{1, 106, 2147483648}
    var v3 float32 = 1e10
    var v4 uint8 = v1
    var v5 int64 = 460
    var v6 int32 = -1
    if (!true) {
        fmt.Println(v4, v1)
        v2 = append(v2, 2147483647)
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v7 := GenS1{F0: (v5 | v5), F1: 316}
    fmt.Println(v5)
    v2 = append(v2, 125)
    v2 = append(v2, 2)
    fmt.Println(v4, v4, v6)
    switch "foo bar" {
    case "":
        switch 2147483647 {
        case 128:
            v2 = append(v2, 156)
        case 32767:
            var v8 string = ""
            _ = v8
        case 649:
            v4 += 68
        }
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v4, v4)
        _ = fi9
    }
    fmt.Println(v3, v4)
    v3 = v3
    v4 = v4
    if ((-32768 == 1000) || (!false)) {
        switch ("" + "hello") {
        case "foo bar":
            fmt.Println(v4)
        case "":
            fmt.Println(v5)
        case "a":
            var v10 int = 42
            _ = v10
        default:
            for fi11 := 0; fi11 < 4; fi11++ {
                fmt.Println(v5, fi11)
                _ = fi11
            }
        }
    }
    fmt.Println(v6)
    fmt.Println(v3)
    if false {
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := GenS1{F0: 0, F1: 1}
                fmt.Println(v6)
                var v15 string = "hello"
                _ = fi13
                _ = v14
                _ = v15
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi12
        }
    } else {
        var v16 string = "foo bar"
        _ = v16
    }
    fmt.Println(v6, v5)
    var v17 int = -966
    if (!false) {
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(v4, v5, v1)
            v19 := GenS1{F0: v5, F1: 811}
            var v20 uint = 18446744073709551615
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = len(v2)
        for fi21 := 0; fi21 < 18; fi21++ {
            if true {
                var v22 uint64 = 2
                switch v17 {
                case 4294967295:
                    var v23 uint16 = 32767
                    _ = v23
                case -9223372036854775808:
                    switch (v17 - v17) {
                    case -656:
                        v24 := GenS2{F0: 2, F1: (v17 & (v17 ^ (fi21 / (((v17 ^ 594)) | 1))))}
                        _ = v24
                    default:
                        switch -203 {
                        case -106:
                            for fi25 := 0; fi25 < 15; fi25++ {
                                fmt.Println(v17, v5, fi21)
                                _ = fi25
                            }
                        case 2147483647:
                            var v26 float32 = -2.25
                            _ = v26
                        }
                    }
                }
                _ = v22
            } else {
                for fi27 := 0; fi27 < 11; fi27++ {
                    if false {
                        v28 := GenS3{F0: 18446744073709551615, F1: (v5 * (v5 + v5))}
                        _ = v28
                    } else {
                        if false {
                            v29 := GenS4{F0: 65535, F1: 651, F2: 65535}
                            _ = v29
                        } else {
                            var v30 int64 = v5
                            fmt.Println(v1, fi21, v3)
                            var v31 float32 = v3
                            _ = v30
                            _ = v31
                        }
                        if (82 < 953) {
                            var v32 int16 = 178
                            _ = v32
                        }
                        _ = len(v2)
                        if true {
                            fmt.Println(v5, v17)
                        } else {
                            switch (fi21 ^ 128) {
                            case 9223372036854775807:
                                v33 := GenS4{F0: 255, F1: 255, F2: 42}
                                _ = v33
                            case 42:
                                fmt.Println(fi27, fi21)
                            }
                            var v34 uint64 = 4294967295
                            _ = v34
                        }
                    }
                    var v35 int = 1
                    var v36 uint = uint(v4)
                    _ = fi27
                    _ = v35
                    _ = v36
                }
            }
            if (true && (0 <= 255)) {
                var v37 uint16 = 255
                if (((-2.25 + -2.25) < (3.14159 * 1e10)) && false) {
                    var v38 float64 = float64(0.0)
                    if true {
                        v37 = v37
                        var v39 string = ""
                        fmt.Println(v38, v37, v38)
                        _ = v39
                    } else {
                        _ = len(v2)
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v38
                } else {
                    _ = len(v2)
                }
                for fi40 := 0; fi40 < 9; fi40++ {
                    var v41 int64 = v5
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v42 int64 = v5
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = v37
            } else {
                v1 -= v1
                var v43 uint32 = 42
                _ = v43
            }
            _ = fi21
        }
    } else {
        v44 := GenS5{F0: 4294967295}
        var v45 float64 = 0.1
        v17 = (v17 ^ (v17 & v17))
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
