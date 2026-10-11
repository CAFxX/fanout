package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 bool }
type GenS4 struct { F0 string; F1 int }
func genFn0(p0 string, p1 int) int {
    fmt.Println(p0)
    if true {
        p0 = "hello"
        p0 = ((p0 + ("hello" + "foo bar")) + p0)
    } else {
        p0 = p0
    }
    v7 := GenS1{F0: 468, F1: 42, F2: 0}
    _ = p0
    _ = p1
    _ = v7
    return (p1 | p1)
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 float64 = (1e10 * 1.5)
    var v5 int32 = 0
    var v6 int8 = -128
    var v8 int64 = 528
    v9 := GenS2{F0: 567, F1: 710}
    if (true && ((^(v6)) < -1)) {
        switch ("foo bar" + ("" + "\n\t")) {
        case "hello":
            var v10 uint16 = 0
            _ = v10
        case "\n\t":
            v11 := GenS1{F0: 0, F1: (v8 % (((v8 ^ (-(v8)))) | 1)), F2: v6}
            _ = v11
        case "":
            v12 := []int{0, 409, 0, 0}
            _ = v12
        default:
            for fi13 := 0; fi13 < 15; fi13++ {
                v8 = (v8 | v8)
                var v14 string = "hello"
                v14 = (v14 + (v14 + ""))
                _ = fi13
                _ = v14
            }
        }
    } else {
        switch -2147483648 {
        case -1:
            v8 += int64(v5)
        }
        fmt.Println(v6)
        if true {
            var v15 uint = 65535
            _ = len(v3)
            v16 := GenS2{F0: 4294967295, F1: -1}
            _ = v15
            _ = v16
        }
    }
    v6 = (v6 << ((0) & 63))
    v17 := v2["a"]
    var v18 int32 = v5
    var v19 uint8 = 127
    v20 := []int{}
    if (!false) {
        if (979 >= 42) {
            v19 = (v19 ^ v19)
        }
        v21 := GenS2{F0: 2, F1: 0}
        for fi22 := 0; fi22 < 17; fi22++ {
            v5 = -2147483648
            for fi23 := 0; fi23 < 4; fi23++ {
                switch v17 {
                case "hello":
                    var v24 float32 = 1e10
                    _ = v24
                case "":
                    v5 = (^(v18))
                default:
                    if (false || true) {
                        switch int(v6) {
                        case 1:
                            v8 = v8
                        case 32767:
                            if true {
                                v8 -= int64(fi23)
                                delete(v2, v17)
                                var v25 int8 = -12
                                _ = v25
                            } else {
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                                v26 := GenS2{F0: 2, F1: -547}
                                _ = v26
                            }
                        default:
                            v6 = -126
                        }
                        if true {
                            v6 = (v6 >> ((4294967295) & 63))
                        } else {
                            v19 += v19
                            var v27 uint16 = 0
                            _ = v27
                        }
                        var v28 uint64 = 807
                        _ = v28
                    } else {
                        v6 = v6
                        var v29 uint64 = 4294967295
                        var v30 uint32 = uint32(v19)
                        _ = v29
                        _ = v30
                    }
                }
                v31 := GenS3{F0: 835, F1: 75, F2: ((!true) && ((v6 >= v6) || (fi23 == (^(fi23)))))}
                _ = fi23
                _ = v31
            }
            _ = fi22
        }
        _ = v21
    } else {
        v32 := GenS4{F0: "a", F1: 4294967295}
        v8 = -1
        v20 = append(v20, 655)
        _ = v32
    }
    fmt.Println(v5, v17)
    switch v17 {
    case "hello":
        switch 0 {
        case 32767:
            var v33 uint16 = 37
            _ = v33
        default:
            var v34 float64 = float64(v1)
            _ = v34
        }
    case "":
        var v35 uint = 1
        _ = v35
    default:
        var v36 int = 75
        _ = v36
    }
    v17 = "foo bar"
    v8 -= v8
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
