package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 bool }

func main() {
    var v1 int64 = -68
    var v2 float32 = 1e10
    var v3 uint = 65535
    var v4 string = ("\n\t" + "")
    v5 := make(map[string]string)
    v6 := make(map[string]int)
    v1 = int64(v3)
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int64 = (v1 >> (((v3 | v3)) & 63))
        v9 := GenS1{F0: 846, F1: 244, F2: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 float64 = 1e10
        v12 := GenS2{F0: (42 >= 550)}
        if false {
            v11 -= 1e10
            if false {
                v1 -= (v1 + (v1 | int64(v3)))
                v11 = 3.14159
                v13 := []int{(fi10 - (fi10 ^ fi10))}
                _ = v13
            }
            fmt.Println(v2, v1)
        } else {
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 int16 = -1
                v16 := []int{}
                switch (v4 + (("a" + v4) + (v4 + v4))) {
                case "a":
                    v17 := GenS3{F0: 103}
                    _ = v17
                case "":
                    var v18 uint64 = 42
                    _ = v18
                }
                var v19 float32 = (v2 - (1.5 + -2.25))
                _ = fi14
                _ = v15
                _ = v16
                _ = v19
            }
            var v20 int16 = -951
            _ = v20
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v1 = v1
    v1 = (v1 % ((-120) | 1))
    v21 := []int{}
    v22 := GenS4{F0: v4}
    for fi23 := 0; fi23 < 2; fi23++ {
        v2 = (v2 - 1.5)
        if (!true) {
            if (-1 < int8(v3)) {
                v4 = v4
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 uint16 = 1
                    var v26 int16 = 1
                    var v27 float32 = ((v2 * ((v2 - v2) - v2)) + v2)
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
            v21 = append(v21, (fi23 % (((fi23 / (((fi23 | fi23)) | 1))) | 1)))
            v1 -= v1
        } else {
            for fi28 := 0; fi28 < 2; fi28++ {
                var v29 float64 = (0.0 * (1e10 + 1e10))
                v30 := GenS5{F0: (fi23 + fi23)}
                if (!false) {
                    if true {
                        var v31 int16 = -412
                        v2 = v2
                        v1 = (v1 >> (((v3 * (v3 ^ 1))) & 63))
                        _ = v31
                    } else {
                        var v32 int16 = 32767
                        v32 = 0
                        _ = v32
                    }
                }
                _ = fi28
                _ = v29
                _ = v30
            }
            switch fi23 {
            case 4294967295:
                var v33 uint64 = 2
                _ = v33
            default:
                if (42 > (v1 << ((v3) & 63))) {
                    var v34 float32 = float32((3.14159 + float64(float32(0.1))))
                    var v35 int64 = v1
                    v36 := GenS1{F0: 0, F1: 4294967295, F2: 213}
                    _ = v34
                    _ = v35
                    _ = v36
                } else {
                    v37 := GenS2{F0: (0 == (v3 * v3))}
                    v38 := GenS2{F0: false}
                    if len(v21) > 0 {
                        v21[0] = fi23
                    }
                    _ = v37
                    _ = v38
                }
            }
        }
        v39 := GenS6{F0: -454, F1: true}
        _ = fi23
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
