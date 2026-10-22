package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint8 }
func genFn0() uint {
    var v6 uint32 = 65535
    v6 = (v6 >> ((561) & 63))
    fmt.Println(v6)
    _ = v6
    return 2
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 0
    var v3 uint16 = 2
    var v4 uint64 = 2
    var v5 uint8 = 42
    var v7 int8 = -1
    delete(v1, "hello")
    v8 := GenS1{F0: 0, F1: (^(v4))}
    for fi9 := 0; fi9 < 6; fi9++ {
        if true {
            v10 := GenS2{F0: ""}
            v5 -= (v5 << ((255) & 63))
            _ = v10
        } else {
            switch fi9 {
            case 2147483647:
                switch ("foo bar" + "hello") {
                case "hello":
                    var v11 string = "\n\t"
                    _ = v11
                default:
                    v7 += v7
                }
            case 81:
                switch -289 {
                case 32767:
                    if (!((fi9 | fi9) != fi9)) {
                        v3 += (v3 % ((65535) | 1))
                    }
                case 402:
                    var v12 uint = genFn0()
                    _ = v12
                }
            case 32768:
                v4 += 4294967295
            }
            v7 = (^(v7))
            fmt.Println(v7, v7, fi9)
        }
        fmt.Println(v7)
        v2 = (v2 % ((v2) | 1))
        _ = fi9
    }
    var v13 string = (("foo bar" + "\n\t") + ("foo bar" + "foo bar"))
    if true {
        for fi14 := 0; fi14 < 15; fi14++ {
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 bool = true
                v1[v13] = fi15
                _ = fi15
                _ = v16
            }
            v17 := GenS3{F0: v4, F1: v3, F2: v3}
            var v18 float64 = 0.1
            _ = fi14
            _ = v17
            _ = v18
        }
    }
    if (704 >= 828) {
        v3 = (v3 | 2)
        var v19 float32 = (0.0 - (1e10 + (-2.25 * 0.1)))
        _ = v19
    }
    if (0 >= 0) {
        v20 := GenS4{F0: false, F1: (v4 * v4), F2: 2147483647}
        _ = v20
    } else {
        v21 := GenS5{F0: 0, F1: 4294967295}
        _ = v21
    }
    switch 42 {
    case 42:
        switch 32768 {
        case -1:
            v13 = (v13 + "\n\t")
        case -656:
            v4 += v4
        case 127:
            for fi22 := 0; fi22 < 3; fi22++ {
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 int64 = 9223372036854775807
                    if (0 != v3) {
                        switch "hello" {
                        case "hello":
                            var v25 uint32 = 720
                            _ = v25
                        }
                        v24 += v24
                        var v26 int = 2147483647
                        _ = v26
                    } else {
                        v24 += v24
                        var v27 uint32 = 833
                        _ = v27
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
        default:
            switch 2147483647 {
            case 32767:
                v28 := GenS1{F0: 65535, F1: v4}
                _ = v28
            case -219:
                var v29 uint64 = (v4 << ((genFn0()) & 63))
                _ = v29
            case -9223372036854775808:
                v30 := v1[v13]
                _ = v30
            default:
                v7 -= (^(v7))
            }
        }
    default:
        v31 := GenS6{F0: 2, F1: true, F2: (^(v4))}
        _ = v31
    }
    v32 := v1[((v13 + v13) + v13)]
    for fi33 := 0; fi33 < 3; fi33++ {
        for fi34 := 0; fi34 < 8; fi34++ {
            var v35 float32 = (((1.5 * 0.0) * 3.14159) - (((0.1 / float32(0.1)) + (0.1 / 1.0)) / 0.1))
            switch v13 {
            case "a":
                var v36 uint64 = (v4 % ((v4) | 1))
                _ = v36
            case "foo bar":
                var v37 uint8 = v5
                _ = v37
            case "hello":
                var v38 uint = uint(v4)
                _ = v38
            default:
                v39 := GenS7{F0: (v5 + v5)}
                _ = v39
            }
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    v3 = 1
    if ((v4 | (v4 | 2)) != (^(v4))) {
        v40 := []int{}
        _ = v40
    } else {
        for fi41 := 0; fi41 < 20; fi41++ {
            v32 -= 2147483647
            fmt.Println(v2)
            var v42 int64 = 0
            _ = fi41
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 14; fi43++ {
        if (!(!true)) {
            v44 := []int{v32, int(v2), v32}
            _ = v44
        }
        var v45 uint8 = 128
        var v46 int8 = -49
        _ = fi43
        _ = v45
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
