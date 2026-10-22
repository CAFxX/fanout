package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int8; F1 uint16; F2 int }
type GenS3 struct { F0 int32; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 string }
func genFn0(p0 uint8) uint64 {
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := []int{(fi4 / ((fi4) | 1)), (fi4 << ((2) & 63)), 42}
        v6 := GenS1{F0: 9223372036854775807, F1: (fi4 - int(p0))}
        v5 = append(v5, (fi4 ^ fi4))
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{-1}
    var v8 uint32 = 456
    if true {
        v9 := GenS2{F0: -1, F1: 32767, F2: 261}
        switch 128 {
        case 802:
            var v10 int = int(p0)
            _ = v10
        default:
            v7 = append(v7, 32768)
        }
        v11 := GenS3{F0: 1, F1: 0, F2: 32767}
        _ = v11
        _ = v9
    }
    var v12 int8 = 127
    _ = p0
    _ = v12
    _ = v7
    _ = v8
    return uint64(v8)
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 uint32 = 1
    if true {
        if false {
            v13 := GenS2{F0: -109, F1: 42, F2: 32767}
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 7; fi15++ {
                v3 += (v3 & 2147483648)
                _ = fi15
            }
            v1["a"] = ("a" + "foo bar")
            v3 -= (^(v3))
            _ = fi14
        }
    }
    switch -9223372036854775808 {
    case -2147483648:
        _ = len(v2)
    case -408:
        if true {
            if (!true) {
                fmt.Println(v3)
                var v16 int = -11
                v17 := GenS2{F0: 116, F1: 655, F2: v16}
                _ = v16
                _ = v17
            }
            switch ((("" + "hello") + "a") + "\n\t") {
            case "hello":
                if false {
                    v3 = (v3 ^ v3)
                    var v18 float32 = ((-2.25 + -2.25) - 1.5)
                    for fi19 := 0; fi19 < 2; fi19++ {
                        var v20 float64 = (-2.25 + 1e10)
                        var v21 int16 = -32768
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = v18
                } else {
                    for fi22 := 0; fi22 < 15; fi22++ {
                        fmt.Println(fi22)
                        _ = fi22
                    }
                }
            case "a":
                for fi23 := 0; fi23 < 11; fi23++ {
                    v3 += v3
                    _ = fi23
                }
            case "":
                fmt.Println(v3)
            }
            v24 := GenS3{F0: -263, F1: 18446744073709551615, F2: 1}
            _ = v24
        }
    case -181:
        switch -89 {
        case 128:
            if false {
                v3 -= 2147483647
                var v25 int64 = 9223372036854775807
                _ = len(v2)
                _ = v25
            }
        }
    default:
        v26 := GenS4{F0: 97, F1: 0, F2: "a"}
        _ = v26
    }
    fmt.Println(v3)
    if ((107 > 42) || (-914 <= -1)) {
        var v27 int32 = 42
        switch "\n\t" {
        case "\n\t":
            var v28 bool = (-2147483648 == -1)
            _ = v28
        }
        v27 += v27
        _ = v27
    } else {
        v3 = 768
    }
    var v29 uint = 0
    v30 := []int{}
    fmt.Println(v29)
    if false {
        if (!true) {
            var v31 uint64 = 2
            var v32 uint = v29
            _ = v31
            _ = v32
        }
        if (0.1 == (float32(0.1) * (-2.25 - 0.0))) {
            var v33 uint8 = 41
            var v34 int16 = 1
            _ = v33
            _ = v34
        }
    } else {
        for fi35 := 0; fi35 < 6; fi35++ {
            v3 = (v3 >> (((v29 / ((1) | 1))) & 63))
            _ = fi35
        }
        if (-1 <= 0) {
            v36 := GenS1{F0: -1, F1: 4294967295}
            v29 -= v29
            _ = v36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
