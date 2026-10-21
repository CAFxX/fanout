package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 string = ""
    var v4 uint16 = 42
    if false {
        v5 := GenS1{F0: -1, F1: 42, F2: 200}
        if true {
            if (0 != -1) {
                v4 += v4
            }
            for fi6 := 0; fi6 < 18; fi6++ {
                for fi7 := 0; fi7 < 19; fi7++ {
                    if (42 != -32768) {
                        var v8 int16 = 42
                        _ = v8
                    }
                    _ = fi7
                }
                if (2147483647 >= 4294967295) {
                    if false {
                        if (!true) {
                            for fi9 := 0; fi9 < 3; fi9++ {
                                switch fi6 {
                                case 1:
                                    var v10 uint32 = 29
                                    _ = v10
                                default:
                                    v11 := GenS2{F0: 2}
                                    _ = v11
                                }
                                _ = fi9
                            }
                        }
                    }
                    v1[v3] = 45
                }
                _ = fi6
            }
        } else {
            v1["foo bar"] = 0
            v12 := GenS1{F0: -32768, F1: 177, F2: 106}
            _ = v12
        }
        _ = v5
    }
    _ = len(v1)
    var v13 int8 = -128
    v1[v3] = 2147483648
    if (false || (-343 > -320)) {
        if false {
            fmt.Println(v13, v3, v13)
        } else {
            if (true && (!false)) {
                v14 := []int{}
                fmt.Println(v4, v3, v4)
                v15 := GenS3{F0: 76, F1: 1}
                _ = v14
                _ = v15
            } else {
                var v16 int = 0
                v17 := GenS1{F0: -660, F1: 128, F2: 255}
                _ = v16
                _ = v17
            }
            fmt.Println(v3)
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 uint = 934
                _ = fi18
                _ = v19
            }
        }
    }
    var v20 uint32 = 951
    v13 = -18
    var v21 string = ((((v3 + v3) + ((v3 + (("" + "hello") + "")) + (v3 + ("" + v3)))) + (((((v3 + v3) + "hello") + "") + "\n\t") + (v3 + ((((v3 + "") + ((v3 + ("foo bar" + v3)) + "")) + "hello") + v3)))) + (v3 + v3))
    v21 = (v3 + (("foo bar" + v3) + (((v3 + v3) + (v3 + v21)) + v21)))
    v20 = 843
    v22 := GenS3{F0: 0, F1: 98}
    v23 := GenS2{F0: 255}
    if (!(560 == 42)) {
        for fi24 := 0; fi24 < 15; fi24++ {
            v13 = (v13 & v13)
            v25 := []int{32767, fi24}
            var v26 uint8 = 23
            _ = fi24
            _ = v25
            _ = v26
        }
        fmt.Println(v13, v21)
    } else {
        switch v21 {
        case "":
            v27 := GenS1{F0: 1000, F1: 127, F2: 1}
            _ = v27
        case "a":
            var v28 uint64 = 520
            _ = v28
        case "foo bar":
            if false {
                var v29 int32 = 1
                if ((1e10 + 0.1) < (1.5 + 3.14159)) {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v31 := []int{(fi30 / ((fi30) | 1)), fi30, (fi30 / ((fi30) | 1)), fi30}
                        v3 = "hello"
                        var v32 int16 = 1
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
                v33 := GenS4{F0: (((v21 + ((v3 + "\n\t") + ("\n\t" + "hello"))) + ("a" + "a")) + "a")}
                _ = v29
                _ = v33
            }
        }
        v21 = ((v3 + (v3 + "\n\t")) + "\n\t")
        if false {
            var v34 int = 32768
            var v35 uint16 = (v4 + (v4 * (^(v4))))
            _ = v34
            _ = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
