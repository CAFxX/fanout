package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 string }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS3 struct { F0 bool; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = true
    v3 := make(map[int]bool)
    v2 = true
    if (!((v2 || false) && (1.5 > ((3.14159 + -2.25) + ((1e10 / float64(1e10)) + 1.5))))) {
        for fi4 := 0; fi4 < 2; fi4++ {
            for fi5 := 0; fi5 < 17; fi5++ {
                fmt.Println(v2)
                fmt.Println(fi4, fi4, v2)
                var v6 string = ("\n\t" + "a")
                _ = fi5
                _ = v6
            }
            _ = fi4
        }
    }
    v2 = (true && (128 < 975))
    fmt.Println(v2)
    v2 = (89 == 1)
    v7 := []int{127, 4294967295, 42}
    v8 := GenS1{F0: 110, F1: 9223372036854775807, F2: "hello"}
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 1, F1: 9223372036854775807, F2: "hello"}
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 255, F1: -327, F2: "\n\t"}
    v12 := GenS1{F0: 158, F1: 42, F2: "foo bar"}
    switch 127 {
    case -326:
        v2 = (176 == -315)
    default:
        var v13 uint32 = 2147483647
        _ = v13
    }
    v2 = (1 != 255)
    switch 9223372036854775807 {
    case 623:
        if true {
            v14 := GenS2{F0: 1, F1: 130, F2: 1}
            _ = v14
        }
    case 127:
        var v15 int8 = 0
        _ = v15
    case -679:
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int16 = -32768
            v7 = append(v7, fi16)
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        if false {
            switch (fi18 | fi18) {
            case 127:
                v19 := GenS1{F0: 0, F1: 463, F2: "a"}
                _ = v19
            case -9223372036854775808:
                v2 = v2
            }
            if (0 == 255) {
                var v20 int64 = -9223372036854775808
                _ = v20
            }
            var v21 uint64 = 18446744073709551615
            _ = v21
        }
        _ = fi18
    }
    v22 := GenS1{F0: 45, F1: 862, F2: ("a" + "\n\t")}
    if (!(true && true)) {
        var v23 string = "foo bar"
        _ = len(v7)
        _ = v23
    }
    v2 = (v2 && v2)
    v24 := GenS3{F0: (-570 > 818), F1: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v7
    _ = v8
}
