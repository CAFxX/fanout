package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int16; F2 uint16 }
func genFn0() int16 {
    if false {
        var v5 uint8 = 220
        v6 := []int{2147483647, 42, 128}
        v5 -= (v5 >> ((65535) & 63))
        _ = v5
        _ = v6
    } else {
        v7 := []int{}
        _ = v7
    }
    var v8 float64 = float64(((((3.14159 / -2.25) + 3.14159) - (0.1 + 1e10)) + 1e10))
    _ = v8
    return 32767
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    var v4 uint8 = 0
    var v9 uint64 = 2
    fmt.Println(v4, v9, v1)
    v10 := GenS1{F0: 42}
    for fi11 := 0; fi11 < 9; fi11++ {
        v9 = (v9 * v9)
        _ = fi11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v1 += (v1 << ((18446744073709551615) & 63))
        _ = fi12
    }
    var v13 float32 = (3.14159 - (0.1 * 0.0))
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint8 = (^(v4))
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := []int{(-(fi14)), 128, 9223372036854775807}
            _ = fi16
            _ = v17
        }
        var v18 uint8 = v4
        _ = fi14
        _ = v15
        _ = v18
    }
    v4 = v4
    switch "" {
    case "foo bar":
        for fi19 := 0; fi19 < 3; fi19++ {
            if (true && true) {
                switch (fi19 - (fi19 >> ((uint(v4)) & 63))) {
                case -895:
                    delete(v2, 0)
                case -1:
                    var v20 string = "\n\t"
                    _ = v20
                case -2147483648:
                    v21 := []int{fi19, fi19}
                    _ = v21
                }
            } else {
                switch fi19 {
                case -691:
                    v2[1] = ("foo bar" <= (((("foo bar" + "foo bar") + "a") + "a") + "hello"))
                case -929:
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 int8 = int8(fi19)
                        var v24 float64 = ((-2.25 - 0.0) * -2.25)
                        if false {
                            switch "" {
                            case "hello":
                                v25 := []int{(-(fi22)), -2147483648, (fi19 << ((1) & 63))}
                                _ = v25
                            case "":
                                fmt.Println(v13, v24)
                            }
                            var v26 int = fi22
                            var v27 uint64 = v9
                            _ = v26
                            _ = v27
                        } else {
                            v28 := []int{(^(fi22)), fi19}
                            var v29 uint = 1
                            v30 := GenS1{F0: 29}
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            }
            if (0.1 != -2.25) {
                var v31 uint8 = (v4 >> ((0) & 63))
                if ((("foo bar" + "\n\t") != "\n\t") && ((true && (-128 > v1)) || (true && true))) {
                    v13 += v13
                }
                _ = v31
            }
            v4 += v4
            _ = fi19
        }
    default:
        var v32 float64 = (3.14159 / float64(float32(-2.25)))
        _ = v32
    }
    v9 -= v9
    var v33 int = 9223372036854775807
    var v34 int = v33
    switch (v34 + (v34 & v33)) {
    case 0:
        for fi35 := 0; fi35 < 9; fi35++ {
            v4 = (v4 ^ 255)
            _ = len(v2)
            _ = fi35
        }
    case 32768:
        v1 = v1
    default:
        var v36 bool = (!(("hello" + "") != ""))
        _ = v36
    }
    v37 := GenS2{F0: ("foo bar" + ("a" + "")), F1: 48, F2: 255}
    var v38 uint16 = 65535
    v39 := GenS1{F0: (v4 & v4)}
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v9
}
