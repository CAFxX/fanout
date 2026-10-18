package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS3 struct { F0 string }
func genFn0() string {
    var v4 uint32 = 404
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(v4)
            v7 := []int{fi5}
            v8 := GenS1{F0: 2}
            _ = fi6
            _ = v7
            _ = v8
        }
        if (false || ((184 != -264) && true)) {
            var v9 float64 = 1.5
            var v10 string = "\n\t"
            _ = v10
            _ = v9
        }
        v11 := []int{(fi5 ^ (fi5 | (fi5 & fi5))), -1}
        _ = fi5
        _ = v11
    }
    _ = v4
    return "a"
}

func main() {
    var v1 uint16 = 32767
    var v2 uint64 = 671
    var v3 uint32 = 42
    v12 := []int{568, 32768}
    for fi13 := 0; fi13 < 17; fi13++ {
        if (100 == -22) {
            v1 -= uint16(v2)
            for fi14 := 0; fi14 < 4; fi14++ {
                v1 -= v1
                _ = fi14
            }
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            if (!(uint(v1) != uint(v2))) {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                if len(v12) > 0 {
                    v12[0] = (fi15 * (fi13 - (fi13 ^ fi15)))
                }
                if false {
                    var v16 int32 = -1
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    fmt.Println(fi15, v1, fi15)
                    _ = v16
                }
            }
            _ = fi15
        }
        v17 := GenS2{F0: false, F1: 42, F2: v3}
        _ = fi13
        _ = v17
    }
    v18 := GenS1{F0: 1}
    v3 += v3
    var v19 int32 = -2147483648
    fmt.Println(v3)
    v2 = v2
    var v20 float64 = (-2.25 - (0.1 + 1e10))
    v2 = (^(v2))
    for fi21 := 0; fi21 < 12; fi21++ {
        if (!true) {
            v20 += 0.1
        } else {
            v3 = (v3 | v3)
        }
        fmt.Println(v20, v2, v2)
        v22 := GenS2{F0: true, F1: 2, F2: (v3 | (^(v3)))}
        _ = fi21
        _ = v22
    }
    v3 += 0
    if (42 <= -2147483648) {
        fmt.Println(v3, v2, v1)
        v23 := GenS3{F0: "foo bar"}
        _ = v23
    }
    if (false && ((823 == 42) || false)) {
        if false {
            var v24 string = ("\n\t" + "\n\t")
            fmt.Println(v19, v3, v1)
            _ = v24
        }
    }
    v12 = append(v12, 2147483647)
    for fi25 := 0; fi25 < 7; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            _ = len(v12)
            switch (fi26 << ((65535) & 63)) {
            case -856:
                fmt.Println(fi25, fi26, fi25)
            }
            _ = fi26
        }
        var v27 uint32 = 2
        fmt.Println(v20, v19)
        _ = fi25
        _ = v27
    }
    v20 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
