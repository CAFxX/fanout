package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint64; F2 bool }
func genFn0(p0 int16, p1 int) bool {
    if (float32(-2.25) == (float32(((((-2.25 / 1.0) * -2.25) * float64(((-2.25 + 0.1) / float32(1e10)))) - 1.5)) / 1.0)) {
        if ("foo bar" != "hello") {
            switch "a" {
            case "\n\t":
                for fi4 := 0; fi4 < 20; fi4++ {
                    for fi5 := 0; fi5 < 1; fi5++ {
                        var v6 int8 = -1
                        v6 = -1
                        _ = fi5
                        _ = v6
                    }
                    v7 := []int{(fi4 ^ fi4), (-(fi4))}
                    _ = fi4
                    _ = v7
                }
            case "a":
                var v8 uint32 = 2147483648
                _ = v8
            case "foo bar":
                var v9 bool = (!(1 >= 255))
                _ = v9
            }
            fmt.Println(p0)
            var v10 uint16 = 58
            _ = v10
        }
        var v11 int = p1
        _ = v11
    }
    if (0 != 42) {
        v12 := []int{(p1 % (((^(p1))) | 1)), (p1 - 127), p1, p1}
        p1 = p1
        _ = v12
    }
    _ = p0
    _ = p1
    return true
}
func genFn1(p0 uint) uint {
    v13 := []int{int(p0), -1, 32767, -2147483648}
    var v14 uint = p0
    _ = p0
    _ = v13
    _ = v14
    return (p0 / (((p0 * 2)) | 1))
}

func main() {
    var v1 bool = true
    var v2 uint8 = 255
    var v3 int8 = 1
    if false {
        if true {
            var v15 uint32 = 42
            var v16 float64 = (3.14159 / 3.14159)
            if (v1 || false) {
                switch (("foo bar" + "foo bar") + "hello") {
                case "":
                    v17 := GenS1{F0: (("\n\t" + "a") + "")}
                    _ = v17
                default:
                    if v1 {
                        for fi18 := 0; fi18 < 11; fi18++ {
                            fmt.Println(v2)
                            _ = fi18
                        }
                        var v19 uint = 171
                        _ = v19
                    }
                }
            }
            _ = v15
            _ = v16
        } else {
            switch "foo bar" {
            case "\n\t":
                v1 = (1 != 65535)
            }
        }
        var v20 bool = (42 <= 2)
        _ = v20
    }
    fmt.Println(v2, v1, v2)
    v2 -= 74
    v2 = 2
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 float64 = (((1.5 / ((((float64(1.5) - (1e10 - (1.5 + ((0.0 * 0.1) - 3.14159)))) + float64((1.5 + 0.1))) / 1.0) * 1e10)) / 0.1) + float64(1e10))
            _ = fi22
            _ = v23
        }
        var v24 uint16 = 0
        switch (fi21 % ((9223372036854775807) | 1)) {
        case -328:
            v25 := GenS2{F0: 842, F1: 0, F2: ((!v1) || (2 != uint32(v2)))}
            _ = v25
        case -9223372036854775808:
            var v26 int32 = -1
            _ = v26
        case 2147483647:
            v27 := GenS2{F0: 2, F1: uint64(v24), F2: v1}
            _ = v27
        }
        _ = fi21
        _ = v24
    }
    var v28 int16 = 600
    var v29 string = "a"
    v2 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
