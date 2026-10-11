package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint16 }
func genFn0(p0 uint8, p1 int, p2 uint32) int {
    var v5 string = (((("a" + "a") + ("\n\t" + "")) + "a") + "hello")
    var v6 int64 = 1
    v7 := []int{}
    v6 -= -9223372036854775808
    if len(v7) > 0 {
        v7[0] = p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (p1 * -2147483648)
}

func main() {
    var v1 int64 = 33
    var v2 int64 = v1
    var v3 uint64 = 0
    var v4 int16 = -32768
    v2 -= -99
    if true {
        v3 = v3
    }
    if false {
        v8 := GenS1{F0: 0}
        v3 += v3
        _ = v8
    } else {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int32 = 1
            var v11 uint64 = v3
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint64 = v3
        for fi13 := 0; fi13 < 14; fi13++ {
            fmt.Println(fi13, v2)
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 uint = 1
                var v16 uint = 4294967295
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 float64 = 1.5
            _ = fi13
            _ = v17
        }
        _ = v12
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := []int{(^(fi18))}
        v19 = append(v19, 877)
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 int32 = 2147483647
            v22 := GenS2{F0: -45, F1: (v2 & (v2 + 1)), F2: 2}
            if true {
                var v23 uint16 = 255
                v24 := GenS1{F0: 1}
                v25 := GenS1{F0: 255}
                _ = v23
                _ = v24
                _ = v25
            } else {
                if (false && true) {
                    v26 := GenS2{F0: 42, F1: -6, F2: 515}
                    var v27 uint32 = 2
                    v21 -= (v21 & (^(v21)))
                    _ = v26
                    _ = v27
                }
                var v28 float32 = (0.0 - float32(3.14159))
                v29 := GenS2{F0: -128, F1: 0, F2: 2}
                _ = v28
                _ = v29
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    var v30 uint64 = (v3 + (v3 | v3))
    if false {
        fmt.Println(v1)
        v31 := []int{221, 1, -2147483648, genFn0(0, 42, 847)}
        _ = v31
    }
    v32 := GenS1{F0: 2}
    for fi33 := 0; fi33 < 14; fi33++ {
        v2 += (v2 - v2)
        switch 42 {
        case 2147483647:
            if false {
                v4 += -354
                v34 := GenS2{F0: 127, F1: (v1 - (v1 % (((v1 - -1)) | 1))), F2: 32767}
                v35 := []int{(fi33 & fi33), genFn0(1, (^(fi33)), 4294967295), genFn0(127, fi33, 2147483648), fi33}
                _ = v34
                _ = v35
            }
        }
        _ = fi33
    }
    v36 := []int{127}
    v37 := GenS1{F0: 128}
    fmt.Println(v3, v3)
    for fi38 := 0; fi38 < 11; fi38++ {
        v30 += 2
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
        _ = fi38
    }
    var v39 int16 = -1
    for fi40 := 0; fi40 < 13; fi40++ {
        if (65535 <= 65535) {
            if true {
                if (true || ((float64(1.5) - float64(0.0)) != 0.0)) {
                    for fi41 := 0; fi41 < 13; fi41++ {
                        fmt.Println(v30)
                        _ = fi41
                    }
                    v3 += v3
                }
            }
        } else {
            v42 := GenS3{F0: 42, F1: 65535, F2: 1}
            var v43 int8 = 1
            _ = v42
            _ = v43
        }
        v4 -= (-(v39))
        switch "" {
        case "":
            v44 := GenS1{F0: 2}
            _ = v44
        case "\n\t":
            v45 := GenS4{F0: 0, F1: 42, F2: 32767}
            _ = v45
        default:
            if len(v36) > 0 {
                v36[0] = (fi40 | (fi40 - (fi40 - fi40)))
            }
        }
        _ = fi40
    }
    if true {
        var v46 uint = 65535
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
        var v47 float32 = 0.1
        _ = v46
        _ = v47
    }
    v48 := GenS1{F0: 1}
    v4 += (v4 ^ v39)
    switch (("\n\t" + "a") + "\n\t") {
    case "hello":
        fmt.Println(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v48
}
