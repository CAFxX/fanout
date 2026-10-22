package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint16; F1 string }
func genFn0() int64 {
    if (3.14159 > (0.0 / -2.25)) {
        var v6 int = 32767
        v7 := GenS1{F0: 153, F1: 467, F2: 1}
        v8 := GenS2{F0: 42, F1: 2147483648}
        _ = v6
        _ = v7
        _ = v8
    } else {
        v9 := GenS3{F0: 0}
        var v10 int32 = -1
        _ = v10
        _ = v9
    }
    var v11 uint = 592
    _ = v11
    return 42
}

func main() {
    var v1 uint16 = 1
    v2 := []uint32{}
    var v3 uint64 = 169
    var v4 int32 = -49
    var v5 uint8 = 255
    for fi12 := 0; fi12 < 20; fi12++ {
        v3 = (v3 & (v3 & v3))
        _ = fi12
    }
    if false {
        var v13 uint = 337
        v4 += 42
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v1 = v1
        switch "" {
        case "\n\t":
            if (!(!false)) {
                for fi15 := 0; fi15 < 20; fi15++ {
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    v16 := GenS3{F0: v1}
                    _ = fi15
                    _ = v16
                }
            } else {
                v3 -= v3
            }
        case "hello":
            v17 := GenS3{F0: 940}
            _ = v17
        }
        _ = fi14
    }
    v4 = (v4 / ((v4) | 1))
    fmt.Println(v5, v5, v4)
    switch 1 {
    case -416:
        var v18 uint = 799
        _ = v18
    case 128:
        if ((true && false) && false) {
            if true {
                v1 -= 0
                v3 += (v3 & (^(v3)))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = len(v2)
            v19 := GenS1{F0: 798, F1: v4, F2: 0}
            _ = v19
        }
    }
    fmt.Println(v4, v1)
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := GenS4{F0: v4, F1: (v3 ^ (^(v3)))}
        for fi22 := 0; fi22 < 1; fi22++ {
            if false {
                var v23 uint64 = (v3 >> ((610) & 63))
                fmt.Println(v23)
                v24 := GenS1{F0: 255, F1: 0, F2: (^(v5))}
                var v25 int32 = v4
                var v26 int64 = -9223372036854775808
                var v27 float32 = -2.25
                v28 := GenS5{F0: v1, F1: "hello"}
                _ = v23
                _ = v24
                _ = v25
                _ = v26
                _ = v27
                _ = v28
            } else {
                v29 := GenS1{F0: 65535, F1: -1, F2: v5}
                var v30 string = (("foo bar" + "a") + ("hello" + ""))
                var v31 uint32 = 815
                _ = v29
                _ = v30
                _ = v31
            }
            if (true || false) {
                v32 := GenS4{F0: (v4 % (((v4 << ((42) & 63))) | 1)), F1: (v3 << ((2) & 63))}
                _ = v32
            } else {
                v33 := GenS3{F0: (v1 << ((4294967295) & 63))}
                v4 = (v4 & (v4 >> ((4294967295) & 63)))
                _ = v33
            }
            _ = len(v2)
            _ = fi22
        }
        if ((2147483647 != (v4 * v4)) || false) {
            if false {
                var v34 int16 = -39
                v34 -= v34
                _ = v34
            }
            fmt.Println(fi20)
            fmt.Println(v5, v1, fi20)
        } else {
            if false {
                v3 = 151
                _ = len(v2)
            }
            fmt.Println(v1, v3)
        }
        _ = fi20
        _ = v21
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = (v1 % ((v1) | 1))
    v5 = v5
    for fi35 := 0; fi35 < 12; fi35++ {
        v2 = append(v2, 42)
        v3 = (v3 / ((v3) | 1))
        _ = fi35
    }
    var v36 int16 = -32768
    v1 = v1
    var v37 uint32 = uint32(v5)
    v4 -= -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
