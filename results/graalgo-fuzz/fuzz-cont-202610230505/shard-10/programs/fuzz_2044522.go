package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 string }
func genFn0() uint64 {
    switch 42 {
    case 4294967295:
        v5 := GenS1{F0: 42, F1: 1, F2: 2}
        _ = v5
    default:
        if (350 >= 0) {
            var v6 uint64 = 0
            _ = v6
        } else {
            if (4294967295 > -616) {
                if true {
                    v7 := []int{42}
                    for fi8 := 0; fi8 < 6; fi8++ {
                        var v9 int32 = -2147483648
                        _ = fi8
                        _ = v9
                    }
                    _ = v7
                }
                v10 := []int{32768, 443}
                _ = v10
            }
            var v11 int64 = 0
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := []int{fi12}
        fmt.Println(fi12)
        if true {
            var v14 uint8 = 127
            if false {
                var v15 uint16 = 65535
                var v16 int64 = -649
                _ = v15
                _ = v16
            } else {
                switch "" {
                case "\n\t":
                    if (false || true) {
                        v14 = (v14 - (v14 + 128))
                    } else {
                        for fi17 := 0; fi17 < 13; fi17++ {
                            v14 = (v14 >> ((320) & 63))
                            fmt.Println(fi12, fi17, v14)
                            _ = fi17
                        }
                    }
                case "hello":
                    switch ("\n\t" + ((("a" + ("foo bar" + "\n\t")) + "hello") + "foo bar")) {
                    case "a":
                        fmt.Println(fi12, fi12)
                    case "foo bar":
                        if true {
                            for fi18 := 0; fi18 < 17; fi18++ {
                                var v19 int32 = int32(fi12)
                                _ = fi18
                                _ = v19
                            }
                            fmt.Println(fi12)
                        }
                    default:
                        for fi20 := 0; fi20 < 4; fi20++ {
                            v21 := GenS2{F0: 1, F1: (false || ((697 <= 1) && (288 == 782)))}
                            var v22 string = ""
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                case "a":
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 float32 = (0.1 - (1e10 * 1.5))
                        v25 := GenS1{F0: 203, F1: 84, F2: 1}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                default:
                    var v26 float64 = (1.5 - 0.1)
                    _ = v26
                }
                fmt.Println(v14, fi12)
                var v27 uint32 = 134
                _ = v27
            }
            var v28 int8 = 123
            _ = v14
            _ = v28
        }
        _ = fi12
        _ = v13
    }
    return 647
}
func genFn1(p0 int64, p1 uint) uint8 {
    v29 := GenS2{F0: (p1 - (p1 ^ 0)), F1: false}
    var v30 uint32 = 4294967295
    v30 = 284
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    return 1
}

func main() {
    var v1 uint32 = 65535
    v2 := []string{("hello" + "a")}
    var v3 uint8 = 127
    v4 := []uint32{}
    v31 := GenS3{F0: ""}
    var v32 uint32 = v1
    var v33 int64 = 0
    v1 = 634
    v1 = (v32 | (v32 | 0))
    v34 := GenS1{F0: v3, F1: -128, F2: 987}
    fmt.Println(v3, v32)
    v35 := GenS1{F0: (^(v3)), F1: -1, F2: 255}
    var v36 bool = false
    var v37 float32 = float32(-2.25)
    v32 = 4294967295
    var v38 string = "foo bar"
    var v39 float64 = 0.1
    var v40 int64 = (v33 - (v33 << ((uint(v32)) & 63)))
    for fi41 := 0; fi41 < 14; fi41++ {
        v37 = v37
        var v42 int64 = int64(v1)
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
