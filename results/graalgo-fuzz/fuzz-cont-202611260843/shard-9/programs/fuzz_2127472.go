package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 string }
func genFn0(p0 int64) uint8 {
    p0 -= p0
    p0 = p0
    v4 := []int{}
    _ = p0
    _ = v4
    return 156
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = (("foo bar" + "foo bar") + "hello")
    var v3 int64 = -72
    switch 32 {
    case 536:
        v2 = (v2 + (v2 + v2))
    case -2147483648:
        if (!false) {
            var v5 uint8 = 96
            switch -283 {
            case -684:
                if (42 >= -128) {
                    fmt.Println(v3, v3)
                    var v6 string = v2
                    v7 := GenS1{F0: -110, F1: 42, F2: 1}
                    _ = v6
                    _ = v7
                } else {
                    if true {
                        v3 = v3
                    } else {
                        v8 := v1[1]
                        fmt.Println(v3, v2)
                        _ = v8
                    }
                    v9 := GenS1{F0: 0, F1: 322, F2: 0}
                    v5 += 130
                    _ = v9
                }
            }
            _ = v5
        } else {
            v2 = v2
            switch 32767 {
            case 40:
                var v10 int8 = 106
                _ = v10
            case 42:
                for fi11 := 0; fi11 < 12; fi11++ {
                    fmt.Println(fi11, v2, fi11)
                    _ = fi11
                }
            case 4294967295:
                v2 = ((v2 + ("foo bar" + "\n\t")) + (v2 + v2))
            }
            if true {
                v2 = "hello"
            }
        }
    }
    var v12 int32 = -160
    if true {
        _ = len(v1)
        for fi13 := 0; fi13 < 17; fi13++ {
            v2 = "\n\t"
            var v14 int = fi13
            _ = fi13
            _ = v14
        }
        var v15 bool = (29 != 42)
        _ = v15
    } else {
        v3 += (-(v3))
    }
    fmt.Println(v12, v2)
    v16 := GenS2{F0: 2, F1: 692, F2: 672}
    if ((true || (3.14159 >= 1e10)) || ((!(!(false && true))) || false)) {
        v17 := GenS3{F0: v12}
        switch ("" + "a") {
        case "a":
            if true {
                var v18 int8 = -1
                _ = v18
            } else {
                fmt.Println(v3, v3)
            }
        }
        _ = v17
    } else {
        v19 := GenS4{F0: -969, F1: v2}
        var v20 bool = ((true || (!(125 <= 1))) || (v2 < v2))
        _ = v19
        _ = v20
    }
    if false {
        fmt.Println(v12, v2, v12)
        v3 = v3
        v3 -= v3
    }
    var v21 string = (v2 + "")
    var v22 uint8 = 127
    var v23 int64 = (-(v3))
    var v24 float32 = 1.5
    v25 := GenS2{F0: 767, F1: 589, F2: 0}
    fmt.Println(v12)
    fmt.Println(v21)
    v26 := []int{-2147483648, 42, 2147483647, 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
