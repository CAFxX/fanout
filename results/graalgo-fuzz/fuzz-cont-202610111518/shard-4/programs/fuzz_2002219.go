package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int16 }
func genFn0(p0 uint16, p1 int32) int {
    fmt.Println(p1, p1)
    if true {
        for fi4 := 0; fi4 < 8; fi4++ {
            for fi5 := 0; fi5 < 2; fi5++ {
                v6 := []int{fi5, fi4, fi5, -2147483648}
                v7 := GenS1{F0: -1, F1: -9223372036854775808}
                v6 = append(v6, fi4)
                _ = fi5
                _ = v6
                _ = v7
            }
            p1 = -1
            _ = fi4
        }
    }
    var v8 uint32 = 475
    for fi9 := 0; fi9 < 13; fi9++ {
        v8 = 118
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return 9223372036854775807
}
func genFn1() uint {
    var v10 int8 = -128
    if false {
        var v11 uint32 = 2147483648
        _ = v11
    } else {
        if true {
            var v12 uint = 0
            var v13 int32 = 1
            _ = v12
            _ = v13
        }
        v14 := []int{-783}
        switch ((("a" + "") + (("foo bar" + ("" + ("foo bar" + "foo bar"))) + ("foo bar" + ("a" + "\n\t")))) + (("hello" + (("\n\t" + ("hello" + "hello")) + "hello")) + ("foo bar" + "\n\t"))) {
        case "":
            v15 := GenS1{F0: 2147483647, F1: -1}
            _ = v15
        case "a":
            var v16 bool = (-518 == genFn0(460, -1))
            _ = v16
        case "hello":
            v10 += v10
        }
        _ = v14
    }
    _ = v10
    return 457
}

func main() {
    v1 := []string{"a", ("foo bar" + "")}
    v2 := make(map[int]uint32)
    var v3 int32 = 0
    var v17 string = "hello"
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            if true {
                var v20 int64 = -9223372036854775808
                v21 := GenS1{F0: 1, F1: 0}
                var v22 uint8 = 255
                if len(v1) > 0 {
                    v1[0] = ((v17 + "hello") + v17)
                }
                _ = v20
                _ = v21
                _ = v22
            } else {
                v17 = "foo bar"
                v23 := GenS1{F0: -1, F1: -1}
                var v24 int = fi19
                _ = v23
                _ = v24
            }
            _ = fi19
        }
        v25 := GenS1{F0: v3, F1: 180}
        if true {
            _ = len(v1)
            for fi26 := 0; fi26 < 16; fi26++ {
                fmt.Println(v17)
                _ = fi26
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                v28 := v2[fi18]
                v17 = ((v17 + (v17 + (v17 + ("a" + "foo bar")))) + v17)
                var v29 int8 = -42
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        _ = fi18
        _ = v25
    }
    var v30 uint64 = 1
    if true {
        if (0.0 != 3.14159) {
            v31 := GenS2{F0: v3}
            _ = v31
        } else {
            v32 := GenS2{F0: 10}
            v17 = (v17 + "foo bar")
            _ = v32
        }
    } else {
        switch "" {
        case "foo bar":
            v3 = 630
        default:
            for fi33 := 0; fi33 < 1; fi33++ {
                var v34 int8 = 111
                v35 := GenS3{F0: (fi33 * (^(fi33))), F1: 32767}
                v17 = v17
                _ = fi33
                _ = v34
                _ = v35
            }
        }
    }
    var v36 uint8 = 158
    switch 9223372036854775807 {
    case 32767:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 127:
        var v37 uint32 = 522
        _ = v37
    }
    if len(v1) > 0 {
        v1[0] = v17
    }
    _ = len(v2)
    if (220 >= v36) {
        v3 = (v3 % ((v3) | 1))
        v30 -= (v30 ^ (v30 + (v30 << ((65535) & 63))))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v36
}
