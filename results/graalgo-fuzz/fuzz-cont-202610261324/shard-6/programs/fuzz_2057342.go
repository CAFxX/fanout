package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 string }
type GenS2 struct { F0 int32; F1 string; F2 bool }
type GenS3 struct { F0 int64; F1 bool; F2 uint64 }
type GenS4 struct { F0 int64; F1 int32; F2 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 32767
    var v3 uint16 = (v2 << ((652) & 63))
    var v4 int16 = 0
    var v5 uint = 18446744073709551615
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: -9223372036854775808, F1: v3, F2: ("" + "")}
    v8 := GenS1{F0: -401, F1: 1, F2: ""}
    switch (("hello" + "foo bar") + "") {
    case "foo bar":
        switch 233 {
        case -781:
            switch 0 {
            case 9223372036854775807:
                fmt.Println(v3, v4, v4)
            case 760:
                var v9 int8 = int8(v4)
                _ = v9
            }
        case -415:
            v10 := GenS1{F0: -1, F1: v2, F2: "foo bar"}
            _ = v10
        case -2147483648:
            if true {
                v5 -= v5
                fmt.Println(v3, v4)
                if true {
                    if (false && true) {
                        v2 -= (v2 + (v3 ^ v3))
                    } else {
                        v11 := GenS1{F0: 32768, F1: (v3 - v3), F2: "a"}
                        v12 := []int{32767, 32768, 32767, 32767}
                        var v13 bool = false
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    fmt.Println(v3, v5)
                }
            }
        default:
            v14 := GenS2{F0: 0, F1: ("foo bar" + "a"), F2: false}
            _ = v14
        }
    case "":
        v15 := []int{1, -424, 32768}
        _ = v15
    }
    switch "hello" {
    case "":
        if true {
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := GenS2{F0: 2147483647, F1: "", F2: true}
                fmt.Println(v2, v4)
                _ = fi16
                _ = v17
            }
            v5 -= v5
        }
    }
    v18 := []int{853}
    var v19 uint32 = 2
    var v20 string = ("hello" + "\n\t")
    for fi21 := 0; fi21 < 8; fi21++ {
        if true {
            v2 = (v3 - (v2 ^ (^(v2))))
            for fi22 := 0; fi22 < 20; fi22++ {
                if (!((v19 - v19) >= v19)) {
                    v4 += (v4 * v4)
                    if (uint16(v5) > (v2 / ((v3) | 1))) {
                        v18 = append(v18, 486)
                        v23 := GenS2{F0: -2147483648, F1: ((("a" + ("\n\t" + v20)) + (v20 + v20)) + ("\n\t" + "")), F2: (32768 < (fi22 % ((505) | 1)))}
                        _ = v23
                    }
                    _ = len(v18)
                }
                v24 := GenS2{F0: -1, F1: "foo bar", F2: false}
                v20 = v20
                var v25 int8 = -1
                delete(v6, 0)
                var v26 uint8 = 2
                _ = fi22
                _ = v24
                _ = v25
                _ = v26
            }
        } else {
            var v27 string = v20
            _ = v27
        }
        _ = fi21
    }
    var v28 int8 = -1
    v29 := GenS1{F0: 258, F1: v2, F2: (v20 + v20)}
    var v30 int16 = v4
    v19 = 65535
    var v31 uint32 = v19
    var v32 int = 515
    fmt.Println(v28)
    v3 += v3
    if true {
        v3 -= 974
    } else {
        for fi33 := 0; fi33 < 14; fi33++ {
            v34 := GenS3{F0: 588, F1: false, F2: 1}
            v35 := GenS4{F0: 1, F1: -2147483648, F2: (v4 / ((v30) | 1))}
            _ = fi33
            _ = v34
            _ = v35
        }
        v5 += (v5 >> ((uint(v2)) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
