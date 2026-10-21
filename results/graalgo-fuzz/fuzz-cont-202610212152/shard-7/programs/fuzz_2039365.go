package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint64 }
func genFn0(p0 string, p1 uint64) int64 {
    var v6 float32 = (-2.25 - 0.0)
    var v7 uint = 4294967295
    var v8 uint64 = 4294967295
    if true {
        v9 := GenS1{F0: 1, F1: "a", F2: (p1 & v8)}
        if true {
            for fi10 := 0; fi10 < 2; fi10++ {
                if (true && false) {
                    switch fi10 {
                    case -881:
                        v11 := []int{fi10, fi10, 9223372036854775807}
                        _ = v11
                    }
                }
                var v12 int8 = 42
                switch (fi10 * fi10) {
                case 42:
                    var v13 float64 = -2.25
                    _ = v13
                case -434:
                    for fi14 := 0; fi14 < 11; fi14++ {
                        fmt.Println(v6)
                        _ = fi14
                    }
                default:
                    for fi15 := 0; fi15 < 19; fi15++ {
                        var v16 float32 = v6
                        fmt.Println(v12)
                        v7 = (v7 % ((694) | 1))
                        _ = fi15
                        _ = v16
                    }
                }
                _ = fi10
                _ = v12
            }
            var v17 uint16 = 2
            var v18 uint32 = 2147483648
            _ = v17
            _ = v18
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 string = (("hello" + "a") + "foo bar")
    v2 := []string{"a"}
    v3 := []string{v1, "foo bar"}
    var v4 bool = true
    var v5 string = (v1 + v1)
    for fi19 := 0; fi19 < 8; fi19++ {
        v2 = append(v2, "foo bar")
        _ = fi19
    }
    switch 0 {
    case -1:
        switch "a" {
        case "\n\t":
            fmt.Println(v5, v4)
        default:
            if v4 {
                v3 = append(v3, v1)
            } else {
                var v20 uint8 = 0
                var v21 uint8 = (v20 << ((1) & 63))
                _ = v20
                _ = v21
            }
        }
    }
    var v22 float64 = 3.14159
    v23 := GenS1{F0: -1, F1: ((v1 + ("hello" + v1)) + (v5 + "a")), F2: 18446744073709551615}
    if (!true) {
        v24 := GenS1{F0: genFn0(v5, 1), F1: (v5 + ((((v1 + (((v5 + v1) + ("hello" + "foo bar")) + v1)) + v5) + v1) + ((v5 + v5) + v5))), F2: 42}
        for fi25 := 0; fi25 < 19; fi25++ {
            if v4 {
                switch (fi25 | 0) {
                case 0:
                    fmt.Println(v5, v5)
                case 4294967295:
                    var v26 uint8 = 0
                    _ = v26
                case -9223372036854775808:
                    fmt.Println(v5, v5)
                default:
                    if len(v3) > 0 {
                        v3[0] = v5
                    }
                }
            } else {
                fmt.Println(v4, v22, v1)
                v4 = (!(!(v4 && v4)))
            }
            if true {
                if false {
                    v3 = append(v3, "")
                    if v4 {
                        v5 = "hello"
                        if false {
                            v22 += ((v22 - 0.1) / 0.1)
                            for fi27 := 0; fi27 < 7; fi27++ {
                                v2 = append(v2, ("\n\t" + "foo bar"))
                                _ = fi27
                            }
                        }
                    } else {
                        v1 = (("\n\t" + "hello") + "hello")
                    }
                    _ = len(v2)
                }
                var v28 float32 = 0.1
                for fi29 := 0; fi29 < 4; fi29++ {
                    v28 = (v28 + v28)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi29
                }
                _ = v28
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi25
        }
        v22 -= ((v22 + (-2.25 * float64(((0.1 * 0.0) - ((0.0 * 0.0) * float32(-2.25)))))) * (v22 + v22))
        _ = v24
    }
    v22 = 0.0
    for fi30 := 0; fi30 < 1; fi30++ {
        var v31 int16 = 42
        var v32 uint = 65535
        var v33 int32 = -1
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    if ((-739 <= -1) && v4) {
        var v34 bool = v4
        _ = v34
    }
    v5 = "\n\t"
    fmt.Println(v1, v22)
    if (2147483647 < 628) {
        v5 = (("hello" + "foo bar") + v5)
    }
    var v35 uint16 = 42
    if (0 != 1) {
        v36 := GenS1{F0: 604, F1: (("hello" + "a") + v5), F2: 115}
        _ = v36
    } else {
        v35 = 255
    }
    var v37 string = (v5 + "a")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
