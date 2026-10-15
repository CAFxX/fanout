package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 string; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int8 }

func main() {
    var v1 uint16 = 544
    var v2 int32 = 1
    var v3 bool = true
    var v4 uint = 0
    v5 := []string{"\n\t", "hello"}
    v6 := make(map[uint32]int)
    var v7 uint32 = 2
    if ("foo bar" >= "\n\t") {
        switch -1 {
        case 32767:
            var v8 float32 = ((1e10 * 1e10) * 1.5)
            _ = v8
        }
        var v9 float64 = -2.25
        v10 := GenS1{F0: (v2 ^ 0)}
        _ = v10
        _ = v9
    } else {
        var v11 string = "\n\t"
        _ = v11
    }
    var v12 int32 = v2
    v3 = v3
    v13 := GenS2{F0: 2147483647, F1: ("\n\t" + ("\n\t" + "foo bar")), F2: 0}
    var v14 float64 = 1e10
    var v15 int32 = v12
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS2{F0: 669, F1: ("a" + "foo bar"), F2: 42}
        var v18 float32 = 3.14159
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint = 65535
    var v20 string = (("\n\t" + ("\n\t" + ("" + "hello"))) + "hello")
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            v23 := GenS3{F0: v7}
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    if v3 {
        var v24 bool = false
        _ = v24
    } else {
        var v25 uint32 = (^(v7))
        for fi26 := 0; fi26 < 3; fi26++ {
            switch fi26 {
            case -1:
                fmt.Println(v12)
            case -2147483648:
                v14 -= (v14 * float64((1.5 / 1.0)))
            case 32768:
                v27 := GenS4{F0: (v4 << ((v4) & 63)), F1: -128}
                _ = v27
            }
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            v15 += (v12 << ((uint(v7)) & 63))
            _ = fi26
        }
        _ = v25
    }
    var v28 uint = (v19 | v4)
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 13; fi30++ {
            delete(v6, 2147483648)
            _ = fi30
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
