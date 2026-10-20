package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 float64 = 1e10
    var v2 int32 = -865
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 19; fi5++ {
        if false {
            v6 := GenS1{F0: -32768, F1: 2, F2: false}
            for fi7 := 0; fi7 < 5; fi7++ {
                v8 := GenS2{F0: 39}
                delete(v4, 42)
                _ = fi7
                _ = v8
            }
            fmt.Println(v2)
            _ = v6
        }
        if (fi5 == 32767) {
            switch ("foo bar" + "a") {
            case "hello":
                var v9 int16 = 32767
                _ = v9
            }
            if (1.5 <= (1.5 + 0.1)) {
                if false {
                    delete(v3, ("hello" + "\n\t"))
                    var v10 int = fi5
                    v11 := GenS3{F0: 1, F1: 9223372036854775807, F2: 641}
                    _ = v10
                    _ = v11
                } else {
                    v1 = float64(1e10)
                    var v12 int32 = (v2 ^ v2)
                    var v13 int8 = 127
                    _ = v12
                    _ = v13
                }
                v14 := []int{(fi5 & (-(fi5))), -1}
                switch "foo bar" {
                case "\n\t":
                    v15 := GenS4{F0: 18446744073709551615, F1: 42, F2: 0}
                    _ = v15
                }
                _ = v14
            } else {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v1 += ((v1 / (v1 * ((1e10 * v1) - float64(1e10)))) + (v1 * (3.14159 + v1)))
                    _ = fi16
                }
                var v17 uint32 = 0
                v2 -= (v2 | int32(v17))
                _ = v17
            }
            switch (fi5 ^ 32767) {
            case -2147483648:
                v18 := []int{(fi5 * fi5)}
                _ = v18
            case 127:
                v2 += v2
            case 4294967295:
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := GenS5{F0: 850}
                    switch "" {
                    case "hello":
                        v21 := GenS4{F0: 42, F1: 0, F2: 1}
                        _ = len(v3)
                        var v22 uint64 = 18446744073709551615
                        _ = v21
                        _ = v22
                    }
                    _ = fi19
                    _ = v20
                }
            default:
                if false {
                    if true {
                        v23 := v3["foo bar"]
                        _ = v23
                    } else {
                        var v24 uint64 = 824
                        v2 = (v2 ^ (^(v2)))
                        _ = v24
                    }
                    if (1 != 208) {
                        v1 = 1.5
                    } else {
                        v1 += 0.1
                        v1 -= ((v1 / v1) - (v1 / v1))
                        fmt.Println(v2)
                    }
                } else {
                    switch "a" {
                    case "hello":
                        v1 += 1.5
                    }
                    for fi25 := 0; fi25 < 3; fi25++ {
                        var v26 uint32 = 65535
                        var v27 bool = true
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
            }
        }
        fmt.Println(v2)
        _ = fi5
    }
    v28 := GenS3{F0: 119, F1: 901, F2: 65535}
    switch ("" + ("" + "a")) {
    case "":
        v1 = 1.5
    case "foo bar":
        fmt.Println(v2, v2)
    case "\n\t":
        var v29 int32 = 0
        _ = v29
    default:
        v2 = 0
    }
    v30 := []int{}
    for fi31 := 0; fi31 < 5; fi31++ {
        if (true || (!((-656 >= -2147483648) || false))) {
            v2 += 1
        } else {
            var v32 int8 = 0
            _ = v32
        }
        _ = fi31
    }
    v1 = 0.1
    v33 := v3["hello"]
    var v34 int8 = 27
    var v35 uint = 18446744073709551615
    v36 := GenS3{F0: v34, F1: 0, F2: 933}
    v37 := GenS2{F0: 88}
    var v38 string = (("hello" + (("hello" + "hello") + "a")) + ("hello" + "hello"))
    fmt.Println(v1)
    var v39 uint16 = uint16(v33)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
