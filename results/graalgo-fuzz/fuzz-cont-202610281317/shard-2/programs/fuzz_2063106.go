package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 string; F2 int8 }
func genFn0() int8 {
    if true {
        var v4 float32 = (float32(((0.0 - 0.1) - float64(1e10))) - (float32(1e10) / -2.25))
        for fi5 := 0; fi5 < 4; fi5++ {
            for fi6 := 0; fi6 < 1; fi6++ {
                v4 -= float32(3.14159)
                var v7 int64 = 1
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
        if (42 <= 0) {
            fmt.Println(v4)
            for fi8 := 0; fi8 < 2; fi8++ {
                fmt.Println(fi8, fi8)
                _ = fi8
            }
            fmt.Println(v4)
        }
        _ = v4
    } else {
        switch "\n\t" {
        case "foo bar":
            var v9 uint64 = 0
            _ = v9
        }
        v10 := GenS1{F0: 2147483647}
        switch 9223372036854775807 {
        case -1:
            var v11 int16 = -32768
            _ = v11
        case 128:
            var v12 int8 = 127
            _ = v12
        case -185:
            v13 := GenS1{F0: -453}
            _ = v13
        default:
            var v14 uint = 569
            _ = v14
        }
        _ = v10
    }
    var v15 uint64 = 1
    v16 := GenS2{F0: 33, F1: "\n\t", F2: 0}
    v15 = (v15 ^ 1)
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: 949}
            if true {
                v15 = v15
            }
            _ = fi18
            _ = v19
        }
        if (true && (!(2147483648 == 364))) {
            if (true && false) {
                v15 = v15
                v15 += v15
                fmt.Println(v15)
            }
            if (1.5 <= 0.1) {
                v20 := []int{fi17}
                if len(v20) > 0 {
                    v20[0] = fi17
                }
                _ = v20
            } else {
                v21 := GenS2{F0: 32767, F1: ("a" + "foo bar"), F2: 42}
                v22 := []int{(-(fi17)), fi17, (^(fi17))}
                if false {
                    var v23 int32 = -2147483648
                    if (!true) {
                        fmt.Println(v15, v15)
                        var v24 int64 = -77
                        for fi25 := 0; fi25 < 17; fi25++ {
                            fmt.Println(v23)
                            _ = fi25
                        }
                        _ = v24
                    } else {
                        var v26 int8 = 100
                        var v27 bool = false
                        _ = v26
                        _ = v27
                    }
                    v15 = (v15 + (v15 * (v15 + (v15 ^ v15))))
                    _ = v23
                }
                _ = v21
                _ = v22
            }
        }
        _ = fi17
    }
    _ = v15
    _ = v16
    return 42
}
func genFn1(p0 bool, p1 bool, p2 uint8) uint64 {
    v28 := []int{-2147483648}
    switch -116 {
    case 9223372036854775807:
        var v29 int64 = -1
        _ = v29
    default:
        var v30 uint32 = 2147483647
        _ = v30
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 float32 = ((3.14159 * -2.25) / ((0.0 * (0.0 + (0.1 - 1e10))) + 0.1))
        _ = fi31
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return 0
}

func main() {
    var v1 uint = 110
    v2 := []int{1, -966, 672, 127}
    v3 := make(map[uint32]string)
    for fi33 := 0; fi33 < 6; fi33++ {
        fmt.Println(fi33, fi33)
        for fi34 := 0; fi34 < 10; fi34++ {
            v35 := GenS1{F0: 42}
            if ((3.14159 >= 1.5) && true) {
                var v36 int64 = -927
                delete(v3, 2147483647)
                _ = v36
            }
            v1 -= 255
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    fmt.Println(v1)
    var v37 uint8 = 42
    fmt.Println(v37)
    v37 += (v37 / ((2) | 1))
    var v38 uint64 = 907
    v39 := v3[2147483647]
    var v40 int = 127
    for fi41 := 0; fi41 < 7; fi41++ {
        var v42 float32 = 0.0
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 11; fi43++ {
        v3[54] = "\n\t"
        var v44 uint8 = (^(v37))
        v1 -= 4294967295
        _ = fi43
        _ = v44
    }
    switch (v39 + "foo bar") {
    case "a":
        var v45 int8 = 42
        _ = v45
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = (v40 / ((v40) | 1))
        }
    case "hello":
        var v46 uint8 = (v37 << ((v1) & 63))
        _ = v46
    }
    v47 := v3[502]
    _ = len(v3)
    fmt.Println(v47, v39)
    v37 += (v37 & 127)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v47
}
