package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 string }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 string }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 17; fi4++ {
        switch "hello" {
        case "\n\t":
            for fi5 := 0; fi5 < 8; fi5++ {
                if true {
                    var v6 bool = (false && false)
                    v7 := []int{(fi4 + fi4)}
                    switch "" {
                    case "\n\t":
                        switch "hello" {
                        case "":
                            if (2147483647 < 42) {
                                v6 = ((!(65535 > 0)) && (0 < 42))
                                v8 := GenS1{F0: 128, F1: 127, F2: (("foo bar" + (("" + "foo bar") + "foo bar")) + "")}
                                _ = v8
                            }
                        case "hello":
                            if true {
                                if len(v7) > 0 {
                                    v7[0] = fi5
                                }
                                var v9 int16 = 42
                                v9 = (v9 ^ -1)
                                v9 = (v9 / (((^(v9))) | 1))
                                _ = v9
                            } else {
                                var v10 float64 = float64(-2.25)
                                switch fi5 {
                                case 32768:
                                    if len(v7) > 0 {
                                        v7[0] = 4294967295
                                    }
                                case 4294967295:
                                    switch 1 {
                                    case 714:
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v11 int32 = 2147483647
                                            _ = v11
                                        case "foo bar":
                                            v6 = v6
                                        case "":
                                            v7 = append(v7, 128)
                                        }
                                    case -31:
                                        v6 = false
                                    default:
                                        switch (("" + "") + "foo bar") {
                                        case "hello":
                                            v6 = v6
                                        case "":
                                            v10 += ((-2.25 + v10) - v10)
                                        default:
                                            var v12 uint32 = 42
                                            _ = v12
                                        }
                                    }
                                case 20:
                                    _ = len(v7)
                                default:
                                    if (!v6) {
                                        v13 := GenS1{F0: 128, F1: int8(fi5), F2: ((((((("" + "a") + ("" + "")) + (("a" + ("hello" + "hello")) + "foo bar")) + ("\n\t" + "\n\t")) + "hello") + "foo bar") + "hello")}
                                        _ = v13
                                    }
                                }
                                v14 := GenS2{F0: ("\n\t" + "hello"), F1: 1}
                                _ = v10
                                _ = v14
                            }
                        default:
                            v6 = (1 != 2)
                            var v15 float64 = 0.0
                            var v16 int8 = -1
                            _ = v15
                            _ = v16
                        }
                    case "hello":
                        _ = len(v7)
                    case "a":
                        v6 = (1 != 2)
                    }
                    _ = v6
                    _ = v7
                } else {
                    var v17 uint8 = 2
                    switch "hello" {
                    case "":
                        v18 := []int{fi5}
                        _ = v18
                    }
                    _ = v17
                }
                _ = fi5
            }
        default:
            v19 := []int{(fi4 / ((fi4) | 1)), fi4, fi4, (fi4 ^ (fi4 % (((fi4 << ((42) & 63))) | 1)))}
            _ = v19
        }
        _ = fi4
    }
    v20 := GenS2{F0: ("foo bar" + ((("foo bar" + "") + ("foo bar" + ("a" + "\n\t"))) + "hello")), F1: 628}
    _ = v20
    return 68
}

func main() {
    var v1 uint16 = 42
    var v2 int64 = -619
    var v3 bool = false
    for fi21 := 0; fi21 < 10; fi21++ {
        v3 = (!(!(!(true || v3))))
        if v3 {
            fmt.Println(v2)
            if (((!true) || v3) || v3) {
                if (1000 >= -583) {
                    v22 := GenS2{F0: ("hello" + "\n\t"), F1: v1}
                    v2 -= (v2 ^ v2)
                    _ = v22
                } else {
                    v23 := []int{}
                    _ = v23
                }
                switch "\n\t" {
                case "":
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v2 -= v2
                        _ = fi24
                    }
                default:
                    v2 -= (v2 * (^(v2)))
                }
                v3 = (!v3)
            }
            v3 = (-333 > v2)
        }
        for fi25 := 0; fi25 < 6; fi25++ {
            v1 -= (v1 / ((v1) | 1))
            var v26 uint = 664
            _ = fi25
            _ = v26
        }
        _ = fi21
    }
    v27 := []int{}
    var v28 uint16 = (^(v1))
    v28 = (v1 * (v28 >> ((0) & 63)))
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    v3 = ((!(!(1e10 == 0.0))) || v3)
    v29 := GenS3{F0: "a"}
    var v30 int16 = 1
    v31 := GenS2{F0: "hello", F1: v28}
    fmt.Println(v3)
    v28 = 65535
    v30 += 32767
    if (genFn0() == 222) {
        for fi32 := 0; fi32 < 1; fi32++ {
            var v33 uint = uint(v28)
            _ = fi32
            _ = v33
        }
        fmt.Println(v2, v1, v3)
    }
    switch 0 {
    case 127:
        v2 += (v2 / (((v2 * v2)) | 1))
    case 1:
        if v3 {
            v3 = (!v3)
        } else {
            v30 += 561
        }
    default:
        v30 -= 1000
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
