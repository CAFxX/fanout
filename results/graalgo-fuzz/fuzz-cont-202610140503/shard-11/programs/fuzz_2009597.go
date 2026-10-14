package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int; F1 int32; F2 int32 }
type GenS4 struct { F0 string; F1 string }
func genFn0(p0 uint64, p1 int32) int16 {
    var v6 uint64 = p0
    var v7 uint16 = 32767
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint = 18446744073709551615
        v6 += (p0 | v6)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: uint8(v7), F1: 550}
    v11 := GenS2{F0: 421, F1: "a"}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    _ = v7
    return -32768
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 1
    var v3 float32 = 1e10
    v4 := make(map[string]uint32)
    var v5 uint32 = 2147483648
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS3{F0: fi12, F1: 2147483647, F2: 0}
        v2 -= (v2 * v2)
        _ = fi12
        _ = v13
    }
    v2 = 18446744073709551615
    var v14 string = "hello"
    var v15 uint = v2
    v16 := v1[v14]
    var v17 string = ((v14 + "hello") + ("\n\t" + v14))
    v15 += v2
    v5 -= v5
    v18 := []int{4294967295, int(v5), 127, -646}
    var v19 uint = v15
    for fi20 := 0; fi20 < 20; fi20++ {
        fmt.Println(v14, v14, v3)
        v21 := GenS3{F0: (fi20 + fi20), F1: 0, F2: -2147483648}
        var v22 string = v17
        _ = fi20
        _ = v21
        _ = v22
    }
    if (!(true || false)) {
        for fi23 := 0; fi23 < 4; fi23++ {
            for fi24 := 0; fi24 < 14; fi24++ {
                if true {
                    var v25 bool = (!false)
                    var v26 int8 = 1
                    fmt.Println(v2)
                    _ = v25
                    _ = v26
                }
                switch "a" {
                case "a":
                    _ = len(v18)
                }
                _ = fi24
            }
            if true {
                _ = len(v18)
                switch v16 {
                case "":
                    var v27 int = (fi23 * 0)
                    _ = v27
                }
                v19 += v2
            }
            _ = fi23
        }
        if true {
            var v28 uint32 = 2147483647
            _ = v28
        } else {
            for fi29 := 0; fi29 < 4; fi29++ {
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 int32 = -939
                    v19 -= 0
                    _ = fi30
                    _ = v31
                }
                var v32 uint64 = 42
                var v33 uint64 = 4294967295
                _ = fi29
                _ = v32
                _ = v33
            }
            v1[v16] = v17
        }
    }
    switch v14 {
    case "foo bar":
        switch -350 {
        case 127:
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        case -117:
            if len(v18) > 0 {
                v18[0] = 32767
            }
        case 1:
            v34 := GenS2{F0: 415, F1: "hello"}
            _ = v34
        default:
            switch v17 {
            case "a":
                if true {
                    var v35 uint64 = uint64(v19)
                    v4[((v17 + "hello") + "a")] = v5
                    if true {
                        var v36 int32 = 743
                        _ = v36
                    } else {
                        fmt.Println(v5, v16)
                        v2 += 18446744073709551615
                        for fi37 := 0; fi37 < 14; fi37++ {
                            switch 128 {
                            case -9223372036854775808:
                                v38 := GenS4{F0: "a", F1: v16}
                                _ = v38
                            }
                            v5 = 42
                            _ = fi37
                        }
                    }
                    _ = v35
                }
            case "foo bar":
                v39 := GenS1{F0: 30, F1: 42}
                _ = v39
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
