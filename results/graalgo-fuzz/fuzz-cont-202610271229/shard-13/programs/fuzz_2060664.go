package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint32 }
func genFn0(p0 int32, p1 int8, p2 int32) int {
    var v4 float64 = (3.14159 - 0.0)
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 float64 = (v4 * v4)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 bool = (false && false)
        switch (^(fi7)) {
        case 9223372036854775807:
            var v9 uint32 = 483
            _ = v9
        case 234:
            switch fi7 {
            case 438:
                v10 := GenS1{F0: 0, F1: 782}
                _ = v10
            default:
                v11 := []int{}
                _ = v11
            }
        default:
            if (((^(fi7)) == (fi7 | fi7)) && v8) {
                var v12 int = -313
                _ = v12
            }
        }
        _ = fi7
        _ = v8
    }
    v4 = ((float64(-2.25) - (v4 / v4)) * ((v4 + v4) - (float64(3.14159) - (v4 / v4))))
    v13 := []int{1, 9223372036854775807, 1, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v4
    return -891
}

func main() {
    var v1 float64 = (0.1 + -2.25)
    var v2 uint = 18446744073709551615
    var v3 int16 = 0
    for fi14 := 0; fi14 < 17; fi14++ {
        if false {
            switch fi14 {
            case -2147483648:
                v15 := []int{}
                _ = v15
            case 32767:
                v1 = ((v1 - (float64(-2.25) + v1)) * 0.1)
            case -330:
                v16 := GenS2{F0: 104, F1: 20}
                _ = v16
            default:
                var v17 int64 = 0
                _ = v17
            }
            v18 := []int{fi14, fi14}
            for fi19 := 0; fi19 < 3; fi19++ {
                if true {
                    var v20 uint32 = 4294967295
                    _ = v20
                }
                if (755 < 547) {
                    var v21 int8 = -91
                    v1 -= float64(((3.14159 + 1.5) + float32(float64(0.0))))
                    v22 := GenS3{F0: 725, F1: true, F2: 4294967295}
                    _ = v21
                    _ = v22
                }
                var v23 uint32 = 0
                _ = fi19
                _ = v23
            }
            _ = v18
        }
        _ = fi14
    }
    switch "a" {
    case "":
        for fi24 := 0; fi24 < 8; fi24++ {
            switch "a" {
            case "hello":
                v2 = v2
            case "":
                var v25 uint = 130
                _ = v25
            case "a":
                fmt.Println(v3)
            }
            for fi26 := 0; fi26 < 15; fi26++ {
                v27 := GenS1{F0: 409, F1: (^(v3))}
                _ = fi26
                _ = v27
            }
            var v28 int = fi24
            _ = fi24
            _ = v28
        }
    case "foo bar":
        v29 := GenS3{F0: 0, F1: false, F2: 1}
        _ = v29
    }
    v30 := []int{}
    var v31 float32 = ((3.14159 - (1.5 + 3.14159)) / (-2.25 + float32(v1)))
    fmt.Println(v3, v2)
    if true {
        var v32 int8 = -82
        _ = v32
    } else {
        v31 = 1e10
        v3 -= v3
        for fi33 := 0; fi33 < 4; fi33++ {
            v3 = (v3 ^ (v3 << ((275) & 63)))
            _ = fi33
        }
    }
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    v2 = 65535
    for fi34 := 0; fi34 < 18; fi34++ {
        for fi35 := 0; fi35 < 8; fi35++ {
            v31 = v31
            v31 -= (v31 * ((0.0 / 1.0) * v31))
            _ = fi35
        }
        v31 = v31
        v3 -= (v3 >> (((v2 - 4294967295)) & 63))
        _ = fi34
    }
    var v36 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
}
