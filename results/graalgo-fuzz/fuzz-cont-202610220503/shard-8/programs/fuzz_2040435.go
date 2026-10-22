package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int16, p1 uint16) uint32 {
    var v4 float64 = 1.5
    var v5 int = -736
    var v6 bool = false
    if (!((p1 | (p1 & p1)) == (p1 - (p1 % ((p1) | 1))))) {
        var v7 int = v5
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := []string{("foo bar" + (("\n\t" + ("" + "a")) + ("" + "foo bar"))), "a", "hello"}
    var v2 int64 = 1
    v3 := []string{("hello" + ("a" + "foo bar")), ("" + ""), ("\n\t" + "\n\t")}
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            if true {
                _ = len(v3)
                var v10 uint16 = 361
                var v11 int64 = (v2 | (v2 - (v2 * 1)))
                v2 = v11
                _ = v10
                _ = v11
            }
            _ = fi9
        }
        v2 += (v2 & (-(v2)))
        if (1 == 255) {
            v2 += v2
            if (472 != 749) {
                v12 := GenS1{F0: 4294967295, F1: genFn0(1000, 32767)}
                _ = v12
            } else {
                for fi13 := 0; fi13 < 7; fi13++ {
                    var v14 int8 = 1
                    _ = fi13
                    _ = v14
                }
                if false {
                    _ = len(v1)
                    var v15 int16 = 42
                    _ = v15
                } else {
                    v16 := GenS1{F0: 65535, F1: 2}
                    v17 := GenS2{F0: 128}
                    fmt.Println(fi8, v2)
                    _ = v16
                    _ = v17
                }
            }
        }
        _ = fi8
    }
    v18 := GenS3{F0: 622}
    v2 -= (v2 % ((v2) | 1))
    v2 = 437
    fmt.Println(v2)
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 int64 = v2
        _ = fi19
        _ = v20
    }
    switch "foo bar" {
    case "\n\t":
        for fi21 := 0; fi21 < 11; fi21++ {
            v22 := GenS2{F0: 255}
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 uint16 = 839
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
    case "hello":
        v25 := GenS2{F0: 175}
        _ = v25
    default:
        var v26 int8 = -128
        _ = v26
    }
    fmt.Println(v2)
    var v27 uint16 = 227
    var v28 int16 = 0
    switch (("a" + "foo bar") + ("hello" + "a")) {
    case "":
        var v29 int16 = (v28 & (v28 ^ v28))
        _ = v29
    }
    var v30 uint16 = (v27 << ((2) & 63))
    var v31 uint64 = 42
    v27 -= v30
    if true {
        var v32 uint64 = v31
        if (32767 >= (v28 / (((^(v28))) | 1))) {
            v33 := GenS1{F0: 4294967295, F1: 465}
            _ = v33
        }
        v2 = v2
        _ = v32
    }
    fmt.Println(v31)
    v2 = (v2 - v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
}
