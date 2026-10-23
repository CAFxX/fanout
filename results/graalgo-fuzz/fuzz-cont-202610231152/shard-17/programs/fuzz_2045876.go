package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int; F2 uint16 }
type GenS5 struct { F0 string; F1 uint64; F2 int64 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 int; F1 uint8; F2 uint }
func genFn0(p0 bool) int16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            switch "" {
            case "hello":
                v6 := GenS1{F0: -2147483648}
                _ = v6
            default:
                p0 = (float32((float64(float32(1e10)) + (-2.25 + 1.5))) != (1e10 * -2.25))
            }
            if false {
                v7 := GenS2{F0: 1}
                v8 := GenS1{F0: 0}
                if p0 {
                    fmt.Println(p0)
                    var v9 int64 = 42
                    _ = v9
                }
                _ = v7
                _ = v8
            }
            _ = fi5
        }
        v10 := GenS3{F0: "hello"}
        v11 := GenS1{F0: 1}
        _ = fi4
        _ = v10
        _ = v11
    }
    p0 = p0
    switch 824 {
    case 128:
        v12 := GenS1{F0: 199}
        _ = v12
    case -403:
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                switch 687 {
                case 61:
                    v15 := []int{(fi13 % ((fi14) | 1)), fi13}
                    fmt.Println(fi14)
                    _ = v15
                }
                switch "foo bar" {
                case "\n\t":
                    p0 = (!true)
                    v16 := []int{(fi14 / (((fi14 - fi14)) | 1)), 32768, (fi14 * fi13), (fi13 / ((fi14) | 1))}
                    var v17 int16 = 32767
                    _ = v16
                    _ = v17
                case "":
                    var v18 int16 = 1000
                    _ = v18
                default:
                    v19 := []int{}
                    v20 := GenS4{F0: 0, F1: fi14, F2: 65535}
                    v21 := GenS4{F0: 9223372036854775807, F1: fi13, F2: 297}
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = fi14
            }
            p0 = p0
            v22 := GenS4{F0: 183, F1: fi13, F2: 255}
            _ = fi13
            _ = v22
        }
    case -1:
        if (!(!p0)) {
            p0 = false
            var v23 int16 = 228
            _ = v23
        } else {
            for fi24 := 0; fi24 < 6; fi24++ {
                var v25 uint8 = 0
                v26 := GenS5{F0: "foo bar", F1: 4294967295, F2: 42}
                _ = fi24
                _ = v25
                _ = v26
            }
            var v27 uint32 = 2147483648
            _ = v27
        }
    }
    _ = p0
    return -249
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 uint32 = 65535
    if true {
        v28 := GenS6{F0: 4294967295, F1: 179}
        v3 = (^(v3))
        var v29 int = -9223372036854775808
        _ = v28
        _ = v29
    }
    fmt.Println(v3)
    v3 = v3
    var v30 int32 = 37
    v30 = v30
    for fi31 := 0; fi31 < 6; fi31++ {
        fmt.Println(v3)
        for fi32 := 0; fi32 < 1; fi32++ {
            var v33 int64 = 9223372036854775807
            var v34 uint = 339
            _ = fi32
            _ = v33
            _ = v34
        }
        v35 := GenS7{F0: fi31, F1: 240, F2: 65535}
        _ = fi31
        _ = v35
    }
    fmt.Println(v30, v3)
    for fi36 := 0; fi36 < 11; fi36++ {
        v3 = v3
        if false {
            var v37 int8 = 100
            v30 = v30
            v37 -= (v37 % ((v37) | 1))
            _ = v37
        } else {
            _ = len(v2)
        }
        _ = fi36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
}
