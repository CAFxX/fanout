package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 uint32; F1 uint64; F2 bool }
type GenS3 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS4 struct { F0 int16 }
func genFn0() int8 {
    var v7 int32 = 653
    fmt.Println(v7)
    switch "\n\t" {
    case "\n\t":
        switch ((((("a" + (("a" + "hello") + "\n\t")) + (("foo bar" + ("foo bar" + "hello")) + "foo bar")) + ("hello" + ("a" + "a"))) + "\n\t") + "\n\t") {
        case "":
            v8 := []int{127, 2147483647, 32767, 744}
            _ = v8
        }
    case "a":
        v9 := GenS1{F0: true, F1: 255, F2: ""}
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 14; fi10++ {
            v7 = -2147483648
            var v11 int64 = -1
            _ = fi10
            _ = v11
        }
    default:
        if true {
            v12 := GenS2{F0: 42, F1: 42, F2: true}
            var v13 string = ("\n\t" + ("hello" + "foo bar"))
            _ = v12
            _ = v13
        } else {
            if ((false || (3.14159 != 3.14159)) && true) {
                var v14 int = 2147483647
                _ = v14
            }
        }
    }
    fmt.Println(v7)
    _ = v7
    return 42
}
func genFn1() uint {
    var v15 uint16 = 65535
    v16 := GenS1{F0: false, F1: v15, F2: "foo bar"}
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            if (42 > 177) {
                var v19 uint = 4294967295
                v20 := []int{4294967295}
                v21 := GenS1{F0: false, F1: (v15 << ((v19) & 63)), F2: "hello"}
                _ = v19
                _ = v20
                _ = v21
            } else {
                v15 = v15
                fmt.Println(v15, v15, fi17)
            }
            _ = fi18
        }
        _ = fi17
    }
    var v22 int32 = 42
    v23 := GenS3{F0: 62, F1: 2, F2: 1000}
    _ = v15
    _ = v16
    _ = v22
    _ = v23
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 127
    var v3 int32 = 0
    var v4 int8 = v2
    v5 := make(map[uint32]int)
    var v6 int = 32767
    fmt.Println(v3, v6)
    v24 := GenS3{F0: 0, F1: 18446744073709551615, F2: 653}
    var v25 float64 = (((1e10 / (1e10 + 1.5)) * float64(0.1)) - 1.5)
    v26 := []int{v6, -281, 248}
    fmt.Println(v2, v2, v4)
    if (1 != 921) {
        v27 := GenS4{F0: 0}
        _ = v27
    } else {
        v6 = -695
    }
    var v28 bool = false
    fmt.Println(v2, v28)
    var v29 int8 = 0
    for fi30 := 0; fi30 < 4; fi30++ {
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 int8 = 127
            v32 = (v32 * v32)
            var v33 float64 = (v25 / (v25 / 3.14159))
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    switch "hello" {
    case "":
        if false {
            var v34 float64 = (3.14159 / v25)
            _ = v34
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
