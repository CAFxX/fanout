package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint) string {
    v7 := []int{-688}
    switch 0 {
    case 42:
        p0 = (p0 ^ (p0 << (((^(p0))) & 63)))
    case -1:
        v8 := GenS1{F0: -1, F1: -2147483648}
        _ = v8
    case -528:
        if false {
            if (-2147483648 <= 875) {
                p0 = (p0 + 1)
            } else {
                switch "" {
                case "foo bar":
                    if false {
                        v9 := GenS1{F0: 9223372036854775807, F1: -351}
                        if (!true) {
                            var v10 int8 = -38
                            var v11 string = ("hello" + "hello")
                            _ = v10
                            _ = v11
                        } else {
                            if true {
                                p0 = (^(p0))
                                if false {
                                    switch (("\n\t" + ((("" + "a") + "") + ("hello" + ("hello" + ("a" + "a"))))) + (("\n\t" + "") + "a")) {
                                    case "":
                                        var v12 uint8 = 3
                                        _ = v12
                                    case "a":
                                        var v13 int16 = 1000
                                        _ = v13
                                    case "\n\t":
                                        if len(v7) > 0 {
                                            v7[0] = 127
                                        }
                                    }
                                }
                            } else {
                                for fi14 := 0; fi14 < 6; fi14++ {
                                    var v15 bool = true
                                    var v16 int32 = 42
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                                var v17 int16 = 0
                                _ = v17
                            }
                            switch "foo bar" {
                            case "":
                                var v18 int32 = -750
                                _ = v18
                            default:
                                p0 -= (p0 ^ 255)
                            }
                        }
                        _ = v9
                    } else {
                        v7 = append(v7, 42)
                        fmt.Println(p0)
                    }
                case "":
                    var v19 uint8 = 2
                    _ = v19
                case "a":
                    fmt.Println(p0)
                default:
                    p0 = (^(p0))
                }
            }
            v7 = append(v7, -43)
            p0 -= 658
        } else {
            var v20 uint32 = 185
            for fi21 := 0; fi21 < 2; fi21++ {
                p0 = p0
                _ = fi21
            }
            _ = v20
        }
    }
    v22 := GenS2{F0: false}
    p0 += p0
    for fi23 := 0; fi23 < 1; fi23++ {
        p0 -= (^(p0))
        _ = fi23
    }
    _ = p0
    _ = v22
    _ = v7
    return (("" + "\n\t") + "hello")
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 0
    var v4 float32 = -2.25
    var v5 uint16 = 2
    var v6 float32 = 3.14159
    var v24 int8 = 28
    var v25 int64 = -786
    fmt.Println(v6)
    if true {
        var v26 uint8 = (v3 + v3)
        if false {
            v1["a"] = (true && true)
            v26 += (v3 - v26)
        }
        v27 := GenS1{F0: -46, F1: int32(v24)}
        _ = v26
        _ = v27
    } else {
        fmt.Println(v3, v24)
        delete(v2, 42)
    }
    v28 := []int{32767, 127, -235}
    if (2147483647 > 467) {
        if true {
            fmt.Println(v25, v25, v24)
            var v29 bool = (float32(0.0) != -2.25)
            _ = v29
        } else {
            var v30 int = -62
            v28 = append(v28, -2147483648)
            v28 = append(v28, v30)
            _ = v30
        }
        if false {
            fmt.Println(v4, v5)
            v31 := GenS2{F0: ((false || true) || (v25 <= 0))}
            _ = v31
        }
    }
    v3 = v3
    fmt.Println(v24, v25)
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    if ((0 == 492) && (!(!(!((v3 % ((1) | 1)) > 49))))) {
        switch 1 {
        case 4294967295:
            v6 -= 0.1
        case 9223372036854775807:
            switch 127 {
            case 2147483647:
                var v32 int = 654
                _ = v32
            case -9223372036854775808:
                var v33 float32 = (-2.25 / float32(float64((3.14159 + 0.0))))
                _ = v33
            case 32767:
                var v34 uint = 531
                _ = v34
            default:
                if true {
                    _ = len(v28)
                    v6 = v6
                } else {
                    delete(v2, 42)
                }
            }
        }
    }
    var v35 int16 = -1
    v4 -= v4
    switch "" {
    case "hello":
        var v36 uint64 = 111
        _ = v36
    default:
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 int32 = 2147483647
            _ = fi37
            _ = v38
        }
    }
    var v39 int16 = (-(v35))
    _ = len(v28)
    fmt.Println(v3, v6)
    var v40 uint8 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
