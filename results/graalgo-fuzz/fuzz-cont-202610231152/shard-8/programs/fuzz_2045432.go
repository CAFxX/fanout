package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() int {
    v5 := GenS1{F0: 0}
    if false {
        v6 := []int{127, 965, -1}
        var v7 uint8 = 111
        _ = v6
        _ = v7
    }
    v8 := []int{}
    var v9 int16 = 1000
    _ = v5
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 int64, p1 int32, p2 uint32) bool {
    v10 := GenS1{F0: -410}
    p2 += (^(p2))
    switch -2147483648 {
    case 2147483647:
        p0 = (p0 << ((0) & 63))
    default:
        if ((p2 & p2) >= (^(p2))) {
            p2 -= (p2 * (p2 % (((^(p2))) | 1)))
            if true {
                v11 := []int{128, genFn0(), 0}
                fmt.Println(p2)
                var v12 float32 = float32(0.1)
                _ = v11
                _ = v12
            }
            p1 = p1
        }
    }
    if false {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 uint32 = (^(p2))
            var v15 uint32 = v14
            _ = fi13
            _ = v14
            _ = v15
        }
        switch "" {
        case "\n\t":
            v16 := []int{1, 4294967295, -528}
            _ = v16
        default:
            var v17 float32 = 0.1
            _ = v17
        }
    } else {
        switch "a" {
        case "a":
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 float32 = 0.1
                for fi20 := 0; fi20 < 13; fi20++ {
                    var v21 int32 = p1
                    var v22 uint = 721
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi18
                _ = v19
            }
        case "\n\t":
            p0 = (p0 << ((115) & 63))
        default:
            var v23 bool = ((((float64(-2.25) * 1e10) * 3.14159) * (0.1 + ((1.5 - -2.25) * 3.14159))) != (0.1 - (1e10 / ((1e10 * (float64(3.14159) * float64(1.5))) / 1.0))))
            _ = v23
        }
        p1 -= (^(p1))
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return false
}

func main() {
    var v1 float64 = (float64(-2.25) + 0.0)
    v2 := []uint32{0}
    v3 := make(map[uint32]uint32)
    v4 := []int{}
    v3[958] = 715
    var v24 int64 = -1
    v25 := GenS1{F0: 128}
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 int32 = 1
        _ = fi26
        _ = v27
    }
    v1 = 0.0
    switch "a" {
    case "hello":
        if false {
            v24 = (v24 >> ((65535) & 63))
        }
    }
    var v28 string = "foo bar"
    switch -899 {
    case -332:
        var v29 bool = (1 < 973)
        _ = v29
    case 233:
        if (!false) {
            v28 = ("hello" + (("\n\t" + v28) + (("" + ("hello" + (v28 + "foo bar"))) + ("a" + ""))))
        } else {
            for fi30 := 0; fi30 < 20; fi30++ {
                if true {
                    v24 = (v24 & -637)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi30
            }
            if true {
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
            } else {
                v28 = v28
                var v31 uint8 = 0
                _ = v31
            }
            _ = len(v2)
        }
    default:
        fmt.Println(v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
