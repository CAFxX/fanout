package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint32) string {
    p0 -= p0
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(fi6)
        v7 := []int{fi6, fi6}
        _ = fi6
        _ = v7
    }
    _ = p0
    return "hello"
}

func main() {
    var v1 uint = 1
    var v2 uint8 = 0
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v5 := make(map[uint32]uint32)
    var v8 float32 = (1.5 / 1e10)
    v9 := GenS1{F0: 127, F1: 255}
    var v10 int32 = 1
    if false {
        var v11 uint32 = 2147483648
        _ = v11
    } else {
        fmt.Println(v8)
        var v12 string = "\n\t"
        _ = v12
    }
    fmt.Println(v1)
    if true {
        v1 = v1
        if (!(v1 >= (v1 << ((69) & 63)))) {
            var v13 int8 = 100
            var v14 int8 = v13
            switch "\n\t" {
            case "a":
                for fi15 := 0; fi15 < 2; fi15++ {
                    v1 = 0
                    _ = fi15
                }
            default:
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 int = (fi16 ^ 42)
                    v8 += v8
                    delete(v4, (v17 & (^(v17))))
                    _ = fi16
                    _ = v17
                }
            }
            _ = v13
            _ = v14
        }
    }
    v10 += v10
    if true {
        var v18 string = "foo bar"
        var v19 uint64 = 459
        _ = v18
        _ = v19
    } else {
        var v20 uint32 = uint32(v2)
        for fi21 := 0; fi21 < 8; fi21++ {
            v2 = (v2 | v2)
            if true {
                v2 = v2
                v2 += v2
            } else {
                for fi22 := 0; fi22 < 19; fi22++ {
                    if false {
                        var v23 int8 = 0
                        v24 := GenS1{F0: fi22, F1: 1}
                        if false {
                            fmt.Println(v8, v8, fi22)
                            if false {
                                v25 := GenS1{F0: 32768, F1: 255}
                                v20 += (v20 | v20)
                                _ = v25
                            }
                        } else {
                            switch "a" {
                            case "\n\t":
                                v26 := GenS2{F0: 65535, F1: uint16(v1)}
                                _ = v26
                            case "hello":
                                var v27 uint16 = 32767
                                _ = v27
                            default:
                                v28 := GenS3{F0: 32767, F1: v2}
                                _ = v28
                            }
                            v29 := GenS4{F0: v10}
                            _ = v29
                        }
                        _ = v23
                        _ = v24
                    } else {
                        v2 += (v2 ^ 127)
                        var v30 string = "foo bar"
                        _ = v30
                    }
                    _ = fi22
                }
            }
            if false {
                for fi31 := 0; fi31 < 12; fi31++ {
                    switch ((("foo bar" + ("\n\t" + ("hello" + "hello"))) + (("foo bar" + "foo bar") + "foo bar")) + ("a" + "\n\t")) {
                    case "a":
                        var v32 float64 = float64((v8 + (v8 - float32(0.0))))
                        _ = v32
                    default:
                        if (uint64(v2) > 0) {
                            var v33 uint16 = 65535
                            v1 = v1
                            fmt.Println(fi31, v20, v1)
                            _ = v33
                        }
                    }
                    _ = fi31
                }
                for fi34 := 0; fi34 < 4; fi34++ {
                    v35 := []int{-400, fi34}
                    v1 = (v1 << ((v1) & 63))
                    var v36 int32 = int32(v2)
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            } else {
                var v37 float32 = v8
                fmt.Println(v2, v10)
                _ = v37
            }
            _ = fi21
        }
        fmt.Println(v20, v10)
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
