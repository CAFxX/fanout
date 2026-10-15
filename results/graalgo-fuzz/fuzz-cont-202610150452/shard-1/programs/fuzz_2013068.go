package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint8, p1 int32) int32 {
    if true {
        for fi4 := 0; fi4 < 11; fi4++ {
            v5 := GenS1{F0: "", F1: (("a" + ("a" + "foo bar")) + "foo bar")}
            _ = fi4
            _ = v5
        }
    }
    v6 := []int{-571, 611}
    var v7 uint64 = 18446744073709551615
    var v8 int64 = 9223372036854775807
    v9 := GenS2{F0: -75, F1: -1, F2: (v7 ^ 117)}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -292
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int32 = 2147483647
    var v3 int16 = 744
    var v10 int64 = -584
    if true {
        v1 = 1
        switch 653 {
        case 4294967295:
            if (!true) {
                v11 := []int{-2147483648, 4294967295, 32767, -9223372036854775808}
                _ = v11
            } else {
                var v12 int64 = 9223372036854775807
                _ = v12
            }
        case 32767:
            v13 := []int{2147483647, int(v1), 0}
            _ = v13
        case 128:
            v1 = v10
        }
        switch -785 {
        case 9223372036854775807:
            for fi14 := 0; fi14 < 10; fi14++ {
                if false {
                    v15 := []int{(fi14 ^ 0)}
                    v15 = append(v15, 2147483647)
                    _ = v15
                }
                v2 -= v2
                _ = fi14
            }
        case 0:
            var v16 uint8 = 42
            _ = v16
        }
    }
    switch ("\n\t" + (("\n\t" + "foo bar") + ("hello" + "a"))) {
    case "":
        if false {
            for fi17 := 0; fi17 < 17; fi17++ {
                if ((0 < 42) || (float64(1.5) >= -2.25)) {
                    v2 -= (v2 + v2)
                }
                var v18 int16 = (v3 >> ((18446744073709551615) & 63))
                _ = fi17
                _ = v18
            }
        } else {
            v19 := GenS2{F0: 2147483647, F1: (^(v2)), F2: 4294967295}
            _ = v19
        }
    case "hello":
        var v20 int64 = 1
        _ = v20
    default:
        v3 += 32767
    }
    var v21 uint32 = 2147483647
    v21 -= (^(v21))
    v22 := GenS3{F0: 24}
    fmt.Println(v21, v10)
    fmt.Println(v3)
    v23 := GenS1{F0: "foo bar", F1: "a"}
    v10 = 710
    switch 42 {
    case -1:
        v24 := GenS2{F0: -1, F1: (^(v2)), F2: 209}
        _ = v24
    case 2147483647:
        v1 += v1
    case 487:
        if (((v2 % ((v2) | 1)) != -2147483648) || ((!((uint(v21) >= 958) && false)) && true)) {
            v1 = v1
        } else {
            var v25 float32 = 1.5
            switch ("" + "") {
            case "a":
                if false {
                    var v26 string = ("hello" + "hello")
                    _ = v26
                } else {
                    fmt.Println(v10)
                    if false {
                        v27 := GenS2{F0: -2147483648, F1: v2, F2: 42}
                        var v28 int64 = int64(v2)
                        _ = v27
                        _ = v28
                    }
                    var v29 int16 = 1
                    _ = v29
                }
            case "hello":
                var v30 float64 = (3.14159 * 0.1)
                _ = v30
            }
            _ = v25
        }
    }
    switch ("foo bar" + (("hello" + ("hello" + ("hello" + "a"))) + ((("hello" + "foo bar") + "foo bar") + "hello"))) {
    case "foo bar":
        if ((true && true) && false) {
            v31 := []int{}
            fmt.Println(v3)
            _ = v31
        } else {
            switch 2147483647 {
            case 442:
                var v32 uint64 = 0
                _ = v32
            case -17:
                var v33 bool = (!false)
                _ = v33
            }
        }
    case "":
        switch "a" {
        case "\n\t":
            fmt.Println(v2, v2)
        }
    }
    v21 -= v21
    switch "\n\t" {
    case "foo bar":
        var v34 int64 = (v10 - v10)
        _ = v34
    case "hello":
        v10 += v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
