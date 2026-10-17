package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint }
func genFn0() int16 {
    v5 := []int{32767}
    var v6 uint16 = 255
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := []int{-9223372036854775808, 1, -9223372036854775808, -133}
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    var v4 uint8 = 2
    var v7 int64 = 9223372036854775807
    v3[0] = 127
    for fi8 := 0; fi8 < 10; fi8++ {
        if (128 == (v4 >> ((891) & 63))) {
            delete(v2, "a")
        }
        _ = fi8
    }
    if true {
        var v9 float32 = 3.14159
        fmt.Println(v7)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 9; fi11++ {
                if ((!false) || (0 != 2)) {
                    v12 := v2["a"]
                    _ = v12
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if false {
                    var v13 int = (fi11 + (fi10 ^ 4294967295))
                    v4 = (^(v4))
                    _ = v13
                } else {
                    v14 := GenS1{F0: false, F1: 1, F2: 1}
                    var v15 string = ("hello" + "")
                    fmt.Println(fi11)
                    _ = v14
                    _ = v15
                }
                switch (^(fi11)) {
                case 32767:
                    if ((!false) && (3.14159 == (0.0 - (float64(0.0) / float64(1.5))))) {
                        v3[1] = fi11
                        delete(v2, "")
                        var v16 float64 = 1.5
                        _ = v16
                    }
                }
                _ = fi11
            }
            _ = fi10
        }
    }
    v7 = 9223372036854775807
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(fi17, fi17, v4)
        _ = fi17
    }
    v18 := v2["hello"]
    v7 += (v7 ^ v7)
    var v19 int32 = 0
    if true {
        v20 := GenS1{F0: (-128 >= 100), F1: 2, F2: 42}
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            switch v18 {
            case "\n\t":
                var v23 int8 = 127
                _ = v23
            case "":
                var v24 uint16 = 83
                _ = v24
            }
            _ = fi22
        }
        var v25 uint64 = 4294967295
        var v26 uint32 = 4294967295
        _ = fi21
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: (!true), F1: 42, F2: 1}
    v18 = v18
    var v28 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v7
}
