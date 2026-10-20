package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 bool; F1 string; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint16; F2 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = ((("\n\t" + (("foo bar" + "\n\t") + ("" + ("hello" + "foo bar")))) + "foo bar") == ("a" + ("" + ("a" + "foo bar"))))
    var v3 float32 = float32((1e10 * 1.5))
    var v4 int64 = -1
    var v5 string = "hello"
    var v6 int = 541
    var v7 int8 = -1
    var v8 int64 = 42
    v9 := GenS1{F0: 1, F1: "a"}
    var v10 int16 = 1000
    fmt.Println(v4, v2, v2)
    if false {
        v2 = v2
        v11 := GenS2{F0: 255, F1: 2, F2: 340}
        _ = v11
    }
    _ = len(v1)
    var v12 int64 = (v4 << ((65535) & 63))
    var v13 uint64 = 0
    v14 := GenS2{F0: 1, F1: 42, F2: 2}
    v7 = v7
    fmt.Println(v2, v7, v8)
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v5)
        if v2 {
            switch v5 {
            case "foo bar":
                delete(v1, ((v5 + (v5 + v5)) + v5))
            default:
                fmt.Println(v6)
            }
        } else {
            fmt.Println(v3, fi15)
            v10 = 42
            var v16 uint32 = 4294967295
            _ = v16
        }
        v6 -= (fi15 & -2147483648)
        _ = fi15
    }
    if true {
        var v17 int = v6
        v18 := GenS2{F0: 42, F1: 0, F2: 1}
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := GenS3{F0: -266, F1: -2147483648}
            v5 = v5
            if v2 {
                switch -376 {
                case 205:
                    if (4294967295 != 2) {
                        if v2 {
                            for fi21 := 0; fi21 < 6; fi21++ {
                                v5 = ""
                                fmt.Println(v10, v6)
                                v22 := GenS2{F0: 91, F1: 255, F2: 65535}
                                _ = fi21
                                _ = v22
                            }
                        } else {
                            var v23 int32 = 366
                            _ = v23
                        }
                    }
                }
                if (!(v5 >= ((v5 + (v5 + "foo bar")) + (v5 + "hello")))) {
                    v6 = 32768
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v25 := GenS4{F0: (1 <= 789), F1: ((v5 + v5) + v5), F2: (fi19 * 852)}
                        var v26 int32 = 1
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 1; fi27++ {
                        v28 := GenS3{F0: v10, F1: -128}
                        v29 := GenS5{F0: v10}
                        delete(v1, v5)
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                } else {
                    v30 := []int{(fi19 / (((v17 | (v17 << ((390) & 63)))) | 1))}
                    _ = v30
                }
                var v31 int64 = v4
                _ = v31
            } else {
                for fi32 := 0; fi32 < 1; fi32++ {
                    fmt.Println(v8, v3)
                    _ = fi32
                }
                switch ("foo bar" + v5) {
                case "\n\t":
                    var v33 int16 = v10
                    _ = v33
                default:
                    v13 = v13
                }
            }
            _ = fi19
            _ = v20
        }
        _ = v17
        _ = v18
    }
    if (v13 <= v13) {
        var v34 bool = ((!(!((v13 / (((^(v13))) | 1)) != v13))) && true)
        var v35 uint = 2
        var v36 int8 = (v7 / ((0) | 1))
        _ = v34
        _ = v35
        _ = v36
    } else {
        if v2 {
            switch v6 {
            case 127:
                v37 := GenS6{F0: 1, F1: 255, F2: v5}
                _ = v37
            case -2147483648:
                var v38 int8 = (^(v7))
                _ = v38
            }
        }
    }
    v39 := v1[("hello" + v5)]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
