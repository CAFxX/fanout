package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int }

func main() {
    v1 := []uint32{1, 292}
    var v2 uint = 0
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 float64 = 3.14159
    var v6 float64 = v5
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = 0
        _ = fi7
        _ = v8
    }
    v6 -= (float64(0.1) / (float64(1e10) / 1e10))
    v9 := GenS1{F0: -83, F1: 32768}
    var v10 string = ("foo bar" + "foo bar")
    var v11 uint = v2
    switch 127 {
    case -2147483648:
        if false {
            v12 := GenS1{F0: -2, F1: -549}
            var v13 int32 = 0
            _ = v12
            _ = v13
        }
    case 42:
        if true {
            var v14 int32 = 912
            var v15 uint16 = 65535
            _ = v14
            _ = v15
        } else {
            v10 = v10
        }
    default:
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int8 = 127
            var v18 uint32 = 2147483647
            delete(v3, (v10 + ((v10 + (v10 + ("" + v10))) + (v10 + v10))))
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    if true {
        if false {
            v19 := GenS1{F0: -1, F1: -9223372036854775808}
            v20 := GenS2{F0: v2, F1: 0}
            v21 := v4[v10]
            _ = v19
            _ = v20
            _ = v21
        }
    }
    if ((-167 == 42) || ((!(694 > -1)) || false)) {
        delete(v4, "a")
        delete(v3, "foo bar")
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi22
    }
    v1 = append(v1, 2147483648)
    v23 := GenS3{F0: 128}
    v11 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
