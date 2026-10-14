package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int32 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint8, p2 bool) uint32 {
    switch "" {
    case "foo bar":
        var v7 uint32 = 0
        _ = v7
    case "hello":
        v8 := GenS1{F0: 557, F1: 894, F2: 0}
        _ = v8
    }
    p2 = (p0 != (p0 % (((p0 % (((p0 - p0)) | 1))) | 1)))
    var v9 uint = 2
    v9 = (v9 ^ 608)
    switch (("\n\t" + ((("a" + "foo bar") + (("a" + "\n\t") + "hello")) + "foo bar")) + ("" + "\n\t")) {
    case "foo bar":
        p2 = ((-2.25 != 1.5) && true)
    default:
        if (4294967295 != 1) {
            var v10 uint32 = 42
            _ = v10
        } else {
            var v11 float32 = -2.25
            var v12 int8 = -21
            v13 := GenS1{F0: v9, F1: 937, F2: p0}
            _ = v11
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2
}
func genFn1(p0 uint64) uint16 {
    p0 -= (p0 << ((65535) & 63))
    if (!(1e10 < 0.1)) {
        var v14 float64 = (1e10 - float64(0.0))
        switch 42 {
        case -1:
            var v15 int32 = 1
            _ = v15
        }
        _ = v14
    } else {
        p0 = (p0 | p0)
        p0 += 42
        fmt.Println(p0)
    }
    v16 := GenS2{F0: ("\n\t" + ""), F1: 18446744073709551615}
    var v17 float32 = (((float32(0.1) / (0.0 + 3.14159)) + float32(-2.25)) * (0.0 + 0.0))
    _ = p0
    _ = v16
    _ = v17
    return 65535
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]bool)
    var v3 float32 = 3.14159
    var v4 float64 = 0.1
    var v5 int = 157
    var v6 uint8 = 127
    var v18 uint32 = 4294967295
    var v19 float64 = -2.25
    v18 = v18
    v20 := []int{(v5 % ((v5) | 1)), v5}
    var v21 int64 = -146
    var v22 int32 = -2147483648
    var v23 int = -1
    if true {
        v4 = float64(v3)
    } else {
        for fi24 := 0; fi24 < 10; fi24++ {
            for fi25 := 0; fi25 < 18; fi25++ {
                if false {
                    var v26 string = "hello"
                    _ = v26
                }
                _ = len(v2)
                if (!(true && false)) {
                    var v27 uint = 65535
                    _ = v27
                }
                v19 -= -2.25
                fmt.Println(v21, fi25, v19)
                v28 := GenS3{F0: (v18 % ((592) | 1))}
                _ = len(v2)
                _ = fi25
                _ = v28
            }
            _ = fi24
        }
    }
    v29 := GenS3{F0: v18}
    v30 := GenS3{F0: 744}
    var v31 float32 = v3
    v32 := GenS2{F0: "", F1: 4294967295}
    switch (v5 | v5) {
    case 32767:
        v20 = append(v20, int(v6))
    case 0:
        v33 := GenS4{F0: 601}
        _ = v33
    }
    var v34 bool = false
    v19 -= 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
