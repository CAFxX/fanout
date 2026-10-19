package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int8; F2 int64 }
func genFn0(p0 int64) int32 {
    if false {
        p0 -= 9223372036854775807
    } else {
        v7 := GenS1{F0: 0, F1: 32767, F2: (("hello" + "hello") + ("foo bar" + ""))}
        if (-2147483648 != -326) {
            v8 := []int{32767, -2147483648}
            var v9 uint32 = 0
            var v10 int8 = 127
            _ = v10
            _ = v8
            _ = v9
        }
        v11 := GenS1{F0: 2, F1: 127, F2: "hello"}
        _ = v11
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return -171
}

func main() {
    v1 := []int{-1}
    var v2 int16 = 32767
    v3 := make(map[string]bool)
    var v4 uint32 = 2147483648
    var v5 uint64 = 0
    var v6 int32 = 42
    delete(v3, (((("hello" + (("a" + "a") + "\n\t")) + ("hello" + ("a" + "foo bar"))) + (("hello" + ("foo bar" + "\n\t")) + ("\n\t" + "\n\t"))) + ("\n\t" + ("" + "hello"))))
    v12 := GenS2{F0: v2}
    fmt.Println(v4)
    if true {
        v13 := GenS3{F0: v5}
        _ = v13
    } else {
        var v14 int16 = (v2 & -324)
        v2 += 517
        switch -9223372036854775808 {
        case 0:
            delete(v3, "hello")
        case -2147483648:
            v15 := GenS2{F0: 699}
            _ = v15
        default:
            v1 = append(v1, 128)
        }
        _ = v14
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        v6 = 42
        delete(v3, "")
        _ = fi16
    }
    v2 = (v2 << ((549) & 63))
    var v17 string = ("a" + "")
    for fi18 := 0; fi18 < 6; fi18++ {
        if (0.1 != float64(float32(0.1))) {
            var v19 int32 = v6
            _ = v19
        } else {
            v4 = (^(v4))
            v2 = (v2 | (v2 >> ((65535) & 63)))
            var v20 float64 = 0.1
            _ = v20
        }
        v17 = v17
        var v21 uint = 517
        _ = fi18
        _ = v21
    }
    var v22 int = 32767
    switch v17 {
    case "a":
        v23 := GenS4{F0: true, F1: 100, F2: int64(v5)}
        _ = v23
    case "foo bar":
        v5 += 42
    }
    fmt.Println(v4, v5, v5)
    var v24 int64 = 0
    if true {
        switch "\n\t" {
        case "hello":
            _ = len(v1)
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    } else {
        fmt.Println(v17)
    }
    var v25 int16 = (v2 | -1)
    var v26 int = (v22 * (v22 ^ -9223372036854775808))
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
