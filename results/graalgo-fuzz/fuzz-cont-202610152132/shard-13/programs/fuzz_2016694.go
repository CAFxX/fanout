package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int8 }
type GenS2 struct { F0 string; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint8 }
func genFn0(p0 int64) uint {
    fmt.Println(p0)
    p0 += p0
    p0 -= p0
    _ = p0
    return 42
}
func genFn1(p0 string, p1 bool, p2 string) uint32 {
    if p1 {
        p0 = ("\n\t" + "")
        if ((((((float64(1.5) - 0.0) + 1.5) / (((3.14159 * (((float64(((0.0 / float32(0.1)) / 1.0)) - ((float64(1e10) / 1.5) + 0.1)) * 0.1) / -2.25)) + 3.14159) + (0.1 + float64(-2.25)))) - 0.1) <= float64((float32(-2.25) + (((0.0 - 3.14159) * -2.25) - 0.1)))) || (-2147483648 >= 0)) {
            var v7 int64 = 629
            var v8 uint8 = 0
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: "a", F1: p1, F2: 100}
        _ = v9
    }
    var v10 int64 = -9223372036854775808
    if true {
        switch ("\n\t" + "hello") {
        case "hello":
            p0 = ((p0 + ((p2 + p0) + ("a" + (p2 + ((p0 + p0) + ("foo bar" + p0)))))) + "")
        case "\n\t":
            p2 = "a"
        case "a":
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := []int{fi11}
                _ = fi11
                _ = v12
            }
        default:
            var v13 float32 = 0.0
            _ = v13
        }
        var v14 int64 = v10
        _ = v14
    } else {
        v15 := GenS1{F0: "foo bar", F1: true, F2: 127}
        v16 := GenS2{F0: "a", F1: 0, F2: v10}
        _ = v15
        _ = v16
    }
    p0 = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 273
}

func main() {
    var v1 int = -697
    v2 := []string{"\n\t", "", "foo bar"}
    var v3 uint8 = 0
    v4 := []uint32{1, 810}
    v5 := []uint32{2147483647, 42, 2147483648}
    var v6 int64 = -1
    var v17 float64 = (3.14159 + 3.14159)
    _ = len(v4)
    for fi18 := 0; fi18 < 16; fi18++ {
        v6 = (v6 + (v6 + (v6 % ((v6) | 1))))
        v3 = (v3 % (((v3 ^ (v3 - (v3 << ((genFn0((v6 * (v6 / ((v6) | 1))))) & 63))))) | 1))
        _ = len(v2)
        _ = fi18
    }
    var v19 string = ""
    v4 = append(v4, 42)
    _ = len(v2)
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint = 4294967295
        if (!(!(1.5 > v17))) {
            var v22 float64 = ((v17 - 1e10) - ((0.1 / (((0.1 * (3.14159 - (v17 * (v17 * v17)))) * v17) / ((v17 + v17) - 0.1))) - float64((1e10 + 0.0))))
            _ = v22
        } else {
            var v23 int8 = 127
            var v24 uint64 = 926
            if true {
                if true {
                    for fi25 := 0; fi25 < 18; fi25++ {
                        var v26 uint8 = (v3 & v3)
                        switch int(v24) {
                        case 2147483647:
                            v27 := GenS3{F0: 32767, F1: 42, F2: v26}
                            _ = v27
                        case 1:
                            var v28 uint32 = 42
                            _ = v28
                        default:
                            if (1 != 1) {
                                fmt.Println(v3)
                            } else {
                                if false {
                                    _ = len(v4)
                                    switch v19 {
                                    case "\n\t":
                                        if true {
                                            var v29 int8 = (-(v23))
                                            switch (fi25 | v1) {
                                            case 214:
                                                v2 = append(v2, (((v19 + "hello") + "foo bar") + ((v19 + v19) + ("foo bar" + v19))))
                                            case -1:
                                                switch fi25 {
                                                case 403:
                                                    fmt.Println(v24, v24, v3)
                                                case -1:
                                                    v30 := GenS2{F0: v19, F1: 67, F2: -9223372036854775808}
                                                    _ = v30
                                                }
                                            }
                                            _ = v29
                                        } else {
                                            v5 = append(v5, 250)
                                        }
                                    }
                                    var v31 uint32 = 135
                                    var v32 float64 = v17
                                    fmt.Println(v23)
                                    _ = v31
                                    _ = v32
                                }
                            }
                        }
                        _ = fi25
                        _ = v26
                    }
                    fmt.Println(v1, v24, v6)
                    v1 = (v1 << ((genFn0(int64(v24))) & 63))
                } else {
                    v1 = (-(v1))
                    var v33 int16 = 32767
                    _ = v33
                }
                fmt.Println(v24, v21)
                v19 = ((v19 + v19) + (("foo bar" + (v19 + ("" + v19))) + (v19 + ("\n\t" + "hello"))))
            } else {
                var v34 bool = false
                _ = v34
            }
            _ = v23
            _ = v24
        }
        v19 = v19
        _ = fi20
        _ = v21
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    fmt.Println(v1)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
