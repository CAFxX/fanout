package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int8, p1 uint32) uint64 {
    if false {
        p1 -= p1
        switch -2147483648 {
        case 713:
            p1 = (^(p1))
        case 0:
            v5 := GenS1{F0: 2}
            _ = v5
        }
        switch int(p0) {
        case -455:
            var v6 float32 = float32((1e10 / 1e10))
            _ = v6
        case 42:
            if ("a" != ("hello" + "foo bar")) {
                p0 = 127
                var v7 uint64 = 4294967295
                if false {
                    var v8 float32 = float32((0.0 + ((((float64(0.1) - 3.14159) * (((3.14159 / 1e10) / ((((-2.25 + (1e10 - 3.14159)) / 1.5) * (-2.25 + 3.14159)) - (-2.25 / 1e10))) / ((((float64((0.1 * 3.14159)) * -2.25) / ((float64(-2.25) / (1.5 / float64(0.1))) / 1e10)) + (1e10 - (float64(0.0) + (3.14159 / ((-2.25 - 1.5) / 1.5))))) + (1e10 + 0.0)))) / 0.1) - (3.14159 + (0.1 * (float64(0.1) + 1.5))))))
                    _ = v8
                }
                _ = v7
            } else {
                if (!true) {
                    fmt.Println(p1, p0)
                }
                for fi9 := 0; fi9 < 12; fi9++ {
                    fmt.Println(p0, p0, p1)
                    p0 -= (p0 | 95)
                    _ = fi9
                }
                var v10 int16 = 948
                _ = v10
            }
        default:
            var v11 uint32 = 284
            _ = v11
        }
    } else {
        fmt.Println(p1, p1)
    }
    p1 += (p1 << ((65535) & 63))
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 65535
    var v2 uint16 = 474
    v3 := make(map[string]int)
    v4 := []int{}
    fmt.Println(v1, v1)
    fmt.Println(v1, v2)
    var v12 float64 = 0.1
    switch "hello" {
    case "foo bar":
        v1 -= v1
    case "":
        for fi13 := 0; fi13 < 9; fi13++ {
            v3[(("foo bar" + "\n\t") + (("foo bar" + (("" + "hello") + ("hello" + "hello"))) + ("hello" + "a")))] = 42
            var v14 int8 = 127
            _ = fi13
            _ = v14
        }
    default:
        fmt.Println(v1, v1)
    }
    if false {
        if false {
            if false {
                var v15 uint = 42
                _ = v15
            }
            if (1 == 176) {
                var v16 float64 = v12
                _ = v16
            } else {
                switch 42 {
                case 406:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 42:
                    var v17 float64 = v12
                    _ = v17
                }
                v1 = (v1 / ((597) | 1))
            }
            if true {
                _ = len(v3)
                v18 := GenS2{F0: 25}
                _ = v18
            } else {
                v2 = v2
            }
        }
        if len(v4) > 0 {
            v4[0] = -341
        }
    } else {
        v12 -= 1e10
        switch 136 {
        case -388:
            fmt.Println(v2)
        case 32767:
            var v19 int8 = 100
            _ = v19
        case -1:
            for fi20 := 0; fi20 < 16; fi20++ {
                switch -389 {
                case -1:
                    delete(v3, "foo bar")
                case 922:
                    var v21 int64 = -9223372036854775808
                    _ = v21
                }
                v12 = 1.5
                var v22 int32 = -1
                _ = fi20
                _ = v22
            }
        }
    }
    var v23 int8 = 100
    var v24 int16 = -511
    v25 := GenS1{F0: 42}
    v12 += (((v12 / (float64(0.0) / v12)) * 3.14159) + v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
