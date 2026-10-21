package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0(p0 uint8) uint8 {
    v6 := []int{563}
    p0 = p0
    _ = p0
    _ = v6
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 42
    v3 := []string{"a", "hello"}
    v4 := make(map[uint32]int)
    var v5 int8 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    delete(v4, 1)
    var v7 int16 = 32767
    v8 := GenS1{F0: "\n\t", F1: false}
    if true {
        _ = len(v3)
    }
    v9 := GenS1{F0: (("\n\t" + "a") + "hello"), F1: false}
    var v10 int16 = (-(v7))
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: "foo bar", F1: false}
        _ = fi11
        _ = v12
    }
    v3 = append(v3, (("\n\t" + "a") + "hello"))
    if true {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int64 = 1
            _ = fi13
            _ = v14
        }
        _ = len(v3)
    }
    fmt.Println(v5, v5)
    switch ("\n\t" + "") {
    case "hello":
        if true {
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 float64 = 1.5
                _ = fi15
                _ = v16
            }
            switch 128 {
            case 32768:
                var v17 int32 = -2147483648
                _ = v17
            case -686:
                var v18 int16 = -32768
                _ = v18
            }
            fmt.Println(v10, v10, v10)
        } else {
            if (4294967295 > 983) {
                var v19 uint64 = 221
                _ = v19
            } else {
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 int16 = -1
                    if (v5 < -1) {
                        fmt.Println(v10, v10)
                        v22 := GenS1{F0: (("hello" + ("foo bar" + ("hello" + ""))) + ((("foo bar" + "hello") + "") + ("hello" + (("\n\t" + "hello") + "a")))), F1: true}
                        _ = v22
                    } else {
                        var v23 uint8 = (v2 + v2)
                        _ = v23
                    }
                    v3 = append(v3, ("a" + "a"))
                    _ = fi20
                    _ = v21
                }
                v2 = (v2 + genFn0(v2))
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                if (!((float64((1.5 + -2.25)) - (float64(((0.0 / 3.14159) + (-2.25 / 1e10))) * 1e10)) <= (float64(((float32(0.0) + (1.5 / 1e10)) - ((0.1 + 0.1) * (3.14159 - 1.5)))) - 1e10))) {
                    var v25 bool = true
                    _ = len(v4)
                    _ = v25
                }
                for fi26 := 0; fi26 < 2; fi26++ {
                    v10 = 42
                    _ = fi26
                }
                fmt.Println(v10, v2)
                _ = fi24
            }
            v27 := GenS1{F0: "\n\t", F1: false}
            _ = v27
        }
    case "foo bar":
        var v28 uint = 42
        _ = v28
    case "":
        var v29 string = "hello"
        _ = v29
    default:
        var v30 bool = (42 > 264)
        _ = v30
    }
    fmt.Println(v2, v10, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
