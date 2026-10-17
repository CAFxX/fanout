package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 int16; F2 uint16 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 bool; F1 int16; F2 uint16 }
func genFn0(p0 int16, p1 string, p2 int16) uint16 {
    v6 := GenS1{F0: 281, F1: -376}
    if (!true) {
        var v7 float64 = 1.5
        _ = v7
    } else {
        p2 -= (^(p2))
        var v8 float32 = (1.5 + (0.0 - 0.1))
        var v9 uint8 = 155
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1(p0 bool, p1 int) bool {
    p0 = ((0.1 >= (-2.25 * (1.5 / 0.1))) || false)
    for fi10 := 0; fi10 < 20; fi10++ {
        p0 = (-1 != 1)
        v11 := []int{(fi10 / (((fi10 - (p1 * (p1 >> ((1) & 63))))) | 1)), fi10, 0}
        _ = fi10
        _ = v11
    }
    fmt.Println(p0)
    var v12 uint8 = 255
    _ = p0
    _ = p1
    _ = v12
    return (float64((1e10 / float32(1.5))) >= -2.25)
}

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    var v3 float64 = 1e10
    v4 := []uint32{301, 927}
    v5 := make(map[uint32]bool)
    fmt.Println(v3)
    switch 32767 {
    case 99:
        v3 = ((v3 - 1e10) * v3)
    case 127:
        var v13 int16 = 1000
        _ = v13
    default:
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint16 = 2
            v16 := GenS2{F0: 919}
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        delete(v5, 791)
        v3 -= 3.14159
        v1 = append(v1, fi17)
        _ = fi17
    }
    var v18 int64 = 1
    v19 := GenS3{F0: -1, F1: 418, F2: genFn0(42, ("a" + "a"), 1000)}
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 uint32 = 1
        v22 := GenS1{F0: 1, F1: 956}
        v18 += (v18 * 0)
        _ = fi20
        _ = v21
        _ = v22
    }
    if ((2 > 914) && (1 != 127)) {
        if (0 <= 2) {
            if (4294967295 > 218) {
                switch (("hello" + "\n\t") + ("hello" + "hello")) {
                case "a":
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v3 = float64(3.14159)
                        var v24 bool = (true || (!true))
                        var v25 string = ("hello" + "")
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case "":
                    fmt.Println(v18, v3)
                case "foo bar":
                    v3 = 1e10
                }
                if (true && true) {
                    switch "foo bar" {
                    case "\n\t":
                        v26 := GenS3{F0: -1, F1: 1, F2: 65535}
                        _ = v26
                    }
                    v27 := GenS4{F0: 42, F1: true}
                    _ = v27
                }
            }
            fmt.Println(v3)
        }
        if true {
            if false {
                switch "foo bar" {
                case "hello":
                    for fi28 := 0; fi28 < 19; fi28++ {
                        v2[-941] = "hello"
                        var v29 uint64 = 42
                        var v30 int32 = 1
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case "\n\t":
                    switch "a" {
                    case "a":
                        var v31 uint16 = 255
                        _ = v31
                    case "\n\t":
                        v32 := GenS5{F0: false, F1: -1, F2: 978}
                        _ = v32
                    case "":
                        var v33 int32 = 0
                        _ = v33
                    default:
                        v4 = append(v4, 65535)
                    }
                default:
                    if (!false) {
                        v18 += v18
                    } else {
                        for fi34 := 0; fi34 < 6; fi34++ {
                            v18 -= 751
                            var v35 uint32 = 4294967295
                            v36 := GenS2{F0: v35}
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    }
                }
                for fi37 := 0; fi37 < 20; fi37++ {
                    v3 = v3
                    _ = fi37
                }
                v18 = v18
            }
            fmt.Println(v18)
            var v38 float64 = float64(1.5)
            _ = v38
        }
    } else {
        if true {
            fmt.Println(v3, v18)
        }
    }
    v3 -= (-2.25 + (float64((3.14159 * (1e10 * -2.25))) / ((v3 / v3) / (v3 + v3))))
    v18 += v18
    fmt.Println(v3, v3)
    v3 -= (((v3 * 1e10) * 0.1) - 0.1)
    var v39 string = ("\n\t" + (("foo bar" + (("foo bar" + "a") + ("foo bar" + "a"))) + ("" + "a")))
    v40 := GenS2{F0: 65535}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
