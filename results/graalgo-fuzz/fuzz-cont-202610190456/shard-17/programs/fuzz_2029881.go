package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 string }
func genFn0() string {
    var v7 uint = 65535
    var v8 uint8 = 86
    v9 := GenS1{F0: 42, F1: 42}
    var v10 float32 = 0.0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ("\n\t" + ("foo bar" + ""))
}
func genFn1(p0 int8, p1 uint64) int32 {
    v11 := GenS2{F0: 733, F1: 42, F2: (0.1 >= (1.5 + 0.1))}
    var v12 int = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 512
}

func main() {
    var v1 uint16 = 65535
    var v2 int8 = 0
    v3 := []string{"\n\t", ("\n\t" + "\n\t"), "foo bar", ""}
    v4 := make(map[int]string)
    var v5 uint16 = v1
    var v6 int8 = (v2 << ((65535) & 63))
    var v13 int8 = (v6 + -1)
    v5 -= v5
    var v14 uint32 = 4294967295
    v15 := GenS2{F0: v5, F1: 255, F2: false}
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 string = ("hello" + "")
        _ = fi16
        _ = v17
    }
    v6 = (v13 - v13)
    v18 := GenS3{F0: 1}
    switch 32768 {
    case 322:
        for fi19 := 0; fi19 < 3; fi19++ {
            v3 = append(v3, "\n\t")
            var v20 uint64 = 18446744073709551615
            fmt.Println(v14, v2)
            _ = fi19
            _ = v20
        }
    case -377:
        v1 = (v1 - (^(v5)))
    }
    fmt.Println(v13, v2)
    var v21 string = ""
    if false {
        v1 = v5
        v14 = (v14 ^ 65535)
        v1 = 65535
    }
    var v22 uint32 = (^(v14))
    for fi23 := 0; fi23 < 13; fi23++ {
        v6 = 124
        v24 := GenS4{F0: 42, F1: (v21 + (v21 + ((v21 + v21) + "foo bar")))}
        _ = fi23
        _ = v24
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v22)
    v25 := v4[9223372036854775807]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
