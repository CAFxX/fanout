package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() uint16 {
    var v6 int32 = 0
    fmt.Println(v6)
    fmt.Println(v6)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int64 = 63
        var v9 uint16 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v6
    return 0
}

func main() {
    var v1 uint64 = 42
    var v2 int = 32768
    var v3 int16 = int16(v2)
    v4 := make(map[string]string)
    var v5 int32 = -724
    if (700 == 886) {
        if (v5 < (v5 - v5)) {
            fmt.Println(v3, v1, v5)
            var v10 int64 = -1
            _ = v10
        }
        var v11 uint = 1
        _ = v11
    }
    var v12 string = "a"
    switch v12 {
    case "":
        for fi13 := 0; fi13 < 13; fi13++ {
            v12 = "a"
            _ = fi13
        }
    case "hello":
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := v4["foo bar"]
            if false {
                v1 += (v1 % ((v1) | 1))
                var v16 int16 = -22
                v2 += (^(fi14))
                _ = v16
            }
            _ = len(v4)
            _ = fi14
            _ = v15
        }
    case "foo bar":
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 3; fi18++ {
                v2 = fi18
                _ = fi18
            }
            _ = fi17
        }
    }
    v1 = (v1 + (v1 ^ v1))
    if (56 > 728) {
        v3 = 0
        var v19 int32 = (v5 ^ 1)
        if true {
            v1 += 2
            v19 = 0
            if (false && true) {
                var v20 uint = 42
                _ = v20
            } else {
                v19 -= (v5 % (((v19 & 399)) | 1))
            }
        }
        _ = v19
    } else {
        var v21 int32 = 42
        var v22 int8 = -128
        v23 := GenS1{F0: 1}
        _ = v21
        _ = v22
        _ = v23
    }
    delete(v4, v12)
    v24 := []int{(v2 - (v2 / ((v2) | 1))), v2}
    for fi25 := 0; fi25 < 9; fi25++ {
        v2 += fi25
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
