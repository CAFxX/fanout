package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []string{(("" + ("foo bar" + ("hello" + ((("foo bar" + ("a" + (("\n\t" + ("hello" + ("hello" + "foo bar"))) + "hello"))) + "foo bar") + "foo bar")))) + "foo bar"), ("foo bar" + ("\n\t" + ""))}
    var v2 bool = true
    v3 := make(map[int]string)
    v4 := []uint32{4294967295, 495}
    v5 := make(map[uint32]int)
    var v6 float32 = (3.14159 / 1.0)
    var v7 uint = 42
    fmt.Println(v7, v2, v6)
    fmt.Println(v7, v6, v6)
    if v2 {
        var v8 int8 = -42
        switch "" {
        case "\n\t":
            fmt.Println(v6, v2, v2)
        case "a":
            for fi9 := 0; fi9 < 10; fi9++ {
                switch fi9 {
                case 821:
                    v7 = (v7 >> (((v7 << (((^(v7))) & 63))) & 63))
                case 127:
                    var v10 uint16 = 32767
                    _ = v10
                case -9223372036854775808:
                    fmt.Println(v7, v8, v8)
                }
                _ = fi9
            }
        case "":
            v8 += v8
        default:
            v11 := GenS1{F0: 537}
            _ = v11
        }
        if v2 {
            for fi12 := 0; fi12 < 7; fi12++ {
                fmt.Println(v6)
                fmt.Println(v6, v6)
                var v13 float32 = 3.14159
                _ = fi12
                _ = v13
            }
            v2 = (false || (v2 || true))
        } else {
            v14 := GenS1{F0: 437}
            if (v2 || (true && (v2 || ((float64(v6) * float64(1e10)) == ((3.14159 * -2.25) - -2.25))))) {
                var v15 int16 = 299
                _ = v15
            } else {
                var v16 int = -922
                _ = v16
            }
            fmt.Println(v2)
            _ = v14
        }
        _ = v8
    } else {
        fmt.Println(v2, v7, v2)
        fmt.Println(v7)
        v2 = (32767 <= 0)
    }
    var v17 int32 = 42
    v2 = (1 > -818)
    if false {
        v18 := GenS1{F0: (v17 % ((v17) | 1))}
        var v19 int16 = 0
        delete(v5, 640)
        _ = v18
        _ = v19
    }
    var v20 float64 = (float64(v6) * (0.1 * float64((0.0 / v6))))
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := v3[4294967295]
        v23 := GenS1{F0: (v17 - (v17 * -1))}
        v24 := GenS2{F0: 65535}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    if v2 {
        v2 = (v2 && v2)
        v25 := v5[2147483648]
        if len(v1) > 0 {
            v1[0] = ("\n\t" + "\n\t")
        }
        _ = v25
    }
    var v26 uint16 = 65535
    v27 := GenS2{F0: 4294967295}
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 string = "\n\t"
        v17 += -307
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        v4 = append(v4, 42)
        v6 = v6
        _ = fi30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
