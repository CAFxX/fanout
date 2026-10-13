package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 string; F1 string; F2 int32 }
type GenS4 struct { F0 uint; F1 int32; F2 int8 }

func main() {
    var v1 string = "hello"
    var v2 uint = 42
    v3 := []string{(v1 + (v1 + v1))}
    var v4 uint = 65535
    var v5 uint64 = 2
    var v6 bool = (1 >= -119)
    v7 := GenS1{F0: -1, F1: true, F2: 1}
    fmt.Println(v2, v4, v5)
    switch 127 {
    case 2147483647:
        v8 := GenS1{F0: 81, F1: ((!(((1e10 * float64(-2.25)) - 1e10) > 1e10)) && v6), F2: 42}
        _ = v8
    case 128:
        fmt.Println(v5, v5)
    default:
        var v9 uint = v4
        _ = v9
    }
    fmt.Println(v6)
    v10 := GenS1{F0: -32768, F1: (-32768 > 42), F2: 65535}
    if v6 {
        if (!v6) {
            var v11 float32 = (0.0 - -2.25)
            v12 := GenS2{F0: 822, F1: v1}
            v1 = (v1 + v1)
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 int16 = -703
                _ = fi13
                _ = v14
            }
        }
        v1 = "hello"
        fmt.Println(v2, v5)
    } else {
        v15 := GenS3{F0: (v1 + ("\n\t" + "hello")), F1: (((((v1 + (((v1 + (v1 + (v1 + v1))) + v1) + ("foo bar" + v1))) + (v1 + v1)) + "\n\t") + v1) + "\n\t"), F2: -1}
        var v16 bool = (!v6)
        fmt.Println(v2)
        _ = v15
        _ = v16
    }
    switch -9223372036854775808 {
    case -876:
        v17 := GenS4{F0: 18446744073709551615, F1: -1, F2: -30}
        _ = v17
    case 32768:
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 uint64 = 2
                if (0 >= 42) {
                    var v21 float64 = (-2.25 + 1e10)
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
            v1 = "a"
            if v6 {
                var v22 string = ("\n\t" + (("\n\t" + v1) + v1))
                _ = v22
            } else {
                v6 = (9223372036854775807 <= fi18)
            }
            _ = fi18
        }
    }
    var v23 uint = (v2 & 255)
    var v24 int32 = -749
    if false {
        v2 = 147
        v23 += (v4 * 40)
    } else {
        v2 = 18446744073709551615
        var v25 uint32 = 42
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
