package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8; F2 string }
type GenS3 struct { F0 uint32; F1 uint8; F2 int16 }
func genFn0(p0 uint32, p1 string) int {
    switch "foo bar" {
    case "hello":
        v7 := []int{1, 128}
        _ = v7
    }
    v8 := []int{1, 224, 32768}
    var v9 uint64 = 137
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -341
}
func genFn1(p0 uint, p1 string, p2 uint8) uint32 {
    fmt.Println(p1, p2)
    var v10 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 600
}

func main() {
    var v1 float64 = 1.5
    var v2 int32 = 1
    var v3 string = ("a" + ("\n\t" + "hello"))
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]string)
    var v6 uint = 380
    if true {
        var v11 int8 = 77
        var v12 int64 = 305
        _ = v11
        _ = v12
    } else {
        if false {
            fmt.Println(v6, v3, v3)
        }
    }
    v13 := []int{0}
    if true {
        if (2147483648 < 4294967295) {
            var v14 int64 = -31
            _ = v14
        } else {
            if len(v13) > 0 {
                v13[0] = 0
            }
        }
    } else {
        v15 := GenS1{F0: 682}
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 11; fi17++ {
            if (!false) {
                var v18 int16 = 42
                _ = v18
            }
            _ = fi17
        }
        _ = fi16
    }
    if (!true) {
        delete(v5, 42)
        var v19 int = -1
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := GenS1{F0: 2}
            var v22 bool = (fi20 > (fi20 ^ (^(fi20))))
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v19
    }
    var v23 uint8 = 127
    switch 32768 {
    case -1:
        var v24 uint8 = v23
        _ = v24
    case 32767:
        v25 := GenS2{F0: -960, F1: 22, F2: v3}
        _ = v25
    default:
        if true {
            v26 := GenS1{F0: 4294967295}
            _ = v26
        }
    }
    var v27 int8 = -17
    v28 := GenS3{F0: genFn1(v6, v3, 171), F1: v23, F2: 1000}
    fmt.Println(v1, v27, v27)
    var v29 int32 = 1
    switch -2147483648 {
    case 0:
        var v30 string = v3
        _ = v30
    }
    v31 := GenS3{F0: genFn1((v6 & v6), v3, v23), F1: (^(v23)), F2: int16(v29)}
    var v32 uint16 = 65535
    var v33 int64 = int64(v6)
    var v34 string = ((v3 + v3) + v3)
    if true {
        v6 -= 42
        var v35 uint64 = 2
        _ = v35
    } else {
        var v36 int = 32768
        if (((1 >= 0) && false) && (((3.14159 + float64(1e10)) - (0.0 / v1)) == float64(1.5))) {
            v4[4294967295] = (!(true && false))
            v37 := GenS2{F0: v36, F1: v27, F2: v34}
            _ = v37
        }
        var v38 bool = ((v33 + v33) == int64(v23))
        _ = v36
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
