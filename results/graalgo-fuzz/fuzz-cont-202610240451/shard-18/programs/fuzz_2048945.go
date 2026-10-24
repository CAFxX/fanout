package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
func genFn0(p0 int8) uint32 {
    p0 = (p0 - (p0 & (p0 | p0)))
    p0 = (p0 >> ((2) & 63))
    if false {
        v5 := []int{}
        if (4294967295 < 370) {
            fmt.Println(p0)
        } else {
            var v6 int16 = 32767
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 float32 = (1e10 + (1.5 - 3.14159))
            var v9 int = fi7
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v5
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 uint = 4294967295
            p0 -= (p0 % (((p0 - p0)) | 1))
            switch int(p0) {
            case -1:
                v12 := []int{(-(fi10)), (-(fi10)), 2147483647, fi10}
                _ = v12
            default:
                fmt.Println(v11)
            }
            _ = fi10
            _ = v11
        }
        var v13 bool = (true || false)
        _ = v13
    }
    _ = p0
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 1000
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    v2 = 1
    v14 := GenS1{F0: -2147483648, F1: 0, F2: 2147483647}
    v15 := []int{2147483647, 32768}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    switch "\n\t" {
    case "a":
        switch "a" {
        case "":
            var v16 bool = ((-(v2)) > v2)
            _ = v16
        case "a":
            if false {
                var v17 string = "a"
                _ = v17
            } else {
                v2 += -627
                v2 += v2
            }
        default:
            fmt.Println(v2)
        }
    case "":
        var v18 float64 = 0.0
        _ = v18
    }
    if false {
        v2 = v2
        v2 = v2
        var v19 uint8 = 127
        _ = v19
    } else {
        var v20 uint16 = 2
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 uint32 = 2147483648
            v24 := GenS1{F0: fi21, F1: v2, F2: 0}
            _ = fi22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            if (((-2.25 - float32(float64(1.5))) > (0.1 / 3.14159)) && true) {
                switch ("hello" + "a") {
                case "\n\t":
                    v26 := v4[2147483648]
                    _ = v26
                case "a":
                    delete(v4, 4294967295)
                default:
                    fmt.Println(fi25)
                }
                v2 = v2
            }
            _ = fi25
        }
        switch fi21 {
        case -517:
            switch 0 {
            case 9223372036854775807:
                var v27 int64 = 1
                _ = v27
            default:
                _ = len(v3)
            }
        case -486:
            if false {
                var v28 uint16 = 871
                v28 = (v28 - 158)
                _ = v28
            } else {
                v2 = v2
            }
        default:
            v29 := GenS2{F0: 65535}
            _ = v29
        }
        _ = fi21
    }
    v2 += -63
    if len(v15) > 0 {
        v15[0] = 4294967295
    }
    v2 -= 42
    fmt.Println(v2)
    v2 += (^(v2))
    v30 := GenS2{F0: 2}
    var v31 uint8 = 0
    if false {
        _ = len(v15)
        fmt.Println(v2, v2)
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        fmt.Println(v31, v2, fi32)
        _ = fi32
    }
    fmt.Println(v31, v31)
    v33 := GenS3{F0: 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
}
