package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
func genFn0() int64 {
    var v6 float32 = 1e10
    v6 -= v6
    if (!false) {
        var v7 int = 2147483647
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{}
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v10 := []int{-2147483648, -36, -958}
    var v11 float32 = (0.0 - v6)
    _ = v10
    _ = v11
    _ = v6
    return 42
}

func main() {
    var v1 string = "a"
    v2 := []uint32{0, 16, 42}
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    var v12 string = (("foo bar" + "hello") + (v1 + v1))
    v1 = "foo bar"
    v1 = v12
    v1 = ("foo bar" + v12)
    v13 := GenS1{F0: 100}
    var v14 float64 = 1e10
    v1 = "hello"
    for fi15 := 0; fi15 < 18; fi15++ {
        v12 = v12
        if true {
            if false {
                var v16 string = ((((v12 + v1) + v12) + v1) + ((v1 + v1) + ""))
                for fi17 := 0; fi17 < 19; fi17++ {
                    v1 = v12
                    _ = fi17
                }
                _ = v16
            }
            v12 = v1
            if (!true) {
                if false {
                    var v18 uint16 = 32767
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        var v20 uint8 = 2
                        if ((fi19 | (fi19 - (fi15 + (fi15 / (((-(fi15))) | 1))))) == fi19) {
                            switch v1 {
                            case "a":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "hello":
                                fmt.Println(v1, v1)
                            default:
                                switch -1 {
                                case -9223372036854775808:
                                    if false {
                                        v21 := GenS2{F0: "hello"}
                                        var v22 int64 = 59
                                        v1 = (("hello" + "") + ("" + (((v1 + "") + v1) + ((v1 + "a") + "\n\t"))))
                                        _ = v21
                                        _ = v22
                                    } else {
                                        var v23 uint16 = 32767
                                        v24 := GenS1{F0: -128}
                                        var v25 float64 = v14
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                case 1:
                                    switch fi19 {
                                    case 0:
                                        v12 = v12
                                    case -9223372036854775808:
                                        v14 -= v14
                                    case 85:
                                        v26 := GenS3{F0: 1}
                                        _ = v26
                                    }
                                case -612:
                                    var v27 uint64 = 1
                                    _ = v27
                                default:
                                    var v28 int = fi19
                                    v14 -= 0.0
                                    _ = v28
                                }
                            }
                        } else {
                            var v29 uint32 = 2147483648
                            var v30 float64 = ((v14 - 0.0) / ((-2.25 - v14) / (v14 * v14)))
                            var v31 uint64 = 585
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = fi19
                        _ = v20
                    }
                }
                v12 = (v1 + (v1 + ""))
            }
        } else {
            v32 := GenS2{F0: v12}
            v1 = (v12 + "a")
            _ = v32
        }
        v1 = v12
        _ = fi15
    }
    fmt.Println(v14, v14, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
