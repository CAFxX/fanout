package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int16; F2 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 int16; F1 int64; F2 uint32 }
func genFn0(p0 int32, p1 uint, p2 int16) uint8 {
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = 0
    var v3 uint8 = 36
    var v4 uint64 = 2
    var v5 string = "hello"
    if ((((392 >= 641) || true) || (v4 <= (v1 * v1))) || false) {
        v6 := []int{128}
        fmt.Println(v3, v4, v4)
        _ = v6
    } else {
        v5 = v5
    }
    fmt.Println(v2, v1)
    v7 := GenS1{F0: -949}
    var v8 int16 = 32767
    for fi9 := 0; fi9 < 9; fi9++ {
        if false {
            if (65535 < 255) {
                switch 128 {
                case -963:
                    v4 = 1
                case -2147483648:
                    for fi10 := 0; fi10 < 11; fi10++ {
                        v11 := GenS1{F0: 597}
                        var v12 int64 = (^(v2))
                        if (4294967295 != (v1 / ((v4) | 1))) {
                            v13 := []int{}
                            v14 := GenS1{F0: fi9}
                            v5 = (v5 + v5)
                            _ = v13
                            _ = v14
                        }
                        var v15 int32 = -2147483648
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v15
                    }
                default:
                    var v16 int32 = -2147483648
                    _ = v16
                }
            } else {
                v17 := GenS1{F0: fi9}
                switch v5 {
                case "hello":
                    v18 := []int{(fi9 >> ((255) & 63))}
                    _ = v18
                case "a":
                    if ((^(v3)) > (^(v3))) {
                        for fi19 := 0; fi19 < 11; fi19++ {
                            v20 := GenS2{F0: (-(v2)), F1: 2, F2: v5}
                            switch (fi9 | (^(fi9))) {
                            case 32767:
                                v1 -= v1
                            case 2147483647:
                                var v21 float32 = ((float32(1e10) + float32(1.5)) - (0.1 * (float32((3.14159 - float64(1e10))) * (3.14159 + 1.5))))
                                fmt.Println(fi19, v1)
                                v4 -= 177
                                _ = v21
                            }
                            _ = fi19
                            _ = v20
                        }
                        switch ("hello" + "a") {
                        case "a":
                            for fi22 := 0; fi22 < 17; fi22++ {
                                v23 := GenS1{F0: fi9}
                                fmt.Println(fi22, fi22, v2)
                                fmt.Println(v5, v8)
                                _ = fi22
                                _ = v23
                            }
                        case "\n\t":
                            v4 += 0
                        case "foo bar":
                            v24 := GenS2{F0: 42, F1: 1, F2: "\n\t"}
                            _ = v24
                        default:
                            v2 = v2
                        }
                        v8 = v8
                    } else {
                        v25 := GenS3{F0: (v3 ^ (^(v3)))}
                        _ = v25
                    }
                case "":
                    for fi26 := 0; fi26 < 20; fi26++ {
                        var v27 float64 = 3.14159
                        var v28 uint = 18446744073709551615
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                default:
                    v29 := GenS4{F0: 0}
                    _ = v29
                }
                switch v5 {
                case "foo bar":
                    v2 += 695
                case "a":
                    for fi30 := 0; fi30 < 15; fi30++ {
                        fmt.Println(v1)
                        _ = fi30
                    }
                case "":
                    switch "foo bar" {
                    case "foo bar":
                        switch (("\n\t" + v5) + "\n\t") {
                        case "":
                            if (4294967295 != 65535) {
                                v31 := GenS5{F0: "a", F1: (v8 * v8), F2: 28}
                                if true {
                                    fmt.Println(v3, v4)
                                }
                                var v32 uint8 = v3
                                _ = v31
                                _ = v32
                            } else {
                                var v33 uint = 70
                                v34 := []int{fi9, fi9, fi9, 1}
                                _ = v33
                                _ = v34
                            }
                        }
                    }
                default:
                    var v35 string = v5
                    _ = v35
                }
                _ = v17
            }
        }
        v36 := []int{fi9, (fi9 + fi9), 127, fi9}
        v8 = 42
        _ = fi9
        _ = v36
    }
    switch v5 {
    case "a":
        for fi37 := 0; fi37 < 1; fi37++ {
            fmt.Println(v5)
            v38 := []int{fi37}
            _ = fi37
            _ = v38
        }
    case "\n\t":
        switch (v5 + ("foo bar" + v5)) {
        case "foo bar":
            v8 -= v8
        case "hello":
            var v39 uint16 = 1
            _ = v39
        }
    }
    if true {
        v2 += v2
        var v40 float32 = ((0.1 - 0.0) - 0.1)
        v41 := GenS6{F0: 42}
        _ = v40
        _ = v41
    }
    v42 := GenS7{F0: 42, F1: v2}
    v43 := GenS8{F0: (v8 % (((v8 * -32768)) | 1)), F1: (v2 * v2), F2: 344}
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v7
    _ = v8
}
