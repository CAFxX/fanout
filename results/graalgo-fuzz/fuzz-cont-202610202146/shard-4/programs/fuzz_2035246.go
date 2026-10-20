package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 int32 = 1
    v2 := make(map[string]bool)
    var v3 uint32 = 4294967295
    fmt.Println(v1)
    for fi4 := 0; fi4 < 9; fi4++ {
        v1 = v1
        _ = fi4
    }
    var v5 int64 = -1
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(v5, v5)
        fmt.Println(v1, v3, v5)
        _ = fi6
    }
    var v7 string = ("hello" + ((("a" + "a") + ("hello" + "")) + "foo bar"))
    v1 = v1
    var v8 int32 = 42
    v2["a"] = true
    v9 := GenS1{F0: (v5 << ((527) & 63)), F1: 4294967295}
    v10 := GenS1{F0: 1, F1: 64}
    if true {
        for fi11 := 0; fi11 < 11; fi11++ {
            if (!false) {
                if true {
                    _ = len(v2)
                    v3 = (v3 % (((v3 - (v3 & 2))) | 1))
                    var v12 int32 = (v1 >> ((4294967295) & 63))
                    _ = v12
                }
                switch fi11 {
                case 128:
                    v13 := GenS2{F0: v3, F1: 245}
                    _ = v13
                case 32768:
                    for fi14 := 0; fi14 < 20; fi14++ {
                        switch int(v8) {
                        case -522:
                            fmt.Println(v8)
                        }
                        var v15 int16 = 821
                        _ = fi14
                        _ = v15
                    }
                default:
                    var v16 int16 = 1
                    _ = v16
                }
            }
            v3 = v3
            _ = fi11
        }
    }
    fmt.Println(v1)
    for fi17 := 0; fi17 < 5; fi17++ {
        if (0 < 18446744073709551615) {
            var v18 float32 = 1.5
            _ = v18
        } else {
            var v19 uint8 = 2
            var v20 int8 = -42
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    v21 := []int{}
    switch 9223372036854775807 {
    case -9223372036854775808:
        if false {
            if (false || false) {
                for fi22 := 0; fi22 < 19; fi22++ {
                    v23 := v2["\n\t"]
                    _ = fi22
                    _ = v23
                }
                for fi24 := 0; fi24 < 17; fi24++ {
                    v25 := GenS1{F0: v5, F1: 42}
                    v7 = "foo bar"
                    _ = fi24
                    _ = v25
                }
            }
            v3 += v3
        } else {
            var v26 int8 = 42
            _ = v26
        }
    default:
        for fi27 := 0; fi27 < 14; fi27++ {
            v28 := v2[v7]
            _ = fi27
            _ = v28
        }
    }
    v21 = append(v21, -624)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
