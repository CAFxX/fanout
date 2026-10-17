package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int; F2 int }
type GenS3 struct { F0 int8; F1 string; F2 int }
type GenS4 struct { F0 uint16; F1 int8 }

func main() {
    v1 := []int{-914}
    var v2 int16 = -32768
    v3 := make(map[string]string)
    if true {
        _ = len(v1)
    }
    fmt.Println(v2)
    if (true && true) {
        v2 = 0
        for fi4 := 0; fi4 < 5; fi4++ {
            fmt.Println(v2, v2)
            _ = fi4
        }
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint = 0
        var v7 uint32 = 2
        _ = fi5
        _ = v6
        _ = v7
    }
    v2 -= 32767
    v8 := GenS1{F0: "\n\t"}
    var v9 uint32 = 94
    _ = len(v1)
    for fi10 := 0; fi10 < 16; fi10++ {
        if (((3.14159 / -2.25) + 0.0) == -2.25) {
            for fi11 := 0; fi11 < 5; fi11++ {
                if (true || ((1.5 / (1e10 * 3.14159)) <= (1.5 - (0.1 * ((float32(((3.14159 - float64(3.14159)) / (float64(0.0) + float64(3.14159)))) + 0.1) * (((0.0 / 0.1) * 1e10) + (((1.5 / (1.5 / 1.0)) + 1.5) * (3.14159 * float32((1.5 - -2.25)))))))))) {
                    fmt.Println(fi11, fi11)
                }
                _ = fi11
            }
        }
        var v12 int64 = 0
        var v13 uint32 = v9
        _ = fi10
        _ = v12
        _ = v13
    }
    delete(v3, "")
    if (!(!(-550 == -530))) {
        for fi14 := 0; fi14 < 18; fi14++ {
            fmt.Println(fi14, v2, v2)
            v15 := GenS2{F0: 100, F1: fi14, F2: 32768}
            if (42 < 1) {
                v9 += v9
                for fi16 := 0; fi16 < 18; fi16++ {
                    switch fi16 {
                    case -50:
                        var v17 int32 = -870
                        _ = v17
                    case 4294967295:
                        fmt.Println(fi16, fi14, v9)
                    case 32767:
                        if len(v1) > 0 {
                            v1[0] = fi14
                        }
                    }
                    _ = fi16
                }
                if false {
                    v18 := GenS3{F0: 1, F1: ("hello" + "\n\t"), F2: fi14}
                    fmt.Println(fi14, v9, fi14)
                    _ = v18
                }
            }
            _ = fi14
            _ = v15
        }
        var v19 string = "\n\t"
        _ = v19
    }
    delete(v3, "")
    fmt.Println(v9)
    var v20 float32 = 0.1
    for fi21 := 0; fi21 < 12; fi21++ {
        if (!(-2.25 < v20)) {
            v2 = v2
            var v22 int = fi21
            _ = v22
        }
        if true {
            fmt.Println(v20, v20, fi21)
        } else {
            v20 = 1e10
            v1 = append(v1, (fi21 & fi21))
        }
        switch fi21 {
        case 455:
            var v23 uint32 = 65535
            _ = v23
        case 1:
            if true {
                var v24 bool = false
                _ = v24
            }
        }
        _ = fi21
    }
    if ((v9 & (v9 % (((^(v9))) | 1))) <= (v9 - (v9 - v9))) {
        v25 := GenS4{F0: 1, F1: 127}
        v26 := GenS4{F0: 65535, F1: -48}
        _ = v25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
