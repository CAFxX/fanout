package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint32) uint16 {
    v5 := GenS1{F0: 0}
    var v6 uint16 = 237
    var v7 int16 = -32768
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (v6 ^ (v6 - 42))
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[int]int)
    var v3 bool = (-1 == 0)
    var v4 uint16 = 0
    v8 := GenS2{F0: 42, F1: 0, F2: 240}
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 float32 = -2.25
            switch -9223372036854775808 {
            case 1:
                switch (fi9 >> ((535) & 63)) {
                case -1:
                    fmt.Println(fi9)
                case 128:
                    v12 := GenS3{F0: 128, F1: 1}
                    _ = v12
                case 32767:
                    fmt.Println(v4, fi10, fi9)
                default:
                    if ((v3 || v3) && (1 > 0)) {
                        v2[(-(fi10))] = fi10
                    }
                }
            case -1:
                fmt.Println(v11, v3, fi10)
            }
            switch fi9 {
            case 32768:
                v4 = 42
            }
            _ = fi10
            _ = v11
        }
        var v13 uint = 65535
        _ = fi9
        _ = v13
    }
    if (float32((3.14159 / 1e10)) > ((((1.5 * (((0.1 + 1.5) * 1.5) + 1e10)) - (-2.25 / (((3.14159 + -2.25) * ((0.0 - 0.0) / 3.14159)) - 0.1))) / ((0.0 + 1e10) * float32(1.5))) - 3.14159)) {
        if (!(-877 >= 35)) {
            fmt.Println(v3, v3)
        } else {
            var v14 int8 = 100
            switch (("foo bar" + "foo bar") + "a") {
            case "":
                v15 := GenS1{F0: -2147483648}
                _ = v15
            case "foo bar":
                fmt.Println(v3)
            }
            v4 -= 65535
            _ = v14
        }
    }
    if (!true) {
        v4 = v4
    } else {
        v4 -= v4
    }
    v1 = append(v1, ("a" + ""))
    var v16 bool = (4294967295 >= 4294967295)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v17 uint8 = 1
    var v18 uint16 = v4
    var v19 uint32 = 714
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
