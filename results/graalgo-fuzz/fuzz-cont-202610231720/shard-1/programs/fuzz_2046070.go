package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int }
func genFn0() int64 {
    switch "" {
    case "hello":
        v7 := GenS1{F0: 1, F1: 2, F2: 255}
        _ = v7
    case "\n\t":
        v8 := GenS2{F0: -32768}
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 4; fi9++ {
            if false {
                switch "hello" {
                case "foo bar":
                    v10 := GenS2{F0: 0}
                    _ = v10
                case "\n\t":
                    var v11 int64 = 9223372036854775807
                    _ = v11
                case "":
                    if (4294967295 < 1) {
                        var v12 int32 = -2147483648
                        _ = v12
                    }
                }
            } else {
                for fi13 := 0; fi13 < 7; fi13++ {
                    var v14 bool = (2147483647 >= 2147483648)
                    if (false || v14) {
                        v14 = false
                    }
                    _ = fi13
                    _ = v14
                }
                v15 := []int{fi9}
                _ = v15
            }
            v16 := GenS2{F0: 0}
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := GenS3{F0: 891, F1: 27, F2: 65535}
                _ = fi17
                _ = v18
            }
            _ = fi9
            _ = v16
        }
    }
    var v19 uint32 = 42
    _ = v19
    return 981
}
func genFn1(p0 string) uint32 {
    for fi20 := 0; fi20 < 5; fi20++ {
        switch 32768 {
        case -2147483648:
            v21 := GenS2{F0: -283}
            _ = v21
        case -154:
            var v22 uint8 = 127
            _ = v22
        case 32768:
            v23 := GenS2{F0: 0}
            _ = v23
        default:
            v24 := GenS4{F0: 323}
            _ = v24
        }
        _ = fi20
    }
    v25 := []int{-1, -58, -881, 128}
    for fi26 := 0; fi26 < 8; fi26++ {
        p0 = p0
        var v27 float32 = 3.14159
        var v28 uint64 = 18446744073709551615
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 int64 = -313
            fmt.Println(v31)
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    _ = p0
    _ = v25
    return 809
}

func main() {
    var v1 float64 = (3.14159 * (0.1 * 3.14159))
    v2 := []string{(("a" + "hello") + ("hello" + "\n\t"))}
    var v3 uint8 = 255
    var v4 float32 = 1e10
    var v5 int32 = 1
    v6 := make(map[int]string)
    if true {
        for fi32 := 0; fi32 < 20; fi32++ {
            var v33 int8 = 127
            var v34 uint = 1
            v6[fi32] = ((("" + "") + ("foo bar" + ("" + ""))) + "foo bar")
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    var v35 uint8 = (v3 ^ 42)
    var v36 uint16 = 0
    var v37 int32 = v5
    var v38 int64 = 9223372036854775807
    for fi39 := 0; fi39 < 5; fi39++ {
        var v40 uint16 = v36
        var v41 uint64 = uint64(v35)
        v35 = (v3 * v3)
        _ = fi39
        _ = v40
        _ = v41
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 = (^(v5))
    delete(v6, 32767)
    var v42 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
