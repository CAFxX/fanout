package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint }
type GenS6 struct { F0 int; F1 int8; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 string = "a"
    if false {
        var v4 string = v3
        v5 := GenS1{F0: false, F1: 2147483647}
        _ = v4
        _ = v5
    } else {
        fmt.Println(v3)
        v3 = ("a" + (v3 + (v3 + (v3 + "foo bar"))))
    }
    v6 := GenS1{F0: true, F1: 647}
    var v7 string = ((v3 + v3) + v3)
    v8 := []int{2147483647}
    _ = len(v1)
    if (v7 == "") {
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 int = fi9
            for fi11 := 0; fi11 < 20; fi11++ {
                v8 = append(v8, (v10 % ((fi11) | 1)))
                _ = fi11
            }
            v12 := GenS1{F0: true, F1: -2147483648}
            _ = fi9
            _ = v10
            _ = v12
        }
        var v13 int8 = 0
        v14 := v1[-191]
        _ = v13
        _ = v14
    }
    if false {
        if len(v8) > 0 {
            v8[0] = 4294967295
        }
    } else {
        v15 := GenS2{F0: -1, F1: 255}
        _ = v15
    }
    v3 = ""
    var v16 float64 = 3.14159
    v17 := GenS1{F0: true, F1: -936}
    if true {
        v18 := v2[-878]
        _ = v18
    }
    var v19 uint8 = 170
    var v20 int64 = -839
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 uint64 = 1
            var v24 int = (fi21 << ((211) & 63))
            delete(v2, 0)
            _ = fi22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            v20 += 0
            if (2147483647 <= -2147483648) {
                var v26 bool = true
                v27 := GenS3{F0: (!true), F1: ("foo bar" + v7)}
                _ = v26
                _ = v27
            } else {
                var v28 uint = 450
                v16 -= 1e10
                switch v7 {
                case "hello":
                    _ = len(v8)
                case "a":
                    v16 = (v16 * v16)
                    var v29 uint32 = 232
                    v30 := GenS4{F0: -393, F1: fi25}
                    _ = v29
                    _ = v30
                case "\n\t":
                    _ = len(v8)
                }
                v16 += (float64(1.5) * v16)
                if false {
                    v31 := GenS5{F0: (^(v19)), F1: 1, F2: (v28 / ((v28) | 1))}
                    _ = v31
                } else {
                    v32 := GenS4{F0: 0, F1: fi25}
                    v33 := GenS6{F0: fi25, F1: 100, F2: 2}
                    _ = v32
                    _ = v33
                }
                _ = v28
            }
            if false {
                v3 = (v3 + ((("a" + v3) + "a") + v7))
                var v34 uint = 839
                _ = v34
            } else {
                var v35 int64 = 1
                _ = v35
            }
            _ = fi25
        }
        _ = len(v8)
        _ = fi21
    }
    fmt.Println(v3, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
