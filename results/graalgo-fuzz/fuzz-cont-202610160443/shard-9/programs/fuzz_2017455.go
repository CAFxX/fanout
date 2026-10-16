package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int; F1 int8; F2 string }
func genFn0() uint16 {
    var v7 uint64 = 2
    var v8 float32 = 1.5
    var v9 bool = false
    v10 := []int{9223372036854775807, 0, 42, 0}
    var v11 int64 = 1
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 uint32, p1 uint64) uint64 {
    fmt.Println(p1, p0)
    if false {
        if (2147483647 > -2147483648) {
            fmt.Println(p1, p0)
            v12 := []int{801}
            p0 += 619
            _ = v12
        }
    }
    v13 := GenS1{F0: (133 <= 4294967295)}
    _ = p0
    _ = p1
    _ = v13
    return p1
}

func main() {
    var v1 bool = (("foo bar" + "foo bar") > ("a" + "hello"))
    var v2 uint = 2
    v3 := []uint32{2147483647}
    var v4 int32 = -2147483648
    v5 := make(map[uint32]string)
    v6 := []string{(("a" + "foo bar") + "hello"), ("foo bar" + "\n\t"), "", "foo bar"}
    if (!v1) {
        if v1 {
            v2 = (v2 & 2)
            v1 = true
            v4 = v4
        } else {
            switch "\n\t" {
            case "hello":
                var v14 uint32 = 89
                _ = v14
            }
            v5[65535] = "a"
        }
    }
    if v1 {
        v15 := GenS1{F0: false}
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS2{F0: v2, F1: 927, F2: 0}
            _ = fi16
            _ = v17
        }
        v18 := GenS3{F0: 127, F1: -2147483648}
        _ = v15
        _ = v18
    }
    switch -122 {
    case 32768:
        var v19 string = "a"
        _ = v19
    }
    var v20 uint8 = 0
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 bool = v1
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := GenS4{F0: fi23, F1: 0, F2: ""}
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v4, v20, v4)
    v25 := GenS4{F0: -1, F1: 1, F2: "foo bar"}
    _ = len(v3)
    for fi26 := 0; fi26 < 4; fi26++ {
        switch fi26 {
        case -395:
            switch fi26 {
            case 0:
                v27 := GenS1{F0: false}
                _ = v27
            }
        }
        var v28 uint = v2
        v28 = 255
        _ = fi26
        _ = v28
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 bool = v1
        _ = fi29
        _ = v30
    }
    fmt.Println(v1)
    delete(v5, 0)
    v20 -= v20
    delete(v5, 970)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
