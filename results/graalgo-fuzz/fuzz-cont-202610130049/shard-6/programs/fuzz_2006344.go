package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint; F2 int64 }
type GenS4 struct { F0 string; F1 int16; F2 uint }
func genFn0() int32 {
    var v5 uint32 = 2147483647
    for fi6 := 0; fi6 < 13; fi6++ {
        v5 = v5
        var v7 float64 = 3.14159
        if (0.1 > ((0.0 / ((((-2.25 - 1.5) * (0.1 * ((float32(v7) - 1e10) - -2.25))) + 1.5) - float32((((1e10 * float64(1e10)) - 0.0) + ((v7 / v7) - 0.0))))) + 0.1)) {
            for fi8 := 0; fi8 < 6; fi8++ {
                switch fi8 {
                case 1:
                    v7 -= v7
                    switch "foo bar" {
                    case "":
                        var v9 float32 = 0.1
                        _ = v9
                    case "\n\t":
                        v10 := []int{fi6, (fi6 | (fi6 | 127)), fi6, 2147483647}
                        _ = v10
                    case "a":
                        v7 -= 0.0
                    }
                case 128:
                    switch "" {
                    case "":
                        v7 -= v7
                    case "\n\t":
                        v7 = float64(((3.14159 * -2.25) / 0.1))
                    }
                case 32767:
                    var v11 string = "a"
                    _ = v11
                }
                v12 := []int{(fi6 * fi6)}
                v5 -= (v5 | (v5 % (((^(v5))) | 1)))
                _ = fi8
                _ = v12
            }
            v5 = v5
            var v13 uint64 = 4294967295
            _ = v13
        }
        _ = fi6
        _ = v7
    }
    var v14 uint = 658
    v15 := GenS1{F0: 9223372036854775807, F1: "foo bar"}
    var v16 int = -882
    _ = v14
    _ = v15
    _ = v16
    _ = v5
    return 2147483647
}

func main() {
    var v1 float64 = 1.5
    var v2 bool = true
    var v3 int32 = 1
    v4 := make(map[string]int)
    if v2 {
        if v2 {
            if false {
                v17 := GenS1{F0: -836, F1: ""}
                fmt.Println(v2, v3, v1)
                _ = v17
            }
        }
        if true {
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 uint16 = 32767
                fmt.Println(v3)
                _ = fi18
                _ = v19
            }
            v1 = v1
            var v20 int16 = 32767
            _ = v20
        } else {
            v2 = v2
            v3 = 1
        }
        if v2 {
            if v2 {
                v3 -= -1
                switch -1 {
                case 2147483647:
                    var v21 int = 989
                    _ = v21
                case -794:
                    if true {
                        if ((-(v3)) != -2147483648) {
                            var v22 string = "hello"
                            switch 4294967295 {
                            case -724:
                                if false {
                                    v2 = false
                                    var v23 int64 = 9223372036854775807
                                    v2 = v2
                                    _ = v23
                                }
                            case 0:
                                v24 := GenS2{F0: 56}
                                _ = v24
                            case -646:
                                v3 = (v3 << ((2) & 63))
                            default:
                                var v25 int16 = 1000
                                _ = v25
                            }
                            _ = v22
                        } else {
                            v26 := []int{-1, 127}
                            v1 = ((v1 + v1) * float64(1.5))
                            _ = v26
                        }
                    } else {
                        var v27 bool = (!v2)
                        v1 = v1
                        var v28 int32 = 0
                        _ = v27
                        _ = v28
                    }
                case 4294967295:
                    var v29 uint32 = 65535
                    _ = v29
                default:
                    var v30 uint16 = 0
                    _ = v30
                }
            }
            var v31 string = ((("a" + ("foo bar" + (("a" + "\n\t") + "hello"))) + "") + "foo bar")
            v32 := GenS1{F0: -9223372036854775808, F1: "foo bar"}
            _ = v31
            _ = v32
        } else {
            v3 = v3
        }
    }
    v3 = v3
    v33 := GenS1{F0: 1, F1: (((((("\n\t" + "a") + "a") + "\n\t") + "\n\t") + "hello") + "hello")}
    v1 = (1.5 * v1)
    v3 += v3
    v34 := v4[""]
    for fi35 := 0; fi35 < 3; fi35++ {
        v36 := GenS3{F0: 1, F1: 223, F2: 42}
        _ = fi35
        _ = v36
    }
    if (v2 && (((3.14159 >= (0.1 * 1.5)) || (216 > 0)) || (true || false))) {
        for fi37 := 0; fi37 < 12; fi37++ {
            v38 := GenS2{F0: 1}
            _ = fi37
            _ = v38
        }
        switch "a" {
        case "\n\t":
            fmt.Println(v3, v34)
        }
    }
    v39 := []int{940, (-(v34)), (v34 ^ 4294967295)}
    var v40 uint32 = 2
    if v2 {
        fmt.Println(v1)
        fmt.Println(v1, v40)
        for fi41 := 0; fi41 < 6; fi41++ {
            v4["a"] = (v34 / (((fi41 ^ (v34 + (v34 - fi41)))) | 1))
            _ = fi41
        }
    } else {
        var v42 uint = 4294967295
        _ = v42
    }
    var v43 bool = (((0.1 - 0.0) / 1.5) > float32(float64((float32((v1 - (v1 * 3.14159))) / (1e10 - float32(float64((1.5 / (1e10 * (1.5 * (0.0 + 1e10)))))))))))
    delete(v4, ("foo bar" + ""))
    var v44 uint32 = 828
    var v45 float32 = 1.5
    v46 := GenS4{F0: "a", F1: int16(v3), F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
