package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint64) uint {
    var v6 int32 = 2147483647
    v6 -= v6
    var v7 uint16 = 2
    v8 := []int{128, 128}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 656
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 0
    v3 := []int{32768, 1}
    v4 := make(map[uint32]int)
    var v5 uint8 = 42
    if len(v3) > 0 {
        v3[0] = 1
    }
    switch 4294967295 {
    case -874:
        _ = len(v4)
    case 209:
        if true {
            fmt.Println(v2, v5)
        } else {
            v9 := GenS1{F0: ("a" + ("\n\t" + "hello"))}
            v2 -= (^(v2))
            _ = v9
        }
    }
    v2 += v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 -= 255
    fmt.Println(v5, v2)
    var v10 int = 2147483647
    fmt.Println(v10)
    var v11 uint8 = (v5 / ((95) | 1))
    v10 = v10
    v11 = v5
    if (42 > 32767) {
        var v12 uint8 = (v5 % (((^(v5))) | 1))
        _ = v12
    } else {
        var v13 string = "\n\t"
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(v5, v5)
            for fi15 := 0; fi15 < 16; fi15++ {
                if (!true) {
                    switch (-(v10)) {
                    case 4294967295:
                        var v16 uint64 = 2
                        _ = v16
                    default:
                        fmt.Println(v13, v5)
                    }
                } else {
                    v2 += (^(v2))
                    if true {
                        delete(v4, 2)
                        v13 = "a"
                        v13 = (("\n\t" + "\n\t") + v13)
                    }
                    v13 = v13
                    _ = len(v1)
                }
                if ((v5 / (((v11 & uint8(v2))) | 1)) != v5) {
                    var v17 uint16 = v2
                    var v18 int = (-(fi14))
                    _ = v17
                    _ = v18
                } else {
                    v11 += (v5 ^ (v5 % (((v5 ^ v11)) | 1)))
                    var v19 uint16 = v2
                    var v20 float64 = (1e10 / 1.5)
                    _ = v19
                    _ = v20
                }
                _ = fi15
            }
            _ = fi14
        }
        _ = v13
    }
    var v21 uint = 267
    v21 -= v21
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int8 = 42
        var v24 uint8 = (v5 * (v5 << ((v21) & 63)))
        _ = fi22
        _ = v23
        _ = v24
    }
    if false {
        v11 += 127
    } else {
        for fi25 := 0; fi25 < 11; fi25++ {
            v5 = v11
            var v26 uint = (^(v21))
            var v27 uint32 = 65535
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    switch "\n\t" {
    case "a":
        switch (("" + "foo bar") + ("hello" + "")) {
        case "foo bar":
            fmt.Println(v21, v5)
        default:
            if false {
                v21 = genFn0(uint64(v2))
            }
        }
    default:
        v11 = v5
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
