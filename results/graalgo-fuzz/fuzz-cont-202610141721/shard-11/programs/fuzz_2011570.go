package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 bool }

func main() {
    v1 := []uint32{2147483648, 65535, 65535}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 uint64 = 0
    var v5 string = "a"
    var v6 uint = 0
    if (true && (27 <= 0)) {
        fmt.Println(v4, v6, v5)
    } else {
        var v7 int = 9223372036854775807
        _ = v7
    }
    v1 = append(v1, 152)
    v8 := GenS1{F0: 4294967295, F1: "", F2: 9223372036854775807}
    v4 -= v4
    if true {
        var v9 string = v5
        _ = len(v1)
        _ = v9
    }
    var v10 int64 = -9223372036854775808
    switch "hello" {
    case "":
        switch "" {
        case "":
            var v11 bool = true
            _ = v11
        case "foo bar":
            switch v5 {
            case "foo bar":
                var v12 uint = (^(v6))
                _ = v12
            default:
                var v13 uint16 = 255
                _ = v13
            }
        }
    case "hello":
        v14 := GenS1{F0: (v4 % (((v4 - (v4 - (v4 & v4)))) | 1)), F1: (v5 + v5), F2: 4294967295}
        _ = v14
    case "\n\t":
        v6 = (v6 ^ (v6 * 255))
    default:
        v15 := GenS1{F0: v4, F1: v5, F2: 2147483647}
        _ = v15
    }
    if (!false) {
        var v16 int16 = 1000
        switch v5 {
        case "\n\t":
            v1 = append(v1, 4294967295)
        }
        _ = v16
    } else {
        switch ((v5 + v5) + v5) {
        case "foo bar":
            if true {
                for fi17 := 0; fi17 < 5; fi17++ {
                    var v18 float64 = 0.0
                    for fi19 := 0; fi19 < 1; fi19++ {
                        var v20 uint8 = 2
                        switch fi17 {
                        case 128:
                            v6 = (v6 << ((v6) & 63))
                        default:
                            var v21 int = -2147483648
                            _ = v21
                        }
                        _ = fi19
                        _ = v20
                    }
                    if (0 == 1) {
                        var v22 uint8 = 42
                        switch (fi17 % ((int(v22)) | 1)) {
                        case 2147483647:
                            v18 = 0.0
                        case 0:
                            for fi23 := 0; fi23 < 1; fi23++ {
                                if true {
                                    v2[2147483648] = 4294967295
                                    fmt.Println(v18, v4, v10)
                                    v24 := GenS2{F0: v4}
                                    _ = v24
                                } else {
                                    var v25 float64 = (v18 + ((0.0 * v18) + v18))
                                    v26 := GenS2{F0: (v4 % (((v4 * v4)) | 1))}
                                    _ = v25
                                    _ = v26
                                }
                                v22 = v22
                                _ = fi23
                            }
                        case 9223372036854775807:
                            var v27 uint32 = 4294967295
                            _ = v27
                        default:
                            for fi28 := 0; fi28 < 15; fi28++ {
                                var v29 string = ("" + v5)
                                _ = fi28
                                _ = v29
                            }
                        }
                        _ = v22
                    } else {
                        v30 := GenS2{F0: v4}
                        delete(v2, 42)
                        _ = v30
                    }
                    _ = fi17
                    _ = v18
                }
                v10 = -1
            }
        case "":
            if true {
                var v31 uint = (v6 | v6)
                _ = v31
            } else {
                if (0 < 882) {
                    var v32 uint64 = v4
                    _ = v32
                }
            }
        }
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        var v34 int32 = 410
        v5 = ""
        _ = fi33
        _ = v34
    }
    v35 := GenS3{F0: int16(v4), F1: (!(!true))}
    v10 = (^(v10))
    fmt.Println(v5, v6, v4)
    var v36 uint16 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
