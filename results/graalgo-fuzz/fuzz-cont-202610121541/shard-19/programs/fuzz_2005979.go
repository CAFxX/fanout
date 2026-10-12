package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int; F2 int32 }
type GenS4 struct { F0 uint; F1 bool; F2 int32 }
func genFn0(p0 string, p1 uint) int32 {
    fmt.Println(p0, p0)
    switch p0 {
    case "foo bar":
        var v5 int8 = 1
        _ = v5
    case "hello":
        var v6 int8 = -5
        _ = v6
    case "\n\t":
        switch p0 {
        case "":
            switch -659 {
            case 32767:
                p1 += (p1 >> (((p1 << (((p1 | p1)) & 63))) & 63))
            case 127:
                v7 := []int{128}
                _ = v7
            case 4294967295:
                p0 = "a"
            }
        case "hello":
            var v8 int32 = -2147483648
            _ = v8
        case "a":
            fmt.Println(p0)
        default:
            v9 := []int{-9223372036854775808, 0, int(p1)}
            _ = v9
        }
    }
    v10 := []int{}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    p0 = ("\n\t" + (p0 + "foo bar"))
    _ = p0
    _ = p1
    _ = v10
    return 694
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 uint8 = 0
    v4 := []string{"a", ("a" + ("a" + "a")), "hello"}
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: 77, F1: 127, F2: -708}
        _ = fi11
        _ = v12
    }
    v3 = (^(v3))
    if false {
        var v13 uint16 = 948
        _ = len(v4)
        _ = v13
    } else {
        for fi14 := 0; fi14 < 10; fi14++ {
            if false {
                v15 := GenS1{F0: 128, F1: 92, F2: genFn0((("\n\t" + "hello") + "a"), 1)}
                _ = v15
            }
            fmt.Println(fi14, fi14)
            var v16 int16 = -32768
            _ = fi14
            _ = v16
        }
        if false {
            var v17 uint = 45
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 bool = (669 >= 2147483647)
                _ = fi18
                _ = v19
            }
            if (true || true) {
                v3 += v3
                v17 -= v17
                switch ("a" + (("" + "\n\t") + ((("" + ("hello" + "\n\t")) + "\n\t") + ""))) {
                case "\n\t":
                    for fi20 := 0; fi20 < 19; fi20++ {
                        fmt.Println(v3, v17)
                        _ = fi20
                    }
                case "foo bar":
                    v21 := v2[-2147483648]
                    _ = v21
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    v22 := GenS2{F0: (v3 | (v3 - v3))}
                    _ = v22
                }
            } else {
                var v23 int = 2147483647
                _ = len(v4)
                v24 := GenS3{F0: (v23 >> ((v17) & 63)), F1: v23, F2: genFn0("\n\t", 758)}
                _ = v23
                _ = v24
            }
            _ = v17
        } else {
            var v25 string = ""
            v26 := GenS2{F0: v3}
            _ = v25
            _ = v26
        }
    }
    v27 := GenS4{F0: 0, F1: ("a" < "\n\t"), F2: -278}
    v3 += v3
    v28 := GenS3{F0: 2147483647, F1: 425, F2: genFn0("", 4294967295)}
    fmt.Println(v3)
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 int = fi29
        _ = fi29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
