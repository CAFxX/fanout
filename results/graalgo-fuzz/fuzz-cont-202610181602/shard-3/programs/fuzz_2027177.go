package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64, p1 int8) bool {
    v6 := GenS1{F0: 9223372036854775807, F1: 42}
    var v7 float64 = 1.5
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return true
}

func main() {
    var v1 string = "a"
    var v2 uint16 = 0
    v3 := []uint32{806, 4294967295}
    v4 := make(map[uint32]int)
    v5 := []int{}
    var v8 uint64 = 0
    v9 := GenS1{F0: 0, F1: 350}
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint = 255
        fmt.Println(fi10, fi10)
        _ = fi10
        _ = v11
    }
    switch -1 {
    case 693:
        if false {
            v1 = "\n\t"
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case -1:
        v12 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
        _ = v12
    }
    if false {
        v13 := GenS2{F0: (^(v8))}
        v1 = (((v1 + "a") + v1) + "hello")
        var v14 int16 = 1
        _ = v13
        _ = v14
    }
    var v15 int64 = -279
    var v16 bool = false
    v4[1] = 1
    switch 42 {
    case -42:
        if false {
            if (2147483648 > 2147483647) {
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 float64 = (0.1 * 0.1)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    for fi19 := 0; fi19 < 13; fi19++ {
                        v20 := GenS1{F0: -205, F1: v15}
                        var v21 float64 = (v18 - 1.5)
                        if len(v3) > 0 {
                            v3[0] = 775
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = fi17
                    _ = v18
                }
                fmt.Println(v1, v15, v1)
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if true {
                    var v22 float64 = 1e10
                    var v23 bool = (255 < 65535)
                    if (!v23) {
                        for fi24 := 0; fi24 < 12; fi24++ {
                            if v23 {
                                var v25 int32 = 2147483647
                                var v26 uint8 = 2
                                v27 := GenS2{F0: v8}
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            _ = fi24
                        }
                    }
                    _ = v22
                    _ = v23
                } else {
                    if true {
                        var v28 string = "hello"
                        if ("a" != v1) {
                            var v29 int32 = 0
                            for fi30 := 0; fi30 < 13; fi30++ {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v31 := v4[4294967295]
                                v28 = "foo bar"
                                _ = fi30
                                _ = v31
                            }
                            _ = v29
                        } else {
                            fmt.Println(v2, v2)
                        }
                        if (!v16) {
                            var v32 int16 = -177
                            _ = v32
                        } else {
                            if true {
                                v2 = 42
                            }
                            v28 = (v1 + v28)
                        }
                        _ = v28
                    }
                    var v33 int = -1
                    v4[128] = v33
                    _ = v33
                }
                switch 4294967295 {
                case 32768:
                    fmt.Println(v2, v2)
                case 128:
                    fmt.Println(v16, v2)
                case 2147483647:
                    var v34 int = 42
                    _ = v34
                default:
                    var v35 int8 = 127
                    _ = v35
                }
            }
            v2 += v2
        }
    default:
        var v36 float64 = (-2.25 * -2.25)
        _ = v36
    }
    var v37 int8 = -114
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
