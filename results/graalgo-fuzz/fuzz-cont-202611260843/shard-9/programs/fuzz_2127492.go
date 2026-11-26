package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 13
    v3 := []string{"\n\t"}
    for fi4 := 0; fi4 < 17; fi4++ {
        v2 = 2147483648
        if false {
            v3 = append(v3, ("\n\t" + ("a" + ("\n\t" + (("foo bar" + (("" + "a") + "")) + ("" + "\n\t"))))))
            for fi5 := 0; fi5 < 13; fi5++ {
                if true {
                    var v6 int32 = 457
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    delete(v1, 42)
                    v7 := v1[v2]
                    _ = v6
                    _ = v7
                }
                _ = fi5
            }
        }
        var v8 float64 = float64(1.5)
        _ = fi4
        _ = v8
    }
    switch ("foo bar" + "hello") {
    case "foo bar":
        v2 += v2
    case "\n\t":
        if (949 > -1) {
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 int16 = 1
                _ = fi9
                _ = v10
            }
        }
    default:
        v11 := GenS1{F0: false}
        _ = v11
    }
    if false {
        var v12 int = -9223372036854775808
        v12 += -440
        v12 -= v12
        _ = v12
    } else {
        switch 42 {
        case 4294967295:
            fmt.Println(v2)
        case 9223372036854775807:
            switch "foo bar" {
            case "":
                for fi13 := 0; fi13 < 8; fi13++ {
                    v2 = v2
                    _ = fi13
                }
            }
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            switch (fi14 & 313) {
            case 2147483647:
                v15 := GenS1{F0: true}
                _ = v15
            case 32767:
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 bool = false
                    _ = fi16
                    _ = v17
                }
            case -1:
                if len(v3) > 0 {
                    v3[0] = (("hello" + (("a" + "a") + ("foo bar" + ("foo bar" + "")))) + "foo bar")
                }
            default:
                fmt.Println(v2, v2)
            }
            v2 = (v2 << ((745) & 63))
            _ = fi14
        }
    }
    v2 = (v2 ^ v2)
    v2 -= 103
    var v18 float64 = 1e10
    if (-2147483648 < -2147483648) {
        var v19 int8 = 127
        _ = v19
    }
    switch 4294967295 {
    case 348:
        var v20 uint16 = 255
        _ = v20
    case 871:
        v18 = 0.0
    case -421:
        if (4294967295 > 4294967295) {
            _ = len(v3)
        }
    default:
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 float32 = 0.1
            v23 := GenS2{F0: 78}
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v24 := GenS2{F0: -109}
    fmt.Println(v2, v2)
    switch 9223372036854775807 {
    case 128:
        if (!(2 < 404)) {
            v25 := GenS2{F0: 46}
            _ = v25
        }
    case 42:
        var v26 uint64 = 222
        _ = v26
    }
    var v27 uint64 = 4294967295
    switch "hello" {
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        if (175 != 0) {
            var v28 float64 = (((float64(1.5) / v18) / ((v18 / (v18 / (1e10 / ((v18 / 1e10) / (3.14159 + (v18 + v18)))))) - 1.5)) + 1e10)
            v2 = 0
            fmt.Println(v2, v27)
            _ = v28
        } else {
            v18 -= v18
            if true {
                for fi29 := 0; fi29 < 10; fi29++ {
                    for fi30 := 0; fi30 < 1; fi30++ {
                        var v31 int64 = 42
                        _ = fi30
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 18; fi32++ {
                        fmt.Println(fi32)
                        v33 := GenS2{F0: 1}
                        _ = fi32
                        _ = v33
                    }
                    v2 -= (v2 << ((1) & 63))
                    _ = fi29
                }
                v34 := GenS3{F0: 709, F1: 1}
                _ = v34
            }
        }
    case "foo bar":
        switch "foo bar" {
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = "a"
            }
        case "a":
            fmt.Println(v18, v2, v2)
        case "":
            var v35 float64 = 0.1
            _ = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v27
    _ = v3
}
