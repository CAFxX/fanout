package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
type GenS3 struct { F0 string; F1 uint16 }

func main() {
    var v1 int = -1
    var v2 uint32 = 2147483647
    v3 := make(map[string]int)
    var v4 int8 = -128
    var v5 int8 = v4
    _ = len(v3)
    v6 := GenS1{F0: 373}
    var v7 uint32 = (^(v2))
    var v8 int32 = 0
    for fi9 := 0; fi9 < 12; fi9++ {
        v4 = (v5 ^ (v4 + v4))
        _ = fi9
    }
    v1 -= v1
    v10 := []int{}
    if (((v2 < (^(v2))) || (("a" + "foo bar") < "a")) && false) {
        v5 += v5
        var v11 float64 = 0.0
        v5 = v4
        _ = v11
    } else {
        var v12 uint16 = 42
        v13 := GenS2{F0: v7, F1: true, F2: 918}
        var v14 bool = true
        _ = v12
        _ = v13
        _ = v14
    }
    fmt.Println(v8, v5, v5)
    if false {
        if true {
            delete(v3, ("" + ("a" + "\n\t")))
        }
        v5 -= v5
    }
    v15 := GenS3{F0: "", F1: 42}
    v4 = (v4 >> ((18446744073709551615) & 63))
    for fi16 := 0; fi16 < 8; fi16++ {
        v7 -= v7
        fmt.Println(v5, fi16)
        v2 = v7
        _ = fi16
    }
    _ = len(v10)
    v2 = (v7 << ((798) & 63))
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 float32 = 3.14159
        for fi19 := 0; fi19 < 10; fi19++ {
            switch 127 {
            case 2147483647:
                v7 = (^(v2))
            case 32767:
                switch "\n\t" {
                case "":
                    var v20 int = fi19
                    _ = v20
                case "foo bar":
                    var v21 uint32 = v2
                    _ = v21
                case "hello":
                    var v22 bool = (false && true)
                    _ = v22
                }
            default:
                if (255 > 141) {
                    delete(v3, "a")
                    fmt.Println(v1, v5)
                    var v23 string = ((("a" + ("foo bar" + "hello")) + ("\n\t" + "foo bar")) + "\n\t")
                    var v24 uint = 2
                    _ = v23
                    _ = v24
                } else {
                    v25 := GenS2{F0: (v7 - v2), F1: false, F2: 32767}
                    var v26 uint64 = 0
                    _ = v25
                    _ = v26
                }
            }
            _ = fi19
        }
        if false {
            var v27 int32 = (v8 & (v8 | v8))
            var v28 float64 = 3.14159
            _ = v27
            _ = v28
        } else {
            if true {
                v18 = ((v18 / (float32(0.1) / ((((v18 / 0.1) + 0.0) * (v18 - v18)) / (-2.25 * 1e10)))) + (1.5 * v18))
            }
            if true {
                fmt.Println(v1, v7)
            }
            var v29 int16 = 341
            _ = v29
        }
        _ = fi17
        _ = v18
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        for fi31 := 0; fi31 < 19; fi31++ {
            v32 := GenS1{F0: 32767}
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 uint16 = 864
            var v35 int64 = 42
            var v36 uint16 = v34
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        switch v1 {
        case 4294967295:
            v37 := v3["hello"]
            _ = v37
        default:
            var v38 int64 = 0
            _ = v38
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
