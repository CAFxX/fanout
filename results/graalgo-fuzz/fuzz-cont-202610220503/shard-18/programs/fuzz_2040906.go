package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
type GenS2 struct { F0 int32; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint64 }
func genFn0() int8 {
    var v5 uint32 = 243
    v6 := GenS1{F0: 0, F1: 128, F2: true}
    v5 -= (^(v5))
    var v7 bool = false
    v8 := GenS2{F0: 42, F1: 2, F2: -1}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int, p1 string) int16 {
    v9 := GenS1{F0: 101, F1: 18446744073709551615, F2: true}
    switch p1 {
    case "a":
        p1 = (p1 + p1)
    case "foo bar":
        v10 := GenS3{F0: 337}
        _ = v10
    default:
        for fi11 := 0; fi11 < 5; fi11++ {
            switch ("hello" + "") {
            case "hello":
                v12 := GenS4{F0: "", F1: 2147483647, F2: 526}
                _ = v12
            case "":
                var v13 uint64 = 577
                _ = v13
            }
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return -32768
}

func main() {
    var v1 int64 = 786
    var v2 uint8 = 255
    var v3 int32 = -925
    v4 := make(map[int]int)
    v14 := v4[127]
    v14 = v14
    for fi15 := 0; fi15 < 1; fi15++ {
        v3 = (v3 >> ((1) & 63))
        _ = fi15
    }
    if (0 != 82) {
        fmt.Println(v3)
        fmt.Println(v14)
    } else {
        fmt.Println(v1)
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int8 = 1
        if true {
            v18 := GenS4{F0: ("hello" + "hello"), F1: 0, F2: 2147483647}
            var v19 string = ""
            _ = v18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    var v20 float32 = 1e10
    switch ("hello" + ((((((("a" + ("" + "")) + "\n\t") + "hello") + "\n\t") + "a") + "hello") + "")) {
    case "":
        var v21 uint8 = (v2 >> ((65535) & 63))
        _ = v21
    case "hello":
        delete(v4, (-(v14)))
    case "foo bar":
        var v22 string = ("a" + ("a" + ""))
        _ = v22
    default:
        var v23 uint32 = 0
        _ = v23
    }
    var v24 uint32 = 1
    v25 := GenS1{F0: 1, F1: 0, F2: (true && true)}
    for fi26 := 0; fi26 < 11; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 uint64 = 0
            v2 += (v2 & 2)
            _ = fi27
            _ = v28
        }
        switch (v14 ^ v14) {
        case -9223372036854775808:
            v3 = (^(v3))
        case -2147483648:
            for fi29 := 0; fi29 < 8; fi29++ {
                fmt.Println(v20, v1, v24)
                _ = fi29
            }
        }
        _ = fi26
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := GenS4{F0: "\n\t", F1: (v24 | (v24 * (v24 | (^(v24))))), F2: v24}
        _ = len(v4)
        var v32 bool = false
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := []int{int(v1)}
    for fi34 := 0; fi34 < 19; fi34++ {
        if len(v33) > 0 {
            v33[0] = (fi34 - (-(v14)))
        }
        _ = fi34
    }
    switch "\n\t" {
    case "a":
        var v35 uint32 = v24
        _ = v35
    default:
        switch ("hello" + ("" + "a")) {
        case "a":
            v36 := v4[v14]
            _ = v36
        default:
            if true {
                for fi37 := 0; fi37 < 2; fi37++ {
                    var v38 int = (v14 - (fi37 % ((fi37) | 1)))
                    v20 -= 0.0
                    _ = fi37
                    _ = v38
                }
            } else {
                fmt.Println(v24, v1)
                switch (v14 << ((18446744073709551615) & 63)) {
                case 4294967295:
                    var v39 uint = 149
                    _ = v39
                case 42:
                    var v40 uint8 = 0
                    _ = v40
                case 32768:
                    v41 := GenS5{F0: v24, F1: (^(v24)), F2: 4294967295}
                    _ = v41
                default:
                    v1 = v1
                }
                fmt.Println(v20)
            }
        }
    }
    v33 = append(v33, v14)
    v3 = (v3 - v3)
    var v42 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v42
}
