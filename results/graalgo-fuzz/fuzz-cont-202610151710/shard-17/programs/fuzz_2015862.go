package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint32, p1 uint16) uint16 {
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        var v11 int = fi8
        _ = fi8
        _ = v11
    }
    var v12 float32 = -2.25
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return (p1 + p1)
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 350
    var v3 uint8 = 25
    v4 := make(map[string]bool)
    var v5 uint8 = (v3 + v3)
    var v6 uint32 = 2147483648
    v13 := GenS1{F0: "hello"}
    if ((-2.25 / 3.14159) == (-2.25 / 0.1)) {
        v6 += v6
        v14 := []int{}
        if len(v14) > 0 {
            v14[0] = 128
        }
        _ = v14
    }
    if (false || true) {
        fmt.Println(v5)
        if true {
            if (42 < 0) {
                v5 += v3
                fmt.Println(v6, v3, v6)
            } else {
                v15 := []int{479}
                _ = v15
            }
            fmt.Println(v6)
            v3 -= v5
        } else {
            for fi16 := 0; fi16 < 8; fi16++ {
                if true {
                    if ((false && (false || true)) || false) {
                        var v17 uint8 = 209
                        if false {
                            var v18 int = fi16
                            if false {
                                var v19 int = 4294967295
                                fmt.Println(v17, v18, v6)
                                _ = v19
                            }
                            v20 := v1[("a" + ("foo bar" + "hello"))]
                            _ = v18
                            _ = v20
                        }
                        _ = v17
                    }
                }
                v5 = 128
                v2 = (v2 - 361)
                _ = fi16
            }
            v21 := GenS1{F0: ("hello" + "a")}
            _ = v21
        }
    } else {
        v3 = v5
        v5 -= (v3 % ((1) | 1))
    }
    var v22 uint = 42
    var v23 uint8 = (v5 | v3)
    if (!(false && false)) {
        var v24 uint64 = (v2 << (((v22 << (((v22 & v22)) & 63))) & 63))
        _ = v24
    } else {
        fmt.Println(v6, v6, v2)
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 string = ("a" + "a")
            _ = fi25
            _ = v26
        }
        delete(v1, (("\n\t" + "\n\t") + ""))
    }
    v27 := GenS1{F0: "a"}
    v6 = 2
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := GenS1{F0: (("" + "") + ("a" + "a"))}
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
