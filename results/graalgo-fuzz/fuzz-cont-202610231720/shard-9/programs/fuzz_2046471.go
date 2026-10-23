package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 int; F2 string }
type GenS3 struct { F0 int32; F1 int8; F2 uint }

func main() {
    var v1 int16 = -1
    v2 := []int{-378, 2147483647, 9223372036854775807, -508}
    v3 := make(map[uint32]int)
    var v4 int = -1
    v5 := make(map[int]int)
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: 329, F1: "a"}
    v1 += (v1 & v1)
    var v8 string = "foo bar"
    var v9 uint64 = 2
    if (!true) {
        var v10 int64 = -12
        v11 := GenS1{F0: 65535, F1: (v8 + v8)}
        var v12 uint64 = 2
        _ = v10
        _ = v11
        _ = v12
    } else {
        var v13 int8 = 6
        if (true || true) {
            var v14 int8 = v13
            for fi15 := 0; fi15 < 20; fi15++ {
                if ((((!true) || true) && false) || true) {
                    v16 := GenS1{F0: 2147483647, F1: ("\n\t" + "hello")}
                    v2 = append(v2, v4)
                    _ = v16
                }
                _ = fi15
            }
            if true {
                switch v4 {
                case 589:
                    v2 = append(v2, v4)
                case 128:
                    var v17 int64 = 1
                    _ = v17
                case 32768:
                    if (false && (168 == 1)) {
                        fmt.Println(v9, v9)
                        switch int(v1) {
                        case -2147483648:
                            v18 := GenS1{F0: 315, F1: v8}
                            _ = v18
                        case 2147483647:
                            v6[(v4 ^ 128)] = 2
                        case 128:
                            var v19 float32 = 1e10
                            _ = v19
                        }
                        switch (v4 | v4) {
                        case 42:
                            delete(v5, v4)
                        case 9223372036854775807:
                            var v20 float64 = 0.1
                            _ = v20
                        }
                    } else {
                        delete(v5, v4)
                        _ = len(v3)
                        var v21 string = v8
                        _ = v21
                    }
                }
            } else {
                if false {
                    for fi22 := 0; fi22 < 1; fi22++ {
                        v4 = fi22
                        _ = fi22
                    }
                }
            }
            _ = v14
        }
        _ = v13
    }
    v23 := GenS2{F0: -1, F1: 32768, F2: v8}
    v24 := GenS1{F0: 1, F1: v8}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v25 uint64 = (^(v9))
    if false {
        for fi26 := 0; fi26 < 20; fi26++ {
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 uint = 4294967295
                _ = fi27
                _ = v28
            }
            switch v4 {
            case -2147483648:
                switch "\n\t" {
                case "foo bar":
                    var v29 uint = 18446744073709551615
                    _ = v29
                case "hello":
                    _ = len(v2)
                case "a":
                    for fi30 := 0; fi30 < 20; fi30++ {
                        v31 := GenS1{F0: 224, F1: (v8 + v8)}
                        _ = fi30
                        _ = v31
                    }
                default:
                    fmt.Println(v8, v1, v4)
                }
            case -28:
                if false {
                    delete(v6, (^(fi26)))
                    v32 := GenS3{F0: 1, F1: 1, F2: 255}
                    _ = v32
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            if len(v2) > 0 {
                v2[0] = v4
            }
            _ = fi26
        }
    } else {
        v33 := GenS1{F0: uint32(v9), F1: ("hello" + "foo bar")}
        _ = len(v3)
        _ = v33
    }
    delete(v5, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
