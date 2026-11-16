package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]bool)
    v3 := []uint32{65535, 2147483648}
    var v4 uint16 = 126
    for fi5 := 0; fi5 < 20; fi5++ {
        if false {
            v4 = 1
        }
        var v6 uint8 = 255
        var v7 uint64 = 28
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 string = "\n\t"
    _ = len(v2)
    if false {
        if true {
            _ = len(v2)
            v3 = append(v3, 42)
            var v9 int = -248
            _ = v9
        } else {
            if false {
                v1 += 1
            }
            v10 := v2[270]
            _ = v10
        }
    }
    v11 := GenS1{F0: -1}
    var v12 uint64 = 1
    if (!(819 >= 0)) {
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 uint8 = 0
                if (((fi13 | fi13) <= 32768) || (0.0 < ((3.14159 / 1e10) / 1.0))) {
                    fmt.Println(v15)
                } else {
                    var v16 int8 = 20
                    _ = v16
                }
                v17 := GenS1{F0: 1}
                _ = fi14
                _ = v15
                _ = v17
            }
            for fi18 := 0; fi18 < 11; fi18++ {
                v8 = v8
                var v19 bool = (("foo bar" + "") <= "foo bar")
                delete(v2, (fi18 - (-(fi13))))
                v20 := GenS2{F0: 4294967295}
                switch fi18 {
                case 610:
                    fmt.Println(v8, v8)
                case 42:
                    fmt.Println(v8, v8)
                case 0:
                    switch fi18 {
                    case -1:
                        var v21 uint64 = (v12 / ((uint64(v4)) | 1))
                        _ = v21
                    case -461:
                        v22 := GenS2{F0: v12}
                        var v23 float32 = 3.14159
                        var v24 int16 = 42
                        _ = v22
                        _ = v23
                        _ = v24
                    case 0:
                        var v25 uint = 65535
                        _ = v25
                    default:
                        switch (v8 + (v8 + "hello")) {
                        case "\n\t":
                            v4 = v4
                            var v26 uint = v1
                            _ = v26
                        case "":
                            v19 = (!(0 <= 546))
                        default:
                            if v19 {
                                v8 = v8
                            }
                        }
                    }
                default:
                    fmt.Println(v4)
                }
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi13
        }
        switch "hello" {
        case "\n\t":
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 uint8 = 0
                var v29 uint16 = v4
                _ = fi27
                _ = v28
                _ = v29
            }
        default:
            var v30 float64 = 0.0
            _ = v30
        }
    } else {
        switch -9 {
        case 42:
            var v31 int32 = 231
            _ = v31
        case -9223372036854775808:
            if (!false) {
                switch v8 {
                case "\n\t":
                    var v32 uint16 = (v4 * (v4 % ((v4) | 1)))
                    _ = v32
                case "":
                    var v33 float64 = -2.25
                    _ = v33
                }
            } else {
                v34 := GenS1{F0: -1}
                var v35 bool = true
                v12 = 1
                _ = v34
                _ = v35
            }
        case 128:
            v8 = (v8 + ("\n\t" + ((v8 + ((v8 + "\n\t") + (v8 + v8))) + "\n\t")))
        default:
            if false {
                v2[4294967295] = true
                v36 := GenS3{F0: v12, F1: 4294967295}
                switch ("" + (v8 + v8)) {
                case "\n\t":
                    v8 = "\n\t"
                }
                _ = v36
            } else {
                _ = len(v2)
            }
        }
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        var v38 int64 = 1
        var v39 int8 = -81
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v8
}
