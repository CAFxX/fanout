package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32 }
func genFn0() int {
    var v6 bool = true
    var v7 int32 = 218
    switch -161 {
    case -2147483648:
        var v8 int64 = -891
        _ = v8
    case -726:
        v9 := GenS1{F0: false, F1: 109}
        _ = v9
    case 127:
        switch (((("a" + "\n\t") + ("\n\t" + "")) + "hello") + ("\n\t" + "")) {
        case "hello":
            v10 := GenS1{F0: true, F1: 127}
            _ = v10
        default:
            switch -2147483648 {
            case 1:
                fmt.Println(v6)
            case 32768:
                if (true || (128 != 42)) {
                    v11 := []int{42, -988, -2147483648}
                    _ = v11
                }
            case 128:
                var v12 uint = 0
                _ = v12
            default:
                var v13 uint64 = 0
                _ = v13
            }
        }
    default:
        var v14 uint8 = 255
        _ = v14
    }
    var v15 int8 = 127
    if true {
        v16 := []int{-1, 32768}
        _ = v16
    }
    _ = v15
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint = 740
    v2 := []uint32{2147483648}
    var v3 float32 = -2.25
    v4 := make(map[uint32]int)
    var v5 uint64 = 1
    var v17 int16 = 0
    var v18 uint8 = 42
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS1{F0: false, F1: 255}
            fmt.Println(v5)
            _ = fi20
            _ = v21
        }
        var v22 bool = (false || true)
        v1 = 1
        _ = fi19
        _ = v22
    }
    v23 := GenS2{F0: 42}
    if true {
        var v24 int16 = -32768
        v25 := GenS2{F0: 999}
        fmt.Println(v17, v5)
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: (0 >= 1), F1: (^(v18))}
    v17 = -460
    var v27 bool = (65535 > 499)
    var v28 float32 = v3
    if len(v2) > 0 {
        v2[0] = 65535
    }
    var v29 int8 = 100
    v4[42] = genFn0()
    fmt.Println(v1, v3, v28)
    var v30 string = ("foo bar" + (("a" + "a") + "\n\t"))
    if false {
        if true {
            if (!(uint64(v1) == (v5 & v5))) {
                fmt.Println(v3)
            }
            for fi31 := 0; fi31 < 8; fi31++ {
                var v32 int64 = 114
                _ = fi31
                _ = v32
            }
        } else {
            if v27 {
                var v33 uint8 = (v18 / ((v18) | 1))
                var v34 int64 = 0
                switch "" {
                case "":
                    _ = len(v4)
                case "hello":
                    fmt.Println(v5, v30, v3)
                case "foo bar":
                    v34 = (^(v34))
                }
                _ = v33
                _ = v34
            } else {
                if (!(true || v27)) {
                    for fi35 := 0; fi35 < 11; fi35++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi35
                    }
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    if ((v27 && ((v5 <= (v5 | (v5 | (v5 << ((uint(v18)) & 63))))) || (v17 > (v17 & (^(v17)))))) || (false && (!(!(!((!((v18 >> ((v1) & 63)) >= v18)) && v27)))))) {
                        var v36 string = (v30 + v30)
                        for fi37 := 0; fi37 < 7; fi37++ {
                            _ = len(v2)
                            v38 := GenS1{F0: v27, F1: (^(v18))}
                            var v39 string = ""
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        v18 -= 20
                        _ = v36
                    } else {
                        var v40 int8 = 42
                        _ = v40
                    }
                } else {
                    var v41 int32 = 439
                    var v42 uint = (v1 | 856)
                    _ = v41
                    _ = v42
                }
                v4[uint32(v5)] = -1
                v17 = v17
            }
        }
    }
    fmt.Println(v27, v17, v17)
    if (true && false) {
        v2 = append(v2, 860)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
