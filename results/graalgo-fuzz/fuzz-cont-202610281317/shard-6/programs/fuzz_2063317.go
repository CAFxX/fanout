package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 51
    v3 := make(map[int]string)
    var v4 uint32 = 4294967295
    var v5 int = -309
    v6 := GenS1{F0: ("foo bar" + "\n\t"), F1: -1}
    v7 := GenS1{F0: "foo bar", F1: -9223372036854775808}
    var v8 int64 = -349
    if (true || (false || true)) {
        var v9 bool = ((v8 << ((905) & 63)) == (v8 / ((v8) | 1)))
        _ = v9
    } else {
        switch "" {
        case "":
            if true {
                v4 += 4294967295
                var v10 int32 = 42
                _ = v10
            } else {
                var v11 int = int(v4)
                switch int(v2) {
                case 226:
                    v8 += (v8 + -1)
                default:
                    if (!((!true) || true)) {
                        v4 += (v4 * (v4 ^ v4))
                        v2 = (v2 * v2)
                    }
                }
                v4 = 940
                _ = v11
            }
        case "foo bar":
            var v12 uint8 = 0
            _ = v12
        default:
            fmt.Println(v5)
        }
    }
    fmt.Println(v2)
    fmt.Println(v2)
    switch (("foo bar" + "a") + "\n\t") {
    case "":
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int = (v5 ^ (fi13 - v5))
            v8 -= (v8 ^ (v8 | v8))
            if false {
                if false {
                    if (!true) {
                        _ = len(v3)
                    } else {
                        if true {
                            var v15 int64 = 1
                            _ = v15
                        } else {
                            v16 := GenS1{F0: ("hello" + "hello"), F1: (v8 << ((643) & 63))}
                            _ = v16
                        }
                    }
                    switch v5 {
                    case 127:
                        var v17 int64 = int64(v5)
                        _ = v17
                    case 32767:
                        if false {
                            var v18 uint16 = 17
                            _ = v18
                        } else {
                            fmt.Println(v8)
                        }
                    case 32768:
                        for fi19 := 0; fi19 < 4; fi19++ {
                            v1[-2147483648] = false
                            if (!true) {
                                v5 -= 218
                                var v20 uint = 4294967295
                                _ = v20
                            }
                            switch fi19 {
                            case -666:
                                var v21 int8 = (v2 % ((100) | 1))
                                _ = v21
                            case 4294967295:
                                var v22 int16 = 42
                                _ = v22
                            }
                            fmt.Println(v4, v2)
                            if (65535 > 1) {
                                v23 := GenS1{F0: "a", F1: int64(v4)}
                                _ = v23
                            } else {
                                _ = len(v1)
                            }
                            fmt.Println(v2, fi13)
                            v24 := []int{fi19, fi19}
                            _ = fi19
                            _ = v24
                        }
                    default:
                        if (127 > (-(v14))) {
                            var v25 float64 = 1.5
                            _ = v25
                        }
                    }
                    if false {
                        v2 = (-(v2))
                        var v26 uint8 = 251
                        _ = v26
                    }
                } else {
                    if (!((true && false) && false)) {
                        var v27 bool = (!false)
                        _ = v27
                    }
                }
                v28 := []int{}
                var v29 uint = 18446744073709551615
                _ = v28
                _ = v29
            }
            _ = fi13
            _ = v14
        }
    case "hello":
        var v30 float32 = 3.14159
        _ = v30
    case "\n\t":
        v31 := []int{}
        _ = v31
    }
    if false {
        var v32 string = "foo bar"
        _ = v32
    }
    v33 := GenS2{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
