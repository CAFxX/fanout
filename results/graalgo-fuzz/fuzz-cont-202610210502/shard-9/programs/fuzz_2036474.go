package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
func genFn0() int64 {
    var v7 bool = true
    var v8 string = "a"
    v7 = true
    var v9 uint64 = 979
    fmt.Println(v9, v9)
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int16) int8 {
    fmt.Println(p0)
    v10 := GenS1{F0: 0, F1: (^(p0)), F2: (("" + "\n\t") + (("a" + "") + "hello"))}
    p0 -= 32767
    var v11 float64 = 3.14159
    var v12 int = -1
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 41
}

func main() {
    var v1 int64 = 1
    v2 := []uint32{}
    var v3 bool = ((0.1 / (3.14159 - 0.0)) <= 0.0)
    v4 := make(map[int]bool)
    var v5 int32 = 42
    var v6 int8 = -35
    if v3 {
        v13 := GenS1{F0: 582, F1: 1000, F2: (((("foo bar" + "") + ("\n\t" + "\n\t")) + ("foo bar" + "hello")) + "hello")}
        switch -967 {
        case -654:
            v6 = (^(v6))
        case -391:
            var v14 float64 = 3.14159
            _ = v14
        case 4294967295:
            v5 -= v5
        default:
            v15 := GenS1{F0: 65535, F1: 1, F2: "a"}
            _ = v15
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            switch ((("a" + "foo bar") + ("foo bar" + "foo bar")) + ("\n\t" + "foo bar")) {
            case "foo bar":
                for fi17 := 0; fi17 < 6; fi17++ {
                    if (true && (65535 < 4294967295)) {
                        var v18 string = "a"
                        _ = v18
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi17
                }
            case "":
                var v19 int = 1
                _ = v19
            }
            switch "hello" {
            case "hello":
                v20 := v4[fi16]
                _ = v20
            }
            _ = fi16
        }
        _ = v13
    } else {
        for fi21 := 0; fi21 < 14; fi21++ {
            switch 2147483647 {
            case 703:
                v3 = ((-9223372036854775808 > fi21) && (!v3))
            }
            _ = fi21
        }
        v22 := GenS1{F0: 255, F1: 1000, F2: "a"}
        _ = v22
    }
    if ((2 <= 132) && true) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v23 := GenS2{F0: 631}
    v3 = false
    var v24 uint64 = 291
    v6 -= (v6 | v6)
    fmt.Println(v1, v1, v6)
    var v25 bool = (!false)
    v1 = (v1 ^ (v1 + v1))
    if (v24 == (v24 / (((v24 << ((46) & 63))) | 1))) {
        switch -1 {
        case -9223372036854775808:
            v1 += (v1 - v1)
        case 32767:
            var v26 uint64 = v24
            _ = v26
        default:
            v27 := GenS2{F0: 132}
            _ = v27
        }
        var v28 float64 = 3.14159
        var v29 uint16 = 65535
        _ = v28
        _ = v29
    }
    var v30 string = "foo bar"
    v31 := GenS3{F0: -666}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
