package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint16, p1 bool, p2 uint64) uint32 {
    v7 := GenS1{F0: 18446744073709551615, F1: 255, F2: -997}
    v8 := GenS1{F0: p2, F1: 11, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 655
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]int)
    var v3 uint8 = 114
    var v4 float64 = 0.0
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]string)
    var v9 uint = 67
    for fi10 := 0; fi10 < 4; fi10++ {
        v9 += 1
        fmt.Println(fi10, v9)
        if ((v1 * float32(0.1)) != v1) {
            if true {
                var v11 float32 = v1
                switch -2147483648 {
                case 127:
                    v12 := GenS2{F0: 65535}
                    _ = v12
                default:
                    v1 -= v1
                }
                _ = v11
            } else {
                if ((float32(v4) - float32((v4 + v4))) > v1) {
                    var v13 bool = false
                    v3 = v3
                    _ = v13
                }
                fmt.Println(fi10)
                v14 := []int{}
                _ = v14
            }
        } else {
            v15 := []int{0, 42, fi10}
            _ = v15
        }
        _ = fi10
    }
    if true {
        for fi16 := 0; fi16 < 9; fi16++ {
            v9 += v9
            v5[fi16] = 65535
            _ = fi16
        }
    }
    var v17 bool = (-1 < 42)
    v18 := GenS3{F0: 42, F1: v17}
    v5[127] = uint32(v3)
    fmt.Println(v9)
    v9 += v9
    v9 = 1
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(v1)
        _ = len(v2)
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := GenS4{F0: 0, F1: -1}
            switch fi19 {
            case 9223372036854775807:
                var v22 int64 = 0
                _ = v22
            }
            if (40 >= 542) {
                switch fi19 {
                case 32768:
                    switch "hello" {
                    case "a":
                        var v23 float32 = (3.14159 + -2.25)
                        _ = v23
                    case "foo bar":
                        switch "hello" {
                        case "a":
                            var v24 string = "a"
                            switch "\n\t" {
                            case "\n\t":
                                v9 += 0
                            case "hello":
                                var v25 int64 = -348
                                _ = v25
                            case "a":
                                v4 = v4
                            default:
                                var v26 string = (v24 + v24)
                                _ = v26
                            }
                            var v27 float32 = v1
                            _ = v24
                            _ = v27
                        case "foo bar":
                            fmt.Println(v9, fi20)
                        default:
                            var v28 float32 = 0.0
                            _ = v28
                        }
                    case "\n\t":
                        var v29 int = (fi19 >> ((v9) & 63))
                        _ = v29
                    default:
                        var v30 uint = v9
                        _ = v30
                    }
                case -533:
                    if (!(float64(v1) < v4)) {
                        var v31 uint32 = uint32(v3)
                        v32 := v2[fi19]
                        _ = v31
                        _ = v32
                    }
                case -529:
                    switch "\n\t" {
                    case "":
                        switch ("a" + "\n\t") {
                        case "\n\t":
                            v33 := v6[uint32(v3)]
                            _ = v33
                        default:
                            fmt.Println(fi19)
                        }
                    case "\n\t":
                        var v34 uint = 42
                        _ = v34
                    default:
                        var v35 int = fi20
                        _ = v35
                    }
                default:
                    v36 := v6[542]
                    _ = v36
                }
                delete(v2, 95)
            } else {
                fmt.Println(fi19, fi19, v3)
                v9 += 455
                var v37 bool = (v1 >= 3.14159)
                _ = v37
            }
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    var v38 uint8 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
