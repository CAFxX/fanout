package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int16 = 746
    var v2 uint32 = 65535
    var v3 int32 = 350
    v4 := []int{9223372036854775807, 42}
    if false {
        v5 := GenS1{F0: 491, F1: -128}
        if (1.5 > (1e10 - 0.1)) {
            v6 := GenS1{F0: 633, F1: 0}
            if true {
                for fi7 := 0; fi7 < 6; fi7++ {
                    var v8 int = 129
                    _ = fi7
                    _ = v8
                }
                v1 = (v1 - v1)
            }
            _ = v6
        }
        _ = v5
    } else {
        v9 := GenS2{F0: -1}
        v10 := GenS2{F0: -1}
        _ = v10
        _ = v9
    }
    if (127 > 127) {
        v2 = v2
    }
    v2 = v2
    if (false || false) {
        var v11 string = ""
        _ = v11
    }
    v12 := GenS1{F0: 42, F1: 1}
    _ = len(v4)
    var v13 uint8 = 42
    var v14 uint16 = 32767
    v13 -= 42
    for fi15 := 0; fi15 < 12; fi15++ {
        v1 -= v1
        _ = len(v4)
        _ = fi15
    }
    v16 := GenS1{F0: (v14 ^ (v14 / (((v14 + v14)) | 1))), F1: 48}
    if true {
        for fi17 := 0; fi17 < 5; fi17++ {
            switch (fi17 | -2147483648) {
            case 32768:
                var v18 uint16 = v14
                _ = v18
            case 4294967295:
                v3 = v3
            }
            if false {
                v19 := GenS1{F0: v14, F1: 55}
                v14 = (^(v14))
                _ = v19
            }
            v2 = (v2 + 4294967295)
            _ = fi17
        }
        var v20 uint = 59
        v1 = -32768
        _ = v20
    } else {
        v21 := GenS2{F0: -9223372036854775808}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
