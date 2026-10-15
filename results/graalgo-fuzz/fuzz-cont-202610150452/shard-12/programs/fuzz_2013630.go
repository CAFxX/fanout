package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 int8; F1 int16; F2 int64 }
type GenS3 struct { F0 uint32; F1 bool }

func main() {
    var v1 int8 = 42
    v2 := make(map[string]uint32)
    var v3 int = -824
    v4 := make(map[string]string)
    var v5 uint32 = 1
    if false {
        v6 := GenS1{F0: 1, F1: "", F2: (v3 < 955)}
        _ = len(v2)
        _ = v6
    }
    var v7 bool = true
    var v8 uint = 1
    for fi9 := 0; fi9 < 14; fi9++ {
        switch "a" {
        case "hello":
            v8 = v8
        case "a":
            for fi10 := 0; fi10 < 8; fi10++ {
                v7 = v7
                var v11 uint8 = 2
                v7 = (127 >= v11)
                _ = fi10
                _ = v11
            }
        default:
            v5 -= 2147483647
        }
        _ = fi9
    }
    v3 -= v3
    var v12 float64 = float64(-2.25)
    v13 := GenS1{F0: 128, F1: "foo bar", F2: (979 >= 0)}
    v14 := GenS1{F0: 255, F1: "\n\t", F2: v7}
    switch v3 {
    case -405:
        var v15 int64 = -859
        _ = v15
    case 32768:
        if (!true) {
            v5 -= (^(v5))
        } else {
            v8 = v8
            v16 := GenS1{F0: 119, F1: "hello", F2: v7}
            _ = v16
        }
    case 2147483647:
        v7 = (!(65535 != (v8 % (((v8 + 18446744073709551615)) | 1))))
    default:
        if false {
            v3 = v3
            var v17 string = "hello"
            _ = v17
        } else {
            v18 := GenS1{F0: 113, F1: (("\n\t" + ("" + "hello")) + ("" + ("foo bar" + ""))), F2: (32767 == 32767)}
            for fi19 := 0; fi19 < 11; fi19++ {
                for fi20 := 0; fi20 < 4; fi20++ {
                    if v7 {
                        v3 -= (v3 | fi19)
                        delete(v4, ((("\n\t" + "\n\t") + "hello") + ((("foo bar" + ("" + "")) + ("a" + "hello")) + ("\n\t" + ("" + ("hello" + "foo bar"))))))
                    } else {
                        fmt.Println(fi20, v8, v8)
                        fmt.Println(fi19, v8)
                        v21 := GenS2{F0: -1, F1: -32768, F2: 1}
                        _ = v21
                    }
                    fmt.Println(v8, v8)
                    _ = fi20
                }
                _ = fi19
            }
            _ = v18
        }
    }
    v22 := GenS1{F0: 128, F1: "", F2: true}
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS3{F0: 4294967295, F1: v7}
        v8 = 815
        var v25 bool = false
        _ = fi23
        _ = v24
        _ = v25
    }
    delete(v4, "foo bar")
    var v26 int16 = -961
    if (-795 >= 345) {
        v26 = v26
        for fi27 := 0; fi27 < 12; fi27++ {
            v12 = v12
            _ = fi27
        }
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 int = v3
            _ = fi28
            _ = v29
        }
    } else {
        fmt.Println(v5, v12, v7)
    }
    v30 := GenS1{F0: 1, F1: "foo bar", F2: (v12 > float64((1.5 / -2.25)))}
    v3 += 127
    v31 := []int{v3, (v3 * v3), 979}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
