package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[int]string)
    var v3 uint = 857
    v4 := make(map[string]int)
    var v5 bool = (2 >= 416)
    var v6 float32 = 0.1
    if (255 == 810) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v7 := GenS1{F0: "\n\t"}
    var v8 uint16 = 0
    for fi9 := 0; fi9 < 8; fi9++ {
        v6 += ((((-2.25 - v6) * 1e10) - v6) + v6)
        var v10 uint32 = 2
        _ = fi9
        _ = v10
    }
    var v11 bool = v5
    if (100 != -21) {
        v6 = (v6 / (((v6 * v6) / 1.5) - v6))
        v6 = (0.0 / v6)
        v1 = append(v1, ("hello" + ""))
    } else {
        fmt.Println(v6, v11, v5)
        v12 := GenS1{F0: "\n\t"}
        _ = v12
    }
    var v13 int8 = 42
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS2{F0: 42}
        var v16 int16 = 1000
        _ = fi14
        _ = v15
        _ = v16
    }
    v6 = 0.1
    v1 = append(v1, ("" + (("foo bar" + "hello") + ("foo bar" + ("a" + ("" + "foo bar"))))))
    var v17 uint = v3
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 int = fi18
        _ = fi18
        _ = v19
    }
    var v20 int32 = 1
    switch 32767 {
    case -9223372036854775808:
        for fi21 := 0; fi21 < 13; fi21++ {
            v4["a"] = fi21
            if (((-2.25 - ((1e10 / -2.25) + (v6 + 0.0))) != float32(0.0)) && v11) {
                switch 32767 {
                case 32768:
                    var v22 float32 = 1e10
                    _ = v22
                case 1:
                    var v23 uint64 = 2
                    _ = v23
                case 32767:
                    var v24 uint = v3
                    _ = v24
                default:
                    switch ("\n\t" + "") {
                    case "hello":
                        var v25 int = fi21
                        _ = v25
                    }
                }
                var v26 int16 = 42
                v11 = v5
                _ = v26
            }
            _ = fi21
        }
    case 0:
        if true {
            v5 = (-692 < 2147483647)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v13
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
