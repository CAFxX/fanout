package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 1
    v3 := []string{(("\n\t" + "") + (("a" + (("hello" + "hello") + "")) + "hello")), ("foo bar" + "foo bar"), ("" + "foo bar")}
    v4 := make(map[int]string)
    v5 := make(map[string]uint32)
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: (false && true)}
        var v8 float64 = ((3.14159 / 1.0) + -2.25)
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: ""}
    v2 += 127
    v2 = v2
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(v2, fi10)
        _ = fi10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        if true {
            var v12 int64 = -417
            if true {
                switch 1 {
                case 42:
                    var v13 string = ""
                    _ = v13
                default:
                    v12 += int64(fi11)
                }
            }
            v14 := GenS1{F0: (!true)}
            _ = v12
            _ = v14
        } else {
            v15 := GenS1{F0: (!true)}
            _ = v15
        }
        _ = fi11
    }
    var v16 uint64 = 650
    v17 := v4[42]
    if true {
        var v18 int8 = -25
        var v19 float32 = (((-2.25 + -2.25) / 3.14159) / -2.25)
        _ = v18
        _ = v19
    }
    var v20 uint = 262
    if true {
        if (true && ((-766 <= -353) || ((!false) || ((v20 / ((v20) | 1)) == v20)))) {
            var v21 int32 = -2147483648
            fmt.Println(v17)
            for fi22 := 0; fi22 < 19; fi22++ {
                for fi23 := 0; fi23 < 4; fi23++ {
                    fmt.Println(fi23, v17)
                    if false {
                        var v24 float64 = -2.25
                        _ = v24
                    }
                    _ = fi23
                }
                fmt.Println(fi22, v20)
                _ = fi22
            }
            _ = v21
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    } else {
        v1[833] = "\n\t"
        var v25 string = (v17 + v17)
        _ = v25
    }
    var v26 int16 = -965
    if true {
        fmt.Println(v26, v26, v16)
    }
    v27 := v5["foo bar"]
    var v28 int64 = 377
    var v29 int = 1
    fmt.Println(v26, v28)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
