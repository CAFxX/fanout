package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 int8 }
func genFn0() uint64 {
    var v5 float32 = (((((0.1 / 1.5) * ((-2.25 - 1.5) / 0.1)) - ((0.1 - -2.25) * 1e10)) / 1.0) * 0.0)
    v5 += 0.0
    var v6 string = (("a" + "\n\t") + "\n\t")
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 string, p1 int) int32 {
    var v7 bool = (false || false)
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint32 = 4294967295
        v10 := GenS1{F0: 138, F1: -1, F2: -2}
        for fi11 := 0; fi11 < 5; fi11++ {
            p1 -= (p1 & fi8)
            if false {
                if (550 < 1000) {
                    switch fi8 {
                    case 2147483647:
                        if false {
                            var v12 uint32 = 2
                            v12 += v9
                            _ = v12
                        }
                    default:
                        v9 += 65535
                    }
                    switch -1 {
                    case -9223372036854775808:
                        var v13 int32 = -696
                        _ = v13
                    }
                }
                p0 = p0
                var v14 uint32 = v9
                _ = v14
            } else {
                v9 += 2147483647
            }
            switch "foo bar" {
            case "":
                var v15 float32 = float32(-2.25)
                _ = v15
            case "hello":
                switch p0 {
                case "hello":
                    p1 -= fi11
                case "\n\t":
                    v16 := []int{fi11, (fi8 - fi11)}
                    _ = v16
                case "":
                    fmt.Println(fi11)
                }
            default:
                v17 := GenS2{F0: -32768}
                _ = v17
            }
            _ = fi11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 = v7
    v18 := []int{p1}
    _ = p0
    _ = p1
    _ = v18
    _ = v7
    return 895
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    v19 := v2[689]
    switch (v19 + (v19 + (("" + "foo bar") + v19))) {
    case "":
        v1 = v1
    default:
        v20 := GenS1{F0: 127, F1: genFn1(((v19 + v19) + ("\n\t" + (("a" + "hello") + ("foo bar" + v19)))), 9223372036854775807), F2: v1}
        _ = v20
    }
    var v21 int16 = (v1 % (((v1 % (((v1 + v1)) | 1))) | 1))
    var v22 int = 534
    fmt.Println(v19)
    switch "\n\t" {
    case "hello":
        switch (v19 + v19) {
        case "hello":
            v23 := GenS2{F0: v21}
            _ = v23
        case "foo bar":
            for fi24 := 0; fi24 < 7; fi24++ {
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 uint8 = 138
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
        }
    case "foo bar":
        for fi27 := 0; fi27 < 3; fi27++ {
            fmt.Println(v22, fi27, v1)
            var v28 uint64 = 4294967295
            _ = fi27
            _ = v28
        }
    case "":
        v29 := GenS3{F0: true}
        _ = v29
    default:
        var v30 int = 127
        _ = v30
    }
    v31 := GenS4{F0: 104, F1: 119}
    _ = len(v3)
    v32 := []int{1, v22, (v22 & 128), 1}
    v21 = 240
    var v33 string = ("\n\t" + "hello")
    v21 -= v21
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
