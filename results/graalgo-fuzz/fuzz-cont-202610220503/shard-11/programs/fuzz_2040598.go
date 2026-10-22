package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint32; F1 bool; F2 bool }
type GenS3 struct { F0 int64; F1 uint }
func genFn0(p0 int64, p1 bool) bool {
    v4 := GenS1{F0: ((408 != 42) && true), F1: -1}
    var v5 int16 = 0
    var v6 int = 751
    var v7 uint = 65535
    v6 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return false
}
func genFn1() uint8 {
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 string = ("a" + "\n\t")
        v9 = v9
        v10 := GenS1{F0: (-93 != -1), F1: 2147483647}
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(fi11)
            var v12 string = "\n\t"
            for fi13 := 0; fi13 < 3; fi13++ {
                if ((!true) || true) {
                    v12 = (v12 + (v12 + ""))
                    v14 := []int{fi13, fi11, -861, fi11}
                    v12 = (v12 + v12)
                    if (42 < 1000) {
                        if len(v14) > 0 {
                            v14[0] = -649
                        }
                        v15 := GenS2{F0: 1, F1: (760 <= 2), F2: true}
                        switch fi11 {
                        case 9223372036854775807:
                            switch ((v12 + (v12 + ("hello" + (v12 + (v12 + ((v12 + ((v12 + "\n\t") + "a")) + v12)))))) + "foo bar") {
                            case "":
                                fmt.Println(v12, v12, fi13)
                            case "a":
                                v16 := GenS1{F0: (42 >= -126), F1: 0}
                                _ = v16
                            case "hello":
                                v17 := GenS2{F0: 0, F1: (!(!true)), F2: false}
                                _ = v17
                            }
                        case 42:
                            var v18 uint16 = 42
                            _ = v18
                        case -2147483648:
                            switch (v12 + ((v12 + v12) + ("hello" + ((v12 + v12) + v12)))) {
                            case "\n\t":
                                v12 = "a"
                            case "a":
                                fmt.Println(fi11)
                            }
                        default:
                            if false {
                                v12 = v12
                            }
                        }
                        _ = v15
                    }
                    if ((1 >= 1) || false) {
                        v14 = append(v14, -649)
                        v19 := GenS2{F0: 204, F1: false, F2: false}
                        _ = v19
                    } else {
                        v20 := GenS2{F0: 1, F1: false, F2: (42 < 1)}
                        v21 := GenS1{F0: false, F1: 1}
                        _ = v20
                        _ = v21
                    }
                    _ = v14
                }
                var v22 uint32 = 340
                _ = fi13
                _ = v22
            }
            _ = fi11
            _ = v12
        }
        var v23 float32 = 1e10
        fmt.Println(v23)
        _ = v23
    } else {
        v24 := GenS3{F0: 42, F1: 42}
        _ = v24
    }
    v25 := GenS3{F0: 1, F1: 185}
    var v26 int8 = 42
    v26 -= (v26 - -59)
    _ = v25
    _ = v26
    return 136
}

func main() {
    v1 := []string{("hello" + ("a" + ("a" + ("a" + ("a" + ("a" + "\n\t")))))), "hello", ("foo bar" + "a")}
    var v2 uint8 = 210
    v3 := make(map[uint32]bool)
    v2 += v2
    v3[824] = true
    var v27 uint32 = 2
    v27 += v27
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := GenS3{F0: -1, F1: 42}
        _ = fi28
        _ = v29
    }
    var v30 uint8 = 42
    var v31 uint32 = v27
    v31 -= (v31 / ((v31) | 1))
    v27 -= (v31 | v31)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
}
