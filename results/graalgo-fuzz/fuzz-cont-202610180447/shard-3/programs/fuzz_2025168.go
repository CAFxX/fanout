package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = -95
    v3 := []uint32{1, 4294967295, 1}
    var v4 uint64 = 1
    v4 = 544
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 string = ((("hello" + "\n\t") + "\n\t") + "hello")
        var v7 bool = ((255 == 32767) && (fi5 > (fi5 >> ((255) & 63))))
        if true {
            switch v6 {
            case "hello":
                if v7 {
                    fmt.Println(v2)
                } else {
                    fmt.Println(v2)
                    for fi8 := 0; fi8 < 12; fi8++ {
                        if (!(!v7)) {
                            switch (fi5 + fi5) {
                            case 127:
                                v4 = (v4 / (((v4 / (((v4 + v4)) | 1))) | 1))
                            }
                            _ = len(v3)
                            switch (v6 + v6) {
                            case "":
                                _ = len(v3)
                            }
                        } else {
                            if v7 {
                                v2 = -107
                            } else {
                                v6 = v6
                            }
                            fmt.Println(fi8, fi5)
                        }
                        var v9 uint = 18446744073709551615
                        var v10 int = 2147483647
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
            }
            fmt.Println(v4, fi5)
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if false {
        var v11 int8 = v2
        v2 -= (^(v2))
        _ = v11
    } else {
        v2 += -128
        switch 127 {
        case -1:
            _ = len(v3)
        case -994:
            v4 = v4
        }
    }
    v4 = 329
    v12 := GenS1{F0: 0, F1: 1, F2: 2147483647}
    switch "hello" {
    case "hello":
        v13 := GenS1{F0: 255, F1: 0, F2: 636}
        _ = v13
    case "foo bar":
        if false {
            var v14 string = (("hello" + ("a" + "\n\t")) + "")
            _ = v14
        }
    case "\n\t":
        v2 = (v2 << ((446) & 63))
    }
    if false {
        switch 984 {
        case 948:
            fmt.Println(v2)
        case 128:
            fmt.Println(v2)
        default:
            switch ("foo bar" + "hello") {
            case "foo bar":
                for fi15 := 0; fi15 < 16; fi15++ {
                    if (243 >= 72) {
                        v3 = append(v3, 2147483648)
                        for fi16 := 0; fi16 < 1; fi16++ {
                            v4 = v4
                            var v17 int64 = 1
                            _ = fi16
                            _ = v17
                        }
                    }
                    _ = fi15
                }
            }
        }
        var v18 float64 = (-2.25 * (0.0 * 3.14159))
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        v2 = v2
        if true {
            if false {
                fmt.Println(v2, v4)
            }
            v4 = v4
        }
        _ = fi19
    }
    var v20 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
