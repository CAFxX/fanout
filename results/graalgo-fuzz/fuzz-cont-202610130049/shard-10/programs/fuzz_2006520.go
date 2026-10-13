package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int; F1 string; F2 int64 }
type GenS5 struct { F0 int32; F1 int16; F2 int }
type GenS6 struct { F0 int8 }

func main() {
    var v1 uint64 = 0
    var v2 float32 = -2.25
    v3 := []string{"hello"}
    var v4 uint64 = 18446744073709551615
    var v5 int64 = 504
    v6 := make(map[uint32]string)
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 20; fi8++ {
        switch fi8 {
        case -922:
            v7 = (^(v1))
        case -2147483648:
            v4 = 2
        }
        v9 := GenS1{F0: (-63 <= 42)}
        var v10 uint64 = (v7 ^ v1)
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 2147483647, F1: (^(v1))}
    v12 := GenS1{F0: true}
    var v13 float64 = 1e10
    v14 := GenS2{F0: 2147483648, F1: (v7 - v7)}
    v3 = append(v3, "a")
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            switch ("" + "foo bar") {
            case "foo bar":
                var v17 bool = true
                _ = v17
            }
            _ = fi16
        }
        if (true && false) {
            v18 := GenS3{F0: fi15, F1: 1}
            _ = v18
        } else {
            fmt.Println(v1)
        }
        _ = fi15
    }
    v19 := GenS4{F0: -848, F1: "hello", F2: v5}
    var v20 float64 = float64(v2)
    var v21 uint32 = 0
    var v22 int8 = 97
    var v23 int16 = int16(v1)
    var v24 int32 = -1
    _ = len(v6)
    var v25 uint16 = 1
    for fi26 := 0; fi26 < 6; fi26++ {
        switch "a" {
        case "\n\t":
            if true {
                var v27 float32 = (float32((v13 - v13)) / (float32(v13) / 0.1))
                _ = v27
            }
        default:
            var v28 uint32 = (v21 >> ((255) & 63))
            _ = v28
        }
        var v29 string = ""
        switch ((v29 + (((v29 + v29) + (v29 + "")) + v29)) + v29) {
        case "":
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if false {
                    fmt.Println(v21, v13)
                    v30 := GenS5{F0: 1, F1: 1000, F2: fi26}
                    v20 = v20
                    _ = v30
                }
            } else {
                for fi31 := 0; fi31 < 1; fi31++ {
                    var v32 float64 = (v13 - v13)
                    v24 -= v24
                    fmt.Println(v32)
                    _ = fi31
                    _ = v32
                }
                v4 += uint64(v21)
                v33 := GenS6{F0: 100}
                _ = v33
            }
        case "\n\t":
            var v34 uint32 = 2147483648
            _ = v34
        }
        _ = fi26
        _ = v29
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 int8 = v22
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
