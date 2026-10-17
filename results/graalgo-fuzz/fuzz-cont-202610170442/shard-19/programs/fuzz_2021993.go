package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint32, p1 uint64, p2 string) int64 {
    var v7 int64 = int64(p0)
    v8 := GenS1{F0: (p2 + p2)}
    var v9 int32 = 2147483647
    var v10 uint8 = 127
    var v11 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return (v7 / ((v7) | 1))
}

func main() {
    var v1 uint8 = 168
    var v2 uint64 = 1
    var v3 uint64 = 2
    var v4 uint8 = uint8(v3)
    var v5 uint16 = 1
    var v6 uint64 = v3
    v12 := GenS1{F0: ((((("a" + "foo bar") + "foo bar") + (("hello" + ("hello" + "\n\t")) + (("foo bar" + "hello") + "a"))) + "foo bar") + ("" + "hello"))}
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := []int{}
            if true {
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                v6 -= v2
                if false {
                    v16 := GenS1{F0: (("foo bar" + ((("foo bar" + "") + ("hello" + "")) + "\n\t")) + "foo bar")}
                    _ = v16
                }
                if ((3.14159 >= 1e10) && ("foo bar" >= ("" + "hello"))) {
                    switch (fi13 + -1) {
                    case 32767:
                        v1 = uint8(v2)
                    case -269:
                        var v17 float32 = 0.0
                        _ = v17
                    case -253:
                        var v18 uint8 = v1
                        _ = v18
                    default:
                        v15 = append(v15, fi13)
                    }
                }
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            } else {
                var v19 bool = false
                v6 += v3
                _ = v19
            }
            v3 -= uint64(v1)
            fmt.Println(fi13)
            var v20 uint = 635
            _ = fi14
            _ = v15
            _ = v20
        }
        var v21 int16 = 1
        _ = fi13
        _ = v21
    }
    v22 := []int{-478, 0, 32767, 9223372036854775807}
    v23 := GenS1{F0: "\n\t"}
    v4 = v4
    fmt.Println(v1, v4)
    if true {
        var v24 bool = true
        if len(v22) > 0 {
            v22[0] = 920
        }
        _ = v24
    } else {
        v1 -= 0
        switch ("" + "") {
        case "":
            if true {
                v22 = append(v22, -2147483648)
                fmt.Println(v4, v2)
                var v25 float64 = float64((1.5 - 0.1))
                _ = v25
            } else {
                for fi26 := 0; fi26 < 18; fi26++ {
                    v2 = 813
                    _ = fi26
                }
                v2 = v2
            }
        case "hello":
            for fi27 := 0; fi27 < 14; fi27++ {
                fmt.Println(v1, v4, v3)
                _ = fi27
            }
        }
    }
    switch ("hello" + "foo bar") {
    case "foo bar":
        var v28 int64 = 977
        _ = v28
    }
    if (0 >= 738) {
        switch -2147483648 {
        case 32767:
            var v29 float32 = (0.1 + 1e10)
            _ = v29
        case -766:
            var v30 uint64 = (^(v6))
            _ = v30
        case -838:
            if (42 > 764) {
                if (false && (2147483647 == -2147483648)) {
                    _ = len(v22)
                    v31 := GenS1{F0: "a"}
                    var v32 int32 = 42
                    _ = v31
                    _ = v32
                } else {
                    fmt.Println(v3)
                    for fi33 := 0; fi33 < 20; fi33++ {
                        var v34 float64 = (0.0 + ((1.5 + ((1.5 + (1.5 + 3.14159)) + 0.1)) + 3.14159))
                        _ = len(v22)
                        if (!true) {
                            _ = len(v22)
                            v34 = v34
                            v1 = v1
                        }
                        _ = fi33
                        _ = v34
                    }
                    v5 -= v5
                }
                v35 := GenS1{F0: ("hello" + "hello")}
                _ = v35
            } else {
                var v36 bool = false
                fmt.Println(v36)
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 16; fi37++ {
            if false {
                var v38 bool = (1e10 <= 3.14159)
                _ = v38
            }
            v39 := GenS2{F0: 963, F1: -813}
            var v40 int = fi37
            _ = fi37
            _ = v39
            _ = v40
        }
        for fi41 := 0; fi41 < 10; fi41++ {
            switch fi41 {
            case 904:
                v2 += (v6 % ((593) | 1))
            case 1:
                fmt.Println(v4, v6)
            case 9223372036854775807:
                v2 -= (v6 + 0)
            default:
                var v42 uint32 = 4294967295
                _ = v42
            }
            _ = fi41
        }
    }
    v1 = (v4 << ((uint(v3)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
