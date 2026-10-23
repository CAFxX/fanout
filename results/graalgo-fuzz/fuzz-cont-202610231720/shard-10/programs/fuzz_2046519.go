package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
func genFn0() string {
    var v7 uint8 = 0
    switch 32768 {
    case 128:
        if (255 > 18446744073709551615) {
            var v8 string = "a"
            if (2 <= 1) {
                v7 -= (^(v7))
                v9 := GenS1{F0: 1, F1: 100, F2: 2}
                _ = v9
            } else {
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 float64 = (1e10 / 1.5)
                    v7 -= (v7 % ((v7) | 1))
                    _ = fi10
                    _ = v11
                }
            }
            var v12 int32 = 42
            _ = v12
            _ = v8
        }
    case 4294967295:
        switch "" {
        case "a":
            var v13 uint64 = 42
            _ = v13
        case "hello":
            v14 := GenS1{F0: 100, F1: -1, F2: 42}
            _ = v14
        }
    case 0:
        var v15 uint32 = 336
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: 35, F2: 686}
    var v17 uint64 = 0
    _ = v16
    _ = v17
    _ = v7
    return "a"
}

func main() {
    v1 := []string{}
    v2 := make(map[int]int)
    v3 := []string{(((((("" + "") + "a") + ("" + "hello")) + "hello") + "hello") + ""), ((("\n\t" + (("" + "\n\t") + "foo bar")) + ("\n\t" + "foo bar")) + ""), ("foo bar" + "")}
    v4 := []string{((("\n\t" + ("a" + "foo bar")) + "hello") + ("hello" + "foo bar")), (("a" + "") + ("\n\t" + "a")), "\n\t"}
    v5 := []uint32{2147483647}
    var v6 uint16 = 157
    v18 := GenS1{F0: 1, F1: -1, F2: 4294967295}
    v6 = v6
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        v6 -= (v6 << ((818) & 63))
    } else {
        for fi19 := 0; fi19 < 19; fi19++ {
            fmt.Println(fi19)
            for fi20 := 0; fi20 < 15; fi20++ {
                fmt.Println(fi19)
                _ = fi20
            }
            if false {
                for fi21 := 0; fi21 < 12; fi21++ {
                    var v22 float32 = 1.5
                    _ = fi21
                    _ = v22
                }
                v6 += v6
                fmt.Println(v6, fi19)
            }
            _ = fi19
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 int8 = 19
            v6 -= (v6 | 255)
            v6 += 0
            _ = fi23
            _ = v24
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        var v25 uint16 = v6
        var v26 uint32 = 507
        if (false && (128 <= -2147483648)) {
            if (990 < 42) {
                _ = len(v5)
                if true {
                    var v27 int32 = 42
                    fmt.Println(v6, v25, v6)
                    v26 = (v26 | v26)
                    _ = v27
                } else {
                    var v28 float64 = (-2.25 * 1.5)
                    fmt.Println(v26, v25)
                    var v29 uint32 = 932
                    _ = v28
                    _ = v29
                }
            }
        } else {
            v30 := v2[264]
            delete(v2, (-(v30)))
            var v31 uint16 = 255
            _ = v30
            _ = v31
        }
        _ = v25
        _ = v26
    }
    v32 := GenS1{F0: 42, F1: 1, F2: 2}
    var v33 uint8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
