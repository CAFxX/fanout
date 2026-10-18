package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint; F2 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 0
    var v3 int16 = 42
    var v4 int32 = -2147483648
    for fi5 := 0; fi5 < 4; fi5++ {
        if true {
            var v6 int = fi5
            v4 -= 42
            _ = v6
        } else {
            if true {
                v7 := GenS1{F0: 42}
                var v8 int16 = (v3 - (v3 - (v3 / (((v3 % (((-(v3))) | 1))) | 1))))
                _ = v7
                _ = v8
            }
            switch "hello" {
            case "hello":
                var v9 bool = true
                _ = v9
            }
            switch "" {
            case "":
                v4 -= -2147483648
            case "foo bar":
                if (42 <= 0) {
                    v10 := []int{(fi5 | 127), (fi5 << (((v2 % ((v2) | 1))) & 63))}
                    v11 := GenS1{F0: 82}
                    _ = v10
                    _ = v11
                }
            default:
                var v12 bool = (v2 >= 4294967295)
                _ = v12
            }
        }
        var v13 uint64 = 42
        _ = fi5
        _ = v13
    }
    if ((193 <= 2) || (!true)) {
        v14 := GenS1{F0: 1}
        var v15 bool = false
        if (2147483648 >= 4294967295) {
            if (((("foo bar" + "a") + "\n\t") + "") > ("hello" + "\n\t")) {
                for fi16 := 0; fi16 < 17; fi16++ {
                    fmt.Println(v4, v3, v3)
                    _ = fi16
                }
            }
            var v17 int8 = 8
            v3 = v3
            _ = v17
        }
        _ = v14
        _ = v15
    }
    if true {
        v4 += (v4 >> (((^(v2))) & 63))
        v3 = v3
        var v18 int = 106
        _ = v18
    } else {
        v19 := GenS2{F0: false, F1: 255, F2: -29}
        fmt.Println(v3, v4, v4)
        _ = v19
    }
    v2 -= (v2 / (((v2 & (v2 & v2))) | 1))
    v20 := GenS2{F0: true, F1: (v2 << ((18446744073709551615) & 63)), F2: -83}
    v21 := []int{}
    if (255 != 0) {
        if true {
            switch 2147483647 {
            case 127:
                switch "foo bar" {
                case "hello":
                    var v22 uint16 = 42
                    _ = v22
                case "a":
                    v4 += (v4 + (^(v4)))
                default:
                    _ = len(v21)
                }
            case 2147483647:
                if (("\n\t" + ("hello" + "hello")) > "") {
                    for fi23 := 0; fi23 < 9; fi23++ {
                        for fi24 := 0; fi24 < 19; fi24++ {
                            fmt.Println(fi23, v4, fi24)
                            var v25 int16 = (v3 + v3)
                            var v26 int64 = 346
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        var v27 string = ("a" + "a")
                        if (-32768 >= 0) {
                            var v28 uint8 = 128
                            _ = v28
                        }
                        _ = fi23
                        _ = v27
                    }
                    v29 := v1["a"]
                    var v30 uint16 = 32767
                    _ = v29
                    _ = v30
                } else {
                    var v31 float32 = 1e10
                    v32 := v1["foo bar"]
                    _ = v31
                    _ = v32
                }
            case 42:
                v3 = v3
            default:
                for fi33 := 0; fi33 < 9; fi33++ {
                    v2 = (v2 | (v2 % ((v2) | 1)))
                    var v34 string = (((("hello" + ("hello" + "hello")) + ("\n\t" + "\n\t")) + ("\n\t" + "hello")) + "hello")
                    var v35 int64 = 514
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            }
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            v4 += (^(v4))
        } else {
            var v36 uint32 = 2147483648
            v37 := v1[("foo bar" + "")]
            _ = v36
            _ = v37
        }
        switch 1 {
        case 9223372036854775807:
            v2 = (v2 << ((v2) & 63))
        case 128:
            var v38 float32 = 0.1
            _ = v38
        case 315:
            v1["foo bar"] = (false && true)
        }
    }
    if false {
        for fi39 := 0; fi39 < 10; fi39++ {
            v2 = (v2 % ((18446744073709551615) | 1))
            for fi40 := 0; fi40 < 20; fi40++ {
                fmt.Println(v4)
                var v41 string = ("a" + ("\n\t" + "a"))
                _ = fi40
                _ = v41
            }
            _ = fi39
        }
        var v42 uint32 = 956
        _ = v42
    }
    for fi43 := 0; fi43 < 4; fi43++ {
        v44 := GenS2{F0: true, F1: v2, F2: 0}
        v4 = -354
        var v45 int16 = (v3 - v3)
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
