package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int64; F2 int32 }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint8, p1 uint16) int16 {
    p0 += (p0 | 128)
    var v7 uint8 = 237
    v8 := []int{2147483647, 32767}
    v8 = append(v8, 127)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1() uint8 {
    var v9 float64 = -2.25
    if true {
        if false {
            v10 := GenS1{F0: (1 != 1)}
            _ = v10
        }
        var v11 float32 = ((1.5 / (-2.25 * (0.1 - 3.14159))) * (0.0 * (1.5 / (float32(1.5) / (float32(v9) - 1e10)))))
        _ = v11
    } else {
        switch -591 {
        case 128:
            v12 := GenS1{F0: (float64(1e10) > v9)}
            _ = v12
        case 1:
            var v13 uint32 = 0
            _ = v13
        }
        v9 -= (3.14159 * 0.1)
        v14 := GenS2{F0: -203, F1: -1, F2: -1}
        _ = v14
    }
    _ = v9
    return 1
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]string)
    v3 := []string{("a" + ("" + "foo bar"))}
    var v4 int32 = 0
    v5 := make(map[uint32]uint32)
    v6 := []string{"a", "", "hello", ("\n\t" + "\n\t")}
    v1 -= (v1 / 0.1)
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(v4, fi15, v4)
        var v16 uint32 = 1
        var v17 int32 = v4
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int32 = (v4 - -342)
    v4 -= v4
    v1 += (v1 / v1)
    var v19 uint8 = 2
    v20 := v2["\n\t"]
    for fi21 := 0; fi21 < 15; fi21++ {
        switch v20 {
        case "foo bar":
            var v22 float32 = 3.14159
            _ = v22
        case "":
            v23 := GenS2{F0: int(v19), F1: 9223372036854775807, F2: v18}
            _ = v23
        case "a":
            v18 += 1
        }
        _ = fi21
    }
    if (true && false) {
        v24 := GenS3{F0: 981, F1: 65535}
        _ = v24
    } else {
        if false {
            for fi25 := 0; fi25 < 15; fi25++ {
                fmt.Println(v19, v18)
                v1 += v1
                _ = fi25
            }
            var v26 uint64 = 18446744073709551615
            _ = v26
        }
        for fi27 := 0; fi27 < 15; fi27++ {
            if false {
                v20 = v20
                for fi28 := 0; fi28 < 13; fi28++ {
                    var v29 uint = 586
                    _ = fi28
                    _ = v29
                }
            } else {
                v1 -= float64(-2.25)
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 uint8 = (v19 >> ((65535) & 63))
                    v6 = append(v6, "\n\t")
                    _ = fi30
                    _ = v31
                }
                v18 += 42
            }
            _ = fi27
        }
    }
    v18 += 42
    if (1 != 1) {
        var v32 int32 = v4
        _ = v32
    }
    if false {
        v4 += -2147483648
    } else {
        if true {
            v33 := v2[(v20 + "hello")]
            v20 = (v20 + ("\n\t" + v33))
            _ = v33
        }
    }
    if false {
        var v34 uint32 = 2147483647
        switch 2147483647 {
        case 32768:
            fmt.Println(v1, v4, v18)
        case 127:
            if false {
                v6 = append(v6, "foo bar")
                v35 := GenS2{F0: 4294967295, F1: -9223372036854775808, F2: 0}
                fmt.Println(v19, v18, v34)
                _ = v35
            }
        }
        switch v20 {
        case "a":
            switch "\n\t" {
            case "foo bar":
                if true {
                    fmt.Println(v34)
                    v36 := GenS1{F0: true}
                    v19 += (v19 | v19)
                    _ = v36
                } else {
                    fmt.Println(v20, v19, v34)
                }
            }
        default:
            var v37 string = ("foo bar" + v20)
            _ = v37
        }
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
