package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 string }
type GenS2 struct { F0 uint; F1 bool; F2 uint }
type GenS3 struct { F0 int64 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int = -1
        v8 = (^(fi7))
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(fi9)
            v10 := GenS1{F0: -1, F1: 42, F2: (("\n\t" + "a") + "hello")}
            var v11 uint8 = 2
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int32 = int32(fi12)
        _ = fi12
        _ = v13
    }
    switch "a" {
    case "foo bar":
        if (("a" + ("a" + "foo bar")) != "foo bar") {
            var v14 float64 = (3.14159 + (1.5 + (0.0 - 0.0)))
            if (!true) {
                if (4294967295 <= 1) {
                    var v15 uint64 = 4294967295
                    _ = v15
                } else {
                    v16 := []int{}
                    _ = v16
                }
            } else {
                var v17 uint32 = 2147483648
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := GenS1{F0: -1, F1: 1, F2: ("hello" + ("foo bar" + "a"))}
                    var v20 uint16 = 32767
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 1; fi21++ {
                    for fi22 := 0; fi22 < 11; fi22++ {
                        v14 += v14
                        switch "\n\t" {
                        case "foo bar":
                            var v23 int32 = 0
                            _ = v23
                        case "":
                            if true {
                                if false {
                                    v24 := GenS2{F0: 1, F1: (!(fi21 == (fi21 >> ((1) & 63)))), F2: 18446744073709551615}
                                    v14 += 3.14159
                                    _ = v24
                                } else {
                                    var v25 bool = true
                                    _ = v25
                                }
                            }
                        default:
                            switch fi22 {
                            case 154:
                                var v26 uint8 = 2
                                _ = v26
                            case 42:
                                var v27 float32 = 1e10
                                _ = v27
                            case -746:
                                v28 := []int{fi22}
                                _ = v28
                            }
                        }
                        var v29 bool = (42 == 0)
                        _ = fi22
                        _ = v29
                    }
                    if (!(((^(fi21)) != fi21) && false)) {
                        fmt.Println(fi21, fi21, v14)
                    }
                    var v30 float64 = v14
                    _ = fi21
                    _ = v30
                }
                _ = v17
            }
            _ = v14
        }
    case "":
        var v31 int = 128
        _ = v31
    }
    var v32 uint8 = 42
    var v33 int16 = -799
    _ = v32
    _ = v33
    return 42
}

func main() {
    var v1 int8 = -28
    var v2 uint64 = 296
    var v3 int32 = 2147483647
    var v4 uint32 = 0
    v5 := make(map[uint32]bool)
    var v6 int = -2147483648
    v4 = v4
    var v34 float64 = 0.1
    for fi35 := 0; fi35 < 16; fi35++ {
        fmt.Println(v4, fi35, v4)
        delete(v5, v4)
        v34 += 3.14159
        _ = fi35
    }
    v3 += v3
    v3 = (v3 >> ((519) & 63))
    delete(v5, (v4 ^ (v4 ^ v4)))
    v5[(v4 | 2)] = true
    v6 = (^(v6))
    fmt.Println(v2)
    v36 := GenS3{F0: -9223372036854775808}
    v37 := GenS3{F0: 468}
    var v38 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
