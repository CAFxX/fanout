package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 2147483648
    v3 := []uint32{4294967295}
    v4 := make(map[uint32]int)
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            var v7 uint32 = v2
            v5 -= v5
            if true {
                v8 := GenS1{F0: "foo bar"}
                _ = v8
            }
            _ = v7
        } else {
            for fi9 := 0; fi9 < 5; fi9++ {
                v1 += v1
                var v10 int16 = 1
                _ = fi9
                _ = v10
            }
            var v11 uint64 = 0
            switch (fi6 + 0) {
            case 128:
                var v12 uint64 = v1
                _ = v12
            }
            _ = v11
        }
        v13 := GenS1{F0: ((("a" + ("" + "")) + "a") + "a")}
        var v14 uint = 0
        _ = fi6
        _ = v13
        _ = v14
    }
    v1 = v1
    v1 = 122
    if ((127 >= 116) || false) {
        var v15 int = 2147483647
        var v16 float32 = -2.25
        if (((v5 % (((v5 - 0)) | 1)) >= v5) || false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v17 uint = 18446744073709551615
            _ = v17
        }
        _ = v15
        _ = v16
    } else {
        fmt.Println(v2)
        fmt.Println(v2, v1)
    }
    if true {
        v5 += (^(v5))
    }
    fmt.Println(v2)
    var v18 int64 = 42
    var v19 int64 = v18
    if (!((1e10 * float32((0.0 / 1e10))) < (0.0 - 0.1))) {
        v20 := GenS1{F0: "a"}
        _ = v20
    }
    var v21 uint16 = 796
    var v22 uint8 = 255
    v1 = (v1 % (((v1 - 18446744073709551615)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
