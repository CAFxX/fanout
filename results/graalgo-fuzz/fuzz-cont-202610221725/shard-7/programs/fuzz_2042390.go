package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int8 }
func genFn0(p0 uint, p1 int16) uint64 {
    fmt.Println(p0)
    v4 := GenS1{F0: (p0 % ((111) | 1)), F1: "\n\t", F2: 316}
    fmt.Println(p1, p1)
    switch 32768 {
    case 9223372036854775807:
        p1 = 32767
    case -521:
        fmt.Println(p1)
    case 127:
        var v5 uint32 = 1
        _ = v5
    default:
        var v6 int8 = 0
        _ = v6
    }
    if true {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float64 = (3.14159 * (((0.0 + 1e10) * float64(((-2.25 - ((0.1 * 0.1) - 1e10)) - (0.0 * -2.25)))) - 3.14159))
            var v9 uint = 42
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(p1, p0)
        fmt.Println(p0, p1)
    } else {
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}

func main() {
    var v1 bool = (true || true)
    var v2 int16 = 0
    var v3 uint16 = 32767
    v10 := GenS2{F0: (^(v3)), F1: 4294967295, F2: 50}
    for fi11 := 0; fi11 < 9; fi11++ {
        fmt.Println(v2, v3)
        _ = fi11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        if (!(((((("hello" + "") + "") + "hello") + "hello") + "\n\t") >= ("hello" + "a"))) {
            for fi13 := 0; fi13 < 1; fi13++ {
                v1 = true
                _ = fi13
            }
            if (0.1 <= 1e10) {
                var v14 int32 = 19
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            if v1 {
                v3 += v3
                var v16 string = "hello"
                var v17 uint32 = 65535
                _ = v16
                _ = v17
            } else {
                fmt.Println(fi12)
                fmt.Println(fi15, v2)
                fmt.Println(fi12)
                var v18 int32 = 42
                var v19 uint = 18446744073709551615
                _ = v18
                _ = v19
            }
            var v20 uint32 = 666
            var v21 bool = (1.5 != 0.0)
            switch "\n\t" {
            case "a":
                var v22 int8 = 0
                _ = v22
            default:
                var v23 float32 = float32(0.0)
                _ = v23
            }
            v21 = (-1 > 42)
            v21 = (-1 < -451)
            _ = fi15
            _ = v20
            _ = v21
        }
        v24 := GenS1{F0: 42, F1: "hello", F2: (v3 - (^(v3)))}
        _ = fi12
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        fmt.Println(v2)
        var v26 uint8 = 244
        fmt.Println(v26, v2)
        _ = fi25
        _ = v26
    }
    switch ("a" + "\n\t") {
    case "\n\t":
        for fi27 := 0; fi27 < 4; fi27++ {
            for fi28 := 0; fi28 < 11; fi28++ {
                fmt.Println(v3)
                v3 += v3
                _ = fi28
            }
            if (1 <= 1) {
                v29 := []int{fi27, (fi27 + fi27)}
                _ = v29
            }
            var v30 int64 = int64(v3)
            _ = fi27
            _ = v30
        }
    default:
        v3 = 1
    }
    v3 = 1
    var v31 float64 = (0.0 * 0.0)
    if v1 {
        v32 := GenS1{F0: 4294967295, F1: (("\n\t" + (("hello" + "\n\t") + "\n\t")) + "a"), F2: (v3 & v3)}
        _ = v32
    } else {
        v2 += v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
}
