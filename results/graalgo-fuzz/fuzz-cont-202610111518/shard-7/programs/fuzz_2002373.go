package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int8; F2 int64 }
func genFn0(p0 int64, p1 int32, p2 int64) int16 {
    switch "" {
    case "a":
        if (42 == 0) {
            fmt.Println(p2, p0, p1)
            switch 0 {
            case 4294967295:
                var v4 int64 = (-(p0))
                _ = v4
            }
        } else {
            for fi5 := 0; fi5 < 9; fi5++ {
                v6 := GenS1{F0: (^(fi5))}
                v7 := GenS2{F0: (fi5 / ((fi5) | 1)), F1: 32767, F2: 32767}
                _ = fi5
                _ = v6
                _ = v7
            }
            var v8 int64 = (p2 / ((p0) | 1))
            _ = v8
        }
    case "":
        v9 := []int{2147483647, -1}
        _ = v9
    default:
        p0 -= p0
    }
    p1 = (p1 ^ p1)
    v10 := []int{32767}
    var v11 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 42
}
func genFn1(p0 bool, p1 int8, p2 int32) int64 {
    p0 = true
    p2 = p2
    var v12 float64 = (1.5 / ((float64(1.5) / 1e10) - float64(3.14159)))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -1
}

func main() {
    var v1 int64 = 0
    var v2 int64 = (v1 | (v1 + v1))
    v3 := []uint32{65535, 65535, 556, 500}
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: 0}
        _ = fi13
        _ = v14
    }
    var v15 int32 = 364
    v15 -= v15
    var v16 int8 = 100
    var v17 uint = 4294967295
    if (!false) {
        v3 = append(v3, 42)
        for fi18 := 0; fi18 < 10; fi18++ {
            switch "" {
            case "hello":
                if len(v3) > 0 {
                    v3[0] = 63
                }
            }
            fmt.Println(v16, v15, fi18)
            var v19 uint8 = 255
            _ = fi18
            _ = v19
        }
        v1 = -1
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint8 = 18
        _ = fi20
        _ = v21
    }
    fmt.Println(v2, v17, v1)
    if (32767 >= -730) {
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 float64 = (1.5 - 3.14159)
            fmt.Println(v23, fi22)
            var v24 float64 = v23
            _ = fi22
            _ = v23
            _ = v24
        }
    } else {
        switch "hello" {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = 2
            }
        case "":
            if (true && (!true)) {
                var v25 uint = (v17 % ((v17) | 1))
                _ = v25
            }
        }
        v17 = (^(v17))
    }
    fmt.Println(v15)
    if (!(!true)) {
        switch "" {
        case "foo bar":
            var v26 uint8 = 211
            _ = v26
        case "a":
            if ((true && false) && ((0 <= 515) && false)) {
                if true {
                    v27 := GenS2{F0: 243, F1: genFn0((v1 ^ 0), v15, v2), F2: 1}
                    _ = v27
                } else {
                    var v28 float32 = (0.1 * 3.14159)
                    fmt.Println(v16, v28)
                    _ = len(v3)
                    _ = v28
                }
                v3 = append(v3, 42)
                v29 := GenS3{F0: 32767, F1: 0, F2: v2}
                _ = v29
            } else {
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 int64 = (v1 << (((^(v17))) & 63))
                    v16 += (-(v16))
                    _ = fi30
                    _ = v31
                }
                var v32 bool = true
                _ = v32
            }
        case "hello":
            var v33 float32 = (((((3.14159 + 1e10) * (1.5 + float32(1.5))) + 0.1) - (-2.25 - 1.5)) / 1e10)
            _ = v33
        }
        v2 = 1
    } else {
        var v34 int16 = genFn0((v2 - v2), -110, (v2 | (v1 / ((genFn1(true, v16, -1)) | 1))))
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        v1 = v1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi35
    }
    v2 = -672
    var v36 int32 = (v15 - (v15 ^ 42))
    var v37 float64 = 0.1
    v16 += v16
    var v38 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
