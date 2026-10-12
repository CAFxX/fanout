package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int16; F1 int8; F2 int64 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int; F1 uint32; F2 bool }
type GenS6 struct { F0 int64 }

func main() {
    var v1 int16 = -1
    var v2 uint16 = 42
    var v3 bool = (false && (!false))
    var v4 int16 = (v1 % ((v1) | 1))
    var v5 string = ("a" + ("" + "a"))
    var v6 string = (((v5 + "foo bar") + ("a" + (((v5 + v5) + (((v5 + v5) + ("foo bar" + "a")) + v5)) + v5))) + "\n\t")
    v3 = (948 > 42)
    v7 := GenS1{F0: (2147483647 >= 127)}
    switch (v6 + v5) {
    case "":
        if v3 {
            var v8 uint = 745
            if (2 < 1) {
                v9 := GenS2{F0: (v2 | 64), F1: 42}
                v10 := GenS3{F0: (v1 ^ v1), F1: 0, F2: 1}
                v4 = -1
                _ = v10
                _ = v9
            } else {
                fmt.Println(v8)
            }
            _ = v8
        } else {
            v11 := GenS2{F0: v2, F1: 42}
            v12 := []int{32767, 128}
            _ = v11
            _ = v12
        }
    }
    var v13 float32 = ((3.14159 / ((-2.25 + (0.1 / 1.5)) - (1e10 - 1e10))) + ((float32(3.14159) / 1.0) + 0.0))
    if v3 {
        switch 32767 {
        case 128:
            var v14 uint64 = 2
            _ = v14
        case 311:
            var v15 uint = 65535
            _ = v15
        case 2147483647:
            v16 := GenS1{F0: (v2 == (v2 | (v2 & v2)))}
            _ = v16
        }
    } else {
        var v17 int = 213
        v18 := GenS3{F0: v4, F1: -118, F2: 42}
        _ = v17
        _ = v18
    }
    var v19 float64 = -2.25
    v5 = v6
    v20 := []int{32767, 127, 42, -9223372036854775808}
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS3{F0: (v4 % (((v1 % ((v4) | 1))) | 1)), F1: -128, F2: -9223372036854775808}
        var v23 int64 = 9223372036854775807
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS4{F0: 474, F1: 1}
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 uint32 = 42
        v5 = v6
        v27 := GenS3{F0: -70, F1: -39, F2: -9223372036854775808}
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 string = v5
        for fi30 := 0; fi30 < 17; fi30++ {
            v4 -= 32767
            _ = len(v20)
            _ = fi30
        }
        if false {
            fmt.Println(v5, v13, v19)
            if (!(35 == 0)) {
                v31 := GenS3{F0: (v1 / ((-32768) | 1)), F1: 54, F2: 1}
                _ = v31
            } else {
                for fi32 := 0; fi32 < 14; fi32++ {
                    v33 := GenS5{F0: fi32, F1: 2147483648, F2: v3}
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    v34 := GenS6{F0: -1}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
        } else {
            _ = len(v20)
            var v35 uint8 = 2
            _ = v35
        }
        _ = fi28
        _ = v29
    }
    if (-912 == -1) {
        var v36 int = 32767
        _ = v36
    }
    v37 := GenS2{F0: (^(v2)), F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
