package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0(p0 string, p1 bool) uint16 {
    p1 = p1
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 13; fi6++ {
        p1 = false
        v7 := []int{}
        p1 = p1
        _ = fi6
        _ = v7
    }
    var v8 uint16 = 0
    v8 += v8
    _ = p0
    _ = p1
    _ = v8
    return 523
}
func genFn1() uint32 {
    v9 := []int{9223372036854775807}
    v10 := GenS1{F0: 18446744073709551615, F1: -1}
    _ = v10
    _ = v9
    return 2147483648
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[uint32]string)
    v3 := []string{"\n\t", "hello", "hello"}
    v4 := make(map[uint32]uint32)
    var v5 float64 = (((0.1 + (1.5 + float64(0.1))) - 0.0) / float64((3.14159 - 0.1)))
    for fi11 := 0; fi11 < 16; fi11++ {
        if len(v3) > 0 {
            v3[0] = ""
        }
        var v12 int = (fi11 / ((2147483647) | 1))
        var v13 uint = 42
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v1, v5)
    var v14 bool = (1 > 775)
    v15 := GenS2{F0: 4294967295, F1: 640}
    var v16 int16 = 658
    if len(v3) > 0 {
        v3[0] = ("foo bar" + (("\n\t" + (("a" + "a") + "foo bar")) + "foo bar"))
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        if false {
            fmt.Println(v16, v1, v1)
            var v18 int64 = 42
            fmt.Println(v5, fi17, v5)
            _ = v18
        }
        fmt.Println(fi17)
        for fi19 := 0; fi19 < 6; fi19++ {
            _ = len(v3)
            delete(v4, 0)
            _ = fi19
        }
        _ = fi17
    }
    _ = len(v3)
    v3 = append(v3, "a")
    for fi20 := 0; fi20 < 17; fi20++ {
        if len(v3) > 0 {
            v3[0] = ("hello" + "")
        }
        _ = fi20
    }
    fmt.Println(v1, v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
