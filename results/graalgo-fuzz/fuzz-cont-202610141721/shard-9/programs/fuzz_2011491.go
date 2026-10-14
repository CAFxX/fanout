package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint8) int16 {
    switch "hello" {
    case "a":
        p0 = p0
    case "":
        switch "hello" {
        case "foo bar":
            v5 := GenS1{F0: 0, F1: 9223372036854775807}
            _ = v5
        case "\n\t":
            v6 := GenS2{F0: 100, F1: 98}
            _ = v6
        default:
            var v7 float32 = ((3.14159 + (3.14159 / 3.14159)) / 3.14159)
            _ = v7
        }
    case "hello":
        p0 += p0
    default:
        var v8 int = -895
        _ = v8
    }
    var v9 int8 = 42
    for fi10 := 0; fi10 < 6; fi10++ {
        if (!true) {
            p0 = (p0 - p0)
            switch "hello" {
            case "foo bar":
                v11 := []int{(fi10 * int(p0)), (fi10 + fi10)}
                _ = v11
            case "":
                var v12 string = ("foo bar" + ("hello" + "foo bar"))
                _ = v12
            default:
                v13 := []int{}
                _ = v13
            }
            v9 = (v9 * (v9 | (v9 | -1)))
        } else {
            if true {
                var v14 int32 = 725
                _ = v14
            } else {
                if true {
                    switch (("hello" + "\n\t") + "") {
                    case "foo bar":
                        v15 := []int{2147483647, fi10}
                        _ = v15
                    case "":
                        if true {
                            fmt.Println(fi10)
                            var v16 int8 = (^(v9))
                            _ = v16
                        } else {
                            switch fi10 {
                            case -56:
                                v17 := GenS3{F0: 4294967295, F1: 42}
                                _ = v17
                            case -889:
                                var v18 int16 = 42
                                _ = v18
                            default:
                                switch -954 {
                                case 32767:
                                    p0 = p0
                                default:
                                    if (false || (-1 < -866)) {
                                        v9 = (^(v9))
                                    } else {
                                        for fi19 := 0; fi19 < 6; fi19++ {
                                            v9 = 127
                                            _ = fi19
                                        }
                                        v20 := GenS3{F0: 2, F1: 255}
                                        _ = v20
                                    }
                                }
                            }
                            v21 := []int{}
                            _ = v21
                        }
                    }
                }
                v22 := GenS2{F0: -128, F1: 42}
                v23 := GenS2{F0: (-(v9)), F1: 775}
                _ = v22
                _ = v23
            }
            p0 = p0
            var v24 float32 = 0.1
            _ = v24
        }
        _ = fi10
    }
    _ = p0
    _ = v9
    return 975
}
func genFn1(p0 string, p1 uint64) uint8 {
    if false {
        if false {
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 float32 = 1e10
                p1 = p1
                _ = fi25
                _ = v26
            }
            v27 := []int{128, 127, 128, 600}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 10; fi28++ {
                p0 = p0
                _ = fi28
            }
            var v29 int = 545
            _ = v29
        }
    }
    p1 -= (p1 % (((p1 ^ 4294967295)) | 1))
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 2
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    v2 = v2
    v30 := GenS2{F0: -4, F1: 460}
    var v31 int8 = 35
    var v32 uint = 112
    v33 := GenS4{F0: -1, F1: -1}
    v32 -= 65535
    v34 := GenS5{F0: (-(v31)), F1: 42}
    _ = len(v1)
    switch "" {
    case "foo bar":
        fmt.Println(v2)
    case "hello":
        for fi35 := 0; fi35 < 1; fi35++ {
            v36 := v4[128]
            var v37 uint32 = 2147483647
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    var v38 int64 = 0
    v39 := GenS2{F0: v31, F1: 2}
    if ((v2 < 127) || true) {
        v38 = (v38 << ((0) & 63))
        if (((v2 >> ((v32) & 63)) <= 42) && false) {
            var v40 int = -144
            _ = v40
        }
    }
    v41 := GenS3{F0: 2147483647, F1: 45}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v41
}
