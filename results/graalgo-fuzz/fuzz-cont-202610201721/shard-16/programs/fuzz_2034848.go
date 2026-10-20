package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 bool }
func genFn0() int64 {
    var v6 int64 = 1
    v7 := []int{2147483647, 32767, -279, 1}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int = (fi8 * 640)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v6 += (^(v6))
        _ = fi10
    }
    fmt.Println(v6)
    _ = v6
    _ = v7
    return v6
}
func genFn1() string {
    if false {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int64 = -853
            _ = fi11
            _ = v12
        }
    } else {
        var v13 uint = 255
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS1{F0: 0, F1: genFn0()}
            v13 -= v13
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    var v16 int32 = -757
    var v17 string = ("" + "")
    _ = v16
    _ = v17
    return (v17 + v17)
}

func main() {
    var v1 int16 = 541
    var v2 int64 = -1
    var v3 uint8 = 42
    v4 := make(map[int]int)
    var v5 bool = (!(65535 == 4294967295))
    var v18 int64 = -9223372036854775808
    if v5 {
        v1 += v1
        var v19 int64 = (v18 ^ genFn0())
        _ = v19
    }
    if true {
        fmt.Println(v3, v3)
        fmt.Println(v18, v1, v18)
        v5 = v5
    }
    v20 := []int{}
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 int32 = 2147483647
        switch fi21 {
        case -240:
            if (v3 != 0) {
                for fi23 := 0; fi23 < 15; fi23++ {
                    fmt.Println(v3, fi23)
                    _ = fi23
                }
                v24 := GenS2{F0: 279, F1: 255}
                _ = v24
            }
        case 32767:
            for fi25 := 0; fi25 < 14; fi25++ {
                v20 = append(v20, (-(fi25)))
                if v5 {
                    _ = len(v20)
                    switch int(v22) {
                    case 0:
                        v5 = v5
                    default:
                        v2 += v18
                    }
                }
                _ = fi25
            }
        default:
            switch -1 {
            case -2147483648:
                v26 := v4[(^(fi21))]
                _ = v26
            case 1:
                var v27 uint16 = 657
                _ = v27
            }
        }
        _ = fi21
        _ = v22
    }
    if len(v20) > 0 {
        v20[0] = 128
    }
    if (v5 || v5) {
        var v28 bool = ((v1 - (v1 + 1000)) >= v1)
        var v29 int8 = 100
        v20 = append(v20, 32767)
        _ = v28
        _ = v29
    }
    if (255 > 65535) {
        v2 = genFn0()
    }
    switch "foo bar" {
    case "":
        v5 = (!false)
    case "\n\t":
        v30 := GenS3{F0: (v1 | 1)}
        _ = v30
    }
    if v5 {
        var v31 uint32 = 42
        var v32 uint32 = v31
        _ = v31
        _ = v32
    }
    var v33 int32 = int32(v3)
    if (65535 > 2147483648) {
        fmt.Println(v33, v18, v5)
        for fi34 := 0; fi34 < 14; fi34++ {
            v1 += (v1 & v1)
            _ = fi34
        }
    } else {
        v33 = (v33 ^ v33)
        var v35 uint32 = 2147483647
        var v36 uint8 = (v3 << ((2) & 63))
        _ = v35
        _ = v36
    }
    var v37 uint16 = 42
    if ((int64(v3) == (v18 + (-(v18)))) && (1 < 1)) {
        v37 = (^(v37))
    } else {
        v5 = ((v3 ^ v3) > v3)
        v38 := GenS4{F0: "foo bar", F1: v5}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
}
