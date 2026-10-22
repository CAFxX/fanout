package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int32 = 881
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    switch ("hello" + "foo bar") {
    case "":
        fmt.Println(v1)
    case "hello":
        v5 := v2[127]
        _ = v5
    default:
        switch ("\n\t" + (("a" + "foo bar") + ("hello" + "\n\t"))) {
        case "hello":
            if (1e10 == float32(1.5)) {
                v6 := []int{}
                _ = v6
            }
        default:
            for fi7 := 0; fi7 < 8; fi7++ {
                v3[2147483647] = ("foo bar" + "\n\t")
                switch fi7 {
                case 22:
                    var v8 int = (-(fi7))
                    _ = v8
                case 2147483647:
                    if true {
                        for fi9 := 0; fi9 < 15; fi9++ {
                            v10 := GenS1{F0: 1, F1: (fi9 ^ (fi7 / ((9223372036854775807) | 1))), F2: 42}
                            _ = fi9
                            _ = v10
                        }
                        var v11 uint8 = 0
                        _ = v11
                    }
                case -1:
                    fmt.Println(v1)
                }
                _ = fi7
            }
        }
    }
    v12 := GenS1{F0: -1, F1: -9223372036854775808, F2: 4294967295}
    var v13 uint = 18446744073709551615
    var v14 uint32 = 739
    fmt.Println(v1, v1)
    v15 := []int{4294967295, -125}
    v16 := GenS1{F0: 773, F1: int(v13), F2: 775}
    for fi17 := 0; fi17 < 3; fi17++ {
        if (uint16(v14) == 42) {
            v1 = v1
        } else {
            for fi18 := 0; fi18 < 20; fi18++ {
                if true {
                    v19 := GenS1{F0: -1, F1: fi18, F2: 513}
                    if false {
                        var v20 uint8 = 2
                        v21 := v3[v14]
                        var v22 uint = v13
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        var v23 int16 = 42
                        _ = v23
                    }
                    _ = v19
                } else {
                    v24 := v3[(v14 / (((v14 / ((uint32(v13)) | 1))) | 1))]
                    var v25 string = v24
                    v26 := GenS2{F0: 255}
                    _ = v24
                    _ = v25
                    _ = v26
                }
                var v27 int64 = -9223372036854775808
                fmt.Println(v14, v27, v1)
                _ = fi18
                _ = v27
            }
            v14 -= v14
            v13 = 18446744073709551615
        }
        var v28 int8 = 0
        _ = fi17
        _ = v28
    }
    _ = len(v4)
    fmt.Println(v13, v13, v13)
    delete(v2, 42)
    var v29 uint64 = 868
    if true {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        var v30 uint16 = 2
        if false {
            switch "a" {
            case "foo bar":
                v31 := GenS1{F0: 42, F1: 1, F2: uint64(v14)}
                _ = v31
            case "a":
                v13 = (v13 % (((v13 - 18446744073709551615)) | 1))
            }
        }
        _ = v30
    } else {
        v32 := GenS2{F0: 2}
        if len(v15) > 0 {
            v15[0] = -692
        }
        fmt.Println(v14)
        _ = v32
    }
    v33 := GenS2{F0: 172}
    for fi34 := 0; fi34 < 2; fi34++ {
        switch "hello" {
        case "":
            v13 = v13
        }
        var v35 int16 = -1
        _ = fi34
        _ = v35
    }
    v13 += uint(v14)
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    delete(v4, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
}
