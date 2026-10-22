package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 string; F2 bool }
type GenS3 struct { F0 uint; F1 bool; F2 string }
func genFn0(p0 uint, p1 uint8) uint {
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    p1 = 1
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 uint16 = 2
    var v2 uint16 = 132
    var v3 int = 616
    var v4 uint = 805
    v5 := make(map[int]string)
    var v9 int16 = 42
    switch "a" {
    case "foo bar":
        v1 = (v2 * (v1 - v1))
    case "hello":
        for fi10 := 0; fi10 < 18; fi10++ {
            fmt.Println(v2, v1)
            v11 := []int{int(v9), (fi10 % ((4294967295) | 1))}
            if (((true && (!(!true))) || (!((int32(v4) == 431) || false))) && (false && (!false))) {
                switch 906 {
                case 42:
                    if ((-(fi10)) <= (-(v3))) {
                        if (127 == 40) {
                            if true {
                                v12 := GenS1{F0: true}
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                _ = v12
                            } else {
                                if (-2147483648 <= v3) {
                                    v3 += 32767
                                }
                                var v13 int64 = -166
                                fmt.Println(v3)
                                _ = v13
                            }
                            var v14 int16 = (v9 & v9)
                            if (false || ((v2 != 274) && (v9 > v9))) {
                                switch "" {
                                case "a":
                                    var v15 float32 = -2.25
                                    _ = v15
                                case "foo bar":
                                    if (127 != 127) {
                                        switch "hello" {
                                        case "\n\t":
                                            if len(v11) > 0 {
                                                v11[0] = (^(fi10))
                                            }
                                        case "hello":
                                            v16 := GenS1{F0: false}
                                            _ = v16
                                        default:
                                            v2 = (v2 & 65535)
                                        }
                                        v14 = 0
                                        switch v3 {
                                        case 32767:
                                            for fi17 := 0; fi17 < 10; fi17++ {
                                                v18 := GenS2{F0: 42, F1: "\n\t", F2: (!(495 == 1))}
                                                var v19 string = "\n\t"
                                                _ = len(v5)
                                                var v20 bool = true
                                                v20 = ((((v19 + v19) + "foo bar") < ("hello" + (v19 + ("\n\t" + (v19 + "hello"))))) || (1.5 > (0.0 + 1e10)))
                                                _ = fi17
                                                _ = v18
                                                _ = v19
                                                _ = v20
                                            }
                                        case 4294967295:
                                            for fi21 := 0; fi21 < 1; fi21++ {
                                                var v22 string = ((("\n\t" + "\n\t") + "foo bar") + ("hello" + "\n\t"))
                                                v23 := GenS1{F0: false}
                                                _ = fi21
                                                _ = v22
                                                _ = v23
                                            }
                                        case -633:
                                            var v24 float64 = 0.1
                                            _ = v24
                                        }
                                    }
                                case "hello":
                                    v5[(fi10 - fi10)] = "hello"
                                default:
                                    _ = len(v5)
                                }
                                v9 = (v9 & v14)
                                var v25 int32 = int32(v4)
                                _ = v25
                            }
                            _ = v14
                        } else {
                            if ((false && false) || false) {
                                var v26 float64 = (0.0 + 0.0)
                                var v27 int64 = 42
                                _ = v26
                                _ = v27
                            }
                            if len(v11) > 0 {
                                v11[0] = (fi10 & 32767)
                            }
                            v28 := GenS3{F0: genFn0(2, 127), F1: (false || (2147483647 < 4294967295)), F2: "hello"}
                            _ = v28
                        }
                        var v29 int8 = 100
                        _ = v29
                    } else {
                        fmt.Println(v9)
                        var v30 int8 = 127
                        _ = v30
                    }
                case -823:
                    if len(v11) > 0 {
                        v11[0] = v3
                    }
                default:
                    v9 = v9
                }
                if (false && (("a" + ("a" + "\n\t")) <= "\n\t")) {
                    var v31 string = "a"
                    _ = v31
                }
            } else {
                v32 := GenS1{F0: (!false)}
                _ = v32
            }
            _ = fi10
            _ = v11
        }
    default:
        var v33 uint16 = (v2 ^ v2)
        _ = v33
    }
    v34 := []int{4294967295, v3, -2147483648}
    v4 += (v4 | (v4 & (v4 * (v4 ^ v4))))
    fmt.Println(v1, v3)
    v9 = (v9 ^ (v9 ^ v9))
    v1 += (v2 & (v2 << ((v4) & 63)))
    v35 := GenS3{F0: genFn0((v4 / ((genFn0(genFn0((v4 + 257), 127), 199)) | 1)), 0), F1: (4294967295 < 0), F2: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
