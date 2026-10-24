package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint; F2 uint8 }
type GenS5 struct { F0 string }
func genFn0(p0 int8, p1 int64) int16 {
    var v4 int16 = 1
    v5 := GenS1{F0: (!false)}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return int16(p1)
}

func main() {
    v1 := []uint32{4294967295, 781, 2147483647}
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    switch ("hello" + "hello") {
    case "\n\t":
        _ = len(v1)
    case "":
        var v6 uint8 = 9
        _ = v6
    default:
        _ = len(v3)
    }
    var v7 int32 = 2147483647
    var v8 int8 = 1
    switch 0 {
    case -732:
        v7 += v7
    default:
        switch "hello" {
        case "foo bar":
            if true {
                if ("foo bar" <= ("foo bar" + ("foo bar" + ""))) {
                    if (!false) {
                        var v9 uint64 = 42
                        _ = v9
                    } else {
                        v10 := GenS2{F0: (false && true), F1: 42, F2: -1}
                        v1 = append(v1, 622)
                        _ = v10
                    }
                } else {
                    var v11 int32 = (v7 - v7)
                    v7 += v11
                    for fi12 := 0; fi12 < 2; fi12++ {
                        v8 += v8
                        v8 = (v8 & v8)
                        _ = fi12
                    }
                    _ = v11
                }
            } else {
                var v13 int64 = 156
                if false {
                    if (v7 != (v7 / (((v7 - -31)) | 1))) {
                        var v14 uint32 = 2147483647
                        var v15 uint32 = 647
                        var v16 uint16 = 65535
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    var v17 uint16 = 141
                    _ = len(v1)
                    _ = v17
                }
                if len(v1) > 0 {
                    v1[0] = 2
                }
                _ = v13
            }
        case "a":
            delete(v2, "\n\t")
        case "":
            v2[""] = (("foo bar" + "foo bar") + "a")
        }
    }
    v18 := GenS1{F0: (32768 != 9223372036854775807)}
    v8 -= v8
    var v19 float32 = 0.1
    var v20 uint32 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v21 := GenS1{F0: true}
    fmt.Println(v7, v8, v8)
    v22 := GenS2{F0: false, F1: 2, F2: 1}
    for fi23 := 0; fi23 < 17; fi23++ {
        if false {
            switch "\n\t" {
            case "":
                v8 += (^(v8))
            case "a":
                v24 := v3[733]
                _ = v24
            case "\n\t":
                delete(v2, ("foo bar" + "foo bar"))
            default:
                for fi25 := 0; fi25 < 20; fi25++ {
                    var v26 uint = 255
                    _ = fi25
                    _ = v26
                }
            }
            v27 := GenS3{F0: 42}
            var v28 int16 = -701
            _ = v27
            _ = v28
        }
        _ = fi23
    }
    if true {
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 float32 = v19
            for fi31 := 0; fi31 < 2; fi31++ {
                var v32 int = fi31
                if false {
                    var v33 float32 = ((v19 * 1.5) * float32(-2.25))
                    if true {
                        v8 -= v8
                        var v34 int16 = 159
                        v35 := GenS4{F0: v7, F1: 18446744073709551615, F2: 127}
                        _ = v34
                        _ = v35
                    }
                    v7 -= v7
                    _ = v33
                }
                var v36 uint64 = 318
                var v37 uint8 = 255
                _ = fi31
                _ = v32
                _ = v36
                _ = v37
            }
            _ = fi29
            _ = v30
        }
    }
    if false {
        var v38 int64 = -437
        v38 = v38
        _ = v38
    } else {
        var v39 int64 = -734
        _ = v39
    }
    switch -2147483648 {
    case 9223372036854775807:
        v1 = append(v1, (v20 ^ 2147483647))
    case 331:
        switch "hello" {
        case "":
            var v40 uint32 = (v20 / (((v20 ^ 666)) | 1))
            _ = v40
        case "hello":
            fmt.Println(v20, v8, v19)
        case "foo bar":
            if false {
                switch -1 {
                case -2147483648:
                    v41 := GenS5{F0: (("hello" + "a") + "hello")}
                    _ = v41
                }
            } else {
                v2[""] = ("" + "")
            }
        default:
            v42 := GenS2{F0: true, F1: 4294967295, F2: 1}
            _ = v42
        }
    case -9223372036854775808:
        v3[v20] = 647
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v7
    _ = v8
}
