package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 bool; F2 string }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 string }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 19; fi4++ {
        v2[fi4] = 2
        for fi5 := 0; fi5 < 8; fi5++ {
            v1 = 0.0
            delete(v2, (fi4 % ((fi5) | 1)))
            v6 := GenS1{F0: 42}
            _ = fi5
            _ = v6
        }
        var v7 int64 = -399
        _ = fi4
        _ = v7
    }
    v1 -= (float32((3.14159 + 0.0)) * ((1.5 / v1) * ((((float32((1.5 * 0.1)) + float32(0.1)) * v1) / v1) / float32(3.14159))))
    delete(v2, -1)
    v8 := []int{42, -9223372036854775808}
    v9 := GenS2{F0: -1, F1: false, F2: ""}
    switch 32768 {
    case 32768:
        v10 := GenS1{F0: 65535}
        _ = v10
    case 1:
        v11 := GenS1{F0: 1}
        _ = v11
    default:
        if (0 != 211) {
            if len(v8) > 0 {
                v8[0] = -1
            }
            if ((1 < 1) && (849 != 683)) {
                if true {
                    v12 := GenS2{F0: 0, F1: (false && (-20 != 0)), F2: "hello"}
                    _ = v12
                }
                v1 = (3.14159 * v1)
            } else {
                var v13 int8 = 42
                var v14 uint = 255
                _ = v13
                _ = v14
            }
        } else {
            v15 := GenS1{F0: 0}
            v16 := GenS3{F0: 0, F1: true}
            _ = v15
            _ = v16
        }
    }
    switch 32767 {
    case 4294967295:
        for fi17 := 0; fi17 < 5; fi17++ {
            v1 += v1
            fmt.Println(fi17, fi17)
            _ = fi17
        }
    case 32768:
        v18 := GenS1{F0: 32767}
        _ = v18
    case 550:
        var v19 int8 = 3
        _ = v19
    default:
        var v20 string = ("hello" + ("hello" + "\n\t"))
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v1 -= 1e10
        v22 := GenS4{F0: false}
        _ = fi21
        _ = v22
    }
    if false {
        v3[(("\n\t" + "hello") + ("" + "foo bar"))] = ("foo bar" + ("a" + ("foo bar" + "a")))
        var v23 uint64 = 2
        _ = v23
    }
    var v24 int = 42
    var v25 uint16 = 65535
    if (!true) {
        var v26 int8 = -1
        v26 += 0
        _ = v26
    } else {
        var v27 string = "hello"
        _ = v27
    }
    var v28 float32 = float32(0.0)
    var v29 uint32 = 2147483647
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := GenS5{F0: false, F1: "a"}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 bool = (!false)
        for fi34 := 0; fi34 < 8; fi34++ {
            var v35 uint16 = v25
            _ = fi34
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    fmt.Println(v28, v25)
    v25 -= (v25 - (^(v25)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v8
    _ = v9
}
