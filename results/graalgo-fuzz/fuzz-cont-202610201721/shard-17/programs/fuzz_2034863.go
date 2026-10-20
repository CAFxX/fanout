package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 uint32 = 953
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 int16 = -32768
    v5 := make(map[uint32]int)
    var v6 uint32 = v1
    var v7 int64 = 0
    var v8 float32 = (-2.25 * 3.14159)
    for fi9 := 0; fi9 < 6; fi9++ {
        if false {
            var v10 uint64 = 4294967295
            _ = v10
        } else {
            var v11 float64 = 0.0
            fmt.Println(v11, v8, v7)
            _ = v11
        }
        var v12 uint64 = 380
        for fi13 := 0; fi13 < 3; fi13++ {
            switch "\n\t" {
            case "":
                delete(v3, (fi9 ^ -9223372036854775808))
            case "foo bar":
                v14 := GenS1{F0: (v12 << ((1) & 63)), F1: -1, F2: 883}
                _ = v14
            }
            var v15 bool = ((((3.14159 + float32(float64(v8))) - 1e10) + 0.0) != v8)
            var v16 int16 = 32767
            _ = fi13
            _ = v15
            _ = v16
        }
        _ = fi9
        _ = v12
    }
    switch 32768 {
    case -1:
        fmt.Println(v7)
    }
    v17 := []int{1}
    delete(v2, (^(v6)))
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 uint16 = 1
        if (42 >= 255) {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            var v20 int8 = 127
            var v21 string = ("\n\t" + ("foo bar" + "\n\t"))
            _ = v20
            _ = v21
        } else {
            var v22 int = 1
            switch ("\n\t" + "\n\t") {
            case "a":
                v4 += -796
            case "\n\t":
                for fi23 := 0; fi23 < 4; fi23++ {
                    v1 += (v1 * v6)
                    v1 -= 0
                    switch (^(fi23)) {
                    case 42:
                        var v24 int64 = (v7 % ((v7) | 1))
                        _ = v24
                    default:
                        var v25 int16 = (v4 ^ 0)
                        _ = v25
                    }
                    var v26 uint = 2
                    if (41 >= -34) {
                        switch (-(v22)) {
                        case 388:
                            fmt.Println(v1)
                        case 242:
                            v27 := GenS2{F0: 4294967295, F1: v6}
                            _ = v27
                        default:
                            v3[(v22 | (fi18 | v22))] = (v6 ^ (v6 | v1))
                        }
                        v8 -= 1e10
                    } else {
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        v8 -= v8
                    }
                    _ = fi23
                    _ = v26
                }
            }
            v22 -= fi18
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    v28 := GenS3{F0: 2, F1: v4, F2: (v4 & v4)}
    v4 += 900
    v29 := GenS4{F0: 4294967295}
    v17 = append(v17, 2147483647)
    var v30 uint = 378
    if (-2147483648 > 2147483647) {
        var v31 string = ("hello" + "")
        _ = v31
    } else {
        if ((v4 - v4) >= v4) {
            delete(v5, v6)
            v6 = (v1 % (((v6 >> ((v30) & 63))) | 1))
        } else {
            var v32 int32 = 42
            _ = v32
        }
    }
    if true {
        v1 = (v6 - v6)
        var v33 uint16 = 1
        _ = v33
    }
    v3[714] = v6
    switch ((("a" + "") + "a") + "\n\t") {
    case "":
        _ = len(v17)
    case "\n\t":
        switch 128 {
        case 0:
            v30 -= (v30 ^ 255)
        }
    }
    v34 := GenS4{F0: 0}
    if (!(-128 != 1)) {
        v35 := GenS5{F0: 171, F1: v6}
        switch "foo bar" {
        case "a":
            v8 = (float32(1e10) + ((v8 + v8) * (v8 - v8)))
        default:
            v36 := GenS4{F0: 611}
            _ = v36
        }
        var v37 int64 = (v7 ^ 143)
        _ = v35
        _ = v37
    } else {
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
