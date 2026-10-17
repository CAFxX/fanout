package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int32; F1 int; F2 bool }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint8, p1 string, p2 uint64) int16 {
    var v6 int8 = 1
    switch 42 {
    case 693:
        fmt.Println(p2)
    default:
        v7 := GenS1{F0: 32767, F1: false}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"foo bar"}
    v3 := []string{"a", "foo bar", "\n\t", ""}
    v4 := make(map[string]uint32)
    v5 := []int{4294967295}
    var v8 uint64 = 0
    v9 := GenS2{F0: -1}
    var v10 uint32 = 188
    var v11 uint8 = 160
    v8 = v8
    v10 += v10
    switch 127 {
    case -9223372036854775808:
        for fi12 := 0; fi12 < 15; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                _ = len(v2)
                v11 = 225
                _ = fi13
            }
            _ = fi12
        }
    case -954:
        var v14 int32 = 720
        _ = v14
    case 9223372036854775807:
        for fi15 := 0; fi15 < 5; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                if (true && (1 < 262)) {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    if false {
                        v17 := GenS3{F0: fi16, F1: true}
                        _ = v17
                    } else {
                        if false {
                            delete(v1, "a")
                            _ = len(v3)
                            fmt.Println(v10)
                        } else {
                            fmt.Println(fi16)
                            v18 := GenS4{F0: 2147483647, F1: (-(fi15)), F2: (!true)}
                            _ = v18
                        }
                        v10 = 4294967295
                    }
                } else {
                    if true {
                        _ = len(v2)
                    } else {
                        var v19 int16 = 1
                        _ = v19
                    }
                }
                v20 := GenS4{F0: 518, F1: (fi15 / ((-494) | 1)), F2: ((v8 % (((v8 + (^(v8)))) | 1)) <= (v8 % (((v8 % ((v8) | 1))) | 1)))}
                if (("\n\t" == "") || true) {
                    fmt.Println(v11, fi15, v8)
                } else {
                    var v21 int8 = 0
                    _ = v21
                }
                _ = fi16
                _ = v20
            }
            var v22 int32 = -1
            _ = fi15
            _ = v22
        }
    }
    fmt.Println(v8, v8, v11)
    v23 := GenS5{F0: genFn0(v11, "foo bar", v8)}
    v24 := GenS2{F0: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
