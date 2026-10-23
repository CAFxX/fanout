package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint }
func genFn0(p0 string) int32 {
    var v4 int32 = 568
    v5 := GenS1{F0: false, F1: 201}
    for fi6 := 0; fi6 < 15; fi6++ {
        switch fi6 {
        case 4294967295:
            fmt.Println(v4, fi6)
        case 1:
            fmt.Println(fi6, v4)
        default:
            v7 := GenS2{F0: 107, F1: 42, F2: 1}
            _ = v7
        }
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: (329 == 65535), F1: 32767}
            if false {
                var v10 uint = 18446744073709551615
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        _ = fi6
    }
    switch 25 {
    case -9223372036854775808:
        p0 = ((p0 + ("\n\t" + (p0 + (p0 + "foo bar")))) + p0)
    case -579:
        v11 := []int{}
        _ = v11
    }
    _ = p0
    _ = v4
    _ = v5
    return (v4 ^ v4)
}

func main() {
    v1 := []string{"foo bar", (((("a" + "") + "hello") + "\n\t") + ((("a" + "\n\t") + "hello") + (("hello" + "\n\t") + "hello"))), "", (("foo bar" + ("a" + "")) + "hello")}
    v2 := []uint32{260}
    v3 := make(map[int]bool)
    switch 1 {
    case 32768:
        if true {
            var v12 float64 = 1.5
            _ = v12
        } else {
            if true {
                var v13 int32 = genFn0("foo bar")
                delete(v3, 128)
                _ = v13
            }
        }
    }
    var v14 int = -1
    var v15 int = v14
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "foo bar")
    }
    _ = len(v3)
    v14 = v14
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v16 uint32 = 998
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 18; fi17++ {
            fmt.Println(v15, v15, v14)
            _ = fi17
        }
        var v18 float32 = (float32((float64(-2.25) / float64(3.14159))) / float32(3.14159))
        v19 := GenS3{F0: 628, F1: 2}
        _ = v18
        _ = v19
    } else {
        if false {
            v14 += -613
            var v20 float32 = 0.1
            _ = v20
        } else {
            var v21 int = (v14 % ((v14) | 1))
            _ = v21
        }
        if false {
            var v22 uint8 = 255
            v14 = int(v22)
            _ = v22
        } else {
            switch 128 {
            case 80:
                fmt.Println(v14)
            case -2147483648:
                v14 -= 127
            case 42:
                var v23 int64 = -950
                _ = v23
            default:
                var v24 uint32 = 0
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            if false {
                var v26 uint = 0
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 uint16 = 374
                    _ = len(v2)
                    fmt.Println(v14, v15)
                    _ = fi27
                    _ = v28
                }
                _ = v26
            } else {
                for fi29 := 0; fi29 < 3; fi29++ {
                    fmt.Println(v14)
                    _ = fi29
                }
            }
            v14 += (v15 >> ((2) & 63))
            fmt.Println(v14, v15, v15)
            _ = fi25
        }
    }
    v15 -= 2147483647
    v14 += v15
    var v30 bool = true
    var v31 int = (v15 << ((18446744073709551615) & 63))
    for fi32 := 0; fi32 < 15; fi32++ {
        if (985 >= 1) {
            fmt.Println(fi32)
            v33 := v3[-873]
            _ = v33
        } else {
            fmt.Println(v14, v14)
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        }
        if true {
            var v34 uint32 = 1
            v31 = 1
            _ = v34
        } else {
            if len(v2) > 0 {
                v2[0] = 65535
            }
            fmt.Println(v15, v30, fi32)
            v3[fi32] = (2147483647 > -234)
        }
        _ = fi32
    }
    var v35 int64 = 42
    var v36 uint16 = 255
    v37 := GenS4{F0: 760}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
}
