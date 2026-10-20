package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int64; F2 string }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 string = ("hello" + ("foo bar" + "a"))
    var v4 string = ("" + (v3 + v3))
    var v5 uint16 = 65535
    for fi6 := 0; fi6 < 10; fi6++ {
        switch (("\n\t" + "\n\t") + (v3 + v3)) {
        case "a":
            var v7 int8 = 0
            _ = v7
        default:
            if true {
                var v8 float64 = 1.5
                if ((true && ((!(18446744073709551615 >= 0)) || true)) || true) {
                    v8 -= v8
                    var v9 int8 = 0
                    _ = v9
                }
                if (2147483648 >= 42) {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        switch fi10 {
                        case 0:
                            v11 := GenS1{F0: 45, F1: 2}
                            _ = v11
                        default:
                            if true {
                                v5 += v5
                                switch v3 {
                                case "\n\t":
                                    var v12 int32 = 0
                                    _ = v12
                                case "hello":
                                    v13 := []int{323, -2147483648, (fi10 >> ((781) & 63)), fi6}
                                    _ = v13
                                case "a":
                                    v8 = (-2.25 * ((v8 / 1.5) / 0.1))
                                }
                                fmt.Println(fi6, fi10)
                                v14 := GenS2{F0: 0}
                                fmt.Println(v5)
                                var v15 bool = false
                                _ = v14
                                _ = v15
                            } else {
                                v5 -= (v5 ^ (v5 >> ((255) & 63)))
                                switch "\n\t" {
                                case "hello":
                                    v16 := GenS3{F0: 2, F1: 1, F2: "hello"}
                                    _ = v16
                                }
                                v17 := GenS2{F0: 2}
                                _ = v17
                            }
                        }
                        var v18 int16 = -1
                        _ = fi10
                        _ = v18
                    }
                }
                _ = v8
            } else {
                for fi19 := 0; fi19 < 6; fi19++ {
                    fmt.Println(fi19, fi19, fi19)
                    _ = len(v2)
                    var v20 float64 = 1e10
                    _ = fi19
                    _ = v20
                }
            }
        }
        _ = fi6
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        switch (^(fi21)) {
        case -2147483648:
            var v22 int16 = -887
            _ = v22
        case 127:
            _ = len(v1)
        case -1:
            switch (v3 + (v4 + v4)) {
            case "hello":
                v23 := []int{(fi21 & (^(fi21))), (fi21 << ((1) & 63))}
                _ = v23
            case "":
                var v24 uint32 = 959
                _ = v24
            default:
                var v25 uint64 = 790
                _ = v25
            }
        }
        var v26 int32 = -964
        fmt.Println(v3, fi21, v5)
        _ = fi21
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        if false {
            v28 := GenS3{F0: 2147483648, F1: 42, F2: v4}
            v29 := GenS3{F0: 2, F1: -9223372036854775808, F2: (((((v3 + v3) + "\n\t") + (v3 + ((v4 + (v3 + ((v4 + v3) + v3))) + v4))) + ((v4 + ("" + v3)) + v3)) + (v4 + "foo bar"))}
            _ = v28
            _ = v29
        }
        fmt.Println(v3, fi27)
        for fi30 := 0; fi30 < 10; fi30++ {
            v5 -= 1
            _ = fi30
        }
        _ = fi27
    }
    v31 := GenS1{F0: 2147483647, F1: 258}
    if false {
        var v32 int = -2147483648
        _ = v32
    } else {
        v5 -= 797
    }
    v5 = (v5 + (^(v5)))
    v3 = "\n\t"
    switch -9223372036854775808 {
    case 32768:
        var v33 float64 = float64(0.1)
        _ = v33
    case 42:
        var v34 string = "hello"
        _ = v34
    default:
        v35 := v2[v3]
        _ = v35
    }
    v36 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v4
    _ = v5
}
