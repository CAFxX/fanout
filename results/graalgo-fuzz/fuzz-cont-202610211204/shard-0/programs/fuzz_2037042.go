package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
func genFn0(p0 int8) uint32 {
    v5 := GenS1{F0: 384}
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            switch "\n\t" {
            case "a":
                fmt.Println(p0, fi7)
            case "\n\t":
                p0 += p0
            case "hello":
                v8 := []int{fi7, (fi7 >> ((930) & 63))}
                _ = v8
            }
            p0 = (p0 + (p0 | (p0 ^ -128)))
            _ = fi7
        }
        _ = fi6
    }
    switch 2147483647 {
    case -408:
        if false {
            v9 := GenS1{F0: int16(p0)}
            v10 := GenS1{F0: 0}
            fmt.Println(p0)
            _ = v10
            _ = v9
        } else {
            var v11 bool = (!false)
            v11 = (!v11)
            var v12 int = -2147483648
            _ = v11
            _ = v12
        }
    case 127:
        fmt.Println(p0)
    default:
        v13 := []int{}
        _ = v13
    }
    v14 := GenS1{F0: 42}
    _ = p0
    _ = v14
    _ = v5
    return 2
}

func main() {
    var v1 int = 128
    var v2 bool = true
    v3 := make(map[string]string)
    v4 := make(map[uint32]bool)
    v15 := GenS1{F0: 0}
    v16 := GenS2{F0: 255, F1: 337, F2: 42}
    for fi17 := 0; fi17 < 18; fi17++ {
        v1 += (fi17 ^ v1)
        _ = fi17
    }
    var v18 string = ""
    if (0.0 > 0.0) {
        v1 = 1
    }
    var v19 int16 = 1
    switch 42 {
    case 32768:
        var v20 string = "hello"
        _ = v20
    case 127:
        delete(v4, 560)
    }
    var v21 float32 = (((0.0 / 1.0) - 0.0) / (1e10 - ((1.5 - 1e10) + 1.5)))
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int8 = 42
        _ = len(v3)
        switch (v18 + v18) {
        case "hello":
            var v24 bool = true
            _ = v24
        default:
            v25 := GenS1{F0: 620}
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    v21 = float32(-2.25)
    switch v18 {
    case "":
        switch (v1 | 574) {
        case 4294967295:
            if v2 {
                var v26 float32 = v21
                var v27 uint16 = 42
                _ = v26
                _ = v27
            }
        case 943:
            for fi28 := 0; fi28 < 15; fi28++ {
                _ = len(v3)
                var v29 bool = (false || ((3.14159 / ((float64(v21) - 0.1) / 1e10)) < 1.5))
                v30 := []int{(fi28 - (fi28 * (fi28 * -997)))}
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            var v31 uint8 = 127
            _ = v31
        }
    case "hello":
        fmt.Println(v19)
    case "a":
        if true {
            var v32 int8 = 100
            for fi33 := 0; fi33 < 8; fi33++ {
                for fi34 := 0; fi34 < 13; fi34++ {
                    v35 := GenS1{F0: 0}
                    if (!(2147483648 != 453)) {
                        var v36 int32 = 4
                        v32 = (v32 % ((v32) | 1))
                        _ = v36
                    }
                    _ = fi34
                    _ = v35
                }
                switch v1 {
                case 942:
                    if v2 {
                        var v37 int = v1
                        fmt.Println(v21, fi33, v19)
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 12; fi38++ {
                            v39 := []int{v1, fi38}
                            var v40 int32 = -1
                            v41 := GenS3{F0: fi33, F1: 42, F2: 263}
                            v4[65535] = v2
                            switch v18 {
                            case "\n\t":
                                var v42 int32 = v40
                                _ = v42
                            default:
                                var v43 uint32 = genFn0(v32)
                                _ = v43
                            }
                            v3[v18] = v18
                            v2 = v2
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        v44 := GenS2{F0: 1, F1: 0, F2: 42}
                        _ = v44
                    }
                case 128:
                    v45 := []int{}
                    _ = v45
                default:
                    for fi46 := 0; fi46 < 13; fi46++ {
                        v47 := GenS1{F0: v19}
                        var v48 uint64 = 433
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                }
                v32 -= v32
                _ = fi33
            }
            var v49 int16 = v19
            _ = v32
            _ = v49
        }
    }
    var v50 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v50
}
