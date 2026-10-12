package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 uint64, p1 uint8) uint {
    if false {
        v5 := []int{32767, 54, -1}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if (false && ((!(!(!true))) && (3.14159 != 0.1))) {
            var v6 int16 = 1000
            var v7 uint32 = 1
            var v8 int = 936
            _ = v6
            _ = v7
            _ = v8
        } else {
            var v9 uint32 = 2
            _ = v9
        }
        _ = v5
    } else {
        switch "hello" {
        case "a":
            v10 := []int{0, 127, 490}
            _ = v10
        case "foo bar":
            v11 := []int{}
            _ = v11
        }
    }
    p1 += p1
    v12 := []int{32768}
    v13 := GenS1{F0: 648}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 2
}
func genFn1(p0 uint8) uint16 {
    fmt.Println(p0)
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        p0 -= (p0 * p0)
    case "\n\t":
        v14 := []int{-2147483648}
        _ = v14
    case "":
        var v15 float32 = -2.25
        _ = v15
    default:
        if false {
            p0 -= 2
            p0 = (p0 % (((p0 | p0)) | 1))
            if false {
                var v16 int64 = 764
                _ = v16
            } else {
                for fi17 := 0; fi17 < 5; fi17++ {
                    p0 = p0
                    _ = fi17
                }
            }
        } else {
            v18 := []int{-2147483648}
            _ = v18
        }
    }
    if false {
        fmt.Println(p0)
    }
    if ((!false) && (2147483647 != -1)) {
        switch "a" {
        case "\n\t":
            if false {
                fmt.Println(p0)
                var v19 float64 = (1.5 - 1.5)
                _ = v19
            }
        }
    }
    _ = p0
    return 65535
}

func main() {
    var v1 uint16 = 744
    v2 := []string{}
    var v3 uint64 = 938
    var v4 uint = uint(v3)
    v3 = (v3 / ((uint64(v4)) | 1))
    v4 -= v4
    var v20 float32 = 3.14159
    if (true || (!(2 == 0))) {
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS1{F0: (v3 / ((v3) | 1))}
            v23 := GenS2{F0: 4294967295, F1: true}
            var v24 string = (("a" + "a") + "foo bar")
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 5; fi25++ {
            v4 = v4
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi25
        }
    } else {
        fmt.Println(v3, v20)
    }
    fmt.Println(v4, v3, v1)
    v26 := GenS3{F0: genFn0(4294967295, 1), F1: 32767}
    if (42 >= 725) {
        var v27 uint32 = 4294967295
        fmt.Println(v1, v3, v4)
        var v28 uint16 = (v1 | (v1 % ((130) | 1)))
        _ = v27
        _ = v28
    }
    v29 := GenS2{F0: 865, F1: (true || false)}
    v1 = (v1 & (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
