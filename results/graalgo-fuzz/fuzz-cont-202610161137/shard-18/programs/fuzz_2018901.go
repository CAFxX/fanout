package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 bool; F2 uint8 }

func main() {
    var v1 uint = 42
    var v2 uint16 = 2
    var v3 uint = (^(v1))
    var v4 uint64 = 4294967295
    v5 := make(map[string]string)
    var v6 int16 = 1
    if false {
        for fi7 := 0; fi7 < 16; fi7++ {
            switch 935 {
            case 9223372036854775807:
                v8 := GenS1{F0: uint(v4)}
                _ = v8
            }
            _ = fi7
        }
        switch 42 {
        case -1:
            v4 = v4
        case 441:
            fmt.Println(v6, v6, v1)
        }
    }
    v1 -= (v1 * (v1 * v1))
    var v9 int16 = (v6 << ((4294967295) & 63))
    for fi10 := 0; fi10 < 17; fi10++ {
        if (!false) {
            switch "hello" {
            case "":
                fmt.Println(v6)
            }
        } else {
            v11 := GenS2{F0: (v3 & v3), F1: (!false), F2: 255}
            _ = v11
        }
        _ = fi10
    }
    fmt.Println(v3)
    switch "hello" {
    case "hello":
        v12 := []int{}
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := []int{fi13, fi13, 1, (fi13 + (fi13 << (((v3 & v1)) & 63)))}
            var v15 int8 = -1
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v16 := GenS1{F0: (v3 % (((v1 * v1)) | 1))}
    for fi17 := 0; fi17 < 13; fi17++ {
        v6 = v6
        var v18 float64 = -2.25
        var v19 uint16 = 42
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int16 = v6
    var v21 bool = (v4 > 18446744073709551615)
    fmt.Println(v1)
    var v22 int8 = 26
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 float64 = 3.14159
        var v25 uint32 = uint32(v4)
        for fi26 := 0; fi26 < 18; fi26++ {
            v20 -= v6
            _ = fi26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    v27 := GenS1{F0: v1}
    for fi28 := 0; fi28 < 20; fi28++ {
        if v21 {
            v29 := GenS2{F0: v1, F1: v21, F2: 0}
            _ = v29
        } else {
            switch fi28 {
            case 0:
                fmt.Println(v4, v2)
            case -399:
                var v30 int8 = -13
                _ = v30
            default:
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 bool = v21
                    if v32 {
                        fmt.Println(v9, v21)
                        var v33 int64 = int64(v3)
                        var v34 int8 = int8(v20)
                        _ = v33
                        _ = v34
                    }
                    fmt.Println(v21, fi28)
                    _ = fi31
                    _ = v32
                }
            }
            if (1 >= 128) {
                v4 = v4
                if (1 < 712) {
                    v2 = v2
                } else {
                    var v35 int64 = int64(v4)
                    if ((!true) || (((^(v2)) > v2) || false)) {
                        v9 = v6
                        fmt.Println(v9, v6)
                    }
                    var v36 float64 = (-2.25 + 0.0)
                    _ = v35
                    _ = v36
                }
                var v37 int16 = -491
                _ = v37
            }
            var v38 uint16 = v2
            _ = v38
        }
        v6 = 947
        _ = fi28
    }
    var v39 int8 = v22
    var v40 int32 = 688
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
