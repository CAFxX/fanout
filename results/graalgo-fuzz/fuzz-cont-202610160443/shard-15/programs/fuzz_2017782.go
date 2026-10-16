package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }

func main() {
    var v1 int = 2147483647
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 4294967295
    var v4 float32 = -2.25
    var v5 uint = 4294967295
    v6 := []int{128, 127, v1, v1}
    var v7 bool = (true && true)
    v8 := GenS1{F0: v7, F1: 2, F2: -81}
    for fi9 := 0; fi9 < 14; fi9++ {
        switch "a" {
        case "a":
            v10 := GenS1{F0: v7, F1: 317, F2: 0}
            _ = v10
        }
        var v11 int8 = 100
        switch fi9 {
        case 32768:
            for fi12 := 0; fi12 < 4; fi12++ {
                if v7 {
                    var v13 int16 = -240
                    _ = v13
                } else {
                    v14 := GenS2{F0: 42, F1: 2, F2: -1}
                    _ = v14
                }
                if true {
                    v1 = v1
                } else {
                    var v15 int64 = -375
                    v11 -= (v11 / (((v11 + (v11 ^ (v11 >> ((v5) & 63))))) | 1))
                    _ = v15
                }
                var v16 int8 = (v11 * -64)
                fmt.Println(v3, v7, v16)
                _ = fi12
                _ = v16
            }
        case 9223372036854775807:
            var v17 int = (fi9 & 42)
            _ = v17
        }
        _ = fi9
        _ = v11
    }
    v18 := GenS3{F0: 65535}
    var v19 uint16 = 65535
    var v20 int8 = 1
    for fi21 := 0; fi21 < 5; fi21++ {
        v7 = v7
        v6 = append(v6, v1)
        _ = fi21
    }
    if false {
        switch "" {
        case "":
            for fi22 := 0; fi22 < 4; fi22++ {
                for fi23 := 0; fi23 < 1; fi23++ {
                    if v7 {
                        v24 := GenS3{F0: v19}
                        _ = v24
                    }
                    _ = fi23
                }
                v1 = fi22
                switch "a" {
                case "a":
                    var v25 bool = ((!(18446744073709551615 == v3)) && (!v7))
                    _ = v25
                case "":
                    var v26 int16 = -130
                    _ = v26
                case "\n\t":
                    v19 = 65535
                }
                _ = fi22
            }
        }
        var v27 int32 = -2147483648
        _ = v27
    } else {
        v28 := GenS2{F0: 128, F1: v19, F2: 0}
        if (1 < -757) {
            var v29 float32 = v4
            if ((4294967295 >= v5) && (uint64(v5) > (v3 * v3))) {
                var v30 int32 = -1
                _ = v30
            } else {
                v31 := GenS4{F0: ""}
                _ = v31
            }
            _ = v29
        }
        if (!true) {
            v3 = (v3 | (v3 - (v3 >> ((v5) & 63))))
            var v32 uint16 = (v19 * (v19 % ((v19) | 1)))
            _ = v32
        } else {
            fmt.Println(v1, v4, v1)
            if false {
                var v33 bool = v7
                _ = v33
            }
            if v7 {
                v19 -= uint16(v5)
            }
        }
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
