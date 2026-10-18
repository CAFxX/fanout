package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }

func main() {
    var v1 uint64 = 202
    var v2 bool = true
    var v3 string = ("hello" + "")
    v4 := make(map[string]uint32)
    var v5 int8 = -128
    for fi6 := 0; fi6 < 1; fi6++ {
        if (!v2) {
            v5 = (v5 | v5)
            fmt.Println(v5, v2)
        } else {
            v7 := GenS1{F0: 32767, F1: 968}
            if v2 {
                v1 += v1
                var v8 float32 = -2.25
                v1 = v1
                _ = v8
            }
            _ = v7
        }
        v9 := GenS2{F0: 1, F1: 127}
        for fi10 := 0; fi10 < 5; fi10++ {
            if v2 {
                var v11 uint8 = uint8(v1)
                switch v3 {
                case "hello":
                    var v12 int = int(v1)
                    _ = v12
                default:
                    fmt.Println(v11, v11, v2)
                }
                v1 = v1
                _ = v11
            } else {
                v13 := GenS2{F0: 2147483648, F1: (-(v5))}
                fmt.Println(v5)
                _ = v13
            }
            _ = fi10
        }
        _ = fi6
        _ = v9
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v2 = (716 < 0)
        var v15 int64 = 0
        var v16 uint32 = 4294967295
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v2 = ((32767 < uint16(v1)) || v2)
        _ = fi17
    }
    v18 := []int{-1}
    if false {
        v19 := v4["\n\t"]
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        if false {
            switch "hello" {
            case "\n\t":
                switch fi20 {
                case -458:
                    fmt.Println(v5)
                case -9223372036854775808:
                    fmt.Println(v5, v3, v1)
                case 9223372036854775807:
                    if (!(0 <= 255)) {
                        v3 = v3
                    }
                default:
                    v21 := GenS1{F0: 874, F1: v1}
                    _ = v21
                }
            case "a":
                var v22 uint16 = 255
                _ = v22
            }
            v3 = (v3 + "")
            var v23 int32 = -1
            _ = v23
        } else {
            v24 := GenS1{F0: 42, F1: v1}
            _ = v24
        }
        _ = fi20
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 20; fi26++ {
            var v27 int32 = 2147483647
            _ = fi26
            _ = v27
        }
        v28 := GenS3{F0: 32767, F1: -1, F2: 775}
        _ = fi25
        _ = v28
    }
    if true {
        var v29 int64 = -432
        _ = v29
    }
    v18 = append(v18, -1)
    var v30 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
