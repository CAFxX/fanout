package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint32; F1 bool; F2 int }
func genFn0() uint64 {
    v5 := []int{}
    var v6 uint32 = 1
    var v7 string = "a"
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 45
    var v3 bool = true
    v4 := []uint32{uint32(v2), 101, 945}
    if true {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint16 = 0
            _ = fi8
            _ = v9
        }
        v2 += v2
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: (v2 ^ v2), F1: 400, F2: 42}
            _ = fi10
            _ = v11
        }
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    var v12 float64 = (1e10 - float64(1e10))
    if (!v3) {
        if v3 {
            var v13 string = ("a" + ("a" + "\n\t"))
            _ = v13
        }
        fmt.Println(v3, v12, v2)
    } else {
        v2 = 79
        v14 := GenS2{F0: 901}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v14
    }
    if len(v4) > 0 {
        v4[0] = 656
    }
    v15 := GenS3{F0: ""}
    v16 := GenS4{F0: 112, F1: "\n\t"}
    switch "a" {
    case "":
        v17 := GenS2{F0: 42}
        _ = v17
    case "a":
        v12 = v12
    default:
        v3 = (false && (!(!(!v3))))
    }
    var v18 uint = 1
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 uint32 = 2147483648
        _ = fi19
        _ = v20
    }
    if (0 < 0) {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 uint64 = 2
            if ("\n\t" <= "\n\t") {
                var v23 uint16 = 2
                _ = v23
            } else {
                v2 += 0
            }
            _ = fi21
            _ = v22
        }
        v2 = (v2 << (((v2 * (^(v18)))) & 63))
        if v3 {
            v1[-1] = 42
        }
    } else {
        var v24 float64 = v12
        var v25 string = ("foo bar" + "")
        v26 := GenS2{F0: 529}
        _ = v24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        for fi28 := 0; fi28 < 17; fi28++ {
            if v3 {
                if v3 {
                    v2 = v2
                    if (2 == 241) {
                        v29 := v1[(fi28 + (fi27 + (fi27 >> ((v2) & 63))))]
                        if ((fi28 & 127) < (fi28 | 0)) {
                            var v30 uint = 208
                            _ = v30
                        } else {
                            if v3 {
                                switch "foo bar" {
                                case "hello":
                                    switch (fi27 + (-(fi28))) {
                                    case 927:
                                        v18 += v18
                                    }
                                case "\n\t":
                                    v18 = (v2 >> (((v2 & (v18 / ((v18) | 1)))) & 63))
                                case "foo bar":
                                    fmt.Println(v2, v29, fi27)
                                default:
                                    var v31 uint8 = 255
                                    _ = v31
                                }
                                var v32 float32 = (0.0 - (3.14159 - 1.5))
                                _ = v32
                            } else {
                                v29 = (v29 + v29)
                                var v33 uint8 = 255
                                _ = v33
                            }
                        }
                        _ = v29
                    } else {
                        v12 -= float64((0.0 + 1e10))
                    }
                    var v34 bool = true
                    _ = v34
                } else {
                    var v35 string = ("hello" + "\n\t")
                    _ = v35
                }
            }
            _ = fi28
        }
        _ = fi27
    }
    v36 := GenS5{F0: 2147483648, F1: (473 == 2147483647), F2: 127}
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 bool = (0 > 157)
        fmt.Println(v2)
        _ = fi37
        _ = v38
    }
    v39 := GenS1{F0: v2, F1: 42, F2: -1}
    for fi40 := 0; fi40 < 9; fi40++ {
        var v41 uint32 = 2147483648
        _ = fi40
        _ = v41
    }
    v12 = ((v12 - float64(float32(3.14159))) + v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
}
