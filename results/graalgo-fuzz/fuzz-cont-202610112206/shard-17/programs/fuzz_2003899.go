package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int16; F1 int32; F2 int }
type GenS3 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int32 = 1
    var v2 uint64 = 848
    var v3 float32 = 0.0
    v4 := make(map[uint32]bool)
    var v5 uint32 = 65535
    var v6 uint32 = uint32(v2)
    v7 := GenS1{F0: 0, F1: 558, F2: 0}
    v5 = v5
    switch "foo bar" {
    case "foo bar":
        if (v6 < v5) {
            var v8 uint = 65535
            if true {
                v9 := GenS1{F0: 2, F1: 2, F2: 0}
                _ = v9
            } else {
                v10 := GenS1{F0: 178, F1: 670, F2: 42}
                _ = v10
            }
            _ = v8
        } else {
            if true {
                switch -876 {
                case 202:
                    fmt.Println(v1, v6)
                case 0:
                    _ = len(v4)
                }
                var v11 uint = 255
                var v12 int8 = -1
                _ = v11
                _ = v12
            } else {
                fmt.Println(v2)
            }
            var v13 int32 = (v1 ^ (v1 * v1))
            switch 933 {
            case -814:
                v2 -= 18446744073709551615
            case 127:
                v1 = (v1 | v1)
            default:
                for fi14 := 0; fi14 < 17; fi14++ {
                    v3 -= (-2.25 * v3)
                    _ = fi14
                }
            }
            _ = v13
        }
    case "\n\t":
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS1{F0: 42, F1: 65535, F2: uint8(v2)}
            if ((v6 ^ (v5 ^ 0)) >= v5) {
                switch "hello" {
                case "\n\t":
                    var v17 float32 = v3
                    _ = v17
                case "foo bar":
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v5 -= 0
                        var v19 uint = 255
                        _ = fi18
                        _ = v19
                    }
                case "hello":
                    fmt.Println(v3, fi15)
                default:
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 bool = (!(!(0.0 > (0.1 / (3.14159 + 1.5)))))
                        _ = fi20
                        _ = v21
                    }
                }
            } else {
                switch "hello" {
                case "foo bar":
                    v22 := GenS1{F0: 42, F1: 42, F2: 70}
                    _ = v22
                default:
                    v2 = 2
                }
                for fi23 := 0; fi23 < 9; fi23++ {
                    v24 := GenS2{F0: -32768, F1: 2147483647, F2: 32768}
                    v25 := GenS3{F0: 25, F1: 226, F2: v1}
                    if false {
                        v6 = 0
                        delete(v4, (v5 ^ (v6 - 1)))
                    } else {
                        var v26 int8 = 0
                        _ = v26
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                switch fi15 {
                case 1:
                    v27 := GenS1{F0: uint16(v2), F1: 894, F2: 0}
                    _ = v27
                default:
                    var v28 bool = (478 > 373)
                    _ = v28
                }
            }
            _ = fi15
            _ = v16
        }
    }
    var v29 bool = false
    v30 := GenS4{F0: 4294967295, F1: 42}
    for fi31 := 0; fi31 < 12; fi31++ {
        switch fi31 {
        case -1:
            if (((0.1 * (((0.0 * v3) + float32(1.5)) + 1e10)) > ((v3 / float32(float64(float32(((3.14159 + 0.1) + 3.14159))))) + 0.1)) || v29) {
                fmt.Println(v1)
            } else {
                v3 = (((v3 + v3) - float32(1e10)) * (v3 / 1.5))
                switch ("foo bar" + "a") {
                case "hello":
                    fmt.Println(v5, fi31, v6)
                case "\n\t":
                    v32 := []int{(fi31 % (((fi31 * fi31)) | 1)), fi31}
                    _ = v32
                default:
                    v33 := v4[v5]
                    _ = v33
                }
                v4[v6] = v29
            }
        case 2147483647:
            fmt.Println(v2, v6)
        case -588:
            if ((v1 ^ (v1 ^ v1)) >= v1) {
                v34 := []int{fi31}
                _ = v34
            } else {
                v1 += (v1 | v1)
            }
        }
        _ = fi31
    }
    var v35 string = "\n\t"
    if (42 < 33) {
        var v36 int = -9223372036854775808
        v2 += 877
        v37 := GenS1{F0: 2, F1: 0, F2: 251}
        _ = v36
        _ = v37
    }
    var v38 uint8 = 15
    v39 := GenS5{F0: (v2 << ((42) & 63)), F1: -639}
    fmt.Println(v38, v29)
    if (v29 && v29) {
        fmt.Println(v1, v38, v2)
        v40 := GenS6{F0: (!(((902 == 0) || v29) || (v1 <= 1))), F1: 1}
        _ = v40
    } else {
        switch 127 {
        case -1:
            var v41 uint = 1
            _ = v41
        }
    }
    _ = len(v4)
    var v42 bool = v29
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
