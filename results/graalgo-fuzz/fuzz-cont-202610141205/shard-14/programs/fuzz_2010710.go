package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int32; F2 int }
func genFn0() bool {
    v7 := []int{-1, -2147483648, -2147483648}
    for fi8 := 0; fi8 < 11; fi8++ {
        if (1 < 0) {
            v9 := GenS1{F0: (("hello" + ("foo bar" + "\n\t")) + ("hello" + "a"))}
            fmt.Println(fi8)
            _ = v9
        }
        var v10 int32 = 42
        _ = fi8
        _ = v10
    }
    _ = v7
    return ((0 != 1) && false)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "foo bar"
    v3 := []int{}
    v4 := []string{((v2 + (("hello" + (v2 + ("hello" + v2))) + (v2 + ""))) + "foo bar")}
    var v5 bool = true
    var v6 int8 = 1
    var v11 float32 = 0.1
    if false {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int = -2147483648
            _ = fi12
            _ = v13
        }
        var v14 uint = 45
        v6 = (v6 / ((v6) | 1))
        _ = v14
    } else {
        if (!v5) {
            fmt.Println(v2, v2)
        }
        fmt.Println(v6, v5)
    }
    if v5 {
        if len(v3) > 0 {
            v3[0] = 859
        }
    }
    fmt.Println(v11, v2, v5)
    v1[504] = 1
    v15 := GenS2{F0: 65535}
    if v5 {
        v11 -= 1.5
    } else {
        v16 := GenS2{F0: 916}
        switch v2 {
        case "foo bar":
            var v17 int8 = 100
            _ = v17
        case "\n\t":
            var v18 uint32 = 2147483648
            _ = v18
        }
        _ = v16
    }
    var v19 string = v2
    for fi20 := 0; fi20 < 1; fi20++ {
        if true {
            for fi21 := 0; fi21 < 2; fi21++ {
                v22 := GenS1{F0: v2}
                _ = fi21
                _ = v22
            }
            if ((3.14159 / 1e10) != (0.0 / (1.5 + 0.0))) {
                if v5 {
                    fmt.Println(v11, v2)
                } else {
                    for fi23 := 0; fi23 < 3; fi23++ {
                        v24 := GenS3{F0: (true && ((161 < 1) || ((0.0 - 3.14159) != 1.5))), F1: 42, F2: 32767}
                        delete(v1, 4294967295)
                        _ = fi23
                        _ = v24
                    }
                }
                v25 := GenS4{F0: 2, F1: int32(v6), F2: (^(fi20))}
                _ = v25
            } else {
                v2 = (v19 + "hello")
                if (!(!v5)) {
                    v26 := GenS3{F0: false, F1: 4294967295, F2: 535}
                    _ = v26
                } else {
                    var v27 uint8 = 127
                    _ = v27
                }
                var v28 uint32 = 42
                _ = v28
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v29 bool = (0 >= -32768)
        _ = fi20
        _ = v29
    }
    var v30 float32 = (float32((1e10 / (-2.25 * 0.1))) * v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
