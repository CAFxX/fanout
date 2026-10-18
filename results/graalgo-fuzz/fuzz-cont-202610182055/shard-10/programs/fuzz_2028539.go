package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint16 }
func genFn0(p0 int, p1 uint) uint32 {
    v6 := []int{32768, p0, (^(p0)), p0}
    p0 += p0
    var v7 uint16 = 583
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 136
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int = 127
    var v3 float64 = (-2.25 * (0.1 - ((3.14159 * -2.25) - (float64(3.14159) / 1e10))))
    v4 := []string{"a", "hello"}
    v5 := make(map[string]string)
    if (43 != 101) {
        _ = len(v5)
    }
    var v8 float32 = 0.0
    for fi9 := 0; fi9 < 17; fi9++ {
        v4 = append(v4, (("a" + "") + "foo bar"))
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v8)
        default:
            var v10 uint32 = 65535
            _ = v10
        }
        _ = fi9
    }
    v11 := GenS1{F0: 42}
    if (1000 > -32768) {
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := GenS1{F0: 255}
            for fi14 := 0; fi14 < 19; fi14++ {
                v3 -= v3
                var v15 uint8 = 2
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    }
    v3 = (-2.25 - (v3 - ((v3 + v3) + 0.0)))
    var v16 int64 = 42
    v17 := v5["\n\t"]
    v8 -= (3.14159 / v8)
    switch v2 {
    case 127:
        if ((v2 << ((1) & 63)) != (v2 - v2)) {
            if false {
                if len(v4) > 0 {
                    v4[0] = v17
                }
                var v18 bool = false
                v16 = (v1 >> ((0) & 63))
                _ = v18
            } else {
                fmt.Println(v2)
                switch (v2 << ((42) & 63)) {
                case 128:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    v2 = v2
                }
                v19 := GenS2{F0: 252, F1: v2, F2: 195}
                _ = v19
            }
            switch ("" + "") {
            case "\n\t":
                v20 := GenS1{F0: 0}
                _ = v20
            case "a":
                v21 := GenS3{F0: 343}
                _ = v21
            }
        } else {
            fmt.Println(v8, v8, v16)
            var v22 int32 = 282
            _ = v22
        }
    case 169:
        if (v16 >= (v1 / ((-9223372036854775808) | 1))) {
            v23 := GenS1{F0: 55}
            if true {
                fmt.Println(v17)
                if true {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        for fi25 := 0; fi25 < 19; fi25++ {
                            v26 := GenS4{F0: (v17 + v17)}
                            _ = fi25
                            _ = v26
                        }
                        _ = fi24
                    }
                    v27 := v5[((v17 + v17) + v17)]
                    for fi28 := 0; fi28 < 9; fi28++ {
                        v29 := GenS2{F0: 32, F1: -9223372036854775808, F2: 42}
                        v5[(v17 + ("" + "hello"))] = ""
                        v30 := GenS5{F0: v2, F1: 65535}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = v27
                }
                v8 += (-2.25 + (v8 - 3.14159))
            }
            _ = v23
        }
    case -1:
        if (((v8 <= ((v8 / (1e10 - v8)) * v8)) && true) || ((0 < 127) || ((true && false) || false))) {
            var v31 uint16 = 972
            _ = v31
        }
    }
    var v32 bool = (v16 != v1)
    if v32 {
        v1 = v16
    } else {
        var v33 int64 = v16
        _ = v33
    }
    v3 -= 0.1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if v32 {
        switch (v2 >> ((4294967295) & 63)) {
        case -2147483648:
            var v34 int8 = 100
            _ = v34
        }
    }
    var v35 bool = v32
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v8
}
