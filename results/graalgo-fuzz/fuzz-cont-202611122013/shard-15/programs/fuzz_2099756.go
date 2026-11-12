package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32; F1 bool; F2 string }
type GenS3 struct { F0 int32; F1 string }
func genFn0(p0 uint, p1 int8, p2 int64) int64 {
    switch -871 {
    case 4294967295:
        var v7 int16 = 1000
        _ = v7
    default:
        for fi8 := 0; fi8 < 15; fi8++ {
            if false {
                p2 += -477
            }
            var v9 uint16 = 32767
            _ = fi8
            _ = v9
        }
    }
    var v10 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (p2 | 550)
}
func genFn1(p0 string) bool {
    var v11 uint16 = 42
    v11 += v11
    _ = p0
    _ = v11
    return false
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 1
    var v3 float32 = (3.14159 + 1e10)
    v4 := []string{("\n\t" + ""), "hello"}
    var v5 uint16 = 42
    var v6 int = 32767
    fmt.Println(v2, v6, v5)
    fmt.Println(v2, v2)
    if (false || ("a" == (("foo bar" + "a") + "foo bar"))) {
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := GenS1{F0: (("\n\t" + "hello") + ("a" + "a")), F1: 680}
            var v14 int = (fi12 + (v6 + v6))
            var v15 float64 = 0.1
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    } else {
        v2 = v2
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi16
    }
    switch "a" {
    case "foo bar":
        for fi17 := 0; fi17 < 12; fi17++ {
            switch fi17 {
            case 2147483647:
                var v18 uint16 = (v5 / ((v5) | 1))
                _ = v18
            default:
                for fi19 := 0; fi19 < 14; fi19++ {
                    fmt.Println(fi19, v3, fi17)
                    _ = fi19
                }
            }
            switch "\n\t" {
            case "":
                v6 += (v6 << ((64) & 63))
            case "foo bar":
                v3 = 3.14159
            case "hello":
                switch (fi17 | (^(fi17))) {
                case 97:
                    if (29 != 2147483648) {
                        var v20 int = v6
                        _ = v20
                    } else {
                        switch 1 {
                        case 347:
                            fmt.Println(v5, v6)
                        }
                        v3 -= (0.1 / float32((1.5 / (0.0 + ((0.1 / 1e10) * 3.14159)))))
                    }
                case -456:
                    v3 = -2.25
                case 42:
                    for fi21 := 0; fi21 < 5; fi21++ {
                        v22 := GenS1{F0: "", F1: uint(v5)}
                        if true {
                            v23 := GenS1{F0: "foo bar", F1: 573}
                            _ = v23
                        }
                        _ = fi21
                        _ = v22
                    }
                default:
                    v24 := GenS1{F0: "a", F1: 0}
                    _ = v24
                }
            }
            _ = fi17
        }
    default:
        var v25 int8 = -116
        _ = v25
    }
    if true {
        var v26 int = v6
        for fi27 := 0; fi27 < 4; fi27++ {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            switch (("" + "a") + "") {
            case "":
                for fi28 := 0; fi28 < 5; fi28++ {
                    v29 := GenS2{F0: 2147483647, F1: true, F2: ("a" + "\n\t")}
                    _ = fi28
                    _ = v29
                }
            }
            if (false && (!(!(v26 >= (v26 - v6))))) {
                v26 = (v26 * (v6 ^ (v26 / (((v26 & 192)) | 1))))
            } else {
                v5 += v5
                v5 -= (v5 - (v5 | (v5 % ((v5) | 1))))
                for fi30 := 0; fi30 < 11; fi30++ {
                    var v31 int8 = -128
                    var v32 int64 = 0
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
            _ = fi27
        }
        v2 -= (-(v2))
        _ = v26
    } else {
        var v33 uint16 = (v5 + v5)
        v33 = 42
        v34 := GenS3{F0: 806, F1: ("" + ("\n\t" + (("\n\t" + ("foo bar" + "a")) + "")))}
        _ = v33
        _ = v34
    }
    fmt.Println(v2)
    v6 += 32768
    var v35 uint32 = 42
    var v36 float64 = 0.1
    v1[v6] = false
    var v37 int = 364
    v37 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
