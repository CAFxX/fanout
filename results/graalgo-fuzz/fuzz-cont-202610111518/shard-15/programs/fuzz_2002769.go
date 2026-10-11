package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint64 }
func genFn0(p0 uint16) uint32 {
    if false {
        var v4 int16 = 456
        for fi5 := 0; fi5 < 1; fi5++ {
            var v6 bool = (false && (-2.25 >= 1.5))
            var v7 uint8 = 0
            var v8 int = 132
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        _ = v4
    }
    fmt.Println(p0)
    v9 := []int{2147483647, -9223372036854775808}
    p0 = p0
    _ = p0
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 uint32 = 4294967295
    var v10 bool = (0.0 != (3.14159 + 0.0))
    _ = len(v2)
    switch 42 {
    case -9223372036854775808:
        v10 = v10
    case 1:
        if v10 {
            fmt.Println(v10, v10)
            if (!(-745 >= 0)) {
                var v11 int16 = -1
                _ = v11
            }
            if ((2147483647 <= 325) && false) {
                var v12 float64 = 1.5
                _ = v12
            } else {
                v3 -= (v3 >> ((4294967295) & 63))
                v13 := GenS1{F0: (100 > 42), F1: "", F2: 4294967295}
                _ = v13
            }
        }
    default:
        v10 = v10
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v14 := v2[-9223372036854775808]
        _ = v14
    default:
        var v15 string = "hello"
        _ = v15
    }
    var v16 bool = v10
    v10 = (!true)
    var v17 float64 = float64(((0.0 + -2.25) / (0.0 + (1.5 * float32((3.14159 + float64((0.1 * (3.14159 / 1e10)))))))))
    var v18 int64 = 915
    v19 := GenS1{F0: v16, F1: "foo bar", F2: 1}
    if (0 == -1) {
        v16 = (2 == 2)
        var v20 uint16 = 42
        delete(v2, 32767)
        _ = v20
    } else {
        v10 = (true || (-224 < -1))
        if (0 > -936) {
            v10 = false
        }
        var v21 uint8 = 181
        _ = v21
    }
    switch 1 {
    case 32768:
        var v22 uint64 = 166
        _ = v22
    case 127:
        var v23 float64 = (v17 + (v17 + (float64(1e10) * v17)))
        _ = v23
    case -1:
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := GenS1{F0: (42 <= 1), F1: "hello", F2: 2}
            _ = fi24
            _ = v25
        }
    default:
        v26 := GenS1{F0: (32767 == 9223372036854775807), F1: ("" + "\n\t"), F2: 201}
        _ = v26
    }
    var v27 uint32 = v3
    if v10 {
        if v16 {
            var v28 int32 = 0
            _ = v28
        } else {
            switch 42 {
            case 42:
                if (536 != 65535) {
                    for fi29 := 0; fi29 < 12; fi29++ {
                        var v30 int32 = 0
                        _ = fi29
                        _ = v30
                    }
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
}
