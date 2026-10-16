package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int16; F1 int64; F2 uint16 }
func genFn0(p0 int64, p1 int8, p2 int64) int16 {
    if (-2.25 < (1.5 - float64(0.1))) {
        p1 += (p1 % (((p1 - p1)) | 1))
    } else {
        if (0.0 >= -2.25) {
            for fi7 := 0; fi7 < 2; fi7++ {
                fmt.Println(p1, p2, p0)
                if ((0 <= 255) && (!true)) {
                    v8 := GenS1{F0: int32(p2)}
                    _ = v8
                }
                _ = fi7
            }
            p1 = 127
        }
    }
    v9 := []int{int(p1), 32767}
    switch 42 {
    case 128:
        var v10 uint16 = 32767
        _ = v10
    case 4294967295:
        switch 32767 {
        case 2147483647:
            v11 := GenS1{F0: 42}
            _ = v11
        }
    case 32768:
        var v12 uint = 255
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 int8 = 100
    v4 := make(map[int]int)
    v5 := []string{("a" + ""), "\n\t"}
    var v6 uint8 = 1
    v1[218] = "\n\t"
    if true {
        v13 := GenS1{F0: 0}
        _ = v13
    }
    v3 = (v3 + 100)
    v3 = v3
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            if (!false) {
                if true {
                    v6 = v6
                    switch ("a" + "hello") {
                    case "a":
                        var v16 int64 = 42
                        _ = v16
                    default:
                        var v17 int8 = v3
                        _ = v17
                    }
                    var v18 uint64 = 42
                    fmt.Println(fi15)
                    _ = v18
                } else {
                    v6 += v6
                }
            } else {
                if ((-2.25 == -2.25) || (v6 == (^(v6)))) {
                    var v19 bool = false
                    v19 = ("" >= (((("\n\t" + "\n\t") + ((("hello" + "foo bar") + "foo bar") + "a")) + ("" + "")) + "a"))
                    var v20 int16 = 341
                    var v21 int8 = (v3 / ((v3) | 1))
                    _ = v19
                    _ = v20
                    _ = v21
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            fmt.Println(fi15, v6)
            _ = len(v5)
            _ = fi15
        }
        for fi22 := 0; fi22 < 3; fi22++ {
            v5 = append(v5, "hello")
            delete(v4, (fi22 & (fi22 >> ((255) & 63))))
            _ = fi22
        }
        _ = fi14
    }
    v23 := GenS2{F0: (v6 - (^(v6))), F1: 65535, F2: (^(v6))}
    v24 := GenS1{F0: 2147483647}
    v4[4294967295] = -1
    var v25 uint = 789
    for fi26 := 0; fi26 < 20; fi26++ {
        fmt.Println(v25, v25, v25)
        v3 -= (v3 & 32)
        switch "a" {
        case "\n\t":
            if (!false) {
                v6 = v6
                var v27 int8 = (v3 - (v3 & (v3 / ((-1) | 1))))
                v28 := GenS3{F0: "foo bar", F1: 253}
                _ = v27
                _ = v28
            }
        }
        _ = fi26
    }
    v29 := GenS4{F0: "a", F1: v3}
    v30 := GenS5{F0: 0, F1: int64(v3), F2: 135}
    var v31 int16 = 42
    if true {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        var v32 float64 = 0.1
        v33 := GenS3{F0: "", F1: 998}
        _ = v32
        _ = v33
    }
    var v34 uint = (^(v25))
    var v35 uint64 = 2
    var v36 uint16 = 65535
    var v37 uint32 = 854
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
