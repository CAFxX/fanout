package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int; F2 uint32 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int32 }

func main() {
    var v1 int16 = 32767
    var v2 uint = 0
    var v3 float64 = 3.14159
    var v4 int16 = (v1 / ((v1) | 1))
    v5 := make(map[int]uint32)
    v6 := v5[1]
    if (18446744073709551615 != 2) {
        v7 := v5[32767]
        _ = v7
    }
    var v8 float32 = 0.0
    v1 = -584
    v5[32768] = v6
    fmt.Println(v3)
    var v9 int16 = v1
    for fi10 := 0; fi10 < 12; fi10++ {
        _ = len(v5)
        if false {
            var v11 int = (fi10 * fi10)
            if true {
                var v12 string = "hello"
                _ = v12
            }
            _ = v11
        } else {
            fmt.Println(fi10, v8, v8)
            var v13 string = "hello"
            if true {
                v14 := GenS1{F0: 1}
                v15 := GenS2{F0: (fi10 > (fi10 * (fi10 | fi10))), F1: int(v6), F2: v6}
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 3; fi16++ {
                    v3 = 0.0
                    var v17 int16 = (v9 * v4)
                    _ = fi16
                    _ = v17
                }
                v5[128] = v6
            }
            _ = v13
        }
        _ = fi10
    }
    var v18 int = 32768
    var v19 uint = (v2 << ((v2) & 63))
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 int = v18
        _ = fi20
        _ = v21
    }
    v22 := GenS3{F0: "\n\t", F1: "\n\t"}
    v3 = (((v3 / ((float64(1e10) + v3) / (-2.25 - v3))) * (1.5 + ((float64(float32((1.5 + ((0.1 * (3.14159 / 1.5)) / ((1e10 - v3) / 1.5))))) * v3) - -2.25))) - 0.1)
    _ = len(v5)
    for fi23 := 0; fi23 < 14; fi23++ {
        if (135 > 71) {
            switch "" {
            case "a":
                if true {
                    v3 -= v3
                    var v24 string = "hello"
                    _ = v24
                }
            }
            delete(v5, (fi23 % ((v18) | 1)))
            if false {
                v5[(^(fi23))] = v6
                v9 = v1
            }
        } else {
            var v25 uint8 = 133
            _ = v25
        }
        _ = fi23
    }
    switch (((("foo bar" + "") + "\n\t") + "\n\t") + "\n\t") {
    case "foo bar":
        if ((^(v6)) > 516) {
            v8 = v8
        }
    default:
        for fi26 := 0; fi26 < 17; fi26++ {
            if true {
                var v27 int = (v18 & fi26)
                v28 := GenS4{F0: (v6 << (((v2 & v2)) & 63))}
                switch 42 {
                case -206:
                    switch "a" {
                    case "hello":
                        v29 := []int{(fi26 ^ fi26), (v18 / ((fi26) | 1))}
                        _ = v29
                    case "a":
                        for fi30 := 0; fi30 < 3; fi30++ {
                            if false {
                                var v31 uint32 = 0
                                var v32 string = "\n\t"
                                var v33 uint8 = 1
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                v34 := []int{v18, (v18 ^ (fi30 | 1)), (-(fi30)), fi30}
                                v35 := GenS5{F0: (v27 & 726), F1: 172}
                                _ = v34
                                _ = v35
                            }
                            switch v18 {
                            case -813:
                                var v36 float32 = (v8 - (1e10 / (((v8 + 0.0) * v8) - v8)))
                                _ = v36
                            case 128:
                                var v37 uint32 = 2147483648
                                _ = v37
                            case 9223372036854775807:
                                _ = len(v5)
                            default:
                                if (true && true) {
                                    var v38 int32 = 2147483647
                                    if (1 == -1) {
                                        var v39 uint8 = 255
                                        delete(v5, (fi26 - (^(fi30))))
                                        v27 = -2147483648
                                        v5[int(v39)] = (v6 & v6)
                                        v1 += (^(v4))
                                        v40 := GenS6{F0: (v38 + (^(v38)))}
                                        var v41 bool = true
                                        var v42 int32 = v38
                                        var v43 int = (fi26 * (v18 & (v27 & (^(v27)))))
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    }
                                    var v44 uint8 = 2
                                    _ = v38
                                    _ = v44
                                }
                            }
                            fmt.Println(v6, v18, v2)
                            fmt.Println(v4, v27)
                            switch fi26 {
                            case -618:
                                v4 += 32767
                            }
                            _ = fi30
                        }
                    default:
                        v8 = float32(3.14159)
                    }
                default:
                    v27 -= v27
                }
                _ = v27
                _ = v28
            }
            v5[v18] = v6
            _ = fi26
        }
    }
    var v45 uint32 = v6
    var v46 string = ("\n\t" + (("a" + "hello") + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
