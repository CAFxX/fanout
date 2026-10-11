package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int64; F1 int; F2 int64 }
type GenS3 struct { F0 int8; F1 int32; F2 int64 }
type GenS4 struct { F0 int64; F1 uint16; F2 int32 }
func genFn0(p0 int8, p1 uint64, p2 uint) string {
    v5 := []int{0, int(p1)}
    p0 -= (p0 << ((p2) & 63))
    p0 = -128
    switch (("a" + "\n\t") + ("hello" + "hello")) {
    case "":
        if true {
            for fi6 := 0; fi6 < 7; fi6++ {
                if false {
                    p0 = (p0 & 85)
                } else {
                    var v7 uint16 = 32767
                    if len(v5) > 0 {
                        v5[0] = (fi6 - (fi6 % (((fi6 << (((p2 << ((0) & 63))) & 63))) | 1)))
                    }
                    _ = v7
                }
                _ = len(v5)
                if len(v5) > 0 {
                    v5[0] = (fi6 >> ((0) & 63))
                }
                _ = fi6
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (127 != 0) {
                var v8 int64 = -1
                for fi9 := 0; fi9 < 19; fi9++ {
                    v10 := GenS1{F0: 1, F1: 994}
                    if true {
                        switch "a" {
                        case "hello":
                            var v11 string = ""
                            _ = v11
                        }
                    } else {
                        if false {
                            var v12 int8 = 125
                            _ = v12
                        } else {
                            fmt.Println(p0)
                            v5 = append(v5, fi9)
                            var v13 int64 = v8
                            _ = v13
                        }
                        v14 := GenS1{F0: -9223372036854775808, F1: -1}
                        _ = v14
                    }
                    p2 -= (p2 ^ p2)
                    _ = fi9
                    _ = v10
                }
                if false {
                    v15 := GenS1{F0: (v8 % ((v8) | 1)), F1: 0}
                    v8 -= (v8 - v8)
                    _ = v15
                }
                _ = v8
            } else {
                v16 := GenS2{F0: -9223372036854775808, F1: -506, F2: 1}
                p1 -= 42
                if false {
                    p0 = p0
                    v17 := GenS3{F0: (-(p0)), F1: 42, F2: 0}
                    var v18 float64 = 3.14159
                    _ = v17
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        p0 = p0
                        _ = fi19
                    }
                    var v20 int64 = 0
                    _ = v20
                }
                _ = v16
            }
        }
    case "foo bar":
        var v21 string = ""
        _ = v21
    case "\n\t":
        v22 := GenS4{F0: -9223372036854775808, F1: 32767, F2: 42}
        _ = v22
    default:
        if (!((p0 << (((p2 + p2)) & 63)) != (p0 & p0))) {
            var v23 uint = uint(p1)
            p1 = (p1 % (((p1 * p1)) | 1))
            _ = v23
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                _ = len(v5)
                v25 := GenS1{F0: 0, F1: -1}
                _ = fi24
                _ = v25
            }
            var v26 int16 = -32768
            _ = v26
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ("hello" + ("hello" + ""))
}
func genFn1(p0 bool) uint16 {
    v27 := GenS2{F0: -17, F1: -66, F2: -620}
    var v28 bool = (p0 || false)
    _ = p0
    _ = v27
    _ = v28
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1.5
    var v3 int64 = 42
    var v4 float64 = float64(v2)
    v3 = (v3 >> ((4294967295) & 63))
    var v29 int16 = 32767
    var v30 string = (("\n\t" + "a") + "a")
    v31 := []int{32768, -1}
    v32 := GenS1{F0: v3, F1: 832}
    v33 := GenS2{F0: v3, F1: 32767, F2: v3}
    var v34 int32 = 42
    for fi35 := 0; fi35 < 4; fi35++ {
        if (2147483647 < 2147483647) {
            v31 = append(v31, -324)
            fmt.Println(v2, v34)
        } else {
            v2 += float32((1e10 + v4))
            v29 = (v29 ^ -32768)
            for fi36 := 0; fi36 < 18; fi36++ {
                var v37 uint16 = 32767
                _ = fi36
                _ = v37
            }
        }
        v2 += v2
        _ = fi35
    }
    fmt.Println(v29, v2, v4)
    for fi38 := 0; fi38 < 14; fi38++ {
        v39 := GenS3{F0: -1, F1: (v34 ^ (v34 * v34)), F2: (^(v3))}
        _ = fi38
        _ = v39
    }
    var v40 int32 = (v34 << ((65535) & 63))
    var v41 int32 = v40
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v41
}
