package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint8 }
func genFn0() uint32 {
    if (2147483647 != 42) {
        if (42 < 42) {
            var v6 int8 = 42
            var v7 int32 = 2147483647
            _ = v6
            _ = v7
        } else {
            v8 := []int{}
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint64 = 4294967295
                v11 := GenS1{F0: 80}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        }
    } else {
        var v12 uint8 = 255
        v12 = 127
        if true {
            v13 := []int{127}
            switch -1 {
            case 236:
                v12 = v12
            }
            _ = v13
        }
        _ = v12
    }
    switch (("" + ("" + (("foo bar" + "\n\t") + ""))) + "") {
    case "foo bar":
        v14 := []int{-9223372036854775808}
        _ = v14
    case "hello":
        var v15 int16 = 42
        _ = v15
    case "":
        v16 := GenS1{F0: -128}
        _ = v16
    }
    var v17 int8 = 127
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 uint = 826
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 uint16 = 0
            var v22 uint32 = 2
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 string = "a"
        _ = fi18
        _ = v19
        _ = v23
    }
    _ = v17
    return 565
}
func genFn1() int32 {
    var v24 float64 = 0.0
    fmt.Println(v24)
    for fi25 := 0; fi25 < 8; fi25++ {
        v24 = float64((0.1 * (1e10 * float32(((v24 * 0.1) - float64(-2.25))))))
        _ = fi25
    }
    v26 := GenS1{F0: -21}
    fmt.Println(v24)
    _ = v24
    _ = v26
    return -2147483648
}

func main() {
    var v1 uint = 2
    v2 := []uint32{81}
    v3 := make(map[string]int)
    var v4 uint16 = 32767
    var v5 int16 = 32767
    switch ("" + "\n\t") {
    case "a":
        v1 = 4294967295
    case "foo bar":
        v5 -= (^(v5))
    }
    var v27 bool = true
    _ = len(v3)
    v28 := GenS1{F0: 1}
    _ = len(v2)
    for fi29 := 0; fi29 < 4; fi29++ {
        v27 = false
        v30 := v3[("\n\t" + "foo bar")]
        var v31 int16 = v5
        _ = fi29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        for fi33 := 0; fi33 < 20; fi33++ {
            if false {
                var v34 uint32 = 42
                v35 := GenS1{F0: -101}
                _ = len(v2)
                v36 := GenS1{F0: -128}
                switch ("" + "a") {
                case "\n\t":
                    var v37 float64 = 3.14159
                    _ = v37
                case "hello":
                    v34 += (v34 / (((v34 ^ (^(v34)))) | 1))
                case "":
                    switch (fi32 << ((v1) & 63)) {
                    case 127:
                        v2 = append(v2, 2147483647)
                    case -2147483648:
                        var v38 float64 = (1.5 * -2.25)
                        _ = v38
                    }
                default:
                    v39 := GenS2{F0: v4, F1: 0}
                    _ = v39
                }
                _ = v34
                _ = v35
                _ = v36
            }
            var v40 int8 = -24
            switch ((("" + "") + "") + "a") {
            case "\n\t":
                var v41 string = "hello"
                if len(v2) > 0 {
                    v2[0] = 2
                }
                v27 = (v40 <= 42)
                _ = v41
            case "a":
                _ = len(v2)
            }
            _ = fi33
            _ = v40
        }
        fmt.Println(v1, v1, v4)
        _ = fi32
    }
    fmt.Println(v5, v1)
    var v42 uint16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
