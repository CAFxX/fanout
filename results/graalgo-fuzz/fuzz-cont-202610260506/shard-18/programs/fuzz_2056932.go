package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int8 }
type GenS2 struct { F0 string; F1 bool; F2 uint }
type GenS3 struct { F0 string; F1 uint; F2 uint }
type GenS4 struct { F0 int; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint8; F1 int32; F2 int }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 string; F1 uint64 }
func genFn0(p0 int16, p1 int32) string {
    switch ("hello" + "foo bar") {
    case "":
        if ((false && (1.5 <= 0.1)) || true) {
            var v4 float32 = float32(1.5)
            v5 := GenS1{F0: 9223372036854775807, F1: 65535, F2: 63}
            _ = v4
            _ = v5
        } else {
            var v6 float32 = 0.1
            var v7 int8 = 8
            var v8 bool = (4294967295 >= int(p1))
            _ = v6
            _ = v7
            _ = v8
        }
    case "foo bar":
        v9 := GenS2{F0: "", F1: false, F2: 1}
        _ = v9
    default:
        v10 := GenS2{F0: "\n\t", F1: (p0 != p0), F2: 0}
        _ = v10
    }
    v11 := []int{}
    v12 := GenS3{F0: ("" + ""), F1: 255, F2: 65535}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return ("a" + "foo bar")
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 0
    var v3 uint16 = uint16(v2)
    fmt.Println(v2, v3)
    fmt.Println(v3, v2)
    if (-1 == -32768) {
        v3 = v3
        if false {
            switch -1 {
            case 2147483647:
                v13 := GenS4{F0: 291, F1: 2147483647, F2: v3}
                _ = v13
            case 128:
                var v14 int = 927
                _ = v14
            }
            switch "foo bar" {
            case "":
                for fi15 := 0; fi15 < 2; fi15++ {
                    if (true || (42 != 42)) {
                        var v16 uint8 = 2
                        v17 := v1[128]
                        v18 := GenS5{F0: v16, F1: -1, F2: fi15}
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 15; fi19++ {
                        if ((32767 >= 42) || (!(true && (!false)))) {
                            v3 += (^(v3))
                            var v20 int8 = 100
                            fmt.Println(fi19, v20, v20)
                            v3 = (v3 - 14)
                            if (!((898 == -990) && true)) {
                                v21 := GenS6{F0: v3, F1: ""}
                                _ = v21
                            }
                            switch (^(fi19)) {
                            case 32767:
                                var v22 int8 = v20
                                _ = v22
                            case 141:
                                fmt.Println(fi19, v20, v3)
                            default:
                                var v23 uint64 = 849
                                _ = v23
                            }
                            v24 := GenS7{F0: "foo bar", F1: 748}
                            _ = v20
                            _ = v24
                        } else {
                            v2 = 51
                        }
                        v3 = v3
                        _ = fi19
                    }
                    var v25 int32 = -2147483648
                    _ = fi15
                    _ = v25
                }
            default:
                v1[127] = ("hello" + "")
            }
            v26 := v1[0]
            _ = v26
        }
        fmt.Println(v2, v2)
    } else {
        v27 := []int{32768}
        var v28 uint16 = v3
        if false {
            switch 156 {
            case -517:
                v28 = v28
            case 1:
                v2 = 2147483648
            default:
                v29 := GenS3{F0: "\n\t", F1: 177, F2: 726}
                _ = v29
            }
            v30 := GenS6{F0: v3, F1: ("hello" + "hello")}
            v28 += v28
            _ = v30
        } else {
            v28 += (v3 - v28)
        }
        _ = v27
        _ = v28
    }
    var v31 bool = false
    v1[9223372036854775807] = ("hello" + "hello")
    var v32 uint16 = v3
    fmt.Println(v32, v31, v31)
    if (-126 >= 1) {
        fmt.Println(v31, v2)
        var v33 uint64 = 1
        _ = v33
    }
    v34 := GenS6{F0: (v32 >> ((42) & 63)), F1: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
}
