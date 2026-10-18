package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() uint32 {
    var v7 int8 = -76
    v7 -= v7
    v8 := []int{449}
    switch 646 {
    case 127:
        var v9 float32 = 1e10
        _ = v9
    case -426:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        switch ("a" + "foo bar") {
        case "foo bar":
            v10 := GenS1{F0: -568}
            _ = v10
        default:
            for fi11 := 0; fi11 < 17; fi11++ {
                v7 = (v7 >> ((4294967295) & 63))
                _ = fi11
            }
        }
    }
    var v12 int16 = -1
    _ = v12
    _ = v7
    _ = v8
    return 2147483648
}
func genFn1() int {
    v13 := GenS1{F0: -247}
    var v14 uint8 = 128
    if (((((("a" + "") + "foo bar") > (("" + "hello") + (((("" + "hello") + "hello") + (("foo bar" + "foo bar") + "hello")) + "\n\t"))) || true) || true) || (4294967295 != 2)) {
        var v15 uint = 4294967295
        _ = v15
    } else {
        v16 := GenS1{F0: -307}
        _ = v16
    }
    _ = v13
    _ = v14
    return 2147483647
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 128
    var v3 int16 = -1
    var v4 float32 = (0.0 * ((-2.25 - 0.1) - -2.25))
    v5 := make(map[uint32]int)
    v6 := []string{}
    var v17 int = (v2 % ((v2) | 1))
    if false {
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS1{F0: -9223372036854775808}
            switch genFn1() {
            case 1:
                var v20 uint8 = 42
                _ = v20
            }
            _ = fi18
            _ = v19
        }
        v2 += v2
        fmt.Println(v17, v2)
    } else {
        var v21 bool = false
        _ = v21
    }
    if ((v4 - (0.0 - ((v4 * float32(0.0)) / float32(1.5)))) > v4) {
        var v22 float64 = (0.0 * ((float64(float32(float64(-2.25))) + float64(v4)) + (0.0 - float64(float32(0.1)))))
        var v23 int8 = 1
        var v24 uint64 = 18446744073709551615
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        delete(v5, 2)
        _ = fi25
    }
    v17 -= -821
    var v26 uint16 = 1
    fmt.Println(v3, v26)
    var v27 uint = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
