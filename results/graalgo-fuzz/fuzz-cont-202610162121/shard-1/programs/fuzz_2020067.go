package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint; F1 uint16; F2 string }
type GenS3 struct { F0 int8; F1 uint }
func genFn0(p0 int16, p1 int16) uint {
    v5 := []int{-41, 578}
    p1 += 1
    var v6 uint = 2
    if len(v5) > 0 {
        v5[0] = 128
    }
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint16 = 255
        v9 := GenS1{F0: "\n\t", F1: false}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := []uint32{65535, 2147483647, 1, 1}
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    v4 := []string{"hello", ("hello" + ("hello" + (((("hello" + "a") + "") + "foo bar") + "a"))), ("" + ""), "\n\t"}
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint16 = 65535
        v12 := v2["foo bar"]
        v2["a"] = -2147483648
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 string = (("foo bar" + (("\n\t" + ("foo bar" + "foo bar")) + ("hello" + ("hello" + "a")))) + "foo bar")
    v13 = v13
    fmt.Println(v13)
    if true {
        _ = len(v4)
    }
    if false {
        for fi14 := 0; fi14 < 15; fi14++ {
            switch v13 {
            case "":
                v13 = ((v13 + (("" + "hello") + "a")) + ((v13 + (v13 + (v13 + v13))) + "a"))
            case "a":
                if true {
                    v3[fi14] = (fi14 / ((128) | 1))
                    for fi15 := 0; fi15 < 18; fi15++ {
                        v16 := GenS1{F0: v13, F1: false}
                        _ = len(v2)
                        v1 = append(v1, 2147483648)
                        _ = fi15
                        _ = v16
                    }
                }
            }
            v17 := GenS1{F0: ((v13 + v13) + ((v13 + v13) + "hello")), F1: ((!true) || false)}
            _ = fi14
            _ = v17
        }
    }
    var v18 int64 = 9223372036854775807
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := v2[v13]
            var v22 float32 = (0.1 * 0.0)
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        if true {
            var v24 bool = ((((-266 < fi23) || false) || (!true)) && (!(true && true)))
            var v25 int = fi23
            var v26 bool = true
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi23
    }
    var v27 uint8 = 2
    fmt.Println(v18)
    if true {
        var v28 uint16 = 837
        for fi29 := 0; fi29 < 8; fi29++ {
            fmt.Println(v13, v18, v13)
            v30 := GenS2{F0: genFn0(903, 0), F1: 2, F2: v13}
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 13; fi31++ {
            if len(v4) > 0 {
                v4[0] = ((v13 + "foo bar") + v13)
            }
            v32 := GenS3{F0: 42, F1: 65535}
            _ = fi31
            _ = v32
        }
        _ = v28
    }
    var v33 uint64 = 42
    for fi34 := 0; fi34 < 15; fi34++ {
        if false {
            for fi35 := 0; fi35 < 10; fi35++ {
                var v36 int = 1
                _ = fi35
                _ = v36
            }
            fmt.Println(v18, fi34, v13)
            fmt.Println(fi34, v33)
        } else {
            v4 = append(v4, (v13 + v13))
            for fi37 := 0; fi37 < 8; fi37++ {
                v4 = append(v4, v13)
                if false {
                    fmt.Println(v27, v13)
                    var v38 int64 = v18
                    _ = v38
                }
                var v39 int16 = 42
                _ = fi37
                _ = v39
            }
        }
        v33 -= (v33 * (v33 | v33))
        fmt.Println(v27, v33, v33)
        _ = fi34
    }
    _ = len(v4)
    var v40 bool = true
    _ = len(v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v40
}
