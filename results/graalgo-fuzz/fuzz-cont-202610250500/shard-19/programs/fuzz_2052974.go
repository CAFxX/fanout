package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 bool }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0() int8 {
    var v5 string = ""
    if false {
        if (true && (!(32767 != -300))) {
            v5 = ((v5 + "\n\t") + "a")
            for fi6 := 0; fi6 < 7; fi6++ {
                for fi7 := 0; fi7 < 7; fi7++ {
                    fmt.Println(v5, fi7)
                    if false {
                        var v8 int64 = 95
                        v9 := GenS1{F0: -1, F1: (v8 != -9223372036854775808), F2: (!(true || false))}
                        v5 = v5
                        _ = v8
                        _ = v9
                    }
                    var v10 int8 = 0
                    _ = fi7
                    _ = v10
                }
                v11 := GenS2{F0: 9223372036854775807, F1: 42}
                _ = fi6
                _ = v11
            }
        } else {
            if (!true) {
                for fi12 := 0; fi12 < 9; fi12++ {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 string = v5
                        switch v14 {
                        case "":
                            var v15 uint8 = 129
                            _ = v15
                        }
                        _ = fi13
                        _ = v14
                    }
                    _ = fi12
                }
                switch "\n\t" {
                case "\n\t":
                    if true {
                        v5 = (v5 + v5)
                    } else {
                        var v16 int32 = -1
                        v16 = v16
                        _ = v16
                    }
                case "":
                    v5 = v5
                case "a":
                    v17 := GenS1{F0: -1, F1: true, F2: true}
                    _ = v17
                default:
                    v5 = v5
                }
                for fi18 := 0; fi18 < 16; fi18++ {
                    for fi19 := 0; fi19 < 2; fi19++ {
                        v5 = "foo bar"
                        var v20 string = "a"
                        _ = fi19
                        _ = v20
                    }
                    var v21 int = (fi18 - fi18)
                    if (!false) {
                        var v22 float64 = (float64(0.1) - 3.14159)
                        _ = v22
                    }
                    _ = fi18
                    _ = v21
                }
            }
            if true {
                v5 = v5
                switch 2147483647 {
                case -2147483648:
                    var v23 uint8 = 1
                    _ = v23
                }
            } else {
                var v24 uint32 = 71
                if (true || false) {
                    fmt.Println(v24)
                } else {
                    if false {
                        if false {
                            fmt.Println(v5, v24)
                            v25 := GenS1{F0: 0, F1: (!(65535 < 255)), F2: ((v24 - v24) <= v24)}
                            _ = v25
                        }
                        v26 := []int{42, -2147483648}
                        _ = v26
                    }
                    v27 := GenS3{F0: -1, F1: 4294967295, F2: 32767}
                    for fi28 := 0; fi28 < 14; fi28++ {
                        var v29 uint = 1
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                }
                _ = v24
            }
            fmt.Println(v5)
        }
    }
    var v30 float32 = 1.5
    v30 += v30
    _ = v30
    _ = v5
    return 1
}

func main() {
    var v1 int16 = 1
    var v2 uint16 = 2
    var v3 string = "\n\t"
    v4 := []uint32{2147483648}
    if false {
        fmt.Println(v1)
    } else {
        var v31 uint8 = uint8(v2)
        for fi32 := 0; fi32 < 11; fi32++ {
            fmt.Println(fi32, v3, v2)
            _ = fi32
        }
        v31 = 255
        _ = v31
    }
    fmt.Println(v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v33 string = ("hello" + ((((v3 + (v3 + v3)) + v3) + "") + (v3 + v3)))
    var v34 float64 = (1e10 / 1.0)
    v35 := GenS4{F0: 527, F1: -759}
    var v36 int = -52
    switch v33 {
    case "\n\t":
        v2 -= v2
    case "":
        switch "foo bar" {
        case "hello":
            v37 := GenS4{F0: 60, F1: 1}
            _ = v37
        }
    case "foo bar":
        v33 = ("a" + (v3 + ((("foo bar" + "a") + (v3 + v33)) + v3)))
    }
    v3 = ""
    _ = len(v4)
    var v38 float32 = float32(((v34 - 1.5) + 0.1))
    var v39 int8 = genFn0()
    v38 += v38
    _ = len(v4)
    fmt.Println(v2, v38, v3)
    v38 += (v38 - 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v39
    _ = v4
}
