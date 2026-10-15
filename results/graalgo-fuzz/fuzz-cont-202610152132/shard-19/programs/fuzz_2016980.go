package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0() uint32 {
    var v6 uint64 = 18446744073709551615
    if ((65535 != 2) || (!false)) {
        fmt.Println(v6)
    }
    _ = v6
    return 632
}

func main() {
    var v1 bool = (false || true)
    v2 := []string{"foo bar", ("\n\t" + "foo bar"), "a", ("foo bar" + "")}
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    var v7 bool = false
    var v8 float64 = (0.1 * -2.25)
    v8 -= v8
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    switch -1 {
    case 294:
        var v9 int32 = 548
        _ = v9
    case -597:
        var v10 int16 = -1
        _ = v10
    default:
        var v11 string = "hello"
        _ = v11
    }
    v1 = ((-2.25 / float32((3.14159 * float64(-2.25)))) < 1e10)
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 uint8 = 2
        _ = fi12
        _ = v13
    }
    var v14 uint32 = 811
    v15 := GenS1{F0: 1}
    v16 := GenS1{F0: 65535}
    _ = len(v4)
    for fi17 := 0; fi17 < 19; fi17++ {
        if v7 {
            var v18 uint16 = 761
            _ = v18
        }
        v8 -= float64((-2.25 / float32(v8)))
        _ = fi17
    }
    var v19 uint32 = (v14 / (((v14 | 2147483647)) | 1))
    var v20 string = ""
    v2 = append(v2, "a")
    v21 := GenS1{F0: 134}
    v22 := GenS2{F0: (v19 ^ (v19 >> ((2) & 63)))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
