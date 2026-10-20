package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint32, p1 string, p2 uint32) uint {
    if (p1 <= (((p1 + ("\n\t" + (p1 + ((p1 + "\n\t") + "\n\t")))) + "a") + p1)) {
        p2 = p0
        var v4 uint8 = 42
        p2 = p0
        _ = v4
    } else {
        p0 = (p2 - 938)
    }
    var v5 int64 = -1
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 436
}

func main() {
    var v1 uint8 = 172
    var v2 int = 32767
    v3 := []string{"hello", "\n\t", "\n\t"}
    var v6 int16 = -1
    v7 := GenS1{F0: 696, F1: v6}
    if true {
        var v8 string = "a"
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 bool = true
            var v11 string = (v8 + ((v8 + "foo bar") + (("hello" + "hello") + v8)))
            switch 42 {
            case 548:
                v8 = v8
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int = v2
        _ = v12
        _ = v8
    }
    v2 = (v2 >> ((630) & 63))
    v6 = v6
    v6 = v6
    var v13 int16 = v6
    var v14 uint8 = 127
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v15 := GenS1{F0: -1, F1: -32768}
    for fi16 := 0; fi16 < 12; fi16++ {
        if false {
            if true {
                v17 := GenS1{F0: -2147483648, F1: v6}
                switch v2 {
                case 921:
                    fmt.Println(v6)
                }
                _ = v17
            }
            v1 -= v1
        }
        _ = fi16
    }
    switch 1 {
    case 32767:
        for fi18 := 0; fi18 < 4; fi18++ {
            if false {
                fmt.Println(v13, v13)
            } else {
                for fi19 := 0; fi19 < 7; fi19++ {
                    v1 -= v14
                    switch fi19 {
                    case -9223372036854775808:
                        v6 -= v6
                    case 2147483647:
                        if true {
                            fmt.Println(fi19, fi19)
                            fmt.Println(v2, v6)
                            v2 = 127
                        } else {
                            v2 -= -426
                        }
                    case 4294967295:
                        v20 := GenS2{F0: 42}
                        _ = v20
                    default:
                        v14 = v14
                    }
                    _ = fi19
                }
                v14 -= v14
                v21 := GenS1{F0: 2147483647, F1: v6}
                _ = v21
            }
            var v22 uint16 = 2
            _ = fi18
            _ = v22
        }
    }
    var v23 float64 = (3.14159 + -2.25)
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v6
    _ = v7
}
