package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
func genFn0() string {
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: 2}
            _ = fi6
            _ = v7
        }
        if (26 >= 0) {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{fi8, (fi8 / (((fi5 + fi5)) | 1)), fi8}
                if true {
                    v10 := GenS1{F0: 2}
                    if ((3.14159 > (float32(3.14159) + ((0.1 * float32(1.5)) * -2.25))) || (true && (3.14159 < (-2.25 + 1e10)))) {
                        var v11 int16 = -185
                        _ = v11
                    }
                    if len(v9) > 0 {
                        v9[0] = -146
                    }
                    v12 := GenS2{F0: 128, F1: 32767, F2: 0}
                    switch "\n\t" {
                    case "\n\t":
                        var v13 float64 = (3.14159 - 3.14159)
                        _ = v13
                    case "":
                        v14 := GenS3{F0: ("\n\t" == "a")}
                        _ = v14
                    case "foo bar":
                        var v15 uint8 = 2
                        _ = v15
                    }
                    _ = v10
                    _ = v12
                } else {
                    var v16 uint32 = 65535
                    _ = v16
                }
                _ = fi8
                _ = v9
            }
            if false {
                var v17 float64 = 0.1
                fmt.Println(fi5)
                v17 += v17
                _ = v17
            }
        }
        _ = fi5
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint = 814
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 65535}
    var v21 int8 = 1
    _ = v20
    _ = v21
    return ("" + "a")
}
func genFn1(p0 int) int64 {
    if ((!false) && (true || true)) {
        switch "hello" {
        case "a":
            p0 = (^(p0))
        default:
            var v22 int16 = 1
            _ = v22
        }
    }
    var v23 int32 = 429
    p0 = (p0 + 128)
    v24 := []int{p0, (p0 & 0)}
    _ = p0
    _ = v23
    _ = v24
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]int)
    var v25 int16 = 32767
    v25 -= v25
    v26 := []int{9223372036854775807, 389, -1, int(v25)}
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 uint8 = 1
        _ = fi27
        _ = v28
    }
    var v29 uint64 = 224
    switch "hello" {
    case "\n\t":
        for fi30 := 0; fi30 < 18; fi30++ {
            v25 = (^(v25))
            var v31 int = 127
            _ = fi30
            _ = v31
        }
    case "a":
        v32 := GenS4{F0: -1}
        _ = v32
    case "hello":
        _ = len(v26)
    default:
        v33 := GenS2{F0: 0, F1: 32767, F2: (-(v3))}
        _ = v33
    }
    var v34 uint = 255
    var v35 uint32 = 65535
    v4[(v35 ^ v35)] = -778
    if true {
        v3 = int64(v25)
        var v36 int = -377
        var v37 int = (^(v36))
        _ = v36
        _ = v37
    } else {
        var v38 int8 = 42
        for fi39 := 0; fi39 < 20; fi39++ {
            var v40 int16 = (v25 >> ((v34) & 63))
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            _ = fi39
            _ = v40
        }
        _ = v38
    }
    _ = len(v26)
    if len(v26) > 0 {
        v26[0] = -758
    }
    var v41 float32 = (0.0 / 1.5)
    v42 := GenS3{F0: true}
    switch "hello" {
    case "":
        if len(v26) > 0 {
            v26[0] = 0
        }
    case "foo bar":
        var v43 float32 = (0.0 * float32((1.5 * (3.14159 + 1e10))))
        _ = v43
    case "hello":
        var v44 uint16 = 65535
        _ = v44
    }
    v25 = v25
    v45 := GenS3{F0: (!(false || false))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v41
    _ = v42
    _ = v45
}
