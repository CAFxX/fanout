package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string }
func genFn0(p0 string) uint16 {
    v5 := GenS1{F0: 42, F1: 1}
    if true {
        v6 := []int{32767, -713, 4294967295}
        fmt.Println(p0)
        _ = v6
    } else {
        switch -1 {
        case -2147483648:
            if true {
                var v7 string = "foo bar"
                _ = v7
            }
        case 127:
            fmt.Println(p0)
        default:
            switch p0 {
            case "hello":
                if true {
                    switch 4294967295 {
                    case -2147483648:
                        v8 := []int{}
                        _ = v8
                    case 32767:
                        v9 := GenS1{F0: 1, F1: 771}
                        _ = v9
                    case 2147483647:
                        p0 = ""
                    }
                }
            default:
                fmt.Println(p0)
            }
        }
        p0 = "a"
    }
    _ = p0
    _ = v5
    return 2
}
func genFn1(p0 string) uint32 {
    if (552 == 65535) {
        p0 = ("" + p0)
    }
    v10 := GenS2{F0: "a"}
    var v11 uint8 = 255
    _ = p0
    _ = v10
    _ = v11
    return 65535
}

func main() {
    var v1 int = 339
    var v2 bool = true
    var v3 int = v1
    v4 := make(map[uint32]bool)
    fmt.Println(v1, v2, v3)
    var v12 string = ("foo bar" + (("" + ("hello" + "")) + "hello"))
    if ((1 >= 255) || v2) {
        v13 := []int{}
        var v14 int64 = 9223372036854775807
        _ = v13
        _ = v14
    }
    switch 405 {
    case 32767:
        switch (v3 ^ -320) {
        case 841:
            var v15 int8 = 100
            _ = v15
        case 128:
            v16 := GenS3{F0: -903}
            _ = v16
        default:
            v3 = v1
        }
    case 955:
        if v2 {
            fmt.Println(v2)
            v17 := v4[1]
            fmt.Println(v17, v1)
            _ = v17
        } else {
            v18 := GenS4{F0: 1}
            var v19 int8 = 88
            v19 -= (v19 | v19)
            _ = v18
            _ = v19
        }
    case 618:
        if true {
            v3 += -9223372036854775808
        } else {
            for fi20 := 0; fi20 < 4; fi20++ {
                if false {
                    switch v12 {
                    case "":
                        for fi21 := 0; fi21 < 10; fi21++ {
                            fmt.Println(v1, fi20, fi20)
                            v22 := GenS3{F0: 1}
                            _ = fi21
                            _ = v22
                        }
                    default:
                        v12 = v12
                    }
                    v12 = v12
                } else {
                    v12 = v12
                    _ = len(v4)
                    v23 := GenS4{F0: 42}
                    _ = v23
                }
                v12 = (v12 + ((((v12 + ("a" + (v12 + ((v12 + v12) + v12)))) + v12) + v12) + "a"))
                fmt.Println(v12, v2, fi20)
                _ = fi20
            }
        }
    }
    if ((4294967295 != 1) || (v2 && (!((v12 + v12) <= (v12 + v12))))) {
        switch -2147483648 {
        case 0:
            v12 = ""
        case 42:
            delete(v4, 76)
        default:
            var v24 int64 = -530
            _ = v24
        }
    }
    v25 := GenS5{F0: "hello"}
    v26 := GenS4{F0: 2}
    v2 = v2
    v2 = (genFn1(v12) <= 1)
    v3 = (v3 | v3)
    var v27 string = v12
    var v28 uint32 = 65535
    var v29 uint64 = 948
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
