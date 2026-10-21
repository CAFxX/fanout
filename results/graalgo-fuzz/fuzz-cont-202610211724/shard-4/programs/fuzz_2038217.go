package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint64; F2 bool }
func genFn0(p0 uint16, p1 uint16, p2 uint) string {
    switch ("a" + (("foo bar" + "") + "hello")) {
    case "a":
        fmt.Println(p0, p1, p0)
    case "":
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := []int{fi6, 392}
            _ = fi6
            _ = v7
        }
    }
    v8 := GenS1{F0: -659, F1: 42, F2: ("hello" + "foo bar")}
    v9 := GenS1{F0: 0, F1: 0, F2: ""}
    switch 127 {
    case -2147483648:
        v10 := []int{127, -9223372036854775808, 9223372036854775807, 4294967295}
        _ = v10
    case 32767:
        v11 := []int{9223372036854775807, 1, 9223372036854775807}
        _ = v11
    default:
        p1 = 1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return ("foo bar" + "")
}
func genFn1() int8 {
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint8 = 1
        v14 := GenS2{F0: ""}
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 uint8 = (v13 << ((2) & 63))
            _ = fi15
            _ = v16
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v17 := []int{4294967295, 825, 9223372036854775807}
    _ = v17
    return 100
}

func main() {
    var v1 int32 = -2147483648
    var v2 float64 = 3.14159
    v3 := make(map[int]uint32)
    var v4 int8 = -1
    var v5 int = 4294967295
    var v18 uint = 255
    v19 := []int{2147483647}
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v2 = 3.14159
    var v20 uint64 = 607
    if false {
        _ = len(v19)
        var v21 int64 = -9223372036854775808
        switch "hello" {
        case "foo bar":
            switch (v5 / (((v5 | (v5 % (((v5 ^ 710)) | 1)))) | 1)) {
            case 1:
                if true {
                    if (!(v2 != v2)) {
                        if (2 == 891) {
                            switch "a" {
                            case "":
                                v1 = (v1 << ((307) & 63))
                            case "\n\t":
                                var v22 float32 = (((((0.1 - 0.0) / 1e10) + 3.14159) - (0.0 - (0.0 * 1.5))) / 3.14159)
                                _ = v22
                            }
                        } else {
                            fmt.Println(v21, v2)
                            var v23 int32 = v1
                            _ = v23
                        }
                        v24 := GenS2{F0: "a"}
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 2; fi25++ {
                        if false {
                            delete(v3, (v5 * -1))
                            var v26 int64 = v21
                            v4 = v4
                            _ = v26
                        }
                        v19 = append(v19, v5)
                        v27 := GenS3{F0: v20, F1: v20, F2: true}
                        _ = fi25
                        _ = v27
                    }
                }
            case 0:
                if ((v21 * v21) > (v21 << ((v18) & 63))) {
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 bool = (v2 < v2)
                        v29 = v29
                        _ = fi28
                        _ = v29
                    }
                }
            }
        case "hello":
            var v30 int = 32767
            _ = v30
        case "":
            v20 = 2
        }
        _ = v21
    }
    v1 -= (^(v1))
    for fi31 := 0; fi31 < 16; fi31++ {
        v4 -= 100
        var v32 string = ("foo bar" + "hello")
        for fi33 := 0; fi33 < 19; fi33++ {
            v32 = (v32 + v32)
            v2 += float64(-2.25)
            var v34 float64 = float64((0.0 / 3.14159))
            _ = fi33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    v35 := GenS1{F0: 1, F1: int16(v20), F2: "\n\t"}
    for fi36 := 0; fi36 < 15; fi36++ {
        delete(v3, (v5 >> ((v18) & 63)))
        if true {
            var v37 uint = (v18 ^ (v18 / (((v18 + (^(v18)))) | 1)))
            v4 += (v4 << ((25) & 63))
            _ = v37
        } else {
            delete(v3, (-(v5)))
            v20 -= (v20 ^ v20)
        }
        v3[v5] = 42
        _ = fi36
    }
    if ((("a" + "") + ("foo bar" + (("" + ("" + "\n\t")) + ""))) >= "hello") {
        var v38 uint64 = v20
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
