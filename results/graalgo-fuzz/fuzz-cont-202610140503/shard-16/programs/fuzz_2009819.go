package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64; F1 int8 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[string]bool)
    var v3 int16 = -1
    var v4 int = int(v1)
    v4 += v4
    var v5 uint8 = 2
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 4294967295, F1: 414, F2: 0}
        _ = fi6
        _ = v7
    }
    if ((v3 / ((900) | 1)) >= (v3 % (((v3 | -32768)) | 1))) {
        var v8 string = "\n\t"
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            if ((v1 <= v1) && false) {
                var v11 string = ("" + "a")
                _ = v11
            }
            v5 -= v5
            v12 := GenS1{F0: 1, F1: 241, F2: 127}
            _ = fi10
            _ = v12
        }
        _ = len(v2)
        _ = fi9
    }
    v13 := v2["a"]
    v3 = 0
    if true {
        v14 := GenS1{F0: 687, F1: 9223372036854775807, F2: 1}
        v15 := GenS1{F0: 18446744073709551615, F1: 664, F2: 42}
        _ = v14
        _ = v15
    } else {
        v1 = (v1 % (((v1 + v1)) | 1))
        var v16 uint64 = 42
        v17 := GenS2{F0: v16, F1: -97}
        _ = v16
        _ = v17
    }
    v4 = v4
    v13 = v13
    fmt.Println(v3)
    switch (^(v4)) {
    case 14:
        for fi18 := 0; fi18 < 6; fi18++ {
            v1 = v1
            var v19 int = v4
            _ = fi18
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS2{F0: 2, F1: -1}
            v2[("foo bar" + "\n\t")] = (704 == v3)
            _ = fi20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        if false {
            v13 = true
            var v23 int32 = 1
            _ = v23
        } else {
            var v24 uint32 = (^(v1))
            for fi25 := 0; fi25 < 9; fi25++ {
                v26 := GenS1{F0: 383, F1: -973, F2: 127}
                switch (fi22 % (((fi22 / ((v4) | 1))) | 1)) {
                case 960:
                    var v27 float64 = -2.25
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
            v3 -= (v3 + v3)
            _ = v24
        }
        delete(v2, "hello")
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
