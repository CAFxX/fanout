package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int8 }
type GenS2 struct { F0 string; F1 int32 }
func genFn0(p0 uint8, p1 int, p2 int) uint16 {
    for fi5 := 0; fi5 < 16; fi5++ {
        if (1 > 827) {
            for fi6 := 0; fi6 < 7; fi6++ {
                fmt.Println(p1)
                p0 = (p0 & (^(p0)))
                if (false || true) {
                    v7 := []int{(p1 >> ((299) & 63)), (p1 - -1)}
                    v8 := GenS1{F0: "", F1: ("a" + ""), F2: 1}
                    switch 0 {
                    case 2147483647:
                        v9 := GenS2{F0: ("\n\t" + "foo bar"), F1: -2147483648}
                        _ = v9
                    }
                    _ = v7
                    _ = v8
                }
                var v10 uint16 = 1
                _ = fi6
                _ = v10
            }
            p0 += (p0 >> ((2) & 63))
            for fi11 := 0; fi11 < 11; fi11++ {
                if true {
                    var v12 int32 = 2147483647
                    var v13 float32 = 1e10
                    _ = v12
                    _ = v13
                }
                v14 := GenS1{F0: "a", F1: "foo bar", F2: -112}
                if (true && (1 == 1)) {
                    var v15 float64 = 1.5
                    _ = v15
                }
                _ = fi11
                _ = v14
            }
        }
        _ = fi5
    }
    v16 := GenS2{F0: ((("hello" + (("\n\t" + ("" + "\n\t")) + (("\n\t" + "a") + "foo bar"))) + ("\n\t" + "hello")) + "hello"), F1: 91}
    var v17 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 0
}
func genFn1(p0 uint64, p1 uint64, p2 int32) uint16 {
    p1 -= p0
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 bool = false
    var v2 string = "hello"
    var v3 string = ((v2 + v2) + "hello")
    var v4 float32 = 3.14159
    var v18 string = v3
    for fi19 := 0; fi19 < 10; fi19++ {
        if v1 {
            for fi20 := 0; fi20 < 12; fi20++ {
                v18 = v2
                v1 = (v4 != v4)
                _ = fi20
            }
        }
        switch v2 {
        case "\n\t":
            v21 := GenS1{F0: (v18 + (v18 + v18)), F1: "foo bar", F2: -100}
            _ = v21
        case "hello":
            v22 := []int{(fi19 << ((42) & 63)), (-(fi19)), fi19}
            _ = v22
        case "":
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := []int{42, 42}
                _ = fi23
                _ = v24
            }
        default:
            fmt.Println(v2)
        }
        _ = fi19
    }
    var v25 int = 303
    v26 := GenS2{F0: v2, F1: 625}
    if false {
        var v27 uint16 = 42
        _ = v27
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 int64 = 9223372036854775807
        for fi30 := 0; fi30 < 3; fi30++ {
            var v31 uint16 = 0
            var v32 int32 = -933
            var v33 uint32 = 65535
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi28
        _ = v29
    }
    if v1 {
        v4 = v4
    } else {
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 uint16 = 255
            _ = fi34
            _ = v35
        }
        v3 = ("hello" + "")
        for fi36 := 0; fi36 < 12; fi36++ {
            v37 := GenS1{F0: v2, F1: v2, F2: 127}
            _ = fi36
            _ = v37
        }
    }
    v38 := []int{v25, v25}
    var v39 uint8 = 127
    if (v3 == ("" + ("" + ""))) {
        var v40 string = v18
        fmt.Println(v40)
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
