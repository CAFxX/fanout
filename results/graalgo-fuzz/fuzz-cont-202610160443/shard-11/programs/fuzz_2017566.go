package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 uint8 }
func genFn0(p0 bool, p1 int64) int {
    p1 -= p1
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            p0 = (110 <= 100)
            var v7 uint64 = 694
            _ = fi6
            _ = v7
        }
        switch "a" {
        case "foo bar":
            v8 := GenS1{F0: 1, F1: 42, F2: 2147483647}
            _ = v8
        case "hello":
            var v9 int = fi5
            _ = v9
        default:
            switch (("" + (("\n\t" + "foo bar") + "\n\t")) + "\n\t") {
            case "foo bar":
                var v10 bool = (((2 > 225) && (p0 || true)) || true)
                _ = v10
            case "a":
                p1 += -501
            default:
                fmt.Println(fi5)
            }
        }
        _ = fi5
    }
    v11 := GenS2{F0: p1, F1: 32}
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    var v1 bool = (42 > -9223372036854775808)
    var v2 uint16 = 0
    var v3 float64 = -2.25
    var v4 int32 = 0
    if true {
        var v12 uint32 = 2
        _ = v12
    } else {
        var v13 int16 = -719
        var v14 uint16 = (v2 & (v2 + (v2 | (v2 + v2))))
        v3 = (v3 - 1e10)
        _ = v13
        _ = v14
    }
    v15 := GenS3{F0: 116, F1: genFn0((2147483648 != 4294967295), 0), F2: 255}
    var v16 int8 = 100
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int64 = 1
        fmt.Println(v3, fi17, v2)
        switch 42 {
        case 482:
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 uint8 = 127
                var v21 int32 = v4
                _ = fi19
                _ = v20
                _ = v21
            }
        case 920:
            if ((v16 - (v16 << ((1) & 63))) >= int8(fi17)) {
                if true {
                    v22 := []int{-9223372036854775808}
                    v23 := GenS3{F0: 2147483647, F1: fi17, F2: 255}
                    fmt.Println(v3, fi17, v16)
                    _ = v22
                    _ = v23
                }
            }
        case 4294967295:
            v4 = v4
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v2, v1, v2)
    var v24 uint8 = 255
    v25 := []int{-141}
    switch -841 {
    case 1:
        var v26 int32 = (v4 ^ (v4 ^ -2147483648))
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
