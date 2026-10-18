package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint; F1 bool; F2 string }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 string; F1 uint16; F2 int8 }
func genFn0() int16 {
    v5 := GenS1{F0: 83, F1: 2, F2: 1}
    for fi6 := 0; fi6 < 12; fi6++ {
        switch (fi6 % ((fi6) | 1)) {
        case 884:
            v7 := []int{}
            _ = v7
        case 42:
            fmt.Println(fi6)
        case 127:
            v8 := GenS2{F0: 873, F1: false, F2: ("hello" + (("" + ("a" + "hello")) + "hello"))}
            _ = v8
        }
        if false {
            var v9 string = ("\n\t" + ("hello" + "a"))
            for fi10 := 0; fi10 < 19; fi10++ {
                v9 = (v9 + v9)
                var v11 uint32 = 2147483647
                v11 -= v11
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        fmt.Println(fi6)
        _ = fi6
    }
    var v12 uint = 1
    _ = v12
    _ = v5
    return 746
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    v3 := make(map[uint32]string)
    var v4 uint32 = 334
    fmt.Println(v2, v4)
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := []int{(fi13 / (((fi13 >> ((544) & 63))) | 1)), fi13, (-(fi13))}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            delete(v3, v4)
            v4 = (v4 | v4)
            if false {
                var v17 uint64 = 1
                v18 := GenS3{F0: -1, F1: 65535}
                _ = v17
                _ = v18
            } else {
                v2 = v2
                v19 := []int{42, fi15, (fi16 / ((fi15) | 1)), -2147483648}
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                if (v4 <= (v4 | 65535)) {
                    var v20 uint = 953
                    _ = v20
                } else {
                    v4 = 2147483647
                }
                _ = v19
            }
            _ = fi16
        }
        _ = fi15
    }
    v21 := GenS2{F0: 4294967295, F1: v2, F2: "\n\t"}
    var v22 int64 = 9223372036854775807
    v23 := GenS4{F0: 0, F1: (v4 + 42)}
    v2 = (2 > (^(v4)))
    if (!true) {
        v24 := GenS5{F0: 1, F1: ""}
        _ = v24
    }
    if false {
        var v25 int8 = 42
        v26 := GenS4{F0: 65535, F1: v4}
        for fi27 := 0; fi27 < 3; fi27++ {
            fmt.Println(v2)
            v3[(v4 + v4)] = ("\n\t" + "\n\t")
            var v28 uint64 = 42
            _ = fi27
            _ = v28
        }
        _ = v25
        _ = v26
    } else {
        var v29 int64 = v22
        v30 := GenS3{F0: 120, F1: 2}
        fmt.Println(v2, v4)
        _ = v29
        _ = v30
    }
    if ((32767 == -9223372036854775808) || v2) {
        var v31 int64 = v22
        fmt.Println(v22, v22)
        for fi32 := 0; fi32 < 10; fi32++ {
            _ = len(v1)
            fmt.Println(v31)
            _ = fi32
        }
        _ = v31
    }
    v33 := GenS1{F0: 42, F1: 147, F2: 183}
    v34 := GenS6{F0: (v22 | (-(v22))), F1: -641}
    var v35 uint = 255
    switch "hello" {
    case "\n\t":
        v22 += 497
    case "":
        if (v22 > -1) {
            fmt.Println(v2, v2)
        }
    case "hello":
        v36 := GenS7{F0: "hello", F1: 32767, F2: -1}
        _ = v36
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        for fi38 := 0; fi38 < 12; fi38++ {
            v2 = v2
            delete(v1, v4)
            _ = fi38
        }
        switch "" {
        case "foo bar":
            v35 = (v35 * 18446744073709551615)
        case "a":
            if (v2 || v2) {
                v39 := []int{fi37, 1, (fi37 >> ((v35) & 63))}
                v4 += 42
                _ = v39
            } else {
                var v40 uint8 = 2
                _ = v40
            }
        default:
            var v41 int64 = v22
            _ = v41
        }
        _ = fi37
    }
    v42 := v3[v4]
    v43 := GenS2{F0: v35, F1: false, F2: "a"}
    v35 -= (v35 * v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v43
}
