package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
func genFn0() uint64 {
    var v4 uint16 = 930
    v4 += 255
    fmt.Println(v4)
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{-701, (-(fi5))}
        _ = fi5
        _ = v6
    }
    var v7 int32 = 1
    _ = v4
    _ = v7
    return 189
}

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 32767
    var v3 uint8 = 127
    var v8 int = 127
    v1 -= 810
    if (-2.25 <= (0.0 + ((-2.25 * (float32(0.1) + ((float32(3.14159) / -2.25) * 0.1))) + float32(0.0)))) {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v8, v1)
        case "hello":
            v1 = v1
        case "a":
            v9 := []int{}
            _ = v9
        }
    } else {
        v2 = (v2 >> ((1) & 63))
        var v10 uint = 1
        _ = v10
    }
    v11 := GenS1{F0: (v8 | (v8 & (v8 * v8))), F1: "foo bar"}
    if false {
        for fi12 := 0; fi12 < 3; fi12++ {
            for fi13 := 0; fi13 < 13; fi13++ {
                v1 -= (v1 << ((1) & 63))
                var v14 int32 = -2147483648
                v1 = v1
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        var v15 bool = (!(3.14159 >= (((((float32(0.1) / 0.1) - float32(float64(1e10))) - 3.14159) / (1e10 - 0.1)) / 1.0)))
        if (v3 != 18) {
            switch ("a" + ("\n\t" + "a")) {
            case "hello":
                v16 := GenS1{F0: v8, F1: "hello"}
                _ = v16
            }
            if v15 {
                var v17 float64 = (1e10 / 3.14159)
                v3 += v3
                if v15 {
                    if (2147483648 <= 4294967295) {
                        var v18 int8 = 1
                        _ = v18
                    }
                }
                _ = v17
            } else {
                switch "foo bar" {
                case "\n\t":
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v1 -= (^(v1))
                        _ = fi19
                    }
                case "hello":
                    var v20 int64 = 9223372036854775807
                    _ = v20
                default:
                    var v21 uint64 = 18446744073709551615
                    _ = v21
                }
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 uint32 = 2147483647
                    _ = fi22
                    _ = v23
                }
                v24 := GenS1{F0: (^(v8)), F1: "hello"}
                _ = v24
            }
            v25 := []int{v8, v8, (v8 >> ((4294967295) & 63)), (v8 << ((0) & 63))}
            _ = v25
        }
        _ = v15
    } else {
        var v26 uint8 = (v3 % (((v3 % ((255) | 1))) | 1))
        fmt.Println(v8, v26, v3)
        for fi27 := 0; fi27 < 19; fi27++ {
            var v28 int = (v8 | 9223372036854775807)
            var v29 int8 = int8(v28)
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = v26
    }
    fmt.Println(v8, v8)
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 string = ("a" + "a")
        v32 := []int{fi30, -215, v8}
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int32 = 42
    if false {
        v34 := []int{}
        _ = v34
    } else {
        fmt.Println(v1, v8)
    }
    v2 -= (v2 / ((0) | 1))
    switch "hello" {
    case "hello":
        var v35 int = v8
        _ = v35
    case "":
        fmt.Println(v2)
    }
    if true {
        if false {
            fmt.Println(v3)
        }
        if true {
            var v36 int8 = 33
            v33 = (v33 / (((v33 * (-(v33)))) | 1))
            v37 := GenS1{F0: (v8 % ((v8) | 1)), F1: ""}
            _ = v36
            _ = v37
        }
        for fi38 := 0; fi38 < 18; fi38++ {
            v1 += v1
            _ = fi38
        }
    } else {
        fmt.Println(v2)
        fmt.Println(v2)
    }
    v39 := GenS1{F0: 356, F1: (("\n\t" + "") + "")}
    v40 := []int{-803, 956, (v8 - (v8 + v8)), v8}
    var v41 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v40
    _ = v41
    _ = v8
}
