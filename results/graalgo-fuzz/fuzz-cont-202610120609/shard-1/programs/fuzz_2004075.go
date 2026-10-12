package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 bool; F2 string }

func main() {
    v1 := []uint32{2147483647, 337}
    var v2 uint8 = 165
    var v3 float64 = ((-2.25 + 3.14159) / 1.0)
    v4 := make(map[string]uint32)
    var v5 int64 = -379
    v6 := []string{(("" + "hello") + ""), ("hello" + ""), "a"}
    var v7 uint16 = 1
    v8 := v4["hello"]
    var v9 bool = true
    var v10 float32 = -2.25
    v8 = (v8 - v8)
    v8 -= v8
    switch "\n\t" {
    case "foo bar":
        var v11 string = "foo bar"
        _ = v11
    case "":
        if (((v3 + -2.25) - 1.5) < (v3 / v3)) {
            var v12 uint32 = (v8 - v8)
            var v13 int16 = 772
            _ = v12
            _ = v13
        } else {
            if ((v8 & v8) == (v8 + v8)) {
                var v14 string = "\n\t"
                v15 := GenS1{F0: (v8 * 2), F1: false}
                _ = v14
                _ = v15
            } else {
                fmt.Println(v8, v3, v7)
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 int8 = 127
                    v2 -= v2
                    v18 := GenS1{F0: (v8 | (v8 & v8)), F1: (!(v7 == 1))}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS1{F0: (v8 % ((969) | 1)), F1: v9}
                fmt.Println(v7)
                _ = fi19
                _ = v20
            }
            v3 = 0.1
        }
    }
    var v21 string = "\n\t"
    v22 := GenS1{F0: v8, F1: (v10 < float32(0.1))}
    v23 := GenS1{F0: (v8 | 4294967295), F1: (!true)}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        v5 -= v5
    } else {
        if v9 {
            for fi24 := 0; fi24 < 17; fi24++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v25 uint64 = 0
                if v9 {
                    if ((98 == -1) && (2 >= 2)) {
                        var v26 int32 = 0
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 5; fi27++ {
                        v28 := GenS1{F0: v8, F1: (-117 != 1)}
                        if (false || v9) {
                            var v29 int16 = 1000
                            if (v21 != (v21 + v21)) {
                                v30 := GenS1{F0: uint32(v25), F1: v9}
                                _ = v30
                            } else {
                                var v31 int8 = int8(fi27)
                                var v32 string = v21
                                _ = v31
                                _ = v32
                            }
                            _ = v29
                        }
                        fmt.Println(v2, v25)
                        _ = fi27
                        _ = v28
                    }
                    if (v7 != uint16(v2)) {
                        fmt.Println(v2)
                        var v33 uint16 = (v7 >> ((546) & 63))
                        _ = v33
                    } else {
                        v3 = float64(v10)
                        var v34 uint16 = (v7 | (v7 * (^(v7))))
                        _ = v34
                    }
                } else {
                    v7 = 1
                    _ = len(v1)
                }
                _ = fi24
                _ = v25
            }
            var v35 int8 = -1
            _ = v35
        }
        var v36 float32 = v10
        var v37 int8 = 1
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 uint8 = v2
        v1 = append(v1, 2147483647)
        var v40 int64 = (^(v5))
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := GenS2{F0: 0}
    fmt.Println(v2)
    fmt.Println(v9)
    if v9 {
        fmt.Println(v21, v9, v21)
        delete(v4, v21)
        v42 := GenS3{F0: 228, F1: v9, F2: v21}
        _ = v42
    } else {
        if (32767 > -2147483648) {
            v43 := GenS1{F0: (v8 ^ 848), F1: v9}
            v8 = (v8 / (((v8 & v8)) | 1))
            _ = v43
        }
        var v44 int32 = 42
        _ = v44
    }
    v5 += -777
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
