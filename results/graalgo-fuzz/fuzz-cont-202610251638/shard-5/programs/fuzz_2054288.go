package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 string) uint32 {
    p1 = ""
    v6 := GenS1{F0: -831, F1: 848, F2: 65535}
    fmt.Println(p0)
    var v7 string = "a"
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}
func genFn1() string {
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := []int{}
        var v10 uint8 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 float32 = 3.14159
    _ = v11
    return (("" + "a") + "\n\t")
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -1
    var v3 float64 = 0.0
    var v4 int = -234
    var v5 string = ("a" + "")
    v12 := GenS2{F0: 0, F1: 109, F2: 255}
    var v13 int = 743
    v2 = v2
    v5 = "hello"
    if true {
        fmt.Println(v4)
        if false {
            v3 += 3.14159
            fmt.Println(v5, v4, v4)
            v14 := GenS2{F0: 2, F1: 2147483647, F2: 0}
            _ = v14
        }
        fmt.Println(v4, v4, v5)
    }
    v15 := GenS3{F0: 978}
    switch v4 {
    case 2147483647:
        v3 += float64((float32(float64(3.14159)) + float32(v3)))
    case 940:
        v1[v5] = v5
    }
    v1[(v5 + v5)] = "foo bar"
    var v16 bool = (234 <= 127)
    for fi17 := 0; fi17 < 2; fi17++ {
        delete(v1, (((v5 + "foo bar") + v5) + v5))
        if (((0.1 - float64((((3.14159 * 1.5) - float32((v3 + v3))) * 3.14159))) - (v3 + (v3 * v3))) == float64(1.5)) {
            var v18 int32 = 1
            _ = v18
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            switch v13 {
            case 0:
                fmt.Println(fi19)
            case -7:
                v3 = ((v3 + 0.1) - (v3 + v3))
            }
            var v20 uint64 = 18446744073709551615
            v20 += v20
            _ = fi19
            _ = v20
        }
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
