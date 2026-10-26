package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool) int64 {
    fmt.Println(p0)
    switch 32767 {
    case 9223372036854775807:
        v5 := []int{32767}
        _ = v5
    case 128:
        var v6 int8 = 0
        _ = v6
    case 0:
        p0 = true
    default:
        p0 = p0
    }
    fmt.Println(p0)
    _ = p0
    return 345
}

func main() {
    v1 := []uint32{2147483647, 65535, 42}
    v2 := make(map[string]bool)
    var v3 uint16 = 65535
    var v4 int = -855
    var v7 uint64 = 1
    for fi8 := 0; fi8 < 6; fi8++ {
        v4 += v4
        _ = fi8
    }
    v9 := v2[""]
    for fi10 := 0; fi10 < 16; fi10++ {
        if v9 {
            v2["hello"] = (v9 && v9)
            v1 = append(v1, 0)
        }
        v11 := GenS1{F0: -139, F1: 353}
        fmt.Println(v9, v4)
        _ = fi10
        _ = v11
    }
    v9 = false
    var v12 int16 = -32768
    v13 := GenS1{F0: genFn0(v9), F1: 4294967295}
    v14 := GenS2{F0: 2147483647, F1: -234}
    var v15 uint32 = 392
    v1 = append(v1, 42)
    v16 := GenS1{F0: 612, F1: 0}
    v17 := GenS1{F0: -1, F1: uint(v15)}
    delete(v2, "")
    var v18 uint = 18446744073709551615
    var v19 int8 = -74
    if false {
        if (v7 > (^(v7))) {
            if ((v4 ^ 1) != v4) {
                v7 = (v7 % ((v7) | 1))
                if v9 {
                    fmt.Println(v3)
                    v20 := GenS1{F0: -725, F1: v18}
                    v3 += v3
                    _ = v20
                } else {
                    v3 -= uint16(v15)
                    if v9 {
                        if (!v9) {
                            for fi21 := 0; fi21 < 14; fi21++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v22 := GenS2{F0: -795, F1: 210}
                                var v23 float32 = 1e10
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            for fi24 := 0; fi24 < 3; fi24++ {
                                var v25 int64 = 1
                                if false {
                                    var v26 uint8 = 255
                                    var v27 int = v4
                                    if false {
                                        if (3.14159 != (-2.25 + 3.14159)) {
                                            fmt.Println(v27, v3)
                                            var v28 uint8 = (v26 & v26)
                                            v12 = (v12 & (v12 - -1))
                                            _ = v28
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v2["\n\t"] = (-2.25 >= (((1e10 * (((0.1 - (-2.25 - 1.5)) * -2.25) + 1e10)) - -2.25) / 3.14159))
                                        }
                                        var v29 int = v27
                                        var v30 uint16 = (^(v3))
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v31 int16 = (v12 << ((4294967295) & 63))
                                    fmt.Println(v25, v18, v12)
                                    _ = v31
                                }
                                _ = fi24
                                _ = v25
                            }
                            for fi32 := 0; fi32 < 17; fi32++ {
                                for fi33 := 0; fi33 < 18; fi33++ {
                                    v2["foo bar"] = (int(v15) == int(v18))
                                    _ = fi33
                                }
                                _ = fi32
                            }
                        }
                    }
                }
            } else {
                v34 := GenS3{F0: (v3 & 2)}
                _ = v34
            }
        }
    } else {
        v35 := v2[("hello" + "")]
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
