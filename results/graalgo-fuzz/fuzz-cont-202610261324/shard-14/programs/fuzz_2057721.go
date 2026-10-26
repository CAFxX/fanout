package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 int32 }
func genFn0(p0 uint16, p1 bool) uint16 {
    switch -2147483648 {
    case 987:
        fmt.Println(p0, p1)
    case -970:
        fmt.Println(p0, p1)
    }
    for fi4 := 0; fi4 < 9; fi4++ {
        p1 = false
        _ = fi4
    }
    var v5 uint64 = 2
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{fi6, (fi6 + fi6), fi6}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return (p0 / ((0) | 1))
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []uint32{(v1 / ((v1) | 1)), 1, v1}
    v3 := []int{0, -9223372036854775808}
    var v8 float32 = 3.14159
    switch ("hello" + (("a" + ("foo bar" + "a")) + (("\n\t" + (((("hello" + "") + "foo bar") + "\n\t") + "a")) + "\n\t"))) {
    case "hello":
        if (!true) {
            if true {
                _ = len(v3)
            } else {
                v8 = v8
                if true {
                    v1 = (^(v1))
                    for fi9 := 0; fi9 < 6; fi9++ {
                        fmt.Println(fi9, v8)
                        _ = fi9
                    }
                    var v10 int16 = 42
                    _ = v10
                } else {
                    var v11 string = ("" + "\n\t")
                    _ = v11
                }
                switch (("a" + "\n\t") + "hello") {
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = int(v1)
                    }
                }
            }
        }
    case "a":
        if len(v2) > 0 {
            v2[0] = (v1 * (v1 % ((2147483647) | 1)))
        }
    case "\n\t":
        v8 = v8
    }
    if false {
        var v12 int = 4294967295
        v12 = v12
        _ = v12
    }
    var v13 int64 = 1
    v13 = v13
    var v14 int8 = -1
    v2 = append(v2, 42)
    var v15 float32 = (v8 / (v8 / float32((-2.25 - (1e10 * 1.5)))))
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            if ((!true) || false) {
                var v18 uint64 = 2
                v8 += (0.0 - 1e10)
                v13 += v13
                _ = v18
            }
            _ = fi17
        }
        v3 = append(v3, 42)
        _ = fi16
    }
    v19 := GenS1{F0: -2147483648, F1: 42}
    var v20 int64 = (v13 % (((v13 * (^(v13)))) | 1))
    switch "" {
    case "a":
        v20 = (v13 + 9223372036854775807)
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        if false {
            v22 := GenS2{F0: genFn0(2, false), F1: (v14 | 112), F2: int32(fi21)}
            _ = v22
        }
        if (!false) {
            var v23 int32 = 1
            v24 := GenS1{F0: 32767, F1: 1}
            _ = v23
            _ = v24
        } else {
            var v25 uint8 = 148
            _ = v25
        }
        switch ("hello" + "hello") {
        case "a":
            v26 := GenS2{F0: 958, F1: v14, F2: -376}
            _ = v26
        }
        _ = fi21
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 int64 = v20
        var v29 uint8 = 1
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
}
