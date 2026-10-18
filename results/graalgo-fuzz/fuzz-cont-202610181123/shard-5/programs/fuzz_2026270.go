package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []string{"hello"}
    var v2 int8 = 100
    var v3 float64 = 0.0
    var v4 uint8 = 0
    var v5 uint64 = 0
    var v6 uint8 = (v4 ^ (v4 % ((v4) | 1)))
    if true {
        var v7 string = "foo bar"
        v7 = (v7 + (v7 + "\n\t"))
        var v8 uint32 = 1
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 903, F1: 2147483648}
        _ = v9
    }
    v5 += v5
    v10 := GenS2{F0: (v6 * v6)}
    for fi11 := 0; fi11 < 16; fi11++ {
        switch fi11 {
        case 42:
            switch 1 {
            case 0:
                for fi12 := 0; fi12 < 6; fi12++ {
                    fmt.Println(fi12, fi11, fi12)
                    _ = fi12
                }
            }
        default:
            var v13 string = (("a" + (("\n\t" + "a") + "")) + ("foo bar" + ("" + "foo bar")))
            _ = v13
        }
        v5 = v5
        v6 = (v6 + v6)
        _ = fi11
    }
    v14 := GenS3{F0: 255, F1: 480, F2: 255}
    if true {
        var v15 int64 = 1
        _ = v15
    } else {
        for fi16 := 0; fi16 < 19; fi16++ {
            if true {
                v3 -= (v3 / -2.25)
                v17 := GenS4{F0: 961}
                _ = v17
            } else {
                if (1 >= int(v4)) {
                    v1 = append(v1, ("\n\t" + ("hello" + "foo bar")))
                    var v18 int32 = 413
                    _ = v18
                }
                fmt.Println(v5, v6)
            }
            switch "\n\t" {
            case "":
                v1 = append(v1, "")
            case "foo bar":
                v3 = ((-2.25 - v3) / -2.25)
            case "hello":
                v19 := GenS3{F0: 42, F1: 635, F2: uint16(v5)}
                _ = v19
            }
            var v20 bool = ((663 == -181) && ((v5 % ((v5) | 1)) == v5))
            _ = fi16
            _ = v20
        }
        if true {
            if false {
                v6 = (v4 - (v6 * v4))
                fmt.Println(v6)
            }
            if true {
                if (0 <= 794) {
                    var v21 int8 = (v2 % (((-(v2))) | 1))
                    v2 += v2
                    _ = v21
                }
                switch 2147483647 {
                case 397:
                    if (1 != -1) {
                        v4 -= v4
                        var v22 int32 = 1
                        _ = v22
                    }
                case -559:
                    var v23 float32 = (0.1 - -2.25)
                    _ = v23
                case 128:
                    for fi24 := 0; fi24 < 2; fi24++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v2)
                        _ = fi24
                    }
                }
                fmt.Println(v5)
            } else {
                fmt.Println(v5)
            }
            v1 = append(v1, (("hello" + ("foo bar" + "a")) + "foo bar"))
        } else {
            v25 := GenS2{F0: 191}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (-2147483648 <= 0) {
                v3 = 1e10
            } else {
                fmt.Println(v5)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v25
        }
    }
    fmt.Println(v4)
    fmt.Println(v5)
    v26 := GenS2{F0: v4}
    var v27 float64 = 0.0
    fmt.Println(v27, v2, v6)
    var v28 uint16 = 2
    fmt.Println(v4, v6, v2)
    var v29 uint = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
