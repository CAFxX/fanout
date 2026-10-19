package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int }
func genFn0() int16 {
    var v5 uint8 = 226
    fmt.Println(v5)
    var v6 uint32 = 192
    v5 += (v5 << ((967) & 63))
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 bool = (42 < 42)
    var v3 float32 = 1.5
    v4 := make(map[string]int)
    v7 := GenS1{F0: 157}
    fmt.Println(v3, v2, v2)
    var v8 uint8 = 42
    switch 4294967295 {
    case 714:
        v4["hello"] = 4294967295
    case 128:
        v9 := GenS2{F0: "\n\t", F1: v2, F2: 0}
        _ = v9
    default:
        var v10 uint32 = 65535
        _ = v10
    }
    switch 992 {
    case 353:
        v1 += v1
    }
    v3 -= v3
    v11 := []int{}
    var v12 int8 = 75
    switch -448 {
    case -1:
        if true {
            if (2 == 84) {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v11 = append(v11, fi13)
                    for fi14 := 0; fi14 < 12; fi14++ {
                        var v15 float64 = 3.14159
                        v3 = float32(v15)
                        _ = fi14
                        _ = v15
                    }
                    v16 := GenS1{F0: 1}
                    _ = fi13
                    _ = v16
                }
                if (!(genFn0() == int16(v12))) {
                    v8 += v8
                } else {
                    if v2 {
                        _ = len(v11)
                        v11 = append(v11, 9223372036854775807)
                        for fi17 := 0; fi17 < 7; fi17++ {
                            v18 := GenS3{F0: 2147483647, F1: 941}
                            _ = fi17
                            _ = v18
                        }
                    } else {
                        switch 894 {
                        case 32768:
                            fmt.Println(v12)
                        case 42:
                            v8 += (v8 << ((1) & 63))
                        case 535:
                            fmt.Println(v1, v3, v8)
                        default:
                            var v19 int = -903
                            _ = v19
                        }
                        var v20 int8 = 105
                        _ = v20
                    }
                }
                v21 := GenS4{F0: 2147483647}
                _ = v21
            } else {
                if v2 {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        _ = len(v4)
                        _ = fi22
                    }
                    var v23 int64 = 42
                    _ = v23
                }
            }
            if v2 {
                v24 := v4["hello"]
                v25 := v4["\n\t"]
                v24 = v25
                _ = v24
                _ = v25
            } else {
                if (v2 && v2) {
                    if v2 {
                        var v26 int8 = (v12 - (v12 + (^(v12))))
                        switch "foo bar" {
                        case "a":
                            v8 = (^(v8))
                        }
                        v8 += (^(v8))
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 6; fi27++ {
                            _ = len(v4)
                            v3 = (v3 * -2.25)
                            v11 = append(v11, fi27)
                            _ = fi27
                        }
                    }
                }
                v28 := GenS4{F0: 61}
                var v29 int64 = 0
                _ = v28
                _ = v29
            }
        }
    case 42:
        v2 = v2
    case -276:
        var v30 int = 32767
        _ = v30
    }
    v1 = -32768
    _ = len(v4)
    var v31 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v7
    _ = v8
}
