package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []string{"\n\t"}
    var v2 int = 32767
    v3 := make(map[uint32]uint32)
    for fi4 := 0; fi4 < 10; fi4++ {
        fmt.Println(v2)
        for fi5 := 0; fi5 < 14; fi5++ {
            switch ("" + "\n\t") {
            case "\n\t":
                v6 := GenS1{F0: -9223372036854775808, F1: 1}
                _ = v6
            case "a":
                switch "foo bar" {
                case "":
                    var v7 int16 = 32767
                    _ = v7
                case "hello":
                    fmt.Println(v2, fi5, v2)
                case "\n\t":
                    v2 += fi5
                }
                v2 = (fi4 & 2147483647)
                if true {
                    v2 -= fi4
                } else {
                    v8 := GenS2{F0: 0}
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(fi4, v2, fi4)
                    }
                    _ = len(v3)
                    _ = v8
                }
            default:
                v2 += -1
            }
            _ = fi5
        }
        var v9 int64 = 339
        _ = fi4
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v2 = v2
        var v11 string = "\n\t"
        _ = fi10
        _ = v11
    }
    switch v2 {
    case 417:
        fmt.Println(v2)
    case -1:
        _ = len(v1)
    }
    v3[223] = 2147483648
    fmt.Println(v2)
    delete(v3, 42)
    v12 := GenS2{F0: 255}
    fmt.Println(v2)
    fmt.Println(v2)
    var v13 bool = true
    var v14 int64 = -1
    for fi15 := 0; fi15 < 9; fi15++ {
        v13 = false
        v13 = (!(42 > 1))
        v2 -= (v2 / (((v2 + (v2 + -2147483648))) | 1))
        _ = fi15
    }
    v16 := GenS2{F0: 174}
    if v13 {
        switch ("foo bar" + (("a" + ("a" + "hello")) + "\n\t")) {
        case "":
            var v17 int16 = 0
            _ = v17
        case "hello":
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 uint8 = 118
                _ = fi18
                _ = v19
            }
        default:
            fmt.Println(v14, v2)
        }
    } else {
        v13 = false
        v20 := GenS3{F0: -2147483648}
        var v21 int64 = v14
        _ = v20
        _ = v21
    }
    v13 = v13
    if true {
        var v22 int8 = -1
        _ = v22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        if v13 {
            fmt.Println(v2, fi23, fi23)
            var v24 int8 = -1
            v13 = ((42 == 2147483647) && (3.14159 > (0.1 / 1.5)))
            _ = v24
        } else {
            var v25 int32 = 2147483647
            _ = v25
        }
        fmt.Println(v13, v13)
        fmt.Println(v14)
        _ = fi23
    }
    if (255 > 65535) {
        fmt.Println(v13, v2)
        var v26 uint16 = 318
        fmt.Println(v26)
        _ = v26
    } else {
        if v13 {
            if false {
                v13 = v13
            }
        }
        var v27 uint64 = 0
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
}
