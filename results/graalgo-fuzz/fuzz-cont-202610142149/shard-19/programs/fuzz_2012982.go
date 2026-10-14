package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []string{("" + "a"), "\n\t", "", ("hello" + "\n\t")}
    v2 := make(map[uint32]string)
    var v3 uint16 = 65535
    var v4 int = 4294967295
    switch "\n\t" {
    case "hello":
        switch v4 {
        case 9223372036854775807:
            v4 = v4
        case 0:
            if false {
                fmt.Println(v4, v3)
                v4 = (v4 >> ((0) & 63))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v3 = (^(v3))
            }
        case -738:
            var v5 uint32 = 2147483648
            _ = v5
        }
    case "foo bar":
        fmt.Println(v3, v4)
    case "\n\t":
        v6 := GenS1{F0: (1 == 42)}
        _ = v6
    default:
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 string = "a"
            switch 371 {
            case 4294967295:
                for fi9 := 0; fi9 < 9; fi9++ {
                    switch (^(fi7)) {
                    case 4294967295:
                        v4 = fi7
                    case 127:
                        v4 = v4
                    default:
                        var v10 uint16 = v3
                        delete(v2, 65535)
                        var v11 int8 = 100
                        _ = v10
                        _ = v11
                    }
                    if (true && (100 < 100)) {
                        var v12 uint16 = (v3 / (((^(v3))) | 1))
                        _ = v12
                    } else {
                        var v13 int64 = -435
                        if ((true && false) || (-128 <= -1)) {
                            if true {
                                switch (v8 + ("\n\t" + v8)) {
                                case "":
                                    v14 := GenS1{F0: ((3.14159 + 1.5) <= 1e10)}
                                    _ = v14
                                }
                                if false {
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                            } else {
                                v13 = -935
                                switch (fi7 << ((42) & 63)) {
                                case -1:
                                    if len(v1) > 0 {
                                        v1[0] = v8
                                    }
                                case 1:
                                    switch v8 {
                                    case "\n\t":
                                        v8 = "hello"
                                    case "hello":
                                        if (18446744073709551615 != 1) {
                                            v8 = v8
                                        }
                                    case "":
                                        var v15 float32 = 1.5
                                        _ = v15
                                    }
                                }
                                var v16 int64 = v13
                                var v17 bool = true
                                _ = v16
                                _ = v17
                            }
                        }
                        switch (v8 + ("hello" + v8)) {
                        case "foo bar":
                            v18 := v2[1]
                            _ = v18
                        case "a":
                            v8 = v8
                            switch v4 {
                            case 820:
                                v19 := GenS1{F0: (false || false)}
                                _ = v19
                            }
                            v20 := GenS2{F0: true, F1: (!(-32768 == -638)), F2: 1}
                            _ = v20
                        }
                        v13 = v13
                        v21 := GenS1{F0: (127 <= 0)}
                        if len(v1) > 0 {
                            v1[0] = v8
                        }
                        var v22 uint16 = (v3 << ((77) & 63))
                        v23 := GenS1{F0: (299 <= 975)}
                        _ = v13
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = fi9
                }
            default:
                var v24 float64 = (3.14159 - 1.5)
                _ = v24
            }
            _ = fi7
            _ = v8
        }
    }
    v2[65535] = ""
    switch (v4 % (((v4 / (((v4 % (((v4 % ((v4) | 1))) | 1))) | 1))) | 1)) {
    case 0:
        delete(v2, 527)
    }
    var v25 int8 = 1
    if true {
        switch v4 {
        case 128:
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 float64 = 0.1
                v2[944] = ("" + "hello")
                _ = fi26
                _ = v27
            }
        }
        _ = len(v2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v28 string = "foo bar"
    for fi29 := 0; fi29 < 6; fi29++ {
        _ = len(v2)
        _ = fi29
    }
    v30 := GenS3{F0: (^(v3))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v30
    _ = v4
}
