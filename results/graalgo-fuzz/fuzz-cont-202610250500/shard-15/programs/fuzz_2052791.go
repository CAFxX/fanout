package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 int32 }
type GenS3 struct { F0 bool; F1 bool; F2 int }
type GenS4 struct { F0 int8; F1 int8 }
func genFn0() uint64 {
    v5 := []int{}
    v6 := GenS1{F0: (("hello" + "") + "hello"), F1: 32767}
    _ = v5
    _ = v6
    return 151
}
func genFn1(p0 int, p1 uint64, p2 string) uint32 {
    var v7 int8 = 0
    if false {
        for fi8 := 0; fi8 < 9; fi8++ {
            switch ("a" + p2) {
            case "a":
                var v9 uint8 = 108
                _ = v9
            case "":
                v7 -= (-(v7))
            case "foo bar":
                var v10 float32 = (0.0 - (((-2.25 - 0.1) / 1.0) - 0.1))
                _ = v10
            default:
                if (((true && (true || true)) || true) && true) {
                    v11 := GenS1{F0: p2, F1: 962}
                    v12 := []int{(p0 >> ((127) & 63)), fi8}
                    _ = v11
                    _ = v12
                } else {
                    var v13 uint64 = 18446744073709551615
                    v14 := []int{fi8, 42, (fi8 % ((int(v7)) | 1)), (fi8 << ((255) & 63))}
                    _ = len(v14)
                    _ = v13
                    _ = v14
                }
            }
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{412, 0, 154, 662}
    v3 := make(map[string]uint32)
    v4 := []int{567, 815, 2147483647, 2147483647}
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int16 = -979
        if true {
            v17 := GenS2{F0: 95, F1: 0, F2: -2147483648}
            _ = v17
        } else {
            var v18 int8 = 42
            fmt.Println(v18, v18)
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint32 = 76
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        if (true || (!(false || true))) {
            var v22 uint16 = 924
            _ = v22
        }
        fmt.Println(fi21)
        var v23 uint64 = 42
        _ = fi21
        _ = v23
    }
    var v24 string = ("a" + "a")
    switch v24 {
    case "foo bar":
        var v25 uint = 89
        _ = v25
    default:
        v26 := GenS3{F0: (false || false), F1: true, F2: -230}
        _ = v26
    }
    fmt.Println(v24)
    for fi27 := 0; fi27 < 4; fi27++ {
        fmt.Println(fi27, v24)
        delete(v3, ((((v24 + ((v24 + (v24 + (v24 + "foo bar"))) + (v24 + (v24 + v24)))) + ("hello" + v24)) + v24) + (v24 + (v24 + v24))))
        _ = fi27
    }
    var v28 int32 = 1
    var v29 uint32 = 131
    fmt.Println(v29, v29, v29)
    fmt.Println(v29)
    if (((((("hello" + v24) + "\n\t") + (v24 + "a")) + "a") + v24) == (v24 + ("foo bar" + v24))) {
        switch 32768 {
        case 0:
            v28 -= (v28 >> ((344) & 63))
        case 32767:
            v30 := GenS4{F0: 42, F1: -45}
            _ = v30
        case -335:
            if false {
                for fi31 := 0; fi31 < 7; fi31++ {
                    delete(v3, v24)
                    var v32 uint64 = genFn0()
                    _ = fi31
                    _ = v32
                }
                var v33 int = 1
                _ = v33
            }
        default:
            v24 = ("\n\t" + v24)
        }
        v24 = "\n\t"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
