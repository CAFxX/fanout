package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint16; F1 int32; F2 bool }
func genFn0() uint32 {
    switch "\n\t" {
    case "a":
        if (!(0 != 0)) {
            v7 := GenS1{F0: 1, F1: 2}
            _ = v7
        } else {
            var v8 uint64 = 1
            v9 := GenS1{F0: -1, F1: (^(v8))}
            fmt.Println(v8)
            _ = v8
            _ = v9
        }
    case "hello":
        var v10 int64 = 0
        _ = v10
    case "\n\t":
        var v11 int16 = 1
        _ = v11
    }
    var v12 uint64 = 2
    var v13 string = (("foo bar" + ("\n\t" + ("hello" + "foo bar"))) + "hello")
    _ = v12
    _ = v13
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -32768
    var v3 uint = 255
    v4 := make(map[int]int)
    var v5 bool = true
    var v6 uint32 = 0
    var v14 uint32 = v6
    delete(v4, 254)
    v15 := []int{}
    switch "hello" {
    case "\n\t":
        v4[42] = 0
    case "":
        v16 := GenS2{F0: v6, F1: v14, F2: v2}
        _ = v16
    default:
        v17 := GenS2{F0: (^(v14)), F1: genFn0(), F2: v2}
        _ = v17
    }
    v18 := GenS3{F0: 65535}
    switch 4294967295 {
    case 1:
        var v19 string = "\n\t"
        _ = v19
    case -884:
        var v20 float32 = 3.14159
        _ = v20
    default:
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 int16 = (v2 & v2)
            var v23 uint8 = 176
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    fmt.Println(v6)
    v6 = genFn0()
    for fi24 := 0; fi24 < 5; fi24++ {
        v25 := GenS2{F0: (v14 * 2), F1: (v6 >> (((v3 & (v3 ^ (v3 | v3)))) & 63)), F2: (v2 & v2)}
        fmt.Println(fi24, v14, v2)
        _ = fi24
        _ = v25
    }
    switch "foo bar" {
    case "a":
        _ = len(v15)
    case "hello":
        var v26 int64 = -429
        _ = v26
    case "\n\t":
        for fi27 := 0; fi27 < 19; fi27++ {
            switch fi27 {
            case 127:
                var v28 float64 = (((0.0 + 1.5) / ((float64(float32((float64(3.14159) + float64(1e10)))) / (-2.25 * (1.5 / -2.25))) + (3.14159 / 3.14159))) / float64(1.5))
                _ = v28
            case -116:
                v29 := GenS2{F0: v6, F1: genFn0(), F2: -1}
                _ = v29
            }
            var v30 uint16 = 255
            if v5 {
                var v31 int16 = v2
                if false {
                    delete(v4, 28)
                } else {
                    v32 := GenS4{F0: (v3 >> (((v3 / (((v3 - (^(v3)))) | 1))) & 63)), F1: -15}
                    _ = v32
                }
                if len(v15) > 0 {
                    v15[0] = fi27
                }
                _ = v31
            }
            _ = fi27
            _ = v30
        }
    default:
        var v33 uint8 = 152
        _ = v33
    }
    if v5 {
        v1["a"] = ((65535 > 0) && (127 > int(v14)))
        v6 = v14
        if (1 == 255) {
            if v5 {
                fmt.Println(v5, v5, v5)
            }
            var v34 int8 = 127
            _ = v34
        } else {
            v6 += v6
        }
    }
    if ((!(!(-949 != -122))) && (v5 && (true && (int32(v14) <= 0)))) {
        var v35 uint16 = 32767
        v36 := GenS5{F0: (^(v35)), F1: -1, F2: true}
        _ = v35
        _ = v36
    }
    var v37 float32 = 0.1
    if (!v5) {
        var v38 int8 = 42
        var v39 bool = (127 < -2147483648)
        _ = v38
        _ = v39
    } else {
        switch "\n\t" {
        case "":
            _ = len(v4)
        case "a":
            v15 = append(v15, -470)
        default:
            v40 := v4[89]
            _ = v40
        }
        for fi41 := 0; fi41 < 13; fi41++ {
            for fi42 := 0; fi42 < 9; fi42++ {
                var v43 int64 = 973
                fmt.Println(fi42)
                _ = fi42
                _ = v43
            }
            var v44 int = (fi41 * (fi41 ^ fi41))
            _ = fi41
            _ = v44
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
