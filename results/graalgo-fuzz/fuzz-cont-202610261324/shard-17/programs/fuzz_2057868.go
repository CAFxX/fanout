package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 bool }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 string; F2 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 0
    v3 := []int{}
    v4 := []string{}
    v5 := GenS1{F0: 309, F1: (((("" + "") + ("foo bar" + "")) + "\n\t") + (("hello" + "") + "\n\t")), F2: true}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 uint32 = 2
    for fi7 := 0; fi7 < 19; fi7++ {
        switch (fi7 - (fi7 & fi7)) {
        case 32768:
            _ = len(v3)
        case 9223372036854775807:
            var v8 uint64 = 18446744073709551615
            _ = v8
        }
        _ = fi7
    }
    var v9 uint16 = 32767
    switch -828 {
    case 4294967295:
        for fi10 := 0; fi10 < 9; fi10++ {
            if len(v3) > 0 {
                v3[0] = fi10
            }
            v11 := GenS2{F0: -2147483648, F1: 42}
            v2 -= uint(v6)
            _ = fi10
            _ = v11
        }
    case 42:
        for fi12 := 0; fi12 < 2; fi12++ {
            v9 -= v9
            _ = fi12
        }
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := v1[(^(v6))]
        v15 := GenS3{F0: uint8(v9)}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 string = ""
    var v17 int64 = 0
    var v18 float32 = 3.14159
    for fi19 := 0; fi19 < 19; fi19++ {
        if true {
            fmt.Println(v17)
        }
        _ = fi19
    }
    var v20 int16 = 92
    fmt.Println(v17, v16)
    switch v16 {
    case "foo bar":
        v21 := GenS4{F0: v2, F1: v16, F2: -740}
        _ = v21
    default:
        for fi22 := 0; fi22 < 20; fi22++ {
            if false {
                delete(v1, uint32(v9))
                if true {
                    v23 := GenS1{F0: 4294967295, F1: v16, F2: (!false)}
                    v2 -= 0
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                    _ = v23
                } else {
                    if (false || true) {
                        v24 := GenS3{F0: 206}
                        var v25 int16 = 708
                        _ = v24
                        _ = v25
                    } else {
                        fmt.Println(v6, v6)
                    }
                }
            }
            _ = fi22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
