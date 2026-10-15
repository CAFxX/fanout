package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int16 }

func main() {
    var v1 float64 = (3.14159 * (0.1 + float64(3.14159)))
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    var v4 float32 = -2.25
    for fi5 := 0; fi5 < 11; fi5++ {
        if true {
            var v6 int = (fi5 & fi5)
            if false {
                var v7 int8 = 42
                _ = v7
            }
            _ = v6
        } else {
            v8 := GenS1{F0: false, F1: 649, F2: 255}
            for fi9 := 0; fi9 < 16; fi9++ {
                fmt.Println(fi5, v1, fi5)
                v10 := GenS1{F0: ("\n\t" > "foo bar"), F1: 238, F2: 224}
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: ((0 < -1) && false), F1: 177, F2: 128}
            _ = v11
            _ = v8
        }
        _ = fi5
    }
    v12 := GenS2{F0: 1}
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 string = "a"
        _ = fi13
        _ = v14
    }
    v1 -= ((0.1 + v1) * float64(v4))
    var v15 uint64 = 1
    for fi16 := 0; fi16 < 2; fi16++ {
        _ = len(v2)
        _ = fi16
    }
    v17 := []int{4294967295}
    fmt.Println(v1, v4, v4)
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    v3[32768] = "\n\t"
    v15 -= (v15 << ((1) & 63))
    for fi18 := 0; fi18 < 3; fi18++ {
        if (!true) {
            if (!false) {
                v15 = v15
                var v19 string = "a"
                _ = v19
            } else {
                v4 -= (v4 / v4)
                switch "foo bar" {
                case "":
                    for fi20 := 0; fi20 < 14; fi20++ {
                        var v21 uint16 = 2
                        _ = fi20
                        _ = v21
                    }
                case "\n\t":
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v4 = ((v4 * ((v4 * float32(1.5)) * v4)) / v4)
                        var v23 int16 = 32767
                        _ = fi22
                        _ = v23
                    }
                default:
                    fmt.Println(v15, v4)
                }
                v24 := v3[fi18]
                _ = v24
            }
            v1 = v1
        }
        _ = fi18
    }
    v25 := GenS3{F0: 841, F1: 1}
    for fi26 := 0; fi26 < 7; fi26++ {
        v1 += v1
        if (false && true) {
            switch ((((("foo bar" + (("hello" + ("foo bar" + "a")) + "")) + ("hello" + "a")) + "a") + "a") + ("" + "")) {
            case "hello":
                switch "hello" {
                case "hello":
                    fmt.Println(v4)
                case "a":
                    var v27 bool = false
                    _ = v27
                default:
                    v1 = v1
                }
            case "\n\t":
                if (359 >= 0) {
                    fmt.Println(v4, fi26, fi26)
                } else {
                    v15 += v15
                    var v28 float64 = (((v1 - (0.0 * float64(float32(v1)))) + (v1 / -2.25)) - v1)
                    var v29 string = ("foo bar" + ("" + "foo bar"))
                    _ = v28
                    _ = v29
                }
            }
        }
        var v30 int8 = 90
        _ = fi26
        _ = v30
    }
    if false {
        var v31 int = -871
        v32 := GenS4{F0: 692}
        switch "hello" {
        case "hello":
            var v33 int8 = -128
            _ = v33
        case "a":
            v1 -= v1
        default:
            v4 -= (v4 + (v4 - v4))
        }
        _ = v31
        _ = v32
    } else {
        fmt.Println(v15, v15, v15)
        var v34 int16 = 311
        delete(v3, 9223372036854775807)
        _ = v34
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
