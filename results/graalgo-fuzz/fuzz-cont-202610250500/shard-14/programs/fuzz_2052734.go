package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int8; F2 int16 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int = 32768
    v2 := []string{"", ("\n\t" + "\n\t"), "a", (("" + "foo bar") + "hello")}
    v3 := []uint32{}
    v4 := make(map[uint32]string)
    v5 := []string{(("a" + "\n\t") + ""), ""}
    var v6 int64 = -1
    switch v1 {
    case 128:
        v1 = v1
    case 2147483647:
        for fi7 := 0; fi7 < 15; fi7++ {
            delete(v4, 2147483648)
            _ = len(v3)
            _ = fi7
        }
    }
    switch "hello" {
    case "":
        v1 += v1
    case "hello":
        if ((3.14159 + 1e10) <= 1e10) {
            fmt.Println(v6, v1)
            var v8 uint64 = 0
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint16 = 0
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            var v11 uint32 = 2147483647
            if false {
                var v12 int = (-(v1))
                v11 = v11
                _ = v12
            }
            _ = v11
        }
    }
    var v13 bool = (1 > 42)
    v14 := GenS1{F0: 376}
    if (867 != 42) {
        switch v1 {
        case -342:
            v15 := GenS1{F0: 2147483647}
            _ = v15
        case -1:
            var v16 int64 = -957
            _ = v16
        default:
            var v17 float32 = ((0.1 - 1e10) - (0.1 - 1e10))
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 float64 = 1e10
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 bool = v13
            _ = fi20
            _ = v21
        }
        v13 = (-722 == -2147483648)
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        switch "hello" {
        case "hello":
            if v13 {
                if ((fi22 / ((fi22) | 1)) < (v1 | v1)) {
                    v23 := GenS1{F0: 0}
                    v24 := GenS1{F0: 2147483647}
                    v13 = ("hello" > "foo bar")
                    _ = v23
                    _ = v24
                } else {
                    var v25 uint64 = 87
                    _ = v25
                }
                v6 -= v6
                v26 := GenS1{F0: 930}
                _ = v26
            } else {
                v27 := GenS1{F0: 900}
                v28 := GenS2{F0: 9223372036854775807, F1: -128, F2: -32768}
                _ = v27
                _ = v28
            }
        }
        _ = fi22
    }
    if v13 {
        switch v1 {
        case 128:
            if (4294967295 == 4294967295) {
                var v29 float32 = ((float32(float64((0.0 / 1.0))) - -2.25) - 3.14159)
                var v30 uint = 819
                fmt.Println(v6, v6)
                _ = v29
                _ = v30
            } else {
                switch v1 {
                case -161:
                    var v31 int64 = v6
                    _ = v31
                }
                switch "hello" {
                case "":
                    v6 -= -671
                case "foo bar":
                    fmt.Println(v6, v6, v1)
                case "\n\t":
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                }
            }
        case 4294967295:
            v1 = (v1 >> ((1) & 63))
        default:
            v1 = v1
        }
        var v32 uint64 = 1
        _ = v32
    }
    v33 := v4[2147483647]
    v34 := GenS3{F0: 255}
    var v35 string = ((("\n\t" + "\n\t") + "") + v33)
    v33 = (v33 + (v35 + v33))
    var v36 uint16 = 1
    v35 = "hello"
    fmt.Println(v35, v35, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
