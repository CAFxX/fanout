package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0() uint64 {
    var v7 bool = false
    v8 := GenS1{F0: 1, F1: 32767, F2: 191}
    v9 := []int{-2147483648, 127}
    v7 = (3.14159 > (1.5 - -2.25))
    _ = v7
    _ = v8
    _ = v9
    return 182
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 uint16 = 42
    var v4 uint32 = 516
    var v5 uint8 = 2
    var v6 uint16 = v3
    fmt.Println(v6, v4, v3)
    if false {
        if true {
            v10 := GenS1{F0: 442, F1: 1, F2: (v5 | v5)}
            if (42 > 1) {
                delete(v2, 65535)
            }
            _ = v10
        }
        if false {
            v5 += (v5 ^ (v5 + v5))
            v6 = v3
        }
    }
    v11 := []int{128}
    for fi12 := 0; fi12 < 12; fi12++ {
        if (!(!(3 == 359))) {
            var v13 int8 = 127
            for fi14 := 0; fi14 < 1; fi14++ {
                if true {
                    if true {
                        switch ("hello" + "hello") {
                        case "a":
                            v6 = v6
                        case "\n\t":
                            var v15 float32 = 1.5
                            _ = v15
                        }
                    }
                    switch "a" {
                    case "\n\t":
                        var v16 int32 = 42
                        _ = v16
                    case "":
                        if ((-(fi12)) < (fi12 >> ((2) & 63))) {
                            if true {
                                if true {
                                    var v17 string = "hello"
                                    fmt.Println(v17, v13)
                                    _ = v17
                                } else {
                                    v4 -= 20
                                    v18 := GenS1{F0: uint64(v4), F1: 980, F2: (v5 + 69)}
                                    v6 = v6
                                    _ = v18
                                }
                                delete(v1, -1)
                            }
                            _ = len(v2)
                            var v19 uint8 = 1
                            _ = v19
                        } else {
                            var v20 int32 = int32(v5)
                            v6 -= v3
                            v4 -= v4
                            _ = v20
                        }
                        v13 = v13
                    case "foo bar":
                        fmt.Println(v5)
                    }
                }
                v3 = (v6 | (v3 & uint16(v4)))
                _ = fi14
            }
            fmt.Println(fi12, v6, v13)
            _ = v13
        } else {
            var v21 string = "hello"
            _ = v21
        }
        var v22 int = 4294967295
        _ = fi12
        _ = v22
    }
    if false {
        v5 = v5
        if false {
            switch "foo bar" {
            case "hello":
                var v23 int32 = 349
                _ = v23
            case "\n\t":
                v24 := GenS1{F0: genFn0(), F1: -820, F2: (v5 << ((93) & 63))}
                _ = v24
            case "":
                v1[990] = ("\n\t" + "a")
            }
            var v25 uint32 = (v4 >> ((1) & 63))
            _ = v25
        }
    } else {
        v11 = append(v11, -1)
        fmt.Println(v6, v5, v6)
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 float32 = (1e10 * 3.14159)
        var v28 uint8 = (^(v5))
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint = 503
    var v30 uint64 = genFn0()
    var v31 uint16 = (v3 % ((v3) | 1))
    v30 = v30
    var v32 string = ((("foo bar" + "\n\t") + "a") + "a")
    var v33 uint32 = 633
    v34 := GenS1{F0: v30, F1: -32768, F2: v5}
    if (484 <= -1) {
        v35 := v2[(v4 << ((243) & 63))]
        _ = v35
    } else {
        fmt.Println(v33, v4, v5)
        var v36 uint64 = (v30 ^ (^(v30)))
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
