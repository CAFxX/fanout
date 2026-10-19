package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    var v1 uint64 = 2
    v2 := []int{32767, 32768, -1, 708}
    var v3 int8 = 1
    v4 := make(map[string]string)
    v3 = (v3 / (((v3 & (^(v3)))) | 1))
    v5 := GenS1{F0: 0, F1: true, F2: ("\n\t" + "hello")}
    v1 += 42
    v3 = (v3 % (((v3 - 1)) | 1))
    var v6 float64 = 0.1
    var v7 int8 = v3
    var v8 uint16 = 1
    v3 += v3
    if false {
        v6 -= v6
        var v9 int16 = 1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v9
    } else {
        v3 = -127
        var v10 bool = true
        _ = v10
    }
    var v11 int8 = (v7 | v7)
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    v12 := GenS1{F0: 65535, F1: true, F2: ""}
    if false {
        var v13 string = "a"
        if false {
            delete(v4, v13)
            var v14 uint16 = v8
            _ = v14
        }
        v15 := GenS1{F0: (v8 & 32767), F1: false, F2: v13}
        _ = v13
        _ = v15
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            if true {
                var v17 uint64 = v1
                _ = v17
            } else {
                var v18 bool = (v7 > v3)
                _ = len(v4)
                switch fi16 {
                case -1:
                    var v19 uint8 = 42
                    _ = v19
                case 128:
                    if true {
                        var v20 uint8 = 1
                        fmt.Println(v18)
                        v21 := GenS1{F0: 65535, F1: ("\n\t" < "foo bar"), F2: "hello"}
                        _ = v20
                        _ = v21
                    } else {
                        var v22 int = fi16
                        _ = v22
                    }
                }
                _ = v18
            }
            v23 := GenS1{F0: v8, F1: false, F2: "\n\t"}
            _ = fi16
            _ = v23
        }
        v11 -= v7
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 bool = true
        var v26 int8 = v11
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: -677}
    var v28 uint = 2
    for fi29 := 0; fi29 < 19; fi29++ {
        switch "hello" {
        case "\n\t":
            for fi30 := 0; fi30 < 7; fi30++ {
                fmt.Println(v1, fi30)
                var v31 int = (fi30 << (((v28 << ((v28) & 63))) & 63))
                _ = fi30
                _ = v31
            }
        case "foo bar":
            var v32 int32 = 539
            _ = v32
        case "a":
            v33 := GenS3{F0: -308, F1: -954}
            _ = v33
        }
        _ = fi29
    }
    v8 -= v8
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
