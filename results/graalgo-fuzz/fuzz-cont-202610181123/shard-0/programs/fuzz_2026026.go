package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint = 75
    var v2 uint64 = 1
    var v3 string = "a"
    v4 := []int{1, 4294967295, -929}
    var v5 int = -9223372036854775808
    var v6 int32 = 633
    var v7 int8 = 0
    if len(v4) > 0 {
        v4[0] = v5
    }
    if len(v4) > 0 {
        v4[0] = (v5 ^ 9223372036854775807)
    }
    if (v5 > (^(v5))) {
        if (!true) {
            if false {
                fmt.Println(v7, v7, v2)
                v2 = (v2 % (((^(v2))) | 1))
            } else {
                var v8 int = -1
                v8 = v5
                _ = v8
            }
        } else {
            var v9 int = (v5 / ((502) | 1))
            _ = v9
        }
        if (1e10 <= -2.25) {
            var v10 string = v3
            _ = v10
        } else {
            switch "\n\t" {
            case "hello":
                var v11 uint8 = 1
                _ = v11
            }
        }
    } else {
        v7 = int8(v1)
        var v12 int32 = (v6 * (v6 - 1))
        _ = v12
    }
    v13 := GenS1{F0: 32767, F1: (v1 >> (((v1 >> ((1) & 63))) & 63)), F2: 961}
    if true {
        v14 := GenS2{F0: (^(v2)), F1: -5}
        var v15 uint = (v1 << (((^(v1))) & 63))
        _ = v14
        _ = v15
    } else {
        _ = len(v4)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v3 = v3
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        if len(v4) > 0 {
            v4[0] = -917
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS3{F0: (v1 % (((v1 >> ((v1) & 63))) | 1))}
            _ = fi17
            _ = v18
        }
        fmt.Println(v3, fi16, v7)
        _ = fi16
    }
    var v19 uint32 = 2
    var v20 int64 = -948
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
