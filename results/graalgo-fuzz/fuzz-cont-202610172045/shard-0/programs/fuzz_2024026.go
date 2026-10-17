package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 string; F2 uint8 }
func genFn0() uint {
    var v6 uint8 = 97
    fmt.Println(v6)
    v6 -= (v6 / ((233) | 1))
    _ = v6
    return 797
}
func genFn1() int32 {
    var v7 bool = (!true)
    v8 := []int{2147483647}
    v9 := GenS1{F0: 1, F1: 750, F2: 0}
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 int8 = 127
    var v2 string = "\n\t"
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    v5 := make(map[uint32]bool)
    var v10 uint16 = 0
    if false {
        v11 := GenS2{F0: (true || true), F1: 0}
        _ = v11
    } else {
        switch ("foo bar" + v2) {
        case "":
            var v12 int64 = 0
            _ = v12
        case "\n\t":
            v1 -= (v1 & (v1 & 100))
        case "foo bar":
            var v13 int8 = 100
            _ = v13
        default:
            switch -730 {
            case 108:
                if (351 == 1) {
                    v14 := GenS3{F0: 42, F1: -2147483648, F2: (v1 | -95)}
                    _ = v14
                }
            case 367:
                if true {
                    var v15 uint8 = 42
                    for fi16 := 0; fi16 < 10; fi16++ {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            var v18 uint64 = 1
                            v19 := GenS3{F0: -1, F1: -1, F2: v1}
                            v20 := GenS1{F0: v1, F1: 42, F2: 1}
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        _ = fi16
                    }
                    var v21 int = -258
                    _ = v15
                    _ = v21
                }
            case -159:
                for fi22 := 0; fi22 < 1; fi22++ {
                    v23 := []int{fi22, (fi22 % ((fi22) | 1)), fi22, fi22}
                    v1 -= (^(v1))
                    _ = fi22
                    _ = v23
                }
            }
        }
        if true {
            for fi24 := 0; fi24 < 11; fi24++ {
                var v25 string = v2
                _ = fi24
                _ = v25
            }
            var v26 int = -107
            v27 := GenS4{F0: 264, F1: genFn0()}
            _ = v26
            _ = v27
        }
        var v28 int16 = 69
        _ = v28
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        if (true || true) {
            switch "a" {
            case "":
                v1 += (v1 & v1)
            case "foo bar":
                v2 = (("\n\t" + v2) + v2)
            case "\n\t":
                switch (("foo bar" + v2) + v2) {
                case "foo bar":
                    fmt.Println(v10)
                case "\n\t":
                    if false {
                        var v30 uint64 = 783
                        var v31 uint8 = 128
                        _ = v30
                        _ = v31
                    } else {
                        v32 := GenS5{F0: 2}
                        v2 = "hello"
                        _ = v32
                    }
                }
            default:
                v10 -= v10
            }
            v10 = (v10 - v10)
            fmt.Println(v1)
        }
        v10 = (v10 - v10)
        _ = fi29
    }
    fmt.Println(v10, v10, v1)
    v2 = (v2 + v2)
    var v33 uint64 = 4294967295
    v34 := []int{342}
    v35 := GenS6{F0: -487, F1: ("foo bar" + (v2 + (v2 + v2))), F2: 2}
    for fi36 := 0; fi36 < 8; fi36++ {
        if false {
            v2 = ""
        }
        switch "foo bar" {
        case "hello":
            for fi37 := 0; fi37 < 20; fi37++ {
                v10 = v10
                v38 := v4[fi36]
                v38 = "hello"
                _ = fi37
                _ = v38
            }
        case "":
            delete(v4, fi36)
        case "a":
            v2 = "hello"
        }
        _ = fi36
    }
    v39 := GenS1{F0: (^(v1)), F1: v33, F2: 1}
    var v40 uint32 = 2147483647
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
