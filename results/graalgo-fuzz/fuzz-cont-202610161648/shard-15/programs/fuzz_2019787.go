package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 bool }

func main() {
    v1 := []uint32{2147483648, 694, 2147483647}
    v2 := []uint32{739}
    v3 := []string{("hello" + "hello"), "a", "a", ""}
    for fi4 := 0; fi4 < 9; fi4++ {
        if true {
            var v5 int64 = -1
            var v6 bool = true
            _ = v5
            _ = v6
        } else {
            v7 := GenS1{F0: -982, F1: 1, F2: (true || true)}
            fmt.Println(fi4)
            _ = v7
        }
        switch ("hello" + "") {
        case "foo bar":
            if true {
                var v8 int8 = 42
                if true {
                    var v9 bool = (fi4 < fi4)
                    for fi10 := 0; fi10 < 2; fi10++ {
                        fmt.Println(v9)
                        v9 = ((v9 && v9) && (!(-1 > -563)))
                        _ = fi10
                    }
                    for fi11 := 0; fi11 < 19; fi11++ {
                        var v12 uint16 = 711
                        _ = fi11
                        _ = v12
                    }
                    _ = v9
                } else {
                    var v13 float32 = -2.25
                    _ = v13
                }
                _ = v8
            }
        case "":
            v2 = append(v2, 9)
        }
        fmt.Println(fi4)
        _ = fi4
    }
    if (!(("a" + ("foo bar" + "a")) == "\n\t")) {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 uint64 = 0
            for fi16 := 0; fi16 < 17; fi16++ {
                _ = len(v1)
                _ = fi16
            }
            switch fi14 {
            case 9223372036854775807:
                fmt.Println(fi14)
            case 89:
                fmt.Println(fi14, v15)
            case 32767:
                v15 += 8
            default:
                if len(v2) > 0 {
                    v2[0] = 293
                }
            }
            _ = fi14
            _ = v15
        }
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS2{F0: 727, F1: 1}
        fmt.Println(fi17)
        _ = fi17
        _ = v18
    }
    var v19 uint16 = 0
    var v20 float64 = ((float64(1e10) / 1.5) / 3.14159)
    switch "\n\t" {
    case "hello":
        for fi21 := 0; fi21 < 9; fi21++ {
            v19 = v19
            _ = fi21
        }
    }
    v22 := GenS3{F0: true}
    v23 := GenS2{F0: 958, F1: 42}
    v19 = v19
    var v24 uint64 = 1
    if true {
        v20 -= 1.5
        var v25 float64 = (0.0 * v20)
        var v26 float64 = ((3.14159 * ((v25 / (float64(0.0) + 0.1)) + (v25 * ((3.14159 - 3.14159) + v25)))) / ((0.0 - 3.14159) - v25))
        _ = v25
        _ = v26
    }
    if ("hello" > "foo bar") {
        v24 = v24
        v20 += (v20 / v20)
    }
    if (-247 == 32767) {
        if len(v3) > 0 {
            v3[0] = "a"
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 uint16 = (v19 | 65535)
            _ = fi27
            _ = v28
        }
        switch "a" {
        case "foo bar":
            fmt.Println(v24, v19)
        case "a":
            var v29 int16 = -1
            _ = v29
        default:
            switch 9223372036854775807 {
            case 32767:
                v30 := GenS1{F0: 1000, F1: 0, F2: false}
                _ = v30
            case 128:
                v20 = 0.1
            case 0:
                var v31 float32 = 0.1
                _ = v31
            }
        }
    }
    _ = len(v2)
    var v32 uint16 = (^(v19))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
}
