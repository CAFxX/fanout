package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 string; F1 int8 }
func genFn0() int64 {
    var v7 int16 = -32768
    var v8 int = 1
    _ = v7
    _ = v8
    return 0
}
func genFn1() int {
    if true {
        v9 := []int{32767, 4294967295}
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        switch 896 {
        case 2147483647:
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := []int{}
                _ = fi11
                _ = v12
            }
        case 255:
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: 32767, F1: 18446744073709551615, F2: 1}
                v15 := []int{(fi10 % ((128) | 1)), 398}
                var v16 int64 = -752
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        case 128:
            v17 := []int{(fi10 << ((1) & 63))}
            _ = v17
        }
        v18 := []int{(fi10 >> ((0) & 63)), 42, (fi10 | fi10)}
        _ = fi10
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        switch "foo bar" {
        case "\n\t":
            v20 := GenS2{F0: ("hello" + ("a" + "")), F1: 0}
            _ = v20
        }
        fmt.Println(fi19)
        v21 := []int{(^(fi19)), 32767, fi19, 960}
        _ = fi19
        _ = v21
    }
    return 0
}

func main() {
    v1 := []int{9223372036854775807, -9223372036854775808}
    var v2 uint16 = 42
    v3 := make(map[string]uint32)
    var v4 string = "a"
    v5 := make(map[int]uint32)
    var v6 float64 = 0.0
    var v22 int = 2147483647
    fmt.Println(v22, v4, v6)
    v23 := GenS2{F0: (v4 + v4), F1: 1}
    v24 := GenS2{F0: ((("foo bar" + (v4 + "hello")) + ((v4 + v4) + v4)) + (v4 + "\n\t")), F1: -1}
    for fi25 := 0; fi25 < 20; fi25++ {
        _ = len(v1)
        if false {
            fmt.Println(fi25, fi25, v22)
            var v26 bool = false
            _ = v26
        } else {
            v4 = "\n\t"
            var v27 string = ("a" + (((v4 + (v4 + "\n\t")) + "\n\t") + v4))
            _ = v27
        }
        var v28 int32 = 1
        _ = fi25
        _ = v28
    }
    var v29 string = ""
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 uint32 = 889
        _ = fi30
        _ = v31
    }
    v22 -= v22
    v6 += v6
    switch -2147483648 {
    case 42:
        v32 := GenS1{F0: 32767, F1: 1, F2: 653}
        _ = v32
    }
    var v33 int = 42
    var v34 uint = 547
    v29 = "a"
    if false {
        for fi35 := 0; fi35 < 7; fi35++ {
            v5[v22] = 4294967295
            v3[v29] = 0
            _ = fi35
        }
        if true {
            var v36 uint16 = v2
            if (-32768 >= -527) {
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
            } else {
                v1 = append(v1, (v22 - (v33 ^ -427)))
                for fi37 := 0; fi37 < 2; fi37++ {
                    var v38 int32 = -2147483648
                    v6 += (v6 - (v6 + 0.0))
                    _ = fi37
                    _ = v38
                }
                var v39 uint64 = 4294967295
                _ = v39
            }
            fmt.Println(v34, v2)
            _ = v36
        }
        v34 = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
