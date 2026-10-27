package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }
func genFn0() bool {
    var v5 string = "\n\t"
    v5 = (v5 + "\n\t")
    _ = v5
    return false
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"a"}
    var v3 float64 = 3.14159
    var v4 int32 = -1
    fmt.Println(v4)
    if true {
        if true {
            v6 := GenS1{F0: 1, F1: 2147483648, F2: 228}
            if false {
                var v7 int = -1
                fmt.Println(v4, v7, v4)
                _ = v7
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            fmt.Println(v3, v3)
            _ = v6
        } else {
            v1[818] = (!(1 <= 276))
            v8 := GenS1{F0: 18446744073709551615, F1: 895, F2: 2}
            if (!(!false)) {
                for fi9 := 0; fi9 < 10; fi9++ {
                    if ((!true) || false) {
                        if (-672 != 106) {
                            var v10 float64 = (0.1 + v3)
                            v3 = v10
                            _ = v10
                        } else {
                            v3 = -2.25
                            _ = len(v2)
                            var v11 int16 = 42
                            _ = v11
                        }
                        v3 += v3
                        v3 += v3
                    } else {
                        v12 := GenS2{F0: -888, F1: 255}
                        fmt.Println(v4, v3, fi9)
                        _ = v12
                    }
                    v13 := GenS3{F0: (v4 ^ v4)}
                    _ = fi9
                    _ = v13
                }
                var v14 int64 = 0
                v3 += v3
                _ = v14
            } else {
                for fi15 := 0; fi15 < 13; fi15++ {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 uint = 18446744073709551615
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 17; fi18++ {
                        var v19 bool = true
                        _ = fi18
                        _ = v19
                    }
                    switch ("a" + "a") {
                    case "a":
                        var v20 int16 = 42
                        _ = v20
                    case "foo bar":
                        v3 = (v3 + v3)
                    case "hello":
                        for fi21 := 0; fi21 < 3; fi21++ {
                            _ = len(v1)
                            var v22 int32 = v4
                            v23 := GenS1{F0: 1, F1: 2147483647, F2: 128}
                            var v24 int8 = 42
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    }
                    _ = fi15
                }
                var v25 float32 = 1.5
                fmt.Println(v4, v3)
                _ = v25
            }
            _ = v8
        }
    }
    v1[4294967295] = true
    for fi26 := 0; fi26 < 16; fi26++ {
        switch "\n\t" {
        case "":
            v3 += v3
        case "a":
            v27 := GenS2{F0: -32768, F1: 44}
            _ = v27
        default:
            v4 = v4
        }
        v4 = v4
        fmt.Println(fi26, fi26, fi26)
        _ = fi26
    }
    v3 += (v3 - float64(3.14159))
    if (0 >= 18446744073709551615) {
        var v28 uint64 = 77
        _ = v28
    }
    if false {
        for fi29 := 0; fi29 < 13; fi29++ {
            for fi30 := 0; fi30 < 11; fi30++ {
                if false {
                    v3 -= float64(0.0)
                    v3 = 1e10
                    fmt.Println(fi30, v4)
                }
                v31 := GenS4{F0: ("hello" >= "a")}
                _ = fi30
                _ = v31
            }
            var v32 string = "a"
            v33 := GenS2{F0: -1, F1: 5}
            _ = fi29
            _ = v32
            _ = v33
        }
        delete(v1, 65535)
        v34 := GenS2{F0: 42, F1: 0}
        _ = v34
    }
    var v35 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
