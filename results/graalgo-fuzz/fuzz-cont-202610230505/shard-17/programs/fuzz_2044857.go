package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint; F1 int8; F2 string }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 string }

func main() {
    var v1 uint = 85
    v2 := make(map[int]uint32)
    var v3 int16 = -1
    var v4 float32 = 0.0
    v5 := make(map[int]string)
    var v6 int32 = -2147483648
    switch 0 {
    case 127:
        v7 := GenS1{F0: 1, F1: ""}
        _ = v7
    case -9223372036854775808:
        v1 = (v1 << (((^(v1))) & 63))
    default:
        v8 := GenS2{F0: 1, F1: 100, F2: ("foo bar" + "")}
        _ = v8
    }
    var v9 int64 = 1
    var v10 uint = (v1 >> ((2) & 63))
    v11 := []int{}
    v6 -= (v6 - (v6 - (v6 << (((v10 * (v10 ^ v10))) & 63))))
    if (1 >= 2) {
        switch "a" {
        case "foo bar":
            var v12 float64 = 0.1
            _ = v12
        case "":
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 string = "hello"
                v15 := GenS2{F0: (v10 | v1), F1: 1, F2: v14}
                fmt.Println(fi13, v6)
                _ = fi13
                _ = v14
                _ = v15
            }
        case "\n\t":
            for fi16 := 0; fi16 < 15; fi16++ {
                if (!((^(v1)) >= v1)) {
                    v2[fi16] = 65535
                    if len(v11) > 0 {
                        v11[0] = 42
                    }
                } else {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    if false {
                        v2[fi16] = 65535
                    } else {
                        v17 := GenS3{F0: 785, F1: 960}
                        for fi18 := 0; fi18 < 17; fi18++ {
                            v10 = v1
                            v6 -= v6
                            v10 -= (v1 >> (((v10 * v1)) & 63))
                            v19 := GenS1{F0: 961, F1: ("hello" + ("hello" + ("" + "a")))}
                            _ = fi18
                            _ = v19
                        }
                        _ = v17
                    }
                }
                _ = fi16
            }
        default:
            if false {
                var v20 uint32 = 0
                _ = v20
            }
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        if true {
            if true {
                v21 := GenS3{F0: 955, F1: 32767}
                var v22 int64 = (v9 | -1)
                for fi23 := 0; fi23 < 20; fi23++ {
                    v2[fi23] = 2147483647
                    if false {
                        v24 := GenS4{F0: 65535, F1: -17}
                        if len(v11) > 0 {
                            v11[0] = fi23
                        }
                        _ = v24
                    } else {
                        var v25 uint = (v10 - (v10 - v10))
                        _ = v25
                    }
                    if false {
                        if true {
                            v26 := GenS4{F0: 65535, F1: 100}
                            _ = v26
                        }
                    } else {
                        for fi27 := 0; fi27 < 8; fi27++ {
                            v3 = (v3 / ((723) | 1))
                            var v28 int8 = 0
                            _ = fi27
                            _ = v28
                        }
                        var v29 string = "foo bar"
                        var v30 bool = true
                        _ = v29
                        _ = v30
                    }
                    _ = fi23
                }
                _ = v21
                _ = v22
            }
            switch -9223372036854775808 {
            case 127:
                var v31 uint = 1
                _ = v31
            }
            var v32 float64 = 0.0
            _ = v32
        }
    }
    v9 = 0
    if (false || true) {
        v33 := GenS5{F0: "a"}
        delete(v5, int(v3))
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
