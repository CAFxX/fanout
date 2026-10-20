package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 string, p1 bool, p2 bool) bool {
    p0 = (("\n\t" + p0) + (p0 + p0))
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: -142, F1: 1}
        _ = fi6
        _ = v7
    }
    fmt.Println(p2, p0, p2)
    v8 := []int{32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p1
}
func genFn1() int64 {
    switch ("a" + ("foo bar" + (("hello" + ("\n\t" + "a")) + "foo bar"))) {
    case "foo bar":
        var v9 float64 = -2.25
        _ = v9
    default:
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                if false {
                    fmt.Println(fi11, fi10)
                    v12 := GenS1{F0: 1, F1: 65535}
                    _ = v12
                }
                var v13 bool = true
                _ = fi11
                _ = v13
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := []int{(fi10 + fi10)}
                var v16 bool = true
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := GenS1{F0: 1, F1: 147}
            _ = fi10
            _ = v17
        }
    }
    v18 := []int{2147483647, -9223372036854775808}
    for fi19 := 0; fi19 < 14; fi19++ {
        fmt.Println(fi19)
        _ = fi19
    }
    _ = v18
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 32767
    v3 := []string{"hello", "hello"}
    v4 := make(map[uint32]bool)
    var v5 uint16 = 255
    v2 += (v2 * v2)
    v5 = (v5 - (v5 | 1))
    _ = len(v1)
    var v20 float32 = -2.25
    v2 -= 32768
    if len(v3) > 0 {
        v3[0] = ""
    }
    fmt.Println(v5)
    if false {
        var v21 string = "foo bar"
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        v2 += -9223372036854775808
        var v23 uint8 = 0
        fmt.Println(fi22, v2)
        _ = fi22
        _ = v23
    }
    if false {
        _ = len(v3)
    }
    if true {
        fmt.Println(v20, v2)
    } else {
        v24 := GenS2{F0: ("a" + ("foo bar" + "a")), F1: v5}
        _ = v24
    }
    _ = len(v1)
    var v25 int64 = genFn1()
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
