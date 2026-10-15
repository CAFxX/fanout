package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 int16, p2 uint8) int32 {
    for fi4 := 0; fi4 < 20; fi4++ {
        switch "hello" {
        case "a":
            for fi5 := 0; fi5 < 20; fi5++ {
                var v6 bool = ((^(p2)) == p2)
                v7 := GenS1{F0: -142, F1: 262, F2: p2}
                _ = fi5
                _ = v6
                _ = v7
            }
        case "":
            p2 = 2
        default:
            for fi8 := 0; fi8 < 1; fi8++ {
                fmt.Println(p1, fi8, p0)
                switch "foo bar" {
                case "\n\t":
                    v9 := []int{p0, p0}
                    _ = v9
                case "foo bar":
                    var v10 float32 = 0.0
                    _ = v10
                case "hello":
                    var v11 int32 = -101
                    _ = v11
                default:
                    p2 = (p2 << ((18446744073709551615) & 63))
                    v12 := GenS1{F0: -2147483648, F1: 850, F2: 127}
                    fmt.Println(fi8, p1, p0)
                    v13 := []int{}
                    var v14 float32 = -2.25
                    var v15 float32 = (float32(1e10) - v14)
                    _ = len(v13)
                    p1 = p1
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi8
            }
        }
        var v16 uint8 = p2
        _ = fi4
        _ = v16
    }
    p2 = (p2 ^ 127)
    for fi17 := 0; fi17 < 2; fi17++ {
        p1 = p1
        _ = fi17
    }
    var v18 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return -2147483648
}

func main() {
    var v1 uint8 = 1
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]bool)
    v1 = (v1 & (v1 & (v1 + 0)))
    v19 := v3[-9223372036854775808]
    v2 = v2
    var v20 int32 = 2147483647
    switch "a" {
    case "foo bar":
        switch "" {
        case "hello":
            for fi21 := 0; fi21 < 19; fi21++ {
                v22 := GenS2{F0: 2, F1: (^(v2))}
                _ = fi21
                _ = v22
            }
        case "foo bar":
            v23 := []int{0}
            _ = v23
        case "":
            if (0 <= 626) {
                delete(v3, -779)
                fmt.Println(v2, v19, v1)
            }
        default:
            var v24 int64 = 9223372036854775807
            _ = v24
        }
    case "hello":
        for fi25 := 0; fi25 < 2; fi25++ {
            _ = len(v3)
            fmt.Println(v20, fi25, v19)
            _ = fi25
        }
    case "":
        var v26 uint64 = (v2 / (((v2 & v2)) | 1))
        _ = v26
    }
    switch 42 {
    case -1:
        if (-1 < 127) {
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 uint32 = 42
                for fi29 := 0; fi29 < 17; fi29++ {
                    if v19 {
                        v30 := GenS2{F0: 255, F1: (v2 ^ 186)}
                        _ = v30
                    } else {
                        delete(v3, (^(fi29)))
                        v19 = v19
                        v31 := GenS2{F0: 0, F1: v2}
                        v32 := GenS2{F0: 51, F1: 91}
                        _ = v31
                        _ = v32
                    }
                    if v19 {
                        var v33 int = (fi27 - -1)
                        fmt.Println(v19, v1, v33)
                        var v34 uint8 = (v1 ^ (v1 / ((v1) | 1)))
                        v35 := []int{}
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    _ = fi29
                }
                fmt.Println(v2)
                _ = fi27
                _ = v28
            }
        } else {
            v1 += v1
        }
    case 332:
        if true {
            fmt.Println(v19)
            v2 = v2
            for fi36 := 0; fi36 < 10; fi36++ {
                var v37 int8 = 64
                if true {
                    v38 := GenS3{F0: 4294967295, F1: (^(v1)), F2: 4294967295}
                    _ = v38
                }
                _ = fi36
                _ = v37
            }
        }
    case -2147483648:
        v39 := v3[-9223372036854775808]
        _ = v39
    }
    switch 1 {
    case 1:
        v40 := GenS4{F0: ("\n\t" + ("\n\t" + ""))}
        _ = v40
    case 32768:
        v41 := []int{4294967295, 533, 32767, 128}
        _ = v41
    case 435:
        for fi42 := 0; fi42 < 17; fi42++ {
            fmt.Println(v2)
            v19 = (-128 != 127)
            _ = fi42
        }
    default:
        v43 := GenS4{F0: ""}
        _ = v43
    }
    fmt.Println(v1, v20, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
