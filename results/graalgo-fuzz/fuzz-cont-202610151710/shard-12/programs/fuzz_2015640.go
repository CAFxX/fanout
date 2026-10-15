package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 int32) uint8 {
    v6 := GenS1{F0: 956}
    if false {
        p0 = p0
        if (128 < -9223372036854775808) {
            var v7 int64 = 1
            _ = v7
        }
    }
    fmt.Println(p0)
    fmt.Println(p0)
    v8 := GenS1{F0: 968}
    _ = p0
    _ = v6
    _ = v8
    return 128
}
func genFn1() int64 {
    var v9 int32 = 2147483647
    v9 -= -625
    if false {
        v9 = v9
        var v10 int32 = 2147483647
        v11 := GenS1{F0: 4294967295}
        _ = v10
        _ = v11
    }
    v9 += v9
    fmt.Println(v9)
    _ = v9
    return 881
}

func main() {
    var v1 string = ((("" + ("" + ("" + (((("" + "foo bar") + ((((("\n\t" + "hello") + "hello") + "\n\t") + "\n\t") + "a")) + "a") + ("foo bar" + ""))))) + ("hello" + "foo bar")) + "a")
    var v2 int8 = -45
    v3 := make(map[string]int)
    var v4 float32 = 1.5
    var v5 uint = 18446744073709551615
    v2 -= 1
    if true {
        v5 = v5
    }
    var v12 uint64 = 708
    v12 = 18446744073709551615
    switch -451 {
    case -1:
        switch 32768 {
        case 4294967295:
            if ((!(-375 == 1)) || (4294967295 >= 2147483647)) {
                fmt.Println(v5)
            } else {
                var v13 uint32 = 651
                for fi14 := 0; fi14 < 13; fi14++ {
                    var v15 uint32 = (v13 & (^(v13)))
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
        case -1:
            if ((false && true) && (!false)) {
                var v16 string = v1
                _ = v16
            }
        default:
            v17 := GenS1{F0: 2147483648}
            _ = v17
        }
    default:
        if (int32(v2) == -100) {
            v5 = (v5 - 18446744073709551615)
            if false {
                for fi18 := 0; fi18 < 16; fi18++ {
                    switch (-(fi18)) {
                    case 9223372036854775807:
                        if true {
                            switch ((((v1 + v1) + "\n\t") + v1) + v1) {
                            case "hello":
                                switch (v1 + v1) {
                                case "\n\t":
                                    var v19 string = (v1 + (v1 + (v1 + "a")))
                                    _ = v19
                                case "":
                                    var v20 uint32 = 2
                                    _ = v20
                                case "a":
                                    var v21 int16 = 659
                                    _ = v21
                                }
                            default:
                                fmt.Println(fi18)
                            }
                        } else {
                            fmt.Println(fi18)
                        }
                    case 1:
                        for fi22 := 0; fi22 < 5; fi22++ {
                            var v23 bool = true
                            delete(v3, v1)
                            _ = fi22
                            _ = v23
                        }
                    default:
                        v24 := GenS1{F0: 842}
                        _ = v24
                    }
                    v2 = v2
                    v25 := []int{}
                    _ = fi18
                    _ = v25
                }
                var v26 int = 4294967295
                fmt.Println(v12, v26)
                _ = v26
            }
            if false {
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 int8 = (v2 | (v2 - v2))
                    _ = fi27
                    _ = v28
                }
                v29 := GenS2{F0: (v12 % ((18446744073709551615) | 1)), F1: (v1 + "foo bar")}
                if false {
                    if false {
                        v30 := GenS2{F0: v12, F1: ""}
                        v31 := []int{451, 32768}
                        var v32 uint32 = 94
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    v2 -= (v2 >> (((^(v5))) & 63))
                }
                _ = v29
            }
        }
    }
    var v33 bool = (true || (42 >= -1))
    var v34 int = 32768
    var v35 int16 = -1
    for fi36 := 0; fi36 < 3; fi36++ {
        fmt.Println(v5, v1, v33)
        v2 += v2
        _ = fi36
    }
    if (42 > 1) {
        var v37 int64 = 1
        var v38 uint64 = v12
        _ = v37
        _ = v38
    }
    fmt.Println(v2, v34, v34)
    fmt.Println(v12, v33, v33)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
