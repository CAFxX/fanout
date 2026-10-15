package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8; F1 bool; F2 string }
func genFn0(p0 uint, p1 int32) uint32 {
    p0 += 257
    p0 -= (p0 % (((p0 - (p0 + p0))) | 1))
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1() string {
    if true {
        v6 := GenS1{F0: 255, F1: "a"}
        _ = v6
    } else {
        var v7 int32 = -2147483648
        _ = v7
    }
    var v8 uint = 255
    _ = v8
    return (("hello" + "\n\t") + "a")
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ("" + "hello")
    var v3 bool = false
    v4 := make(map[uint32]uint32)
    var v9 int32 = -1
    var v10 uint16 = 782
    var v11 uint64 = 18446744073709551615
    fmt.Println(v11)
    delete(v1, int(v11))
    if false {
        delete(v1, -396)
        v12 := GenS1{F0: 70, F1: ((((v2 + "foo bar") + "a") + v2) + v2)}
        var v13 uint = 2
        _ = v12
        _ = v13
    }
    v2 = v2
    fmt.Println(v10, v3, v11)
    var v14 float32 = (1.5 * ((1e10 / 3.14159) / float32(((-2.25 + 0.1) + 1e10))))
    if v3 {
        var v15 uint = 1
        v10 = v10
        v11 = (v11 / ((v11) | 1))
        _ = v15
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int16 = -32768
            v14 += v14
            fmt.Println(fi16)
            _ = fi16
            _ = v17
        }
        v18 := []int{796}
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v3 = (v14 >= v14)
        if false {
            if v3 {
                v3 = (798 >= v10)
                fmt.Println(v11, v9, v9)
            } else {
                fmt.Println(v14)
            }
            v20 := []int{}
            v11 -= v11
            _ = v20
        } else {
            for fi21 := 0; fi21 < 2; fi21++ {
                v22 := GenS2{F0: 0, F1: (-9223372036854775808 < 42), F2: v2}
                switch ("" + (v2 + v2)) {
                case "a":
                    fmt.Println(v3, v10, v11)
                case "foo bar":
                    var v23 string = v2
                    var v24 uint64 = (v11 - (v11 + v11))
                    fmt.Println(v23, fi21, v10)
                    _ = v23
                    _ = v24
                default:
                    var v25 int8 = int8(v10)
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
        }
        _ = fi19
    }
    v11 -= 4294967295
    fmt.Println(v10, v2)
    var v26 uint32 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
