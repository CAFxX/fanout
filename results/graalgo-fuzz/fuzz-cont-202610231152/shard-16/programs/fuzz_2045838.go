package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int8; F2 int16 }
type GenS4 struct { F0 uint; F1 bool; F2 uint32 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v4 float64 = 3.14159
    v5 := []int{-9223372036854775808, 128}
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 979}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            v4 = v4
            _ = fi9
        }
        _ = fi6
    }
    v10 := GenS2{F0: ""}
    switch (("\n\t" + "hello") + "hello") {
    case "foo bar":
        if (0.1 > ((3.14159 + float32(1e10)) / 1.0)) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            switch 128 {
            case 128:
                var v11 uint32 = 569
                _ = v11
            case -348:
                v4 = 0.0
            }
        }
    case "":
        var v12 float32 = 3.14159
        _ = v12
    default:
        v13 := GenS3{F0: -32768, F1: 46, F2: 36}
        _ = v13
    }
    fmt.Println(v4)
    var v14 string = ""
    v15 := GenS3{F0: -1, F1: 42, F2: -139}
    switch v14 {
    case "a":
        for fi16 := 0; fi16 < 11; fi16++ {
            switch (v14 + "hello") {
            case "":
                v17 := GenS4{F0: 4294967295, F1: true, F2: 65535}
                _ = v17
            case "a":
                if (((1 == 1) || false) && true) {
                    for fi18 := 0; fi18 < 8; fi18++ {
                        _ = len(v5)
                        var v19 uint64 = 42
                        _ = fi18
                        _ = v19
                    }
                }
            }
            v4 += v4
            switch "" {
            case "foo bar":
                v1[""] = (fi16 << ((222) & 63))
            default:
                delete(v3, 12)
            }
            _ = fi16
        }
    case "":
        for fi20 := 0; fi20 < 15; fi20++ {
            v14 = v14
            fmt.Println(fi20)
            _ = fi20
        }
    case "foo bar":
        _ = len(v5)
    }
    v21 := GenS1{F0: 0}
    var v22 bool = (2 == 1)
    fmt.Println(v22)
    v14 = ("" + "")
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            switch "a" {
            case "\n\t":
                v25 := GenS5{F0: -380}
                _ = v25
            case "foo bar":
                switch fi24 {
                case 1:
                    v26 := GenS1{F0: 42}
                    _ = v26
                case 4294967295:
                    v22 = false
                    var v27 uint64 = 753
                    _ = v27
                }
            }
            v28 := GenS3{F0: 516, F1: -128, F2: 0}
            _ = fi24
            _ = v28
        }
        var v29 uint32 = 2147483647
        delete(v2, 134)
        _ = fi23
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        delete(v2, 1)
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
