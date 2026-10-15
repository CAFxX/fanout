package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint16, p1 uint64, p2 int32) uint8 {
    v4 := GenS1{F0: 65535, F1: (p1 <= (^(p1))), F2: 100}
    v5 := []int{42, -883, 42, -2147483648}
    p0 += uint16(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 127
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    v3 := []uint32{648}
    if true {
        v6 := GenS2{F0: 128}
        var v7 string = ("hello" + (((("a" + "hello") + ((((("\n\t" + ("" + "a")) + "\n\t") + "foo bar") + "a") + "")) + "\n\t") + "hello"))
        _ = v6
        _ = v7
    }
    v8 := GenS3{F0: 42, F1: -656}
    v9 := GenS3{F0: 591, F1: 42}
    var v10 uint = 42
    if true {
        if false {
            if true {
                fmt.Println(v10)
                if (("foo bar" + "\n\t") < "a") {
                    v11 := GenS4{F0: 847, F1: 841}
                    _ = v11
                } else {
                    var v12 string = "a"
                    var v13 bool = (!false)
                    switch 42 {
                    case 32768:
                        var v14 int16 = 1
                        _ = v14
                    }
                    _ = v12
                    _ = v13
                }
                var v15 float32 = -2.25
                _ = v15
            }
            var v16 string = "foo bar"
            fmt.Println(v16, v16)
            _ = v16
        }
    }
    if true {
        var v17 uint8 = 110
        delete(v1, "hello")
        _ = v17
    }
    v10 = v10
    v10 = 153
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
