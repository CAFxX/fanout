package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint8 = 255
    var v2 int16 = 42
    v3 := make(map[string]int)
    var v4 int16 = (v2 % ((-1) | 1))
    var v5 uint16 = 32767
    v6 := make(map[string]uint32)
    delete(v6, "hello")
    var v7 int64 = 0
    delete(v3, (("\n\t" + "foo bar") + "a"))
    var v8 int32 = -1
    _ = len(v3)
    fmt.Println(v8, v7)
    if (!(((32767 < 0) || true) || true)) {
        v4 -= (v2 + v4)
        var v9 bool = true
        _ = v9
    } else {
        var v10 int8 = 24
        _ = v10
    }
    v1 = (^(v1))
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 float32 = (-2.25 - (float32(3.14159) / 3.14159))
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 float32 = v12
            v8 = (-(v8))
            var v15 uint64 = 2
            var v16 uint64 = 500
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        v17 := GenS1{F0: 0, F1: 620, F2: (^(v5))}
        _ = fi11
        _ = v12
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 uint64 = 18446744073709551615
        var v20 uint16 = v5
        _ = fi18
        _ = v19
        _ = v20
    }
    switch "hello" {
    case "\n\t":
        v21 := v6["foo bar"]
        _ = v21
    case "a":
        for fi22 := 0; fi22 < 2; fi22++ {
            if true {
                v8 = 1
            } else {
                for fi23 := 0; fi23 < 11; fi23++ {
                    switch ("a" + ("hello" + "foo bar")) {
                    case "\n\t":
                        v24 := GenS1{F0: 1, F1: 255, F2: (^(v5))}
                        _ = v24
                    case "a":
                        v1 = v1
                    case "hello":
                        var v25 bool = true
                        _ = v25
                    }
                    fmt.Println(v2)
                    _ = fi23
                }
                v4 = v2
                switch -866 {
                case 642:
                    v26 := []int{}
                    _ = v26
                }
            }
            _ = fi22
        }
    default:
        for fi27 := 0; fi27 < 5; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 string = "hello"
                v30 := GenS1{F0: 4294967295, F1: 758, F2: (v5 & v5)}
                if true {
                    v8 = (v8 >> ((18446744073709551615) & 63))
                    v1 += v1
                    v3["hello"] = -503
                }
                var v31 int16 = -1
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            if (3.14159 == (0.0 / (0.1 * 1e10))) {
                v1 = (^(v1))
                fmt.Println(v5)
                fmt.Println(v2)
            } else {
                if false {
                    v32 := GenS1{F0: 0, F1: 2, F2: 312}
                    var v33 bool = (18446744073709551615 <= 42)
                    _ = v32
                    _ = v33
                }
                v34 := GenS2{F0: 0}
                _ = v34
            }
            v7 = v7
            _ = fi27
        }
    }
    switch "hello" {
    case "hello":
        for fi35 := 0; fi35 < 11; fi35++ {
            var v36 bool = false
            switch 2147483647 {
            case -2147483648:
                for fi37 := 0; fi37 < 10; fi37++ {
                    var v38 int64 = v7
                    v2 += (^(v2))
                    _ = fi37
                    _ = v38
                }
            case 515:
                v2 += v2
            }
            _ = fi35
            _ = v36
        }
    default:
        var v39 int16 = (v4 / ((0) | 1))
        _ = v39
    }
    if (!(("hello" + "") < "a")) {
        for fi40 := 0; fi40 < 6; fi40++ {
            var v41 string = ((("a" + ("foo bar" + ("" + "hello"))) + ((("\n\t" + "\n\t") + "\n\t") + "foo bar")) + "foo bar")
            _ = fi40
            _ = v41
        }
    } else {
        delete(v6, "foo bar")
        fmt.Println(v7)
    }
    v42 := GenS1{F0: 4294967295, F1: 2, F2: v5}
    v7 = (v7 ^ v7)
    v43 := []int{-571, -819}
    v2 -= (v2 & (v2 ^ v4))
    switch "hello" {
    case "":
        v1 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
