package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint32; F2 int8 }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = uint32(v1)
    var v3 int8 = 100
    v4 := GenS1{F0: (v2 + 65535), F1: false}
    fmt.Println(v2, v1)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{(^(fi5))}
        var v7 int64 = -1
        _ = fi5
        _ = v6
        _ = v7
    }
    if (569 == 886) {
        fmt.Println(v3, v3)
        v1 = (v1 - v1)
    } else {
        var v8 int = -231
        if false {
            fmt.Println(v3, v1, v3)
        } else {
            fmt.Println(v3, v3, v1)
            var v9 uint64 = 4294967295
            var v10 bool = false
            _ = v10
            _ = v9
        }
        _ = v8
    }
    v11 := []int{4294967295, 1, int(v3), 32768}
    _ = len(v11)
    v12 := GenS1{F0: 731, F1: (true || true)}
    for fi13 := 0; fi13 < 7; fi13++ {
        if (false || false) {
            var v14 float64 = 0.0
            fmt.Println(v1, v2)
            v3 = (v3 & (v3 / ((v3) | 1)))
            _ = v14
        }
        fmt.Println(v1)
        var v15 float32 = 1.5
        _ = fi13
        _ = v15
    }
    var v16 float32 = 1e10
    for fi17 := 0; fi17 < 4; fi17++ {
        switch fi17 {
        case -306:
            v1 = v1
        }
        if false {
            switch ("a" + ("\n\t" + "hello")) {
            case "foo bar":
                if ((-2147483648 <= 1) && (true || ((true || false) && (55 < 235)))) {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                } else {
                    v18 := GenS2{F0: -1}
                    if (0 == 812) {
                        v16 = (float32(0.0) + 0.0)
                        for fi19 := 0; fi19 < 15; fi19++ {
                            if false {
                                var v20 int16 = 915
                                var v21 uint16 = 218
                                fmt.Println(v16, fi17, v20)
                                _ = v20
                                _ = v21
                            }
                            fmt.Println(v3, v3, v16)
                            _ = fi19
                        }
                        for fi22 := 0; fi22 < 9; fi22++ {
                            v3 = (^(v3))
                            _ = fi22
                        }
                    }
                    var v23 int16 = 0
                    _ = v18
                    _ = v23
                }
            case "\n\t":
                switch fi17 {
                case 42:
                    for fi24 := 0; fi24 < 20; fi24++ {
                        if (v2 >= (v2 - (v2 & (v2 * (v2 >> ((4294967295) & 63)))))) {
                            _ = len(v11)
                            fmt.Println(v16)
                        }
                        _ = fi24
                    }
                }
            default:
                v25 := GenS3{F0: (!(4294967295 != 734)), F1: 2147483647, F2: (v3 - v3)}
                _ = v25
            }
            v2 -= v2
        }
        _ = fi17
    }
    switch -909 {
    case 1:
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 uint16 = 615
            fmt.Println(v3, v2, v3)
            _ = fi26
            _ = v27
        }
    case 42:
        v28 := GenS2{F0: 455}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
