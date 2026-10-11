package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 int32 }
func genFn0() uint32 {
    var v5 uint8 = 0
    var v6 bool = ((true && (((true && true) && (-128 >= -17)) || (!(-128 >= -128)))) && false)
    v6 = (!((true || true) || (!((!(2 < 18446744073709551615)) && (1.5 != 3.14159)))))
    var v7 uint64 = 0
    v7 += (v7 | v7)
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int8 = 42
    var v3 uint8 = 19
    v4 := make(map[int]bool)
    fmt.Println(v3, v1)
    v1 += v1
    if true {
        v8 := []int{-727, 868, int(v3), 2147483647}
        v2 = v2
        _ = v8
    } else {
        switch -542 {
        case 32767:
            for fi9 := 0; fi9 < 12; fi9++ {
                v2 -= v2
                switch ("\n\t" + "hello") {
                case "foo bar":
                    v3 = (v3 ^ 238)
                case "":
                    v10 := GenS1{F0: 255, F1: "hello", F2: (v3 / ((v3) | 1))}
                    _ = v10
                }
                var v11 bool = true
                _ = fi9
                _ = v11
            }
        case 9223372036854775807:
            v3 -= v3
        case 32768:
            v2 += (v2 / (((-(v2))) | 1))
        }
        fmt.Println(v3, v1)
    }
    v12 := []int{-425, 128}
    if false {
        fmt.Println(v1, v2, v2)
        v1 = v1
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            switch fi13 {
            case 908:
                var v14 int32 = 565
                _ = v14
            default:
                v15 := GenS1{F0: 255, F1: "", F2: (v3 | (v3 / ((1) | 1)))}
                _ = v15
            }
            var v16 int = (fi13 * 2147483647)
            if true {
                fmt.Println(v3, v3, v1)
            }
            _ = fi13
            _ = v16
        }
    }
    v3 = v3
    fmt.Println(v3, v2, v3)
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int16 = -1
            v1 = (^(v1))
            v2 -= v2
            _ = fi18
            _ = v19
        }
        delete(v4, 1)
        v1 = v1
        _ = fi17
    }
    v20 := GenS1{F0: 42, F1: (("" + "a") + "a"), F2: v3}
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 13; fi22++ {
            _ = len(v12)
            var v23 uint = 18446744073709551615
            _ = fi22
            _ = v23
        }
        var v24 float64 = (3.14159 - -2.25)
        if true {
            v25 := GenS2{F0: 42}
            _ = v25
        }
        _ = fi21
        _ = v24
    }
    var v26 uint32 = 2147483647
    var v27 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
