package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 int16, p1 uint, p2 int32) int {
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: "", F1: 2}
        p1 += p1
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: "", F1: p1}
    var v7 uint32 = 397
    p2 = 341
    var v8 uint8 = 172
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float32 = 1.5
    v2 := []int{}
    var v3 string = ("a" + ("\n\t" + "foo bar"))
    if true {
        fmt.Println(v3)
        switch "hello" {
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = -564
            }
        default:
            if false {
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 bool = ((true && (1 <= 121)) || true)
                    var v11 uint16 = 2
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                for fi12 := 0; fi12 < 7; fi12++ {
                    v13 := GenS1{F0: v3, F1: 42}
                    var v14 int16 = -575
                    if false {
                        var v15 uint64 = 1
                        v16 := GenS2{F0: 161, F1: 0}
                        var v17 string = v3
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            } else {
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := GenS1{F0: "a", F1: 255}
                    var v20 string = "a"
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
        }
        v1 += v1
    } else {
        v3 = "hello"
    }
    if false {
        var v21 uint = 429
        _ = v21
    } else {
        v2 = append(v2, -771)
        v22 := GenS1{F0: ("foo bar" + ""), F1: 394}
        _ = v22
    }
    var v23 string = v3
    var v24 uint16 = 351
    v3 = v23
    if len(v2) > 0 {
        v2[0] = -1
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        _ = len(v2)
        for fi26 := 0; fi26 < 11; fi26++ {
            switch ("" + "hello") {
            case "foo bar":
                switch "hello" {
                case "a":
                    v2 = append(v2, (fi25 * -949))
                }
            case "":
                var v27 int64 = 1
                _ = v27
            case "hello":
                var v28 uint8 = 2
                _ = v28
            default:
                switch ((v23 + ((v23 + ((((v3 + "foo bar") + v23) + v23) + ((v23 + v3) + "\n\t"))) + "\n\t")) + v3) {
                case "\n\t":
                    v29 := GenS1{F0: "\n\t", F1: 0}
                    _ = v29
                default:
                    switch fi25 {
                    case 4294967295:
                        v24 += (^(v24))
                    case 1:
                        var v30 uint32 = 2147483648
                        _ = v30
                    }
                }
            }
            fmt.Println(v1, fi26, v1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi26
        }
        _ = fi25
    }
    var v31 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
}
