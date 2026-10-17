package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
type GenS2 struct { F0 int8 }
func genFn0() uint16 {
    var v7 int64 = 618
    v8 := []int{-9223372036854775808, -76}
    var v9 string = ("hello" + "\n\t")
    var v10 int32 = 2147483647
    if (2 < 2) {
        fmt.Println(v7, v10, v10)
    } else {
        v10 = v10
        if true {
            var v11 uint8 = 255
            _ = v11
        } else {
            var v12 float32 = float32((0.0 + ((1e10 - -2.25) * (3.14159 * 0.0))))
            v12 = ((v12 - (float32(0.0) / ((v12 * (0.0 + 3.14159)) * v12))) + v12)
            switch "hello" {
            case "hello":
                switch ("foo bar" + (v9 + v9)) {
                case "":
                    var v13 int8 = -1
                    _ = v13
                case "foo bar":
                    v7 = v7
                default:
                    var v14 float32 = v12
                    _ = v14
                }
            case "":
                var v15 bool = true
                _ = v15
            }
            _ = v12
        }
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{9223372036854775807, 42, 32768}
    v3 := []string{("\n\t" + "foo bar"), "foo bar", "foo bar", ("a" + "")}
    v4 := []string{}
    var v5 bool = (1 < -1)
    v6 := []int{32767, 128, 4294967295, 159}
    var v16 uint64 = 422
    delete(v1, 32767)
    v16 += v16
    delete(v1, 32767)
    v17 := GenS1{F0: 0, F1: -121, F2: 942}
    switch 32768 {
    case -9223372036854775808:
        var v18 int64 = 42
        _ = v18
    }
    v3 = append(v3, "hello")
    v19 := GenS2{F0: -1}
    v20 := GenS2{F0: 0}
    v21 := GenS2{F0: 35}
    for fi22 := 0; fi22 < 8; fi22++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(fi22, v5, fi22)
        _ = fi22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 int8 = 42
        _ = fi23
        _ = v24
    }
    if ((!v5) || v5) {
        for fi25 := 0; fi25 < 13; fi25++ {
            fmt.Println(v5, v5, fi25)
            var v26 uint32 = 62
            var v27 string = ""
            _ = fi25
            _ = v26
            _ = v27
        }
        fmt.Println(v16, v16)
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        v16 += v16
        _ = len(v4)
        _ = fi28
    }
    switch 9223372036854775807 {
    case 128:
        for fi29 := 0; fi29 < 3; fi29++ {
            v5 = (v5 && v5)
            fmt.Println(v16, v16)
            _ = fi29
        }
    case 427:
        v16 = v16
    case 9223372036854775807:
        v16 += v16
    default:
        v16 = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
