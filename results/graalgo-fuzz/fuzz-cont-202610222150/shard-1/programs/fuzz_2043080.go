package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint32 }
type GenS2 struct { F0 string; F1 string }

func main() {
    var v1 uint16 = 32767
    var v2 string = "foo bar"
    var v3 uint16 = v1
    v4 := GenS1{F0: v2, F1: 9223372036854775807, F2: 393}
    var v5 float32 = 0.0
    v2 = v2
    if (0 > 4294967295) {
        var v6 uint = 843
        _ = v6
    } else {
        fmt.Println(v2)
        fmt.Println(v1, v5, v3)
        fmt.Println(v5, v3)
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int8 = 0
        if ((v5 / (0.1 - 3.14159)) != v5) {
            var v9 float32 = float32(0.1)
            v10 := []int{fi7}
            _ = v10
            _ = v9
        } else {
            v1 = (v3 << ((1) & 63))
            v8 = (v8 & 0)
        }
        v8 -= 1
        _ = fi7
        _ = v8
    }
    var v11 int64 = 42
    if true {
        v12 := GenS1{F0: v2, F1: 2147483647, F2: 395}
        for fi13 := 0; fi13 < 6; fi13++ {
            v11 -= v11
            _ = fi13
        }
        var v14 uint8 = 106
        _ = v12
        _ = v14
    }
    switch -231 {
    case -866:
        if (126 != 132) {
            v15 := GenS1{F0: v2, F1: 42, F2: 2}
            v5 = v5
            _ = v15
        } else {
            var v16 uint8 = 139
            for fi17 := 0; fi17 < 9; fi17++ {
                fmt.Println(v16, v16)
                _ = fi17
            }
            v1 = v3
            _ = v16
        }
    case -2147483648:
        for fi18 := 0; fi18 < 12; fi18++ {
            if false {
                var v19 int = fi18
                _ = v19
            } else {
                v20 := GenS2{F0: ((v2 + v2) + "a"), F1: (v2 + (("foo bar" + "foo bar") + v2))}
                v3 += (v3 << ((1) & 63))
                _ = v20
            }
            if (false || true) {
                if (!(!false)) {
                    var v21 int32 = -1
                    var v22 float32 = ((3.14159 * v5) + v5)
                    var v23 int32 = v21
                    _ = v21
                    _ = v22
                    _ = v23
                } else {
                    var v24 uint64 = 1
                    var v25 string = ""
                    switch (fi18 % (((fi18 * 127)) | 1)) {
                    case 9223372036854775807:
                        v24 = (v24 + v24)
                    case 32767:
                        v26 := []int{(fi18 << ((976) & 63)), fi18}
                        _ = v26
                    case 1:
                        v27 := []int{}
                        _ = v27
                    default:
                        if false {
                            for fi28 := 0; fi28 < 17; fi28++ {
                                v2 = ("\n\t" + (v25 + "hello"))
                                fmt.Println(v11)
                                v25 = ("" + (v25 + v25))
                                _ = fi28
                            }
                        } else {
                            fmt.Println(v25, v3, v11)
                            for fi29 := 0; fi29 < 2; fi29++ {
                                v2 = "hello"
                                v30 := []int{(fi18 << ((2) & 63)), (fi18 | int(v24)), -967}
                                _ = fi29
                                _ = v30
                            }
                            fmt.Println(v1)
                        }
                    }
                    _ = v24
                    _ = v25
                }
            } else {
                v3 = v1
            }
            var v31 int64 = v11
            _ = fi18
            _ = v31
        }
    }
    var v32 uint8 = 204
    var v33 uint16 = v3
    var v34 string = v2
    for fi35 := 0; fi35 < 8; fi35++ {
        for fi36 := 0; fi36 < 18; fi36++ {
            v5 += -2.25
            var v37 float32 = float32(1e10)
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
