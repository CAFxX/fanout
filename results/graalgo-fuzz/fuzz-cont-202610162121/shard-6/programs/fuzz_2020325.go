package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint64; F1 uint16; F2 int }
type GenS4 struct { F0 int16; F1 string }
func genFn0() uint {
    var v7 int = 0
    var v8 string = "a"
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS1{F0: 4294967295}
        var v11 int16 = -1
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 uint, p1 string, p2 uint8) uint16 {
    switch (p1 + "a") {
    case "a":
        v12 := GenS1{F0: uint64(p0)}
        _ = v12
    default:
        v13 := GenS1{F0: 1}
        _ = v13
    }
    v14 := GenS2{F0: ("\n\t" + p1), F1: (^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 65535
}

func main() {
    v1 := []int{4294967295}
    v2 := []uint32{1}
    v3 := make(map[string]bool)
    v4 := []string{"\n\t"}
    var v5 float32 = (1.5 / (1.5 - (0.1 + 1.5)))
    var v6 float64 = 0.0
    v15 := GenS1{F0: 42}
    delete(v3, "hello")
    v6 = (0.1 / (v6 - (1.5 / float64(float32(float64(-2.25))))))
    fmt.Println(v5, v6)
    if false {
        if false {
            if true {
                switch 128 {
                case 2147483647:
                    v6 += v6
                default:
                    for fi16 := 0; fi16 < 14; fi16++ {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            var v18 int32 = 2147483647
                            v6 = v6
                            _ = fi17
                            _ = v18
                        }
                        _ = fi16
                    }
                }
            } else {
                v19 := GenS3{F0: 0, F1: 1, F2: -9223372036854775808}
                var v20 uint64 = 42
                _ = v19
                _ = v20
            }
        } else {
            if len(v4) > 0 {
                v4[0] = ""
            }
            fmt.Println(v5, v5)
            fmt.Println(v5, v5)
        }
        v21 := GenS4{F0: 0, F1: "foo bar"}
        if true {
            fmt.Println(v6, v6)
            v6 -= v6
            v6 += v6
        } else {
            for fi22 := 0; fi22 < 1; fi22++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v23 uint = 4294967295
                var v24 uint32 = 0
                _ = fi22
                _ = v23
                _ = v24
            }
            delete(v3, "")
        }
        _ = v21
    }
    v6 = 3.14159
    fmt.Println(v5)
    v25 := GenS3{F0: 901, F1: 688, F2: 0}
    v5 = v5
    fmt.Println(v6)
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    var v26 int8 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
