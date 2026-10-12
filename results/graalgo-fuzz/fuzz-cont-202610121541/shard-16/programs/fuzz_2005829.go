package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 string, p1 string) uint32 {
    fmt.Println(p0)
    v6 := GenS1{F0: 2147483647, F1: 0}
    p1 = ("a" + "a")
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1() uint16 {
    var v7 float64 = (1.5 * (float64((1.5 + (float32(-2.25) + 0.0))) * (0.1 + (-2.25 * ((float64((1.5 * (1e10 - 3.14159))) - 0.1) - 1.5)))))
    v7 = (v7 * v7)
    v8 := []int{-878, 32768, 32767, -2147483648}
    var v9 bool = ("hello" == ("hello" + ""))
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 323
    v3 := []string{""}
    var v4 int16 = 32767
    var v5 float32 = (1.5 / float32(-2.25))
    if false {
        if ((("hello" + "hello") + ("a" + "foo bar")) < "\n\t") {
            if true {
                switch "foo bar" {
                case "a":
                    var v10 uint32 = 2
                    _ = v10
                case "":
                    fmt.Println(v4, v4, v2)
                default:
                    var v11 bool = true
                    _ = v11
                }
                var v12 int8 = 0
                _ = v12
            }
        }
    }
    fmt.Println(v2, v2, v2)
    var v13 int8 = 116
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int8 = v13
        if true {
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS1{F0: 2147483647, F1: 842}
                var v18 float32 = float32(-2.25)
                v3 = append(v3, "\n\t")
                _ = fi16
                _ = v17
                _ = v18
            }
            switch "\n\t" {
            case "\n\t":
                for fi19 := 0; fi19 < 7; fi19++ {
                    v20 := GenS1{F0: 770, F1: 1}
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                    var v21 int8 = v15
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        }
        _ = fi14
        _ = v15
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v2 = v2
        v23 := GenS2{F0: 42}
        _ = fi22
        _ = v23
    }
    var v24 float32 = v5
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 int32 = 444
        _ = fi25
        _ = v26
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
