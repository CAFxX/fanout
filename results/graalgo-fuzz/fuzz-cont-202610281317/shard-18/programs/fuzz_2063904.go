package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
func genFn0() string {
    if (32767 == 255) {
        if (true || ((1e10 < 1.5) && (42 < -15))) {
            var v4 uint32 = 42
            _ = v4
        }
    } else {
        v5 := GenS1{F0: -2147483648, F1: 2147483647}
        var v6 float32 = (0.1 + -2.25)
        _ = v5
        _ = v6
    }
    v7 := []int{}
    if (!(2 > 4294967295)) {
        v8 := GenS2{F0: 2, F1: -729}
        _ = v8
    } else {
        var v9 uint = 65535
        v9 -= 2
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int64 = 1
        if len(v7) > 0 {
            v7[0] = (fi10 + (fi10 % (((-(fi10))) | 1)))
        }
        var v12 uint8 = 127
        _ = fi10
        _ = v11
        _ = v12
    }
    switch ("\n\t" + (("a" + "hello") + "a")) {
    case "foo bar":
        var v13 bool = (!false)
        _ = v13
    }
    _ = v7
    return ("a" + "")
}
func genFn1(p0 uint, p1 int64, p2 int16) uint16 {
    p2 = (p2 & 32767)
    p0 = p0
    p0 = p0
    v14 := GenS1{F0: 4294967295, F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 65535
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = float32((1.5 * (3.14159 / 1.5)))
    v3 := []int{127, 128}
    var v15 string = ("a" + "hello")
    switch -295 {
    case 4294967295:
        if true {
            if (-1 >= 100) {
                v1[1] = false
            }
            var v16 float32 = float32((-2.25 + ((-2.25 + -2.25) / 1.5)))
            _ = v16
        } else {
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS1{F0: 32768, F1: (fi17 + (fi17 + (^(fi17))))}
                v19 := GenS2{F0: 255, F1: -1}
                v20 := GenS2{F0: 42, F1: 286}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        }
    case 298:
        var v21 int16 = -7
        _ = v21
    }
    if true {
        var v22 int16 = -1
        var v23 float64 = 1.5
        _ = v22
        _ = v23
    }
    if false {
        var v24 float64 = float64(v2)
        for fi25 := 0; fi25 < 3; fi25++ {
            if false {
                v26 := GenS3{F0: -737}
                v1[219] = false
                _ = v26
            }
            _ = fi25
        }
        var v27 uint = 42
        _ = v24
        _ = v27
    } else {
        v2 = v2
        switch "foo bar" {
        case "foo bar":
            var v28 string = "hello"
            _ = v28
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = -892
            }
        case "a":
            var v29 int8 = 127
            _ = v29
        }
        v15 = (("hello" + v15) + "a")
    }
    var v30 int32 = -391
    v31 := GenS4{F0: 9223372036854775807}
    v15 = "a"
    var v32 float64 = -2.25
    var v33 string = v15
    var v34 uint = 425
    v33 = (v33 + v33)
    var v35 uint32 = 42
    var v36 uint32 = 741
    v2 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
