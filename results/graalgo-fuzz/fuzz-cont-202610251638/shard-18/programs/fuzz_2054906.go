package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }
type GenS2 struct { F0 int16 }
func genFn0() int16 {
    var v6 float64 = ((((-2.25 + 1.5) - 1e10) - -2.25) / 1.0)
    v7 := GenS1{F0: "", F1: 42, F2: 42}
    v6 = v6
    v8 := GenS1{F0: ("a" + ""), F1: 1, F2: -1}
    _ = v6
    _ = v7
    _ = v8
    return 1000
}
func genFn1(p0 uint64) int8 {
    v9 := []int{}
    p0 = p0
    var v10 uint16 = uint16(p0)
    _ = p0
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 bool = (-118 < 127)
    v2 := []string{("\n\t" + ("\n\t" + "hello")), "", ("\n\t" + "foo bar"), "a"}
    v3 := []int{-450}
    v4 := []string{"", (("foo bar" + "hello") + ("foo bar" + ("\n\t" + (("" + "") + "foo bar"))))}
    v5 := make(map[string]string)
    if (32768 == 457) {
        var v11 uint16 = 1
        _ = v11
    } else {
        delete(v5, "a")
        for fi12 := 0; fi12 < 15; fi12++ {
            if v1 {
                v1 = v1
            }
            v4 = append(v4, ("" + "a"))
            _ = fi12
        }
    }
    delete(v5, "hello")
    if v1 {
        fmt.Println(v1)
    } else {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint8 = 66
            v14 -= v14
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 bool = true
                v17 := GenS1{F0: "\n\t", F1: 487, F2: -32768}
                if true {
                    v18 := GenS2{F0: genFn0()}
                    v2 = append(v2, "")
                    _ = v18
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi13
            _ = v14
        }
    }
    v1 = (1 > 1)
    if (3 < 190) {
        var v19 uint16 = 32767
        var v20 int16 = -32768
        var v21 int = 32767
        _ = v19
        _ = v20
        _ = v21
    } else {
        switch -9223372036854775808 {
        case 9223372036854775807:
            v1 = (-1 >= -2147483648)
        case 1:
            v22 := GenS2{F0: -613}
            _ = v22
        case -519:
            if len(v3) > 0 {
                v3[0] = 32768
            }
        default:
            var v23 string = ("a" + "")
            _ = v23
        }
        v24 := GenS2{F0: -32768}
        _ = v24
    }
    v1 = (65535 > 0)
    var v25 int16 = 32767
    for fi26 := 0; fi26 < 17; fi26++ {
        if v1 {
            v27 := GenS1{F0: "foo bar", F1: 42, F2: (v25 & v25)}
            _ = len(v4)
            var v28 uint16 = 255
            _ = v27
            _ = v28
        } else {
            v25 += (^(v25))
            var v29 uint16 = 363
            _ = v29
        }
        var v30 int = fi26
        _ = fi26
        _ = v30
    }
    var v31 uint16 = 2
    fmt.Println(v25, v1, v31)
    if (!((v31 - v31) == (v31 << ((542) & 63)))) {
        fmt.Println(v31)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
