package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 int32) uint16 {
    v6 := GenS1{F0: 42, F1: -474}
    p0 = -1
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 uint64 = 42
    var v2 uint16 = 932
    var v3 bool = false
    var v4 string = ""
    var v5 int32 = -2147483648
    v5 = -809
    v2 = 198
    v7 := []int{2147483647, -204}
    fmt.Println(v1)
    fmt.Println(v3, v3)
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS2{F0: -9223372036854775808, F1: 4294967295}
            var v11 string = "hello"
            var v12 uint = 298
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    var v13 uint = 0
    var v14 float64 = 1e10
    v13 -= v13
    switch v4 {
    case "foo bar":
        fmt.Println(v14, v14)
    }
    v1 = v1
    fmt.Println(v13, v3, v2)
    v4 = ""
    var v15 int16 = 0
    if v3 {
        v3 = (4294967295 < 92)
        switch -958 {
        case -9223372036854775808:
            v3 = (!v3)
        case -85:
            for fi16 := 0; fi16 < 14; fi16++ {
                v14 = float64(-2.25)
                _ = fi16
            }
        default:
            v5 = (v5 >> (((v13 ^ 42)) & 63))
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 int8 = 127
            _ = fi17
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 18; fi19++ {
            switch "foo bar" {
            case "":
                v20 := GenS3{F0: v3, F1: 65535}
                _ = v20
            case "a":
                var v21 int8 = 127
                _ = v21
            default:
                v1 += 0
            }
            if len(v7) > 0 {
                v7[0] = (fi19 ^ 734)
            }
            _ = fi19
        }
        v22 := GenS3{F0: v3, F1: 664}
        _ = v22
    }
    v5 = v5
    v23 := GenS2{F0: -596, F1: 2147483648}
    var v24 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
