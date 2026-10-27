package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0(p0 uint8) uint {
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := []int{(^(fi4)), fi4, 960}
        switch fi4 {
        case 4294967295:
            var v6 int8 = 127
            _ = v6
        case 32767:
            var v7 bool = false
            _ = v7
        default:
            for fi8 := 0; fi8 < 1; fi8++ {
                fmt.Println(p0, p0, fi8)
                _ = fi8
            }
        }
        _ = fi4
        _ = v5
    }
    var v9 uint8 = p0
    if (0 > 687) {
        var v10 int64 = 1
        _ = v10
    }
    switch "foo bar" {
    case "":
        if true {
            v11 := GenS1{F0: 65535, F1: ""}
            _ = v11
        }
    case "a":
        v12 := []int{700, -1, 2147483647, -9223372036854775808}
        _ = v12
    case "\n\t":
        var v13 int64 = -1
        _ = v13
    default:
        if true {
            v14 := GenS1{F0: 84, F1: "a"}
            v15 := GenS1{F0: 241, F1: ("" + ("" + (((("foo bar" + ("hello" + "\n\t")) + ("a" + (("a" + "foo bar") + "foo bar"))) + "") + "foo bar")))}
            if ((3.14159 >= 1e10) || false) {
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 uint16 = 42
                    _ = fi16
                    _ = v17
                }
                var v18 uint8 = (v9 | (v9 % (((p0 % ((p0) | 1))) | 1)))
                if true {
                    var v19 int8 = 0
                    v20 := GenS2{F0: 372}
                    _ = v19
                    _ = v20
                }
                _ = v18
            } else {
                var v21 int = 32767
                _ = v21
            }
            _ = v14
            _ = v15
        }
    }
    _ = p0
    _ = v9
    return 4294967295
}

func main() {
    var v1 float32 = 3.14159
    v2 := []string{("foo bar" + "")}
    v3 := make(map[int]string)
    v1 = 3.14159
    v1 += v1
    v1 -= ((v1 - v1) - 1e10)
    for fi22 := 0; fi22 < 12; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            if false {
                v24 := GenS3{F0: 456, F1: 0}
                _ = v24
            } else {
                fmt.Println(fi22, v1)
                var v25 float64 = 1.5
                _ = v25
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v26 uint8 = 128
            _ = fi23
            _ = v26
        }
        _ = fi22
    }
    v1 = 1e10
    var v27 uint16 = 42
    switch ("hello" + ("\n\t" + "foo bar")) {
    case "hello":
        fmt.Println(v27, v1)
    }
    var v28 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
