package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2}
    var v3 int = -1
    var v4 int = (v3 / (((v3 ^ (v3 + v3))) | 1))
    fmt.Println(v3, v4)
    switch ("hello" + "foo bar") {
    case "":
        if true {
            switch "" {
            case "foo bar":
                v3 += (v4 | v4)
            case "hello":
                fmt.Println(v4)
            case "a":
                v3 -= 127
            }
            var v5 int64 = -9223372036854775808
            v6 := GenS1{F0: 0}
            _ = v5
            _ = v6
        } else {
            var v7 int = 127
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS2{F0: -439, F1: 65535, F2: 65535}
                _ = len(v2)
                var v10 bool = false
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: 996}
            _ = v11
            _ = v7
        }
    }
    var v12 uint16 = 65535
    for fi13 := 0; fi13 < 6; fi13++ {
        switch "a" {
        case "hello":
            v12 += (v12 * 472)
        }
        var v14 int64 = 0
        switch v3 {
        case -2147483648:
            v15 := GenS3{F0: 453}
            _ = v15
        case 0:
            var v16 uint16 = (v12 + v12)
            _ = v16
        case 42:
            v14 -= -278
        }
        _ = fi13
        _ = v14
    }
    var v17 float32 = -2.25
    v18 := GenS1{F0: 353}
    _ = len(v2)
    var v19 uint = 1
    v1[0] = 1
    var v20 int = v3
    v21 := GenS1{F0: 902}
    for fi22 := 0; fi22 < 11; fi22++ {
        delete(v1, 0)
        _ = fi22
    }
    var v23 int16 = 444
    fmt.Println(v17)
    var v24 uint8 = 128
    if ((1e10 / ((((float32(0.0) - v17) + 1e10) * float32(float64((v17 + (v17 + v17))))) * v17)) == v17) {
        switch "" {
        case "foo bar":
            _ = len(v2)
        case "a":
            var v25 int64 = int64(v12)
            _ = v25
        case "\n\t":
            for fi26 := 0; fi26 < 10; fi26++ {
                _ = len(v1)
                v27 := GenS3{F0: (v19 << (((v19 >> (((v19 - 170)) & 63))) & 63))}
                var v28 int16 = (v23 % (((-(v23))) | 1))
                _ = fi26
                _ = v27
                _ = v28
            }
        }
        var v29 uint8 = (v24 | (v24 ^ (v24 / ((30) | 1))))
        v30 := GenS4{F0: (v12 ^ v12), F1: 2, F2: 776}
        _ = v29
        _ = v30
    } else {
        fmt.Println(v20, v23, v17)
        for fi31 := 0; fi31 < 8; fi31++ {
            fmt.Println(v12, v24, v23)
            _ = fi31
        }
        switch ("hello" + ((("" + "foo bar") + "\n\t") + "\n\t")) {
        case "\n\t":
            if true {
                var v32 float64 = 3.14159
                _ = v32
            } else {
                switch "foo bar" {
                case "hello":
                    var v33 float32 = 1e10
                    _ = v33
                case "":
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                case "foo bar":
                    var v34 string = "foo bar"
                    _ = v34
                }
                v3 = v3
            }
        case "a":
            if false {
                var v35 uint8 = v24
                _ = v35
            }
        case "foo bar":
            var v36 int8 = -34
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        fmt.Println(v12, v17)
        for fi38 := 0; fi38 < 15; fi38++ {
            v12 -= v12
            _ = fi38
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
