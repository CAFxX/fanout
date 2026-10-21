package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 string }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 string = ""
    v2 := []int{404, 837}
    var v3 int32 = -1
    var v4 string = "hello"
    var v5 float32 = 1.5
    v6 := make(map[uint32]string)
    var v7 uint8 = 2
    v5 = v5
    v1 = v1
    v8 := GenS1{F0: 2, F1: 1, F2: (v4 + "")}
    for fi9 := 0; fi9 < 17; fi9++ {
        v5 += (0.1 + v5)
        fmt.Println(v3, v3)
        _ = fi9
    }
    if ((2147483648 >= 287) && ((true || (65535 > 1)) && true)) {
        if (!(false || true)) {
            v10 := GenS1{F0: 2147483648, F1: 255, F2: "hello"}
            _ = v10
        }
    }
    if (0.0 < 0.0) {
        if true {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 int16 = 1
                v3 += v3
                v13 := GenS2{F0: 1}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v14 := GenS1{F0: 65535, F1: 65535, F2: v1}
        var v15 uint8 = v7
        _ = v14
        _ = v15
    }
    var v16 int8 = 127
    for fi17 := 0; fi17 < 16; fi17++ {
        fmt.Println(v1, v16)
        fmt.Println(v4, v3)
        for fi18 := 0; fi18 < 7; fi18++ {
            v2 = append(v2, (fi17 >> ((268) & 63)))
            var v19 string = ((("" + "a") + "foo bar") + v1)
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 float32 = (-2.25 + 3.14159)
    var v21 uint32 = 42
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 bool = (uint(v21) == uint(v21))
        _ = fi22
        _ = v23
    }
    var v24 int8 = v16
    for fi25 := 0; fi25 < 2; fi25++ {
        v4 = ((v1 + v1) + ("hello" + (v4 + (v4 + v1))))
        _ = fi25
    }
    v16 += (v24 | (v24 * -128))
    v6[(v21 ^ v21)] = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
