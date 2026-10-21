package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint16) string {
    p0 = (p0 << ((4294967295) & 63))
    v5 := []int{-2147483648}
    v6 := GenS1{F0: 42, F1: 42}
    switch 1 {
    case 1:
        if false {
            var v7 uint32 = 2147483648
            _ = len(v5)
            _ = v7
        }
    default:
        p0 += (p0 & 255)
    }
    _ = p0
    _ = v5
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -793
    v3 := make(map[uint32]int)
    var v4 bool = true
    var v8 uint8 = 255
    v9 := v1[704]
    delete(v1, 0)
    var v10 uint16 = 255
    var v11 int = 32767
    fmt.Println(v9)
    switch "a" {
    case "foo bar":
        v12 := []int{(v11 - (v11 & (v11 * 326)))}
        _ = v12
    default:
        var v13 uint8 = 1
        _ = v13
    }
    if ((v11 < v11) || true) {
        v11 = (^(v11))
        if (((!false) || v4) || (false && (829 <= 1000))) {
            v10 += v10
            fmt.Println(v2, v11)
        }
        v14 := []int{}
        _ = v14
    }
    if v9 {
        switch (v11 / ((v11) | 1)) {
        case -622:
            if (true || true) {
                v15 := GenS1{F0: -1, F1: 2147483647}
                _ = v15
            } else {
                v16 := []int{(v11 % ((int(v2)) | 1)), (v11 >> ((0) & 63)), (v11 * 9223372036854775807)}
                var v17 uint64 = 42
                switch v11 {
                case -9223372036854775808:
                    delete(v1, 0)
                default:
                    if true {
                        v17 -= v17
                        fmt.Println(v17, v2)
                        var v18 bool = ((90 < -1) && ("" <= "foo bar"))
                        _ = v18
                    }
                }
                _ = v16
                _ = v17
            }
        case -1:
            fmt.Println(v9, v9)
        default:
            v9 = v9
        }
    } else {
        v19 := GenS1{F0: (^(v2)), F1: (^(v2))}
        if (v8 == 0) {
            var v20 float32 = -2.25
            v20 = 0.1
            _ = v20
        } else {
            var v21 int32 = v2
            fmt.Println(v8)
            switch (("hello" + ("a" + "")) + "foo bar") {
            case "":
                for fi22 := 0; fi22 < 6; fi22++ {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        v21 -= 0
                        var v24 string = ""
                        if v9 {
                            var v25 uint = 2
                            _ = v25
                        }
                        _ = fi23
                        _ = v24
                    }
                    fmt.Println(v4, v4)
                    v26 := []int{fi22, 32767, v11, (v11 | (-(fi22)))}
                    _ = fi22
                    _ = v26
                }
            case "a":
                var v27 bool = v9
                _ = v27
            default:
                v28 := GenS2{F0: (v10 ^ v10)}
                _ = v28
            }
            _ = v21
        }
        _ = v19
    }
    v29 := GenS2{F0: (v10 << ((4294967295) & 63))}
    var v30 int8 = 1
    switch v11 {
    case 36:
        for fi31 := 0; fi31 < 1; fi31++ {
            v32 := v1[4294967295]
            _ = fi31
            _ = v32
        }
    case 32767:
        v33 := GenS2{F0: 65535}
        _ = v33
    case 30:
        v34 := []int{-2147483648, v11, v11}
        _ = v34
    default:
        v2 += int32(v30)
    }
    var v35 int = v11
    v36 := []int{(v35 ^ v35), (^(v11)), (v35 + -874)}
    fmt.Println(v30)
    var v37 float32 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v8
    _ = v9
}
