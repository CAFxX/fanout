package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint64; F2 int8 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    var v3 int8 = -66
    var v4 int16 = 0
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            if v2 {
                v7 := GenS1{F0: 0}
                fmt.Println(v2, v4, fi6)
                v4 = 501
                _ = v7
            } else {
                v4 -= v4
                v8 := []int{}
                _ = v8
            }
            var v9 float32 = 0.0
            _ = fi6
            _ = v9
        }
        if false {
            var v10 int64 = 0
            fmt.Println(v3, fi5)
            _ = v10
        } else {
            var v11 int = (^(fi5))
            _ = v11
        }
        fmt.Println(v3)
        _ = fi5
    }
    var v12 uint = 18446744073709551615
    v13 := []int{354, 921, 0, -823}
    var v14 int8 = (v3 & (v3 % ((-128) | 1)))
    switch int(v14) {
    case 1:
        _ = len(v1)
    }
    fmt.Println(v3, v2)
    v4 += v4
    var v15 int16 = v4
    var v16 uint32 = 1
    if (!(("a" + ("a" + "")) != "")) {
        for fi17 := 0; fi17 < 6; fi17++ {
            v15 = (v4 / ((-32768) | 1))
            if true {
                var v18 int16 = (^(v4))
                var v19 uint16 = 32767
                _ = v18
                _ = v19
            }
            _ = fi17
        }
        if (1.5 == 0.0) {
            if false {
                v12 = (v12 & (v12 >> ((889) & 63)))
                v20 := GenS2{F0: "foo bar", F1: 589, F2: v3}
                _ = v20
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                v22 := GenS3{F0: "hello", F1: 1}
                _ = fi21
                _ = v22
            }
            v23 := GenS3{F0: "", F1: (-(v15))}
            _ = v23
        }
        v24 := v1[("foo bar" + ("" + ""))]
        _ = v24
    } else {
        if (0 != uint64(v16)) {
            v2 = v2
        }
        _ = len(v1)
        v16 = v16
    }
    switch ("\n\t" + "") {
    case "":
        v12 = uint(v16)
    case "\n\t":
        v25 := GenS4{F0: 42, F1: v16}
        _ = v25
    }
    if v2 {
        var v26 uint16 = 32767
        _ = v26
    } else {
        v2 = false
        v14 = v3
        var v27 int16 = (v4 & v4)
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
