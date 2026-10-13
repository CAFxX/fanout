package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint16; F1 int; F2 int16 }
func genFn0(p0 int32, p1 string) int16 {
    v6 := []int{4294967295}
    p1 = p1
    fmt.Println(p1)
    for fi7 := 0; fi7 < 1; fi7++ {
        _ = len(v6)
        for fi8 := 0; fi8 < 6; fi8++ {
            p1 = p1
            _ = fi8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint16 = 255
            _ = fi9
            _ = v10
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return 20
}
func genFn1(p0 int16) int32 {
    p0 += p0
    v11 := GenS1{F0: 0, F1: 1, F2: 42}
    v12 := GenS1{F0: 2147483648, F1: 127, F2: (p0 >> ((0) & 63))}
    var v13 int = 32768
    v14 := GenS2{F0: 1}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]string)
    var v3 float32 = (((((3.14159 + 0.0) * (0.1 + (-2.25 / 1.5))) + 3.14159) / float32(-2.25)) + 0.0)
    var v4 float32 = v3
    v5 := make(map[string]int)
    fmt.Println(v4, v3, v4)
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    v17 := GenS3{F0: "foo bar", F1: (!false), F2: 4294967295}
    switch 127 {
    case -2147483648:
        if (156 > 248) {
            v3 = ((v4 - (1.5 + ((v4 * v4) * (-2.25 - (v4 / 1.5))))) + v4)
        }
    case 125:
        var v18 uint64 = 2
        _ = v18
    case 2147483647:
        v3 = v3
    default:
        if (!false) {
            fmt.Println(v4, v3, v4)
            if false {
                if true {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        for fi20 := 0; fi20 < 9; fi20++ {
                            var v21 uint16 = 0
                            var v22 int8 = -1
                            v23 := GenS4{F0: 255, F1: fi20, F2: -1}
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        v24 := GenS3{F0: "\n\t", F1: false, F2: 77}
                        v2[(("\n\t" + "hello") + ("foo bar" + "a"))] = ("a" + "foo bar")
                        _ = fi19
                        _ = v24
                    }
                    v25 := GenS1{F0: 1, F1: 42, F2: 1}
                    v4 = 1e10
                    _ = v25
                } else {
                    v26 := GenS2{F0: 65535}
                    _ = v26
                }
                var v27 int32 = 42
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 uint = 328
                    v27 += (^(v27))
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
        } else {
            switch 127 {
            case 2147483647:
                var v30 float32 = float32((0.0 - v1))
                _ = v30
            }
            if ((0 <= 42) && true) {
                var v31 int8 = 1
                _ = v31
            }
            _ = len(v2)
        }
    }
    delete(v5, "")
    var v32 int64 = 42
    if true {
        delete(v5, "hello")
    }
    fmt.Println(v4)
    var v33 uint64 = 2
    delete(v2, "a")
    var v34 uint = 393
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
