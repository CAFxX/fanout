package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8, p1 int16) string {
    p1 -= (p1 - p1)
    fmt.Println(p0)
    var v7 string = "a"
    if ((-(p1)) >= (p1 & (^(p1)))) {
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v7
    return (("a" + (v7 + "foo bar")) + v7)
}

func main() {
    var v1 float32 = (-2.25 / float32((1.5 * 3.14159)))
    v2 := []int{}
    v3 := make(map[uint32]int)
    var v4 uint = 2
    var v5 uint32 = 65535
    v6 := make(map[string]bool)
    for fi8 := 0; fi8 < 6; fi8++ {
        _ = len(v3)
        fmt.Println(v1)
        _ = fi8
    }
    var v9 int = int(v5)
    if true {
        v5 = uint32(v4)
        v10 := v6["a"]
        switch (v9 >> ((864) & 63)) {
        case 80:
            var v11 float64 = 1e10
            _ = v11
        case 51:
            var v12 int8 = -128
            _ = v12
        case 762:
            _ = len(v3)
        default:
            v13 := GenS1{F0: 0, F1: ""}
            _ = v13
        }
        _ = v10
    } else {
        if (1 != 127) {
            v14 := GenS1{F0: 1, F1: ("foo bar" + (("foo bar" + "a") + "hello"))}
            _ = v14
        }
        fmt.Println(v4)
        var v15 uint16 = 1
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        if (17 >= -52) {
            fmt.Println(v4, fi16, v9)
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
            v17 := GenS1{F0: 127, F1: "hello"}
            _ = v17
        }
        _ = fi16
    }
    var v18 int32 = -988
    v1 = (v1 + v1)
    var v19 uint32 = v5
    if (747 > 9223372036854775807) {
        v20 := GenS2{F0: -1, F1: "a"}
        v21 := GenS3{F0: 418}
        _ = v20
        _ = v21
    } else {
        switch ("\n\t" + "\n\t") {
        case "a":
            _ = len(v2)
        }
        var v22 float32 = (v1 - v1)
        _ = v22
    }
    v19 = (v5 >> (((v4 & v4)) & 63))
    if true {
        var v23 uint64 = 4294967295
        _ = v23
    } else {
        for fi24 := 0; fi24 < 15; fi24++ {
            if false {
                var v25 uint16 = 1
                switch (fi24 / ((48) | 1)) {
                case 127:
                    var v26 float32 = v1
                    _ = v26
                case 128:
                    var v27 int8 = 107
                    _ = v27
                }
                _ = v25
            }
            _ = fi24
        }
    }
    var v28 uint32 = v5
    if true {
        var v29 uint = v4
        _ = v29
    }
    v9 -= v9
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
