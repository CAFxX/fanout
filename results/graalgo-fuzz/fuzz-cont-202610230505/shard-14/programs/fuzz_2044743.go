package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint64; F2 int16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 string) uint {
    p0 = p0
    fmt.Println(p0)
    if ((42 > -1) || (100 > -1)) {
        if true {
            switch 32767 {
            case 4294967295:
                var v5 uint32 = 4294967295
                _ = v5
            case 0:
                v6 := GenS1{F0: 42}
                _ = v6
            default:
                v7 := []int{1, -1, 9223372036854775807, 32767}
                _ = v7
            }
        } else {
            p0 = p0
            var v8 bool = false
            _ = v8
        }
        switch p0 {
        case "foo bar":
            fmt.Println(p0)
        case "hello":
            var v9 int8 = 100
            _ = v9
        case "a":
            if false {
                for fi10 := 0; fi10 < 2; fi10++ {
                    var v11 int16 = 1000
                    _ = fi10
                    _ = v11
                }
            }
        }
    } else {
        v12 := GenS1{F0: -351}
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint64 = 711
            v15 := GenS2{F0: 42, F1: v14, F2: 1000}
            var v16 int16 = 0
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        p0 = (p0 + ((p0 + "foo bar") + ""))
        _ = v12
    }
    var v17 int64 = 9223372036854775807
    _ = p0
    _ = v17
    return 187
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 182
    var v3 int64 = 926
    var v4 uint64 = 0
    var v18 int = 32767
    var v19 int16 = 32767
    v19 += (^(v19))
    var v20 bool = true
    v4 = v4
    if true {
        v21 := GenS2{F0: (-(v3)), F1: (v4 - v4), F2: v19}
        v22 := GenS3{F0: 127}
        _ = v21
        _ = v22
    } else {
        _ = len(v1)
        v2 += (v2 | 1)
        v2 -= v2
    }
    var v23 uint32 = 65535
    v20 = v20
    v24 := GenS1{F0: (v3 | 266)}
    if v20 {
        var v25 bool = v20
        _ = v25
    } else {
        var v26 uint16 = 65535
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
