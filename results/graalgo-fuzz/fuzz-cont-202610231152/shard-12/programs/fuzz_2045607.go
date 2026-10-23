package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint16; F2 uint16 }
func genFn0(p0 bool) uint16 {
    if p0 {
        v7 := []int{-1, -9223372036854775808, 128, 881}
        var v8 float64 = 3.14159
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 95}
        _ = v9
    }
    fmt.Println(p0)
    var v10 float64 = 0.0
    _ = p0
    _ = v10
    return 772
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{42, 65535, 42}
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    v5 := []string{"hello", ("\n\t" + "a"), "foo bar"}
    v6 := []uint32{0, 2, 65535, 216}
    var v11 string = ""
    switch 42 {
    case 1:
        fmt.Println(v11)
    case 128:
        for fi12 := 0; fi12 < 20; fi12++ {
            if (!((-2.25 - 3.14159) == 0.0)) {
                var v13 float32 = 1e10
                v14 := GenS1{F0: 42}
                v11 = v11
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 3; fi15++ {
                    if false {
                        if len(v5) > 0 {
                            v5[0] = (v11 + "hello")
                        }
                        v11 = (v11 + "\n\t")
                        fmt.Println(fi15)
                    } else {
                        v11 = v11
                        v16 := GenS1{F0: 2}
                        _ = v16
                    }
                    v11 = v11
                    _ = fi15
                }
                var v17 int32 = -1
                v11 = (("a" + (v11 + "a")) + "foo bar")
                _ = v17
            }
            _ = fi12
        }
    case -507:
        v18 := GenS1{F0: 255}
        _ = v18
    }
    var v19 int8 = -1
    v11 = v11
    v20 := GenS1{F0: 66}
    for fi21 := 0; fi21 < 8; fi21++ {
        v19 += v19
        fmt.Println(fi21, v19, v19)
        fmt.Println(v19, v11)
        _ = fi21
    }
    delete(v4, 0)
    v22 := GenS1{F0: 42}
    switch int(v19) {
    case 32767:
        _ = len(v6)
    case 440:
        var v23 int8 = (v19 & (v19 & (-(v19))))
        _ = v23
    default:
        var v24 bool = false
        _ = v24
    }
    fmt.Println(v11)
    v11 = ""
    switch 32768 {
    case -1:
        v25 := GenS2{F0: 42, F1: 1}
        _ = v25
    case 641:
        v19 -= v19
    default:
        switch 676 {
        case 127:
            switch 1 {
            case 4294967295:
                for fi26 := 0; fi26 < 16; fi26++ {
                    if len(v6) > 0 {
                        v6[0] = 329
                    }
                    var v27 int = (fi26 % ((32767) | 1))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi26
                    _ = v27
                }
            }
        }
    }
    if (!(!(1 > 649))) {
        v19 -= 1
        if false {
            delete(v3, 65535)
            fmt.Println(v11, v19)
            if (v11 < v11) {
                v28 := GenS3{F0: ("hello" + v11)}
                v19 = v19
                _ = v28
            } else {
                v5 = append(v5, v11)
                var v29 int16 = -253
                for fi30 := 0; fi30 < 1; fi30++ {
                    v29 -= v29
                    v11 = v11
                    _ = fi30
                }
                _ = v29
            }
        }
        var v31 uint16 = 0
        _ = v31
    }
    if false {
        for fi32 := 0; fi32 < 5; fi32++ {
            v11 = v11
            _ = fi32
        }
        v33 := GenS4{F0: false, F1: 1, F2: 117}
        _ = v33
    }
    switch -1 {
    case 32767:
        for fi34 := 0; fi34 < 19; fi34++ {
            for fi35 := 0; fi35 < 20; fi35++ {
                switch "hello" {
                case "\n\t":
                    var v36 int = (fi34 - fi35)
                    _ = v36
                case "a":
                    switch 2147483647 {
                    case 0:
                        v37 := GenS2{F0: 4294967295, F1: 811}
                        _ = v37
                    case 892:
                        var v38 int32 = -2147483648
                        _ = v38
                    case 390:
                        v11 = "hello"
                    }
                case "hello":
                    var v39 int64 = -457
                    _ = v39
                }
                fmt.Println(fi34)
                _ = fi35
            }
            if len(v6) > 0 {
                v6[0] = 435
            }
            if (935 <= 0) {
                var v40 uint8 = 127
                v11 = ("a" + "")
                _ = v40
            }
            _ = fi34
        }
    case 128:
        v2 = append(v2, 2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
