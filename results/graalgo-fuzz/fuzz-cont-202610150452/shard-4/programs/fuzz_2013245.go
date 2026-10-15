package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
func genFn0() int16 {
    var v7 uint64 = 1
    if (!(228 <= 1)) {
        for fi8 := 0; fi8 < 18; fi8++ {
            v7 = v7
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 int = fi8
                var v11 float32 = 0.1
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
    }
    v12 := []int{-1}
    var v13 bool = false
    _ = v12
    _ = v13
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = (true && false)
    var v3 bool = (659 != -2147483648)
    var v4 uint32 = 65535
    v5 := make(map[uint32]string)
    v6 := []uint32{}
    var v14 float32 = float32(float64(((0.1 - 1e10) - 0.0)))
    switch "\n\t" {
    case "\n\t":
        var v15 bool = (!v2)
        _ = v15
    case "foo bar":
        var v16 int64 = -9223372036854775808
        _ = v16
    default:
        if true {
            v4 += v4
            v6 = append(v6, (^(v4)))
        }
    }
    if (v3 && true) {
        fmt.Println(v3, v2)
        var v17 string = "hello"
        switch "foo bar" {
        case "\n\t":
            switch (((v17 + "\n\t") + "\n\t") + "") {
            case "foo bar":
                if (-2.25 <= (3.14159 - 3.14159)) {
                    var v18 uint16 = 878
                    if len(v6) > 0 {
                        v6[0] = 2147483647
                    }
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        if true {
                            var v20 uint64 = uint64(v4)
                            _ = v20
                        } else {
                            v21 := GenS1{F0: 181, F1: -291}
                            _ = v21
                        }
                        if (v3 || (!false)) {
                            var v22 int = (^(fi19))
                            v23 := v1[127]
                            _ = v22
                            _ = v23
                        } else {
                            v1[fi19] = ("" + v17)
                        }
                        _ = fi19
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v4 -= (^(v4))
                }
            default:
                v24 := GenS1{F0: 90, F1: -1}
                _ = v24
            }
        }
        _ = v17
    }
    fmt.Println(v2, v2)
    if (1 == 514) {
        delete(v5, (v4 >> ((970) & 63)))
    }
    fmt.Println(v14, v14)
    var v25 float32 = 1.5
    if (1.5 > 3.14159) {
        var v26 string = ""
        _ = v26
    } else {
        var v27 uint = 65535
        v28 := v5[2147483648]
        _ = v27
        _ = v28
    }
    v25 -= v25
    if (v2 && v3) {
        v2 = true
        for fi29 := 0; fi29 < 3; fi29++ {
            var v30 uint = 255
            var v31 uint = (v30 & v30)
            _ = fi29
            _ = v30
            _ = v31
        }
    } else {
        var v32 float32 = v25
        _ = v32
    }
    if ((0.0 * 1.5) != (0.1 * -2.25)) {
        v33 := GenS1{F0: 42, F1: 1}
        v34 := GenS1{F0: 128, F1: -570}
        var v35 string = ("hello" + ("hello" + ""))
        _ = v33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        v14 = v14
        _ = len(v6)
        _ = fi36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
