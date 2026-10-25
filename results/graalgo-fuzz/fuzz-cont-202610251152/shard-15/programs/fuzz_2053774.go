package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int32; F1 int32; F2 uint32 }
type GenS3 struct { F0 string; F1 uint; F2 int8 }
func genFn0() uint8 {
    var v6 string = "hello"
    var v7 bool = (1 != 828)
    for fi8 := 0; fi8 < 16; fi8++ {
        switch "foo bar" {
        case "\n\t":
            v7 = (2 >= 2147483648)
        case "hello":
            fmt.Println(v7, v7, fi8)
        }
        fmt.Println(fi8, v6)
        v9 := []int{(fi8 | 37), fi8}
        _ = fi8
        _ = v9
    }
    fmt.Println(v6, v6)
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 bool, p1 uint) uint64 {
    p1 = (p1 / (((^(p1))) | 1))
    p1 = p1
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint = p1
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        switch ("\n\t" + ("hello" + ("\n\t" + ("foo bar" + "a")))) {
        case "foo bar":
            var v13 bool = p0
            _ = v13
        default:
            v14 := []int{fi12}
            _ = v14
        }
        p1 -= p1
        _ = fi12
    }
    p1 = (p1 >> ((p1) & 63))
    _ = p0
    _ = p1
    return uint64(p1)
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    var v3 int8 = -34
    var v4 uint64 = 1
    var v5 string = (("" + ("" + "a")) + (("foo bar" + ("foo bar" + "\n\t")) + "foo bar"))
    switch -9223372036854775808 {
    case 9223372036854775807:
        switch ((((("a" + (v5 + ("foo bar" + ((v5 + (v5 + v5)) + (v5 + v5))))) + v5) + ("" + v5)) + (v5 + v5)) + v5) {
        case "hello":
            var v15 bool = (0.0 <= ((1e10 / ((1.5 * (0.0 - (float64(((float32((float64(((0.1 / 1.0) + (0.1 / float32(0.1)))) - 1e10)) + (3.14159 - 3.14159)) - 0.1)) * 0.1))) * 0.1)) - 0.0))
            _ = v15
        case "a":
            for fi16 := 0; fi16 < 18; fi16++ {
                v5 = v5
                v17 := GenS1{F0: 2147483648, F1: (v3 / (((v3 & 42)) | 1))}
                if len(v1) > 0 {
                    v1[0] = v5
                }
                _ = fi16
                _ = v17
            }
        }
    case 42:
        v18 := GenS1{F0: 831, F1: (v3 % ((100) | 1))}
        _ = v18
    case -9223372036854775808:
        _ = len(v2)
    default:
        var v19 uint16 = 65535
        _ = v19
    }
    delete(v2, "hello")
    if false {
        switch 128 {
        case 128:
            var v20 uint64 = (^(v4))
            _ = v20
        default:
            var v21 float32 = 3.14159
            _ = v21
        }
        fmt.Println(v4, v5, v5)
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        if true {
            v23 := GenS1{F0: 1, F1: (v3 * 100)}
            v5 = (v5 + v5)
            delete(v2, v5)
            _ = v23
        } else {
            for fi24 := 0; fi24 < 18; fi24++ {
                var v25 int = (fi24 * (fi22 * fi24))
                _ = fi24
                _ = v25
            }
            v3 = -1
        }
        var v26 int = fi22
        fmt.Println(fi22)
        _ = fi22
        _ = v26
    }
    switch (v5 + (v5 + v5)) {
    case "a":
        v3 = (v3 + v3)
    case "\n\t":
        v5 = v5
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v27 := GenS2{F0: 919, F1: 806, F2: 130}
        _ = v27
    }
    var v28 string = (v5 + "\n\t")
    v3 += v3
    v29 := v2[v5]
    for fi30 := 0; fi30 < 3; fi30++ {
        var v31 int32 = -83
        _ = fi30
        _ = v31
    }
    v28 = ("hello" + v5)
    for fi32 := 0; fi32 < 16; fi32++ {
        if v29 {
            var v33 uint64 = (v4 << ((2) & 63))
            var v34 uint = 364
            _ = v33
            _ = v34
        }
        _ = fi32
    }
    var v35 uint = 0
    v36 := GenS3{F0: (("hello" + ((v28 + "\n\t") + v5)) + "a"), F1: v35, F2: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
