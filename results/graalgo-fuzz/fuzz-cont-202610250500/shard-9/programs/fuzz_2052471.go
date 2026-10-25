package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 bool }
func genFn0(p0 uint64, p1 int32) uint {
    p1 = 42
    v7 := GenS1{F0: 127, F1: 65535}
    _ = p0
    _ = p1
    _ = v7
    return 444
}

func main() {
    var v1 uint8 = 144
    var v2 uint16 = 65535
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    var v5 int32 = -99
    v6 := []string{((("a" + "foo bar") + ((("a" + "foo bar") + ("\n\t" + "")) + "hello")) + "foo bar"), "foo bar", ((((("\n\t" + "foo bar") + "foo bar") + "\n\t") + "") + ("a" + (("foo bar" + "\n\t") + "")))}
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int32 = (v5 + (v5 >> ((genFn0(2, v5)) & 63)))
        v10 := GenS1{F0: (fi8 + fi8), F1: 1}
        v1 = v1
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 32767, F1: "\n\t"}
    v5 = -1
    if (128 < v1) {
        v5 = v5
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 uint64 = 1
            v14 := GenS3{F0: fi12}
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        var v15 int32 = 2147483647
        _ = v15
    }
    v2 = (^(v2))
    v4[128] = true
    if true {
        v5 = 469
        for fi16 := 0; fi16 < 18; fi16++ {
            fmt.Println(fi16, v5, fi16)
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 float32 = 1.5
                if true {
                    v1 -= 222
                } else {
                    if true {
                        var v19 int16 = 42
                        var v20 uint16 = v2
                        _ = v19
                        _ = v20
                    } else {
                        v21 := GenS4{F0: 2147483647, F1: true}
                        _ = v21
                    }
                    var v22 int64 = -9223372036854775808
                    v1 = v1
                    _ = v22
                }
                _ = fi17
                _ = v18
            }
            switch (fi16 * fi16) {
            case 32768:
                fmt.Println(fi16)
            case 1:
                fmt.Println(fi16, v2, v2)
            case 650:
                _ = len(v6)
            }
            _ = fi16
        }
        v23 := GenS1{F0: 42, F1: 325}
        _ = v23
    }
    var v24 int16 = -32768
    v4[4294967295] = (9223372036854775807 != 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
