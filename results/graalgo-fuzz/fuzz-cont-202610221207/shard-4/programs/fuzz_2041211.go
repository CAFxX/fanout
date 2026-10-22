package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int16 = -725
    v2 := []uint32{}
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 42
    var v5 int64 = -9223372036854775808
    var v6 uint = 4294967295
    var v7 float64 = 3.14159
    if ((v5 | -9223372036854775808) < (v5 - (v5 | v5))) {
        var v8 int8 = 100
        fmt.Println(v5, v7)
        _ = v8
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v1 = -32768
        v5 = v5
        _ = fi9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 string = "foo bar"
            var v13 int = int(v5)
            var v14 bool = (-1 < 42)
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        if false {
            var v15 uint64 = 18446744073709551615
            v15 = 42
            _ = v15
        } else {
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 int64 = v5
                _ = fi16
                _ = v17
            }
        }
        _ = fi10
    }
    var v18 int32 = -2147483648
    v19 := GenS1{F0: (v4 + v4)}
    var v20 int32 = v18
    v4 = (^(v4))
    v2 = append(v2, (v4 % (((v4 % (((v4 & v4)) | 1))) | 1)))
    switch -1 {
    case 32767:
        fmt.Println(v1, v20)
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        if false {
            var v22 string = "\n\t"
            _ = v22
        } else {
            switch (-(fi21)) {
            case 122:
                fmt.Println(v1, v5, v6)
            default:
                v3[734] = 884
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint = 255
                v18 = v18
                _ = fi23
                _ = v24
            }
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
