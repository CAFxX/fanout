package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint64; F1 int32 }
func genFn0(p0 string, p1 uint, p2 string) int8 {
    var v5 bool = (230 < 1)
    var v6 bool = false
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: p1, F1: 18446744073709551615, F2: 4294967295}
        switch "\n\t" {
        case "\n\t":
            if false {
                for fi9 := 0; fi9 < 13; fi9++ {
                    fmt.Println(v5, p2)
                    var v10 int64 = -185
                    _ = fi9
                    _ = v10
                }
                v5 = true
                p2 = p2
            }
        case "":
            v11 := []int{fi7, (fi7 - 42)}
            _ = v11
        default:
            v12 := GenS1{F0: (^(p1)), F1: (^(p1)), F2: 670}
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 23
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 0
    var v3 int64 = 9223372036854775807
    var v4 float32 = 1e10
    fmt.Println(v3, v3)
    var v13 int8 = 45
    v14 := []int{}
    var v15 string = "hello"
    if len(v14) > 0 {
        v14[0] = 42
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        if false {
            var v17 uint32 = 42
            switch -2147483648 {
            case 918:
                v18 := GenS2{F0: (v17 / (((v17 + v17)) | 1)), F1: v17, F2: -471}
                _ = v18
            }
            _ = v17
        }
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi16
    }
    v4 -= (v4 * 3.14159)
    var v19 int = -1
    switch v19 {
    case 32768:
        fmt.Println(v4, v3, v19)
    }
    if (42 < -32768) {
        fmt.Println(v15, v3)
    } else {
        var v20 uint8 = 255
        v21 := v1[771]
        switch (v19 ^ (-(v19))) {
        case 32767:
            fmt.Println(v2)
        case -2147483648:
            fmt.Println(v20, v13, v19)
        default:
            var v22 float64 = (0.1 + 1e10)
            _ = v22
        }
        _ = v20
        _ = v21
    }
    v23 := v1[65535]
    fmt.Println(v13, v4, v15)
    v2 = v2
    if true {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        v1[2147483648] = v23
        switch ("hello" + ((v15 + "\n\t") + "hello")) {
        case "a":
            switch (v15 + "") {
            case "a":
                var v24 float32 = (3.14159 + (float32(3.14159) + 1.5))
                _ = v24
            case "hello":
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := GenS3{F0: genFn0(v15, 883, v15), F1: ((v19 - 127) == (-(v19)))}
                    var v27 uint8 = 128
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        case "foo bar":
            var v28 int32 = 42
            _ = v28
        case "hello":
            fmt.Println(v2, v23, v19)
        default:
            switch (v23 | (v23 | 681)) {
            case 4294967295:
                fmt.Println(v3)
            case -1:
                if false {
                    var v29 uint = 157
                    _ = v29
                }
            case -503:
                _ = len(v1)
            }
        }
    } else {
        v2 += -846
        v30 := GenS4{F0: 940, F1: 563}
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
