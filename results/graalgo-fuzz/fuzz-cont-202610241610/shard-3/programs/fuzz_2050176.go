package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint8, p1 int64) int {
    fmt.Println(p1)
    p1 += 9223372036854775807
    var v6 uint16 = 32767
    var v7 float64 = 0.1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -310
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint32 = (v1 / (((v1 + (v1 % ((v1) | 1)))) | 1))
    var v3 uint16 = 1
    var v4 int8 = -1
    v5 := make(map[uint32]string)
    if (-2147483648 != -1) {
        delete(v5, v2)
        v3 = uint16(v2)
        v5[v2] = "hello"
    }
    if false {
        v2 -= v2
        v3 += (v3 | (v3 | v3))
    } else {
        v8 := GenS1{F0: (-(v4)), F1: uint(v1), F2: true}
        _ = v8
    }
    var v9 int32 = 1
    if ((((((3.14159 - -2.25) * (float32(0.1) - 1.5)) / 1.0) + -2.25) + (-2.25 - (1e10 / 1.0))) > (((-2.25 * ((3.14159 - -2.25) + -2.25)) * (0.0 - float32((((-2.25 - 3.14159) - -2.25) + 1e10)))) * (0.1 - 1e10))) {
        if (false || (!true)) {
            switch (("\n\t" + ("" + ("" + "a"))) + ("foo bar" + "\n\t")) {
            case "":
                v9 = (v9 | v9)
            case "foo bar":
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 uint8 = 128
                    _ = fi10
                    _ = v11
                }
            case "a":
                v2 = (^(v2))
            default:
                delete(v5, 2147483647)
            }
            v12 := v5[v2]
            _ = v12
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := []int{(fi13 >> ((270) & 63)), int(v4)}
            var v15 string = ""
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    delete(v5, 0)
    v16 := v5[v1]
    v9 += v9
    var v17 string = v16
    if (42 == 4294967295) {
        v2 = v1
        var v18 float64 = float64(3.14159)
        var v19 uint = 1
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 float64 = 3.14159
            v1 += 2147483648
            if true {
                for fi22 := 0; fi22 < 18; fi22++ {
                    fmt.Println(v9, v1)
                    _ = fi22
                }
                var v23 float64 = 1e10
                _ = v23
            } else {
                v24 := GenS2{F0: (v9 | v9)}
                _ = v24
            }
            _ = fi20
            _ = v21
        }
        if (2 > 0) {
            v25 := GenS2{F0: (v9 ^ 689)}
            v26 := v5[(v1 * 2147483647)]
            v27 := []int{}
            _ = v25
            _ = v26
            _ = v27
        }
    }
    var v28 int8 = (v4 | (-(v4)))
    v9 = (v9 - v9)
    var v29 uint16 = (v3 ^ (v3 + v3))
    v30 := GenS2{F0: (^(v9))}
    var v31 bool = (false || true)
    if v31 {
        switch v17 {
        case "":
            v32 := v5[v1]
            _ = v32
        case "foo bar":
            var v33 uint32 = 2147483647
            _ = v33
        }
    } else {
        for fi34 := 0; fi34 < 17; fi34++ {
            var v35 uint64 = 423
            _ = fi34
            _ = v35
        }
    }
    v36 := GenS2{F0: int32(v1)}
    var v37 float64 = (1e10 / ((1e10 + ((1e10 - 0.0) + float64(-2.25))) / (-2.25 - (1.5 * (3.14159 / 3.14159)))))
    for fi38 := 0; fi38 < 7; fi38++ {
        if v31 {
            _ = len(v5)
        }
        var v39 string = "hello"
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
