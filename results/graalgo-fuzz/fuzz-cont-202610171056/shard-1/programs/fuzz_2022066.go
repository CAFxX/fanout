package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint32 = 2147483648
    var v2 int16 = 0
    var v3 int16 = v2
    v4 := make(map[int]string)
    v5 := []uint32{(v1 << ((1) & 63))}
    var v6 uint16 = 0
    v5 = append(v5, 4294967295)
    v7 := v4[4294967295]
    if (0 >= 5) {
        v5 = append(v5, v1)
        delete(v4, 127)
        v1 += (v1 * v1)
    } else {
        var v8 uint = 18446744073709551615
        switch "foo bar" {
        case "\n\t":
            v8 += (v8 / (((^(v8))) | 1))
        case "a":
            var v9 uint = 65535
            _ = v9
        default:
            switch 127 {
            case 0:
                v8 -= v8
            case -1:
                v8 = v8
            case 42:
                var v10 int16 = 32767
                _ = v10
            }
        }
        v11 := GenS1{F0: (-2.25 == float64(0.1)), F1: v8}
        _ = v11
        _ = v8
    }
    fmt.Println(v7)
    var v12 float32 = float32(1e10)
    fmt.Println(v12, v2, v6)
    if true {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint16 = (v6 >> ((65535) & 63))
            if (int8(v3) < 1) {
                var v15 uint64 = 1
                _ = v15
            } else {
                var v16 float32 = v12
                _ = v16
            }
            v17 := GenS1{F0: ((false && (!(true && (-2147483648 >= 413)))) && false), F1: 18446744073709551615}
            _ = fi13
            _ = v14
            _ = v17
        }
        v18 := GenS1{F0: ("foo bar" != ((v7 + "foo bar") + ((v7 + ((v7 + v7) + (v7 + v7))) + v7))), F1: 840}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS2{F0: (v6 / (((v6 * (v6 >> ((255) & 63)))) | 1))}
            _ = fi19
            _ = v20
        }
        if (127 > 1) {
            if false {
                var v21 uint = 4294967295
                v22 := GenS2{F0: (^(v6))}
                var v23 uint = v21
                _ = v21
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                switch (fi24 - (fi24 % ((2147483647) | 1))) {
                case 78:
                    v3 = v2
                case 32767:
                    fmt.Println(v6, v3, fi24)
                default:
                    v12 = float32((3.14159 - float64(0.1)))
                }
                var v25 int32 = 0
                v2 = (v3 * -32768)
                _ = fi24
                _ = v25
            }
        } else {
            var v26 bool = (3.14159 < 3.14159)
            if v26 {
                for fi27 := 0; fi27 < 13; fi27++ {
                    v12 = v12
                    var v28 int8 = 58
                    switch v7 {
                    case "hello":
                        var v29 string = ""
                        _ = v29
                    case "\n\t":
                        var v30 int64 = 42
                        _ = v30
                    case "foo bar":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        var v31 uint16 = uint16(v1)
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
            } else {
                v2 = v2
                fmt.Println(v12)
                var v32 int64 = 9223372036854775807
                _ = v32
            }
            v4[32768] = "hello"
            _ = v26
        }
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 uint32 = 698
            v35 := GenS2{F0: v6}
            v1 = v1
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    switch v7 {
    case "\n\t":
        for fi36 := 0; fi36 < 15; fi36++ {
            if true {
                if (!true) {
                    fmt.Println(v3, v2)
                } else {
                    var v37 float64 = float64(v12)
                    _ = v37
                }
                if (true || true) {
                    switch fi36 {
                    case -9223372036854775808:
                        for fi38 := 0; fi38 < 18; fi38++ {
                            var v39 float32 = v12
                            v39 += (float32((1.5 * 1e10)) * ((v12 / v39) + v39))
                            _ = fi38
                            _ = v39
                        }
                    case -1:
                        if ((fi36 + fi36) != -2147483648) {
                            for fi40 := 0; fi40 < 10; fi40++ {
                                fmt.Println(v3, fi36)
                                if (4294967295 < 1) {
                                    switch fi40 {
                                    case -778:
                                        v41 := GenS1{F0: true, F1: 197}
                                        _ = v41
                                    case 1:
                                        v2 = int16(fi36)
                                    case 127:
                                        switch (fi40 * fi36) {
                                        case 318:
                                            v5 = append(v5, v1)
                                        case -140:
                                            fmt.Println(v12, v6, v6)
                                        case 0:
                                            var v42 int32 = int32(v1)
                                            _ = v42
                                        }
                                    default:
                                        v43 := GenS3{F0: (0 < 1)}
                                        _ = v43
                                    }
                                } else {
                                    if (v7 < v7) {
                                        var v44 int8 = 20
                                        v44 = (v44 & v44)
                                        switch (fi36 >> ((2) & 63)) {
                                        case 42:
                                            var v45 bool = ((v1 >> ((255) & 63)) >= 65535)
                                            _ = v45
                                        }
                                        _ = v44
                                    }
                                }
                                _ = fi40
                            }
                            var v46 uint8 = 128
                            v3 -= v3
                            _ = v46
                        }
                    case 719:
                        v3 += (v2 + 32767)
                    default:
                        fmt.Println(v7, fi36)
                    }
                    fmt.Println(v1)
                }
                fmt.Println(v7, v3, v7)
            } else {
                if (!(false || false)) {
                    v7 = v7
                    switch ((v7 + v7) + v7) {
                    case "hello":
                        switch (fi36 % (((fi36 - 128)) | 1)) {
                        case 127:
                            if true {
                                for fi47 := 0; fi47 < 5; fi47++ {
                                    if len(v5) > 0 {
                                        v5[0] = v1
                                    }
                                    v4[-795] = v7
                                    v3 -= (v2 % ((-394) | 1))
                                    _ = fi47
                                }
                                v12 = (v12 + v12)
                                v48 := GenS1{F0: (!true), F1: 663}
                                _ = v48
                            } else {
                                v1 += v1
                                v49 := GenS2{F0: v6}
                                for fi50 := 0; fi50 < 11; fi50++ {
                                    v7 = (v7 + v7)
                                    _ = fi50
                                }
                                _ = v49
                            }
                        case 32767:
                            for fi51 := 0; fi51 < 15; fi51++ {
                                var v52 int32 = 702
                                _ = fi51
                                _ = v52
                            }
                        case 32768:
                            if true {
                                v2 -= -32768
                                var v53 float64 = 1e10
                                for fi54 := 0; fi54 < 9; fi54++ {
                                    fmt.Println(fi54, v6, v1)
                                    v1 -= (v1 - 2147483647)
                                    _ = fi54
                                }
                                _ = v53
                            }
                        }
                    case "foo bar":
                        fmt.Println(v2)
                    default:
                        var v55 float32 = ((v12 / (v12 / (3.14159 + (v12 / float32((1.5 - 3.14159)))))) + -2.25)
                        _ = v55
                    }
                }
                switch "a" {
                case "hello":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            _ = fi36
        }
    case "foo bar":
        fmt.Println(v12, v2, v2)
    default:
        var v56 int8 = 127
        _ = v56
    }
    var v57 uint64 = 42
    var v58 int32 = -471
    if len(v5) > 0 {
        v5[0] = uint32(v57)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
    _ = v7
}
