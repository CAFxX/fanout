package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string }
func genFn0(p0 uint8, p1 string) string {
    var v5 uint = 65535
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int32 = -403
        fmt.Println(p1)
        fmt.Println(p0, fi6)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 0, F1: 127}
    p1 = ("foo bar" + "hello")
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return "hello"
}
func genFn1() uint32 {
    var v9 float32 = 1.5
    var v10 int = 1
    for fi11 := 0; fi11 < 16; fi11++ {
        if (114 < 0) {
            fmt.Println(fi11, fi11)
        } else {
            switch v10 {
            case -650:
                v10 += v10
            default:
                v9 = ((1e10 / v9) * -2.25)
            }
            v12 := []int{fi11, (fi11 / ((fi11) | 1))}
            _ = v12
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            v10 = v10
            _ = fi13
        }
        if true {
            v14 := GenS1{F0: -1, F1: 42}
            _ = v14
        } else {
            v15 := GenS2{F0: "\n\t"}
            _ = v15
        }
        _ = fi11
    }
    fmt.Println(v10, v9)
    fmt.Println(v10, v10)
    _ = v10
    _ = v9
    return 690
}

func main() {
    var v1 uint8 = 42
    var v2 int64 = 629
    var v3 int32 = 1
    var v4 uint8 = v1
    fmt.Println(v3, v1, v1)
    v16 := []int{}
    v17 := GenS2{F0: ("a" + (("hello" + "") + ("hello" + ((("foo bar" + "foo bar") + "") + (("a" + "\n\t") + "")))))}
    fmt.Println(v2)
    v1 -= (v4 - v1)
    if false {
        fmt.Println(v2)
        var v18 uint = 94
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 bool = true
            v1 = 82
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        for fi21 := 0; fi21 < 8; fi21++ {
            v2 = (v2 - 681)
            _ = fi21
        }
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    var v22 int8 = 100
    v23 := GenS2{F0: "hello"}
    if (876 < 355) {
        if len(v16) > 0 {
            v16[0] = 673
        }
        var v24 string = "foo bar"
        _ = v24
    } else {
        v22 -= (v22 ^ v22)
        _ = len(v16)
        var v25 uint32 = 2
        _ = v25
    }
    var v26 uint8 = 127
    if false {
        if len(v16) > 0 {
            v16[0] = 32767
        }
        if len(v16) > 0 {
            v16[0] = 120
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
