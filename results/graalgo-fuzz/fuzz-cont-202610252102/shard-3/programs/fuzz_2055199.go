package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 string }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = (255 >= 2)
    var v3 bool = true
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 11; fi5++ {
        if (v2 && (-2147483648 == -776)) {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 bool = v3
                delete(v4, 35)
                if (v7 && true) {
                    fmt.Println(fi6, v2, v2)
                    v1[(^(fi5))] = ("" + "\n\t")
                } else {
                    if v3 {
                        fmt.Println(v7, v2, v7)
                    }
                }
                _ = fi6
                _ = v7
            }
            _ = len(v4)
        } else {
            for fi8 := 0; fi8 < 10; fi8++ {
                v3 = (!(!(!(!true))))
                var v9 uint64 = 42
                v10 := []int{733, 636, fi8}
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                var v11 float64 = (float64(1.5) * 0.1)
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v3 = v2
            switch "a" {
            case "foo bar":
                v12 := []int{-458, fi5, fi5, (^(fi5))}
                _ = v12
            case "\n\t":
                v13 := []int{(-(fi5)), 1}
                _ = v13
            case "hello":
                v2 = (42 > 0)
            default:
                var v14 uint8 = 0
                _ = v14
            }
        }
        _ = fi5
    }
    var v15 float32 = -2.25
    var v16 uint64 = 4294967295
    if (((957 == -87) || v3) || v3) {
        var v17 uint = 2
        _ = v17
    } else {
        if (v15 == 1.5) {
            var v18 float32 = float32(0.1)
            _ = v18
        }
    }
    fmt.Println(v2)
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v19 int = 42
        _ = v19
    case 128:
        var v20 float32 = 3.14159
        _ = v20
    case -2147483648:
        v21 := GenS1{F0: 65535, F1: (-117 == -128), F2: ""}
        _ = v21
    default:
        fmt.Println(v2, v2, v15)
    }
    v16 = (v16 * 4294967295)
    var v22 int64 = 42
    for fi23 := 0; fi23 < 10; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            _ = len(v4)
            var v25 string = "\n\t"
            switch (^(fi23)) {
            case -213:
                fmt.Println(v22, fi24, v3)
            case -47:
                v26 := GenS2{F0: -553}
                _ = v26
            case 9223372036854775807:
                switch (fi24 + fi23) {
                case 4294967295:
                    var v27 int32 = 42
                    _ = v27
                case 128:
                    var v28 bool = false
                    _ = v28
                case -880:
                    v29 := []int{-557, fi24}
                    var v30 uint8 = 0
                    _ = v29
                    _ = v30
                }
            default:
                fmt.Println(v16)
            }
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v22 = int64(v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
