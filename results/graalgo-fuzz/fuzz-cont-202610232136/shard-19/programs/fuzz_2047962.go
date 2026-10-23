package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("foo bar" + "hello")
    var v3 bool = false
    v4 := []int{58, -1}
    var v5 int16 = 1
    for fi6 := 0; fi6 < 17; fi6++ {
        delete(v1, 722)
        _ = fi6
    }
    if (!v3) {
        var v7 uint16 = 2
        _ = v7
    } else {
        v3 = (((255 < -1) || true) || (9223372036854775807 <= -1))
        v8 := GenS1{F0: (-107 <= v5)}
        _ = v8
    }
    if v3 {
        v5 = -224
    } else {
        v9 := GenS1{F0: v3}
        switch (v2 + v2) {
        case "foo bar":
            v10 := v1[684]
            _ = v10
        case "a":
            if v3 {
                switch ("" + ("\n\t" + ("hello" + ("foo bar" + ("\n\t" + ("\n\t" + v2)))))) {
                case "foo bar":
                    var v11 float64 = 3.14159
                    _ = v11
                case "":
                    v3 = (v3 && false)
                default:
                    v3 = (((1 != 2147483647) && (272 < 2147483647)) && v3)
                }
                var v12 uint = 134
                _ = v12
            }
        case "hello":
            if (true && false) {
                v13 := GenS1{F0: (1 < -9223372036854775808)}
                if (!(v3 || ((v3 && v3) && (!(868 == 0))))) {
                    v14 := v1[611]
                    v5 -= 42
                    _ = v14
                } else {
                    var v15 float32 = 0.0
                    _ = v15
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v13
            } else {
                if ((363 <= 434) && v3) {
                    v5 = (v5 / ((1000) | 1))
                    v2 = v2
                    v2 = v2
                } else {
                    v16 := v1[2]
                    _ = v16
                }
                v3 = v3
                var v17 string = "foo bar"
                _ = v17
            }
        default:
            if v3 {
                if (303 <= 1) {
                    v2 = v2
                } else {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 int32 = 42
                        v2 = v2
                        _ = fi18
                        _ = v19
                    }
                    _ = len(v4)
                    v3 = true
                }
                var v20 int8 = -107
                switch 127 {
                case -1:
                    fmt.Println(v2, v5, v5)
                default:
                    v21 := GenS1{F0: (1.5 < ((-2.25 + 1.5) - (3.14159 / (-2.25 - 1.5))))}
                    _ = v21
                }
                _ = v20
            }
        }
        fmt.Println(v2)
        _ = v9
    }
    v22 := GenS1{F0: (794 != 2)}
    for fi23 := 0; fi23 < 18; fi23++ {
        if false {
            v5 += 655
        } else {
            v2 = ("\n\t" + (v2 + (v2 + v2)))
            for fi24 := 0; fi24 < 20; fi24++ {
                v3 = v3
                _ = fi24
            }
        }
        _ = fi23
    }
    if v3 {
        var v25 string = "foo bar"
        _ = len(v4)
        _ = v25
    } else {
        fmt.Println(v3)
        v3 = v3
        var v26 string = v2
        _ = v26
    }
    var v27 int8 = -1
    v28 := GenS2{F0: 42}
    var v29 uint8 = 127
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
