package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 string = (("hello" + "a") + "foo bar")
    v2 := []int{}
    v3 := make(map[string]int)
    v2 = append(v2, 2147483647)
    v1 = "\n\t"
    var v4 uint64 = 2
    if true {
        v3[v1] = -317
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 int16 = 202
            _ = fi5
            _ = v6
        }
        var v7 int32 = -2147483648
        _ = v7
    } else {
        for fi8 := 0; fi8 < 7; fi8++ {
            if ((109 >= -56) || (1.5 == ((((0.1 * 0.0) - float64(3.14159)) + 3.14159) / -2.25))) {
                switch v1 {
                case "a":
                    var v9 uint32 = 1
                    _ = v9
                default:
                    v4 += (^(v4))
                }
            }
            var v10 string = ((v1 + v1) + (v1 + v1))
            v3[v1] = fi8
            _ = fi8
            _ = v10
        }
    }
    if true {
        v1 = v1
    }
    var v11 string = ("" + (("a" + ("foo bar" + (v1 + v1))) + ("\n\t" + v1)))
    if true {
        if true {
            _ = len(v3)
        } else {
            for fi12 := 0; fi12 < 7; fi12++ {
                fmt.Println(fi12, v1, fi12)
                var v13 uint64 = v4
                _ = fi12
                _ = v13
            }
            v14 := v3[("\n\t" + ((("" + v11) + "") + v1))]
            if (v4 == v4) {
                fmt.Println(v14, v11)
                v1 = v11
                fmt.Println(v11)
            } else {
                v1 = v1
                if (true || false) {
                    if true {
                        v14 -= v14
                        v15 := GenS1{F0: 820}
                        _ = v15
                    }
                } else {
                    fmt.Println(v1, v4)
                }
                var v16 float32 = float32(0.1)
                _ = v16
            }
            _ = v14
        }
        fmt.Println(v1)
        if (!true) {
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 bool = (false && (-418 < 1))
                for fi19 := 0; fi19 < 4; fi19++ {
                    v11 = ((("foo bar" + (("a" + v1) + v11)) + v1) + ("a" + v11))
                    _ = fi19
                }
                _ = fi17
                _ = v18
            }
            v4 -= 4294967295
            v20 := GenS1{F0: 606}
            _ = v20
        }
    } else {
        if ((v11 + v1) >= v1) {
            v21 := GenS2{F0: v4, F1: v4}
            for fi22 := 0; fi22 < 15; fi22++ {
                switch v11 {
                case "foo bar":
                    for fi23 := 0; fi23 < 3; fi23++ {
                        var v24 int = fi22
                        _ = fi23
                        _ = v24
                    }
                case "":
                    if (!true) {
                        var v25 uint32 = 2
                        if (true && (32767 < 329)) {
                            var v26 float64 = -2.25
                            _ = v26
                        }
                        if false {
                            var v27 uint8 = 250
                            v4 = v4
                            _ = v27
                        }
                        _ = v25
                    }
                default:
                    var v28 int64 = -14
                    _ = v28
                }
                _ = fi22
            }
            _ = v21
        } else {
            var v29 uint = 18446744073709551615
            for fi30 := 0; fi30 < 3; fi30++ {
                if true {
                    v4 = (v4 + v4)
                }
                _ = fi30
            }
            _ = v29
        }
        v1 = v11
    }
    v11 = v1
    switch -1 {
    case 42:
        for fi31 := 0; fi31 < 3; fi31++ {
            fmt.Println(v4)
            v4 = (v4 / (((^(v4))) | 1))
            _ = fi31
        }
    case -430:
        v11 = ("foo bar" + v1)
    case 713:
        var v32 uint32 = 884
        _ = v32
    }
    var v33 uint64 = (v4 >> ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
