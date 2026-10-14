package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 string }
func genFn0(p0 int32, p1 uint8, p2 int16) uint64 {
    var v5 uint16 = 1
    switch "" {
    case "a":
        var v6 uint64 = 0
        _ = v6
    case "hello":
        var v7 string = (("hello" + ((("" + "") + "") + "")) + (("foo bar" + "a") + "foo bar"))
        _ = v7
    case "\n\t":
        p2 -= 32767
    default:
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: (^(fi8)), F1: 1}
            switch (fi8 | (fi8 >> ((0) & 63))) {
            case -9223372036854775808:
                v10 := []int{42, fi8, (fi8 % ((32767) | 1))}
                _ = v10
            default:
                for fi11 := 0; fi11 < 4; fi11++ {
                    v12 := GenS2{F0: 685, F1: (v5 << ((388) & 63))}
                    if (4294967295 <= 1) {
                        fmt.Println(fi11, fi11, p2)
                    } else {
                        p0 += p0
                    }
                    fmt.Println(p2, p0)
                    _ = fi11
                    _ = v12
                }
            }
            _ = fi8
            _ = v9
        }
    }
    var v13 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    return v13
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]uint32)
    var v3 float64 = (3.14159 / (3.14159 / 1.5))
    v4 := []string{"foo bar", "\n\t", ("a" + "\n\t"), "hello"}
    fmt.Println(v3)
    var v14 string = (("" + "hello") + "\n\t")
    var v15 uint32 = 42
    if (61 <= 18446744073709551615) {
        var v16 int64 = 1
        _ = v16
    } else {
        var v17 float32 = 1.5
        if (500 != 23) {
            v1 -= 42
            var v18 float32 = 0.0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v18
        } else {
            var v19 uint = 42
            fmt.Println(v15, v3)
            v14 = (((v14 + v14) + v14) + "hello")
            _ = v19
        }
        _ = v17
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 int8 = 6
        if true {
            var v22 bool = false
            v23 := GenS2{F0: -976, F1: 0}
            _ = v22
            _ = v23
        }
        fmt.Println(v3, v3)
        _ = fi20
        _ = v21
    }
    var v24 uint8 = (v1 * uint8(v15))
    v24 += (v1 % (((v1 * 2)) | 1))
    switch v14 {
    case "hello":
        switch -1 {
        case 2147483647:
            fmt.Println(v14, v3, v3)
        case 4294967295:
            var v25 bool = ((0 != 2) && false)
            _ = v25
        case -1:
            fmt.Println(v24, v15)
        }
    case "foo bar":
        v26 := GenS3{F0: ((v14 + "a") + (v14 + ("hello" + (("" + v14) + ("\n\t" + "")))))}
        _ = v26
    case "a":
        for fi27 := 0; fi27 < 5; fi27++ {
            if (false || false) {
                v28 := GenS2{F0: 1, F1: 255}
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 string = v14
                    _ = fi29
                    _ = v30
                }
                v15 += 65535
                _ = v28
            }
            var v31 uint32 = (^(v15))
            _ = fi27
            _ = v31
        }
    }
    _ = len(v2)
    var v32 uint16 = 65535
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    delete(v2, 42)
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 uint64 = 2
        var v35 int64 = -9223372036854775808
        switch v14 {
        case "":
            v32 += 2
        case "\n\t":
            v36 := GenS1{F0: 1, F1: -1}
            _ = v36
        default:
            var v37 int = fi33
            _ = v37
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    v14 = ""
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v4
}
