package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string }
func genFn0() string {
    for fi6 := 0; fi6 < 19; fi6++ {
        switch "" {
        case "foo bar":
            var v7 int32 = -540
            _ = v7
        case "a":
            var v8 string = "hello"
            _ = v8
        default:
            for fi9 := 0; fi9 < 18; fi9++ {
                v10 := []int{1, 1}
                var v11 uint8 = 127
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        v12 := GenS1{F0: 12, F1: 0}
        var v13 int64 = 42
        _ = fi6
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: -1, F1: 42}
    if (-9223372036854775808 >= -39) {
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 int32 = 1
            if (true && (fi15 < fi15)) {
                var v17 uint = 18446744073709551615
                v16 -= (^(v16))
                v16 -= (v16 >> (((v17 * v17)) & 63))
                _ = v17
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 int8 = -84
                _ = fi18
                _ = v19
            }
            _ = fi15
            _ = v16
        }
        v20 := []int{574, -620, 2147483647}
        _ = v20
    }
    var v21 uint16 = 1
    _ = v14
    _ = v21
    return "hello"
}

func main() {
    var v1 int8 = -128
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    var v4 float32 = 3.14159
    var v5 int8 = v2
    for fi22 := 0; fi22 < 4; fi22++ {
        if (42 >= v1) {
            v23 := v3[1]
            _ = v23
        } else {
            v1 -= (v2 + (-(v1)))
        }
        v24 := []int{(fi22 >> ((4294967295) & 63)), (-(fi22)), -9223372036854775808}
        if false {
            for fi25 := 0; fi25 < 5; fi25++ {
                v4 -= v4
                var v26 int16 = 388
                _ = fi25
                _ = v26
            }
            if (v4 == 1e10) {
                fmt.Println(v5)
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
            }
        } else {
            if (0 > 14) {
                if true {
                    for fi27 := 0; fi27 < 9; fi27++ {
                        v28 := GenS1{F0: 1, F1: 7}
                        var v29 bool = true
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
            }
            _ = len(v24)
            v30 := GenS1{F0: 42, F1: 230}
            _ = v30
        }
        _ = fi22
        _ = v24
    }
    if true {
        if (!false) {
            var v31 int16 = 1
            v31 -= (v31 << ((4294967295) & 63))
            _ = v31
        } else {
            v1 -= v2
        }
    } else {
        if false {
            for fi32 := 0; fi32 < 19; fi32++ {
                v33 := []int{}
                _ = fi32
                _ = v33
            }
            var v34 bool = false
            _ = v34
        } else {
            v35 := []int{}
            for fi36 := 0; fi36 < 14; fi36++ {
                v4 += v4
                var v37 string = (("hello" + ("hello" + "foo bar")) + ("a" + ""))
                _ = fi36
                _ = v37
            }
            _ = v35
        }
        var v38 float32 = ((v4 * v4) - float32(((0.1 * (-2.25 + -2.25)) / 3.14159)))
        v4 -= (((v4 + v4) / (float32(1e10) + v4)) / (0.1 + 3.14159))
        _ = v38
    }
    v39 := []int{}
    var v40 uint8 = 1
    var v41 int = 127
    v42 := GenS1{F0: -25, F1: (v40 - v40)}
    for fi43 := 0; fi43 < 3; fi43++ {
        var v44 bool = (((32767 > 32767) || false) || true)
        _ = fi43
        _ = v44
    }
    v45 := GenS2{F0: "a"}
    v2 -= -128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v5
}
