package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int32; F1 int }
func genFn0(p0 uint16, p1 uint32) int8 {
    fmt.Println(p1, p0)
    var v4 uint = 4294967295
    v5 := []int{-9223372036854775808, 32768, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1() bool {
    v6 := []int{9223372036854775807}
    var v7 uint = 18446744073709551615
    if len(v6) > 0 {
        v6[0] = 32768
    }
    if false {
        fmt.Println(v7)
        for fi8 := 0; fi8 < 7; fi8++ {
            fmt.Println(fi8, fi8)
            var v9 int = (fi8 >> ((v7) & 63))
            _ = fi8
            _ = v9
        }
    }
    var v10 uint32 = 42
    _ = v10
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 int64 = 271
    var v2 uint8 = 2
    v3 := make(map[string]uint32)
    v1 += (v1 << ((1) & 63))
    v11 := []int{-9223372036854775808, 368}
    v12 := GenS1{F0: 42, F1: 127, F2: 0}
    if false {
        v2 = v2
        var v13 uint = 4294967295
        fmt.Println(v1)
        _ = v13
    } else {
        switch ("" + "\n\t") {
        case "a":
            if ((1 < 65535) || (!false)) {
                var v14 int = 431
                v15 := GenS2{F0: 884}
                _ = v14
                _ = v15
            } else {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
        case "foo bar":
            fmt.Println(v1)
        }
        _ = len(v3)
        v16 := GenS3{F0: 32767, F1: -2147483648}
        _ = v16
    }
    v17 := GenS4{F0: -1, F1: -867}
    delete(v3, "a")
    v1 += (v1 % ((v1) | 1))
    v18 := GenS5{F0: 42, F1: 32768}
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            _ = len(v3)
            _ = fi20
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            fmt.Println(fi21)
            var v22 int32 = 2147483647
            _ = fi21
            _ = v22
        }
        fmt.Println(v2, v1, fi19)
        _ = fi19
    }
    v2 = v2
    var v23 string = ("foo bar" + "foo bar")
    if true {
        if (32768 >= 32767) {
            for fi24 := 0; fi24 < 3; fi24++ {
                for fi25 := 0; fi25 < 13; fi25++ {
                    var v26 float32 = 1.5
                    if true {
                        var v27 int32 = 341
                        _ = v27
                    }
                    _ = fi25
                    _ = v26
                }
                v11 = append(v11, (fi24 & fi24))
                _ = fi24
            }
            _ = len(v11)
        }
        v28 := GenS2{F0: 2}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
