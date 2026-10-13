package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int64; F2 int32 }

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]string)
    v3 := make(map[string]bool)
    var v4 float32 = -2.25
    switch -1 {
    case 127:
        v5 := GenS1{F0: 2147483647, F1: 42}
        _ = v5
    case 32767:
        var v6 uint8 = (v1 | v1)
        _ = v6
    }
    v1 += (v1 - (v1 / ((v1) | 1)))
    v7 := []int{795, 32768}
    v8 := GenS1{F0: 0, F1: 42}
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 string = ""
        v4 = (v4 - v4)
        var v11 float64 = 0.1
        _ = fi9
        _ = v10
        _ = v11
    }
    v1 -= v1
    if (true || (454 >= 32767)) {
        v4 -= v4
        v2[2] = ("a" + "hello")
        var v12 int64 = -9223372036854775808
        _ = v12
    }
    var v13 uint16 = 42
    switch -2147483648 {
    case -541:
        var v14 uint8 = v1
        _ = v14
    case 824:
        if false {
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 string = "a"
                var v17 float64 = 3.14159
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 float32 = v4
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 float64 = 3.14159
                var v21 int16 = -309
                for fi22 := 0; fi22 < 8; fi22++ {
                    v21 = -336
                    v1 -= (v1 + v1)
                    if (-648 < -1) {
                        switch -585 {
                        case -370:
                            v4 -= float32((v20 - v20))
                        case -9223372036854775808:
                            v23 := GenS1{F0: fi19, F1: (-(v21))}
                            _ = v23
                        default:
                            v24 := GenS1{F0: fi19, F1: 300}
                            _ = v24
                        }
                    } else {
                        if len(v7) > 0 {
                            v7[0] = (fi19 << ((255) & 63))
                        }
                    }
                    if (true && (127 >= 42)) {
                        delete(v2, 455)
                    } else {
                        switch "hello" {
                        case "foo bar":
                            v13 = v13
                        case "hello":
                            v20 = (v20 / v20)
                        }
                        var v25 uint64 = 1
                        _ = v25
                    }
                    _ = fi22
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = v18
        } else {
            var v26 int16 = 1000
            _ = v26
        }
    case 128:
        switch 281 {
        case 1:
            var v27 string = "hello"
            _ = v27
        default:
            v4 += v4
        }
    }
    fmt.Println(v1)
    if (false && true) {
        var v28 uint16 = 65535
        var v29 int = 42
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        if true {
            var v31 string = ""
            _ = len(v2)
            fmt.Println(v31, v13)
            _ = v31
        } else {
            if false {
                switch (fi30 % ((fi30) | 1)) {
                case 127:
                    if false {
                        v32 := GenS2{F0: (^(v1))}
                        var v33 uint64 = 2
                        if true {
                            var v34 uint64 = (v33 % ((v33) | 1))
                            _ = v34
                        }
                        _ = v32
                        _ = v33
                    } else {
                        if (((^(fi30)) > (fi30 % ((791) | 1))) || (int(v13) >= fi30)) {
                            var v35 int32 = 47
                            if true {
                                v1 -= v1
                                fmt.Println(v1, v35, fi30)
                            }
                            _ = v35
                        } else {
                            v2[490] = "foo bar"
                            v4 = v4
                            if ("\n\t" <= "foo bar") {
                                var v36 float32 = v4
                                var v37 float64 = 0.1
                                for fi38 := 0; fi38 < 7; fi38++ {
                                    v4 = (float32(0.1) / (v4 * v4))
                                    var v39 string = (("hello" + "a") + ("hello" + ""))
                                    v40 := v3[(v39 + ("" + v39))]
                                    switch v39 {
                                    case "":
                                        v41 := GenS1{F0: fi30, F1: 302}
                                        _ = v41
                                    case "a":
                                        _ = len(v2)
                                    default:
                                        var v42 bool = (v40 || true)
                                        _ = v42
                                    }
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    v37 -= (v37 * -2.25)
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                                _ = v36
                                _ = v37
                            } else {
                                switch ("" + "foo bar") {
                                case "foo bar":
                                    v43 := GenS3{F0: -546, F1: 0, F2: 371}
                                    _ = v43
                                }
                            }
                        }
                        fmt.Println(fi30, fi30)
                    }
                case -1:
                    for fi44 := 0; fi44 < 16; fi44++ {
                        fmt.Println(fi44, fi44)
                        var v45 int8 = 0
                        _ = fi44
                        _ = v45
                    }
                case 42:
                    v13 = 0
                default:
                    var v46 int8 = -128
                    _ = v46
                }
            } else {
                v4 = float32((3.14159 / 1.5))
                v1 -= (v1 | v1)
                delete(v3, "foo bar")
            }
            for fi47 := 0; fi47 < 7; fi47++ {
                var v48 uint32 = 42
                var v49 int64 = 38
                switch "foo bar" {
                case "foo bar":
                    v48 = 42
                case "":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                default:
                    fmt.Println(v48, v49)
                    if (int16(v48) == 1000) {
                        var v50 string = "\n\t"
                        _ = v50
                    }
                }
                _ = fi47
                _ = v48
                _ = v49
            }
        }
        for fi51 := 0; fi51 < 9; fi51++ {
            v52 := GenS1{F0: 861, F1: 1000}
            v13 += 486
            var v53 int = (fi51 / ((fi30) | 1))
            _ = fi51
            _ = v52
            _ = v53
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
