package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int16, p1 int16, p2 int16) uint8 {
    p2 = (p1 % (((p2 / (((p1 + 32767)) | 1))) | 1))
    v7 := []int{}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 40
}

func main() {
    var v1 uint = 255
    v2 := []int{}
    v3 := make(map[uint32]string)
    var v4 int = -1000
    var v5 uint32 = 2147483648
    v6 := make(map[uint32]uint32)
    fmt.Println(v1, v1, v4)
    for fi9 := 0; fi9 < 5; fi9++ {
        if ("a" < "hello") {
            v4 += (v4 & v4)
            var v10 float64 = 3.14159
            _ = v10
        }
        delete(v3, (v5 % ((1) | 1)))
        _ = fi9
    }
    v1 += v1
    for fi11 := 0; fi11 < 5; fi11++ {
        if false {
            switch "foo bar" {
            case "hello":
                v4 += (-(v4))
            case "\n\t":
                v12 := GenS1{F0: v1}
                _ = v12
            case "":
                var v13 int = v4
                _ = v13
            default:
                var v14 int = int(v5)
                _ = v14
            }
            var v15 int8 = 42
            _ = v15
        }
        fmt.Println(fi11, fi11, v4)
        v5 = v5
        _ = fi11
    }
    fmt.Println(v1)
    v1 -= v1
    v4 += (v4 >> ((v1) & 63))
    for fi16 := 0; fi16 < 8; fi16++ {
        if true {
            _ = len(v2)
        } else {
            var v17 string = ""
            var v18 uint32 = 2
            _ = v17
            _ = v18
        }
        v4 += fi16
        v4 = v4
        _ = fi16
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(v5, v1, fi19)
        v20 := GenS1{F0: v1}
        var v21 uint = v1
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
