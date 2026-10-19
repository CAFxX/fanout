package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 bool }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0(p0 int8) int16 {
    if false {
        v5 := GenS1{F0: "", F1: 42, F2: ("a" < "foo bar")}
        _ = v5
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int = (fi6 | fi6)
        var v8 uint = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{42, 319}
    var v10 string = ("a" + "a")
    v11 := GenS2{F0: p0, F1: p0, F2: 1}
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 0.0
    v3 := []string{"\n\t", "hello", "", "\n\t"}
    var v4 int64 = 464
    if len(v3) > 0 {
        v3[0] = ("a" + "foo bar")
    }
    if ((!(-2.25 == float32((-2.25 * 1.5)))) && (!(4294967295 > 616))) {
        var v12 bool = (1e10 > -2.25)
        v13 := v1[""]
        fmt.Println(v12)
        _ = v12
        _ = v13
    } else {
        if true {
            if (42 == 26) {
                _ = len(v1)
                var v14 uint16 = 1
                _ = v14
            }
            v15 := GenS2{F0: 3, F1: 1, F2: 65535}
            var v16 uint32 = 126
            _ = v15
            _ = v16
        }
    }
    var v17 float32 = v2
    v18 := GenS2{F0: 90, F1: 0, F2: 0}
    var v19 bool = (true || (true && false))
    if (!(!v19)) {
        switch 956 {
        case -331:
            v20 := GenS3{F0: 0, F1: 2147483648, F2: v4}
            _ = v20
        }
        fmt.Println(v2, v4, v2)
        switch -2147483648 {
        case -954:
            var v21 uint32 = 2147483648
            _ = v21
        }
    } else {
        if v19 {
            var v22 int32 = -2147483648
            _ = v22
        }
    }
    if (((0.0 - 0.1) - 1e10) < -2.25) {
        if true {
            v1[(("hello" + "") + "")] = "hello"
            fmt.Println(v2, v17, v19)
        } else {
            fmt.Println(v19, v2)
            if false {
                if true {
                    v23 := GenS2{F0: 42, F1: 1, F2: 0}
                    if v19 {
                        delete(v1, "a")
                    } else {
                        v17 += v2
                        var v24 int = 4294967295
                        _ = v24
                    }
                    _ = v23
                }
            } else {
                fmt.Println(v19, v17, v2)
            }
            v17 = v17
        }
        v17 = (((v2 + v2) - v17) / (v2 - v2))
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint16 = 1
            if v19 {
                for fi27 := 0; fi27 < 9; fi27++ {
                    switch 0 {
                    case 32768:
                        v28 := GenS4{F0: 205, F1: 2147483647, F2: 2}
                        _ = v28
                    case 513:
                        _ = len(v1)
                    default:
                        if v19 {
                            fmt.Println(v26, v19)
                        }
                    }
                    v17 = v2
                    _ = fi27
                }
                var v29 int64 = (v4 % ((-9223372036854775808) | 1))
                if v19 {
                    var v30 uint32 = 0
                    fmt.Println(v4)
                    _ = v30
                }
                _ = v29
            }
            v19 = v19
            _ = fi25
            _ = v26
        }
    } else {
        fmt.Println(v4, v2, v2)
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        v2 += v2
        fmt.Println(fi31, v19, v19)
        v4 += v4
        _ = fi31
    }
    fmt.Println(v17, v17)
    v1[(("a" + ("foo bar" + "a")) + "\n\t")] = ("" + "foo bar")
    var v32 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
