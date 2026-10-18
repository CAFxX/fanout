package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint64) uint16 {
    switch 4294967295 {
    case 0:
        for fi5 := 0; fi5 < 19; fi5++ {
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 int32 = 1
                v8 := GenS1{F0: 42}
                var v9 int32 = 820
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            v10 := []int{(fi5 - (-(fi5))), 899}
            p0 -= 775
            _ = fi5
            _ = v10
        }
    case -9223372036854775808:
        p1 = (p1 + (p1 & 0))
    }
    var v11 uint64 = (^(p1))
    if false {
        var v12 uint8 = 150
        p0 = 1
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v11
    return 293
}

func main() {
    v1 := []string{("\n\t" + ("hello" + ("hello" + ((("hello" + "hello") + "") + ""))))}
    var v2 uint8 = 2
    v3 := []string{"\n\t", (("a" + "\n\t") + "hello"), ("\n\t" + ""), ("foo bar" + "a")}
    v4 := make(map[int]string)
    v4[910] = "hello"
    fmt.Println(v2)
    v13 := GenS1{F0: 297}
    switch ("a" + ("foo bar" + ("\n\t" + "foo bar"))) {
    case "hello":
        switch 127 {
        case 338:
            if true {
                delete(v4, 128)
                v14 := GenS1{F0: 323}
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint32 = 635
                    var v17 bool = (-1 < 0)
                    v17 = v17
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = v14
            } else {
                if true {
                    var v18 bool = true
                    var v19 float64 = (0.1 + (0.1 - ((3.14159 * -2.25) * 1e10)))
                    _ = v18
                    _ = v19
                }
                var v20 int32 = 2147483647
                var v21 float32 = float32(1e10)
                _ = v20
                _ = v21
            }
        case 9223372036854775807:
            var v22 uint64 = 42
            _ = v22
        case 32768:
            var v23 uint8 = v2
            _ = v23
        default:
            var v24 uint16 = 255
            _ = v24
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 uint32 = 1
        v26 -= 65535
        v2 -= 42
        _ = fi25
        _ = v26
    }
    v2 = v2
    if true {
        v2 += v2
        var v27 int8 = -1
        var v28 string = ("foo bar" + "a")
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: 620}
    var v30 uint64 = 4294967295
    switch 182 {
    case 42:
        v3 = append(v3, ("hello" + "hello"))
    case 32768:
        var v31 int8 = 42
        _ = v31
    case 9223372036854775807:
        fmt.Println(v2, v30)
    }
    var v32 uint32 = 1
    v33 := GenS2{F0: 370, F1: 68}
    if (v32 == v32) {
        var v34 int64 = -1
        _ = v34
    } else {
        v35 := GenS3{F0: 32767}
        _ = v35
    }
    var v36 int64 = 593
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v36
    _ = v4
}
