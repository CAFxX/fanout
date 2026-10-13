package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32) uint32 {
    p0 += (p0 % (((-(p0))) | 1))
    p0 += -348
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int8 = 1
    var v3 float64 = 0.1
    var v4 float64 = 1e10
    v5 := []int{-1}
    v6 := make(map[string]int)
    v1 = genFn0(-446)
    for fi10 := 0; fi10 < 15; fi10++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch fi10 {
        case 425:
            var v11 uint8 = 2
            _ = v11
        case -959:
            fmt.Println(v2, v4, v1)
        }
        _ = fi10
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v1)
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := GenS1{F0: (true || ("" <= (("\n\t" + "hello") + ""))), F1: "hello"}
            v2 = (v2 >> ((448) & 63))
            if false {
                var v15 uint = 18446744073709551615
                fmt.Println(fi12)
                switch (^(fi12)) {
                case 4294967295:
                    v16 := GenS1{F0: (((0 > 734) || (65535 >= 65535)) || false), F1: ""}
                    _ = v16
                }
                _ = v15
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v5 = append(v5, -872)
    fmt.Println(v3)
    fmt.Println(v1, v4)
    if true {
        for fi17 := 0; fi17 < 6; fi17++ {
            if (!false) {
                for fi18 := 0; fi18 < 5; fi18++ {
                    v19 := GenS1{F0: ((v3 >= 0.1) || (!false)), F1: "a"}
                    v20 := GenS1{F0: ((-1 >= 42) && false), F1: "a"}
                    if ((fi17 << ((4294967295) & 63)) != (fi18 ^ fi17)) {
                        var v21 int16 = -212
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v3 += v4
                v1 = v1
            }
            v3 -= v4
            v22 := GenS2{F0: 2147483647}
            _ = fi17
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 14; fi23++ {
            if true {
                for fi24 := 0; fi24 < 5; fi24++ {
                    v4 = v4
                    _ = fi24
                }
            } else {
                delete(v6, "")
                if true {
                    v4 = 3.14159
                    var v25 uint32 = v1
                    _ = v25
                } else {
                    v2 -= (v2 % ((0) | 1))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                v4 = (float64(-2.25) + v4)
            }
            _ = fi23
        }
    }
    v3 = 3.14159
    for fi26 := 0; fi26 < 19; fi26++ {
        for fi27 := 0; fi27 < 18; fi27++ {
            if true {
                switch (((("\n\t" + "\n\t") + "") + "\n\t") + ("\n\t" + "a")) {
                case "":
                    _ = len(v6)
                case "a":
                    v2 = (v2 << ((164) & 63))
                case "hello":
                    var v28 int64 = -506
                    _ = v28
                default:
                    if (!true) {
                        var v29 int32 = 1
                        if (!true) {
                            v5 = append(v5, (-(fi27)))
                            if (42 > 1) {
                                var v30 float64 = (((v4 - v3) / (v3 * (v4 * float64(float32(v4))))) - ((float64(3.14159) / v4) / -2.25))
                                _ = v30
                            } else {
                                var v31 string = "\n\t"
                                _ = v31
                            }
                            fmt.Println(v4, v3)
                        }
                        _ = v29
                    }
                }
            }
            if len(v5) > 0 {
                v5[0] = (fi27 * -2147483648)
            }
            var v32 int8 = (-(v2))
            v4 = ((((((0.0 / (v3 - (((v4 - float64((0.1 * 1.5))) / -2.25) + 1e10))) - v4) / v3) / (v3 / 1e10)) / float64(-2.25)) * (v4 / v4))
            _ = fi27
            _ = v32
        }
        for fi33 := 0; fi33 < 5; fi33++ {
            var v34 int = fi33
            fmt.Println(v4)
            _ = fi33
            _ = v34
        }
        _ = fi26
    }
    v3 = v3
    switch 32768 {
    case 1:
        var v35 int32 = 2147483647
        _ = v35
    case 765:
        fmt.Println(v2)
    default:
        var v36 int32 = -1
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
