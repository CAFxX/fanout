package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int16; F2 int64 }
func genFn0(p0 int16, p1 string) uint64 {
    p0 = p0
    switch p1 {
    case "hello":
        for fi5 := 0; fi5 < 15; fi5++ {
            p0 = p0
            for fi6 := 0; fi6 < 17; fi6++ {
                switch p1 {
                case "":
                    var v7 int32 = 0
                    _ = v7
                case "foo bar":
                    var v8 uint64 = 4294967295
                    _ = v8
                }
                v9 := []int{}
                _ = fi6
                _ = v9
            }
            p0 += p0
            _ = fi5
        }
    default:
        fmt.Println(p1, p0)
    }
    v10 := []int{2147483647, 41, 283, 32768}
    var v11 uint32 = 635
    v12 := GenS1{F0: p1, F1: (p0 << ((18446744073709551615) & 63)), F2: -1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 float64 = 3.14159
    v2 := []uint32{4294967295}
    var v3 uint64 = 18446744073709551615
    var v4 uint16 = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 int8 = 42
        if (42 >= 420) {
            if (!(1e10 <= (float32(v1) + 3.14159))) {
                if true {
                    var v15 uint = 255
                    switch (("hello" + "") + ("\n\t" + "\n\t")) {
                    case "a":
                        if false {
                            var v16 bool = true
                            var v17 uint = v15
                            _ = v16
                            _ = v17
                        } else {
                            if (true || (v4 != 2)) {
                                var v18 uint = v15
                                var v19 int8 = -128
                                _ = len(v2)
                                _ = v18
                                _ = v19
                            } else {
                                if (true || (false || ((v14 % (((v14 & (^(v14)))) | 1)) == (v14 | (^(v14)))))) {
                                    _ = len(v2)
                                    v15 -= 42
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                } else {
                                    v3 += v3
                                    var v20 float32 = 3.14159
                                    _ = v20
                                }
                            }
                            v2 = append(v2, 42)
                        }
                    default:
                        v3 += 1
                    }
                    _ = v15
                }
                v14 -= 0
                v21 := GenS1{F0: ("foo bar" + "foo bar"), F1: -32768, F2: -9223372036854775808}
                _ = v21
            } else {
                v2 = append(v2, 2147483648)
            }
            v22 := GenS2{F0: uint32(v4)}
            var v23 int64 = 0
            _ = v22
            _ = v23
        } else {
            v24 := GenS1{F0: (("foo bar" + "a") + "\n\t"), F1: 1, F2: 1}
            _ = v24
        }
        _ = fi13
        _ = v14
    }
    var v25 int64 = 0
    var v26 int8 = -97
    v26 = 100
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        v27 := GenS3{F0: 2, F1: -32768, F2: (v25 + (-(v25)))}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v28 := GenS1{F0: (("" + ("a" + "foo bar")) + "foo bar"), F1: int16(v26), F2: (v25 | v25)}
        _ = v27
        _ = v28
    } else {
        for fi29 := 0; fi29 < 8; fi29++ {
            for fi30 := 0; fi30 < 6; fi30++ {
                v26 -= v26
                var v31 uint64 = genFn0(-32768, (("" + ((("" + ("a" + "hello")) + ("foo bar" + ("foo bar" + "\n\t"))) + "")) + ""))
                v26 = 127
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
        fmt.Println(v1, v4, v1)
        var v32 uint = 4294967295
        _ = v32
    }
    switch 2147483647 {
    case 0:
        v4 -= 65535
    case 32768:
        v1 = v1
    case -9223372036854775808:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v33 uint = 0
    for fi34 := 0; fi34 < 1; fi34++ {
        for fi35 := 0; fi35 < 18; fi35++ {
            var v36 bool = true
            _ = fi35
            _ = v36
        }
        if true {
            var v37 string = ("foo bar" + ("\n\t" + "foo bar"))
            v38 := GenS2{F0: 587}
            _ = v37
            _ = v38
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
}
