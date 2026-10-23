package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint32) string {
    p0 = 1
    var v4 float32 = 1.5
    _ = p0
    _ = v4
    return ""
}
func genFn1(p0 int, p1 uint16) int8 {
    var v5 uint32 = 1
    v6 := GenS1{F0: false}
    v7 := []int{0, p0}
    var v8 float64 = 0.1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int8 = 1
    var v2 uint8 = 42
    var v3 int = 128
    var v9 float64 = 1.5
    for fi10 := 0; fi10 < 10; fi10++ {
        switch "\n\t" {
        case "foo bar":
            v11 := []int{(v3 / (((v3 + (v3 ^ v3))) | 1)), fi10}
            _ = v11
        case "\n\t":
            var v12 uint32 = 2147483647
            _ = v12
        }
        _ = fi10
    }
    v13 := []int{32767}
    switch (v3 % ((-1) | 1)) {
    case 9223372036854775807:
        v14 := GenS2{F0: (v1 | v1)}
        _ = v14
    case 42:
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    default:
        var v15 int32 = -570
        _ = v15
    }
    switch "" {
    case "foo bar":
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    fmt.Println(v1, v1)
    v16 := GenS1{F0: true}
    for fi17 := 0; fi17 < 13; fi17++ {
        v1 = (-(v1))
        _ = fi17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v1 = (v1 >> ((42) & 63))
        var v19 int64 = 579
        var v20 uint = 42
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint64 = 1
    v21 = v21
    for fi22 := 0; fi22 < 17; fi22++ {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        switch (fi22 - 4294967295) {
        case -377:
            var v23 uint8 = v2
            _ = v23
        case 9223372036854775807:
            var v24 bool = (((v21 != 0) && (false && (false || (!(!(true && (false && true))))))) || (1000 == -1))
            _ = v24
        default:
            for fi25 := 0; fi25 < 12; fi25++ {
                var v26 uint = 65535
                _ = fi25
                _ = v26
            }
        }
        _ = fi22
    }
    var v27 uint32 = 721
    v9 = float64(-2.25)
    for fi28 := 0; fi28 < 15; fi28++ {
        switch "foo bar" {
        case "hello":
            switch (v3 % ((9223372036854775807) | 1)) {
            case -918:
                fmt.Println(v3, v3, v3)
            case -653:
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            case 32768:
                var v29 int16 = 32767
                _ = v29
            default:
                if false {
                    v30 := GenS3{F0: v1, F1: (v27 << ((255) & 63))}
                    var v31 float64 = 3.14159
                    _ = v30
                    _ = v31
                } else {
                    var v32 int64 = 42
                    _ = len(v13)
                    _ = v32
                }
            }
        default:
            v3 = (^(v3))
        }
        switch ("foo bar" + "") {
        case "":
            if false {
                if false {
                    switch (fi28 * 128) {
                    case 1:
                        v33 := GenS4{F0: ("\n\t" + "a")}
                        _ = v33
                    default:
                        v34 := GenS1{F0: (true && true)}
                        _ = v34
                    }
                }
                v13 = append(v13, (fi28 / (((fi28 * v3)) | 1)))
                v3 += v3
            } else {
                v35 := GenS5{F0: 32767, F1: 9223372036854775807}
                v21 -= (v21 * uint64(v2))
                v9 = 1e10
                _ = v35
            }
        case "hello":
            var v36 float32 = float32((-2.25 - ((0.1 * 0.1) / (v9 + -2.25))))
            _ = v36
        case "\n\t":
            var v37 int32 = 845
            _ = v37
        default:
            v2 = (^(v2))
        }
        _ = len(v13)
        _ = fi28
    }
    v9 = (v9 * float64(1.5))
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v9
}
