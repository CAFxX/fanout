package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint }
func genFn0(p0 uint16) uint {
    fmt.Println(p0)
    var v5 uint8 = 117
    var v6 int32 = -676
    if (580 >= 0) {
        switch 526 {
        case 888:
            if false {
                v6 -= v6
                p0 -= p0
                switch (("a" + "\n\t") + ("hello" + "hello")) {
                case "":
                    if (790 <= 0) {
                        if true {
                            var v7 uint64 = 999
                            _ = v7
                        }
                    }
                }
            }
        case -250:
            v6 = (v6 + v6)
        }
        for fi8 := 0; fi8 < 20; fi8++ {
            switch (-(fi8)) {
            case 127:
                for fi9 := 0; fi9 < 16; fi9++ {
                    v10 := []int{fi8, (fi8 & -598), 127}
                    _ = fi9
                    _ = v10
                }
            case -656:
                var v11 uint = 0
                _ = v11
            default:
                if false {
                    v12 := GenS1{F0: 0, F1: 1, F2: fi8}
                    _ = v12
                } else {
                    v13 := GenS1{F0: -271, F1: 117, F2: fi8}
                    v14 := []int{}
                    _ = v13
                    _ = v14
                }
            }
            var v15 string = "a"
            _ = fi8
            _ = v15
        }
    } else {
        for fi16 := 0; fi16 < 19; fi16++ {
            switch "" {
            case "a":
                v5 = (v5 | v5)
            case "\n\t":
                var v17 int16 = -897
                _ = v17
            case "hello":
                if false {
                    v18 := []int{fi16, (^(fi16))}
                    _ = len(v18)
                    _ = v18
                }
            default:
                if (!(!((v6 | (v6 >> ((794) & 63))) > 1))) {
                    v5 = v5
                    v19 := GenS1{F0: -613, F1: 0, F2: 32768}
                    _ = v19
                }
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 bool = ((v5 * v5) <= (v5 | v5))
                _ = fi20
                _ = v21
            }
            _ = fi16
        }
        v5 += (^(v5))
    }
    _ = p0
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := []string{"hello", ("\n\t" + "")}
    var v2 uint16 = 65535
    var v3 bool = (42 <= 127)
    var v4 uint8 = 127
    var v22 uint64 = 305
    v23 := GenS1{F0: -621, F1: 0, F2: -1}
    var v24 float32 = 0.0
    var v25 int32 = -916
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS2{F0: ""}
        v28 := GenS3{F0: ((v25 / ((v25) | 1)) <= v25), F1: 18446744073709551615}
        v22 -= (v22 ^ 0)
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v4, v3)
    _ = len(v1)
    v29 := GenS2{F0: ""}
    for fi30 := 0; fi30 < 13; fi30++ {
        if len(v1) > 0 {
            v1[0] = ("" + ("foo bar" + "foo bar"))
        }
        var v31 bool = true
        var v32 uint16 = (v2 << ((uint(v2)) & 63))
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
