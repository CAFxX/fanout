package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 int32 = 0
    var v2 int32 = v1
    var v3 uint64 = 1
    var v4 uint16 = 65535
    var v5 uint = 255
    v6 := []uint32{469, 0}
    var v7 uint16 = (^(v4))
    v3 = (^(v3))
    v8 := GenS1{F0: 1000, F1: (v4 - 0), F2: 1}
    if (v1 > v2) {
        if false {
            var v9 string = ""
            _ = v9
        }
        switch "foo bar" {
        case "a":
            var v10 int = int(v7)
            _ = v10
        case "foo bar":
            v7 += (v4 * (v4 + (v4 / (((v7 % (((v7 % ((v7) | 1))) | 1))) | 1))))
        case "hello":
            var v11 int8 = 42
            _ = v11
        }
        _ = len(v6)
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 uint32 = 2147483647
        _ = fi12
        _ = v13
    }
    if true {
        var v14 float64 = (float64(1e10) + ((float64(0.1) / 1e10) * (3.14159 + 1.5)))
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := GenS1{F0: -1, F1: 42, F2: -9223372036854775808}
            if ((fi15 <= int(v5)) || (!(!true))) {
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 int64 = 42
                    _ = fi17
                    _ = v18
                }
                if false {
                    v1 = v1
                }
            }
            v2 = v2
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        v3 = (v3 << (((v5 ^ (v5 * (v5 & (v5 | (v5 ^ (v5 >> (((v5 - (v5 << (((v5 ^ (v5 ^ v5))) & 63)))) & 63)))))))) & 63))
    }
    if (-70 == 0) {
        if true {
            v19 := GenS1{F0: 43, F1: (^(v4)), F2: 1}
            _ = v19
        } else {
            v3 += v3
            v2 += 42
            var v20 int8 = 100
            _ = v20
        }
    }
    v21 := GenS2{F0: 42}
    v7 += v4
    var v22 float32 = 0.1
    var v23 float64 = -2.25
    fmt.Println(v23, v23, v4)
    if false {
        fmt.Println(v3, v1, v2)
    }
    if true {
        v24 := GenS3{F0: 2147483648, F1: 4294967295}
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        v2 = 459
        var v26 uint8 = 0
        _ = fi25
        _ = v26
    }
    var v27 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
