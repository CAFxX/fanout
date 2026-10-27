package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{}
    var v3 int32 = 2147483647
    v4 := make(map[uint32]int)
    if len(v2) > 0 {
        v2[0] = ""
    }
    fmt.Println(v3)
    _ = len(v1)
    var v5 int8 = 100
    for fi6 := 0; fi6 < 19; fi6++ {
        if (((2147483647 >= 506) && true) || true) {
            var v7 int16 = 1
            v7 -= v7
            v5 -= int8(v7)
            _ = v7
        } else {
            var v8 bool = true
            _ = v8
        }
        v9 := GenS1{F0: v5}
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        if ((!(true || false)) || true) {
            for fi11 := 0; fi11 < 17; fi11++ {
                if false {
                    if false {
                        v3 -= v3
                        var v12 bool = (255 > 255)
                        _ = v12
                    }
                } else {
                    switch "\n\t" {
                    case "":
                        v5 -= (v5 >> ((2) & 63))
                    default:
                        switch "hello" {
                        case "foo bar":
                            switch (fi11 | fi11) {
                            case -1:
                                v3 += (v3 / ((v3) | 1))
                            case 316:
                                v3 -= (v3 << ((42) & 63))
                            default:
                                switch 42 {
                                case 128:
                                    var v13 int8 = (v5 << ((208) & 63))
                                    _ = v13
                                case 42:
                                    var v14 uint = 1
                                    _ = v14
                                case 4294967295:
                                    var v15 string = ""
                                    _ = v15
                                }
                            }
                        default:
                            v16 := GenS1{F0: v5}
                            _ = v16
                        }
                    }
                }
                var v17 int16 = 1
                var v18 int32 = v3
                _ = fi11
                _ = v17
                _ = v18
            }
            v3 -= -2147483648
            switch "foo bar" {
            case "a":
                for fi19 := 0; fi19 < 18; fi19++ {
                    var v20 int16 = 0
                    if false {
                        var v21 int = (fi10 & (fi19 >> ((255) & 63)))
                        _ = v21
                    }
                    fmt.Println(fi10)
                    _ = fi19
                    _ = v20
                }
            }
        } else {
            v22 := GenS1{F0: 1}
            _ = v22
        }
        _ = fi10
    }
    v23 := v1[2147483647]
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch ("foo bar" + ((("foo bar" + ((("" + "\n\t") + ("a" + ("a" + "foo bar"))) + "")) + "foo bar") + ("" + "a"))) {
    case "":
        if (!(!(((0.1 > 0.0) || (732 >= 881)) && (("" + "\n\t") >= "hello")))) {
            v24 := v1[4294967295]
            var v25 float64 = ((0.0 / 3.14159) / (((float64(((1.5 + (-2.25 * (0.1 - (float32((0.0 - 1e10)) - ((0.0 / -2.25) / 1e10))))) * (0.0 - 1.5))) + -2.25) / 1.0) / 1.5))
            _ = v24
            _ = v25
        } else {
            v5 -= (v5 % ((v5) | 1))
            if true {
                if true {
                    for fi26 := 0; fi26 < 10; fi26++ {
                        for fi27 := 0; fi27 < 18; fi27++ {
                            v28 := v4[(v23 | (v23 & v23))]
                            v29 := GenS1{F0: v5}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        v30 := GenS2{F0: 1}
                        var v31 int64 = 1
                        _ = fi26
                        _ = v30
                        _ = v31
                    }
                    v3 += v3
                    var v32 int8 = 0
                    _ = v32
                }
            }
            _ = len(v1)
        }
    }
    var v33 int32 = v3
    switch "hello" {
    case "a":
        var v34 uint16 = 42
        _ = v34
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = ""
        }
    case "foo bar":
        v33 = 0
    }
    if (1 > 42) {
        var v35 uint16 = 881
        for fi36 := 0; fi36 < 11; fi36++ {
            v37 := GenS3{F0: 159, F1: -1, F2: (v35 / (((v35 >> ((0) & 63))) | 1))}
            _ = fi36
            _ = v37
        }
        _ = v35
    } else {
        fmt.Println(v3, v3, v3)
        v38 := GenS4{F0: -2147483648, F1: 42, F2: 1}
        v5 = int8(v23)
        _ = v38
    }
    v39 := GenS4{F0: (v3 << ((703) & 63)), F1: 2, F2: 128}
    var v40 float64 = (0.0 / 0.1)
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
