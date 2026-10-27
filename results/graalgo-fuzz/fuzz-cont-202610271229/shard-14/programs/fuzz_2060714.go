package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 uint8, p2 uint16) int {
    var v6 bool = false
    for fi7 := 0; fi7 < 12; fi7++ {
        if v6 {
            if (uint(p1) != 0) {
                v6 = (-1 >= -1)
            } else {
                switch fi7 {
                case -643:
                    v8 := GenS1{F0: "foo bar", F1: 0, F2: 1}
                    _ = v8
                case 32768:
                    var v9 float32 = 0.1
                    _ = v9
                case -1:
                    for fi10 := 0; fi10 < 1; fi10++ {
                        switch "a" {
                        case "":
                            v6 = false
                        default:
                            fmt.Println(fi10)
                        }
                        var v11 uint32 = 817
                        p1 += 215
                        _ = fi10
                        _ = v11
                    }
                default:
                    if ((42 <= 404) && true) {
                        v12 := GenS1{F0: "a", F1: 0, F2: fi7}
                        fmt.Println(v6)
                        _ = v12
                    }
                }
                var v13 int64 = 1
                _ = v13
            }
            var v14 int64 = 0
            v15 := GenS2{F0: (fi7 | fi7)}
            _ = v14
            _ = v15
        }
        v16 := GenS1{F0: ("\n\t" + ("hello" + "")), F1: 0, F2: fi7}
        if true {
            var v17 int = (fi7 >> ((0) & 63))
            var v18 int16 = 1
            fmt.Println(fi7, p2, p0)
            _ = v17
            _ = v18
        }
        _ = fi7
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1(p0 int8, p1 string, p2 uint64) int {
    v19 := []int{616, -976, 127, 2147483647}
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    p2 -= 42
    if (false && false) {
        if false {
            p1 = p1
        } else {
            var v20 int64 = -835
            p2 -= p2
            var v21 int64 = v20
            _ = v20
            _ = v21
        }
        v22 := GenS3{F0: ((("\n\t" + "a") + "\n\t") + "a")}
        _ = v22
    } else {
        var v23 uint8 = 2
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return -1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{}
    var v3 int = -1
    v4 := make(map[string]int)
    v5 := []int{}
    if true {
        v3 = -264
    }
    var v24 int8 = -128
    fmt.Println(v3, v3)
    v24 = (^(v24))
    switch "a" {
    case "":
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 int64 = -1
            v27 := v4[""]
            _ = fi25
            _ = v26
            _ = v27
        }
    case "foo bar":
        var v28 float64 = (1.5 * ((1.5 - 1e10) * ((1.5 / -2.25) + float64((((3.14159 + float32(((0.1 * 1e10) * (((0.1 - (1.5 - (1.5 - 0.1))) + 0.0) * (0.0 / 1.0))))) * float32(0.0)) - 0.1)))))
        _ = v28
    case "\n\t":
        v24 += v24
    default:
        if true {
            v24 = 127
        }
    }
    var v29 uint8 = 1
    if (1 >= 268) {
        v29 -= 127
        v24 += (v24 >> ((18446744073709551615) & 63))
        v30 := GenS1{F0: "\n\t", F1: 208, F2: (v3 + v3)}
        _ = v30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        delete(v4, ("\n\t" + "foo bar"))
        switch (("\n\t" + "foo bar") + "hello") {
        case "\n\t":
            switch "\n\t" {
            case "":
                if (false && true) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(v3, fi31, fi31)
                } else {
                    switch "\n\t" {
                    case "a":
                        var v32 int8 = (v24 & (v24 ^ -83))
                        _ = v32
                    case "":
                        _ = len(v2)
                    default:
                        for fi33 := 0; fi33 < 11; fi33++ {
                            var v34 int16 = 32767
                            v24 = v24
                            _ = fi33
                            _ = v34
                        }
                    }
                    var v35 int64 = 42
                    _ = v35
                }
            case "hello":
                v36 := GenS2{F0: fi31}
                _ = v36
            }
        case "foo bar":
            v37 := GenS1{F0: "a", F1: 2147483648, F2: fi31}
            _ = v37
        case "":
            v3 = genFn0(1, v29, 501)
        }
        _ = fi31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
